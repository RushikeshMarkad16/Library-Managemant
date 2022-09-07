## GET Book API
Description : This API will be used to get all books in library

### HTTP Request
`GET/book`

### URL Parameters
/books/

### Query Parameters
N/A


### Request Headers
```
Content-Type: application/x-www-form-urlencoded
```

### Request Body
N/A


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
    "Message": Book data fetched Successfully "
```

### Bad Request Response when book fetching failed
```
{
    "Message": "Book data fetching failed, please try again"
}
```

### Forbidden Response when role doesn't match
```
{
    "Message": "Access restricted"
}
```
