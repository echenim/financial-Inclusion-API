// Package api provide routing to the api-end-points
// created by myron echenim on 06/may/2017
package api

//

import (
	"encoding/json"
	"fmt"
	"net/http"

	entity "github.com/echenim/financial-inclusion-api/models"
)

//returnAllAvailableTelcos history
//endpoint to return all transaction history
func returnAllAvailableTelcos(writeresponse http.ResponseWriter, request *http.Request) {

	if request.Method == "GET" {
		data := entity.Banks{
			entity.Bank{ID: "1", Name: "Airtel", Code: "ATT"},
			entity.Bank{ID: "2", Name: "Glo", Code: "GLO"},
			entity.Bank{ID: "3", Name: "MTN", Code: "MTN"},
			entity.Bank{ID: "4", Name: "Etisalat", Code: "ETS"},
		}
		fmt.Printf("EndPoint : List Of Available Telecom provider")
		json.NewEncoder(writeresponse).Encode(data)

	} else {
		http.Error(writeresponse, "Error 405", http.StatusMethodNotAllowed)
	}
}

//vendAirTime vend air time
func vendAirTime(writeresponse http.ResponseWriter, request *http.Request) {

	if request.Method == "POST" {
		decoder := json.NewDecoder(request.Body)
		var requestObj entity.AirTimeVending
		succeed := decoder.Decode(&requestObj)
		if succeed == nil {
			// process data sent to repository
			fmt.Println(requestObj.Username + " | " + requestObj.Telco + "  |  " + requestObj.Amount)

		} else {
			http.Error(writeresponse, "Error 500", http.StatusMethodNotAllowed)
		}

	} else {
		http.Error(writeresponse, "Error 405", http.StatusMethodNotAllowed)
	}
}
