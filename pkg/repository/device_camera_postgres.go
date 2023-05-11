package repository

import (
	"fmt"
	"github.com/ValeryVlasov/Smarthouse_server"
	"github.com/jmoiron/sqlx"
	"strings"
)

type DeviceCameraPostgres struct {
	db *sqlx.DB
}

func NewDeviceCameraPostgres(db *sqlx.DB) *DeviceCameraPostgres {
	return &DeviceCameraPostgres{db: db}
}

func (r *DeviceCameraPostgres) Create(userId int, camera Smarthouse_server.DeviceCamera) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var cameraId int
	createCameraQuery := fmt.Sprintf("INSERT INTO %s (name, place) values ($1, $2) RETURNING id", deviceCamerasTable)

	row := tx.QueryRow(createCameraQuery, camera.Name, camera.Place)
	err = row.Scan(&cameraId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	createListCamerasQuery := fmt.Sprintf("INSERT INTO %s (user_id, camera_id) values ($1, $2)", usersCamerasTable)
	_, err = tx.Exec(createListCamerasQuery, userId, cameraId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return cameraId, tx.Commit()
}

func (r *DeviceCameraPostgres) GetAll(userId int) ([]Smarthouse_server.DeviceCamera, error) {
	var cameras []Smarthouse_server.DeviceCamera
	query := fmt.Sprintf(`SELECT dc.id, dc.name, dc.place FROM %s dc INNER JOIN %s uc on dc.id = uc.camera_id WHERE uc.user_id = $1`,
		deviceCamerasTable, usersCamerasTable)
	if err := r.db.Select(&cameras, query, userId); err != nil {
		return nil, err
	}

	return cameras, nil
}

func (r *DeviceCameraPostgres) GetById(userId, cameraId int) (Smarthouse_server.DeviceCamera, error) {
	var camera Smarthouse_server.DeviceCamera
	query := fmt.Sprintf(`SELECT dc.id, dc.name, dc.place FROM %s dc INNER JOIN %s uc on dc.id = uc.camera_id WHERE uc.id = $1 AND uc.user_id = $2`,
		deviceCamerasTable, usersCamerasTable)
	if err := r.db.Get(&camera, query, cameraId, userId); err != nil {
		return camera, err
	}

	return camera, nil
}

func (r *DeviceCameraPostgres) Delete(userId, cameraId int) error {
	query := fmt.Sprintf(`DELETE FROM %s dc USING %s uc WHERE dc.id = uc.camera_id AND uc.user_id = $1 AND dc.id = $2`,
		deviceCamerasTable, usersCamerasTable)
	_, err := r.db.Exec(query, userId, cameraId)
	return err
}

func (r *DeviceCameraPostgres) Update(userId, cameraId int, input Smarthouse_server.UpdateCameraInput) error {
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

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf(`UPDATE %s dc SET %s FROM %s uc
									WHERE dc.id = uc.camera_id AND uc.user_id = $%d AND dc.id = $%d`,
		deviceCamerasTable, setQuery, usersCamerasTable, argId, argId+1)
	args = append(args, userId, cameraId)

	_, err := r.db.Exec(query, args...)
	return err
}
