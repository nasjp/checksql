package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"github.com/rs/zerolog"
	sqldblogger "github.com/simukti/sqldb-logger"
	"github.com/simukti/sqldb-logger/logadapter/zerologadapter"
)

const dsn = "root:@(localhost:4529)/app?parseTime=true&multiStatements=true&charset=utf8mb4"

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

type User struct {
	ID        int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *User) String() string {
	return fmt.Sprintf(`User{
 ID: %d,
 Name: %s,
 CreatedAt: %s,
 UpdatedAt: %s,
}`, u.ID, u.Name, u.CreatedAt.UTC().Format(time.RFC3339), u.UpdatedAt.UTC().Format(time.RFC3339))
}

type Term struct {
	From time.Time `json:"from"`
	To   time.Time `json:"to"`
}

func (t *Term) String() string {
	return fmt.Sprintf(`Term{
 From: %s,
 To: %s,
}`, t.From.UTC().Format(time.RFC3339), t.To.UTC().Format(time.RFC3339))
}

func run() error {
	db, err := conn()
	if err != nil {
		return err
	}

	defer db.Close()

	jsonBytes := []byte(`
{
   "from":"2022-01-01T10:00:00+09:00",
   "to":"2022-02-01T08:59:59+09:00"
}
`)

	term := &Term{}
	if err := json.Unmarshal(jsonBytes, term); err != nil {
		return err
	}

	fmt.Println(term)

	users, err := GetUserByTerm(db, term)
	if err != nil {
		return err
	}

	for _, user := range users {
		fmt.Println(user)
	}

	return nil
}

func conn() (*sql.DB, error) {

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	db = sqldblogger.OpenDriver(
		dsn,
		db.Driver(),
		zerologadapter.New(zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout, NoColor: false})),
	)

	return db, db.Ping()
}

func GetUserByTerm(db *sql.DB, term *Term) ([]*User, error) {
	rows, err := db.Query(`SELECT * FROM users WHERE created_at BETWEEN ? AND ?`, term.From, term.To)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := make([]*User, 0)

	for rows.Next() {
		var (
			id        int
			name      string
			createdAt time.Time
			updatedAt time.Time
		)

		if err := rows.Scan(&id, &name, &createdAt, &updatedAt); err != nil {
			return nil, err
		}

		users = append(users, &User{ID: id, Name: name, CreatedAt: createdAt, UpdatedAt: updatedAt})
	}

	return users, nil
}
