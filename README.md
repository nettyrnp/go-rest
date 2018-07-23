# Simple RESTful CRUD application with role-based access control

Features:

* RESTful endpoints
* Standard CRUD operations (all transactional)
* JWT-based authentication
* Paging


## How to launch the application

Start PostgreSQL database server. In the PostgreSQL database named `postgres` (it is a default one) execute the SQL statements given in the file `migrate/migration.sql`.
Database connection information (see file `config/app.yaml`):
* server address: `127.0.0.1` (localhost)
* server port: `5432`
* database name: `postgres`
* username: `postgres`
* password: `postgres`

Install the application from the Terminal:
```shell
go get github.com/nettyrnp/go-rest
```

Run the application from within Intellij IDEA (`Run` button with server.go file open) or from the Terminal:
```shell
make run
```

The application runs as an HTTP server at port 8080. It provides the following RESTful endpoints:

* `POST /u2/auth`: authenticate a user
* `GET /u2/users`: returns a paginated list of the users
* `GET /u2/users/<id>`: returns the detailed information of a user
* `POST /u2/users`: creates a new user
* `DELETE /u2/users/<id>`: deletes a user [NB: available only for users with role `admin`]

Now in a separate Terminal run the following commands (or use Postman etc):

```shell
# authenticate as user via: POST /u2/auth
curl -X POST -H "Content-Type: application/json" -d '{"username": "regular_user", "password": "pass"}' http://localhost:8080/u2/auth
# should return a JWT token like: {"token":"...eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."}

# authenticate as admin via: POST /u2/auth
curl -X POST -H "Content-Type: application/json" -d '{"username": "admin", "password": "pass2"}' http://localhost:8080/u2/auth
# should return a JWT token like: {"token":"...eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."}

# with the received JWT token, you can access the user resources, e.g.:
curl -X GET -H "Authorization: Bearer <your_JWT_token>" http://localhost:8080/u2/users
# should return a list of user records in the JSON format
```
