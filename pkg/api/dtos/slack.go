package dtos

type SlackChannel struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
type SlackChannels struct {
	Channels []SlackChannel `json:"channels"`
}

type ShareRequest struct {
	ChannelIds      []string `json:"channelIds"`
	Message         string   `json:"message,omitempty"`
	ImagePreviewUrl string   `json:"imagePreviewUrl"`
	PanelTitle      string   `json:"panelTitle,omitempty"`
	ResourcePath    string   `json:"resourcePath"`
}
