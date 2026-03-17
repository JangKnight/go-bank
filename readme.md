# Go Bank

## EZ start

1. Clone the repository
2. Navigate to the project directory
3. run `docker compose up --build`
4. Open your browser and go to `http://localhost:8000` to see the application in action.

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
