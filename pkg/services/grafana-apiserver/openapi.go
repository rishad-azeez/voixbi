package grafanaapiserver

import (
	"golang.org/x/exp/maps"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	common "k8s.io/kube-openapi/pkg/common"
	spec "k8s.io/kube-openapi/pkg/validation/spec"

	"github.com/grafana/grafana/pkg/apis"
)

// This should eventually live in grafana-app-sdk
func getOpenAPIDefinitions(builders []apis.APIGroupBuilder) common.GetOpenAPIDefinitions {
	return func(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
		defs := getStandardOpenAPIDefinitions(ref)
		for _, builder := range builders {
			g := builder.GetOpenAPIDefinitions()
			if g != nil {
				out := g(ref)
				maps.Copy(defs, out)
			}
		}
		return defs
	}
}

func getStandardOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"k8s.io/apimachinery/pkg/apis/meta/v1.APIGroup":                  schema_pkg_apis_meta_v1_APIGroup(ref),
		"k8s.io/apimachinery/pkg/apis/meta/v1.APIGroupList":              schema_pkg_apis_meta_v1_APIGroupList(ref),
		"k8s.io/apimachinery/pkg/apis/meta/v1.APIResource":               schema_pkg_apis_meta_v1_APIResource(ref),
		"k8s.io/apimachinery/pkg/apis/meta/v1.APIResourceList":           schema_pkg_apis_meta_v1_APIResourceList(ref),
		"k8s.io/apimachinery/pkg/apis/meta/v1.APIVersions":               schema_pkg_apis_meta_v1_APIVersions(ref),
		"k8s.io/apimachinery/pkg/apis/meta/v1.ApplyOptions":              schema_pkg_apis_meta_v1_ApplyOptions(ref),
		"k8s.io/apimachinery/pkg/apis/meta/v1.Condition":                 schema_pkg_apis_meta_v1_Condition(ref),
		"k8s.io/apimachinery/pkg/apis/meta/v1.CreateOptions":             schema_pkg_apis_meta_v1_CreateOptions(ref),
		"k8s.io/apimachinery/pkg/apis/meta/v1.DeleteOptions":             schema_pkg_apis_meta_v1_DeleteOptions(ref),
		"k8s.io/apimachinery/pkg/apis/meta/v1.Duration":                  schema_pkg_apis_meta_v1_Duration(ref),
		"k8s.io/apimachinery/pkg/apis/meta/v1.FieldsV1":                  schema_pkg_apis_meta_v1_FieldsV1(ref),
		"k8s.io/apimachinery/pkg/apis/meta/v1.GetOptions":                schema_pkg_apis_meta_v1_GetOptions(ref),
		"k8s.io/apimachinery/pkg/apis/meta/v1.GroupKind":                 schema_pkg_apis_meta_v1_GroupKind(ref),
		"k8s.io/apimachinery/pkg/apis/meta/v1.GroupResource":             schema_pkg_apis_meta_v1_GroupResource(ref),
		"k8s.io/apimachinery/pkg/apis/meta/v1.GroupVersion":              schema_pkg_apis_meta_v1_GroupVersion(ref),
		"k8s.io/apimachinery/pkg/apis/meta/v1.GroupVersionForDiscovery":  schema_pkg_apis_meta_v1_GroupVersionForDiscovery(ref),
		"k8s.io/apimachinery/pkg/apis/meta/v1.GroupVersionKind":          schema_pkg_apis_meta_v1_GroupVersionKind(ref),
		"k8s.io/apimachinery/pkg/apis/meta/v1.GroupVersionResource":      schema_pkg_apis_meta_v1_GroupVersionResource(ref),
		"k8s.io/apimachinery/pkg/apis/meta/v1.InternalEvent":             schema_pkg_apis_meta_v1_InternalEvent(ref),
		"k8s.io/apimachinery/pkg/apis/meta/v1.LabelSelector":             schema_pkg_apis_meta_v1_LabelSelector(ref),
		"k8s.io/apimachinery/pkg/apis/meta/v1.LabelSelectorRequirement":  schema_pkg_apis_meta_v1_LabelSelectorRequirement(ref),
		"k8s.io/apimachinery/pkg/apis/meta/v1.List":                      schema_pkg_apis_meta_v1_List(ref),
		"k8s.io/apimachinery/pkg/apis/meta/v1.ListMeta":                  schema_pkg_apis_meta_v1_ListMeta(ref),
		"k8s.io/apimachinery/pkg/apis/meta/v1.ListOptions":               schema_pkg_apis_meta_v1_ListOptions(ref),
		"k8s.io/apimachinery/pkg/apis/meta/v1.ManagedFieldsEntry":        schema_pkg_apis_meta_v1_ManagedFieldsEntry(ref),
		"k8s.io/apimachinery/pkg/apis/meta/v1.MicroTime":                 schema_pkg_apis_meta_v1_MicroTime(ref),
		"k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta":                schema_pkg_apis_meta_v1_ObjectMeta(ref),
		"k8s.io/apimachinery/pkg/apis/meta/v1.OwnerReference":            schema_pkg_apis_meta_v1_OwnerReference(ref),
		"k8s.io/apimachinery/pkg/apis/meta/v1.PartialObjectMetadata":     schema_pkg_apis_meta_v1_PartialObjectMetadata(ref),
		"k8s.io/apimachinery/pkg/apis/meta/v1.PartialObjectMetadataList": schema_pkg_apis_meta_v1_PartialObjectMetadataList(ref),
		"k8s.io/apimachinery/pkg/apis/meta/v1.Patch":                     schema_pkg_apis_meta_v1_Patch(ref),
		"k8s.io/apimachinery/pkg/apis/meta/v1.PatchOptions":              schema_pkg_apis_meta_v1_PatchOptions(ref),
		"k8s.io/apimachinery/pkg/apis/meta/v1.Preconditions":             schema_pkg_apis_meta_v1_Preconditions(ref),
		"k8s.io/apimachinery/pkg/apis/meta/v1.RootPaths":                 schema_pkg_apis_meta_v1_RootPaths(ref),
		"k8s.io/apimachinery/pkg/apis/meta/v1.ServerAddressByClientCIDR": schema_pkg_apis_meta_v1_ServerAddressByClientCIDR(ref),
		"k8s.io/apimachinery/pkg/apis/meta/v1.Status":                    schema_pkg_apis_meta_v1_Status(ref),
		"k8s.io/apimachinery/pkg/apis/meta/v1.StatusCause":               schema_pkg_apis_meta_v1_StatusCause(ref),
		"k8s.io/apimachinery/pkg/apis/meta/v1.StatusDetails":             schema_pkg_apis_meta_v1_StatusDetails(ref),
		"k8s.io/apimachinery/pkg/apis/meta/v1.Table":                     schema_pkg_apis_meta_v1_Table(ref),
		"k8s.io/apimachinery/pkg/apis/meta/v1.TableColumnDefinition":     schema_pkg_apis_meta_v1_TableColumnDefinition(ref),
		"k8s.io/apimachinery/pkg/apis/meta/v1.TableOptions":              schema_pkg_apis_meta_v1_TableOptions(ref),
		"k8s.io/apimachinery/pkg/apis/meta/v1.TableRow":                  schema_pkg_apis_meta_v1_TableRow(ref),
		"k8s.io/apimachinery/pkg/apis/meta/v1.TableRowCondition":         schema_pkg_apis_meta_v1_TableRowCondition(ref),
		"k8s.io/apimachinery/pkg/apis/meta/v1.Time":                      schema_pkg_apis_meta_v1_Time(ref),
		"k8s.io/apimachinery/pkg/apis/meta/v1.Timestamp":                 schema_pkg_apis_meta_v1_Timestamp(ref),
		"k8s.io/apimachinery/pkg/apis/meta/v1.TypeMeta":                  schema_pkg_apis_meta_v1_TypeMeta(ref),
		"k8s.io/apimachinery/pkg/apis/meta/v1.UpdateOptions":             schema_pkg_apis_meta_v1_UpdateOptions(ref),
		"k8s.io/apimachinery/pkg/apis/meta/v1.WatchEvent":                schema_pkg_apis_meta_v1_WatchEvent(ref),
		"k8s.io/apimachinery/pkg/runtime.RawExtension":                   schema_k8sio_apimachinery_pkg_runtime_RawExtension(ref),
		"k8s.io/apimachinery/pkg/runtime.TypeMeta":                       schema_k8sio_apimachinery_pkg_runtime_TypeMeta(ref),
		"k8s.io/apimachinery/pkg/runtime.Unknown":                        schema_k8sio_apimachinery_pkg_runtime_Unknown(ref),
		"k8s.io/apimachinery/pkg/version.Info":                           schema_k8sio_apimachinery_pkg_version_Info(ref),
	}
}

