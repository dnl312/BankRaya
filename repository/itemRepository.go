package repository

import (
	"BankRaya/entity"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
)

type ItemRepository struct{
	DB *sql.DB
}

func NewItemRepository(db *sql.DB) *ItemRepository {
	return &ItemRepository{DB: db}
}

func (i *ItemRepository) FindAll() ([]entity.Item, error) {
	rows, err := i.DB.Query("SELECT * FROM ITEM")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items := []entity.Item{}
	for rows.Next() {
		var item entity.Item
		err := rows.Scan(&item.ItemCode, &item.ItemName, &item.ItemPrice)
		if err != nil {
			return nil, err
		}

		items = append(items, item)
	}

	return items, nil
}

func (i *ItemRepository) Insert(Item entity.Item) error {
	_, err := i.DB.Exec("INSERT INTO item (item_code, item_name, item_price) VALUES ($1, $2, $3)",  uuid.New().String(), Item.ItemName, Item.ItemPrice)
	if err != nil {
		return fmt.Errorf("error inserting Item: %v", err)
	}
	return nil
}

func (pr *ItemRepository) Update(Item entity.Item) error {
	result, err := pr.DB.Exec("UPDATE item SET item_name=$1, item_price=$2 WHERE item_code=$3", Item.ItemName, Item.ItemPrice, Item.ItemCode)
	if err != nil {
		return fmt.Errorf("error updating Item: %v", err)
	}

	rowsAffected,_ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("Item with Code: %d not found", Item.ItemCode)
	}

	return nil
}
