// Package api provide routing to the api-end-points
// created by myron echenim on 20/may/2017

package models

//AirTimeVending vending
type AirTimeVending struct {
	Telco       string `json:"telco"`
	PhoneNumber string `json:"phonenumber"`
	Amount      string `json:"amount"`
	Username    string `json:"username"`
	Latitude    string `json:"latitude"`
	Longitude   string `json:"longitude"`
}
