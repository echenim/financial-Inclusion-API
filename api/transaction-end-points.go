// Package api provide routing to the api-end-points
// created by myron echenim on 04/may/2017
package api

//

import (
	"encoding/json"
	"net/http"

	entity "github.com/echenim/financial-inclusion-api/models"
)

//returnAllTransactionHistories history
//endpoint to return all transaction history
func returnAllTransactionHistories(writeresponse http.ResponseWriter, request *http.Request) {

	if request.Method == "GET" {
		transactions := entity.Transactions{
			entity.Transaction{ID: "1", PersonID: "1", ProductName: "Bank:FundTransfer:CashToCash", PreviousBalance: 20000, CurrentBalance: 40000, DebitedAount: 3939, TransFee: 23},
			entity.Transaction{ID: "2", PersonID: "1", ProductName: "Telcos:AirTimeVending:MTN", PreviousBalance: 20000, CurrentBalance: 40000, DebitedAount: 3939, TransFee: 23},
		}

		json.NewEncoder(writeresponse).Encode(transactions)

	} else {
		http.Error(writeresponse, "Error 405", http.StatusMethodNotAllowed)
	}

}
