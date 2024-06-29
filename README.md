# SPYCAT

[![Go Report Card](https://goreportcard.com/badge/github.com/markraiter/spycat)](https://goreportcard.com/report/github.com/markraiter/spycat)

## Description

This is a simple yet comprehensive CRUD application for creating Spy Cats, Missions and Targets for them.
There is also authentication functionality, so only registered users can make operations with Cats/Missions/Targets/

## Installation

To install and run this project, follow these steps:

1. Clone the repository: `git clone https://github.com/markraiter/spycat.git`
2. Install the dependencies with `go mod download`
3. Create `.env` file and copy values from `.env_example`
4. Follow the instructions to install [Taskfile](https://taskfile.dev/ru-ru/installation/) utility
5. Follow the instructions to install [Golang Migrate](https://github.com/golang-migrate/migrate) utility
6. Run migrations with `task migrateup`
7. Run the app with `task run`
8. You can check Swagger docs after run on `localhost:8000/swagger`

**ATTENTION!!!** By default the app will run on port `localhost:8000`, or in any other you provide in your `.env` file.

### Built With

- [Go](https://golang.org/) - The programming language used.
- [Fiber](https://gofiber.io/) - Framework used for transport layer implementations.
- [REST](https://en.wikipedia.org/wiki/Representational_state_transfer) - Architectural style for the API.
- [Clean Architecture](https://8thlight.com/blog/uncle-bob/2012/08/13/the-clean-architecture.html) - Architectural pattern used.
- [Postgres](https://www.postgresql.org/) - Database used.
- [Golang-Migrate](https://github.com/golang-migrate/migrate) - Database migrations tool.
- [JWT](https://jwt.io/) - Used for authentication.
