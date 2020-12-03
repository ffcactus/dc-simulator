package dto

// Memory object in response.
type Memory struct {
	Location     string `json:"location"`
	Model        string `json:"model"`
	SerialNumber string `json:"serialNumber"`
	Manufacture  string `json:"manufacture"`
	PartNumber   string `json:"partNumber"`
	Capacity     int    `json:"capacity"`
}

// MemoryResponse is response.
type MemoryResponse struct {
	Items []Memory `json:"items"`
}
