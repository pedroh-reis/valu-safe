# ValuSafe Backend

## 1 - Endpoints

### 1.1 - /
It is the home default endpoint that returns a welcoming message.
#### 1.1.1 - Request Body
There is no request body.

#### 1.1.2 - Response Body
With status code 200, it returns a welcoming message.
```json
{
    "message": "Welcome to ValuSafe Backend!"
}
```

### 1.2 - /locker
Responsible for changing locker state, from locker to unlocked and vice-versa.

#### 1.2.1 - Request body
```json
{
    "id": string
}
```

#### 1.2.2 - Response body
When the locker is found (status code = 200).
```json
{
    "message": "Locker <id> (unlocked|locked) with success."
}
```

When the locker is not found (status code = 404).
```json
{
    "message": "Locker not found."
}
```

### 1.3 - /locker/state
Responsible for returning the locker state in a certain point of time.

#### 1.3.1 - Request Body
```json
{
    "id": string
}
```

#### 1.3.2 - Response Body
When the locker is found (status code = 200).
```json
{
    "isUnlocked": boolean
}
```

### 1.4 - /locker/stats
Responsible for returning the locker's historic of state changes.
#### 1.4.1 - Request Body
```json
{
    "id": string,
    "timeframe": "second" | "minute" | "hour" | "day" | "month" | "year",
    "value": integer
}
```
For example, if id = 0, timeframe = hour and value = 4, it will return the historic from 4 hour ago until now.

#### 1.4.2 - Response Body
When the locker is found (status code = 200)

```json
{
    "timesUnlocked": integer,
    "history": [
        {
            "unlocked": boolean,
            "timestamp": datetime
        }
    ]
}
```

Example 

```json
{
    "timesUnlocked": 2,
    "history": [
        {
            "unlocked": false,
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