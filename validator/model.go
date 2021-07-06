package validator

type Payload struct {
	Auths map[string]struct {
		Auth  string `json:"auth"`
		Email string `json:"email,omitempty"`
	}
}
