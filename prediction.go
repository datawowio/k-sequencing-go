package main

type GetPrediction struct {
	Data PredictionData `json:"data"`
	Meta Meta           `json:"meta"`
}

type GetPredictions struct {
	Data Predictions `json:"data"`
	Meta Meta        `json:"meta"`
}

type PostPrediction struct {
	Data Prediction `json:"data"`
	Meta Meta       `json:"meta"`
}

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

type Predictions struct {
	Images []Prediction `json:"images"`
}

type PredictionData struct {
	Image Prediction `json:"image"`
}
