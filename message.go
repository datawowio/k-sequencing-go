package main

type GetMessage struct {
	Data MessageData `json:"data"`
	Meta Meta        `json:"meta"`
}

type GetMessages struct {
	Data Messages `json:"data"`
	Meta Meta     `json:"meta"`
}

type PostMessage struct {
	Data Message `json:"data"`
	Meta Meta    `json:"meta"`
}

type Message struct {
	ID            string `json:"id"`
	Answer        string `json:"answer"`
	CreditCharged int    `json:"credit_charged"`
	CustomID      string `json:"custom_id"`
	Source        string `json:"data"`
	Instruction   string `json:"instruction"`
	PostbackURL   string `json:"postback_url"`
	ProcessedAt   string `json:"processed_at"`
	ProjectID     int    `json:"project_id"`
	Status        string `json:"status"`
}

type Messages struct {
	Images []Message `json:"images"`
}

type MessageData struct {
	Image Message `json:"image"`
}
