package middleware

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"go-postgres/models"
	"log"
	"net/http"
	"os"

	"github.com/google/uuid"
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

type response struct {
	ID      uuid.UUID `json:"id,omitempty"`
	Message string    `json:"message,omitempty"`
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	user := models.User{}

	// Decode the user from the request body
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Fatalf("Error reading user from request: %v", err)
	}

	insertID, err := createUser(user)
	if err != nil {
		log.Fatalf("Error inserting user: %v", err)
	}

	res := response{
		ID:      insertID,
		Message: "User inserted successfully!",
	}

	json.NewEncoder(w).Encode(res)
}

func createUser(user models.User) (uuid.UUID, error) {
	db := createConnection()
	defer db.Close()

	sqlQuery := "INSERT INTO users (id, name, age) VALUES ($1, $2, $3) RETURNING id"

	var id uuid.UUID

	err := db.QueryRow(sqlQuery, user.ID, user.Name, user.Age).Scan(&id)
	if err != nil {
		log.Fatalf("Error inserting user!: %v", err)
		// return Nil uuid
		return uuid.Nil, err
	}

	log.Printf("Inserted a single record with id: %v", id)

	return id, nil
}

// Update user with the given ID
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id := params["id"]

	user := models.User{}

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Fatalf("Error decoding user from request: %v", err)
	}

	rowsAffected, err := updateUser(id, user)
	if err != nil {
		log.Fatalf("Error updating the user: %v", err)
	}

	updateUUID, _ := uuid.Parse(id)

	res := response{
		ID:      updateUUID,
		Message: fmt.Sprintf("Updated rows successfully: %v", rowsAffected),
	}

	json.NewEncoder(w).Encode(res)
}

func updateUser(ID string, user models.User) (int64, error) {
	db := createConnection()
	defer db.Close()

	sqlQuery := `UPDATE users SET name=$2, age=$3 WHERE id=$1`

	res, err := db.Exec(sqlQuery, ID, user.Name, user.Age)
	if err != nil {
		log.Fatalf("Error updating user!: %v", err)
		// return Nil uuid
		return 0, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Fatalf("Error while checking the affected rows: %v", err)
	}

	fmt.Printf("Total rows/record affected %v", rowsAffected)

	return rowsAffected, nil
}

// Delete user with the given ID
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id := params["id"]

	rowsAffected, err := deleteUser(id)
	if err != nil {
		log.Fatalf("Error deleting the user: %v", err)
	}

	deleteUUID, _ := uuid.Parse(id)

	res := response{
		ID:      deleteUUID,
		Message: fmt.Sprintf("Deleted rows successfully: %v", rowsAffected),
	}

	json.NewEncoder(w).Encode(res)
}

func deleteUser(ID string) (int64, error) {
	db := createConnection()
	defer db.Close()

	sqlQuery := `DELETE FROM users WHERE id=$1`

	res, err := db.Exec(sqlQuery, ID)
	if err != nil {
		log.Fatalf("Error deleting user!: %v", err)
		// return Nil uuid
		return 0, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Fatalf("Error while checking the affected rows: %v", err)
	}

	fmt.Printf("Total rows/record affected %v", rowsAffected)

	return rowsAffected, nil
}
