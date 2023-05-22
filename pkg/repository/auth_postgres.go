package repository

import (
	"fmt"
	"github.com/ValeryVlasov/Smarthouse_server"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/cast"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) IsSameUser(login, password interface{}) bool {
	var user Smarthouse_server.User2

	fmt.Println(r.db.Get(&user, "SELECT * FROM users WHERE id=1"))
	if err := r.db.Get(&user, "SELECT * FROM users WHERE id=1"); err != nil {
		return false
	}
	if user.Password_hash != password || user.Username != login {
		fmt.Println("dbPass = " + user.Password_hash + ", realPass = " + cast.ToString(password) + ", dbLog = " + user.Username + ", realLog = " + cast.ToString(login))
		return false
	}
	return true
}

func (r *AuthPostgres) CreateUser(user Smarthouse_server.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) values ($1, $2, $3) RETURNING id", usersTable)
	row := r.db.QueryRow(query, user.Name, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *AuthPostgres) GetUser(username, password string) (Smarthouse_server.User, error) {
	var user Smarthouse_server.User
	query := fmt.Sprintf("SELECT * FROM %s WHERE username=$1 AND password_hash=$2", usersTable)
	err := r.db.Get(&user, query, username, password)

	return user, err
}

func (r *AuthPostgres) GetUser2(username, password string) (Smarthouse_server.User2, error) {
	var user Smarthouse_server.User2
	query := fmt.Sprintf("SELECT * FROM %s WHERE username=$1 AND password_hash=$2", usersTable)
	err := r.db.Get(&user, query, username, password)

	return user, err
}
