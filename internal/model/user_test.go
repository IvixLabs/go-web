package model

import "testing"

func TestUser(t *testing.T) {
	const testEmail = "test@email.com"
	const testPassword = "testPassword"
	const testAddress = "Test address"

	user := NewUser(testEmail, testPassword, testAddress)

	if user.GetPassword() != testPassword {
		t.Fatal("Wrong password")
	}

	if user.GetEmail() != testEmail {
		t.Fatal("Wrong email")
	}

	if user.GetAddress() != testAddress {
		t.Fatal("Wrong address")
	}

	if user.GetId() == "" {
		t.Fatal("Id is empty")
	}
}
