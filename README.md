# Articles API

A RESTful API built in Go using **Gin** and **GORM**, responsible for managing:

- Users (`Customer`)
- Articles (`Article`)
- Keywords (`Keyword`)


## Project Structure

```
.
├── cmd/server/main.go        # Entry point
├── internal/                 # Application core
│   ├── db/                   # Database connection
│   ├── handler/              # HTTP handlers
│   ├── model/                # Data models
│   ├── repository/           # Database access (GORM)
│   ├── router/               # HTTP routes
│   └── service/              # Business logic
```

## API Routes

### **`/customer`**

| Method | Route | Description |
|--------|-------|-------------|
| GET    | `/customer` | List all users |
| GET    | `/customer/:id` | Get user by ID |
| POST   | `/customer` | Create a new user |
| PUT    | `/customer/:id` | Update a user |
| DELETE | `/customer/:id` | Delete a user |

### **`/article`**

| Method | Route | Description |
|--------|-------|-------------|
| GET    | `/article` | List all articles |
| GET    | `/article/:id` | Get article by ID |
| POST   | `/article` | Create a new article |
| PUT    | `/article/:id` | Update an article |
| DELETE | `/article/:id` | Delete an article |
| GET    | `/article/customer/:customerId` | Get articles by user ID |
| GET    | `/article/title?title=...` | Search by title (LIKE) |
| GET    | `/article/content?content=...` | Search by content (LIKE) |
| GET    | `/article/keywords?kw=kw1&kw=kw2` | Search articles containing *any* keyword |
| GET    | `/article/keywords/filter?kw=...` | Search articles containing *all* keywords |

### **`/keyword`**

| Method | Route | Description |
|--------|-------|-------------|
| GET    | `/keyword` | List all keywords |
| GET    | `/keyword/:id` | Get keyword by ID |
| POST   | `/keyword` | Create a new keyword |
| PUT    | `/keyword/:id` | Update a keyword |
| DELETE | `/keyword/:id` | Delete a keyword |
| GET    | `/keyword/article/:articleId` | List keywords for an article |

## Setup

1. Set variables in `.env`
2. Create the database using `articles_db.sql`
3. Import requests using `requests.postman_collection.json`
4. Run the app:

```bash
go run cmd/server/main.go
```
