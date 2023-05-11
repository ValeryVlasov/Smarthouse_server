package repository

import (
	"fmt"
	"github.com/ValeryVlasov/Smarthouse_server"
	"github.com/jmoiron/sqlx"
	"strings"
)

type DeviceDetectorPostgres struct {
	db *sqlx.DB
}

func NewDeviceDetectorPostgres(db *sqlx.DB) *DeviceDetectorPostgres {
	return &DeviceDetectorPostgres{db: db}
}

func (r *DeviceDetectorPostgres) Create(userId int, detector Smarthouse_server.DeviceDetector) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var detectorId int
	createDetectorQuery := fmt.Sprintf("INSERT INTO %s (name, place) values ($1, $2) RETURNING id", deviceDetectorsTable)

	row := tx.QueryRow(createDetectorQuery, detector.Name, detector.Place)
	err = row.Scan(&detectorId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	createListDetectorsQuery := fmt.Sprintf("INSERT INTO %s (user_id, detector_id) values ($1, $2)", usersDetectorsTable)
	_, err = tx.Exec(createListDetectorsQuery, userId, detectorId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return detectorId, tx.Commit()
}

func (r *DeviceDetectorPostgres) GetAll(userId int) ([]Smarthouse_server.DeviceDetector, error) {
	var detectors []Smarthouse_server.DeviceDetector
	query := fmt.Sprintf(`SELECT dd.id, dd.name, dd.place, dd.statement FROM %s dd INNER JOIN %s ud on dd.id = ud.detector_id WHERE ud.user_id = $1`,
		deviceDetectorsTable, usersDetectorsTable)
	if err := r.db.Select(&detectors, query, userId); err != nil {
		return nil, err
	}

	return detectors, nil
}

func (r *DeviceDetectorPostgres) GetById(userId, detectorId int) (Smarthouse_server.DeviceDetector, error) {
	var detector Smarthouse_server.DeviceDetector
	query := fmt.Sprintf(`SELECT dd.id, dd.name, dd.place, dd.statement FROM %s dd INNER JOIN %s ud on dd.id = ud.detector_id WHERE ud.id = $1 AND ud.user_id = $2`,
		deviceDetectorsTable, usersDetectorsTable)
	if err := r.db.Get(&detector, query, detectorId, userId); err != nil {
		return detector, err
	}

	return detector, nil
}

func (r *DeviceDetectorPostgres) Delete(userId, detectorId int) error {
	query := fmt.Sprintf(`DELETE FROM %s dd USING %s ud WHERE dd.id = ud.detector_id AND ud.user_id = $1 AND dd.id = $2`,
		deviceDetectorsTable, usersDetectorsTable)
	_, err := r.db.Exec(query, userId, detectorId)
	return err
}

func (r *DeviceDetectorPostgres) Update(userId, detectorId int, input Smarthouse_server.UpdateDetectorInput) error {
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

	if input.Statement != nil {
		setValues = append(setValues, fmt.Sprintf("statement=$%d", argId))
		args = append(args, *input.Statement)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf(`UPDATE %s dd SET %s FROM %s ud
									WHERE dd.id = ud.detector_id AND ud.user_id = $%d AND dd.id = $%d`,
		deviceDetectorsTable, setQuery, usersDetectorsTable, argId, argId+1)
	args = append(args, userId, detectorId)

	_, err := r.db.Exec(query, args...)
	return err
}
