package kseq

// GetMessage represents the response MessageData object that returned from Get Image
// Message API. (GET `/api/images/message`)
type GetMessage struct {
	Data MessageData `json:"data"`
	Meta Meta        `json:"meta"`
}

// GetMessages represents the response of Messages object that returned from Get Image
// Messages API. (GET `/api/images/messages`)
type GetMessages struct {
	Data Messages `json:"data"`
	Meta Meta     `json:"meta"`
}

// PostMessage represents the created object that returned from Create Image Message API.
// (POST `/api/images/messages`)
type PostMessage struct {
	Data Message `json:"data"`
	Meta Meta    `json:"meta"`
}

// Message represents the Image Message object.
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

// Messages represents list of Image Message object.
type Messages struct {
	Images []Message `json:"images"`
}

// MessageData refers to Message object;
type MessageData struct {
	Image Message `json:"image"`
}
