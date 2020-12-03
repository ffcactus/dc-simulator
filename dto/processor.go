package dto

// Processor object in response.
type Processor struct {
	Location     string `json:"location"`
	Model        string `json:"model"`
	SerialNumber string `json:"SerialNumber"`
	Manufacture  string `json:"Manufacture"`
	PartNumber   string `json:"PartNumber"`
}

// ProcessorResponse is response.
type ProcessorResponse struct {
	Items []Processor `json:"items"`
}
