package validator

type ValidationErrorResponse struct {
	Filed string `json:"filed"`
	Rule  string `json:"rule"`
	Value string `json:"value"`
}
