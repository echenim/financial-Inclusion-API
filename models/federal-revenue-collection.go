package models

//FedCollection revenue collection
type FedCollection struct {
	CollectionName string `json:"telco"`
	Payername      string `json:"payername"`
	Amount         string `json:"amount"`
	Username       string `json:"username"`
	Latitude       string `json:"latitude"`
	Longitude      string `json:"longitude"`
}
