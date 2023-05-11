package repository

import (
	"fmt"
	"github.com/ValeryVlasov/Smarthouse_server"
	"github.com/jmoiron/sqlx"
	"strings"
)

type DeviceLightPostgres struct {
	db *sqlx.DB
}

func NewDeviceLightPostgres(db *sqlx.DB) *DeviceLightPostgres {
	return &DeviceLightPostgres{db: db}
}

func (r *DeviceLightPostgres) Create(userId int, light Smarthouse_server.DeviceLight) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}
	var lightId int
	createLightQuery := fmt.Sprintf("INSERT INTO %s (name, place) values ($1, $2) RETURNING id", deviceLightsTable)

	row := tx.QueryRow(createLightQuery, light.Name, light.Place)
	err = row.Scan(&lightId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	createListLightsQuery := fmt.Sprintf("INSERT INTO %s (user_id, light_id) VALUES ($1, $2)", usersLightsTable)
	_, err = tx.Exec(createListLightsQuery, userId, lightId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return lightId, tx.Commit()
}

func (r *DeviceLightPostgres) GetAll(userId int) ([]Smarthouse_server.DeviceLight, error) {
	var lights []Smarthouse_server.DeviceLight
	query := fmt.Sprintf(`SELECT dl.id, dl.name, dl.place, dl.condition FROM %s dl INNER JOIN %s ul on dl.id = ul.light_id WHERE ul.user_id = $1`,
		deviceLightsTable, usersLightsTable)
	if err := r.db.Select(&lights, query, userId); err != nil {
		return nil, err
	}

	return lights, nil
}

func (r *DeviceLightPostgres) GetById(userId, lightId int) (Smarthouse_server.DeviceLight, error) {
	var light Smarthouse_server.DeviceLight
	query := fmt.Sprintf(`SELECT dl.id, dl.name, dl.place, dl.condition FROM %s dl INNER JOIN %s ul on dl.id = ul.light_id WHERE ul.id = $1 AND ul.user_id = $2`,
		deviceLightsTable, usersLightsTable)
	if err := r.db.Get(&light, query, lightId, userId); err != nil {
		return light, err
	}

	return light, nil
}

func (r *DeviceLightPostgres) Delete(userId, lightId int) error {
	query := fmt.Sprintf(`DELETE FROM %s dl USING %s ul WHERE dl.id = ul.light_id AND ul.user_id = $1 AND dl.id = $2`,
		deviceLightsTable, usersLightsTable)
	_, err := r.db.Exec(query, userId, lightId)
	return err
}

func (r *DeviceLightPostgres) Update(userId, lightId int, input Smarthouse_server.UpdateLightInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Name != nil {
		setValues = append(setValues, fmt.Sprintf("name=$%d", argId))
		args = append(args, *input.Name)
		argId++
	}

	if input.Place != nil {
		setValues = append(setValues, fmt.Sprintf("place=$%d", argId))
		args = append(args, *input.Place)
		argId++
	}

	if input.Condition != nil {
		setValues = append(setValues, fmt.Sprintf("condition=$%d", argId))
		args = append(args, *input.Condition)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf(`UPDATE %s dl SET %s FROM %s ul
									WHERE dl.id = ul.light_id AND ul.user_id = $%d AND dl.id = $%d`,
		deviceLightsTable, setQuery, usersLightsTable, argId, argId+1)
	args = append(args, userId, lightId)

	_, err := r.db.Exec(query, args...)
	return err
}
