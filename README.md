# Modal
## User
```
{
    ID: int64,
    Name: string,
    Email: string,
    Password: string,
}
```

## Product
```
{
    ID: int64,
    Name: string,
    Price: uint64,
    UserID: int64,
    User: User
}
```

# Documentation
## health 
to check current server is alive:

<b>GET</b>
```
https://localhost/api/check/health
```
  
Response (Status: 200)
```
{
   message: "OK!"
}
```

## register
Registering a new user

<b>POST</b>
```
https://localhost/api/auth/register
```

Request Body
```
{
    "name": "Vinh Luong",
    "email": "Vinh@gmail.com",
    "password": "Vinh@123456"
}
```

Response success (Status: 201)
```
{
    "status": true,
    "message": "OK!",
    "errors": null,
    "data": {
        "id": 2,
        "name": "Vinh Luong",
        "email": "Vinh@gmail.com",
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiMiIsImV4cCI6MTY1MTgyMDAwMCwiaWF0IjoxNjIwMjg0MDAwLCJpc3MiOiJhZG1pbiJ9.HtnuWlBaevEO3fHAI4McH5W8axvw_3Og47RUI3m9IyI"
    }
}
```

Response error (Status : 400) [Depends on what error]
```
{
    "status": false,
    "message": "Failed to process request",
    "errors": [
        "Key: 'RegisterRequest.Name' Error:Field validation for 'Name' failed on the 'required' tag"
    ],
    "data": {}
}
```

## login
Authenticate by email and password

<b>POST</b>
```
https://localhost/api/auth/login
```

Request body
```
{
    "email": "Vinh@gmail.com",
    "password": "Vinh@123456"
}
```

Response Success (Status: 200)
```
{
    "status": true,
    "message": "OK!",
    "errors": null,
    "data": {
        "id": 1,
        "name": "Vinh Luong",
        "email": "Vinh@gmail.com",
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiMiIsImV4cCI6MTY1MTgyMDAwMCwiaWF0IjoxNjIwMjg0MDAwLCJpc3MiOiJhZG1pbiJ9.HtnuWlBaevEO3fHAI4McH5W8axvw_3Og47RUI3m9IyI"
    }
}
```

Response error, wrong credential (Status: 401)
```
{
    "status": false,
    "message": "Failed to login",
    "errors": [
        "failed to login. check your credential"
    ],
    "data": {}
}
```


## Get Profile
Get current info from logged user

<b>GET</b>
```
https://localhost/api/user/profile
```

Headers
```
Authorization: yourToken
```

Response success (status: 200)
```
{
    "status": true,
    "message": "OK",
    "errors": null,
    "data": {
        "id": 1,
        "name": "Vinh Luong",
        "email": "Vinh@gmail.com"
    }
}
```

## Update profile
Update user data who logged in

<b>PUT</b>
```
https://localhost/api/user/profile
```

Headers
```
Authorization: yourToken
```

Request Body
```
{
    "name": "Vinh Luong",
    "email": "Vinh@gmail.com"
}
```

Response success (Status: 200)
```
{
    "status": true,
    "message": "OK",
    "errors": null,
    "data": {
        "id": 1,
        "name": "Vinh Luong",
        "email": "Vinh@gmail.com"
    }
}
```


<b>=============================================</b>
## All products (based on user who logged in)
Only shows products by user who logged in

<b>GET</b>
```
https://localhost/api/product
```


Headers
```
Authorization: yourToken
```

Response success (Status: 200)
```
{
    "status": true,
    "message": "OK!",
    "errors": null,
    "data": [
        {
            "id": 2,
            "product_name": "Xiaomi Redmi 10",
            "price": 3000,
            "user": {
                "id": 1,
                "name": "Vinh Luong",
                "email": "Vinh@gmail.com"
            }
        },
        {
            "id": 3,
            "product_name": "Iphone 11 Pro Max",
            "price": 2500,
            "user": {
                "id": 1,
                "name": "Vinh Luong",
                "email": "Vinh@gmail.com"
            }
        }
    ]
}
```

## Create product
Create a product with owner is the user who logged in

<b>POST</b>
```
https://localhost/api/product
```

Headers
```
Authorization: yourToken
```

Request body
```
{
    "name": "Xiaomi Redmi 5",
    "price": 3000
}
```

Response success (Status: 201)
```
{
    "status": true,
    "message": "OK!",
    "errors": null,
    "data": {
        "id": 1,
        "product_name": "Xiaomi Redmi 5",
        "price": 3000,
        "user": {
            "id": 1,
            "name": "Vinh Luong",
            "email": "Vinh@gmail.com"
        }
    }
}
```

## Find one product by id
Find product by id

<b>GET</b>
```
https://localhost/api/product/{id}
```

Headers
```
Authorization: yourToken
```


Response success (Status: 200)
```
{
    "status": true,
    "message": "OK!",
    "errors": null,
    "data": {
        "id": 1,
        "product_name": "Xiaomi Redmi 5",
        "price": 3000,
        "user": {
            "id": 1,
            "name": "Vinh Luong",
            "email": "Vinh@gmail.com"
        }
    }
}
```

## Update product
<b>You can only update your own product If you are trying to update someone else product, it will return error. </b>  

<b>PUT</b>
```
https://localhost/api/product/{id}
```

Request body
```
{
    "name": "Xiaomi Redmi 5 Plus",
    "price": 5000
}
```

Response success (Status: 200)
```
{
    "status": true,
    "message": "OK!",
    "errors": null,
    "data": {
        "id": 1,
        "product_name": "Xiaomi Redmi 5 Plus",
        "price": 5000,
        "user": {
            "id": 1,
            "name": "Vinh Luong",
            "email": "Vinh@gmail.com"
        }
    }
}
```


## Delete product
You can only delete your own product

<b>DELETE</b>
```
https://localhost/api/product/{id}
```

Response success (Status: 200)
```
{
    "status": true,
    "message": "OK!",
    "errors": null,
    "data": {}
}
```

# Schema Image
![schema](https://github.com/luongdinhkhanhvinh/codebase-golang/blob/main/docs/schema.png/?raw=true)

# Swagger Doc 
## Run when new api and override docs
```
swag init  
```

# Swagger Doc 
## Change environment in file .env (use Postgres)
```
    DB_USER=
    DB_PASSWORD=
    DB_NAME=
    DB_HOST=
    DB_PORT=
    JWT_SECRET=
```

## Run command or use Makefile run
```
go run main.go
```

## Build command or use Makefile run
```
go build -o bin/main main.go
```


## Makefile 
```
build:
	go build -o bin/main main.go

run:
	go run main.go
```
