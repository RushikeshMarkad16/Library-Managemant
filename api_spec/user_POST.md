## Create new user
Description : This API will create a new user by the superadmin and admin

### HTTP Request
`POST /user`

### URL Parameters
/user/create

### Query Parameters
N/A


### Request Headers
```
Content-Type: application/x-www-form-urlencoded
```

### Request Body
| Parameter | Format | Description                                |
|-----------|--------|--------------------------------------------|
| Email     | String | Email Id of user requesting password reset |
| first_name   | String | first_name of new user      |
| last_name   | String | last_name of new user      |
| Password   | String | set password for user      |
| Gender   | String | Gender of user      |
| DOB   | Date  |   Birthdate of user   |
| Address   | String    | Adress of the user    |
| Role   | String | role of user (admin,superadmin,enduser)     |
| Mob_no   | longint |Conact Number of user      |


### Sample cURL request
```
```

### Status codes and errors
| Value | Description           |
|-------|-----------------------|
| 200   | OK                    |
| 400   | Bad Request           |
| 403   | Forbidden             |
| 410   | Gone                  |
| 500   | Internal Server Error |

### Response Headers
N/A

### Success Response Body
```
{
    "Message": User Created Successfully "
}
```

### Bad Request Response when wrong info entered
```
{
    "Message": "Invalid data.Please try again"
}
```

### Bad Request Response when user already exists
```
{
    "Message": "username already exists...please login."
}
```

### Forbidden Response when field is empty
```
{
    "Message": Please enter valid credentials"
}
```
