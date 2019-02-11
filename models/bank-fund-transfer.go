// Package api provide routing to the api-end-points
// created by myron echenim on 20/may/2017

package models

//CashTransfer cash transfer
type CashTransfer struct {
	Sender            string `json:"sender"`
	SenderPhoneNumber string `json:"phonenumber"`
	Receiver          string `json:"receiver"`
	Amount            string `json:"amount"`
	Tranfee           string `json:"transfee"`
	TransferType      string `json:"transfertype"`
	Username          string `json:"username"`
	Latitude          string `json:"latitude"`
	Longitude         string `json:"longitude"`
}
