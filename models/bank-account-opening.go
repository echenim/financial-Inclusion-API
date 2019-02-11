// Package api provide routing to the api-end-points
// created by myron echenim on 20/may/2017

package models

//AccountOpening vending
type AccountOpening struct {
	Firstname        string `json:"firstname"`
	Lastname         string `json:"lastname"`
	Othername        string `json:"othername"`
	Gender           string `json:"gender"`
	Phonenumber      string `json:"phonenumber"`
	Address          string `json:"address"`
	Dateofbirth      string `json:"dateofbirth"`
	Placeofbirth     string `json:"placeofbirth"`
	Nextofkinname    string `json:"nextofkinname"`
	Nextofkinphone   string `json:"nextofkinphone"`
	Nextofkinaddress string `json:"nextofkinaddress"`
	Product          string `json:"product"`
	Cardserialnumber string `json:"cardserialnumber"`
	Username         string `json:"username"`
	Latitude         string `json:"latitude"`
	Longitude        string `json:"longitude"`
}
