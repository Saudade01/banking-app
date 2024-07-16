package services

import (
	"banking-app/config"
	"banking-app/models"
	"errors"
)

func CreateTransfer(transfer models.Transfer) (int64, error) {
	tx, err := config.DB.Begin()
	if err != nil {
		return 0, err
	}

	var fromAccount models.Account
	query := "SELECT id, balance FROM accounts WHERE id = ?"
	if err := tx.QueryRow(query, transfer.FromAccountID).Scan(&fromAccount.ID, &fromAccount.Balance); err != nil {
		tx.Rollback()
		return 0, err
	}

	if fromAccount.Balance < transfer.Amount {
		tx.Rollback()
		return 0, errors.New("insufficient funds")
	}

	query = "UPDATE accounts SET balance = balance - ? WHERE id = ?"
	if _, err := tx.Exec(query, transfer.Amount, transfer.FromAccountID); err != nil {
		tx.Rollback()
		return 0, err
	}

	query = "UPDATE accounts SET balance = balance + ? WHERE id = ?"
	if _, err := tx.Exec(query, transfer.Amount, transfer.ToAccountID); err != nil {
		tx.Rollback()
		return 0, err
	}

	query = "INSERT INTO transfers (from_account_id, to_account_id, amount) VALUES (?, ?, ?)"
	result, err := tx.Exec(query, transfer.FromAccountID, transfer.ToAccountID, transfer.Amount)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	if err := tx.Commit(); err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

func ListTransfers() ([]models.Transfer, error) {
	rows, err := config.DB.Query("SELECT id, from_account_id, to_account_id, amount FROM transfers")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	transfers := []models.Transfer{}
	for rows.Next() {
		var transfer models.Transfer
		if err := rows.Scan(&transfer.ID, &transfer.FromAccountID, &transfer.ToAccountID, &transfer.Amount); err != nil {
			return nil, err
		}
		transfers = append(transfers, transfer)
	}
	return transfers, nil
}
