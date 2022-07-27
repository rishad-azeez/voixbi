package store

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/grafana/grafana/pkg/api/response"
	"github.com/grafana/grafana/pkg/api/routing"
	"github.com/grafana/grafana/pkg/middleware"
	"github.com/grafana/grafana/pkg/models"
	ac "github.com/grafana/grafana/pkg/services/accesscontrol"
	"github.com/grafana/grafana/pkg/util"
	"github.com/grafana/grafana/pkg/web"
)

func UploadErrorToStatusCode(err error) int {
	switch {
	case errors.Is(err, ErrStorageNotFound):
		return 404

	case errors.Is(err, ErrUnsupportedStorage):
		return 400

	case errors.Is(err, ErrValidationFailed):
		return 400

	case errors.Is(err, ErrFileAlreadyExists):
		return 400

	case errors.Is(err, ErrAccessDenied):
		return 403

	default:
		return 500
	}
}

func (s *standardStorageService) RegisterHTTPRoutes(storageRoute routing.RouteRegister) {
	storageRoute.Get("/list/", routing.Wrap(s.list))
	storageRoute.Get("/list/*", routing.Wrap(s.list))
	storageRoute.Get("/read/*", routing.Wrap(s.read))

	// Write paths
	reqGrafanaAdmin := middleware.ReqGrafanaAdmin
	storageRoute.Post("/delete/*", reqGrafanaAdmin, routing.Wrap(s.delete))
	storageRoute.Post("/upload", reqGrafanaAdmin, routing.Wrap(s.upload))
	storageRoute.Post("/createFolder", reqGrafanaAdmin, routing.Wrap(s.createFolder))
	storageRoute.Post("/deleteFolder", reqGrafanaAdmin, routing.Wrap(s.deleteFolder))
	storageRoute.Get("/config", reqGrafanaAdmin, routing.Wrap(s.getConfig))
}

func (s *standardStorageService) upload(c *models.ReqContext) response.Response {
	// assumes we are only uploading to the SQL database - TODO: refactor once we introduce object stores
	quotaReached, err := s.quotaService.CheckQuotaReached(c.Req.Context(), "file", nil)
	if err != nil {
		return response.Error(500, "Internal server error", err)
	}

	if quotaReached {
		return response.Error(400, "File quota reached", errors.New("file quota reached"))
	}

	type rspInfo struct {
		Message string `json:"message,omitempty"`
		Path    string `json:"path,omitempty"`
		Count   int    `json:"count,omitempty"`
		Bytes   int    `json:"bytes,omitempty"`
		Error   bool   `json:"err,omitempty"`
	}
	rsp := &rspInfo{Message: "uploaded"}

	c.Req.Body = http.MaxBytesReader(c.Resp, c.Req.Body, MAX_UPLOAD_SIZE)
	if err := c.Req.ParseMultipartForm(MAX_UPLOAD_SIZE); err != nil {
		rsp.Message = fmt.Sprintf("Please limit file uploaded under %s", util.ByteCountSI(MAX_UPLOAD_SIZE))
		rsp.Error = true
		return response.JSON(400, rsp)
	}
	message := getMultipartFormValue(c.Req, "message")
	overwriteExistingFile := getMultipartFormValue(c.Req, "overwriteExistingFile") != "false" // must explicitly overwrite
	folder := getMultipartFormValue(c.Req, "folder")

	for k, fileHeaders := range c.Req.MultipartForm.File {
		path := getMultipartFormValue(c.Req, k+".path") // match the path with a file
		if len(fileHeaders) > 1 {
			path = ""
		}
		if path == "" && folder == "" {
			rsp.Message = "please specify the upload folder or full path"
			rsp.Error = true
			return response.JSON(400, rsp)
		}

		for _, fileHeader := range fileHeaders {
			// restrict file size based on file size
			// open each file to copy contents
			file, err := fileHeader.Open()
			if err != nil {
				return response.Error(500, "Internal Server Error", err)
			}
			err = file.Close()
			if err != nil {
				return response.Error(500, "Internal Server Error", err)
			}
			data, err := ioutil.ReadAll(file)
			if err != nil {
				return response.Error(500, "Internal Server Error", err)
			}

			if path == "" {
				path = folder + "/" + fileHeader.Filename
			}

			entityType := EntityTypeJSON
			mimeType := http.DetectContentType(data)
			if strings.HasPrefix(mimeType, "image") || strings.HasSuffix(path, ".svg") {
				entityType = EntityTypeImage
			}

			err = s.Upload(c.Req.Context(), c.SignedInUser, &UploadRequest{
				Contents:              data,
				EntityType:            entityType,
				Path:                  path,
				OverwriteExistingFile: overwriteExistingFile,
				Properties: map[string]string{
					"message": message, // the commit/changelog entry
				},
			})

			if err != nil {
				return response.Error(UploadErrorToStatusCode(err), err.Error(), err)
			}
			rsp.Count++
			rsp.Bytes += len(data)
			rsp.Path = path
		}
	}

	return response.JSON(200, rsp)
}

