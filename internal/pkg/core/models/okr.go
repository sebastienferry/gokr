package models

type Okr struct {
	Id          int64       `json:"id"`
	OrgId       int64       `json:"org_id"`
	ParentId    int64       `json:"parent_id"`
	Title       string      `json:"title"`
	Description string      `json:"description"`
	KeyResults  []KeyResult `json:"key_results"`
}

type KeyResult struct {
	Id          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
