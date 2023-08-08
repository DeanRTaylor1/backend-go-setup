# Go Test Api

This README provides a guide to set up your development environment and run your application. Our project is using Go, PostgreSQL and Docker.

## Prerequisites

Before you can run this project, you will need the following installed:

- Docker: Please follow the official Docker [guide](https://docs.docker.com/get-docker/) to install it.
- Go: Visit the official Go [documentation](https://golang.org/doc/install) for installation instructions.
- Docker Compose: Follow the official Docker Compose [installation guide](https://docs.docker.com/compose/install/).

## Setting Up Your Development Environment

1. Clone the repository and navigate to the root directory.

2. Now, you can start the development server:

```bash
make dev
```

Your API should now be accessible at `http://localhost:8080`.

## Running Tests

You can run the test suite by executing:

```bash
make test-suite
```

## Clean Up

To stop all services and clean up the environment, run:

```bash
make cleanup
```

## Migrating Up and Down

Use the following command to migrate your database schema up:

```bash
make migrateup
```

To migrate down, use:

```bash
make migratedown
```

## Generate SQL Models

To generate models based on your SQL schema:

```bash
make sqlc
```

**Note**: Please make sure that the environment variable `DB_URL` is set to `postgresql://root:secret@localhost:5432/test_db?sslmode=disable` before running migration or SQL generation commands.
