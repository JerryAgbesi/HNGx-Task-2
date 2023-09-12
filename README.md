# HNGx Backend Task 2

A REST API capable of performing CRUD operations on a "Person" resource.

## Getting Started
This project is built with Gin Gonic, a web framework for Go. It uses a SQLite database and GORM as the ORM.

**BaseURL**: `https://hngxx-task-2.onrender.com/`

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
| Name | Request Method | Endpoint | Route Parameters | Response code |
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

| Name | Request Method | Endpoint | Route Parameters | Response code |
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

| Name | Request Method | Endpoint | Route Parameters | Response code |
| --- | --- | --- | --- | --- |
|Update details of a person | PUT | /api/{id} | id | 200 OK|

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
    "name": "Bertram Gilfoyle",
}
```

| Name | Request Method | Endpoint | Route Parameters | Response code |
| --- | --- | --- | --- | --- |
|Delete a person| DELETE | /api/{id} | id | 204 NO CONTENT|

- Request Body : None

- Sample Response
```json
{
    "id": 1,
    "name": "Jerry Agbesi",
}
```
## Testing
Test the live API with a collection in [Postman](https://www.postman.com/)

[![Run in Postman](https://run.pstmn.io/button.svg)](https://elements.getpostman.com/redirect?entityId=21533336-c9cc0288-e477-405d-9166-f4c183e0e9f5&entityType=collection)
