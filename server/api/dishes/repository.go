package dishes

import (
	"database/sql"
	"fmt"
	"strings"
)

// Dish - model for dish
type Dish struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Price int64  `json:"price"`
}

// Store - application store struct
type Store struct {
	Db *sql.DB
}

// NewStore - creates application store connected to db
func NewStore(db *sql.DB) *Store {
	return &Store{Db: db}
}

func transformRowsToDishList(rows *sql.Rows) ([]*Dish, error) {
	if rows == nil {
		return make([]*Dish, 0), nil
	}

	defer rows.Close()

	var res []*Dish
	for rows.Next() {
		var d Dish
		if err := rows.Scan(&d.ID, &d.Name, &d.Price); err != nil {
			return nil, err
		}
		res = append(res, &d)
	}
	if res == nil {
		res = make([]*Dish, 0)
	}
	return res, nil
}

// ListDishes - get dishes from DB
func (s *Store) ListDishes() ([]*Dish, error) {
	rows, err := s.Db.Query("SELECT id, name, price FROM dishes LIMIT 200")
	if err != nil {
		return nil, err
	}

	return transformRowsToDishList(rows)
}

// GetDishesByIds - get dishesbyId
func (s *Store) GetDishesByIds(ids []int64) ([]*Dish, error) {
	idsString := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(ids)), ","), "[]")

	q := "select * from dishes where id in (" + idsString + ");"

	rows, err := s.Db.Query(q)

	if err != nil {
		return nil, err
	}

	return transformRowsToDishList(rows)
}
