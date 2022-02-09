package model

type Version struct {
	Version_id     int    `json:"version_id"`
	Version_string string `json:"version_string"`
	Code_version   int    `json:"code_version"`
	Create_time    string `json:"create_time"`
	Platform       string `json:"platform"`
}
