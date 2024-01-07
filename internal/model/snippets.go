package model

import (
	"database/sql"
	"errors"
	"fmt"
	"time"
)

type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

type SnippetModel struct {
	DB *sql.DB
}

// Insert a new snippet into the database
func (m *SnippetModel) Insert(title string, content string, expires int) (int, error) {

	//Insert query.
	//$3 is explicitly converted to an interval value by ::interval
	//If you are writing this statement in a console then write it as:
	// INSERT INTO snippets (title, content, created, expires)
	// Values ($1,$2, NOW(), NOW() + INTERVAL '$3 DAY')
	//replacing the placeholders with their respective values.
	//If we write stmt like that it will result in an error
	//pq: got 3 parameters but the statement requires 2
	//The $3 won't be recognized as a placeholder hence the statement will
	//only have 2 parameters but our DB.Exec() has three parameters
	//If we omit the single quotes, you'll get the error
	// pq: syntax error at or near "$3"

	stmt := `INSERT INTO snippets (title, content, created, expires) 
			 Values ($1,$2, NOW(), NOW() + $3::interval)`

	//Interval value
	interval := fmt.Sprintf("%d DAY", expires)

	result, err := m.DB.Exec(stmt, title, content, interval)
	if err != nil {
		return 0, err
	}

	id, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

// Get returns specific snippet based on ID
func (m *SnippetModel) Get(id int) (*Snippet, error) {
	stmt := `SELECT id, title, content, created, expires FROM snippets 
			 WHERE expires > NOW() AND id = $1`

	row := m.DB.QueryRow(stmt, id)

	s := &Snippet{}
	err := row.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoRecord
		} else {
			return nil, err
		}
	}

	return s, nil
}

// Latest fetches the 10 most recent snippet entries
func (m *SnippetModel) Latest() ([]*Snippet, error) {
	stmt := `SELECT id, title, content, created, expires FROM snippets 
			 WHERE expires > NOW() ORDER BY id DESC LIMIT 10`

	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	//Empty slice to hold the snippets struct
	var snippets []*Snippet
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
