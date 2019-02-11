// Package api provide routing to the api-end-points
// created by myron echenim on 20/may/2017

package models

//FederalCollection models
type FederalCollection struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
}

//FederalCollections collection
type FederalCollections []FederalCollection
