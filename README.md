# Go Auth System (JWT & Bitmask Permissions)

![Go](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)
![License](https://img.shields.io/badge/license-MIT-blue.svg)
![Build Status](https://img.shields.io/badge/build-passing-brightgreen.svg)
![Maintenance](https://img.shields.io/badge/Maintained%3F-yes-green.svg)

This project is a secure, extensible, and professional authentication and authorization system developed using Go (Golang).

It simulates the **Bitmask** authorization structure used in platforms like Discord. This allows multiple permissions to be stored and managed within a single integer.

## Features

*   **JWT (JSON Web Token)** based authentication.
*   **Bcrypt** for secure password hashing.
*   **Bitmask Authorization System**: Flexible and performant permission control.
*   **Middleware** structure for protected routes.
*   **Clean Architecture**: Separation of Models, Handlers, and Utils.
*   **Unit Tests**: Comprehensive test scenarios.

## Installation

1.  Clone or download the project.
2.  Install the necessary Go modules:
    ```bash
    go mod tidy
    ```
3.  Start the server:
    ```bash
    go run main.go
    ```

## Permission System

In this system, permissions are defined as powers of 2 (bits). You can assign permissions to a user by summing the desired values.

| Permission Name | Value (Decimal) | Bit Value | Description |
| :--- | :--- | :--- | :--- |
| `PermRead` | 1 | `0001` | Read permission |
| `PermWrite` | 2 | `0010` | Write permission |
| `PermDelete` | 4 | `0100` | Delete permission |
| `PermAdmin` | 8 | `1000` | Admin permission |

**Example Combinations:**
*   **Read Only**: `1`
*   **Read + Write**: `1 + 2 = 3`
*   **Full Access (Admin)**: `1 + 2 + 4 + 8 = 15`

## API Usage

### 1. Register

Creates a new user. Permissions are set using the `permissions` field.

*   **URL**: `/register`
*   **Method**: `POST`
*   **Body**:
    ```json
    {
        "username": "user",
        "password": "password123",
        "permissions": 3  // (Read + Write)
    }
    ```

### 2. Login

Logs in and returns a JWT token.

*   **URL**: `/login`
*   **Method**: `POST`
*   **Body**:
    ```json
    {
        "username": "user",
        "password": "password123"
    }
    ```
*   **Response**:
    ```json
    {
        "token": "eyJhbGciOiJIUzI1NiIs..."
    }
    ```

### 3. Accessing Protected Routes

To access the following routes, the token must be sent in the `Authorization` header.

**Header:**
`Authorization: Bearer <TOKEN>`

| Route | Method | Required Permission | Description |
| :--- | :--- | :--- | :--- |
| `/home` | GET | - | Accessible by anyone logged in. |
| `/users` | GET | `PermRead (1)` | Lists all users. |
| `/admin` | GET | `PermAdmin (8)` | Accessible only by admins. |

## Testing

To run the unit tests in the project:

```bash
go test ./... -v
```

## Project Structure

*   `main.go`: Server settings and routes.
*   `models/`: Data structures and database simulation.
*   `handlers/`: Functions handling HTTP requests.
*   `utils/`: Helper tools (JWT, etc.).
