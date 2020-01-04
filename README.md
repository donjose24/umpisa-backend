### Description
This is a backend api for this (app)[https://github.com/jmramos02/umpisa-frontend]. This is a rest api that allows users to login and register, topup, view balance, get account details and view transaction history. This is already running on `heroku` for demonstration purposes


### Requirements
1. `Go 1.12` or higher
2. `postgres` database setup
3. envar as: `UMPISA_DB_URL="host={dbhost} user={dbuser} dbname={dbname}sslmode=disable"` Please edit as follows and remove the curly braces


### Running
1. run `go run main.go` on your terminal
