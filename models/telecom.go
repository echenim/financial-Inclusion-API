// Package api provide routing to the api-end-points
// created by myron echenim on 20/may/2017

package models

//Telco models
type Telco struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
}

//Telcos collection
type Telcos []Telco
