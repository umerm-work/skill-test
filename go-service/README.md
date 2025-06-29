# Go Service

This Go service is designed to handle API requests and generate PDF reports for students. Below are the instructions to set up and run the application.

## Prerequisites

- Go (version 1.20 or later)
- Node.js (for the backend service)
- A running instance of the Node.js backend service on `http://localhost:5007`

## Setup Instructions

1. Clone the repository:
   ```bash
   git clone <repository-url>
   cd <repository-folder>
   ```

2. Navigate to the `go-service` directory:
   ```bash
   cd go-service
   ```

3. Install dependencies:
   ```bash
   go mod tidy
   ```

4. Run the Go service:
   ```bash
   go run main.go
   ```

The service will start and listen on `http://localhost:8081`.

## API Endpoints

### Generate Student Report

**Endpoint:**
```
GET /api/v1/students/:id/report
```

**Description:**
Generates a PDF report for a student with the given `id`.

**Example Request:**
```
GET http://localhost:8081/api/v1/students/1/report
```

**Response:**
- Content-Type: `application/pdf`
- A downloadable PDF file named `report.pdf`.

## Notes

- Ensure the Node.js backend service is running and accessible at `http://localhost:5007`.
- The Go service fetches student data from the backend before generating the report.

## Troubleshooting

- If you encounter a `404 Not Found` error, ensure the URL is correctly formatted and the backend service is running.
- Check the logs for detailed error messages.

## License

This project is licensed under the MIT License.
