package database

import (
	"database/sql"
	"errors"
	"todo-gotth/models"

	"github.com/mattn/go-sqlite3"
)

var (
	ErrDuplicate    = errors.New("record already exists")
	ErrNotExists    = errors.New("record does not exist")
	ErrUpdateFailed = errors.New("update failed")
	ErrDelteFailed  = errors.New("delete failed")
)

type SQLiteDatabase struct {
	db *sql.DB
}

func NewSQLiteDatabase(db *sql.DB) *SQLiteDatabase {
	return &SQLiteDatabase{
		db: db,
	}
}

func (srv *SQLiteDatabase) Migrate() error {
	query := `
		CREATE TABLE IF NOT EXISTS todos(
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			description TEXT NOT NULL UNIQUE,
			completed BOOLEAN NOT NULL DEFAULT FALSE
	);
	`
	_, err := srv.db.Exec(query)

	return err
}

func (srv *SQLiteDatabase) Create(todo models.ToDo) (*models.ToDo, error) {
	res, err := srv.db.Exec("INSERT INTO todos(description, completed) values(? ?)", todo.Description, todo.Completed)
	if err != nil {
		var sqliteErr sqlite3.Error

		if errors.As(err, &sqliteErr) {
			if errors.Is(sqliteErr.ExtendedCode, sqlite3.ErrConstraintUnique) {
				return nil, ErrDuplicate
			}
		}
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	todo.ID = id

	return &todo, nil
}


func (srv *SQLiteDatabase) All() ([]models.ToDo, error) {
	rows, err := srv.db.Query("SELECT * FROM todos")

	if err != nil {
			return nil, err
	}

	defer rows.Close()

	var all []models.ToDo
	for rows.Next(){
		var todo models.ToDo
		if err := rows.Scan(&todo.ID, &todo.Description, &todo.Completed); err != nil {
			return nil, err
		}
		all = append(all, todo)
	}

	return all, nil
}


func (srv *SQLiteDatabase) GetByID(id int64) (*models.ToDo, error){
	res := srv.db.QueryRow("SELECT * FROM todos WHERE id = ?", id)
	
	var todo models.ToDo

	if err := res.Scan(&todo.ID, &todo.Description, &todo.Completed); err != nil {
		if errors.Is(err, sql.ErrNoRows){
			return nil, ErrNotExists
		}
		return nil, err
	}
	
	return &todo, nil
}


func (srv *SQLiteDatabase) Update(id int64, update models.ToDo) (*models.ToDo, error) {
	
	if id == 0 {
		return nil, errors.New("invalid id")
	}

	res, err := srv.db.Exec("Update todos SET description = ?, completed = ? WHERE id = ?", update.Description, update.Completed, update.ID) 
	if err != nil {
		return nil, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}
	
	if rowsAffected == 0 {
		return nil, ErrUpdateFailed
	}

	return &update, nil

}

func (srv *SQLiteDatabase) Delete(id int64) error {
	res, err := srv.db.Exec("DELTE FROM todos WHERE id = ?", id)

	if err != nil {
		return err
	}

	rowAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowAffected == 0 {
		return ErrDelteFailed
	}

	return err
}
