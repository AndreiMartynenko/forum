package main

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

//     _________users_________________________________________________
//     |  id      |  email    |  username  |  password  |  sessionId  |
//     |  INTEGER |  TEXT     |  TEXT      |  TEXT      |  TEXT       |

func createUsersTable() error {

	users_table := `CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY,
        username TEXT NOT NULL UNIQUE,
		email TEXT NOT NULL UNIQUE,
        password TEXT NOT NULL,
		sessionID TEXT,
		);`

	statement, err := db.Prepare(users_table)
	if err != nil {
		// fmt.Println(err)
		// return
		return err
	}
	defer statement.Close()
	statement.Exec()
	return nil
}

// We are passing db reference connection from main to our method with other parameters
func insertUser(username string, email string, password string, sessionId string) error {
	//log.Println("Inserting user record ...")
	insertUser := `INSERT INTO users(username, email, password, sessionID) VALUES (?, ?, ?, ?)`
	statement, err := db.Prepare(insertUser) // Prepare statement.
	// This is good to avoid SQL injections
	if err != nil {
		fmt.Println(err)
		return err
	}
	//defer statement.Close()
	emailLowCase := strings.ToLower(email)
	_, err = statement.Exec(emailLowCase, username, email, password)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
	//defer statement.Close()
}

//Display user

// getting data from table
func displayUsers() {
	//row, err := db.Query("SELECT * FROM users ORDER BY name")
	row, err := db.Query("SELECT * FROM users")
	if err != nil {
		log.Fatal(err)
		//return
	}
	defer row.Close()
	/*
		The row variable is closed using the defer statement,
		which ensures that the row.Close() method is called at
		the end of the function to close the result cursor and
		free up resources.
	*/
	for row.Next() { // Iterate and fetch the records from result cursor
		user := User{}
		err = row.Scan(&(user.Id), &(user.Email), &(user.Username), &(user.Password), &(user.SessionId))
		/*
		   Inside the loop, a User struct is created to store the data of the current row.
		   The row.Scan() method is used to scan the values from the current row into
		   the fields of the User struct, which represents the columns of the users table
		   in the same order. The &user.ID, &user.Username, &user.Email, and &user.Password
		   are pointers to the fields of the User struct, where the corresponding
		   column values from the current row are scanned.
		*/
		if err != nil {
			log.Fatal(err)
			//row.Scan(&username, &email, &password)
			fmt.Println("User: ", user.Username, " ", user.Email, " ", user.Password)
		}
	}
}