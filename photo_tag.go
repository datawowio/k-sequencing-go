package kseq

type GetPhotoTag struct {
	Data PhotoTagData `json:"data"`
	Meta Meta         `json:"meta"`
}

type GetPhotoTags struct {
	Data PhotoTags `json:"data"`
	Meta Meta      `json:"meta"`
}

type PostPhotoTag struct {
	Data PhotoTag `json:"data"`
	Meta Meta     `json:"meta"`
}

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

type PhotoTags struct {
	Images []PhotoTag `json:"images"`
}

type PhotoTagData struct {
	Image PhotoTag `json:"image"`
}
