# Project Register application

## Getting Started

### Commands

- Start the application with the current codebase
```bash
cp .env.example .env
docker compose run -p 8090:8090 -v $(pwd)/uploads:/uploads --rm go run cmd/main.go
```
Interroupt the application with `Ctrl+C`.

- install new dependencies
```bash
docker compose run --rm go get -u github.com/gorilla/mux
```

- Run the database migration
```bash
docker compose run --rm migrate -source file://db/migration -database "postgres://projectregister:password@db:5432/projectregister_development?sslmode=disable" up
```

- Drop the database
```bash
docker compose run --rm migrate -source file://db/migration -database "postgres://projectregister:password@db:5432/projectregister_development?sslmode=disable" drop -f
```

- Create a new migration file
```bash
docker compose run --rm migrate create -ext sql -dir db/migration -seq create_users_table
```

- Run the tests
```bash
docker compose run --rm go test ./...
```

- Run the tests with coverage
```bash
docker compose run --rm go test -coverprofile=coverage/coverage.out ./...
```

- generate html coverage report to the coverage directory
```bash
docker compose run --rm go tool cover -html=coverage/coverage.out -o coverage/index.html
```

### Environment variables

Add a new environment variable.

- Add it to the .env.example file.
- Add it to the .env file, or create a new one based on the .env.example file.
- Add the default value to the `config/defaults.go` file.
- Add the env name to the `config/defaults.go` file.
- Add the variable to the Environment struct in the `config/environment.go` file.
- Implement the parsing of the environment variable in the `config/environment.go` file.
- Implement the getter method in the `config/environment.go` file.

### Create migrations

Add a new migration.

- Create a new migration file in the `migration` directory. Take care for the incremental id in the filename.

### Create a new resource
- create a new migration file: copy an existing migration file and change the table name, sequence number in the filenames (up and down migrations). Write the SQL to create the new table.
- create a new model: copy an existing model and change the struct name, fields, interface.
- create a new repository: copy an existing repository and change the struct name, function inputs, queries to implement the model repository.
- create a new controller: copy an existing controller and change the function names, parameter handling, validation, add new messages (variables.go in controller pkg).
- create a new route: copy an existing route and change the route path, controller function calls.
- update the controller test, add repository mock.
- create a new template: copy an existing template and change the template name, fields, form fields, submit button.

## Frontend

Pages:
- Listing
- Form (create, update)
- Detail

UI:
Left menu strip
- Top menu is the open - close button
- The rest of the items are the menu items
Right content
- Page title
- Page action buttons
- Page content
