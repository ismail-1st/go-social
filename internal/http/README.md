# HTTP Layer (Handlers)

This folder contains the **handlers** responsible for handling HTTP requests and responses.  
Handlers are the **entry point** of the application — they interact with the outside world (HTTP clients) and delegate business logic to the **services**.

---

## Purpose
- Receive HTTP requests.
- Decode/parse request bodies, query params, headers, etc.
- Perform **basic validation** (e.g. malformed JSON).
- Call the appropriate **service** method.
- Encode and return the HTTP response.

---

## Responsibilities
- **What belongs here**:
  - Request parsing (`json.NewDecoder`).
  - Running validation helpers before calling services.
  - Translating service errors into HTTP status codes.
  - Writing HTTP responses (JSON, status code, headers).

- **What does NOT belong here**:
  - Business rules (belongs in **services**).
  - Database queries (belongs in **repositories**).
  - Direct transformations of domain logic.

---

## Benefits
- Clear separation of concerns: HTTP layer only handles request/response logic.
- Services remain agnostic of HTTP and can be reused by other interfaces (e.g. CLI, gRPC).
- Easier to test routes/handlers in isolation.

---