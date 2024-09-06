package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {

	client, err := NewClient("John Doe", "j@j.com")
	assert.Nil(t, err)
	assert.NotNil(t, client)
	assert.Equal(t, "John Doe", client.Name)
	assert.Equal(t, "j@j.com", client.Email)
}

func TestCreatedNewClientWhenArgsAreInvalid(t *testing.T) {
	client, err := NewClient("", "")

	assert.NotNil(t, err)
	assert.Nil(t, client)
}

func TestUpdateClient(t *testing.T) {

	client, _ := NewClient("John Doe", "j@j.com")
	err := client.Update("Jane Doe Updated", "j@j.com")

	assert.Nil(t, err)
	assert.Equal(t, "Jane Doe Updated", client.Name)
	assert.Equal(t, "j@j.com", client.Email)
}

func TestUpdateClientWithInvalidArgs(t *testing.T) {

	client, _ := NewClient("John Doe", "j@j.com")
	err := client.Update("", "j@j.com")

	assert.Error(t, err, "name is required")
}
