# TODO Web application
This web application is a microsoft todo clone.
Feature Description:
- Registration / authorisation
- Ability to create lists
- Possibility to create tasks in lists

## Stack: 
Golang, PostgreSQL

## Endpoints:

Registration / authorisation:
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
