package handler

import (
	"database/sql"
	"errors"
	"ftgo-8-phase-1-pair-project/entity"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockDB struct {
	users map[string]string
}

func (m *mockDB) Exec(query string, args ...interface{}) (sql.Result, error) {
	if query == "INSERT INTO users (email, password, role) VALUES ($1, $2, 'customer')" {
		email, _ := args[0].(string)
		m.users[email] = args[1].(string)
	}
	return nil, nil
}

func (m *mockDB) QueryRow(query string, args ...interface{}) *sql.Row {
	return nil
}

func TestUserRegistration(t *testing.T) {
	testCases := []struct {
		name     string
		email    string
		password string
		wantErr  bool
	}{
		{
			name:     "Valid Registration",
			email:    "test@example.com",
			password: "password",
			wantErr:  false,
		},
		{
			name:     "Empty Email",
			email:    "",
			password: "password",
			wantErr:  true,
		},
		{
			name:     "Empty Password",
			email:    "test@example.com",
			password: "",
			wantErr:  true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockDB := &mockDB{users: make(map[string]string)}

			err := registerUser(mockDB, tc.email, tc.password)

			if tc.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				_, exists := mockDB.users[tc.email]
				assert.True(t, exists)
			}
		})
	}
}

func TestUserLogin(t *testing.T) {
	testCases := []struct {
		name     string
		email    string
		password string
		wantErr  bool
	}{
		{
			name:     "Successful Login",
			email:    "existing@example.com",
			password: "correctpassword",
			wantErr:  false,
		},
		{
			name:     "Invalid Credentials",
			email:    "nonexistent@example.com",
			password: "wrongpassword",
			wantErr:  true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockDB := &mockDB{users: make(map[string]string)}
			mockDB.users["existing@example.com"] = "correctpassword"

			user, err := loginUser(mockDB, tc.email, tc.password)

			if tc.wantErr {
				assert.Error(t, err)
				assert.Nil(t, user)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, user)
				assert.Equal(t, tc.email, user.Email)
			}
		})
	}
}

func registerUser(db *mockDB, email, password string) error {
	if email == "" || password == "" {
		return errors.New("email and password cannot be empty")
	}
	_, err := db.Exec(
		"INSERT INTO users (email, password, role) VALUES ($1, $2, 'customer')",
		email, password)
	return err
}

func loginUser(db *mockDB, email, password string) (*entity.User, error) {
	storedPassword, exists := db.users[email]
	if !exists {
		return nil, errors.New("user not found")
	}
	if storedPassword != password {
		return nil, errors.New("invalid password")
	}

	return &entity.User{
		Email: email,
		Role:  "customer",
	}, nil
}
