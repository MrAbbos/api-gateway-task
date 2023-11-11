package models

type UpdatePatch struct {
	ID   string                 `json:"id"`
	Data map[string]interface{} `json:"data"`
}

type Empty struct{}