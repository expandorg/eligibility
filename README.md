# Eligibility Service
 
Backend service for worker-job eligibility

## Getting started 

### Prerequisities:

- Install Go (On OS X with Homebrew you can just run `brew install go`.)
- [optional to debug] Postman

### Setup the project

Clone the repository with: 

`go get -u github.com/gemsorg/eligibility`

OR 

create a directory `$GOPATH/src/github/gemsorg` and execute: git clone git@github.com:gemsorg/eligibility.git 

Run the project dependencies (db, etc.) with `make up`

Run the latest migration with `make migrate-latest`

Run the project with `make run`

### Dependencies

We use `dep` to manage our dependencies.

To add a new vendor, use: 

`deps ensure -add DEPENDENCY`

To update vendors for built project, run:

`make update-deps`

## CI / CD
We use Google Cloud for CI/CD:

*note: please don't modify the following files unless you know what you're doing :)*

**cloudbuild.cd.yaml & cloudbuild.cd.staging.yaml:** this effectively our CD, it run tests, builds and pushes the image to the container registry and deploys to production on every Master commit, so master has to be always clean. 

**k8s.yaml & k8s.staging.yaml:** this is the kubernetes setup, including workload and service setup. cloudbuild.cd uses this file to deploy.

#### Production Deployment
This happens automatically when you merge or push to master.

#### Staging Deployment
To push any branch, including master, to staging (dev.expand.org), you need to add a tag by using: `make tag-staging` and push it to github. This will automatically trigger a build and deployment. 

```
# please don't use --tags flag, use the following to push to origin:
git push origin <BRANCH NAME> --follow-tags
```

## Database

### Add a new migration

```make add-migration name="migration_name"```

For migration names be descriptive and start with verbs: `create_`, `drop_`, `add_`, etc.

This will look at the latest migrated version (1, 2, 3) and creates 2 files with new version:

`2_migration_name.up.sql` and `2_migration_name.down.sql`

### Migrate

You can migrate to latest:

```make migrate-latest```

OR 

You can migrate up and migrate down a version:

```make run-migrations action="goto" version="1"```

When you migrate up, you can see in the `schema_migrations` the last migrated version. When you migrate down, it updates the the version column in `schema_migrations`.

## Tests
```make run-tests```

### Unit tests
We keep all unit tests close to the code and withing the same package. For example, if you want to test the service package, then you would add the tests in that folder marked `package service`.

### Functional

We keep all functional tests in `tests/` folder. Create a new test file for every function. 

