package apiresponse

import (
	mapi "admin/models/api"
)

//AuthResponse : response when client success login
type AuthResponse struct {
	Username string             `json:"Username"`
	DeviceID string             `json:"DeviceID"`
	Token    string             `json:"Token"`
	Response mapi.BasicResponse `json:"Response"`
}
