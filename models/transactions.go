// Package api provide routing to the api-end-points
// created by myron echenim on 20/may/2017

package models

//Transaction models
type Transaction struct {
	ID              string  `json:"id"`
	PersonID        string  `json:"personid"`
	ProductName     string  `json:"productname"`
	PreviousBalance float64 `json:"previousbalance"`
	CurrentBalance  float64 `json:"currentbalance"`
	DebitedAount    float64 `json:"debitedamount"`
	TransFee        float64 `json:"transfee"`
	Username        string  `json:"username"`
	Latitude        string  `json:"latitude"`
	Longitude       string  `json:"longitude"`
}

//Transactions collection
type Transactions []Transaction
