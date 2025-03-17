package repository

import (
	"BankRaya/entity"
	"database/sql"
)

type ItemRepository struct{
	DB *sql.DB
}

func NewItemRepository(db *sql.DB) *ItemRepository {
	return &ItemRepository{DB: db}
}

func (pr *ItemRepository) FindAll() ([]entity.Item, error) {
	rows, err := pr.DB.Query("SELECT * FROM ITEM")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items := []entity.Item{}
	for rows.Next() {
		var item entity.Item
		err := rows.Scan(&item.ItemCode, &item.ItemName, &item.ItemQty)
		if err != nil {
			return nil, err
		}

		items = append(items, item)
	}

	return items, nil
}