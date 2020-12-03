package dto

// Fan object in response.
type Fan struct {
	Location     string `json:"location"`
	Model        string `json:"model"`
	SerialNumber string `json:"SerialNumber"`
	Manufacture  string `json:"Manufacture"`
	PartNumber   string `json:"PartNumber"`
}

// FanResponse is response.
type FanResponse struct {
	Items []Fan `json:"items"`
}
