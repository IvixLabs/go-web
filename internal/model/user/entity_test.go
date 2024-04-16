package user

import "testing"

func TestUser(t *testing.T) {
	const testEmail = "test@email.com"
	const testPassword = "testPassword"
	const testAddress = "Test address"

	user := New(testEmail, testPassword, testAddress)

	if user.Password() != testPassword {
		t.Fatal("Wrong password")
	}

	if user.Email() != testEmail {
		t.Fatal("Wrong email")
	}

	if user.Address() != testAddress {
		t.Fatal("Wrong address")
	}

	if user.Id() == "" {
		t.Fatal("Id is empty")
	}
}
