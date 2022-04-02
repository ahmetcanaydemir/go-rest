# Go Rest

You can use this [heroku URL](https://go-restful-app.herokuapp.com/) for testing the API server in this repository.

[![CI](https://github.com/ahmetcanaydemir/go-rest/actions/workflows/go.yml/badge.svg)](https://github.com/ahmetcanaydemir/go-rest/actions/workflows/go.yml)

## Libraries

- [stretchr/testify](https://github.com/stretchr/testify): Mocking
- [mongodb/mongo-go-driver](https://github.com/mongodb/mongo-go-driver): Database driver

## Run

1. Add `MONGO_URI` as environment variable. You can also add `PORT` but it is not required (default 8080).

    `export MONGO_URI="mongodb+srv://mongouri..."`

2. You can run locally or with docker. Choose one of the following commands.

    - **Docker:** `docker compose up`
    - **Local:** `make run`

## Commands

    make Command

| Command  | Description                                            |
|----------|--------------------------------------------------------|
| run      | Runs API server                                        |
| all      | Runs fmt, lint, test  and build with order             |
| build    | Builds API Server to build folder                      |
| clean    | Removes ./bin folder                                   |
| test     | Runs tests without coverage                            |
| coverage | Runs all tests and shows total coverage                |
| fmt      | Checks code format                                     |
| lint     | Runs lint tools                                        |
| help     | Shows this message                                     |
| install  | Downloads dependencies                                 |

## Project Structure 

```
├── pkg
│   ├── api
│   │   ├── controller
│   │   ├── repository
│   │   └── service
│   ├── configs
│   ├── db 
│   └── model
└── test
    ├── integration
    │   └── repository
    └── unit
        ├── controller
        ├── repository
        └── service
```

## Endpoints

### in-memory

> POST /in-memory

This endpoint saves the key and value to in memory database and echoes the request body.

#### Request Body

| Field               | Type   | Description                   |
|---------------------|--------|-------------------------------|
| `key`    (required) | string | Unique key for save in DB.    |
| `value`(required)   | string | String value.                 |

#### Example Successful Request

```json
POST /in-memory
{
    "key":"test-key",
    "value":"test-value"
}
```

#### Example Successful Response

```json
{
    "key": "test-key",
    "value": "test-value"
}
``` 

> GET /in-memory

This endpoint retrives the value from in memory db with using given key.

#### Query Params

| Field               | Type   | Description                               |
|---------------------|--------|-------------------------------------------|
| `key`    (required) | string | Unique key for retrive the value from DB. |

#### Example Successful Request

```json
GET /in-memory?key=test-key
```

#### Example Successful Response

```json
{
    "key": "test-key",
    "value": "test-value"
}
``` 

### mongo

> POST /mongo

This endpoint fetch data and filter with given values.

#### Request Body

| Field                   | Type                   | Description                                   |
|-------------------------|------------------------|-----------------------------------------------|
| `startDate` (required)  | string (YYY-MM-DD)     | The record is created after this date.        |
| `endDate` (required)    | string (YYY-MM-DD)     | The record is created before this date.       |
| `minCount` (required)   | int                    | TotalCount value is bigger than this value.   |
| `maxCount` (required)   | int                    | TotalCount value is less than this value.     |

#### Example Successful Request

```json
POST /mongo
{
    "startDate":"2016-01-26",
    "endDate":"2018-02-02",
    "minCount":10,
    "maxCount":100
}
```

#### Example Successful Response

```json
{
    "code": 0,
    "msg": "Success",
    "records": [
        {
            "key": "gWUDiUcV",
            "createdAt": "2016-12-29T22:37:44.688Z",
            "totalCount": 81
        }
    ]
}
``` 
