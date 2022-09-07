Description : This API will delete the book

### HTTP Request
`DELETE /{bookid}`

### URL Parameters
/book/delete/{bookid}

### Query Parameters
N/A


### Request Headers
```
Content-Type: application/x-www-form-urlencoded
```

### Request Body
| Parameter | Format | Description                                |
|-----------|--------|--------------------------------------------|
| Name     | String | name of book |
| id   | String | id of specific book that is to be deleted     |


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
    "Message": Book Deleted Successfuly"
}
```

### Bad Request Response when Book deletion failed
```
{
    "Message": "Invalid data, please try again"
}
```

### Bad Request Response when book doesn't exist
```
{
    "Message": "Book does not exist...please enter valid book"
}
```

### Forbidden Response when auth failed
```
{
    "Message": "Access Denied"
}
```