func schema_pkg_apis_meta_v1_APIGroup(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "APIGroup contains the name, the supported versions, and the preferred version of a group.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"name": {
						SchemaProps: spec.SchemaProps{
							Description: "name is the name of the group.",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"versions": {
						SchemaProps: spec.SchemaProps{
							Description: "versions are the versions supported in this group.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: map[string]interface{}{},
										Ref:     ref("k8s.io/apimachinery/pkg/apis/meta/v1.GroupVersionForDiscovery"),
									},
								},
							},
						},
					},
					"preferredVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "preferredVersion is the version preferred by the API server, which probably is the storage version.",
							Default:     map[string]interface{}{},
							Ref:         ref("k8s.io/apimachinery/pkg/apis/meta/v1.GroupVersionForDiscovery"),
						},
					},
					"serverAddressByClientCIDRs": {
						SchemaProps: spec.SchemaProps{
							Description: "a map of client CIDR to server address that is serving this group. This is to help clients reach servers in the most network-efficient way possible. Clients can use the appropriate server address as per the CIDR that they match. In case of multiple matches, clients should use the longest matching CIDR. The server returns only those CIDRs that it thinks that the client can match. For example: the master will return an internal IP CIDR only, if the client reaches the server using an internal IP. Server looks at X-Forwarded-For header or X-Real-Ip header or request.RemoteAddr (in that order) to get the client IP.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: map[string]interface{}{},
										Ref:     ref("k8s.io/apimachinery/pkg/apis/meta/v1.ServerAddressByClientCIDR"),
									},
								},
							},
						},
					},
				},
				Required: []string{"name", "versions"},
			},
		},
		Dependencies: []string{
			"k8s.io/apimachinery/pkg/apis/meta/v1.GroupVersionForDiscovery", "k8s.io/apimachinery/pkg/apis/meta/v1.ServerAddressByClientCIDR"},
	}
}

func schema_pkg_apis_meta_v1_APIGroupList(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "APIGroupList is a list of APIGroup, to allow clients to discover the API at /apis.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"groups": {
						SchemaProps: spec.SchemaProps{
							Description: "groups is a list of APIGroup.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: map[string]interface{}{},
										Ref:     ref("k8s.io/apimachinery/pkg/apis/meta/v1.APIGroup"),
									},
								},
							},
						},
					},
				},
				Required: []string{"groups"},
			},
		},
		Dependencies: []string{
			"k8s.io/apimachinery/pkg/apis/meta/v1.APIGroup"},
	}
}

func schema_pkg_apis_meta_v1_APIResource(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "APIResource specifies the name of a resource and whether it is namespaced.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"name": {
						SchemaProps: spec.SchemaProps{
							Description: "name is the plural name of the resource.",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"singularName": {
						SchemaProps: spec.SchemaProps{
							Description: "singularName is the singular name of the resource.  This allows clients to handle plural and singular opaquely. The singularName is more correct for reporting status on a single item and both singular and plural are allowed from the kubectl CLI interface.",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"namespaced": {
						SchemaProps: spec.SchemaProps{
							Description: "namespaced indicates if a resource is namespaced or not.",
							Default:     false,
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
					"group": {
						SchemaProps: spec.SchemaProps{
							Description: "group is the preferred group of the resource.  Empty implies the group of the containing resource list. For subresources, this may have a different value, for example: Scale\".",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"version": {
						SchemaProps: spec.SchemaProps{
							Description: "version is the preferred version of the resource.  Empty implies the version of the containing resource list For subresources, this may have a different value, for example: v1 (while inside a v1beta1 version of the core resource's group)\".",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "kind is the kind for the resource (e.g. 'Foo' is the kind for a resource 'foo')",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"verbs": {
						SchemaProps: spec.SchemaProps{
							Description: "verbs is a list of supported kube verbs (this includes get, list, watch, create, update, patch, delete, deletecollection, and proxy)",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: "",
										Type:    []string{"string"},
										Format:  "",
									},
								},
							},
						},
					},
					"shortNames": {
						SchemaProps: spec.SchemaProps{
							Description: "shortNames is a list of suggested short names of the resource.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: "",
										Type:    []string{"string"},
										Format:  "",
									},
								},
							},
						},
					},
					"categories": {
						SchemaProps: spec.SchemaProps{
							Description: "categories is a list of the grouped resources this resource belongs to (e.g. 'all')",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: "",
										Type:    []string{"string"},
										Format:  "",
									},
								},
							},
						},
					},
					"storageVersionHash": {
						SchemaProps: spec.SchemaProps{
							Description: "The hash value of the storage version, the version this resource is converted to when written to the data store. Value must be treated as opaque by clients. Only equality comparison on the value is valid. This is an alpha feature and may change or be removed in the future. The field is populated by the apiserver only if the StorageVersionHash feature gate is enabled. This field will remain optional even if it graduates.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{"name", "singularName", "namespaced", "kind", "verbs"},
			},
		},
	}
}

func schema_pkg_apis_meta_v1_APIResourceList(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "APIResourceList is a list of APIResource, it is used to expose the name of the resources supported in a specific group and version, and if the resource is namespaced.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"groupVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "groupVersion is the group and version this APIResourceList is for.",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"resources": {
						SchemaProps: spec.SchemaProps{
							Description: "resources contains the name of the resources and if they are namespaced.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: map[string]interface{}{},
										Ref:     ref("k8s.io/apimachinery/pkg/apis/meta/v1.APIResource"),
									},
								},
							},
						},
					},
				},
				Required: []string{"groupVersion", "resources"},
			},
		},
		Dependencies: []string{
			"k8s.io/apimachinery/pkg/apis/meta/v1.APIResource"},
	}
}

func schema_pkg_apis_meta_v1_APIVersions(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "APIVersions lists the versions that are available, to allow clients to discover the API at /api, which is the root path of the legacy v1 API.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"versions": {
						SchemaProps: spec.SchemaProps{
							Description: "versions are the api versions that are available.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: "",
										Type:    []string{"string"},
										Format:  "",
									},
								},
							},
						},
					},
					"serverAddressByClientCIDRs": {
						SchemaProps: spec.SchemaProps{
							Description: "a map of client CIDR to server address that is serving this group. This is to help clients reach servers in the most network-efficient way possible. Clients can use the appropriate server address as per the CIDR that they match. In case of multiple matches, clients should use the longest matching CIDR. The server returns only those CIDRs that it thinks that the client can match. For example: the master will return an internal IP CIDR only, if the client reaches the server using an internal IP. Server looks at X-Forwarded-For header or X-Real-Ip header or request.RemoteAddr (in that order) to get the client IP.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: map[string]interface{}{},
										Ref:     ref("k8s.io/apimachinery/pkg/apis/meta/v1.ServerAddressByClientCIDR"),
									},
								},
							},
						},
					},
				},
				Required: []string{"versions", "serverAddressByClientCIDRs"},
			},
		},
		Dependencies: []string{
			"k8s.io/apimachinery/pkg/apis/meta/v1.ServerAddressByClientCIDR"},
	}
}

func schema_pkg_apis_meta_v1_ApplyOptions(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ApplyOptions may be provided when applying an API object. FieldManager is required for apply requests. ApplyOptions is equivalent to PatchOptions. It is provided as a convenience with documentation that speaks specifically to how the options fields relate to apply.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"dryRun": {
						SchemaProps: spec.SchemaProps{
							Description: "When present, indicates that modifications should not be persisted. An invalid or unrecognized dryRun directive will result in an error response and no further processing of the request. Valid values are: - All: all dry run stages will be processed",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: "",
										Type:    []string{"string"},
										Format:  "",
									},
								},
							},
						},
					},
					"force": {
						SchemaProps: spec.SchemaProps{
							Description: "Force is going to \"force\" Apply requests. It means user will re-acquire conflicting fields owned by other people.",
							Default:     false,
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
					"fieldManager": {
						SchemaProps: spec.SchemaProps{
							Description: "fieldManager is a name associated with the actor or entity that is making these changes. The value must be less than or 128 characters long, and only contain printable characters, as defined by https://golang.org/pkg/unicode/#IsPrint. This field is required.",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{"force", "fieldManager"},
			},
		},
	}
}

