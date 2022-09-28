## Update User Info
Description : This API will update the user information

### HTTP Request
`PUT/users`

### URL Parameters


### Query Parameters
N/A


### Request Headers
```
Content-Type: application/x-www-form-urlencoded
```

### Request Body
| Parameter | Format | Description                                |
|-----------|--------|--------------------------------------------|
| first_name     | String | first name of user requesting first name reset |
| last_name   | String | user last name that is to be updated     |
| id   | String | id of the user    |


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
    "Message": Updated Successfully "
}
```

### Bad Request Response when field updation failed
```
{
    "Message": "Invalid data. "
}
```

### Forbidden Response when user not present in request
```
{
    "Message": "Authentication failed. Please contact your administrator"
}
```