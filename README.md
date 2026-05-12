# Industrix Full Stack Todo App

A full-stack Todo application built with Go (Gin, GORM), PostgreSQL, and React (Vite, Ant Design). Features include category management, complex filtering (full-text search, status, category, priority), and robust pagination.

## Prerequisites

- [Go](https://golang.org/doc/install) (1.20+)
- [Node.js](https://nodejs.org/en/download/) (18+)
- [PostgreSQL](https://www.postgresql.org/download/) (15+)
- *(Optional)* [Docker](https://docs.docker.com/get-docker/) & Docker Compose

## Quick Start (Docker)

If you have Docker installed, you can start the database, backend API, and frontend UI instantly:

```bash
docker-compose up --build -d
```
*Note: This automatically runs database migrations on startup.*

- **Frontend:** Available at `http://localhost:3000`
- **Backend API:** Available at `http://localhost:8080/api`

---

## Manual Setup

### 1. Database Setup
Ensure PostgreSQL is running. Create a database named `industrix_todo`:
```sql
CREATE DATABASE industrix_todo;
```

### 2. Backend Setup
```bash
cd backend
go mod tidy
```
Install the [golang-migrate](https://github.com/golang-migrate/migrate) CLI tool if you don't have it:
```bash
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```
Run migrations:
```bash
migrate -database "postgres://postgres:postgres@localhost:5432/industrix_todo?sslmode=disable" -path migrations up
```
Start the server:
```bash
go run cmd/main.go
```
*The API will be available at `http://localhost:8080/api`*

### 3. Frontend Setup
```bash
cd frontend
npm install
npm run dev
```
*The app will be available at `http://localhost:5174`*

### 4. Running Tests
To run the Go unit tests for the service layer:
```bash
cd backend
go test ./... -v
```

---

## API Documentation

### Todos
- **`GET /api/todos`** - List todos
  - *Query Params:* `page` (default 1), `limit` (default 10), `search` (full-text), `status` (completed/incomplete), `category_id`, `priority`, `sort_by`, `sort_order`.
- **`POST /api/todos`** - Create a todo
  - *Body:* `{ "title": "string", "description": "string", "priority": "high|medium|low", "category_id": int, "due_date": "ISO8601" }`
- **`GET /api/todos/:id`** - Get single todo
- **`PUT /api/todos/:id`** - Update todo
- **`DELETE /api/todos/:id`** - Delete todo
- **`PATCH /api/todos/:id/complete`** - Toggle completion status

### Categories
- **`GET /api/categories`** - List all categories
- **`POST /api/categories`** - Create a category
  - *Body:* `{ "name": "string", "color": "#HEX" }`
- **`PUT /api/categories/:id`** - Update category
- **`DELETE /api/categories/:id`** - Delete category

---

## Technical Questions

**1. Explain your database schema design. Why did you choose this structure?**
I chose a normalized, two-table design (`categories` and `todos`) to fulfill the requirements without over-engineering. The `todos` table holds a foreign key to `categories` (`category_id`) with `ON DELETE SET NULL`, meaning if a category is deleted, the associated todos remain intact but lose their category tag. 

**2. How did you handle full-text search?**
Instead of using a slow `ILIKE '%term%'` query, I implemented PostgreSQL's Native Full-Text Search. I created a GIN index on the `title` column `USING gin(to_tsvector('english', title))`. In the GORM repository, I query this using `to_tsvector('english', title) @@ plainto_tsquery(?)`. This allows for fast, index-backed searching.

**3. Explain your backend architecture.**
I implemented a Clean Architecture pattern:
- **Domain:** Defines the interfaces and data structs (`Todo`, `TodoRepository`, `TodoService`).
- **Repository (GORM):** Speaks strictly to the database.
- **Service:** Contains business logic (validation, defaults) and knows nothing about HTTP or the DB implementation.
- **Handler (Gin):** Parses HTTP requests, calls the service, and returns standard JSON responses.
This separation of concerns makes the service layer easily unit-testable using Mock Repositories.

**4. How did you implement pagination?**
Pagination is implemented in the `TodoRepository.List` method. It receives `Page` and `Limit` in the `TodoFilter` struct. Crucially, I calculate the `OFFSET` via `(page - 1) * limit`. I execute `.Count(&total)` on the query *before* applying `.Preload("Category")` to ensure accurate total counts, then apply `Offset()` and `Limit()` to fetch the slice. The frontend Ant Design Table receives this data and syncs its state via the Context API.

**5. Why did you choose React Context API for state management?**
The requirements specified standard React features without external libraries like Redux. The Context API (`TodoContext.jsx`) allows me to elevate the global state (`todos`, `categories`, `filters`, `pagination`) so it can be accessed by nested components (`TodoFilters`, `TodoList`, `TodoForm`) without prop-drilling. It acts as the single source of truth for the frontend application.

**6. How do you handle error states?**
On the Backend, errors bubble up from the DB/Service to the Handler, which formats them into a standardized JSON envelope (`{ "error": "type", "message": "msg" }`). On the Frontend, Axios calls are wrapped in `try/catch` blocks within the Context. If an error occurs, I use Ant Design's `message.error()` to display a non-intrusive toast notification to the user.

**7. If you had more time, what would you improve?**
1. **Caching:** Implement Redis to cache the `GET /api/categories` endpoint, as category data changes infrequently.
2. **Authentication:** Add JWT-based user authentication so multiple users can have private Todo lists.
3. **Frontend Testing:** Add Vitest and React Testing Library to test component rendering and context logic.

**8. What was the most challenging part of this implementation?**
Ensuring accurate pagination alongside dynamic, multi-field filtering. GORM can behave unpredictably if you apply a `.Count()` to a query that has a `JOIN` or `.Preload()` attached, as it attempts to count the joined rows. Separating the base filtering query to get the `Count`, and *then* applying the Preload and Pagination limits required careful structuring in `todo_repo.go`.
