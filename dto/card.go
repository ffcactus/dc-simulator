package dto

// Card object in response.
type Card struct {
	Location     string `json:"location"`
	Model        string `json:"model"`
	SerialNumber string `json:"serialNumber"`
	Manufacture  string `json:"manufacture"`
	PartNumber   string `json:"partNumber"`
}

// CardResponse is response.
type CardResponse struct {
	Items []Card `json:"items"`
}
