## Login for a library management role
Description : This api will allow various roles (endusers, admin and superadmin )to login based on their roles


### HTTP Request
`POST /login`

### URL Parameters
N/A

### Query Parameters
N/A


### Request Headers
```
Content-Type: application/x-www-form-urlencoded   //read
```

### Request Body
| Parameter | Format | Description                                |
|-----------|--------|--------------------------------------------|
| Email     | String | to get Email Id of user  |
|Password | String |to get Password of user       |  //check format


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
    "Message":Login Successfull  " , "Token Id " :"eyJhbGciOiJSUzI1NiIsImtpZCI6Ijk0ZTkyNDgwMzc0OTdkYTg5Nzk1MzBlOWZlNThlNjQ0MDAzZmFlNTAiLCJ0eXAiOiJKV1QifQ.eyJBY2Nlc3NUb2tlbiI6ImxldmVsMSIsIkF1dGgiOiJMT0NBTCIsIkVtYWlsIjoic2FnYXIuc29ud2FuZUBqb3Noc29mdHVJRCI6IjE1Mzk2YmNlLTE2ZjYtNDUxOS04N2FiLTQ2MGFkNzNmMzAyMSIsImV4cCI6MTY0NDg0NTMwNSwiaWF0IjoxNjQzNjM1NzA1fQ.LEAE71bSpk-KAM3f5fQjbs7MmlDueiPOHCK1WYh9ZMEcJy2ylJcAhSmKoYNjMa-r2sJnlYFz8rSU2VJb3Jei-LnB_qOSpAfvuAQrdopF3SZOwjfOZyF82pN0EowWYCgIyqwvgEO78vCg9KpFUuwdWq6Ho1zwYKtDUWr9cwf0s6YnzTF7uVCLl5-E8gZNkXQCJtLfkNMCKcvgsY_uFygFqr55munNtgyM2DqWhOsY011zX2jcXpgGWxyX7MhaZXot3yW0JxwoMhKV_YqbYbciMGRhdLttGclKzykBaQ-s-jHLQT-lbUNc_IpmL-LHnUB9XHfE4ilVE3c0d8aVbZJmiQ"
}
```
### Bad Request Response when email doesn't exist
```
{
    "Message": "email does not exist...please create account"
}
```

### Forbidden Response when Password validation failed
```
{
    "Message": "Invalid credentials... Please try again"
}
```