# gogolf-api

⛳ Welcome to the code repository for gogolf's API! gogolf is a toy project I've created to demonstrate understanding of full-stack software development, DevOps and having a *crazy* high handicap. Having said that, I will hopefully get it up-and-running as a web app at some point.

You're currently looking at the code for the *backend* - that's the API for the gogolf service which ~~contains~~ *will hopefully at some point contain* logic and handlers for:

- creating unique golfers in the gogolf database
- tracking the clubs in a golfer's bag and their yardages for that golfer
- adding courses to the gogolf database, including their course ratings, slope, yardages and total par.
- calculating a golfer's handicap index, and their playing handicap for any course in the database
- any other things I can dream up as I mindlessly top balls down the driving range

## 🏌️ Related Repositories (Coming Soon™)

- gogolf-cli: the code repository for the CLI tool that interacts with the API
- gogolf-web: the code repository for the front-end
- gogolf-qa: code for automated test suites that run against the service
- gogolf-local: docker-compose files and scripts for running the service locally
- gogolf-terraform: the infrastructure-as-code used to deploy the gogolf website

## 🛺 What's in the box? 

gogolf's API is written, shockingly, in Go. The database queries and schemas are generated and handled by the [sqlc](https://github.com/sqlc-dev/sqlc) and the [goose database migration tool](https://github.com/pressly/goose). API documentation is generated from the `gogolf-openapi-spec.yml` file and can be viewed [here](https://samw94.github.io/gogolf-api/) - it will be updated automatically on every push to main. Other relevant documentation can be found in the docs directory of this project.

## 🚩 How do I run it?

### Pre-requisites

- These instructions are for Linux/Unix/WSL. I'm not writing a Windows guide, and you can't make me.
- You must have the [Go toolchain](https://go.dev/doc/install) installed
- You must have [Docker](https://www.docker.com/get-started/) installed
- You must have a [PostgreSQL](https://www.postgresql.org/) database installed, or use the [postgres docker image](https://hub.docker.com/_/postgres)
- You must have the [goose CLI tool](https://github.com/pressly/goose?tab=readme-ov-file#install) installed.
- For the easiest experience, you should have [Docker Compose](https://docs.docker.com/compose/install/) installed.

### Using `docker compose up`

1. Clone this repository.

    `git clone https://github.com/SamW94/gogolf-api.git`

2. Change directory into the root of the project.

    `cd gogolf-api`

3. Create a `.env` from the `sample-env` file.

    `cp sample-env .env`

4. Edit your `.env` file with your favourite text editor to change the values in there by default to whatever you wish. The `DB_URL` and `GOOSE_DBSTRING` variables must both be the same and be in this format:

    `postgres://<user>:<password>@localhost:5432/<postgres-db-name>?sslmode=disable`

    for example:

    `postgres://user:password@localhost:5432/gogolf?sslmode=disable`

5. Your `.env` file should also contain something like this to ensure the API container can communicate with the database (`localhost` will not work as the DB URL if you're using docker compose): 

    `DB_URL_DOCKER="postgres://<user>:<password>@gogolf-postgres:5432/gogolf?sslmode=disable"`

6. Ensure you Docker engine is running, and run `docker-compose up` or start your Postgres container and the API container. 

6. Open another terminal window. From the root of the project, change into the `sql/schema` directory and run the goose migration. 

    ```
    cd sql/schema
    goose up -env ../../.env
    ```

    You should see something like this if you've done everything right.

    ```
    2025/10/17 20:43:29 OK   001_golfers.sql (13.89ms)
    2025/10/17 20:43:29 goose: successfully migrated database to version: 1
    ```

8. The API is serving traffic on `http://localhost:8080` - test it out using Postman, curl or the [gogolf-cli!](https://github.com/SamW94/gogolf-cli)
