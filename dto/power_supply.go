package dto

// PowerSupply object in response.
type PowerSupply struct {
	Location     string `json:"location"`
	Model        string `json:"model"`
	SerialNumber string `json:"serialNumber"`
	Manufacture  string `json:"manufacture"`
	PartNumber   string `json:"partNumber"`
}

// PowerSupplyResponse is response.
type PowerSupplyResponse struct {
	Items []PowerSupply `json:"items"`
}
