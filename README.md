# job-board-backend-project

## Description
A Backend project that simulates a job board. Applicants can view jobs and apply to them, also includes Administrator roles that can edit postings and users.
This project is currently in progress and should be completed soon. Come back soon to see updates and more documentation.

## Table of Contents

- [User Endpoints](#user-endpoints)
- [Job Post Endpoints](#job-post-endpoints)
- [Application Endpoints](#application-endpoints)

## Diagram for Applicant APIs
![Applicant Endpoints](https://github.com/davisshriver/job-board-backend-project/assets/18060803/5e0ceb55-237a-449b-8803-92302b0f55b7)


## Diagram for Admin APIs
![Admin Endpoints](https://github.com/davisshriver/job-board-backend-project/assets/18060803/036f8268-cb11-4202-a772-1993791964d6)

# User Endpoints
<a name="user-endpoints"></a>
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

## Get User: `/users/:user_id`

Get user details by ID.

### Access Level
`USER`

### Method

`GET`

### URL Parameters

| Parameter | Type   | Description            |
|-----------|--------|------------------------|
| `user_id`      | string | User ID           |

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

## Edit User: `/users/:user_id`

Edit a user's information. A user can only edit their own information, but administrators can edit any user's information.

### Access Level
`USER`
### Method

`PATCH`

### URL Parameters

| Parameter | Type   | Description            |
|-----------|--------|------------------------|
| `user_id`      | string | User ID           |

### Request Parameters

| Parameter   | Type      | Description          |
|-------------|-----------|----------------------|
| `email`     | string    | User's email address |
| `password`  | string    | User's password      |
| `first_name`| string    | User's first name    |
| `last_name` | string    | User's last name     |
| `phone`     | string    | User's phone number  |
| `user_type` | string    | User's role (ADMIN/USER) |
| `updated_at`| time.Time | Updated timestamp    |


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

## Delete User: `/users/:user_id`

Deletes a user from the database.

### Access Level
`ADMIN`

### Method

`DELETE`

### URL Parameters

| Parameter | Type   | Description            |
|-----------|--------|------------------------|
| `user_id`      | string | User ID           |

**Response Body Example:**

```json
{
    "success": "User deleted from the database"
}
```

# Job Post Endpoints
<a name="job-post-endpoints"></a>

## Get Posts: `/posts`

Get all job posts.

### Access Level
`USER`

### Method

`GET`

**Response Body Example:**

```json
[
    {
        "PostID": 1,
        "role": "Job1",
        "description": "Cool Job",
        "requirements": "1 Year Experience",
        "created_by": "Admin",
        "location": "Fort Worth, TX",
        "wage": 12,
        "created_at": "2023-08-18T10:00:00Z",
        "expires_at": "2023-08-31T23:59:59Z"
    },
    {
        "PostID": 2,
        "role": "Job2",
        "description": "Exciting Position",
        "requirements": "Bachelor's Degree",
        "created_by": "HR Manager",
        "location": "New York, NY",
        "wage": 20,
        "created_at": "2023-08-19T09:30:00Z",
        "expires_at": "2023-09-15T23:59:59Z"
    }
]

```

## Get Post: `/posts/:post_id`

Get post by post id.

### Access Level
`USER`

### Method

`GET`

### URL Parameters

| Parameter | Type   | Description            |
|-----------|--------|------------------------|
| `post_id`      | string    | Post ID        |

**Response Body Example:**

```json
 {
        "PostID": 2,
        "role": "Job2",
        "description": "Exciting Position",
        "requirements": "Bachelor's Degree",
        "created_by": "HR Manager",
        "location": "New York, NY",
        "wage": 20,
        "created_at": "2023-08-19T09:30:00Z",
        "expires_at": "2023-09-15T23:59:59Z"
    }
```

# Application Endpoints
<a name="application endpoints"></a>
