package kseq

type TextClosedQuestion struct {
	ID          string `json:"id"`
	Data        string `json:"data"`
	CustomID    string `json:"custom_id"`
	Answer      string `json:"answer"`
	PostbackURL string `json:"postback_url"`
	ProcessedAt string `json:"processed_at"`
	ProjectID   int    `json:"project_id"`
	Status      string `json:"status"`
}

type TextClosedQuestionGet struct {
	Data Conversation `json:"data"`
	Meta Meta         `json:"meta"`
}

type TextClosedQuestionList struct {
	Data []TextClosedQuestion `json:"data"`
	Meta Meta                 `json:"meta"`
}

type TextClosedQuestionCreate struct {
	Data TextClosedQuestion `json:"data"`
	Meta Meta               `json:"meta"`
}
