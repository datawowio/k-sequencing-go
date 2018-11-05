package kseq

// GetImageCheck represents the response object that returned from Get Image Check
// API. (GET `/api/images/check`)
type GetImageCheck struct {
	Data ImageCheckData `json:"data"`
	Meta Meta           `json:"meta"`
}

// GetImageChecks represents the response object that returned from
// Get List of Image Check API. (GET `/api/images/check`)
type GetImageChecks struct {
	Data ImageChecks `json:"data"`
	Meta Meta        `json:"meta"`
}

// PostImageCheck respresents the created object that returned from Create Image Check
// API. (POST `/api/images/check`)
type PostImageCheck struct {
	Data ImageCheckData `json:"data"`
	Meta Meta           `json:"meta"`
}

// ImageCheck represents Image Check object.
type ImageCheck struct {
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

// ImageChecks represents list of ImageCheck object.
type ImageChecks struct {
	Images []ImageCheck `json:"images"`
}

// ImageCheckData refers to ImageCheck object
type ImageCheckData struct {
	Image ImageCheck `json:"image"`
}
