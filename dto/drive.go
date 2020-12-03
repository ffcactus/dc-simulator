package dto

// Drive object in response.
type Drive struct {
	Location     string `json:"location"`
	Model        string `json:"model"`
	SerialNumber string `json:"serialNumber"`
	Manufacture  string `json:"manufacture"`
	PartNumber   string `json:"partNumber"`
	Capacity     int    `json:"capacity"`
}

// DriveResponse is response.
type DriveResponse struct {
	Items []Drive `json:"items"`
}