func schema_pkg_apis_meta_v1_Condition(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Condition contains details for one aspect of the current state of this API Resource.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"type": {
						SchemaProps: spec.SchemaProps{
							Description: "type of condition in CamelCase or in foo.example.com/CamelCase.",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Description: "status of the condition, one of True, False, Unknown.",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"observedGeneration": {
						SchemaProps: spec.SchemaProps{
							Description: "observedGeneration represents the .metadata.generation that the condition was set based upon. For instance, if .metadata.generation is currently 12, but the .status.conditions[x].observedGeneration is 9, the condition is out of date with respect to the current state of the instance.",
							Type:        []string{"integer"},
							Format:      "int64",
						},
					},
					"lastTransitionTime": {
						SchemaProps: spec.SchemaProps{
							Description: "lastTransitionTime is the last time the condition transitioned from one status to another. This should be when the underlying condition changed.  If that is not known, then using the time when the API field changed is acceptable.",
							Default:     map[string]interface{}{},
							Ref:         ref("k8s.io/apimachinery/pkg/apis/meta/v1.Time"),
						},
					},
					"reason": {
						SchemaProps: spec.SchemaProps{
							Description: "reason contains a programmatic identifier indicating the reason for the condition's last transition. Producers of specific condition types may define expected values and meanings for this field, and whether the values are considered a guaranteed API. The value should be a CamelCase string. This field may not be empty.",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"message": {
						SchemaProps: spec.SchemaProps{
							Description: "message is a human readable message indicating details about the transition. This may be an empty string.",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{"type", "status", "lastTransitionTime", "reason", "message"},
			},
		},
		Dependencies: []string{
			"k8s.io/apimachinery/pkg/apis/meta/v1.Time"},
	}
}

func schema_pkg_apis_meta_v1_CreateOptions(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "CreateOptions may be provided when creating an API object.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"dryRun": {
						SchemaProps: spec.SchemaProps{
							Description: "When present, indicates that modifications should not be persisted. An invalid or unrecognized dryRun directive will result in an error response and no further processing of the request. Valid values are: - All: all dry run stages will be processed",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: "",
										Type:    []string{"string"},
										Format:  "",
									},
								},
							},
						},
					},
					"fieldManager": {
						SchemaProps: spec.SchemaProps{
							Description: "fieldManager is a name associated with the actor or entity that is making these changes. The value must be less than or 128 characters long, and only contain printable characters, as defined by https://golang.org/pkg/unicode/#IsPrint.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"fieldValidation": {
						SchemaProps: spec.SchemaProps{
							Description: "fieldValidation instructs the server on how to handle objects in the request (POST/PUT/PATCH) containing unknown or duplicate fields. Valid values are: - Ignore: This will ignore any unknown fields that are silently dropped from the object, and will ignore all but the last duplicate field that the decoder encounters. This is the default behavior prior to v1.23. - Warn: This will send a warning via the standard warning response header for each unknown field that is dropped from the object, and for each duplicate field that is encountered. The request will still succeed if there are no other errors, and will only persist the last of any duplicate fields. This is the default in v1.23+ - Strict: This will fail the request with a BadRequest error if any unknown fields would be dropped from the object, or if any duplicate fields are present. The error returned from the server will contain all unknown and duplicate fields encountered.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
			},
		},
	}
}

func schema_pkg_apis_meta_v1_DeleteOptions(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "DeleteOptions may be provided when deleting an API object.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"gracePeriodSeconds": {
						SchemaProps: spec.SchemaProps{
							Description: "The duration in seconds before the object should be deleted. Value must be non-negative integer. The value zero indicates delete immediately. If this value is nil, the default grace period for the specified type will be used. Defaults to a per object value if not specified. zero means delete immediately.",
							Type:        []string{"integer"},
							Format:      "int64",
						},
					},
					"preconditions": {
						SchemaProps: spec.SchemaProps{
							Description: "Must be fulfilled before a deletion is carried out. If not possible, a 409 Conflict status will be returned.",
							Ref:         ref("k8s.io/apimachinery/pkg/apis/meta/v1.Preconditions"),
						},
					},
					"orphanDependents": {
						SchemaProps: spec.SchemaProps{
							Description: "Deprecated: please use the PropagationPolicy, this field will be deprecated in 1.7. Should the dependent objects be orphaned. If true/false, the \"orphan\" finalizer will be added to/removed from the object's finalizers list. Either this field or PropagationPolicy may be set, but not both.",
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
					"propagationPolicy": {
						SchemaProps: spec.SchemaProps{
							Description: "Whether and how garbage collection will be performed. Either this field or OrphanDependents may be set, but not both. The default policy is decided by the existing finalizer set in the metadata.finalizers and the resource-specific default policy. Acceptable values are: 'Orphan' - orphan the dependents; 'Background' - allow the garbage collector to delete the dependents in the background; 'Foreground' - a cascading policy that deletes all dependents in the foreground.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"dryRun": {
						SchemaProps: spec.SchemaProps{
							Description: "When present, indicates that modifications should not be persisted. An invalid or unrecognized dryRun directive will result in an error response and no further processing of the request. Valid values are: - All: all dry run stages will be processed",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: "",
										Type:    []string{"string"},
										Format:  "",
									},
								},
							},
						},
					},
				},
			},
		},
		Dependencies: []string{
			"k8s.io/apimachinery/pkg/apis/meta/v1.Preconditions"},
	}
}

func schema_pkg_apis_meta_v1_Duration(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Duration is a wrapper around time.Duration which supports correct marshaling to YAML and JSON. In particular, it marshals into strings, which can be used as map keys in json.",
				Type:        v1.Duration{}.OpenAPISchemaType(),
				Format:      v1.Duration{}.OpenAPISchemaFormat(),
			},
		},
	}
}

func schema_pkg_apis_meta_v1_FieldsV1(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "FieldsV1 stores a set of fields in a data structure like a Trie, in JSON format.\n\nEach key is either a '.' representing the field itself, and will always map to an empty set, or a string representing a sub-field or item. The string will follow one of these four formats: 'f:<name>', where <name> is the name of a field in a struct, or key in a map 'v:<value>', where <value> is the exact json formatted value of a list item 'i:<index>', where <index> is position of a item in a list 'k:<keys>', where <keys> is a map of  a list item's key fields to their unique values If a key maps to an empty Fields value, the field that key represents is part of the set.\n\nThe exact format is defined in sigs.k8s.io/structured-merge-diff",
				Type:        []string{"object"},
			},
		},
	}
}

func schema_pkg_apis_meta_v1_GetOptions(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "GetOptions is the standard query options to the standard REST get call.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"resourceVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "resourceVersion sets a constraint on what resource versions a request may be served from. See https://kubernetes.io/docs/reference/using-api/api-concepts/#resource-versions for details.\n\nDefaults to unset",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
			},
		},
	}
}

func schema_pkg_apis_meta_v1_GroupKind(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "GroupKind specifies a Group and a Kind, but does not force a version.  This is useful for identifying concepts during lookup stages without having partially valid types",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"group": {
						SchemaProps: spec.SchemaProps{
							Default: "",
							Type:    []string{"string"},
							Format:  "",
						},
					},
					"kind": {
						SchemaProps: spec.SchemaProps{
							Default: "",
							Type:    []string{"string"},
							Format:  "",
						},
					},
				},
				Required: []string{"group", "kind"},
			},
		},
	}
}

func schema_pkg_apis_meta_v1_GroupResource(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "GroupResource specifies a Group and a Resource, but does not force a version.  This is useful for identifying concepts during lookup stages without having partially valid types",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"group": {
						SchemaProps: spec.SchemaProps{
							Default: "",
							Type:    []string{"string"},
							Format:  "",
						},
					},
					"resource": {
						SchemaProps: spec.SchemaProps{
							Default: "",
							Type:    []string{"string"},
							Format:  "",
						},
					},
				},
				Required: []string{"group", "resource"},
			},
		},
	}
}

func schema_pkg_apis_meta_v1_GroupVersion(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "GroupVersion contains the \"group\" and the \"version\", which uniquely identifies the API.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"group": {
						SchemaProps: spec.SchemaProps{
							Default: "",
							Type:    []string{"string"},
							Format:  "",
						},
					},
					"version": {
						SchemaProps: spec.SchemaProps{
							Default: "",
							Type:    []string{"string"},
							Format:  "",
						},
					},
				},
				Required: []string{"group", "version"},
			},
		},
	}
}

func schema_pkg_apis_meta_v1_GroupVersionForDiscovery(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "GroupVersion contains the \"group/version\" and \"version\" string of a version. It is made a struct to keep extensibility.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"groupVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "groupVersion specifies the API group and version in the form \"group/version\"",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"version": {
						SchemaProps: spec.SchemaProps{
							Description: "version specifies the version in the form of \"version\". This is to save the clients the trouble of splitting the GroupVersion.",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{"groupVersion", "version"},
			},
		},
	}
}

func schema_pkg_apis_meta_v1_GroupVersionKind(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "GroupVersionKind unambiguously identifies a kind.  It doesn't anonymously include GroupVersion to avoid automatic coercion.  It doesn't use a GroupVersion to avoid custom marshalling",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"group": {
						SchemaProps: spec.SchemaProps{
							Default: "",
							Type:    []string{"string"},
							Format:  "",
						},
					},
					"version": {
						SchemaProps: spec.SchemaProps{
							Default: "",
							Type:    []string{"string"},
							Format:  "",
						},
					},
					"kind": {
						SchemaProps: spec.SchemaProps{
							Default: "",
							Type:    []string{"string"},
							Format:  "",
						},
					},
				},
				Required: []string{"group", "version", "kind"},
			},
		},
	}
}

