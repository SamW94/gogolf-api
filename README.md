# gogolf-api

⛳ Welcome to the code repository for gogolf's API! gogolf is a toy project I've created to demonstrate understanding of full-stack software development, DevOps and having a *crazy* high handicap. Having said that, I will hopefully get it up-and-running as a useable app at some point.

You're currently looking at the code for the *backend* - that's the API for the gogolf service which ~~contains~~ *will hopefully at some point contain* logic and handlers for:

- creating unique golfers in the gogolf database
- tracking the clubs in a golfer's bag and their yardages for that golfer
- adding courses to the gogolf database, including their course ratings, slope, yardages and total par.
- calculating a golfer's handicap index, and their playing handicap for any course in the database
- any other things I can dream up as I mindlessly top balls down the driving range

## 🏌️ Related Repositories (More Coming Soon™)

- [gogolf-cli](https://github.com/SamW94/gogolf-cli): CLI tool that interacts with the API
- [gogolf-local](https://github.com/SamW94/gogolf-local): Tools for running the gogolf API on your local machine with Docker


## 🛺 What's in the box? 

gogolf's API is written, shockingly, in Go. The database queries and schemas are generated and handled by the [sqlc](https://github.com/sqlc-dev/sqlc) and the [goose database migration tool](https://github.com/pressly/goose). API documentation is generated from the `gogolf-openapi-spec.yml` file and can be viewed [here](https://samw94.github.io/gogolf-api/) - it will be updated automatically on every push to main. Other relevant documentation can be found in the docs directory of this project.

## 🚩 How do I run it?

The *easiest way* to try the gogolf API out locally is to follow the instructions in the README [here](https://github.com/SamW94/gogolf-local). 

Alternatively, there are some steps below for running using `go run` if you prefer.

### Pre-requisites

- These instructions are for Linux/Unix/WSL. I'm not writing a Windows guide, and you can't make me.
- You must have the [Go toolchain](https://go.dev/doc/install) installed
- You must have [Docker](https://www.docker.com/get-started/) installed
- You must have a [PostgreSQL](https://www.postgresql.org/) database installed, or use the [postgres docker image](https://hub.docker.com/_/postgres)
- You must have the [goose CLI tool](https://github.com/pressly/goose?tab=readme-ov-file#install) installed.
- For the easiest experience, you should have [Docker Compose](https://docs.docker.com/compose/install/) installed.

### Using `go run`

1. Clone this repository.

    `git clone https://github.com/SamW94/gogolf-api.git`

2. Change directory into the root of the project.

    `cd gogolf-api`

3. Configure and start your Postgres database, if you're running it locally. 

3. Create an `.env` file.

4. Edit your `.env` file with your favourite text editor to change the values in there by default to whatever you wish. The `DB_URL` and `GOOSE_DBSTRING` variables must both be the same and be in this format:

    `postgres://<user>:<password>@localhost:5432/<postgres-db-name>?sslmode=disable`

    for example:

    `postgres://user:password@localhost:5432/gogolf?sslmode=disable`

Your `.env` file **must contain**:

- DB_URL
- GOOSE_DBSTRING
- GOOSE_DRIVER
- INTERNAL_PORT

6. Open another terminal window. From the root of the project, change into the `sql/schema` directory and run the goose migration. 

    ```
    cd sql/schema
    goose up -env ../../.env
    ```

7. From the root of the project, change into the `src` directory. Run the code with `go run .` and you should see a message like below.

    ```
    cd src
    go run .
    ```

    You should see a message like this if you've done everything right.

    ```
    2025/10/17 20:46:24 Serving on internal port: <the internal port from your .env file>
    ```

8. The API is serving traffic on `http://localhost:<your-port>` - test it out using Postman or curl!