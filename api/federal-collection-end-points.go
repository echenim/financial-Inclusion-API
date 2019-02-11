// Package api provide routing to the api-end-points
// created by myron echenim on 07/may/2017
package api

//

import (
	"encoding/json"
	"fmt"
	"net/http"

	entity "github.com/echenim/financial-inclusion-api/models"
)

//returnAllAvailableCollection
//endpoint to return all kind of collections
func returnAllAvailableCollection(writeresponse http.ResponseWriter, request *http.Request) {

	if request.Method == "GET" {
		data := entity.Banks{
			entity.Bank{ID: "1", Name: "FIRS", Code: "FIRS"},
			entity.Bank{ID: "2", Name: "JAMB", Code: "JAMB"},
		}
		fmt.Printf("EndPoint : List Of Available Collections")
		json.NewEncoder(writeresponse).Encode(data)

	} else {
		http.Error(writeresponse, "Error 405", http.StatusMethodNotAllowed)
	}
}
