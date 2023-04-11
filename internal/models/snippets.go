package models

import (
	"database/sql"
	"time"
	"errors"
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


func (m *SnippetModel) Get (id int) (*Snippet, error){
	stmt := `SELECT id, title, content, created, expires FROM snippets
	WHERE expires > UTC_TIMESTAMP() AND id=?`

	row := m.DB.QueryRow(stmt,id)//QueryRow with stmt, id
	//Initialise new zeroed snippet struct
	
	s := &Snippet{}

	err := row.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
	if err != nil{
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoRecord
		} else {
			return nil, err
		}

	}
	return s, nil
}

func (m *SnippetModel) Latest() ([]*Snippet, error) {

	stmt := `SELECT id, title, content, created, expires FROM snippets
	WHERE expires > UTC_TIMESTAMP() ORDER BY id DESC LIMIT 10`

	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil,err
	}

	//defer is called after error check from Query to avoid tryna close a nil resultset resulting in a panic
	//as long as resultset is open it keeps the underlying database connection open
	defer rows.Close()

	//new empty slice to hold snippets
	snippets := []*Snippet{}

	for rows.Next() {
		s := &Snippet{}
		err = rows.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
		if err != nil {
			return nil, err
		}
		snippets = append(snippets, s)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}


	return snippets, nil
}