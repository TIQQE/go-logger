# CW Logger

A small library intended to be used for logging to Cloud Watch Logs. 

By default it logs as JSON. 

The only required property in the struct is Message, it will later be used to set titles of Alarms. 

The `ILogEntry` interface can be used to extend the log message further with additional properties.

The default log message looks like the following.
```go
type LogEntry struct {
    Message string `json:"message"`
    ErrorMessage string `json:"errorMessage,omitempty"`
}
```

It will marshal into 
```json
{
    "message": "A message",
    "errorMessage": "something"
}
```

Leaving ErrorMessage empty will omit the property from the marshalled message.