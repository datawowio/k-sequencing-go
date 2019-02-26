package kseq

// GetAiConsensus represents the response object that returned from Get AI Consensus API.
type GetAiConsensus struct {
	Data AiConsensusData `json:"data"`
	Meta Meta            `json:"meta"`
}

// GetAiConsensuses represents the response object that returned from Get List of AI Consensus API.
type GetAiConsensuses struct {
	Data AiConsensuses `json:"data"`
	Meta Meta          `json:"meta"`
}

// PostAiConsensus respresents the created object that returned from Create AI Consensus API.
type PostAiConsensus struct {
	Data AiConsensusData `json:"data"`
	Meta Meta            `json:"meta"`
}

// AiConsensus represents AI Consensus object.
type AiConsensus struct {
	ID            string  `json:"id"`
	Answer        string  `json:"answer"`
	CreditCharged float32 `json:"credit_charged"`
	CustomID      string  `json:"custom_id"`
	Source        string  `json:"data"`
	PostbackURL   string  `json:"postback_url"`
	ProcessedAt   string  `json:"processed_at"`
	ProjectID     int     `json:"project_id"`
	Status        string  `json:"status"`
	Mode          string  `json:"mode"`
}

// AiConsensuses represents list of AiConsensus object.
type AiConsensuses struct {
	Images []AiConsensus `json:"images"`
}

// AiConsensusData refers to AiConsensus object
type AiConsensusData struct {
	Image AiConsensus `json:"image"`
}
