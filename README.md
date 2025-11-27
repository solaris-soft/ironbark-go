# Ironbark

A modern practice management application built with Go, featuring a clean architecture and efficient development workflow.

## Overview

Ironbark is a practice management system designed for Australian mining and legal practices. This project serves as both a production application and a template for building Go web applications with best practices.

## Tech Stack

- **Framework**: [Echo](https://echo.labstack.com/) - High performance web framework
- **Database**: PostgreSQL with [pgx](https://github.com/jackc/pgx) driver
- **Templating**: [Templ](https://templ.guide/) - Type-safe templating for Go
- **Code Generation**: [SQLC](https://sqlc.dev/) - Type-safe SQL queries
- **Development**: [Air](https://github.com/cosmtrek/air) - Live reload for Go apps
- **Styling**: Tailwind CSS

## Features

- Contact management system
- Organization management
- Modern, responsive UI
- Type-safe database queries
- Hot reload development environment

## Getting Started

### Prerequisites

- Go 1.24.3 or later
- Docker and Docker Compose
- [Air](https://github.com/cosmtrek/air) (for development)
- [SQLC](https://sqlc.dev/) (for code generation)
- [Templ](https://templ.guide/) (for template generation)

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/solaris-soft/ironbark-go.git
   cd ironbark-go
   ```

2. Start the database:
   ```bash
   docker compose up -d db
   ```

3. Run database migrations:
   ```bash
   make migrate
   ```

4. Generate database code:
   ```bash
   make generate-db
   ```

5. Generate UI templates:
   ```bash
   make generate-ui
   ```

6. Start the development server:
   ```bash
   make dev
   ```

The application will be available at `http://localhost:8080`.

## Development

### Available Make Commands

- `make migrate` - Run database migrations
- `make generate-ui` - Generate Templ templates
- `make generate-db` - Generate SQLC code
- `make dev` - Start development server with hot reload

### Project Structure

```
ironbark-go/
├── cmd/http/          # Application entry point
├── config/            # Configuration management
├── db/                # Database queries and models
├── handlers/          # HTTP handlers
├── services/          # Business logic
├── serializers/       # Data serialization
├── schema/            # Database schema and migrations
├── ui/                # Templ templates and static assets
└── compose.yaml       # Docker Compose configuration
```

## Project Links

- **Project Board**: [Ironbark Practice Management](https://linear.app/solarissoftware/project/ironbark-practice-management-952d275548ed/issues)
- **Initiative**: [Australian Mining and Legal Practice Suite](https://linear.app/solarissoftware/initiative/australian-mining-and-legal-practice-suite-86e0e9b0825e/projects)

