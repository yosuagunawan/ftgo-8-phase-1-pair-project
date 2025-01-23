package handler

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockOrderDB struct {
	games  map[int]gameStock
	orders []orderRecord
}

type gameStock struct {
	stock int
	price float64
}

type orderRecord struct {
	userID   int
	gameID   int
	quantity int
	total    float64
}

func (m *mockOrderDB) Begin() (*mockTx, error) {
	return &mockTx{db: m}, nil
}

type mockTx struct {
	db *mockOrderDB
}

func (tx *mockTx) QueryRow(query string, args ...interface{}) *sql.Row {
	gameID := args[0].(int)
	_, exists := tx.db.games[gameID]
	if !exists {
		return nil
	}
	return nil
}

func (tx *mockTx) Exec(query string, args ...interface{}) (sql.Result, error) {
	if query == "INSERT INTO orders (user_id, game_id, quantity, total) VALUES ($1, $2, $3, $4)" {
		tx.db.orders = append(tx.db.orders, orderRecord{
			userID:   args[0].(int),
			gameID:   args[1].(int),
			quantity: args[2].(int),
			total:    args[3].(float64),
		})
	}
	return nil, nil
}

func (tx *mockTx) Commit() error {
	return nil
}

func (tx *mockTx) Rollback() error {
	return nil
}

func TestPlaceOrder(t *testing.T) {
	testCases := []struct {
		name           string
		userID         int
		gameID         int
		quantity       int
		availableStock int
		gamePrice      float64
		wantErr        bool
	}{
		{
			name:           "Successful Order",
			userID:         1,
			gameID:         10,
			quantity:       2,
			availableStock: 5,
			gamePrice:      49.99,
			wantErr:        false,
		},
		{
			name:           "Insufficient Stock",
			userID:         1,
			gameID:         10,
			quantity:       10,
			availableStock: 5,
			gamePrice:      49.99,
			wantErr:        true,
		},
		{
			name:           "Invalid Game ID",
			userID:         1,
			gameID:         100,
			quantity:       10,
			availableStock: 20,
			gamePrice:      55.99,
			wantErr:        true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockDB := &mockOrderDB{
				games: map[int]gameStock{
					10: {stock: tc.availableStock, price: tc.gamePrice},
				},
				orders: []orderRecord{},
			}

			err := placeOrder(mockDB, tc.userID, tc.gameID, tc.quantity)

			if tc.wantErr {
				assert.Error(t, err)
				assert.Len(t, mockDB.orders, 0)
			} else {
				assert.NoError(t, err)
				assert.Len(t, mockDB.orders, 1)
				lastOrder := mockDB.orders[0]
				assert.Equal(t, tc.userID, lastOrder.userID)
				assert.Equal(t, tc.gameID, lastOrder.gameID)
				assert.Equal(t, tc.quantity, lastOrder.quantity)
				assert.Equal(t, tc.gamePrice*float64(tc.quantity), lastOrder.total)
			}
		})
	}
}

func placeOrder(db *mockOrderDB, userID, gameID, quantity int) error {
	game, exists := db.games[gameID]
	if !exists || game.stock < quantity {
		return errors.New("order placement failed")
	}

	total := game.price * float64(quantity)

	db.orders = append(db.orders, orderRecord{
		userID:   userID,
		gameID:   gameID,
		quantity: quantity,
		total:    total,
	})

	return nil
}

func TestViewOrders(t *testing.T) {
	testCases := []struct {
		name           string
		userID         int
		existingOrders []orderRecord
		expectedCount  int
	}{
		{
			name:   "User with Multiple Orders",
			userID: 1,
			existingOrders: []orderRecord{
				{userID: 1, gameID: 100, quantity: 2, total: 99.98},
				{userID: 1, gameID: 200, quantity: 1, total: 49.99},
			},
			expectedCount: 2,
		},
		{
			name:           "User with No Orders",
			userID:         10,
			existingOrders: []orderRecord{},
			expectedCount:  0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockDB := &mockOrderDB{
				orders: tc.existingOrders,
			}

			userOrders := viewUserOrders(mockDB, tc.userID)

			assert.Len(t, userOrders, tc.expectedCount)
		})
	}
}

func viewUserOrders(db *mockOrderDB, userID int) []orderRecord {
	var userOrders []orderRecord
	for _, order := range db.orders {
		if order.userID == userID {
			userOrders = append(userOrders, order)
		}
	}
	return userOrders
}
