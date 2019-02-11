// Package api provide routing to the api-end-points
// created by myron echenim on 04/may/2017
package api

//

import (
	"encoding/json"
	"fmt"
	"net/http"

	entity "github.com/echenim/financial-inclusion-api/models"
)

//returnAllTransactionHistories endpoint for returning all avaialble transaction history
func returnAllAvailableBanks(writeresponse http.ResponseWriter, request *http.Request) {

	if request.Method == "GET" {
		data := entity.Banks{
			entity.Bank{ID: "1", Name: "Access Bank Plc", Code: "ACB"},
			entity.Bank{ID: "2", Name: "Fidelity Bank Plc", Code: "FIB"},
			entity.Bank{ID: "3", Name: "First City Monument Bank Plc", Code: "FCM"},
			entity.Bank{ID: "4", Name: "First Bank of Nigeria Limited", Code: "FBN"},
			entity.Bank{ID: "5", Name: "Guaranty Trust Bank Plc", Code: "GTB"},
			entity.Bank{ID: "6", Name: "Union Bank of Nigeria Plc", Code: "UBN"},
			entity.Bank{ID: "7", Name: "United Bank for Africa Plc", Code: "UBA"},
			entity.Bank{ID: "8", Name: "Zenith Bank Plc", Code: "ZNB"},
			entity.Bank{ID: "9", Name: "Citibank Nigeria Limited", Code: "CNB"},
			entity.Bank{ID: "10", Name: "Ecobank Nigeria Plc", Code: "ECB"},
			entity.Bank{ID: "11", Name: "Heritage Banking Company Limited", Code: "HCB"},
			entity.Bank{ID: "12", Name: "Keystone Bank Limited", Code: "KSB"},
			entity.Bank{ID: "13", Name: "Polaris Bank Limited", Code: "POB"},
			entity.Bank{ID: "14", Name: "Stanbic IBTC Bank Plc", Code: "SIB"},
			entity.Bank{ID: "15", Name: "Standard Chartered", Code: "SCB"},
			entity.Bank{ID: "16", Name: "Sterling Bank Plc", Code: "STB"},
			entity.Bank{ID: "17", Name: "Unity Bank Plc", Code: "UNB"},
		}
		fmt.Printf("EndPoint : List Of Available Banks")
		json.NewEncoder(writeresponse).Encode(data)

	} else {
		http.Error(writeresponse, "Error 405", http.StatusMethodNotAllowed)
	}
}

//accountOpening endpoint for account opening
func accountOpening(writeresponse http.ResponseWriter, request *http.Request) {

	if request.Method == "POST" {
		decoder := json.NewDecoder(request.Body)
		var requestObj entity.AccountOpening
		succeed := decoder.Decode(&requestObj)
		if succeed == nil {
			// process data sent to repository
			fmt.Println(requestObj.Username)

		} else {
			http.Error(writeresponse, "Error 500", http.StatusMethodNotAllowed)
		}

	} else {
		http.Error(writeresponse, "Error 405", http.StatusMethodNotAllowed)
	}
}

//fundTransfer endpoint for all fund transfer  operation
func fundTransfer(writeresponse http.ResponseWriter, request *http.Request) {

	if request.Method == "POST" {
		decoder := json.NewDecoder(request.Body)
		var requestObj entity.Transaction
		succeed := decoder.Decode(&requestObj)
		if succeed == nil {
			// process data sent to repository
			fmt.Println(requestObj.Username)

		} else {
			http.Error(writeresponse, "Error 500", http.StatusMethodNotAllowed)
		}

	} else {
		http.Error(writeresponse, "Error 405", http.StatusMethodNotAllowed)
	}
}
