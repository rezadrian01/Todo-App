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

### Database Design Questions

**1. What database tables did you create and why?**
- **Describe each table and its purpose:** I created two tables: `categories` and `todos`. The `categories` table stores categorical tags with unique names and colors, allowing users to group tasks visually. The `todos` table is the core of the application, storing task details like `title`, `description`, `priority`, `completed` status, and `due_date`.
- **Explain the relationships between tables:** There is a one-to-many relationship between `categories` and `todos`. The `todos` table contains a `category_id` foreign key referencing `categories.id`. I used `ON DELETE SET NULL` for this relationship so that deleting a category simply removes the tag from associated tasks rather than deleting the tasks themselves.
- **Why did you choose this structure?:** This normalized, 2-table schema satisfies all project constraints without over-engineering. It is simple to maintain, fast to query, and accurately represents the domain logic.

**2. How did you handle pagination and filtering in the database?**
- **What queries did you write for filtering and sorting?:** Filtering and sorting are dynamically constructed using GORM. I check filter presence (e.g., status, priority, category) and chain `.Where()` clauses. For searching, I wrote a PostgreSQL native query using `to_tsvector` and `plainto_tsquery` to match against both the title and description. Sorting is handled by appending `.Order("column ASC/DESC")`.
- **How do you handle pagination efficiently?:** I calculate an `offset = (page - 1) * limit`. Crucially, I execute a `.Count()` on the filtered query *before* applying `.Preload()` or the `Offset`/`Limit` to fetch the accurate total count for the frontend, and then fetch the paginated subset of data.
- **What indexes (if any) did you add and why?:** I added B-Tree indexes on `completed` and `category_id` to speed up filtering on those fields. Most importantly, I added a `GIN` index (`idx_todos_title_desc_fts`) over a combined `to_tsvector` of the `title` and `description` to enable blazing-fast full-text searches.

### Technical Decision Questions

**1. How did you implement responsive design?**
- **What breakpoints did you use and why?:** I primarily targeted a mobile breakpoint (`<768px`) for stacked layouts and a desktop breakpoint (`>768px`) for side-by-side or expansive grid layouts, mirroring standard tablet/mobile device widths.
- **How does the UI adapt on different screen sizes?:** On smaller screens, the filter controls wrap into multiple lines using flexbox `wrap`, and the main Todo table enables horizontal scrolling to prevent the UI from breaking or squishing columns.
- **Which Ant Design components helped with responsiveness?:** The `<Space wrap>` component effortlessly handled responsive flow for the filter inputs. The `<Table>` component's `scroll={{ x: 800 }}` property was vital for preserving table formatting on mobile viewports.

**2. How did you structure your React components?**
- **Explain your component hierarchy:** `App` acts as the layout shell. Inside, `TodoFilters` sits at the top for controls, `CategoryManager` opens as a Drawer for categorical CRUD, and `TodoList` renders the main data table. `TodoForm` is an independent Modal triggered by either the "New Task" button or the edit actions in `TodoList`.
- **How did you manage state between components?:** I used the **React Context API** (`TodoContext.jsx`) instead of prop-drilling. The Context holds the global list of todos, categories, current filters, and pagination state, making it accessible to any component.
- **How did you handle the filtering and pagination state?:** The `filters` state is stored as an object in Context. When a filter changes via `TodoFilters`, it triggers a state update, resets the Context's pagination `current` page to 1, and fires a `useEffect` hook to fetch new data from the API.

**3. What backend architecture did you choose and why?**
- **How did you organize your API routes?:** Using Gin, I created an `/api` router group and mounted RESTful resource paths: `/api/todos` and `/api/categories`, standardizing the URL scheme.
- **How did you structure your code?:** I adopted a **Clean Architecture** pattern. The code is split into `domain` (interfaces/structs), `repository` (GORM database queries), `service` (business logic), and `handler` (Gin HTTP request parsing/responses). This enforces a separation of concerns and makes testing trivial.
- **What error handling approach did you implement?:** Errors propagate up from the repository to the service, and finally to the handler. The handler formats the error into a standardized JSON envelope (`{ "error": "error_type", "message": "specific details" }`) and returns the appropriate HTTP status code (e.g., 400, 404, 500).

**4. How did you handle data validation?**
- **Where do you validate data?:** Data is validated on both the **frontend** and **backend**.
- **What validation rules did you implement?:** Frontend validation (via Ant Design's `Form.Item` rules) prevents submission of empty titles. Backend validation in the Service layer enforces that `title` is not empty, assigns default priorities if missing, and restricts `priority` to explicit enums (`high`, `medium`, `low`).
- **Why did you choose this approach?:** Frontend validation provides immediate, friendly feedback to the user, ensuring a good UX. Backend validation is the ultimate source of truth, protecting the database from malformed data injected via cURL or Postman.

### Testing & Quality Questions

**1. What did you choose to unit test and why?**
- **Which functions/methods have tests?:** I unit tested the `TodoService` and `CategoryService` methods (e.g., `CreateTodo`, `ToggleComplete`, `DeleteCategory`).
- **What edge cases did you consider?:** I tested missing required fields (empty title), invalid enumerations (unsupported priority), and operating on non-existent IDs (deleting or updating a todo that doesn't exist).
- **How did you structure your tests?:** Because of the Clean Architecture, I created a `mock_repo.go` that implements the `TodoRepository` interface entirely in memory. This allowed my unit tests to execute the service-layer business logic swiftly without spinning up a real PostgreSQL database.

**2. If you had more time, what would you improve or add?**
- **What technical debt would you address?:** I would migrate the React frontend from plain JavaScript to strict **TypeScript**. Currently, React Context provides no auto-completion for component consumers, which could cause bugs as the app scales.
- **What features would you add?:** I would add user authentication (JWT) so multiple users can have isolated task lists, and implement a drag-and-drop system to let users manually reorder task priorities.
- **What would you refactor?:** I would extract the filtering parameters out of the React Context and sync them directly with the URL query parameters (e.g., using React Router). This would make specific filter views shareable and bookmarkable for users.
