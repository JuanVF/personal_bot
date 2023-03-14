# DB

## Postgres

Right now this project has a simple docker image for a Postgres database since I don't need something so complex. Make sure to create a file called `Dockerfile` with the contents inside `Dockerfile.template` and change the `POSTGRES_PASSWORD` env variable.

By setting up the DB the migration files will be executed automatically.
