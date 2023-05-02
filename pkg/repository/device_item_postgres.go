package repository

import (
	"fmt"
	"github.com/ValeryVlasov/Smarthouse_server"
	"github.com/jmoiron/sqlx"
	"strings"
)

type DeviceItemPostgres struct {
	db *sqlx.DB
}

func NewDeviceItemPostgres(db *sqlx.DB) *DeviceItemPostgres {
	return &DeviceItemPostgres{db: db}
}

func (r *DeviceItemPostgres) Create(listId int, item Smarthouse_server.DeviceItem) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var itemId int
	createItemQuery := fmt.Sprintf("INSERT INTO %s (title, description) values ($1, $2) RETURNING id", deviceItemsTable)

	row := tx.QueryRow(createItemQuery, item.Title, item.Description)
	err = row.Scan(&itemId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	createListItemsQuery := fmt.Sprintf("INSERT INTO %s (list_id, item_id) values ($1, $2)", listsItemsTable)
	_, err = tx.Exec(createListItemsQuery, listId, itemId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return itemId, tx.Commit()
}

func (r *DeviceItemPostgres) GetAll(userId, listId int) ([]Smarthouse_server.DeviceItem, error) {
	var items []Smarthouse_server.DeviceItem
	query := fmt.Sprintf(`SELECT di.id, di.title, di.description, di.isPowerOn FROM %s di INNER JOIN %s li on li.item_id = di.id
									INNER JOIN %s ul on ul.list_id = li.list_id WHERE li.list_id = $1 AND ul.user_id = $2`,
		deviceItemsTable, listsItemsTable, usersListsTable)
	if err := r.db.Select(&items, query, listId, userId); err != nil {
		return nil, err
	}

	return items, nil
}

func (r *DeviceItemPostgres) GetById(userId, itemId int) (Smarthouse_server.DeviceItem, error) {
	var item Smarthouse_server.DeviceItem
	query := fmt.Sprintf(`SELECT di.id, di.title, di.description, di.isPowerOn FROM %s di INNER JOIN %s li on li.item_id = di.id
									INNER JOIN %s ul on ul.list_id = li.list_id WHERE di.id = $1 AND ul.user_id = $2`,
		deviceItemsTable, listsItemsTable, usersListsTable)
	if err := r.db.Get(&item, query, itemId, userId); err != nil {
		return item, err
	}

	return item, nil
}

func (r *DeviceItemPostgres) Delete(userId, itemId int) error {
	query := fmt.Sprintf(`DELETE FROM %s di USING %s li, %s ul 
									WHERE di.id = li.item_id AND li.list_id = ul.list_id AND ul.user_id = $1 AND di.id = $2`,
		deviceItemsTable, listsItemsTable, usersListsTable)
	_, err := r.db.Exec(query, userId, itemId)
	return err
}

func (r *DeviceItemPostgres) Update(userId, itemId int, input Smarthouse_server.UpdateItemInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *input.Title)
		argId++
	}

	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, *input.Description)
		argId++
	}

	if input.IsPowerOn != nil {
		setValues = append(setValues, fmt.Sprintf("done=$%d", argId))
		args = append(args, *input.IsPowerOn)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf(`UPDATE %s di SET %s FROM %s li, %s ul
									WHERE di.id = li.item_id AND li.list_id = ul.list_id AND ul.user_id = $%d AND di.id = $%d`,
		deviceItemsTable, setQuery, listsItemsTable, usersListsTable, argId, argId+1)
	args = append(args, userId, itemId)

	_, err := r.db.Exec(query, args...)
	return err
}
