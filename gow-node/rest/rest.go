package rest

import (
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/just-do-halee/gowcoin/tree/main/gow-core/bank/sdbox"
	"github.com/just-do-halee/gowcoin/tree/main/gow-node/cmn"
	"github.com/just-do-halee/gowcoin/tree/main/gow-node/single"
)

var port string

type url string

type urlDescription struct {
	URL         url    `json:"url"`
	Method      string `json:"method"`
	Description string `json:"description"`
	Payload     string `json:"payload,omitempty"`
}

func (u url) MarshalText() ([]byte, error) {
	url := fmt.Sprintf("http://localhost%s%s", port, u)
	return []byte(url), nil
}

func documentation(w http.ResponseWriter, r *http.Request) {
	data := []urlDescription{
		{
			URL:         "/",
			Method:      "GET",
			Description: "See Documentation",
		},
		{
			URL:         "/vaults",
			Method:      "POST",
			Description: "Add a new vault",
			Payload:     "vault:Vault",
		},
		{
			URL:         "/vaults/{index}",
			Method:      "GET",
			Description: "See A Vault",
		},
	}
	cmn.HandleError(json.NewEncoder(w).Encode(data))
}

type addVaultBody struct {
	Message string
}

type errorResponse struct {
	ErrorMessage string `json:"errorMessage"`
}

func vaults(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		cmn.HandleError(json.NewEncoder(w).Encode(single.Bank.AllVaults()))
	case "POST":
		// {"vault": ...}
		var addBlockBody addVaultBody
		cmn.HandleError(json.NewDecoder(r.Body).Decode(&addBlockBody))
		single.Bank.AddOwnerVault(&sdbox.Sdboxes{})
		w.WriteHeader(http.StatusCreated)
	}
}

func vault(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	varIndex := vars["index"]
	index, err := strconv.Atoi(varIndex)
	cmn.HandleError(err)
	vault, err := single.Bank.GetVault(big.NewInt(int64(index)))
	encoder := json.NewEncoder(w)
	if err != nil {
		err = encoder.Encode(errorResponse{err.Error()})
	} else {
		err = encoder.Encode(vault)
	}
	cmn.HandleError(err)
}

func jsonContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func Start(aPort int) {
	port = fmt.Sprintf(":%d", aPort)
	router := mux.NewRouter()
	router.Use(jsonContentTypeMiddleware)
	router.HandleFunc("/", documentation).Methods("GET")
	router.HandleFunc("/vaults", vaults).Methods("GET", "POST")
	router.HandleFunc("/vaults/{index:[0-9]+}", vault).Methods("GET")
	fmt.Println("Rest server listening on http://localhost" + port)
	log.Fatal(http.ListenAndServe(port, router))
}
