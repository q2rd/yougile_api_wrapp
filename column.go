package yougile_api_wrapp

type Column struct {
	*YouGileClient
	Id string `json:"id,omitempty"`
}
