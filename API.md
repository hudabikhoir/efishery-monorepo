
# Efishery Tech Test Documentation
This documentation contains the features available in the Efishery Tech Test repository, complete with sample requests and sample responses. 


***Some of the available variables are as follows:***

| Key | Value | Type |
| --- | ------|-------------|
| url_auth | http://127.0.0.1:9000/ |  |
| url_fetch | http://127.0.0.1:9001/ |  |
| token | eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJmcmVzaCI6ZmFsc2UsImlhdCI6MTYyMjk4ODkyMSwianRpIjoiMjJlYzFhYjItZDA2Yi00NzkzLTk3NDktNTI5YjhmODc4ZDQ3IiwidHlwZSI6ImFjY2VzcyIsInN1YiI6IjEyMzQyMzEyMjAiLCJuYmYiOjE2MjI5ODg5MjEsImV4cCI6MTYyMjk4OTgyMX0.jd_wtjRpBsWIrYboQPwU78mmAJuMGgxuuM0BUfTvIEg |  |

<br>

## Indices

* [Auth](#auth)
  * [login](#1-login)
  * [me](#2-me)
  * [registration](#3-registration)

* [Fetch commodity](#fetch-commodity)
  * [all commodities](#1-all-commodities)
  * [report commodities](#2-report-commodities)


--------


## Auth
Auth consists of login, register, and profile. The following features will be used by the user before accessing fetching data


### 1. login
At this endpoint, the user is asked to enter the mobile number and password that was generated at the time of registration.

***Endpoint:***

```bash
Method: POST
Type: RAW
URL: http://127.0.0.1:9000/login
```

***Body:***

```js        
{
    "phone": "081230839311",
    "password": "EP95"
}
```

***More example Requests/Responses:***

##### I. Example Request: success_response

***Body:***

```js        
{
    "phone": "081230839311",
    "password": "EP95"
}
```

##### I. Example Response: success_response
```js
{
    "message": "Logged with phone number 081230839311",
    "access_token": "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJmcmVzaCI6ZmFsc2UsImlhdCI6MTYyMjk5NTk4MSwianRpIjoiMzBhZmI2N2ItYWM1Zi00YWVlLWFjMGItYmY0NDY2MGNiMzc1IiwidHlwZSI6ImFjY2VzcyIsInN1YiI6IjA4MTIzMDgzOTMxMSIsIm5iZiI6MTYyMjk5NTk4MSwiZXhwIjoxNjIyOTk2ODgxfQ.WPGIGMItiQILhQsmglu5KwxqxbATQkDk1C9ZcjdiwrA",
    "refresh_token": "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJmcmVzaCI6ZmFsc2UsImlhdCI6MTYyMjk5NTk4MSwianRpIjoiZDRlOTc4YzgtZGFjOS00ZmU2LTgwNDItOTYwZTc2ZGU0ZjM5IiwidHlwZSI6InJlZnJlc2giLCJzdWIiOiIwODEyMzA4MzkzMTEiLCJuYmYiOjE2MjI5OTU5ODEsImV4cCI6MTYyNTU4Nzk4MX0.gYYxZ-Oki-gyqLLWe6JurLVGzbBnrZUCHQKh6_PtAYs"
}
```

***Status Code:*** 200

<br>

##### II. Example Request: incorrect_username_password_response

***Body:***

```js        
{
    "phone": "081230839311",
    "password": "EP96"
}
```

##### II. Example Response: incorrect_username_password_response
```js
{
    "code": 500,
    "message": "Incorrect username or password"
}
```

***Status Code:*** 200

<br>



### 2. me
Endpoint profile user when they are login. This endpoint is to ensure that the logged in user is correct. You must include JWT token that created when login.

***Endpoint:***

```bash
Method: GET
Type: 
URL: http://127.0.0.1:9000/me
```

***More example Requests/Responses:***

##### I. Example Request: success_response

##### I. Example Response: success_response
```js
{
    "code": "00",
    "message": "Success",
    "data": {
        "phone": "081230839311",
        "name": "Nur Huda Bikhoir",
        "role": 1
    }
}
```

***Status Code:*** 200

<br>

##### II. Example Request: invalid_token_response

##### II. Example Response: invalid_token_response
```js
{
    "msg": "Signature verification failed"
}
```

***Status Code:*** 422

<br>

### 3. registration
At this endpoint the user is asked to enter a name, cellphone number, and role. Password will be generated automatically after registration is complete


***Endpoint:***

```bash
Method: POST
Type: RAW
URL: http://127.0.0.1:9000/registration
```

***Body:***

```js        
{
    "name": "Nur Huda Bikhoir",
    "phone": "081230839311",
    "role": "1"
}
```

***More example Requests/Responses:***

##### I. Example Request: success_response

***Body:***

```js        
{
    "name": "Nur Huda Bikhoir",
    "phone": "081230839311",
    "role": "1"
}
```

##### I. Example Response: success_response
```js
{
    "code": "00",
    "message": "User Nur Huda Bikhoir was created",
    "data": {
        "access_token": "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJmcmVzaCI6ZmFsc2UsImlhdCI6MTYyMjk5NTcyNiwianRpIjoiZWUyMTc1NTMtZTVjMS00ODM0LWE0ZmUtMmZmZDZlYzBiYWQyIiwidHlwZSI6ImFjY2VzcyIsInN1YiI6IjA4MTIzMDgzOTMxMSIsIm5iZiI6MTYyMjk5NTcyNiwiZXhwIjoxNjIyOTk2NjI2fQ.1baqCf2ZC6r_4pvzi2YcIagA4UFndfFkGAGOTj2A5EQ",
        "refresh_token": "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJmcmVzaCI6ZmFsc2UsImlhdCI6MTYyMjk5NTcyNiwianRpIjoiZTc4YmE0OGYtYTJmZS00Yzg4LWJhOWItOWZhY2MwZTIzOTk3IiwidHlwZSI6InJlZnJlc2giLCJzdWIiOiIwODEyMzA4MzkzMTEiLCJuYmYiOjE2MjI5OTU3MjYsImV4cCI6MTYyNTU4NzcyNn0.D8MplEFvd0qAS4VwjaA2N7iZyacUnJ9Pr_nMgE6GPs8",
        "password": "EP95"
    }
}
```

***Status Code:*** 200

<br>

##### II. Example Request: duplicate_phone_number_response

***Body:***

```js        
{
    "name": "Nur Huda Bikhoir",
    "phone": "081230839311",
    "role": "1"
}
```

##### II. Example Response: duplicate_phone_number_response
```js
{
    "code": 500,
    "message": "User Nur Huda Bikhoir with number phone 081230839311 is already exists"
}
```

***Status Code:*** 200

<br>


## Fetch commodity
Displays all commodity data from all over Indonesia along with prices in rupiah and US dollars

### 1. all commodities

***Endpoint:***

```bash
Method: GET
Type: 
URL: http://127.0.0.1:9001/v1/commodities
```

***More example Requests/Responses:***

##### I. Example Request: missing_header_response

##### I. Example Response: missing_header_response
```js
{
    "code": 403,
    "message": "Missing header data"
}
```

***Status Code:*** 403

<br>

##### II. Example Request: token_has_expired_response

##### II. Example Response: token_has_expired_response
```js
{
    "code": 403,
    "message": "Token has expired"
}
```

***Status Code:*** 403

<br>

##### III. Example Request: success_response

##### III. Example Response: success_response
```js
{
    "code": "00",
    "message": "Success",
    "data": [
        {
            "uuid": "0c192840-7ee4-11ea-b3e1-e335da5df3hj",
            "commodity": "Cupang leher kanan",
            "province": "JAWA BARAT",
            "city": "CIMAHI",
            "size": "101",
            "idr": "20100",
            "usd": "1.41",
            "parsed_at": "2021-06-01T23:05:14+07:00",
            "timestamp": "1622563514"
        },
        {
            "uuid": "06d33985-cc46-4829-bc46-ca81efd3d936",
            "commodity": "Arwana",
            "province": "BANTEN",
            "city": "PANDEGLANG",
            "size": "40",
            "idr": "20000",
            "usd": "1.40",
            "parsed_at": "2021-04-14T13:04:55+07:00",
            "timestamp": "1618380295"
        }
    ]
}
```

***Status Code:*** 200

<br>

### 2. report commodities
Displays commodity reports by province and weekly


***Endpoint:***

```bash
Method: GET
Type: 
URL: http://127.0.0.1:9001/v1/commodities/report
```



***More example Requests/Responses:***


##### I. Example Request: success_response



##### I. Example Response: success_response
```js
{
    "code": "00",
    "message": "Success",
    "data": [
        {
            "province": "ACEH",
            "min": 30,
            "max": 50,
            "median": 7,
            "average": 33.57142857142857
        },
        {
            "province": "BALI",
            "min": 30,
            "max": 222,
            "median": 8,
            "average": 77
        },
        {
            "province": "BANDUNG",
            "min": 40,
            "max": 40,
            "median": 0,
            "average": 40
        }
    ]
}
```


***Status Code:*** 200

<br>



##### II. Example Request: invalid_token_response



##### II. Example Response: invalid_token_response
```js
{
    "code": 403,
    "message": "Token has expired"
}
```


***Status Code:*** 403

<br>
