package client

type Metadata struct {
	Identifier *string `json:"identifier"`
	Stage      *string `json:"stage"`
}

type OpenAPI struct {
	Base64Encoded *string `json:"base64Encoded"`
}

type Spec struct {
	OpenAPI *OpenAPI `json:"openApi"`
}

type ProjectPublish struct {
	Metadata *Metadata `json:"metadata"`
	Spec     *Spec     `json:"spec"`
}

type ProjectRetire struct {
	Metadata *Metadata `json:"metadata"`
}

type ProjectPublishResponse struct {
	Code    *int    `json:"code"`
	Message *string `json:"message"`
	Details *[]struct {
		Type *string `json:"@type"`
	} `json:"details"`
}

type ProjectRetireResponse struct {
	Code    *int    `json:"code"`
	Message *string `json:"message"`
	Details *[]struct {
		Type *string `json:"@type"`
	} `json:"details"`
}
