## Get user information
Description : This API will display user's information

### HTTP Request
`GET/user/{id}`

### URL Parameters
/users/{id}

### Query Parameters
id


### Request Headers
```
Content-Type: application/x-www-form-urlencoded
```

### Request Body
| Parameter | Format | Description                                |
|-----------|--------|--------------------------------------------|
| Email     | String | Email Id of user |



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
    "Message": Users found...details are as follows "
}
```

### Bad Request Response when username validation failed
```
{
    "Message": "Invalid username."
}
```


### Forbidden Response when Email is being accessed by unauthorized user 
```
{
    "Message": "Authentication denied.... Please contact administrator"
}
```