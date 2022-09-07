## Update User Info
Description : This API will update the user information

### HTTP Request
`PUT/user`

### URL Parameters
/user/{userid}

### Query Parameters
N/A


### Request Headers
```
Content-Type: application/x-www-form-urlencoded
```

### Request Body
| Parameter | Format | Description                                |
|-----------|--------|--------------------------------------------|
| Email     | String | Email Id of user requesting email reset |
| Password   | String | user password that is to be updated     |
| field   | String | field that is to be updated    |


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
    "Message": Field Updated Successfully "
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