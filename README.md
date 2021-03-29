# Overview
This application implements REST API to send email using different mail providers like mailjet.    
P.S: sendgrid or amazon ses for future implementation

## Algorithm / Control Flow
----
1. Uses middleware for basic API authentication   
    If authentication fails, return message as 'Unauthorised',    
    else continue to step 2
2. Uses middleware to log request URL, request and response timestamp and time taken to serve the HTTP request
3. Executes the main handler function to send email using the mail provider passed in the request query parameter.   
    If no query parameter is passed, default is considered as mailjet


## Usage
----
1. Navigate to the repository directory.   
For instance 
```
/Users/saurabhagarwal/go/src/github.com/saurabh-arch/send-email
```
2. Run
```
go build
```

3. Run
```
./send-email
```
  
## API's Defination
----
POST /sendemail   
POST /sendemail?provider=sendgrid

Send email using the mail provided passed as the query parameter. If no query parameter is passed, default is considered as mailjet

**Request Payload**
A JSON object that describes information regarding email to be send.
|Name   |Type       |Description                |Example|
|-----  |-------    |------------                   |-------|
|to     |JSON Object Array|name and email id of recipients| [{"email" : "saurabhagg301@gmail.com"}]
|cc|JSON Object Array|name and email id of recipients| [{"email" : "saurabhagr.developer@gmail.com"}]
|bcc|JSON Object Array|name and email id of recipients| [{"email" : "saurabhagr.developer@gmail.com"}]
|subject| string| email subject | test123
|textPart| string | Text body of email | test 123
|htmlPart| string | HTML body of email | <h3>Dear passenger 1, welcome to <a href=\"https://www.mailjet.com/\">Mailjet</a>!</h3><br />May the delivery force be with you!

## Example / Make your first call
----
```
curl -u abc:123 -sX POST http://localhost:8080/sendemail -d '{"to" : [{"email" : "saurabhagg301@gmail.com"}], "cc" : [{"email" : "saurabhagr.developer@gmail.com"}], "bcc" : [{"email" : "saurabhagarwal532@gmail.com"}], "subject": "test email123", "textPart": "test mail body123"}'
```

```
curl -H "Authorization: Basic YWJjOjEyMw==" -sX POST http://localhost:8080/sendemail -d '{"to" : [{"email" : "saurabhagg301@gmail.com"}], "cc" : [{"email" : "saurabhagr.developer@gmail.com"}], "bcc" : [{"email" : "saurabhagarwal532@gmail.com"}], "subject": "test email123", "textPart": "test mail body123", "htmlPart": "<h3>Dear passenger 1, welcome to <a href=\"https://www.mailjet.com/\">Mailjet</a>!</h3><br />May the delivery force be with you!"}'
```
**Note:** 
1. YWJjOjEyMw== is the base64 encoded string for abc:123
2. If both 'textPart' and 'htmlPart' key are present in the request payload, htmlPart will be taken into consideration 


## Configuration
---
All config related info are stored in 
```
/send-email/config/config.go
```

## Few Additional Points
---
1. As of now, the log output is being directed to stdout but we can direct it to syslog/system.log file or any custom file like send-email.log. The code for the same is kept as commented in "/send-email/logger/logger.go"