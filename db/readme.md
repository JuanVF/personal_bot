# DB

## Postgres

Right now this project has a simple docker image for a Postgres database since I don't need something so complex. Make sure to create a file called `docker-compose.yaml` with the contents inside `docker-compose.template.yaml` and change the `password` field.

By setting up the DB the migration files will be executed automatically.
