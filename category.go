package kseq

type Category struct {
	ID           string   `json:"id"`
	Title        string   `json:"title"`
	Categories   []string `json:"categories"`
	CustomID     string   `json:"custom_id"`
	Conversation []Conv   `json:"conversation"`
	PostbackURL  string   `json:"postback_url"`
	ProcessedAt  string   `json:"processed_at"`
	ProjectID    int      `json:"project_id"`
	Status       string   `json:"status"`
}

type CategoryGet struct {
	Data Category `json:"data"`
	Meta Meta     `json:"meta"`
}

type CategoryCreate struct {
	Data Category `json:"data"`
	Meta Meta     `json:"meta"`
}

type CategoryList struct {
	Data []Category `json:"data"`
	Meta Meta       `json:"meta"`
}

type Conv struct {
	Name    string `json:"name"`
	Message string `json:"message"`
}
