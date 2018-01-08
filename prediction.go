package kseq

// GetPrediction represents the response object that returned from Get Image prediction API.
// (GET `/api/prime/predictions/{id}`)
type GetPrediction struct {
	Data PredictionData `json:"data"`
	Meta Meta           `json:"meta"`
}

// GetPredictions represents the response object that returned from Get List Image
// Prediction API. (GET `/api/prime/predictions`)
type GetPredictions struct {
	Data Predictions `json:"data"`
	Meta Meta        `json:"meta"`
}

// PostPrediction represents the response object that returned from Create Image Prediction
// API. (POST `/api/prime/predictions`)
type PostPrediction struct {
	Data Prediction `json:"data"`
	Meta Meta       `json:"meta"`
}

// Prediction represents Image Prediction object.
type Prediction struct {
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

// Predictions represents list of Prediction object.
type Predictions struct {
	Images []Prediction `json:"images"`
}

// PredictionData refers to Prediction object.
type PredictionData struct {
	Image Prediction `json:"image"`
}
