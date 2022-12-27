package webfinger

const WebFingerPath = "/.well-known/webfinger"

var Domain = "test.meromeromeiro.top"

type Link struct {
	HRef       string             `json:"href,omitempty"`
	Type       string             `json:"type,omitempty"`
	Rel        string             `json:"rel,omitempty"`
	Template   string             `json:"template,omitempty"`
	Properties map[string]*string `json:"properties,omitempty"`
	Titles     map[string]string  `json:"titles,omitempty"`
}

type Resource struct {
	Subject    string            `json:"subject,omitempty"`
	Aliases    []string          `json:"aliases,omitempty"`
	Properties map[string]string `json:"properties,omitempty"`
	Links      []Link            `json:"links"`
}
