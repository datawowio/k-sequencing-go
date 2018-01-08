package kseq

// GetChoice represents the response object that returned from Get Image Choice API.
// (GET `/api/images/choice`)
type GetChoice struct {
	Data ChoiceData `json:"data"`
	Meta Meta       `json:"meta"`
}

// GetChoices represents the response object that returned from Get List of Image Choice API.
// (GET `/api/images/choices`)
type GetChoices struct {
	Data Choices `json:"data"`
	Meta Meta    `json:"meta"`
}

// PostChoice represents the created response object that returned from Create Image Choice
// API. (POST `/api/images/choices`)
type PostChoice struct {
	Data Choice `json:"data"`
	Meta Meta   `json:"meta"`
}

// Choice represents Image Choice object.
type Choice struct {
	ID            string   `json:"id"`
	AllowEmpty    bool     `json:"allow_empty"`
	Answer        []string `json:"answer"`
	Categories    []string `json:"categories"`
	CreditCharged int      `json:"credit_charged"`
	CustomID      string   `json:"custom_id"`
	Source        string   `json:"data"`
	Instruction   string   `json:"instruction"`
	Multiple      bool     `json:"multiple"`
	PostbackURL   string   `json:"postback_url"`
	ProcessedAt   string   `json:"processed_at"`
	ProjectID     int      `json:"project_id"`
	Status        string   `json:"status"`
}

// Choices represents list of Image Choice object.
type Choices struct {
	Images []Choice `json:"images"`
}

// ChoiceData refers Image Choice object.
type ChoiceData struct {
	Image Choice `json:"image"`
}