func schema_pkg_apis_meta_v1_GroupVersionResource(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "GroupVersionResource unambiguously identifies a resource.  It doesn't anonymously include GroupVersion to avoid automatic coercion.  It doesn't use a GroupVersion to avoid custom marshalling",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"group": {
						SchemaProps: spec.SchemaProps{
							Default: "",
							Type:    []string{"string"},
							Format:  "",
						},
					},
					"version": {
						SchemaProps: spec.SchemaProps{
							Default: "",
							Type:    []string{"string"},
							Format:  "",
						},
					},
					"resource": {
						SchemaProps: spec.SchemaProps{
							Default: "",
							Type:    []string{"string"},
							Format:  "",
						},
					},
				},
				Required: []string{"group", "version", "resource"},
			},
		},
	}
}

func schema_pkg_apis_meta_v1_InternalEvent(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "InternalEvent makes watch.Event versioned",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"Type": {
						SchemaProps: spec.SchemaProps{
							Default: "",
							Type:    []string{"string"},
							Format:  "",
						},
					},
					"Object": {
						SchemaProps: spec.SchemaProps{
							Description: "Object is:\n * If Type is Added or Modified: the new state of the object.\n * If Type is Deleted: the state of the object immediately before deletion.\n * If Type is Bookmark: the object (instance of a type being watched) where\n   only ResourceVersion field is set. On successful restart of watch from a\n   bookmark resourceVersion, client is guaranteed to not get repeat event\n   nor miss any events.\n * If Type is Error: *api.Status is recommended; other types may make sense\n   depending on context.",
							Ref:         ref("k8s.io/apimachinery/pkg/runtime.Object"),
						},
					},
				},
				Required: []string{"Type", "Object"},
			},
		},
		Dependencies: []string{
			"k8s.io/apimachinery/pkg/runtime.Object"},
	}
}

func schema_pkg_apis_meta_v1_LabelSelector(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "A label selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty label selector matches all objects. A null label selector matches no objects.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"matchLabels": {
						SchemaProps: spec.SchemaProps{
							Description: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is \"key\", the operator is \"In\", and the values array contains only \"value\". The requirements are ANDed.",
							Type:        []string{"object"},
							AdditionalProperties: &spec.SchemaOrBool{
								Allows: true,
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: "",
										Type:    []string{"string"},
										Format:  "",
									},
								},
							},
						},
					},
					"matchExpressions": {
						SchemaProps: spec.SchemaProps{
							Description: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: map[string]interface{}{},
										Ref:     ref("k8s.io/apimachinery/pkg/apis/meta/v1.LabelSelectorRequirement"),
									},
								},
							},
						},
					},
				},
			},
			VendorExtensible: spec.VendorExtensible{
				Extensions: spec.Extensions{
					"x-kubernetes-map-type": "atomic",
				},
			},
		},
		Dependencies: []string{
			"k8s.io/apimachinery/pkg/apis/meta/v1.LabelSelectorRequirement"},
	}
}

func schema_pkg_apis_meta_v1_LabelSelectorRequirement(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "A label selector requirement is a selector that contains values, a key, and an operator that relates the key and values.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"key": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-patch-merge-key": "key",
								"x-kubernetes-patch-strategy":  "merge",
							},
						},
						SchemaProps: spec.SchemaProps{
							Description: "key is the label key that the selector applies to.",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"operator": {
						SchemaProps: spec.SchemaProps{
							Description: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"values": {
						SchemaProps: spec.SchemaProps{
							Description: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: "",
										Type:    []string{"string"},
										Format:  "",
									},
								},
							},
						},
					},
				},
				Required: []string{"key", "operator"},
			},
		},
	}
}

func schema_pkg_apis_meta_v1_List(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "List holds a list of objects, which may not be known by the server.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Description: "Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Default:     map[string]interface{}{},
							Ref:         ref("k8s.io/apimachinery/pkg/apis/meta/v1.ListMeta"),
						},
					},
					"items": {
						SchemaProps: spec.SchemaProps{
							Description: "List of objects",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: map[string]interface{}{},
										Ref:     ref("k8s.io/apimachinery/pkg/runtime.RawExtension"),
									},
								},
							},
						},
					},
				},
				Required: []string{"items"},
			},
		},
		Dependencies: []string{
			"k8s.io/apimachinery/pkg/apis/meta/v1.ListMeta", "k8s.io/apimachinery/pkg/runtime.RawExtension"},
	}
}

func schema_pkg_apis_meta_v1_ListMeta(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"selfLink": {
						SchemaProps: spec.SchemaProps{
							Description: "Deprecated: selfLink is a legacy read-only field that is no longer populated by the system.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"resourceVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "String that identifies the server's internal version of this object that can be used by clients to determine when objects have changed. Value must be treated as opaque by clients and passed unmodified back to the server. Populated by the system. Read-only. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"continue": {
						SchemaProps: spec.SchemaProps{
							Description: "continue may be set if the user set a limit on the number of items returned, and indicates that the server has more data available. The value is opaque and may be used to issue another request to the endpoint that served this list to retrieve the next set of available objects. Continuing a consistent list may not be possible if the server configuration has changed or more than a few minutes have passed. The resourceVersion field returned when using this continue value will be identical to the value in the first response, unless you have received this token from an error message.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"remainingItemCount": {
						SchemaProps: spec.SchemaProps{
							Description: "remainingItemCount is the number of subsequent items in the list which are not included in this list response. If the list request contained label or field selectors, then the number of remaining items is unknown and the field will be left unset and omitted during serialization. If the list is complete (either because it is not chunking or because this is the last chunk), then there are no more remaining items and this field will be left unset and omitted during serialization. Servers older than v1.15 do not set this field. The intended use of the remainingItemCount is *estimating* the size of a collection. Clients should not rely on the remainingItemCount to be set or to be exact.",
							Type:        []string{"integer"},
							Format:      "int64",
						},
					},
				},
			},
		},
	}
}

