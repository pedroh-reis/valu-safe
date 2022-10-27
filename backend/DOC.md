# ValuSafe Backend

## Endpoints

### Home
- endpoint: `/`
- statusCode: 200
- response:
```json
{
    "message": "Welcome to ValuSafe Backend!"
}
```

### Change Locker State
- endpoint: `/locker`
- input:
```json
{
    "id": "0" or "1"
}
```
- response:

When the state is changed with sucess (status code = 200)
```json
{
    "message": "Locker 0 unlocked with success."
}
```

When the locker is not found (status code = 404)
```json
{
    "message": "Locker not found."
}
```

When the state fail to be changed (status code = 406)
```json
{
    "message": "Close locker 0's door before locking it."
}
```