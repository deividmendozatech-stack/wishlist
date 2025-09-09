[![Go CI](https://github.com/deividmendozatech-stack/wishlist/actions/workflows/ci.yml/badge.svg)](https://github.com/deividmendozatech-stack/wishlist/actions/workflows/ci.yml)

# 📚 Wishlist API

REST API written in **Go 1.25** to manage book wishlists.  
It includes **SQLite + GORM** for persistence, authentication with **JWT** (in progress), documentation with **Swagger**, and a **Dockerfile** for simple deployment.

Repository: [https://github.com/deividmendozatech-stack/wishlist](https://github.com/deividmendozatech-stack/wishlist)

---

## 🚀 Run locally

```bash
# Clone the repository
git clone https://github.com/deividmendozatech-stack/wishlist.git
cd wishlist

# Install dependencies
go mod tidy

# Run the server
go run cmd/API/main.go

By default the server exposes:
API: http://localhost:8080/api
Swagger UI: http://localhost:8080/swagger/index.html

🐳 Run with Docker:
docker build -t wishlist-api .
docker run -p 8080:8080 wishlist-api

📖 Swagger Documentation:
Once running, open in your browser:
http://localhost:8080/swagger/index.html

🔑 Main Endpoints:
| Method | Path                                | Description               |
| ------ | ----------------------------------- | ------------------------- |
| POST   | `/api/users/register`               | Register a user           |
| GET    | `/api/users`                        | List registered users     |
| POST   | `/api/wishlist`                     | Create wishlist           |
| GET    | `/api/wishlist`                     | List user wishlists       |
| DELETE | `/api/wishlist/{id}`                | Delete wishlist           |
| POST   | `/api/wishlist/{id}/books`          | Add book to wishlist      |
| GET    | `/api/wishlist/{id}/books`          | List wishlist books       |
| DELETE | `/api/wishlist/{id}/books/{bookID}` | Remove book from wishlist |
| GET    | `/api/books/search?q=<query>`       | Search books (Google API) |



📦 Request/Response Examples:
| Operation       | Request (JSON)                                | Response (JSON)                                                                                                                               |
| --------------- | --------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------- |
| Register user   | `{ "username": "david", "password": "1234" }` | `201 Created`                                                                                                                                 |
| Create wishlist | `{ "name": "Pending Books" }`                 | `201 Created`                                                                                                                                 |
| List wishlists  | *N/A* (GET)                                   | `[{"id":1,"name":"Pending Books"}]`                                                                                                           |
| Delete wishlist | *N/A* (DELETE)                                | `204 No Content`                                                                                                                              |
| Add book        | `{ "title": "Go 101", "author": "Anon" }`     | `201 Created`                                                                                                                                 |
| List books      | *N/A* (GET)                                   | `[{"id":1,"title":"Go 101","author":"Anon"}]`                                                                                                 |
| Delete book     | *N/A* (DELETE)                                | `204 No Content`                                                                                                                              |
| Search books    | `GET /api/books/search?q=golang`              | `[{"title":"The Go Programming Language","author":["Alan Donovan","Brian Kernighan"]},{"title":"Go in Action","author":["William Kennedy"]}]` |



🧪 Run Tests
go test ./... -v
Unit tests included for handlers, services, and repositories with simple mocks.

Coverage example:
Storage Layer: ~70%
Services: ~50%
Handlers: ~37%

📂 Project Structure:
cmd/API           # main.go, entry point
internal/handler  # HTTP handlers and routes
internal/service  # business logic, Models
internal/storage  # repositories (SQLite + GORM)
pkg/auth          # JWT helpers (in progress)
docs              # Swagger auto-generated files


✅ CI/CD:
GitHub Actions runs go test ./... on every push to main.
Build status is visible in the badge above.

🔄 Save and push changes to GitHub:
# Check which files were modified
git status

# Add changes to the staging area
git add .

# Commit with a descriptive message
git commit -m "Brief description of changes"

# Push to the main branch on GitHub
git push origin main

🔎 Google Books Search
The API integrates with Google Books API to fetch public book data.

📡 Endpoint
| Method | Path                | Description                     |
| ------ | ------------------- | ------------------------------- |
| GET    | `/api/books/search` | Search books using Google Books |

📥 Query Parameters
| Name | Type   | Required | Description                 |
| ---- | ------ | -------- | --------------------------- |
| `q`  | string | Yes      | Search term (e.g. `golang`) |

📤 Example Request
curl "http://localhost:8080/api/books/search?q=golang"

📄 Response (200)
[
  {
    "title": "The Go Programming Language",
    "author": ["Alan Donovan", "Brian Kernighan"]
  },
  {
    "title": "Go in Action",
    "author": ["William Kennedy"]
  }
]

❌ Errors
| Code | Message                   | Reason                                                  |
| ---- | ------------------------- | ------------------------------------------------------- |
| 400  | `"missing query param q"` | Missing search parameter `q`                            |
| 500  | `"error ..."`             | Failure to connect or process Google Books API response |



📷 Swagger View
The following shows how the **`/api/books/search`** endpoint is displayed in Swagger UI:
![Swagger Google Books](docs/images/Swagger.png)


✍️ Autor
David Mendoza – @deividmendozatech-stack