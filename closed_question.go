package main

type ClosedQuestion struct {
	Data Data `json:"data"`
	Meta Meta `json:"meta"`
}

type Image struct {
	ID            string `json:"id"`
	Answer        string `json:"answer"`
	CreditCharged int    `json:"credit_charged"`
	CustomID      string `json:"custom_id"`
	Source        string `json:"data"`
	PostbackUrl   string `json:"postback_url"`
	ProcessedAt   string `json:"processed_at"`
	ProjectID     int    `json:"project_id"`
	Status        string `json:"status"`
}

type Meta struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Data struct {
	Image Image `json:"image"`
}
