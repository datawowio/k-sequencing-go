package kseq

type Profanity struct {
	ID          string `json:"id"`
	Data        string `json:"data"`
	CustomID    string `json:"custom_id"`
	Answer      string `json:"answer"`
	PostbackURL string `json:"postback_url"`
	ProcessedAt string `json:"processed_at"`
	ProjectID   int    `json:"project_id"`
	Status      string `json:"status"`
}

type ProfanityPost struct {
	Data Profanity `json:"data"`
	Meta Meta      `json:"meta"`
}
