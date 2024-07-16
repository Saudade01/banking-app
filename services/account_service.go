package services

import (
	"banking-app/config"
	"banking-app/models"
	"database/sql"
	"errors"
)

func CreateAccount(account models.Account) (int64, error) {
	var existingAccount models.Account
	query := "SELECT id FROM accounts WHERE owner = ?"
	err := config.DB.QueryRow(query, account.Owner).Scan(&existingAccount.ID)
	if err == nil {
		return 0, errors.New("owner already exists")
	}
	if err != sql.ErrNoRows {
		return 0, err
	}

	query = "INSERT INTO accounts (owner, balance, currency) VALUES (?, ?, ?)"
	result, err := config.DB.Exec(query, account.Owner, account.Balance, account.Currency)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

func GetAccount(id int64) (*models.Account, error) {
	account := &models.Account{}
	query := "SELECT id, owner, balance, currency FROM accounts WHERE id = ?"
	err := config.DB.QueryRow(query, id).Scan(&account.ID, &account.Owner, &account.Balance, &account.Currency)
	if err != nil {
		return nil, err
	}
	return account, nil
}

func ListAccounts() ([]models.Account, error) {
	rows, err := config.DB.Query("SELECT id, owner, balance, currency FROM accounts")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	accounts := []models.Account{}
	for rows.Next() {
		var account models.Account
		if err := rows.Scan(&account.ID, &account.Owner, &account.Balance, &account.Currency); err != nil {
			return nil, err
		}
		accounts = append(accounts, account)
	}
	return accounts, nil
}
