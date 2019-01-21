# CW Logger

A small library intended to be used for logging to Cloud Watch Logs. 

By default it logs as JSON. 

The only required property in the struct is Message, it will later be used to set titles of Alarms. 

The `ILogEntry` interface can be used to extend the log message further with additional properties.

The AWSRequestID needs to be set using the Init() function and LogLevel will be set depending on which logging function is used.

The default log message looks like the following.
```go
type LogEntry struct {
	Message      string                 `json:"message"`
	ErrorMessage string                 `json:"errorMessage,omitempty"`
	RequestID    string                 `json:"requestId"`
	LogLevel     string                 `json:"logLevel"`
	Keys         map[string]interface{} `json:"keys,omitempty"`
}
```

It will marshal into 
```json
{
    "message": "A message",
    "errorMessage": "something",
    "requestId": "INFO",
    "logLevel": "asdasd-asd123-sasad-asd",
    "keys": {
        "somekey": "asd",
    }
}
```

Leaving ErrorMessage empty will omit the property from the marshalled message.

A map called Keys can be used to add custom fields to the log message. 
There is a helper function to set the values in the map when needed.
