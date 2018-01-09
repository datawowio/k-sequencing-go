package kseq

// GetPhotoTag represents the response object that returned from Get Image Photo Tag API.
// (GET `/api/images/photo_tag`)
type GetPhotoTag struct {
	Data PhotoTagData `json:"data"`
	Meta Meta         `json:"meta"`
}

// GetPhotoTags represents the response object that returned from Get List of Image Photo
// Tag API. (GET `/api/images/photo_tags`)
type GetPhotoTags struct {
	Data PhotoTags `json:"data"`
	Meta Meta      `json:"meta"`
}

// PostPhotoTag represents the created object that returned from create Image Photo Tag
// API. (POST `/api/images/photo_tags`)
type PostPhotoTag struct {
	Data PhotoTag `json:"data"`
	Meta Meta     `json:"meta"`
}

// PhotoTag represents the Image Photo Tag object.
type PhotoTag struct {
	ID            string   `json:"id"`
	Answer        []string `json:"answer"`
	CreditCharged int      `json:"credit_charged"`
	CustomID      string   `json:"custom_id"`
	Source        string   `json:"data"`
	Instruction   string   `json:"instruction"`
	PostbackURL   string   `json:"postback_url"`
	ProcessedAt   string   `json:"processed_at"`
	ProjectID     int      `json:"project_id"`
	Status        string   `json:"status"`
}

// PhotoTags represents the list of Image Photo Tag.
type PhotoTags struct {
	Images []PhotoTag `json:"images"`
}

// PhotoTagData refers to PhotoTag object.
type PhotoTagData struct {
	Image PhotoTag `json:"image"`
}
