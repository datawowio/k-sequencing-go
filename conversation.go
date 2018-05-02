package kseq

type Conversation struct {
	ID                    string   `json:"id"`
	Conversation          []string `json:"conversation"`
	Answers               []string `json:"answers"`
	CustomConversationIDs []string `json:"custom_conversation_ids"`
	CustomID              string   `json:"custom_id"`
	PostbackURL           string   `json:"postback_url"`
	ProcessedAt           string   `json:"processed_at"`
	ProjectID             int      `json:"project_id"`
	Status                string   `json:"status"`
}

type ConversationGet struct {
	Data Conversation `json:"data"`
	Meta Meta         `json:"meta"`
}

type ConversationCreate struct {
	Data Conversation `json:"data"`
	Meta Meta         `json:"meta"`
}