func schema_pkg_apis_meta_v1_ListOptions(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ListOptions is the query options to a standard REST list call.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"labelSelector": {
						SchemaProps: spec.SchemaProps{
							Description: "A selector to restrict the list of returned objects by their labels. Defaults to everything.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"fieldSelector": {
						SchemaProps: spec.SchemaProps{
							Description: "A selector to restrict the list of returned objects by their fields. Defaults to everything.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"watch": {
						SchemaProps: spec.SchemaProps{
							Description: "Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion.",
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
					"allowWatchBookmarks": {
						SchemaProps: spec.SchemaProps{
							Description: "allowWatchBookmarks requests watch events with type \"BOOKMARK\". Servers that do not implement bookmarks may ignore this flag and bookmarks are sent at the server's discretion. Clients should not assume bookmarks are returned at any specific interval, nor may they assume the server will send any BOOKMARK event during a session. If this is not a watch, this field is ignored.",
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
					"resourceVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "resourceVersion sets a constraint on what resource versions a request may be served from. See https://kubernetes.io/docs/reference/using-api/api-concepts/#resource-versions for details.\n\nDefaults to unset",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"resourceVersionMatch": {
						SchemaProps: spec.SchemaProps{
							Description: "resourceVersionMatch determines how resourceVersion is applied to list calls. It is highly recommended that resourceVersionMatch be set for list calls where resourceVersion is set See https://kubernetes.io/docs/reference/using-api/api-concepts/#resource-versions for details.\n\nDefaults to unset",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"timeoutSeconds": {
						SchemaProps: spec.SchemaProps{
							Description: "Timeout for the list/watch call. This limits the duration of the call, regardless of any activity or inactivity.",
							Type:        []string{"integer"},
							Format:      "int64",
						},
					},
					"limit": {
						SchemaProps: spec.SchemaProps{
							Description: "limit is a maximum number of responses to return for a list call. If more items exist, the server will set the `continue` field on the list metadata to a value that can be used with the same initial query to retrieve the next set of results. Setting a limit may return fewer than the requested amount of items (up to zero items) in the event all requested objects are filtered out and clients should only use the presence of the continue field to determine whether more results are available. Servers may choose not to support the limit argument and will return all of the available results. If limit is specified and the continue field is empty, clients may assume that no more results are available. This field is not supported if watch is true.\n\nThe server guarantees that the objects returned when using continue will be identical to issuing a single list call without a limit - that is, no objects created, modified, or deleted after the first request is issued will be included in any subsequent continued requests. This is sometimes referred to as a consistent snapshot, and ensures that a client that is using limit to receive smaller chunks of a very large result can ensure they see all possible objects. If objects are updated during a chunked list the version of the object that was present at the time the first list result was calculated is returned.",
							Type:        []string{"integer"},
							Format:      "int64",
						},
					},
					"continue": {
						SchemaProps: spec.SchemaProps{
							Description: "The continue option should be set when retrieving more results from the server. Since this value is server defined, clients may only use the continue value from a previous query result with identical query parameters (except for the value of continue) and the server may reject a continue value it does not recognize. If the specified continue value is no longer valid whether due to expiration (generally five to fifteen minutes) or a configuration change on the server, the server will respond with a 410 ResourceExpired error together with a continue token. If the client needs a consistent list, it must restart their list without the continue field. Otherwise, the client may send another list request with the token received with the 410 error, the server will respond with a list starting from the next key, but from the latest snapshot, which is inconsistent from the previous list results - objects that are created, modified, or deleted after the first list request will be included in the response, as long as their keys are after the \"next key\".\n\nThis field is not supported when watch is true. Clients may start a watch from the last resourceVersion value returned by the server and not miss any modifications.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"sendInitialEvents": {
						SchemaProps: spec.SchemaProps{
							Description: "`sendInitialEvents=true` may be set together with `watch=true`. In that case, the watch stream will begin with synthetic events to produce the current state of objects in the collection. Once all such events have been sent, a synthetic \"Bookmark\" event  will be sent. The bookmark will report the ResourceVersion (RV) corresponding to the set of objects, and be marked with `\"k8s.io/initial-events-end\": \"true\"` annotation. Afterwards, the watch stream will proceed as usual, sending watch events corresponding to changes (subsequent to the RV) to objects watched.\n\nWhen `sendInitialEvents` option is set, we require `resourceVersionMatch` option to also be set. The semantic of the watch request is as following: - `resourceVersionMatch` = NotOlderThan\n  is interpreted as \"data at least as new as the provided `resourceVersion`\"\n  and the bookmark event is send when the state is synced\n  to a `resourceVersion` at least as fresh as the one provided by the ListOptions.\n  If `resourceVersion` is unset, this is interpreted as \"consistent read\" and the\n  bookmark event is send when the state is synced at least to the moment\n  when request started being processed.\n- `resourceVersionMatch` set to any other value or unset\n  Invalid error is returned.\n\nDefaults to true if `resourceVersion=\"\"` or `resourceVersion=\"0\"` (for backward compatibility reasons) and to false otherwise.",
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
				},
			},
		},
	}
}

func schema_pkg_apis_meta_v1_ManagedFieldsEntry(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ManagedFieldsEntry is a workflow-id, a FieldSet and the group version of the resource that the fieldset applies to.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"manager": {
						SchemaProps: spec.SchemaProps{
							Description: "Manager is an identifier of the workflow managing these fields.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"operation": {
						SchemaProps: spec.SchemaProps{
							Description: "Operation is the type of operation which lead to this ManagedFieldsEntry being created. The only valid values for this field are 'Apply' and 'Update'.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the version of this resource that this field set applies to. The format is \"group/version\" just like the top-level APIVersion field. It is necessary to track the version of a field set because it cannot be automatically converted.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"time": {
						SchemaProps: spec.SchemaProps{
							Description: "Time is the timestamp of when the ManagedFields entry was added. The timestamp will also be updated if a field is added, the manager changes any of the owned fields value or removes a field. The timestamp does not update when a field is removed from the entry because another manager took it over.",
							Ref:         ref("k8s.io/apimachinery/pkg/apis/meta/v1.Time"),
						},
					},
					"fieldsType": {
						SchemaProps: spec.SchemaProps{
							Description: "FieldsType is the discriminator for the different fields format and version. There is currently only one possible value: \"FieldsV1\"",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"fieldsV1": {
						SchemaProps: spec.SchemaProps{
							Description: "FieldsV1 holds the first JSON version format as described in the \"FieldsV1\" type.",
							Ref:         ref("k8s.io/apimachinery/pkg/apis/meta/v1.FieldsV1"),
						},
					},
					"subresource": {
						SchemaProps: spec.SchemaProps{
							Description: "Subresource is the name of the subresource used to update that object, or empty string if the object was updated through the main resource. The value of this field is used to distinguish between managers, even if they share the same name. For example, a status update will be distinct from a regular update using the same manager name. Note that the APIVersion field is not related to the Subresource field and it always corresponds to the version of the main resource.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
			},
		},
		Dependencies: []string{
			"k8s.io/apimachinery/pkg/apis/meta/v1.FieldsV1", "k8s.io/apimachinery/pkg/apis/meta/v1.Time"},
	}
}

func schema_pkg_apis_meta_v1_MicroTime(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "MicroTime is version of Time with microsecond level precision.",
				Type:        v1.MicroTime{}.OpenAPISchemaType(),
				Format:      v1.MicroTime{}.OpenAPISchemaFormat(),
			},
		},
	}
}

func schema_pkg_apis_meta_v1_ObjectMeta(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"name": {
						SchemaProps: spec.SchemaProps{
							Description: "Name must be unique within a namespace. Is required when creating resources, although some resources may allow a client to request the generation of an appropriate name automatically. Name is primarily intended for creation idempotence and configuration definition. Cannot be updated. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names#names",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"generateName": {
						SchemaProps: spec.SchemaProps{
							Description: "GenerateName is an optional prefix, used by the server, to generate a unique name ONLY IF the Name field has not been provided. If this field is used, the name returned to the client will be different than the name passed. This value will also be combined with a unique suffix. The provided value has the same validation rules as the Name field, and may be truncated by the length of the suffix required to make the value unique on the server.\n\nIf this field is specified and the generated name exists, the server will return a 409.\n\nApplied only if Name is not specified. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#idempotency",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"namespace": {
						SchemaProps: spec.SchemaProps{
							Description: "Namespace defines the space within which each name must be unique. An empty namespace is equivalent to the \"default\" namespace, but \"default\" is the canonical representation. Not all objects are required to be scoped to a namespace - the value of this field for those objects will be empty.\n\nMust be a DNS_LABEL. Cannot be updated. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"selfLink": {
						SchemaProps: spec.SchemaProps{
							Description: "Deprecated: selfLink is a legacy read-only field that is no longer populated by the system.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"uid": {
						SchemaProps: spec.SchemaProps{
							Description: "UID is the unique in time and space value for this object. It is typically generated by the server on successful creation of a resource and is not allowed to change on PUT operations.\n\nPopulated by the system. Read-only. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names#uids",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"resourceVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "An opaque value that represents the internal version of this object that can be used by clients to determine when objects have changed. May be used for optimistic concurrency, change detection, and the watch operation on a resource or set of resources. Clients must treat these values as opaque and passed unmodified back to the server. They may only be valid for a particular resource or set of resources.\n\nPopulated by the system. Read-only. Value must be treated as opaque by clients and . More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"generation": {
						SchemaProps: spec.SchemaProps{
							Description: "A sequence number representing a specific generation of the desired state. Populated by the system. Read-only.",
							Type:        []string{"integer"},
							Format:      "int64",
						},
					},
					"creationTimestamp": {
						SchemaProps: spec.SchemaProps{
							Description: "CreationTimestamp is a timestamp representing the server time when this object was created. It is not guaranteed to be set in happens-before order across separate operations. Clients may not set this value. It is represented in RFC3339 form and is in UTC.\n\nPopulated by the system. Read-only. Null for lists. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata",
							Default:     map[string]interface{}{},
							Ref:         ref("k8s.io/apimachinery/pkg/apis/meta/v1.Time"),
						},
					},
					"deletionTimestamp": {
						SchemaProps: spec.SchemaProps{
							Description: "DeletionTimestamp is RFC 3339 date and time at which this resource will be deleted. This field is set by the server when a graceful deletion is requested by the user, and is not directly settable by a client. The resource is expected to be deleted (no longer visible from resource lists, and not reachable by name) after the time in this field, once the finalizers list is empty. As long as the finalizers list contains items, deletion is blocked. Once the deletionTimestamp is set, this value may not be unset or be set further into the future, although it may be shortened or the resource may be deleted prior to this time. For example, a user may request that a pod is deleted in 30 seconds. The Kubelet will react by sending a graceful termination signal to the containers in the pod. After that 30 seconds, the Kubelet will send a hard termination signal (SIGKILL) to the container and after cleanup, remove the pod from the API. In the presence of network partitions, this object may still exist after this timestamp, until an administrator or automated process can determine the resource is fully terminated. If not set, graceful deletion of the object has not been requested.\n\nPopulated by the system when a graceful deletion is requested. Read-only. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata",
							Ref:         ref("k8s.io/apimachinery/pkg/apis/meta/v1.Time"),
						},
					},
					"deletionGracePeriodSeconds": {
						SchemaProps: spec.SchemaProps{
							Description: "Number of seconds allowed for this object to gracefully terminate before it will be removed from the system. Only set when deletionTimestamp is also set. May only be shortened. Read-only.",
							Type:        []string{"integer"},
							Format:      "int64",
						},
					},
					"labels": {
						SchemaProps: spec.SchemaProps{
							Description: "Map of string keys and values that can be used to organize and categorize (scope and select) objects. May match selectors of replication controllers and services. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels",
							Type:        []string{"object"},
							AdditionalProperties: &spec.SchemaOrBool{
								Allows: true,
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: "",
										Type:    []string{"string"},
										Format:  "",
									},
								},
							},
						},
					},
					"annotations": {
						SchemaProps: spec.SchemaProps{
							Description: "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations",
							Type:        []string{"object"},
							AdditionalProperties: &spec.SchemaOrBool{
								Allows: true,
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: "",
										Type:    []string{"string"},
										Format:  "",
									},
								},
							},
						},
					},
					"ownerReferences": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-patch-merge-key": "uid",
								"x-kubernetes-patch-strategy":  "merge",
							},
						},
						SchemaProps: spec.SchemaProps{
							Description: "List of objects depended by this object. If ALL objects in the list have been deleted, this object will be garbage collected. If this object is managed by a controller, then an entry in this list will point to this controller, with the controller field set to true. There cannot be more than one managing controller.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: map[string]interface{}{},
										Ref:     ref("k8s.io/apimachinery/pkg/apis/meta/v1.OwnerReference"),
									},
								},
							},
						},
					},
					"finalizers": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-patch-strategy": "merge",
							},
						},
						SchemaProps: spec.SchemaProps{
							Description: "Must be empty before the object is deleted from the registry. Each entry is an identifier for the responsible component that will remove the entry from the list. If the deletionTimestamp of the object is non-nil, entries in this list can only be removed. Finalizers may be processed and removed in any order.  Order is NOT enforced because it introduces significant risk of stuck finalizers. finalizers is a shared field, any actor with permission can reorder it. If the finalizer list is processed in order, then this can lead to a situation in which the component responsible for the first finalizer in the list is waiting for a signal (field value, external system, or other) produced by a component responsible for a finalizer later in the list, resulting in a deadlock. Without enforced ordering finalizers are free to order amongst themselves and are not vulnerable to ordering changes in the list.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: "",
										Type:    []string{"string"},
										Format:  "",
									},
								},
							},
						},
					},
					"managedFields": {
						SchemaProps: spec.SchemaProps{
							Description: "ManagedFields maps workflow-id and version to the set of fields that are managed by that workflow. This is mostly for internal housekeeping, and users typically shouldn't need to set or understand this field. A workflow can be the user's name, a controller's name, or the name of a specific apply path like \"ci-cd\". The set of fields is always in the version that the workflow used when modifying the object.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: map[string]interface{}{},
										Ref:     ref("k8s.io/apimachinery/pkg/apis/meta/v1.ManagedFieldsEntry"),
									},
								},
							},
						},
					},
				},
			},
		},
		Dependencies: []string{
			"k8s.io/apimachinery/pkg/apis/meta/v1.ManagedFieldsEntry", "k8s.io/apimachinery/pkg/apis/meta/v1.OwnerReference", "k8s.io/apimachinery/pkg/apis/meta/v1.Time"},
	}
}

