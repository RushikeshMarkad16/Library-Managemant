## Add Book API
Description : This API will be used to get details of book from library

### HTTP Request
`GET/book/{book_id}`

### URL Parameters
/book/create/{book_id}

### Query Parameters
book_id


### Request Headers
```
Content-Type: application/x-www-form-urlencoded
```

### Request Body
| Parameter | Format | Description                                |
|-----------|--------|--------------------------------------------|
| Name     | String | Name of Book |
| Id   | String | Unique ID of book       |

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
    "Message": Found Book Successfully "
```

### Bad Request Response when book addition failed
```
{
    "Message": "Fetching Book failed, please try again"
}
```

### Forbidden Response when role doesn't match
```
{
    "Message": "Access restricted"
}
```
