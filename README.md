# kyc-integration

Go HTTP service that creates Didit identity verification sessions.

## Requirements

- Go 1.27.0
- A Didit application API key and a published workflow ID

## Setup

1. Clone
2. Create `kyc-integration/.env` with `API_KEY=<your Didit API key>`
3. Install dependencies
   ```bash
   cd kyc-integration && go mod download
   ```
4. Run the server
   ```bash
   cd kyc-integration && go run .
   ```

## Environment Variables

| Name      | Description                     | Required |
| --------- | ------------------------------- | -------- |
| `API_KEY` | Didit application API key       | Yes      |

Loaded from `kyc-integration/.env` via `godotenv`.

## Endpoints

| Method | Path               | Description                                        |
| ------ | ------------------ | -------------------------------------------------- |
| POST   | `/kyc/sessions`    | Creates a Didit verification session and returns it |
| POST   | `/kyc/sessions/:id`| Returns the session's current status               |
| GET    | `/web/`            | Test UI for creating sessions                      |

Request body:

```json
{
  "workflow_id": "<workflow UUID>",
  "vendor_data": "user-123",
  "callback": "https://example.com/verification/callback"
}
```

Only `workflow_id` is required. On success the service responds `201` with the session object from Didit, including the hosted verification `url` to redirect the end user to.

## Project Structure

```
kyc-integration/
├── main.go       # server setup and route wiring
├── models/
│   └── kyc.go    # Didit API client, POST /v3/session/
└── web/
    └── index.html  # test UI for creating sessions
```

## Testing

Run the server from the `kyc-integration/` directory, since the static file path is relative to the working directory:

```bash
cd kyc-integration && go run .
```

Then open **http://localhost:8080/web/** and paste a published KYC `workflow_id` into the form. The page creates a session and opens Didit's hosted verification flow in a modal.

The same page checks a session's status: creating a session fills in the `session_id` field, which you can also paste over. Click **Check status** to read the current status.

The UI must be served over HTTP rather than opened as a file: the verification flow needs camera access, which browsers only allow in a secure context.
