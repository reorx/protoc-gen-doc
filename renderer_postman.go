package gendoc

import (
	"encoding/json"
	"time"
)

type PostmanRequest struct {
	Name    string `json:"name"`
	Request struct {
		Method string `json:"method"`
		Header []struct {
			Key   string `json:"key"`
			Value string `json:"value"`
			Type  string `json:"type"`
		} `json:"header"`
		Body struct {
			Mode    string    `json:"mode"`
			Raw     time.Time `json:"raw"`
			Options struct {
				Raw struct {
					Language string `json:"language"`
				} `json:"raw"`
			} `json:"options"`
		} `json:"body"`
		URL struct {
			Raw  string   `json:"raw"`
			Host []string `json:"host"`
			Path []string `json:"path"`
		} `json:"url"`
	} `json:"request"`
	Response []interface{} `json:"response"`
}

type PostmanFolder struct {
	Name string           `json:"name"`
	Item []PostmanRequest `json:"item"`
}

type PostmanCollection struct {
	Info struct {
		PostmanID string `json:"_postman_id"`
		Name      string `json:"name"`
		Schema    string `json:"schema"`
	} `json:"info"`
	Item []PostmanFolder `json:"item"`
}

type postmanRenderer struct{}

func (r *postmanRenderer) Apply(template *Template) ([]byte, error) {
	return json.MarshalIndent(template, "", "  ")
}
