# Annora Auth Service

---

A **production-grade authentication microservice** built in **Go**, designed with a strong focus on **security**, **clean architecture**, and **real-world auth workflows**.

This service provides:

- JWT-based authentication using **RS256**
- **Email verification** using Redis-backed one-time tokens
- **Refresh token rotation**
- **Rate-limited email flows**
- **JWKS endpoint for public key discovery**
- Clear separation of concerns (Handler / Service / Repository)

The service is designed to be **independently deployable**, **frontend-agnostic**, and **microservice-ready**.

---

## ✨ Key Features

- 🔐 **Asymmetric JWT signing (RS256)**  
  Auth service signs tokens using a private key; other services verify using public keys.

- 🔑 **JWKS (`/.well-known/jwks.json`) support**  
  Enables safe key rotation and standard token verification.

- ✉️ **Email verification (Redis-based)**  
  Short-lived, single-use tokens stored in Redis with TTL.

- 🔄 **Refresh token rotation**  
  Refresh tokens are opaque, stored in PostgreSQL, and revoked on reuse or logout.

- 🚫 **Login blocked until email verification**  
  Simplifies downstream authorisation and trust model.

- 🧹 **Automatic token cleanup**  
  Redis TTL + background cleanup jobs for persistent tokens.

- 🧱 **Clean layered architecture**  
  Business logic is isolated from transport and infrastructure concerns.

---

## 🧱 Architecture Overview

Handler → Service → Repository
↓
Mailer / Redis

- **Handlers**: HTTP transport & response mapping
- **Services**: Business rules, validation, workflows
- **Repositories**: PostgreSQL / Redis access
- **Mailer**: SMTP-based side-effect boundary

---

## 🔐 Authentication Model

| Token Type               | Purpose                | Lifetime            | Storage       |
| ------------------------ | ---------------------- | ------------------- | ------------- |
| Access Token (JWT)       | API authorization      | Short (e.g. 30 min) | Client memory |
| Refresh Token            | Session continuation   | Long (e.g. 15 days) | PostgreSQL    |
| Email Verification Token | Proof of email control | Short (e.g. 4h)     | Redis         |

---

## 🌐 API Endpoints

### 1️⃣ Register

`POST /register`

Creates a new user and sends a verification email.

**Request**

```json
{
  "email": "user@example.com",
  "password": "password"
}
```

**Response**
| Status                      | Meaning                                             |
| --------------------------- | --------------------------------------------------- |
| `201 Created`               | User created successfully. Verification email sent. |
| `400 Bad Request`           | Invalid email or password format.                   |
| `409 Conflict`              | Email already registered.                           |


**Notes**

- User is created with `email_verified = false`
- Verification email is sent asynchronously

---

### 2️⃣ Verify Email (Clickable Link)

`GET /verify-email?token=...`

Validates email verification token

**Response**
| Status                      | Meaning                    |
| --------------------------- | -------------------------- |
| `200 OK`                    | Email Verified Succesfully |
| `400 Bad Request`           | Expired or Used Token      |


**Notes**

- Token is single-use
- Token is stored in Redis with TTL

---

### 3️⃣ Resend Verification Email

`POST /resend-verification`

Resends verification email (rate-limited).

**Request**

```json
{
  "email": "user@example.com"
}
```

**Response**
| Status                      | Meaning                                 |
| --------------------------- | --------------------------------------- |
| `204 No Content`            | Verification email resent successfully. |
| `400 Bad Request`           | Invalid email format.                   |
| `404 Not Found`             | No account exists for this email.       |
| `429 Too Many Requests`     | Resend limit exceeded.                  |


**Notes**

- Rate-limited using Redis counters
- No user/session required

---

### 4️⃣ Login

`POST /login`

Authenticates a verified user.

**Request**

```json
{
  "email": "user@example.com",
  "password": "password"
}
```

**Success**

```json
{
  "access_token": "...",
  "refresh_token": "..."
}
```

**Response**
| Status                      | Meaning                            |
| --------------------------- | ---------------------------------- |
| `200 OK`                    | Login successful. Tokens returned. |
| `401 Unauthorised`          | Invalid email or password.         |
| `403 Forbidden`             | Email not verified.                |
| `400 Bad Request`           | Invalid request payload.           |

**Notes**

- Login is blocked until the email is verified
- Access token is JWT (RS256)
- Refresh token is opaque

---

### 5️⃣ Refresh Access Token

`POST /refresh`

Rotates the refresh token and issues a new access token.

**Request**

```json
{
  "refresh_token": "..."
}
```

**Success**

```json
{
  "access_token": "...",
  "refresh_token": "..."
}
```

**Response**
| Status                      | Meaning                             |
| --------------------------- | ----------------------------------- |
| `200 OK`                    | New access & refresh tokens issued. |
| `401 Unauthorised`          | Refresh token invalid or revoked.   |
| `400 Bad Request`           | Missing or malformed token.         |


**Notes**

- Refresh token rotation prevents replay attacks
- The old token is revoked on use

---

### 6️⃣ Logout

`POST /logout`

Revokes a refresh token.

**Request**

```json
{
  "refresh_token": "..."
}
```

**Response**
| Status                      | Meaning                         |
| --------------------------- | ------------------------------- |
| `204 No Content`            | Logout successful (idempotent). |
| `400 Bad Request`           | Invalid request payload.        |


**Notes**

- Idempotent
- Only affects the refresh token

---

### 7️⃣ JWKS (Public Keys)

`GET /.well-known/jwks.json`

Returns public keys for JWT verification.

**Success**
```json
{
  "keys": [
    {
      "kty": "RSA",
      "kid": "auth-key-2025-01",
      "use": "sig",
      "alg": "RS256",
      "n": "...",
      "e": "AQAB"
    }
  ]
}
```

**Response**
| Status                      | Meaning                            |
| --------------------------- | ---------------------------------- |
| `200 OK`                    | Public keys returned successfully. |

**Notes**

- Enables key rotation
- Standard-compliant (RFC 7517)

---

## 🔒 Security Considerations

- Passwords hashed using bcrypt
- JWT signed with RSA private key
- Public key exposure via JWKS
- Refresh tokens stored server-side
- Email reset tokens are:
  - Opaque
  - Short-lived
  - Single-use
- Rate limiting using Redis
- No sensitive identifiers exposed to the client

---

## 🚀 Tech Stack

- Go
- Chi (HTTP routing)
- PostgreSQL (persistent state)
- Redis (ephemeral state)
- SMTP (email delivery)
- JWT (RS256)

---

## 📌 Final Notes

It is intentionally designed to be **boring, predictable, and correct** — exactly what an auth service should be.

---

## 📜 License

MIT
