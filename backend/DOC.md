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
- input body:
```json
{
    "id": "0" | "1"
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

When the state fail to be changed (status code = 425)
```json
{
    "message": "Locker 0 is changing its state. Try it again!"
}
```

### Get Statistics
- endpoint: `/locker/stats`
- input body:
```json
{
    "id": "0" | "1",
    "timeframe": "second" | "minute" | "hour" | "day" | "month" | "year",
    "value": int
}
```
- reponese:

When the statistics is returned (status code = 200)
```json
{
    "timesunlocked": 2,
    "history": [
        {
            "locked": false,
            "timestamp": "2022-10-27T05:06:51.108072Z"
        },
        {
            "locked": true,
            "timestamp": "2022-10-27T05:07:00.723334Z"
        },
        {
            "locked": false,
            "timestamp": "2022-10-27T05:07:06.973785Z"
        }
    ]
}
```

When the timeframe is invalid (status code = 400)
```json
{
    "message": "Invalid timeframe. Check the documentation."
}
```