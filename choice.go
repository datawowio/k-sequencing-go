package main

type GetChoice struct {
	Data ChoiceData `json:"data"`
	Meta Meta       `json:"meta"`
}

type GetChoices struct {
	Data Choices `json:"data"`
	Meta Meta    `json:"meta"`
}

type PostChoice struct {
	Data Choice `json:"data"`
	Meta Meta   `json:"meta"`
}

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

type Choices struct {
	Images []Choice `json:"images"`
}

type ChoiceData struct {
	Image Choice `json:"image"`
}
