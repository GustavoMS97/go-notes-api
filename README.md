# Go Notes API

A simple and complete Notes API written in **Go**, following clean architecture principles, built for learning and portfolio purposes.

This project simulates a real-world backend service with:

- ✅ JWT authentication (access + refresh token strategy)
- ✅ Full CRUD for Notes
- ✅ MongoDB for persistence
- ✅ Swagger documentation
- ✅ Environment variables support via `.env`
- ✅ Modular code with proper folder structure (like NestJS-style)
- ✅ Dockerized for easy deployment (e.g. Render, Railway, etc.)

---

## 📌 Goal

> Apply real-world backend architecture using Go.

The goal of this project is to gain hands-on experience with Go by building a production-like API that covers:

- Authentication with hashed passwords
- Middleware for auth protection
- A service-repo pattern
- MongoDB integration with models
- A clean and scalable file/folder structure
- Deployment using Docker

---

## 🧱 Stack

| Layer         | Tech / Tool         |
|---------------|----------------------|
| Language      | Go (v1.21+)          |
| Framework     | Fiber (Express-like) |
| Database      | MongoDB              |
| Auth          | JWT (Access + Refresh) |
| Docs          | Swagger (via swaggo) |
| Validation    | go-playground/validator |
| Logging       | Built-in / log.Println |
| Env Handling  | godotenv             |
| Container     | Docker               |

---

## 📁 Project Structure

```bash
internal/
  auth/              # JWT, password hashing
  user/              # User domain (entity, controller, service, repo)
  note/              # Note domain (entity, controller, service, repo)
  database/          # MongoDB setup
  routes/            # Route registration
  di/                # Dependency injection setup
config/              # Application config
cmd/                 # App bootstrap (main.go)
docs/                # Swagger auto-generated
.env                 # Environment variables
Dockerfile           # Deployment config
Makefile             # Helpers (run, swagger, etc.)
```

---

## 🚀 Running Locally

### 1. Clone the repo
```bash
git clone https://github.com/GustavoMS97/go-notes-api
cd go-notes-api
```

### 2. Setup environment
Create a `.env` file:
```env
PORT=4000
ENV="development"
DATABASE_URL="mongodb+srv://<user>:<pass>@cluster.mongodb.net"
DATABASE_NAME="go-notes-api"
JWT_SECRET="your_super_secret_key"
SWAGGER_HOST="127.0.0.1:4000/api"
```

### 3. Run the app
```bash
make run
```

Swagger docs will be available at: `http://localhost:4000/docs/index.html`

---

## 🐳 Docker

### 1. Build image
```bash
docker build -t go-notes-api .
```

### 2. Run container
```bash
docker run -p 4000:4000 --env-file .env go-notes-api
```

You can now deploy this on platforms like **Render**, **Railway**, **Fly.io**, etc.

---

## 🧠 Future Improvements
- ✅ Add support for user ownership (done)
- ✅ Create, list, update, delete notes (done)
- ✅ Full authentication flow (done)
- [ ] Pagination for notes
- [ ] Public/private toggle on notes
- [ ] Unit testing & e2e tests

---

## ✨ Author
Created by [Gustavo Martins](https://github.com/GustavoMS97) – built as a personal learning project and portfolio demonstration.

