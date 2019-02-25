package kseq

// GetAiConsensus represents the response object that returned from Get Image Check
// API. (GET `/api/images/check/:id`)
type GetAiConsensus struct {
	Data AiConsensusData `json:"data"`
	Meta Meta            `json:"meta"`
}

// GetAiConsensuses represents the response object that returned from
// Get List of Image Check API. (GET `/api/images/check`)
type GetAiConsensuses struct {
	Data AiConsensuses `json:"data"`
	Meta Meta          `json:"meta"`
}

// PostAiConsensus respresents the created object that returned from Create Image Check
// API. (POST `/api/images/check`)
type PostAiConsensus struct {
	Data AiConsensusData `json:"data"`
	Meta Meta            `json:"meta"`
}

// AiConsensus represents Image Check object.
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
}

// AiConsensuses represents list of ImageCheck object.
type AiConsensuses struct {
	Images []AiConsensus `json:"images"`
}

// AiConsensusData refers to ImageCheck object
type AiConsensusData struct {
	Image AiConsensus `json:"image"`
}
