package main

type Project struct {
	Count     int    `json:"count"`
	Href      string `json:"href"`
	BuildType []struct {
		ID          string `json:"id"`
		Name        string `json:"name"`
		ProjectName string `json:"projectName"`
		ProjectID   string `json: "projectId"`
		Href        string `json:"href"`
		WebURL      string `json:"weburl"`
	}
}
