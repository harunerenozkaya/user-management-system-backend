
<h1 align="center"> User Management System Backend</h1> <br> 
<p align="center">
<img src = "img/umg_logo.png" height=200> &nbsp;
</p>

## Table of Contents

- [Table of Contents](#table-of-contents)
- [Introduction](#introduction)
- [Navigation](#navigation)
- [Prerequisites](#prerequisites)
- [Getting Started](#getting-started)
- [API Information](#api-information)
- [Technologies](#technologies)

## Introduction

The User Management System is a system that provides a mechanism for storing and managing user information. Managers can create, edit, and delete user profiles. This document describes the backend component of the User Management System project, which offers the necessary APIs and services for user data management.

## Navigation
- Backend : this
- Fronted : https://github.com/harunerenozkaya/user-management-system-frontend

## Prerequisites

Before running this project, make sure you have the following installed:

- Go
- Sqlite

## Getting Started

1. Clone the repository:

    ```bash
    git clone https://github.com/your-username/user-management-system-backend.git
    ```

2. Install the dependencies:

    ```bash
    go mod download
    ```

4. Run the application:

    ```bash
    cd ./cmd/app
    go run main.go
    ```

5. Default port

- As a default , server runs on 8080 port.

## API Information
```
user : {
	"id "       int
	"name"      string
	"surname"   string
	"email"     string
	"created_at" string
	"updated_at" string
}
```

GET /users : Get all users<br>
Response 200 : [users]

POST /users  : Create a new user<br>
Request Body : user<br>
Response 200 : user<br>

GET /users/id  : Get a user with id<br>
Response 200 : user<br>

PUT /users/id  : Update a user with id<br>
Response 200 : user<br> 

DELETE /users/id  : Delete a user with id<br>
Response 200 : id<br> 

## Technologies

- Go
- SQLite
- Gorilla Mux   
