//Package api
// created by myron echenim on 13/may/2017

package api

import (
	"log"
	"net/http"
)

//HandleRequest request
func HandleRequest() {

	http.HandleFunc("/api/availablebalance", returnAllTransactionHistories)
	http.HandleFunc("/api/history", returnAllTransactionHistories)
	http.HandleFunc("/api/history/byAgentId", returnAllTransactionHistories)
	http.HandleFunc("/api/available/banks", returnAllAvailableBanks)
	http.HandleFunc("/api/available/telcos", returnAllTransactionHistories)
	http.HandleFunc("/api/available/fedralcollections", returnAllTransactionHistories)
	http.HandleFunc("/api/telco/vending", vendAirTime)
	http.HandleFunc("/api/bank/transfer", vendAirTime)
	http.HandleFunc("/api/bank/accountopening", vendAirTime)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
