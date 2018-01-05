package kseq

type GetClosedQuestion struct {
	Data ClosedQuestionData `json:"data"`
	Meta Meta               `json:"meta"`
}

type GetClosedQuestions struct {
	Data ClosedQuestions `json:"data"`
	Meta Meta            `json:"meta"`
}

type PostClosedQuestion struct {
	Data ClosedQuestion `json:"data"`
	Meta Meta           `json:"meta"`
}

type ClosedQuestion struct {
	ID            string `json:"id"`
	Answer        string `json:"answer"`
	CreditCharged int    `json:"credit_charged"`
	CustomID      string `json:"custom_id"`
	Source        string `json:"data"`
	PostbackURL   string `json:"postback_url"`
	ProcessedAt   string `json:"processed_at"`
	ProjectID     int    `json:"project_id"`
	Status        string `json:"status"`
}

type ClosedQuestions struct {
	Images []ClosedQuestion `json:"images"`
}

type ClosedQuestionData struct {
	Image ClosedQuestion `json:"image"`
}
