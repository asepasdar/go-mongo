package apirequest

//AuthRequest : request auth model
type AuthRequest struct {
	Username  string `json:"Username"`
	DeviceID  string `json:"DeviceID"`
	ClientKey string `json:"ClientKey"`
}
