package db

import (
	"context"
	"testing"

	_ "github.com/lib/pq"
)

func createRandomUser(t *testing.T) User {
	arg := CreateUserParams{
		Username:       "test",
		HashedPassword: "testpassword",
		FirstName:      "test",
		SecondName:     "name",
		Email:          "testmail@test.com",
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	if err != nil {
		t.Fatalf("Failed to create user: %v", err)
	}

	if user.Email != arg.Email {
		t.Fatalf("Emails do not match: %s", user.Email)
	}

	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	user := createRandomUser(t)

	user, err := testQueries.GetUser(context.Background(), user.Username)
	if err != nil {
		t.Fatalf("Failed to get user: %v", user.Username)
	}
}

func TestUpdateUser(t *testing.T) {
	user := createRandomUser(t)

	newUser, err := testQueries.UpdateUser(context.Background(), UpdateUserParams{
		ID:             user.ID,
		Username:       user.Username,
		Email:          user.Email,
		FirstName:      "NEW FIRSTNAME",
		SecondName:     "NEW SECONDNAME",
		HashedPassword: user.HashedPassword,
	})

	if err != nil {
		t.Fatalf("Error updating user %v", user.Email)
	}

	if user.FirstName == newUser.FirstName {
		t.Fatalf("Error updating users firstname %v, should be %v", user.FirstName, newUser.FirstName)
	}
}
