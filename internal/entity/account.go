package entity

import (
	"time"

	"github.com/google/uuid"
)

type Account struct {
	id        string
	Client    *Client
	Balance   float64
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewAccount(client *Client) *Account {
	if client == nil {
		return nil
	}
	account := &Account{
		id:        uuid.New().String(),
		Client:    client,
		Balance:   0,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	return account
}

func (account *Account) Credit(amount float64) {
	account.Balance += amount
	account.UpdatedAt = time.Now()
}
