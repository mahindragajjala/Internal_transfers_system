# **Internal Transfers Application**

A Golang application that facilitates **internal financial transactions between accounts**, with **PostgreSQL** as the database and **Docker Compose** for seamless deployment across operating systems.

---

## **Features**

* Create new accounts with an initial balance
* Query account details and current balance
* Transfer funds between accounts (transaction logging included)
* REST API built with **Gin framework**
* Containerized with **Docker Compose** (App + Postgres)

---

## **Prerequisites**

Before running the project, ensure you have:

* **Docker** (>= 20.10)
  [Install Docker](https://docs.docker.com/get-docker/)
* **Docker Compose** (v2 recommended)
  [Install Docker Compose](https://docs.docker.com/compose/install/)
* **Git** (optional, if cloning from repo)
  [Install Git](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git)

> **Note:** You do **not** need Go installed locally; the app runs inside Docker.

---

## **Project Structure**

```
cmd/ – Entry point(s) for the application

cmd/server/main.go – Starts the HTTP server and initializes dependencies

configs/ – Configuration management

config.go – Loads configuration from files or environment variables

config.yaml – Centralized configuration (e.g., DB connection, ports, log level)

pkg/ – Reusable libraries (not business-specific)

pkg/logger/ – Structured logging setup

pkg/utils/ – Common utility functions (e.g., standardized API responses)

internal/ – Core business logic (protected from external import)

models/ – Domain models/entities

account.go – Account entity definition

transaction.go – Transaction entity definition

repository/ – Database access layer

account_repository.go – DB logic for accounts

transaction_repository.go – DB logic for transactions

services/ – Business logic layer

account_service.go – Account operations (create, update balance, etc.)

transaction_service.go – Transaction operations (transfer, logs, etc.)

handlers/ – HTTP controllers

account_handler.go – API endpoints for account operations

transaction_handler.go – API endpoints for transaction operations

api/ – HTTP routing and middleware

routes.go – Central place for defining API routes and mapping handlers

middleware.go – Middleware for auth, request logging, error recovery

db/ – Database-related files

migrations/ – Incremental SQL migration files (e.g., 001_init.sql, 002_add_constraints.sql)

seed/ – Optional seed data for development/testing

scripts/ – Automation scripts

run_local.sh – Script to run app locally with environment setup

migrate.sh – Script to run database migrations easily

tests/ – Testing structure

unit/ – Unit tests for services (business logic tests)

integration/ – API + DB integration tests

build/ – Build and CI/CD configurations

Dockerfile – Production-ready Dockerfile

(Can also hold GitHub Actions / CI configs)

deployments/ – Deployment manifests

docker-compose.yml – Local multi-container setup (Postgres + App)

k8s-deployment.yaml – Kubernetes deployment configuration

Miscellaneous Files

go.mod / go.sum – Go module files

README.md – Documentation on setup and usage
```

---

## **How to Run the Application**

### 1. **Clone Repository**

```bash
git clone https://github.com/mahindragajjala/Internal_transfers_system.git
cd internal-transfers
```

---

### 2. **Build and Start Containers**

```bash
docker-compose up --build
```

* This will:

  * Build the Go application Docker image
  * Start a PostgreSQL container
  * Run database migrations automatically
  * Start the Gin API server on port `8080`

---

### 3. **Verify Containers Are Running**

```bash
docker ps
```

Expected output:

* `internal_transfers_app` (Go API)
* `internal_transfers_db` (PostgreSQL)

---

## **API Endpoints**

### 1. **Create Account**

**POST** `/accounts`

**Request:**

```json
{
  "account_id": 123,
  "balance": 100.23344
}
```

**Example:**

```bash
curl -X POST http://localhost:8080/accounts \
-H "Content-Type: application/json" \
-d '{"account_id":123,"balance":100.23344}'
```

---

### 2. **Get Account Balance**

**GET** `/accounts/{account_id}`

**Example:**

```bash
curl http://localhost:8080/accounts/123
```

**Response:**

```json
{
  "account_id": 123,
  "balance": 100.23344
}
```

---

### 3. **Transfer Funds**

**POST** `/transactions`

**Request:**

```json
{
  "source_account_id": 123,
  "destination_account_id": 456,
  "amount": 50.12345
}
```

**Example:**

```bash
curl -X POST http://localhost:8080/transactions \
-H "Content-Type: application/json" \
-d '{"source_account_id":123,"destination_account_id":456,"amount":50.12345}'
```

---

## **Database**

* Database runs in `postgres` container.
* Volume is persisted (`postgres_data`) so data is retained even if containers stop.
* Initial tables are created from `db/migrations.sql`.

---

## **Environment Variables**

Set in `docker-compose.yml`:

```yaml
DB_HOST=postgres
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=transfers
```

Modify these if you want custom credentials.

---

## **Stopping the Application**

```bash
docker-compose down
```

This stops and removes containers (keeps volume data).

To remove volume (wipe DB data):

```bash
docker-compose down -v
```

---

### No response from API

* Ensure app container is running:

```bash
docker-compose logs app
```

---
