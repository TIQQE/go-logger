# CW Logger

A small library intended to be used for logging to Cloud Watch Logs. 

By default it logs as JSON. 

Initialise the logger with the AWS Request ID and Lambda name, you can find the Request ID in the lambda context and the Lambda name you can get from the runtime environment variable `AWS_LAMBDA_FUNCTION_NAME`.

There is no need to set a value in `EventTime` the logger will set the current time in the correct format by it self.

For INFO messages the only required field to fill in is Message, all the other required fields will be added by the logger.

For ERROR messages the `ErrorCode` field should also be filled in, if it is not then it will default to the contents of the `Message` property which is not ideal.

When an ERROR is logged the `Action` field is also required, however it will default to `Open` if it is not set, so you only need too set it if you intend to use any other action such as `Update` or `Close`.

The default log message looks like the following.

```go
type LogEntry struct {
	EventTime    string                 `json:"eventTime"`
	Message      string                 `json:"message"`
	SourceName   string                 `json:"sourceName"`
	ErrorCode    string                 `json:"errorCode,omitempty"`
	ErrorMessage string                 `json:"errorMessage,omitempty"`
	RequestID    string                 `json:"requestId"`
	LogLevel     string                 `json:"logLevel"`
	Keys         map[string]interface{} `json:"keys,omitempty"`
	Action       AlertAction            `json:"alertAction,omitempty"`
}
```

It will marshal into 
```json
{
    "eventTime": "2019-01-02T15:04:05.12345Z",
    "message": "A message",
    "sourceName": "The Source",
    "errorMessage": "something",
    "requestId": "asdasd-asd123-sasad-asd",
    "logLevel": "INFO",
    "keys": {
        "somekey": "asd",
    }
}
```

Leaving ErrorMessage empty will omit the property from the marshalled message.

A map called Keys can be used to add custom fields to the log message. 
There is a helper function to set the values in the map when needed.
