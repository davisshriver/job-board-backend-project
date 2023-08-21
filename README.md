# job-board-backend-project

## Description
A Backend project that simulates a job board. Applicants can view jobs and apply to them, also includes Administrator roles that can edit postings and users.
This is designed for a single organization's hiring needs. I plan to scale this project to include the capabilities to add organizations and work as a 1-click apply job board.

## Table of Contents
### [User Endpoints](#user-endpoints)
- [Sign up](#create-user)
- [Login](#login)
- [Get Users](#get-users)
- [Get User](#get-user)
- [Edit User](#edit-user)
- [Delete User](#delete-user)
### [Job Post Endpoints](#job-post-endpoints)
- [Get Posts](#get-posts)
- [Get Post](#get-post)
- [Create Post](#create-post)
- [Edit Post](#edit-post)
- [Delete Post](#delete-post)
### [Application Endpoints](#application-endpoints)
- [Get Applications](#get-applications)
- [Get Application](#get-application)
- [Create Application](#create-application)
- [Edit Application](#edit-application)
- [Delete Application](#delete-application)
  
## Diagram for Applicant APIs
![Applicant Endpoints](https://github.com/davisshriver/job-board-backend-project/assets/18060803/5e0ceb55-237a-449b-8803-92302b0f55b7)


## Diagram for Admin APIs
![Admin Endpoints](https://github.com/davisshriver/job-board-backend-project/assets/18060803/036f8268-cb11-4202-a772-1993791964d6)

# User Endpoints
<a name="user-endpoints"></a>
<a name="create-user"></a>
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
<a name="login"></a>
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

<a name="get-users"></a>
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

<a name="get-user"></a>
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

<a name="edit-user"></a>
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

<a name="delete-user"></a>
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
<a name="get-posts"></a>
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

<a name="get-post"></a>
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

<a name="create-post"></a>
## Create Post: `/posts`

Create a new job listing that can be applied to by users.

### Access Level
`ADMIN`
### Method

`POST`

### Request Parameters

| Parameter     | Type      | Description             |
|---------------|-----------|-------------------------|
| `role`        | string    | Role of the job         |
| `description` | string    | Job description         |
| `requirements`| string    | Job requirements        |
| `created_by`  | string    | Created by              |
| `location`    | string    | Job location            |
| `wage`        | int       | Wage                    |
| `expires_at`  | time.Time | Expiration timestamp    |


**Response Body Example:**

```json
{
    "PostID": 2,
    "role": "Marketing Specialist",
    "description": "Exciting opportunity in marketing field.",
    "requirements": "Bachelor's degree in Marketing, strong communication skills.",
    "created_by": "Marketing Director",
    "location": "New York, NY",
    "wage": 60000,
    "created_at": "2023-08-19T13:39:47.5982026-05:00",
    "expires_at": "2023-09-30T23:59:59Z"
}
```

<a name="edit-post"></a>
## Edit Post: `/posts/:post_id`

Edit a job post's details.

### Access Level
`ADMIN`
### Method

`PATCH`

### URL Parameters

| Parameter | Type   | Description            |
|-----------|--------|------------------------|
| `post_id`      | int | post ID              |

### Request Parameters

| Parameter     | Type       | Description                    |
|---------------|------------|--------------------------------|
| `role`        | string     | Updated role of the job        |
| `description` | string     | Updated job description        |
| `requirements`| string     | Updated job requirements       |
| `wage`        | int        | Updated wage                   |
| `expires_at`  | time.Time  | Updated expiration timestamp   |

**Response Body Example:**

```json
{
    "PostID": 2,
    "role": "Marketing Intern",
    "description": "Exciting entry level opportunity in marketing field.",
    "requirements": "Currently pursuing Bachelor's degree in Marketing, strong communication skills.",
    "created_by": "Marketing Director",
    "location": "New York, NY",
    "wage": 50000,
    "created_at": "2023-08-19T13:39:47.598202-05:00",
    "expires_at": "2023-09-30T23:59:59Z"
}
```

<a name="delete-post"></a>
## Delete Post: `/posts/:post_id`

Deletes a post from the database.

### Access Level
`ADMIN`

### Method

`DELETE`

### URL Parameters

| Parameter | Type   | Description            |
|-----------|--------|------------------------|
| `post_id`      | int | Post ID              |

**Response Body Example:**

```json
{
    "success": "Post deleted from the database"
}
```


# Application Endpoints
<a name="application endpoints"></a>
<a name="get-applications"></a>
## Get Applications: `/users/:user_id/applications`

Get all of a specific user's applications.

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
[
    {
        "application_id": 3,
        "user_id": 1,
        "post_id": 3,
        "first_name": "John",
        "last_name": "Doe",
        "email": "john.doe@example.com",
        "phone": "1234567890",
        "address": "123 Main St",
        "city": "New York",
        "state": "NY",
        "postal_code": "10001",
        "cover_letter": "This is my cover letter.",
        "resume_url": "https://example.com/resume.pdf",
        "linkedin_url": "https://linkedin.com/in/johndoe",
        "portfolio_url": "https://johndoe.portfolio.com",
        "referrals": [
            {
                "name": "Jane Smith",
                "phone": "9876543210",
                "relation": "Colleague",
                "title": "Project Manager"
            },
            {
                "name": "Michael Johnson",
                "phone": "5555555555",
                "relation": "Supervisor",
                "title": "Lead Developer"
            }
        ],
        "desired_salary": 60000,
        "availability": "Full-time",
        "education": [
            {
                "degree": "Bachelor's Degree",
                "school": "University of Example",
                "location": "New York, NY",
                "grad_year": 2020
            }
        ],
        "work_history": [
            {
                "position": "Software Engineer",
                "company": "Tech Innovators",
                "location": "San Francisco, CA",
                "start_year": 2020,
                "end_year": 2022,
                "responsibilities": [
                    "Developed web applications",
                    "Collaborated with team"
                ]
            }
        ],
        "created_at": "2023-08-20T18:47:46.774061-05:00",
        "expires_at": "2023-08-20T18:47:46.774061-05:00"
    },
    {
        "application_id": 4,
        "user_id": 1,
        "post_id": 3,
        "first_name": "John",
        "last_name": "Doe",
        "email": "john.doe@example.com",
        "phone": "1234567890",
        "address": "123 Main St",
        "city": "New York",
        "state": "NY",
        "postal_code": "10001",
        "cover_letter": "This is my cover letter.",
        "resume_url": "https://example.com/resume.pdf",
        "linkedin_url": "https://linkedin.com/in/johndoe",
        "portfolio_url": "https://johndoe.portfolio.com",
        "referrals": [
            {
                "name": "Jane Smith",
                "phone": "9876543210",
                "relation": "Colleague",
                "title": "Project Manager"
            },
            {
                "name": "Michael Johnson",
                "phone": "5555555555",
                "relation": "Supervisor",
                "title": "Lead Developer"
            }
        ],
        "desired_salary": 60000,
        "availability": "Full-time",
        "education": [
            {
                "degree": "Bachelor's Degree",
                "school": "University of Example",
                "location": "New York, NY",
                "grad_year": 2020
            }
        ],
        "work_history": [
            {
                "position": "Data Scientist",
                "company": "AI Innovations",
                "location": "San Francisco, CA",
                "start_year": 2020,
                "end_year": 2022,
                "responsibilities": [
                    "Analyzed data patterns",
                    "Developed predictive models"
                ]
            }
        ],
        "created_at": "2023-08-20T18:47:47.975755-05:00",
        "expires_at": "2023-08-20T18:47:47.975755-05:00"
    }
]
```

<a name="get-application"></a>
## Get Application: `/users/:user_id/applications/:application_id`

Get a specific application by application id. User's can only get their own applications.

### Access Level
`USER`

### Method

`GET`

### URL Parameters

| Parameter | Type   | Description            |
|-----------|--------|------------------------|
| `user_id`      | string | User ID           |
| `application_id` | string | Application ID  |


**Response Body Example:**

```json
{
        "application_id": 3,
        "user_id": 1,
        "post_id": 3,
        "first_name": "John",
        "last_name": "Doe",
        "email": "john.doe@example.com",
        "phone": "1234567890",
        "address": "123 Main St",
        "city": "New York",
        "state": "NY",
        "postal_code": "10001",
        "cover_letter": "This is my cover letter.",
        "resume_url": "https://example.com/resume.pdf",
        "linkedin_url": "https://linkedin.com/in/johndoe",
        "portfolio_url": "https://johndoe.portfolio.com",
        "referrals": [
            {
                "name": "Jane Smith",
                "phone": "9876543210",
                "relation": "Colleague",
                "title": "Project Manager"
            },
            {
                "name": "Michael Johnson",
                "phone": "5555555555",
                "relation": "Supervisor",
                "title": "Lead Developer"
            }
        ],
        "desired_salary": 60000,
        "availability": "Full-time",
        "education": [
            {
                "degree": "Bachelor's Degree",
                "school": "University of Example",
                "location": "New York, NY",
                "grad_year": 2020
            }
        ],
        "work_history": [
            {
                "position": "Software Engineer",
                "company": "Tech Innovators",
                "location": "San Francisco, CA",
                "start_year": 2020,
                "end_year": 2022,
                "responsibilities": [
                    "Developed web applications",
                    "Collaborated with team"
                ]
            }
        ],
        "created_at": "2023-08-20T18:47:46.774061-05:00",
        "expires_at": "2023-08-20T18:47:46.774061-05:00"
    }
```

<a name="create-application"></a>
## Create Application: `/users/:user_id/posts/:post_id/applications`

Create a application for a job post. User's can only post applications for themselves.

### Access Level
`USER`
### Method

`POST`

### Request Parameters

#### ApplicationInput

| Field          | Type      | Description                 |
|----------------|-----------|-----------------------------|
| `first_name`   | string    | First name of applicant     |
| `last_name`    | string    | Last name of applicant      |
| `email`        | string    | Email address of applicant  |
| `phone`        | string    | Phone number of applicant   |
| `address`      | string    | Address of applicant        |
| `city`         | string    | City of applicant           |
| `state`        | string    | State of applicant          |
| `postal_code`  | string    | Postal code of applicant    |
| `cover_letter` | string    | Cover letter of applicant   |
| `resume_url`   | string    | URL to applicant's resume   |
| `linkedin_url` | string    | LinkedIn profile URL        |
| `portfolio_url`| string    | Portfolio URL of applicant  |
| `desired_salary` | float64 | Desired salary of applicant |
| `availability` | string    | Availability status         |
| `education`    | array     | Educational background      |
| `referrals`    | array     | Referral information        |
| `work_history` | array     | Work experience details     |

#### Referral

| Field     | Type   | Description          |
|-----------|--------|----------------------|
| `name`    | string | Name of the referral |
| `phone`   | string | Phone of the referral|
| `relation`| string | Relation to applicant|
| `title`   | string | Title of the referral |

#### EducationInfo

| Field     | Type   | Description        |
|-----------|--------|--------------------|
| `degree`  | string | Degree             |
| `school`  | string | School attended    |
| `location`| string | Location of school |
| `grad_year`| int   | Graduation year    |

#### WorkExperience

| Field         | Type   | Description            |
|---------------|--------|------------------------|
| `position`   | string | Position               |
| `company`    | string | Company                |
| `location`   | string | Location               |
| `start_year` | int    | Starting year          |
| `end_year`   | int    | Ending year            |
| `responsibilities`| array | Responsibilities    |

**Response Body Example:**

```json
{
    "application_id": 8,
    "user_id": 1,
    "post_id": 2,
    "first_name": "John",
    "last_name": "Doe",
    "email": "john.doe@example.com",
    "phone": "1234567890",
    "address": "123 Main St",
    "city": "Example City",
    "state": "Example State",
    "postal_code": "12345",
    "cover_letter": "This is my cover letter.",
    "resume_url": "https://example.com/resume.pdf",
    "linkedin_url": "https://linkedin.com/in/johndoe",
    "portfolio_url": "https://johndoe.portfolio.com",
    "referrals": [
        {
            "name": "Jane Smith",
            "phone": "9876543210",
            "relation": "Colleague",
            "title": "Project Manager"
        },
        {
            "name": "Michael Johnson",
            "phone": "5555555555",
            "relation": "Supervisor",
            "title": "Lead Developer"
        }
    ],
    "desired_salary": 60000,
    "availability": "Full-time",
    "education": [
        {
            "degree": "Bachelor's Degree",
            "school": "Example University",
            "location": "Example City",
            "grad_year": 2020
        }
    ],
    "work_history": [
        {
            "position": "Software Engineer",
            "company": "Tech Company",
            "location": "Tech City",
            "start_year": 2020,
            "end_year": 2022,
            "responsibilities": [
                "Developed web applications",
                "Collaborated with team"
            ]
        }
    ],
    "created_at": "2023-08-21T11:29:26.3212879-05:00",
    "expires_at": "2023-08-21T11:29:26.3212879-05:00"
}
```

<a name="edit-application"></a>
## Edit Application: `/users/:user_id/applications/:application:_id`

Edit a job post's details. User's can only edit their own applications.

### Access Level
`USER`
### Method

`PATCH`

### URL Parameters

| Parameter | Type   | Description            |
|-----------|--------|------------------------|
| `post_id`      | string | post ID           |

### Request Parameters

#### ApplicationUpdateInput

| Field          | Type      | Description                 |
|----------------|-----------|-----------------------------|
| `first_name`   | string    | First name of applicant     |
| `last_name`    | string    | Last name of applicant      |
| `email`        | string    | Email address of applicant  |
| `phone`        | string    | Phone number of applicant   |
| `address`      | string    | Address of applicant        |
| `city`         | string    | City of applicant           |
| `state`        | string    | State of applicant          |
| `postal_code`  | string    | Postal code of applicant    |
| `cover_letter` | string    | Cover letter of applicant   |
| `resume_url`   | string    | URL to applicant's resume   |
| `linkedin_url` | string    | LinkedIn profile URL        |
| `portfolio_url`| string    | Portfolio URL of applicant  |
| `desired_salary` | float64 | Desired salary of applicant |
| `availability` | string    | Availability status         |
| `education`    | array     | Educational background      |
| `referrals`    | array     | Referral information        |
| `work_history` | array     | Work experience details     |

#### Referral

| Field     | Type   | Description          |
|-----------|--------|----------------------|
| `name`    | string | Name of the referral |
| `phone`   | string | Phone of the referral|
| `relation`| string | Relation to applicant|
| `title`   | string | Title of the referral |

#### EducationInfo

| Field     | Type   | Description        |
|-----------|--------|--------------------|
| `degree`  | string | Degree             |
| `school`  | string | School attended    |
| `location`| string | Location of school |
| `grad_year`| int   | Graduation year    |

#### WorkExperience

| Field         | Type   | Description            |
|---------------|--------|------------------------|
| `position`   | string | Position               |
| `company`    | string | Company                |
| `location`   | string | Location               |
| `start_year` | int    | Starting year          |
| `end_year`   | int    | Ending year            |
| `responsibilities`| array | Responsibilities    |

**Response Body Example:**

```json
{
    "application_id": 3,
    "user_id": 1,
    "post_id": 3,
    "first_name": "Jonathan",
    "last_name": "Doeton",
    "email": "john.doe@example.com",
    "phone": "1234567890",
    "address": "123 Main St",
    "city": "Example City",
    "state": "Example State",
    "postal_code": "12345",
    "cover_letter": "This is my cover letter.",
    "resume_url": "https://example.com/resume.pdf",
    "linkedin_url": "https://linkedin.com/in/johndoe",
    "portfolio_url": "https://johndoe.portfolio.com",
    "referrals": [
        {
            "name": "Cool Referral 1",
            "phone": "",
            "relation": "",
            "title": ""
        },
        {
            "name": "Cool Referral 2",
            "phone": "",
            "relation": "",
            "title": ""
        }
    ],
    "desired_salary": 60000,
    "availability": "Full-time",
    "education": [
        {
            "degree": "Bachelor's Degree",
            "school": "Example University",
            "location": "Example City",
            "grad_year": 2020
        }
    ],
    "work_history": [
        {
            "position": "Software Engineer",
            "company": "Tech Company",
            "location": "Tech City",
            "start_year": 2020,
            "end_year": 2022,
            "responsibilities": [
                "Developed web applications",
                "Collaborated with team"
            ]
        }
    ],
    "created_at": "2023-08-20T18:47:46.774061-05:00",
    "expires_at": "2023-08-21T11:30:57.4266581-05:00"
}
```

<a name="delete-application"></a>
## Delete Application: `/users/:users_id/applications/:application_id`

Allows a user to delete their own applications.

### Access Level
`USER`

### Method

`DELETE`

### URL Parameters

| Parameter | Type   | Description            |
|-----------|--------|------------------------|
| `post_id`      | string | Post ID           |
| `application_id` | string | Application ID  |

**Response Body Example:**

```json
{
    "success": "Application deleted from the database"
}
```
