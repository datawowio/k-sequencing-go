package kseq

type Category struct {
	ID           string         `json:"id"`
	Title        string         `json:"title"`
	Categories   []string       `json:"categories"`
	CustomID     string         `json:"custom_id"`
	Conversation []Conversation `json:"conversation"`
	PostbackURL  string         `json:"postback_url"`
	ProcessedAt  string         `json:"processed_at"`
	ProjectID    int            `json:"project_id"`
	Status       string         `json:"status"`
}

type CategoryPost struct {
	Data Category `json:"data"`
	Meta Meta     `json:"meta"`
}

type Conv struct {
	Name    string
	Message string
}