func getMultipartFormValue(req *http.Request, key string) string {
	v, ok := req.MultipartForm.Value[key]
	if !ok || len(v) != 1 {
		return ""
	}
	return v[0]
}

func (s *standardStorageService) read(c *models.ReqContext) response.Response {
	// full path is api/storage/read/upload/example.jpg, but we only want the part after read
	scope, path := getPathAndScope(c)
	file, err := s.Read(c.Req.Context(), c.SignedInUser, scope+"/"+path)
	if err != nil {
		return response.Error(400, "cannot call read", err)
	}

	if file == nil || file.Contents == nil {
		return response.Error(404, "file does not exist", err)
	}

	// set the correct content type for svg
	if strings.HasSuffix(path, ".svg") {
		c.Resp.Header().Set("Content-Type", "image/svg+xml")
	}
	return response.Respond(200, file.Contents)
}

func (s *standardStorageService) delete(c *models.ReqContext) response.Response {
	// full path is api/storage/delete/upload/example.jpg, but we only want the part after upload
	scope, path := getPathAndScope(c)

	err := s.Delete(c.Req.Context(), c.SignedInUser, scope+"/"+path)
	if err != nil {
		return response.Error(400, "failed to delete the file: "+err.Error(), err)
	}
	return response.JSON(200, map[string]interface{}{
		"message": "Removed file from storage",
		"success": true,
		"path":    path,
	})
}

func (s *standardStorageService) deleteFolder(c *models.ReqContext) response.Response {
	body, err := io.ReadAll(c.Req.Body)
	if err != nil {
		return response.Error(500, "error reading bytes", err)
	}

	cmd := &DeleteFolderCmd{}
	err = json.Unmarshal(body, cmd)
	if err != nil {
		return response.Error(400, "error parsing body", err)
	}

	if cmd.Path == "" {
		return response.Error(400, "empty path", err)
	}

	// full path is api/storage/delete/upload/example.jpg, but we only want the part after upload
	_, path := getPathAndScope(c)
	if err := s.DeleteFolder(c.Req.Context(), c.SignedInUser, cmd); err != nil {
		return response.Error(400, "failed to delete the folder: "+err.Error(), err)
	}

	return response.JSON(200, map[string]interface{}{
		"message": "Removed folder from storage",
		"success": true,
		"path":    path,
	})
}

func (s *standardStorageService) createFolder(c *models.ReqContext) response.Response {
	body, err := io.ReadAll(c.Req.Body)
	if err != nil {
		return response.Error(500, "error reading bytes", err)
	}

	cmd := &CreateFolderCmd{}
	err = json.Unmarshal(body, cmd)
	if err != nil {
		return response.Error(400, "error parsing body", err)
	}

	if cmd.Path == "" {
		return response.Error(400, "empty path", err)
	}

	if err := s.CreateFolder(c.Req.Context(), c.SignedInUser, cmd); err != nil {
		return response.Error(400, "failed to create the folder: "+err.Error(), err)
	}

	return response.JSON(200, map[string]interface{}{
		"message": "Folder created",
		"success": true,
		"path":    cmd.Path,
	})
}

func (s *standardStorageService) list(c *models.ReqContext) response.Response {
	params := web.Params(c.Req)
	path := params["*"]
	frame, err := s.List(c.Req.Context(), c.SignedInUser, path)
	if err != nil {
		return response.Error(400, "error reading path", err)
	}
	if frame == nil {
		return response.Error(404, "not found", nil)
	}
	return response.JSONStreaming(http.StatusOK, frame)
}

func (s *standardStorageService) getConfig(c *models.ReqContext) response.Response {
	roots := make([]RootStorageMeta, 0)
	orgId := c.OrgId
	t := s.tree
	t.assureOrgIsInitialized(orgId)
	for _, f := range t.rootsByOrgId[ac.GlobalOrgID] {
		roots = append(roots, f.Meta())
	}
	if orgId != ac.GlobalOrgID {
		for _, f := range t.rootsByOrgId[orgId] {
			roots = append(roots, f.Meta())
		}
	}
	return response.JSON(200, roots)
}
