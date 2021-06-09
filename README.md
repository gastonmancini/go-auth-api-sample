# go-auth-api-sample

This is just an example on how to implement a simple API in Go, with basic authentication using JWT tokens and user management capabilities.

## 💡 Motivation

The solution is pretty naive and was developed only for trying out [Fiber](https://gofiber.io/) and [GORM](https://gorm.io/).
Also the idea is leave the API ready to be consumed by frontend apps to be developed in newer versions of React and Angular (just to try them out too :)).

## 👀 Overview

In short, the API allows you to register new users, login users, get and edit users, get and edit user roles, get permissions, get and edit user profile, change current user passwords and upload user profile image.

Endpoints:

```

GET /api/users/{id}
GET /api/users?page={pageNumber}
POST /api/users
PUT /api/users/{id}
DELETE /api/users/{id}

GET /api/roles/{id}
GET /api/roles?page={pageNumber}
POST /api/roles
PUT /api/roles/{id}
DELETE /api/roles/{id}

GET /api/permissions?page={pageNumber}

POST /api/register
POST /api/login
POST /api/logout

GET /api/me
PUT /api/me
PUT /api/me/password
POST /api/me/image

```

## 🧬 Development

The application is written purely in golang. MySql is used to persiste the application data.

### Layout

```tree
├── controllers
│   ├── authController.go
│   ├── permissionController.go
│   ├── roleController.go
│   └── userController.go
├── database
│   └── connect.go
├── documentation
│   └── go-auth-api-sample.postman_collection.json
├── middlewares
│   ├── authenticationMiddleware.go
│   └── authorizationMiddleware.go
├── models
│   ├── paginated.go
│   ├── permission.go
│   ├── role.go
│   └── user.go
├── routes
│   ├── routes.go
│   ├── authRoutes.go
│   ├── permissionRoutes.go
│   ├── roleRoutes.go
│   └── uerRoutes.go
├── uploads
├── third_party
│   └── README.md
└── util
│   ├── cookie.go
│   └── jwt.go
├── .air.toml
├── .gitignore
├── LICENSE
├── README.md
├── go.mod
├── go.sum
└── main.go
```

A brief description of the layout:

* `controllers` contains the application controllers
* `database` contains the database migration and connection
* `documentation` the documentation and other useful assets
* `middlewares` contains the authentication and authorization middlewares
* `models` the domain models
* `routes` define the api routing
* `uploads` folder to serve static files
* `util` utilities

## 📖 Database

* Uses [GORM](https://gorm.io/index.html) as ORM and MySql.
* GORM Auto Migration is enabled. The database schema is created by the app.
* Refer to this [link](https://github.com/go-sql-driver/mysql#dsn-data-source-name) for details on how to set the data source name
* Example: dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
* Tables created in the db: users, roles, permissions, role_permissions

## ⚙️ Run

If Air is installed just run the command `air`. If not you can run it with `go run main.go`.

[Air](https://github.com/cosmtrek/air) is setup to be used for live reload.

## ☕ To do

- [] Refactor to apply Uncle Bob - Clean Architecture
- [] Add unit tests

## ⚠️ Warning

The app was developed for educational purposes only. Do not use it in prod :)
