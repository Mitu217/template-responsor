package repository

import (
	"database/sql"
	"errors"
	"log"

	"github.com/Mitu217/template-responsor/repository/model"
	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
)

type SQLiteResponseRepository struct {
	path string
	db   *sql.DB
}

func NewSQLiteResponseRepository(p string) (*SQLiteResponseRepository, error) {
	model := &SQLiteResponseRepository{
		path: p,
		db:   nil,
	}

	if err := model.init(); err != nil {
		return nil, err
	}

	return model, nil
}

func (t *SQLiteResponseRepository) open() error {
	if t.db != nil {
		return errors.New("already open database")
	}
	db, err := sql.Open("sqlite3", t.path)
	t.db = db
	return err
}

func (t *SQLiteResponseRepository) close() error {
	if t.db == nil {
		return errors.New("not open database")
	}
	err := t.db.Close()
	t.db = nil
	return err
}

func (t *SQLiteResponseRepository) init() error {
	if err := t.open(); err != nil {
		return err
	}
	defer t.close()

	query := `
		create table if not exists responses (uuid text, name text, access integer, encoding integer, route text, data text)
	`
	if _, err := t.db.Exec(query); err != nil {
		return err
	}
	return nil
}

func (t *SQLiteResponseRepository) Gets() ([]*model.Response, error) {
	if err := t.open(); err != nil {
		return nil, err
	}
	defer t.close()

	query := `
		select uuid, name, access, encoding, route, data from responses
	`
	rows, err := t.db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	ress := []*model.Response{}
	for rows.Next() {
		res := &model.Response{}
		err = rows.Scan(&res.Uuid, &res.Name, &res.RequestMethod, &res.ContentType, &res.Route, &res.Data)
		if err != nil {
			return nil, err
		}
		ress = append(ress, res)
	}
	return ress, nil
}

func (t *SQLiteResponseRepository) Get(uuid string) (*model.Response, error) {
	return nil, nil
}

func (t *SQLiteResponseRepository) Create(response *model.Response) error {
	if err := t.open(); err != nil {
		return err
	}
	defer t.close()

	query := `
		insert into responses values(?, ?, ?, ?, ?, ?)
	`
	_, err := t.db.Exec(query, uuid.New(), response.Name, response.RequestMethod, response.ContentType, response.Route, response.Data)
	return err
}

func (t *SQLiteResponseRepository) Update(response *model.Response) error {
	return nil
}

func (t *SQLiteResponseRepository) Delete(uuid string) error {
	return nil
}
