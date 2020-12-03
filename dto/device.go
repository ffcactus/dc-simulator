package dto

// Device object in response.
type Device struct {
	Model        string `json:"model"`
	SerialNumber string `json:"SerialNumber"`
	Manufacture  string `json:"Manufacture"`
	PartNumber   string `json:"PartNumber"`
}
