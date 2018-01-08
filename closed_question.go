package kseq

// GetClosedQuestion represents the response object that returned from Get Image Closed
// Question API. (GET `/api/images/closed_question`)
type GetClosedQuestion struct {
	Data ClosedQuestionData `json:"data"`
	Meta Meta               `json:"meta"`
}

// GetClosedQuestions represents the response object that returned from
// Get List of Image Closed Questions API. (GET `/api/images/closed_questions`)
type GetClosedQuestions struct {
	Data ClosedQuestions `json:"data"`
	Meta Meta            `json:"meta"`
}

// PostClosedQuestion respresents the created object that returned from Create Image Closed
// Question API. (POST `/api/images/closed_question`)
type PostClosedQuestion struct {
	Data ClosedQuestion `json:"data"`
	Meta Meta           `json:"meta"`
}

// ClosedQuestion represents Image Closed Question object.
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

// ClosedQuestions represents list of ClosedQuestion object.
type ClosedQuestions struct {
	Images []ClosedQuestion `json:"images"`
}

// ClosedQuestionData refers to ClosedQuestion object
type ClosedQuestionData struct {
	Image ClosedQuestion `json:"image"`
}
