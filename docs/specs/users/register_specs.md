### Register User API

#### Endpoint

`POST` /register

#### Description

API to register a new user with name, username, and password information.

`Request` Body

```json
{
  "name": "string",
  "username": "string",
  "password": "string"
}
```

- `name` (string, required): Full name of the user.
- `username` (string, required): Unique username.
- `password` (string, required): Password for the user account.
  
`Response` Status Code :

- HTTP Status Code: `201 Created` if registration is successful.
- HTTP Status Code: `400 Bad Request` if the request body is invalid.
- HTTP Status Code: `500 Internal Server Error` if there is a server error.

`Response` Body Success :

```json
{
  "message": "register successfully",
  "data": {
    "Name": "string",
    "Username": "string"
  }
}
```

- `message` (string): Success message from the registration operation.
- `data.Name` (string): Full name of the newly registered user.
- `data.Username` (string): Username of the newly registered user.

`Response` Body (Error) :

```json
{
  "error": "error message"
}
```
