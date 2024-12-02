package package_manager

import "strings"

const (
	L_github = "/archive/refs/heads/main.zip"
)

const (
	D_github = "github.com"
)

func Domain(url string) (string, string) {
	url = strings.Replace(url, "http://", "", 1)
	url = strings.Replace(url, "https://", "", 1)
	sp := strings.Split(url, "/")
	domain := sp[0]
	username := sp[1]
	projectname := sp[2]
	if domain == D_github {
		return projectname + ".zip", "https://" + D_github + "/" + username + "/" + projectname + L_github
	}
	return "", ""
}