func schema_pkg_apis_meta_v1_OwnerReference(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "OwnerReference contains enough information to let you identify an owning object. An owning object must be in the same namespace as the dependent, or be cluster-scoped, so there is no namespace field.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "API version of the referent.",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"name": {
						SchemaProps: spec.SchemaProps{
							Description: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names#names",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"uid": {
						SchemaProps: spec.SchemaProps{
							Description: "UID of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names#uids",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"controller": {
						SchemaProps: spec.SchemaProps{
							Description: "If true, this reference points to the managing controller.",
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
					"blockOwnerDeletion": {
						SchemaProps: spec.SchemaProps{
							Description: "If true, AND if the owner has the \"foregroundDeletion\" finalizer, then the owner cannot be deleted from the key-value store until this reference is removed. See https://kubernetes.io/docs/concepts/architecture/garbage-collection/#foreground-deletion for how the garbage collector interacts with this field and enforces the foreground deletion. Defaults to false. To set this field, a user needs \"delete\" permission of the owner, otherwise 422 (Unprocessable Entity) will be returned.",
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
				},
				Required: []string{"apiVersion", "kind", "name", "uid"},
			},
			VendorExtensible: spec.VendorExtensible{
				Extensions: spec.Extensions{
					"x-kubernetes-map-type": "atomic",
				},
			},
		},
	}
}

func schema_pkg_apis_meta_v1_PartialObjectMetadata(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "PartialObjectMetadata is a generic representation of any object with ObjectMeta. It allows clients to get access to a particular ObjectMeta schema without knowing the details of the version.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Description: "Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata",
							Default:     map[string]interface{}{},
							Ref:         ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_meta_v1_PartialObjectMetadataList(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "PartialObjectMetadataList contains a list of objects containing only their metadata",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Description: "Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Default:     map[string]interface{}{},
							Ref:         ref("k8s.io/apimachinery/pkg/apis/meta/v1.ListMeta"),
						},
					},
					"items": {
						SchemaProps: spec.SchemaProps{
							Description: "items contains each of the included items.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: map[string]interface{}{},
										Ref:     ref("k8s.io/apimachinery/pkg/apis/meta/v1.PartialObjectMetadata"),
									},
								},
							},
						},
					},
				},
				Required: []string{"items"},
			},
		},
		Dependencies: []string{
			"k8s.io/apimachinery/pkg/apis/meta/v1.ListMeta", "k8s.io/apimachinery/pkg/apis/meta/v1.PartialObjectMetadata"},
	}
}

func schema_pkg_apis_meta_v1_Patch(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Patch is provided to give a concrete name and type to the Kubernetes PATCH request body.",
				Type:        []string{"object"},
			},
		},
	}
}

func schema_pkg_apis_meta_v1_PatchOptions(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "PatchOptions may be provided when patching an API object. PatchOptions is meant to be a superset of UpdateOptions.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"dryRun": {
						SchemaProps: spec.SchemaProps{
							Description: "When present, indicates that modifications should not be persisted. An invalid or unrecognized dryRun directive will result in an error response and no further processing of the request. Valid values are: - All: all dry run stages will be processed",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: "",
										Type:    []string{"string"},
										Format:  "",
									},
								},
							},
						},
					},
					"force": {
						SchemaProps: spec.SchemaProps{
							Description: "Force is going to \"force\" Apply requests. It means user will re-acquire conflicting fields owned by other people. Force flag must be unset for non-apply patch requests.",
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
					"fieldManager": {
						SchemaProps: spec.SchemaProps{
							Description: "fieldManager is a name associated with the actor or entity that is making these changes. The value must be less than or 128 characters long, and only contain printable characters, as defined by https://golang.org/pkg/unicode/#IsPrint. This field is required for apply requests (application/apply-patch) but optional for non-apply patch types (JsonPatch, MergePatch, StrategicMergePatch).",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"fieldValidation": {
						SchemaProps: spec.SchemaProps{
							Description: "fieldValidation instructs the server on how to handle objects in the request (POST/PUT/PATCH) containing unknown or duplicate fields. Valid values are: - Ignore: This will ignore any unknown fields that are silently dropped from the object, and will ignore all but the last duplicate field that the decoder encounters. This is the default behavior prior to v1.23. - Warn: This will send a warning via the standard warning response header for each unknown field that is dropped from the object, and for each duplicate field that is encountered. The request will still succeed if there are no other errors, and will only persist the last of any duplicate fields. This is the default in v1.23+ - Strict: This will fail the request with a BadRequest error if any unknown fields would be dropped from the object, or if any duplicate fields are present. The error returned from the server will contain all unknown and duplicate fields encountered.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
			},
		},
	}
}

func schema_pkg_apis_meta_v1_Preconditions(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Preconditions must be fulfilled before an operation (update, delete, etc.) is carried out.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"uid": {
						SchemaProps: spec.SchemaProps{
							Description: "Specifies the target UID.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"resourceVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "Specifies the target ResourceVersion",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
			},
		},
	}
}

func schema_pkg_apis_meta_v1_RootPaths(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "RootPaths lists the paths available at root. For example: \"/healthz\", \"/apis\".",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"paths": {
						SchemaProps: spec.SchemaProps{
							Description: "paths are the paths available at root.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: "",
										Type:    []string{"string"},
										Format:  "",
									},
								},
							},
						},
					},
				},
				Required: []string{"paths"},
			},
		},
	}
}

func schema_pkg_apis_meta_v1_ServerAddressByClientCIDR(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ServerAddressByClientCIDR helps the client to determine the server address that they should use, depending on the clientCIDR that they match.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"clientCIDR": {
						SchemaProps: spec.SchemaProps{
							Description: "The CIDR with which clients can match their IP to figure out the server address that they should use.",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"serverAddress": {
						SchemaProps: spec.SchemaProps{
							Description: "Address of this server, suitable for a client that matches the above CIDR. This can be a hostname, hostname:port, IP or IP:port.",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{"clientCIDR", "serverAddress"},
			},
		},
	}
}

func schema_pkg_apis_meta_v1_Status(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Status is a return value for calls that don't return other objects.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Description: "Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Default:     map[string]interface{}{},
							Ref:         ref("k8s.io/apimachinery/pkg/apis/meta/v1.ListMeta"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Description: "Status of the operation. One of: \"Success\" or \"Failure\". More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"message": {
						SchemaProps: spec.SchemaProps{
							Description: "A human-readable description of the status of this operation.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"reason": {
						SchemaProps: spec.SchemaProps{
							Description: "A machine-readable description of why this operation is in the \"Failure\" status. If this value is empty there is no information available. A Reason clarifies an HTTP status code but does not override it.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"details": {
						SchemaProps: spec.SchemaProps{
							Description: "Extended data associated with the reason.  Each reason may define its own extended details. This field is optional and the data returned is not guaranteed to conform to any schema except that defined by the reason type.",
							Ref:         ref("k8s.io/apimachinery/pkg/apis/meta/v1.StatusDetails"),
						},
					},
					"code": {
						SchemaProps: spec.SchemaProps{
							Description: "Suggested HTTP return code for this status, 0 if not set.",
							Type:        []string{"integer"},
							Format:      "int32",
						},
					},
				},
			},
		},
		Dependencies: []string{
			"k8s.io/apimachinery/pkg/apis/meta/v1.ListMeta", "k8s.io/apimachinery/pkg/apis/meta/v1.StatusDetails"},
	}
}

