# mallus_api

># Api responsible for mantaining Mallus service.
## Techs
- Postgresql
- Go Lang

## Installation and running locally
This API requires `Docker` to run.

### First step before installing Docker.
1. docker-compose up -d -build;

2. Now you can make api requests on localhost:8080.


## Endpoints
### Product
Create product:
```bash
curl -X POST http://localhost:8080/product -d '{"name": "gin tonic", "description": "cold drink"}' -H 'Content-Type: application/json'
```

List products:
```bash
curl -X GET http://localhost:8080/product -d -H 'Content-Type: application/json'
```

Remove product:
```bash
curl -X DELETE http://localhost:8080/product/:id -d -H 'Content-Type: application/json'
```

Update product:
```bash
curl -X PUT http://localhost:8080/product/:id -d '{"name": "gin gonic", "description": "code drink"}' -H 'Content-Type: application/json' -H 'Content-Type: application/json'
```

### Auth
Create account:
```bash
curl -X POST http://localhost:8080/signup -d '{"name": "Mateus", "email": "mateus@email.com", "password": "mysimplepassword"}' -H 'Content-Type: application/json' -H 'Content-Type: application/json'
```
