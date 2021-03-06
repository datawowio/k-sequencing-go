package kseq

type Filter struct {
	ID         string   `json:"id"`
	FilterSet  []string `json:"filter_set"`
	UseDefault bool     `json:"use_default"`
}

type FilterCreate struct {
	Data Filter `json:"data"`
	Meta Meta   `json:"meta"`
}
