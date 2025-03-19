# go_net-http_HS256

Simple Go API example using **net/http** and **JWT HS256** for token generation and verification.

## Features
✅ JWT HS256 Token Generation  
✅ JWT HS256 Token Verification (with expired check)  
✅ Simple RESTful API with Go standard library (net/http)  

---

## 📂 Repository
**Repo Name:** `go_net-http_HS256`

## 🚀 How to Run

### 1. Clone the Repository
```bash
git clone https://github.com/your-username/go_net-http_HS256.git
cd go_net-http_HS256
```

### 2. Build Docker Image
```bash
docker build -t jwt-hs256-example .
```

### 3. Run Docker Container
```bash
docker run -p 8088:8088 jwt-hs256-example
```

---

## 📌 API Endpoints

### ✅ Generate Token
**GET** `http://localhost:8088/generate`

**Response:**
```
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9....
```

### ✅ Verify Token
**GET** `http://localhost:8088/verify?token=PASTE_YOUR_TOKEN`

**Response (If Valid):**
```
✅ Token Valid
Email: vario_n_ramadhan_x@telkomsel.co.id
UserID: fb3fa8a8-5768-11ee-af6c-005056978071
PrivateKey: ae7b72f896f54e649403ec1a53e6f1a4
```

**Response (If Invalid/Expired):**
```
HTTP 401 Unauthorized
Token invalid or expired
```

---

## 🛠 JWT Secret
The JWT secret is hardcoded for this example. In production, use environment variables or a secret manager.

```
Secret Key (example):
ae7b72f896f54e649403ec1a53e6f1a4f7c9b334e1224b3bc9d2d5f2c0f739f1
```

---

## 🧠 Notes
- Token expires in **10 minutes**
- Signing algorithm: **HS256**
- Library: [github.com/golang-jwt/jwt/v5](https://github.com/golang-jwt/jwt)

---

## ✅ Example Curl Commands

### Generate Token
```bash
curl http://localhost:8088/generate
```

### Verify Token
```bash
curl "http://localhost:8088/verify?token=PASTE_YOUR_TOKEN"
```

---
