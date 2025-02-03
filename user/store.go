package user

import (
	"database/sql"
	"log"

	"github.com/Komilov31/weatherInfoBot/model"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetUserByName(userName string) (*model.User, error) {
	var user model.User

	row := s.db.QueryRow("SELECT * FROM test WHERE username = $1", userName)
	err := row.Scan(&user.Id, &user.UserName, &user.City)
	if err != nil {
		log.Println(err)
		return &user, err
	}

	return &user, nil
}

func (s *Store) SetLocation(userName, city string) error {
	sqlStatement := `
	INSERT INTO test (username, city)
	VALUES ($1, $2)`

	if userExists(s.db, userName) {
		sqlStatement = `
		UPDATE test
		SET city = $1
		WHERE username = $2
		`
		_, err := s.db.Exec(sqlStatement, city, userName)
		if err != nil {
			return err
			// log.Fatal("Something went wrong while inserting new user go DB")
		}

		return nil
	}

	_, err := s.db.Exec(sqlStatement, userName, city)
	if err != nil {
		return err
		// log.Fatal("Something went wrong while inserting new user go DB")
	}
	return nil
}

func userExists(db *sql.DB, userName string) bool {
	sqlStmt := `SELECT username FROM test WHERE username = $1`
	err := db.QueryRow(sqlStmt, userName).Scan(&userName)
	if err != nil {
		if err != sql.ErrNoRows {
			// a real error happened! you should change your function return
			// to "(bool, error)" and return "false, err" here
			log.Print(err)
		}

		return false
	}

	return true
}
