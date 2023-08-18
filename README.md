# job-board-backend-project

## Description
A Backend project that simulates a job board. Applicants can view jobs and apply to them, also includes Administrator roles that can edit postings and users.
This project is currently in progress and should be completed soon. Come back soon to see updates and more documentation.

## Diagram for Applicant APIs
![Applicant Endpoints](https://github.com/davisshriver/job-board-backend-project/assets/18060803/5e0ceb55-237a-449b-8803-92302b0f55b7)


## Diagram for Admin APIs
![Admin Endpoints](https://github.com/davisshriver/job-board-backend-project/assets/18060803/036f8268-cb11-4202-a772-1993791964d6)

# User Endpoints
## Get User: `/users/:id`

Get user details by ID.

### Access Level
`USER`

### Method

`GET`

### URL Parameters

| Parameter | Type   | Description            |
|-----------|--------|------------------------|
| `id`      | int    | User ID                |

**Response Body Example:**

```json
{
    "user_id": 2,
    "first_name": "John",
    "last_name": "Doe",
    "email": "john.doe@gmail.com",
    "phone": "9403334444",
    "user_type": "USER",
    "created_at": "2023-08-08T12:44:01.597775-05:00",
    "updated_at": "2023-08-08T12:44:01.597775-05:00"
}
```

## Get Users: `/users`

Get all users.

### Access Level
`ADMIN`
### Method

`GET`

**Response Body Example:**

```json
[
    {
        "user_id": 1,
        "first_name": "Alice",
        "last_name": "Johnson",
        "email": "alice@example.com",
        "phone": "555-1234",
        "user_type": "ADMIN",
        "created_at": "2023-08-10T09:30:15.12345Z",
        "updated_at": "2023-08-10T09:30:15.12345Z"
    },
    {
        "user_id": 2,
        "first_name": "Bob",
        "last_name": "Smith",
        "email": "bob@example.com",
        "phone": "555-5678",
        "user_type": "USER",
        "created_at": "2023-08-11T15:20:45.67890Z",
        "updated_at": "2023-08-11T15:20:45.67890Z"
    }
]
```

## Sign Up: `/users/signup`

Sign up for the job site.

### Access Level
`USER`
### Method

`POST`

### Request Parameters

| Parameter   | Type   | Description             |
|-------------|--------|-------------------------|
| `first_name`| string | User's first name       |
| `last_name` | string | User's last name        |
| `password`  | string | User's password         |
| `email`     | string | User's email address    |
| `phone`     | string | User's phone number     |
| `user_type` | string | User's role (ADMIN/USER)|

**Response Body Example:**

```json
{
    "user_id": 3,
    "first_name": "Emma",
    "last_name": "Williams",
    "email": "emma@example.com",
    "phone": "555-9876",
    "user_type": "USER"
}
```

## Login: `/users/login`

Log into the job site and receive a JWT.

### Access Level
`USER`
### Method

`POST`

### Request Parameters

| Parameter   | Type   | Description             |
|-------------|--------|-------------------------|
| `email`     | string | User's email address    |
| `password`  | string | User's password         |

**Response Body Example:**

```json
{
    "user_id": 1,
    "token": "eyJhbGciOiJIUzR4NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6IjE5ZHNocml2ZXJAZ21haWwuY29tIiwiRmlyc3ROYW1lIjoiRGF2aXMiLCJMYXN0TmFtZSI6IlNocml2ZXIiLCJVaWQiT6IiLCJVc2VyVHlwZSI6IkFETUlOIiwiZXhwIjoxNjkyNDY3Mjg0fQ.Qn3Xg7343gBWM0mT7xgRCE5sOA7NgS4Dk7HaAB4orLM"
}
```
