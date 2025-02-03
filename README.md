# Mock SES API


This is a simple mock API that simulates the AWS Simple Email Service (SES) using Go and the Gin framework.

## Deployment URLs
Frontend URL: https://aws-ses-api.vercel.app/

Backend URL: https://aws-ses-api-10.onrender.com/

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

go run main.go

The server will run on http://localhost:8080.

## API Endpoints
1. Send Email
Endpoint: POST /sendEmail

Request Body:


{
  "to": "recipient@example.com",
  "subject": "Test Email",
  "body": "This is a test email."
}

Response:

{
  "message": "Email sent successfully",
  "to": "recipient@example.com"
}


2. Get Statistics
Endpoint: GET /stats

Response:

{
  "emails_sent": 1
}

## Special Rules

Email Warming-Up Logic
In a real-world scenario, new email accounts often have limits on the number of emails that can be sent in the first few weeks. This mock API implements a simple warming-up logic:

->>Initial Limit: Only 10 emails can be sent in the first 168 hours (1 week).

->>After Warming-Up: The email sending limit is lifted, and unlimited emails can be sent.

Error Response for Warming-Up Limit:
If the email sending limit is exceeded during the warming-up period, the API will return:

{
  "error": "Email sending limit exceeded during warming-up period. Please try again later."
}

## Error Codes
Here are the possible error codes returned by the API:

400 Bad Request: Invalid input parameters (e.g., missing or malformed fields in the request body).

429 Too Many Requests: Email sending limit exceeded during the warming-up period.

500 Internal Server Error: An unexpected error occurred on the server.

## Example Usage
Sending an Email

curl -X POST https://aws-ses-api-10.onrender.com/sendEmail \
-H "Content-Type: application/json" \
-d '{
  "to": "recipient@example.com",
  "subject": "Test Email",
  "body": "This is a test email."
}'

Retrieving Statistics


curl -X GET https://aws-ses-api-10.onrender.com/stats


## Future Enhancements
1. Rate Limiting: Implement rate limiting to prevent abuse of the API.

2. Email Templates: Add support for email templates.

3. Authentication: Add API key-based authentication for secure access.

4. Logging: Implement logging for debugging and monitoring purposes.

5. Configuration: Allow configuration of warming-up limits and duration via environment variables.

## Contributing
Contributions are welcome! If you'd like to contribute, please follow these steps:

1. Fork the repository.

2. Create a new branch for your feature or bugfix.

3. Submit a pull request with a detailed description of your changes.

## License
This project is licensed under the MIT License. See the LICENSE file for details.