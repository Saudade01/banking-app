package handlers

import (
	"banking-app/models"
	"banking-app/services"
	"banking-app/utils"
	"encoding/json"
	"net/http"
)

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

func ListTransfersHandler(w http.ResponseWriter, r *http.Request) {
	transfers, err := services.ListTransfers()
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, transfers)
}
