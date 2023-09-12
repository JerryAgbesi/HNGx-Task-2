# HNGx Backend Task 2

A REST API capable of performing CRUD operations on a "Person" resource.

## Getting Started
This project is built with Gin Gonic, a web framework for Go. It uses a SQLite database and GORM as the ORM.

BaseURL: `https://hngxx-task-2.onrender.com/`

## UML Diagram
## Usage
- Clone the repository

    ```bash
    git clone git@github.com:JerryAgbesi/HNGx-Task-2.git
    ```
- Run the Application locally
    
    ```bash
    go run main.go
    ```

## Endpoints
| Name | Request Method | Endpoint | Parameters | Response code |
| --- | --- | --- | --- | --- |
|Create a person | POST | /api | None | 201 CREATED|

- Request Body

```json
{
    "name": "Jerry Agbesi",
}
```
- Sample Response
```json
{
    "id": 1,
    "name": "Jerry Agbesi",
}
```

| Name | Request Method | Endpoint | Parameters | Response code |
| --- | --- | --- | --- | --- |
|Retrieve a person| GET | /api/{id} | id | 200 OK|

- Request Body : None

- Sample Response
```json
{
    "id": 1,
    "name": "Jerry Agbesi",
}
```

| Name | Request Method | Endpoint | Parameters | Response code |
| --- | --- | --- | --- | --- |
|Update details of a person | PUT | /api/{id} | None | 200 OK|

- Request Body

```json
{
    "name": "Betram Gilfoyle",
}
```
- Sample Response
```json
{
    "id": 1,
    "name": "BeRtram Gilfoyle",
}
```

| Name | Request Method | Endpoint | Parameters | Response code |
| --- | --- | --- | --- | --- |
|Delete a person| GET | /api/{id} | id | 204 NO CONTENT|

- Request Body : None

- Sample Response
```json
{
    "id": 1,
    "name": "Jerry Agbesi",
}
```
## Testing

