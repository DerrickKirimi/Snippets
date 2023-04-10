package models

import (
	"database/sql"
	"time"
)

type Snippet struct {
	ID		int
	Title	string
	Content string
	Created	time.Time
	Expires time.Time
}

//SnippetModel wraps a sql.DB connection pool
type SnippetModel struct {
	DB *sql.DB
}

//stmt split into two lines which is why backquotes are used over normal double quotes
func (m *SnippetModel) Insert(title string, content string, expires int) (int, error) {
	stmt := `INSERT INTO snippets (title, content, created, expires) 
	VALUES(?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`

	result, err := m.DB.Exec(stmt, title, content, expires)
	if err != nil{
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil{
		return 0, err
	}

	return int(id), nil

}

func (m *SnippetModel) Get (id int) (*Snippet, error) {
	return nil,nil
}

func (m *SnippetModel) Latest() ([]*Snippet, error) {
	return nil, nil
}
