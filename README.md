## API Doc:

### [POST] `/v1/text/recognize` - recognize text data and return audio/text response

Required: <br/>
```Content-Type header 'application/json'```

Input params:<br/>

```
    {
        Text    string  "text"
        UserID  int     "userID"
    }
```

Response example:<br/>

```
  {
    Status  enum('success','failure')   "status"
    Text    string                      "text"
    URI     string                      "uri"       #uri to audio file with speech
  }
```

### [POST] `/v1/speech/recognize` - recognize ogg data and return audio/text response

Required: <br/>
```Content-Type header 'audio/ogg'```

Input params:<br/>
```Raw ogg data```

Response example:<br/>

```
  {
    Status  enum('success','failure')   "status"
    Text    string                      "text"
    URI     string                      "uri"       #uri to audio file with speech
  }
```
