# TODO Web application
This web application is a microsoft todo clone.
Feature Description:
- Registration / authorization
- Ability to create lists
- Possibility to create tasks in lists

## Stack: 
Golang, PostgreSQL

## Endpoints:

Registration / authorization:
```
POST /auth/sign-up
POST /auth/sgn-in
```

Actions with lists:
```
GET  /lists
GET  /lists/{id}
POST /lists
PUT  /lists/{id}
DELETE /lists/{id}
GET  /lists/{id}/tasks
POST /lists/{id}/tasks
```
Actions with tasks:
```
PUT /tasks/{id}
GET /tasks/{id}
DELETE /tasks/{id}
```

DB:
```
docker run --name=todo-db -e POSTGRES_PASSWORD='qwerty' -p 5432:5432 -d --rm postgres
```

Install migrate:
```
scoop install migrate 
```

Creating an init for the DB:
```
migrate create -ext sql -dir ./schema -seq init
```

Migrations to DB:
```
migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5432/postgres?sslmode=disable' up

migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5432/postgres?sslmode=disable' down
```

To view the created tables:
```
docker exec -it {id_conainer} /bin/bash
psql -U postgres
\d          //To view the created tables

\q          //To exit
exit
```
