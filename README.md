# HNG12 Stage One

This Project is an API that takes a number and returns interesting mathematical properties about it, along with a fun fact.

## Requirements

1. Technology stack
    - Use any programming language or framework of your choice (See Sharp (C #), PHP :elephant:, Python :snake:, Go :runner::skin-tone-5:, Java :coffee:, JS/TS :nauseated_face:)
    - Must be deployed to a publicly accessible endpoint
    - Must handle CORS (Cross-Origin Resource Sharing)
    - Must return responses in JSON format

2. Version Control
    - Code must be hosted on Github
    - Repository must be public
    
## API Specification

- Endpoint: **GET** <your-domain.com>/api/classify-number?number=371**
- Required JSON Response Format (200 OK):
```json
{
    "number": 371,
    "is_prime": false,
    "is_perfect": false,
    "properties": ["armstrong", "odd"],
    "digit_sum": 11,  // sum of its digits
    "fun_fact": "371 is an Armstrong number because 3^3 + 7^3 + 1^3 = 371" //gotten from the numbers API
}
```
- Required JSON Response Format (400 Bad Request)
```json
{
    "number": "alphabet",
    "error": true
}
```

## Acceptance Criteria

### Functionality

- Accepts GET requests with a number parameter.
- Returns JSON in the specified format.
- Accepts all valid integers as the only possible inputs
- Provides appropriate HTTP status codes.

### Code Quality

- Organized code structure.
- Basic error handling and input validation.
- Avoids hardcoded values.

### Documentation

- Complete README.

### Deployment

- Publicly accessible and stable API.
- Fast response time (< 500ms).