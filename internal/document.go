package internal

type Document struct {
	Content   string `json:"content,omitempty"`
	Title     string `json:"title,omitempty"`
	Author    string `json:"author,omitempty"`
	Topic     string `json:"topic,omitempty"`
	Watermark string `json:"watermark,omitempty"`
}

type Filter struct {
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
}

type Status string

const (
	Pending    Status = "Pending"
	Started    Status = "Started"
	InProgress Status = "InProgress"
	Finished   Status = "Finished"
	Failed     Status = "Failed"
)