func schema_pkg_apis_meta_v1_StatusCause(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "StatusCause provides more information about an api.Status failure, including cases when multiple errors are encountered.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"reason": {
						SchemaProps: spec.SchemaProps{
							Description: "A machine-readable description of the cause of the error. If this value is empty there is no information available.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"message": {
						SchemaProps: spec.SchemaProps{
							Description: "A human-readable description of the cause of the error.  This field may be presented as-is to a reader.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"field": {
						SchemaProps: spec.SchemaProps{
							Description: "The field of the resource that has caused this error, as named by its JSON serialization. May include dot and postfix notation for nested attributes. Arrays are zero-indexed.  Fields may appear more than once in an array of causes due to fields having multiple errors. Optional.\n\nExamples:\n  \"name\" - the field \"name\" on the current resource\n  \"items[0].name\" - the field \"name\" on the first array entry in \"items\"",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
			},
		},
	}
}

func schema_pkg_apis_meta_v1_StatusDetails(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "StatusDetails is a set of additional properties that MAY be set by the server to provide additional information about a response. The Reason field of a Status object defines what attributes will be set. Clients must ignore fields that do not match the defined type of each attribute, and should assume that any attribute may be empty, invalid, or under defined.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"name": {
						SchemaProps: spec.SchemaProps{
							Description: "The name attribute of the resource associated with the status StatusReason (when there is a single name which can be described).",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"group": {
						SchemaProps: spec.SchemaProps{
							Description: "The group attribute of the resource associated with the status StatusReason.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "The kind attribute of the resource associated with the status StatusReason. On some operations may differ from the requested resource Kind. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"uid": {
						SchemaProps: spec.SchemaProps{
							Description: "UID of the resource. (when there is a single resource which can be described). More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names#uids",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"causes": {
						SchemaProps: spec.SchemaProps{
							Description: "The Causes array includes more details associated with the StatusReason failure. Not all StatusReasons may provide detailed causes.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: map[string]interface{}{},
										Ref:     ref("k8s.io/apimachinery/pkg/apis/meta/v1.StatusCause"),
									},
								},
							},
						},
					},
					"retryAfterSeconds": {
						SchemaProps: spec.SchemaProps{
							Description: "If specified, the time in seconds before the operation should be retried. Some errors may indicate the client must take an alternate action - for those errors this field may indicate how long to wait before taking the alternate action.",
							Type:        []string{"integer"},
							Format:      "int32",
						},
					},
				},
			},
		},
		Dependencies: []string{
			"k8s.io/apimachinery/pkg/apis/meta/v1.StatusCause"},
	}
}

func schema_pkg_apis_meta_v1_Table(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Table is a tabular representation of a set of API resources. The server transforms the object into a set of preferred columns for quickly reviewing the objects.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Description: "Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Default:     map[string]interface{}{},
							Ref:         ref("k8s.io/apimachinery/pkg/apis/meta/v1.ListMeta"),
						},
					},
					"columnDefinitions": {
						SchemaProps: spec.SchemaProps{
							Description: "columnDefinitions describes each column in the returned items array. The number of cells per row will always match the number of column definitions.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: map[string]interface{}{},
										Ref:     ref("k8s.io/apimachinery/pkg/apis/meta/v1.TableColumnDefinition"),
									},
								},
							},
						},
					},
					"rows": {
						SchemaProps: spec.SchemaProps{
							Description: "rows is the list of items in the table.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: map[string]interface{}{},
										Ref:     ref("k8s.io/apimachinery/pkg/apis/meta/v1.TableRow"),
									},
								},
							},
						},
					},
				},
				Required: []string{"columnDefinitions", "rows"},
			},
		},
		Dependencies: []string{
			"k8s.io/apimachinery/pkg/apis/meta/v1.ListMeta", "k8s.io/apimachinery/pkg/apis/meta/v1.TableColumnDefinition", "k8s.io/apimachinery/pkg/apis/meta/v1.TableRow"},
	}
}

func schema_pkg_apis_meta_v1_TableColumnDefinition(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "TableColumnDefinition contains information about a column returned in the Table.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"name": {
						SchemaProps: spec.SchemaProps{
							Description: "name is a human readable name for the column.",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"type": {
						SchemaProps: spec.SchemaProps{
							Description: "type is an OpenAPI type definition for this column, such as number, integer, string, or array. See https://github.com/OAI/OpenAPI-Specification/blob/master/versions/2.0.md#data-types for more.",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"format": {
						SchemaProps: spec.SchemaProps{
							Description: "format is an optional OpenAPI type modifier for this column. A format modifies the type and imposes additional rules, like date or time formatting for a string. The 'name' format is applied to the primary identifier column which has type 'string' to assist in clients identifying column is the resource name. See https://github.com/OAI/OpenAPI-Specification/blob/master/versions/2.0.md#data-types for more.",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"description": {
						SchemaProps: spec.SchemaProps{
							Description: "description is a human readable description of this column.",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"priority": {
						SchemaProps: spec.SchemaProps{
							Description: "priority is an integer defining the relative importance of this column compared to others. Lower numbers are considered higher priority. Columns that may be omitted in limited space scenarios should be given a higher priority.",
							Default:     0,
							Type:        []string{"integer"},
							Format:      "int32",
						},
					},
				},
				Required: []string{"name", "type", "format", "description", "priority"},
			},
		},
	}
}

func schema_pkg_apis_meta_v1_TableOptions(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "TableOptions are used when a Table is requested by the caller.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"includeObject": {
						SchemaProps: spec.SchemaProps{
							Description: "includeObject decides whether to include each object along with its columnar information. Specifying \"None\" will return no object, specifying \"Object\" will return the full object contents, and specifying \"Metadata\" (the default) will return the object's metadata in the PartialObjectMetadata kind in version v1beta1 of the meta.k8s.io API group.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
			},
		},
	}
}

func schema_pkg_apis_meta_v1_TableRow(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "TableRow is an individual row in a table.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"cells": {
						SchemaProps: spec.SchemaProps{
							Description: "cells will be as wide as the column definitions array and may contain strings, numbers (float64 or int64), booleans, simple maps, lists, or null. See the type field of the column definition for a more detailed description.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"object"},
										Format: "",
									},
								},
							},
						},
					},
					"conditions": {
						SchemaProps: spec.SchemaProps{
							Description: "conditions describe additional status of a row that are relevant for a human user. These conditions apply to the row, not to the object, and will be specific to table output. The only defined condition type is 'Completed', for a row that indicates a resource that has run to completion and can be given less visual priority.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: map[string]interface{}{},
										Ref:     ref("k8s.io/apimachinery/pkg/apis/meta/v1.TableRowCondition"),
									},
								},
							},
						},
					},
					"object": {
						SchemaProps: spec.SchemaProps{
							Description: "This field contains the requested additional information about each object based on the includeObject policy when requesting the Table. If \"None\", this field is empty, if \"Object\" this will be the default serialization of the object for the current API version, and if \"Metadata\" (the default) will contain the object metadata. Check the returned kind and apiVersion of the object before parsing. The media type of the object will always match the enclosing list - if this as a JSON table, these will be JSON encoded objects.",
							Default:     map[string]interface{}{},
							Ref:         ref("k8s.io/apimachinery/pkg/runtime.RawExtension"),
						},
					},
				},
				Required: []string{"cells"},
			},
		},
		Dependencies: []string{
			"k8s.io/apimachinery/pkg/apis/meta/v1.TableRowCondition", "k8s.io/apimachinery/pkg/runtime.RawExtension"},
	}
}

func schema_pkg_apis_meta_v1_TableRowCondition(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "TableRowCondition allows a row to be marked with additional information.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"type": {
						SchemaProps: spec.SchemaProps{
							Description: "Type of row condition. The only defined value is 'Completed' indicating that the object this row represents has reached a completed state and may be given less visual priority than other rows. Clients are not required to honor any conditions but should be consistent where possible about handling the conditions.",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Description: "Status of the condition, one of True, False, Unknown.",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"reason": {
						SchemaProps: spec.SchemaProps{
							Description: "(brief) machine readable reason for the condition's last transition.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"message": {
						SchemaProps: spec.SchemaProps{
							Description: "Human readable message indicating details about last transition.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{"type", "status"},
			},
		},
	}
}

func schema_pkg_apis_meta_v1_Time(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON.  Wrappers are provided for many of the factory methods that the time package offers.",
				Type:        v1.Time{}.OpenAPISchemaType(),
				Format:      v1.Time{}.OpenAPISchemaFormat(),
			},
		},
	}
}

