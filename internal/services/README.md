# Business Logic Layer (Services)

This folder contains the **services** that implement the application's core business logic.  
Services sit between the **handlers/controllers** (HTTP layer) and the **repositories** (data access layer).  

---

## Purpose
- Encapsulate **business rules** (e.g. password hashing, validation beyond basic field checks).
- Transform **DTOs → domain models → DTOs** where necessary.
- Orchestrate multiple repositories if a use case requires data from more than one source.
- Keep handlers/controllers thin by moving all non-HTTP logic here.

---

## Responsibilities
- **Example (User Service)**:
  - Receive a `UserRegister` DTO from the handler.
  - Hash the password using bcrypt.
  - Construct a `models.User` entity.
  - Pass it to the repository for persistence.
  - Return the created domain object (or a response DTO) back to the handler.

- **What NOT to do here**:
  - Do not decode/encode JSON (belongs in handlers).
  - Do not talk directly to the database (belongs in repositories).
  - Do not contain request/response objects.

---

## Benefits
- Keeps the system layered and modular.
- Easier to test (services can be unit-tested without HTTP or DB).
- Allows reusing business logic across different delivery mechanisms (e.g. HTTP API, CLI, gRPC).

---

## Example
```go
// UserService handles business rules for users.
type UserService struct {
    Store store.UserRepository
}

func (s UserService) Create(ctx context.Context, user *dto.UserRegister) (*models.User, error) {
    // Hash password
    hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        return nil, err
    }

    u := models.User{
        Username: user.Username,
        Email:    user.Email,
        Password: string(hashed),
    }

    // Persist via repository
    if err := s.Store.Create(ctx, &u); err != nil {
        return nil, err
    }

    return &u, nil
}
