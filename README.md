# go-postgres

This is a template example for a CRUD appication (REST API) in Go with Postgres as a database authored by [ojas](https://github.com/ojasww).

Following is the API spec of the application:

| API            | Method | Description                         |
| -------------- | ------ | ----------------------------------- |
| /api/user/{id} | GET    | Get a user given an uuid            |
| /api/users     | GET    | Get all users                       |
| /api/newuser   | PUT    | Create a new user given the details |
| /api/user/{id} | UPDATE | Update a user given an uuid         |
| /api/user/{id} | DELETE | Delete a user given an uuid         |


## Third party packages

The following go third-party packages are used:

- [Gorilla Mux](https://github.com/gorilla/mux) - Go router for incoming HTTP requests.
- [Google UUID](https://pkg.go.dev/github.com/google/uuid) - Google's well maintained UUID golang package.
- [Godotenv](https://github.com/joho/godotenv) - Pakcage to load local environment files.
- [Postgres Golang Driver](https://github.com/lib/pq) - Go driver for Postgres. Check [here](https://go.dev/wiki/SQLDrivers) for a full list of drivers. 


## Model

The Postgres database contains a single table called `users`:

```go
type User struct {
    ID   uuid.UUID `json:"id"`
    Name string    `json:"name"`
    Age  int64     `json:"age"`
}
```

## Testing

Any testing is done through [Curl](https://curl.se/). A sample test for creating a new user is as follows API handler:

```bash
curl localhost:8080/api/newuser -d '{"id":"0a99e24d-ae00-49a9-8faa-c6c07ee70384","name":"mickey","age":50}' -X POST
```
Response:

```bash
{"id":"0a99e24d-ae00-49a9-8faa-c6c07ee70384","message":"User inserted successfully!"}
```

## Commits

The following table is the commit-by-commit explanation of the development:

| Commit                                                                                            | Message                                    |
| ------------------------------------------------------------------------------------------------- | ------------------------------------------ |
| [commit 1](https://github.com/ojasww/go-postgres/commit/f4bd2b298607c59e935c0daa6b505837b13b8add) | Ping!                                      |
| [commit 2](https://github.com/ojasww/go-postgres/commit/b2361575426d5bb3a8fd39b0468824a41bbed552) | chore: add file structure/intialize router |
| [commit 3](https://github.com/ojasww/go-postgres/commit/66fed40e68abef31fbe5aafb3244aa42f172109b) | feat: add users model and GetUsers API     |
| [commit 4](https://github.com/ojasww/go-postgres/commit/8115b1460c44364d4c46ec2af6b13a24b7131b0f) | feat: add getAllUsers API                  |
| [commit 5](https://github.com/ojasww/go-postgres/commit/b984eb97bca6e3a75d5c94860b380611805ebd0a) | feat: add newuser API to create a user     |
| [commit 6](https://github.com/ojasww/go-postgres/commit/f7208e5978b0a624c14c807500a2d9ae1a26ebf0) | feat: add UpdateUsers API                  |
| [commit 7](https://github.com/ojasww/go-postgres/commit/6c7bc4a6da73526fdb43cc8a0975f2b9ab001ac1) | feat: add DeleteUsers API                  |
| [commit 8]()                                                                                      | chore: update documentation                |

## Programming paradigms:

- [conventional-commits](https://www.conventionalcommits.org/en/v1.0.0/)

## References

- [go-postgres](https://github.com/schadokar/go-postgres)
- [Golang Standard Library docs](https://pkg.go.dev/std)