package middleware

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"go-postgres/models"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

// createConnection uses database/sql to establish a database connection
func createConnection() *sql.DB {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	db, err := sql.Open("postgres", os.Getenv("POSTGRES_CONNECTION_STRING"))
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Succesfully created connection!")

	return db
}

// getUser fetches one user from the database which matches the id
func getUser(id string) (models.User, error) {
	db := createConnection()
	defer db.Close()

	user := models.User{}

	sqlQuery := "SELECT * FROM users WHERE id=$1"

	row := db.QueryRow(sqlQuery, id)
	err := row.Scan(&user.ID, &user.Name, &user.Age)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return user, nil
	case nil:
		return user, nil
	default:
		log.Fatalf("Unable to scan the row. %v", err)
	}

	return user, err
}

// GetUser will return a single user from the id
func GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	// UUID will be a string
	id := params["id"]

	user, err := getUser(id)

	if err != nil {
		log.Fatalf("Unable to get the user : %v!", err)
	}

	json.NewEncoder(w).Encode(user)
}

// getAllusers fetches all users from the database
func getAllUsers() ([]models.User, error) {
	db := createConnection()
	defer db.Close()

	users := []models.User{}

	sqlQuery := "SELECT * FROM users"

	rows, err := db.Query(sqlQuery)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	defer rows.Close()

	// rows.Next() returns a new boolean for no row present
	for rows.Next() {
		var user models.User

		// unmarshal the row object to user
		err = rows.Scan(&user.ID, &user.Name, &user.Age)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		// append the user in the users slice
		users = append(users, user)
	}

	return users, err
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := getAllUsers()

	if err != nil {
		log.Fatalf("Unable to get the all users : %v!", err)
	}

	json.NewEncoder(w).Encode(users)
}
