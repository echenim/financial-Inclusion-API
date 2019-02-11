// Package api provide routing to the api-end-points
// created by myron echenim on 20/may/2017

package models

//Bank models
type Bank struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
}

//Banks collection
type Banks []Bank
