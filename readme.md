# Go Bank

## EZ start

1. Clone the repository
2. Run `docker-compose up -d` to start the PostgreSQL database
3. Run `go run cmd/api/main.go` to start the API server

<br>

OR

<br>

3. `go install github.com/air-verse/air@latest` then run `air` to start the API server with live reload

4. Access the application at `http://localhost:8080`

## Endpoints

| Method | Endpoint | Description          |
| ------ | -------- | -------------------- |
| POST   | /signup  | Register a new user  |
| POST   | /login   | Authenticate a user  |
| GET    | /signup  | Show the signup form |
| GET    | /        | Homepage             |

## Structure

````lua
├── cmd
│   └── api
├── docker-compose.yml
├── favicon.ico
├── go-bank
├── go.mod
├── go.sum
├── go_bank.txt
├── internal
│   ├── api
│   ├── auth
│   ├── crypto
│   ├── db
│   ├── domain
│   ├── repository
│   └── service
├── migrations
│   ├── 20260305153001_create_users_table.sql
│   ├── 20260305153919_create_accounts_table.sql
│   └── migrations.go
├── templates
│   ├── home.html
│   ├── profile.html
│   └── signup.html
├── tests
│   └── api
└── tmp
    ├── build-errors.log
    └── main
    ```
````
