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

type PublishAPIEndpoint struct {
	Metadata *Metadata `json:"metadata"`
	Spec     *Spec     `json:"spec"`
}

type RetireAPIEndpoint struct {
	Metadata *Metadata `json:"metadata"`
}

type Response struct {
	Code    *int    `json:"code"`
	Message *string `json:"message"`
	Details *[]struct {
		Type *string `json:"@type"`
	} `json:"details"`
}