func schema_pkg_apis_meta_v1_Timestamp(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Timestamp is a struct that is equivalent to Time, but intended for protobuf marshalling/unmarshalling. It is generated into a serialization that matches Time. Do not use in Go structs.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"seconds": {
						SchemaProps: spec.SchemaProps{
							Description: "Represents seconds of UTC time since Unix epoch 1970-01-01T00:00:00Z. Must be from 0001-01-01T00:00:00Z to 9999-12-31T23:59:59Z inclusive.",
							Default:     0,
							Type:        []string{"integer"},
							Format:      "int64",
						},
					},
					"nanos": {
						SchemaProps: spec.SchemaProps{
							Description: "Non-negative fractions of a second at nanosecond resolution. Negative second values with fractions must still have non-negative nanos values that count forward in time. Must be from 0 to 999,999,999 inclusive. This field may be limited in precision depending on context.",
							Default:     0,
							Type:        []string{"integer"},
							Format:      "int32",
						},
					},
				},
				Required: []string{"seconds", "nanos"},
			},
		},
	}
}

func schema_pkg_apis_meta_v1_TypeMeta(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "TypeMeta describes an individual object in an API response or request with strings representing the type of the object and its API schema version. Structures that are versioned or persisted should inline TypeMeta.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
			},
		},
	}
}

func schema_pkg_apis_meta_v1_UpdateOptions(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "UpdateOptions may be provided when updating an API object. All fields in UpdateOptions should also be present in PatchOptions.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"dryRun": {
						SchemaProps: spec.SchemaProps{
							Description: "When present, indicates that modifications should not be persisted. An invalid or unrecognized dryRun directive will result in an error response and no further processing of the request. Valid values are: - All: all dry run stages will be processed",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: "",
										Type:    []string{"string"},
										Format:  "",
									},
								},
							},
						},
					},
					"fieldManager": {
						SchemaProps: spec.SchemaProps{
							Description: "fieldManager is a name associated with the actor or entity that is making these changes. The value must be less than or 128 characters long, and only contain printable characters, as defined by https://golang.org/pkg/unicode/#IsPrint.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"fieldValidation": {
						SchemaProps: spec.SchemaProps{
							Description: "fieldValidation instructs the server on how to handle objects in the request (POST/PUT/PATCH) containing unknown or duplicate fields. Valid values are: - Ignore: This will ignore any unknown fields that are silently dropped from the object, and will ignore all but the last duplicate field that the decoder encounters. This is the default behavior prior to v1.23. - Warn: This will send a warning via the standard warning response header for each unknown field that is dropped from the object, and for each duplicate field that is encountered. The request will still succeed if there are no other errors, and will only persist the last of any duplicate fields. This is the default in v1.23+ - Strict: This will fail the request with a BadRequest error if any unknown fields would be dropped from the object, or if any duplicate fields are present. The error returned from the server will contain all unknown and duplicate fields encountered.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
			},
		},
	}
}

func schema_pkg_apis_meta_v1_WatchEvent(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Event represents a single event to a watched resource.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"type": {
						SchemaProps: spec.SchemaProps{
							Default: "",
							Type:    []string{"string"},
							Format:  "",
						},
					},
					"object": {
						SchemaProps: spec.SchemaProps{
							Description: "Object is:\n * If Type is Added or Modified: the new state of the object.\n * If Type is Deleted: the state of the object immediately before deletion.\n * If Type is Error: *Status is recommended; other types may make sense\n   depending on context.",
							Default:     map[string]interface{}{},
							Ref:         ref("k8s.io/apimachinery/pkg/runtime.RawExtension"),
						},
					},
				},
				Required: []string{"type", "object"},
			},
		},
		Dependencies: []string{
			"k8s.io/apimachinery/pkg/runtime.RawExtension"},
	}
}

func schema_k8sio_apimachinery_pkg_runtime_RawExtension(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "RawExtension is used to hold extensions in external versions.\n\nTo use this, make a field which has RawExtension as its type in your external, versioned struct, and Object in your internal struct. You also need to register your various plugin types.\n\n// Internal package:\n\n\ttype MyAPIObject struct {\n\t\truntime.TypeMeta `json:\",inline\"`\n\t\tMyPlugin runtime.Object `json:\"myPlugin\"`\n\t}\n\n\ttype PluginA struct {\n\t\tAOption string `json:\"aOption\"`\n\t}\n\n// External package:\n\n\ttype MyAPIObject struct {\n\t\truntime.TypeMeta `json:\",inline\"`\n\t\tMyPlugin runtime.RawExtension `json:\"myPlugin\"`\n\t}\n\n\ttype PluginA struct {\n\t\tAOption string `json:\"aOption\"`\n\t}\n\n// On the wire, the JSON will look something like this:\n\n\t{\n\t\t\"kind\":\"MyAPIObject\",\n\t\t\"apiVersion\":\"v1\",\n\t\t\"myPlugin\": {\n\t\t\t\"kind\":\"PluginA\",\n\t\t\t\"aOption\":\"foo\",\n\t\t},\n\t}\n\nSo what happens? Decode first uses json or yaml to unmarshal the serialized data into your external MyAPIObject. That causes the raw JSON to be stored, but not unpacked. The next step is to copy (using pkg/conversion) into the internal struct. The runtime package's DefaultScheme has conversion functions installed which will unpack the JSON stored in RawExtension, turning it into the correct object type, and storing it in the Object. (TODO: In the case where the object is of an unknown type, a runtime.Unknown object will be created and stored.)",
				Type:        []string{"object"},
			},
		},
	}
}

func schema_k8sio_apimachinery_pkg_runtime_TypeMeta(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "TypeMeta is shared by all top level objects. The proper way to use it is to inline it in your type, like this:\n\n\ttype MyAwesomeAPIObject struct {\n\t     runtime.TypeMeta    `json:\",inline\"`\n\t     ... // other fields\n\t}\n\nfunc (obj *MyAwesomeAPIObject) SetGroupVersionKind(gvk *metav1.GroupVersionKind) { metav1.UpdateTypeMeta(obj,gvk) }; GroupVersionKind() *GroupVersionKind\n\nTypeMeta is provided here for convenience. You may use it directly from this package or define your own with the same fields.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"kind": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
				},
			},
		},
	}
}

func schema_k8sio_apimachinery_pkg_runtime_Unknown(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Unknown allows api objects with unknown types to be passed-through. This can be used to deal with the API objects from a plug-in. Unknown objects still have functioning TypeMeta features-- kind, version, etc. metadata and field mutatation.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"kind": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"ContentEncoding": {
						SchemaProps: spec.SchemaProps{
							Description: "ContentEncoding is encoding used to encode 'Raw' data. Unspecified means no encoding.",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"ContentType": {
						SchemaProps: spec.SchemaProps{
							Description: "ContentType  is serialization method used to serialize 'Raw'. Unspecified means ContentTypeJSON.",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{"ContentEncoding", "ContentType"},
			},
		},
	}
}

func schema_k8sio_apimachinery_pkg_version_Info(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Info contains versioning information. how we'll want to distribute that information.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"major": {
						SchemaProps: spec.SchemaProps{
							Default: "",
							Type:    []string{"string"},
							Format:  "",
						},
					},
					"minor": {
						SchemaProps: spec.SchemaProps{
							Default: "",
							Type:    []string{"string"},
							Format:  "",
						},
					},
					"gitVersion": {
						SchemaProps: spec.SchemaProps{
							Default: "",
							Type:    []string{"string"},
							Format:  "",
						},
					},
					"gitCommit": {
						SchemaProps: spec.SchemaProps{
							Default: "",
							Type:    []string{"string"},
							Format:  "",
						},
					},
					"gitTreeState": {
						SchemaProps: spec.SchemaProps{
							Default: "",
							Type:    []string{"string"},
							Format:  "",
						},
					},
					"buildDate": {
						SchemaProps: spec.SchemaProps{
							Default: "",
							Type:    []string{"string"},
							Format:  "",
						},
					},
					"goVersion": {
						SchemaProps: spec.SchemaProps{
							Default: "",
							Type:    []string{"string"},
							Format:  "",
						},
					},
					"compiler": {
						SchemaProps: spec.SchemaProps{
							Default: "",
							Type:    []string{"string"},
							Format:  "",
						},
					},
					"platform": {
						SchemaProps: spec.SchemaProps{
							Default: "",
							Type:    []string{"string"},
							Format:  "",
						},
					},
				},
				Required: []string{"major", "minor", "gitVersion", "gitCommit", "gitTreeState", "buildDate", "goVersion", "compiler", "platform"},
			},
		},
	}
}
