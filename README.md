# Mock SES API


This is a simple mock API that simulates the AWS Simple Email Service (SES) using Go and the Gin framework.

## Features

- Send emails (simulated)
- Track the number of emails sent
- Retrieve statistics

## Prerequisites

- Go (version 1.16 or later)
- Gin framework

## Installation

1. Clone the repository:
   ```bash
   git clone <https://github.com/ramudu314/AWS-SES-API.git>
   cd mock-ses-api/backend


2. Install dependencies:

   ````bash
go get -u github.com/gin-gonic/gin
go get -u github.com/gin-contrib/cors


## Running the API
To run the API, use the following command:

<<<<<<< HEAD
bash
=======

  ``bash

>>>>>>> 4c6869e69d57c71ee12949ddec8bba7ad0135638
go run main.go
The server will run on http://localhost:8080.

## API Endpoints
Send Email
Endpoint: POST /sendEmail

Request Body:

json
{
  "to": "recipient@example.com",
  "subject": "Test Email",
  "body": "This is a test email."
}
Response:
json

{
  "message": "Email sent successfully",
  "to": "recipient@example.com"
}
  ## Get Statistics
Endpoint: GET /stats

Response:

json

{
  "emails_sent": 1
}


## Special Rules
Email Warming Up: In a real-world scenario, new email accounts may have limits on the number of emails that can be sent in the first few weeks. This mock API implements a simple warming-up logic where only 10 emails can be sent in the first week (168 hours). After that, email limits are lifted.


## Error Codes
400 Bad Request: Invalid input parameters.
