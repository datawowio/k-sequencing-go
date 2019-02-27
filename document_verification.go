package kseq

type DocumentVerification struct {
	ID          string                       `json:"id"`
	Data        string                       `json:"data"`
	Info        map[string]map[string]string `json:"info"`
	CustomID    string                       `json:"custom_id"`
	PostbackURL string                       `json:"postback_url"`
	ProcessedAt string                       `json:"processed_at"`
	ProjectID   int                          `json:"project_id"`
	Status      string                       `json:"status"`
}

type DocumentVerificationGet struct {
	Data DocumentVerification `json:"data"`
	Meta Meta                 `json:"meta"`
}

type DocumentVerificationCreate struct {
	Data DocumentVerification `json:"data"`
	Meta Meta                 `json:"meta"`
}
