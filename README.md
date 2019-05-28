# Eligibility Service
 
Backend service for worker-job eligibility

## Getting started 

### Prerequisities:

- Install Go (On OS X with Homebrew you can just run `brew install go`.)
- [optional to debug] Postman

### Setup the project

Clone the repository with: 

`go get -u github.com/gemsorg/eligibility`

OR create a directory `$GOPATH/src/github/gemsorg` and execute: git clone git@github.com:gemsorg/eligibility.git 

Build the project with `make build`

Run the project with `make up`

To see logs running on the server run 

`docker-compose logs -f`

To add a new vendor, use: 

`go get ABC`

To update vendors for built project, run:

`make update-deps`

## Deploying

## Database

### Add a new migration

```make add-migration name="migration_name"```

For migration names be descriptive and start with verbs: `create_`, `drop_`, `add_`, etc.

This will look at the latest migrated version (1, 2, 3) and creates 2 files with new version:

`2_migration_name.up.sql` and `2_migration_name.down.sql`

### Migrate

You can migrate up and migrate down a version:

`make run-migrations action="up" version="1"` and `make run-migrations action="down" version="1"`

When you migrate up, you can see in the `schema_migrations` the last migrated version. When you migrate down, it deletes the the row in `schema_migrations`.
