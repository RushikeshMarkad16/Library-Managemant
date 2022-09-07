## Update User Info
Description : This API will update the book information

### HTTP Request
`PUT/book/{book_id}`

### URL Parameters
/book/{book_id}

### Query Parameters
book_id


### Request Headers
```
Content-Type: application/x-www-form-urlencoded
```

### Request Body
| Parameter | Format | Description                                |
|-----------|--------|--------------------------------------------|
| book_id     | Int | id of the book to be updated  |
| name   | String | name of the book to be updated      |
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
    "Message": Book Field Updated Successfully "
}
```

### Bad Request Response when field updation failed
```
{
    "Message": "Invalid data. "
}
```

### Forbidden Response when authentication failed
```
{
    "Message": "Authentication failed. Please contact your administrator"
}
```

### Forbidden Response when book not present in request
```
{
    "Message": "Book not found"
}
```