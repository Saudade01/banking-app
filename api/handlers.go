package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"banking-app/models"
	"banking-app/services"
	"banking-app/utils"
)

func CreateAccountHandler(w http.ResponseWriter, r *http.Request) {
	var account models.Account
	json.NewDecoder(r.Body).Decode(&account)

	id, err := services.CreateAccount(account)
	if err != nil {
		if err.Error() == "owner already exists" {
			utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		} else {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}
	utils.RespondWithJSON(w, http.StatusCreated, map[string]int64{"id": id})
}

func GetAccountHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		utils.RespondWithError(w, http.StatusBadRequest, "Missing account ID")
		return
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid account ID")
		return
	}

	account, err := services.GetAccount(id)
	if err != nil {
		utils.RespondWithError(w, http.StatusNotFound, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, account)
}

func CreateTransferHandler(w http.ResponseWriter, r *http.Request) {
	var transfer models.Transfer
	json.NewDecoder(r.Body).Decode(&transfer)

	id, err := services.CreateTransfer(transfer)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusCreated, map[string]int64{"id": id})
}

func ListAccountsHandler(w http.ResponseWriter, r *http.Request) {
	accounts, err := services.ListAccounts()
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, accounts)
}

func ListTransfersHandler(w http.ResponseWriter, r *http.Request) {
	transfers, err := services.ListTransfers()
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, transfers)
}
