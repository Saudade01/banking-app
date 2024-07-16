package main

import (
	"banking-app/config"
	"banking-app/models"
	"banking-app/services"
	"math/rand"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	// Testler için gerekli başlangıç ayarları
	config.LoadConfig()
	config.InitDB()
	rand.Seed(time.Now().UnixNano()) // Rastgelelik için tohum

	// Testleri çalıştır
	code := m.Run()

	// Veritabanı bağlantısını kapat
	if config.DB != nil {
		config.DB.Close()
	}

	os.Exit(code)
}

func TestCreateRandomUsers(t *testing.T) {
	for i := 0; i < 5; i++ {
		user := models.Account{
			Owner:    "User" + strconv.Itoa(i) + strconv.Itoa(rand.Intn(1000)), // Benzersiz kullanıcı adı
			Balance:  rand.Float64() * 1000,
			Currency: "USD",
		}

		id, err := services.CreateAccount(user)
		assert.Nil(t, err)
		t.Logf("Created account with ID: %d", id)
	}
}

func TestRandomTransfers(t *testing.T) {
	// Kullanıcıları oluştur
	var accountIDs []int64
	for i := 0; i < 5; i++ {
		user := models.Account{
			Owner:    "User-" + strconv.Itoa(i) + strconv.Itoa(rand.Intn(1000)), // Benzersiz kullanıcı adı
			Balance:  rand.Float64() * 1000,
			Currency: "USD",
		}
		id, err := services.CreateAccount(user)
		assert.Nil(t, err)
		accountIDs = append(accountIDs, id)
	}

	// Rastgele transferler yap
	for i := 0; i < 5; i++ {
		fromIndex := rand.Intn(len(accountIDs))
		toIndex := rand.Intn(len(accountIDs))
		for fromIndex == toIndex {
			toIndex = rand.Intn(len(accountIDs))
		}

		fromAccount, err := services.GetAccount(accountIDs[fromIndex])
		assert.Nil(t, err)
		toAccount, err := services.GetAccount(accountIDs[toIndex])
		assert.Nil(t, err)

		amount := rand.Float64() * 100
		if fromAccount.Balance < amount {
			t.Logf("Insufficient funds for transfer from account ID: %d to account ID: %d", fromAccount.ID, toAccount.ID)
			continue
		}

		transfer := models.Transfer{
			FromAccountID: accountIDs[fromIndex],
			ToAccountID:   accountIDs[toIndex],
			Amount:        amount,
		}

		id, err := services.CreateTransfer(transfer)
		assert.Nil(t, err)
		t.Logf("Created transfer with ID: %d", id)
	}
}
