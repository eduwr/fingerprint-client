# Fingerprint Monorepo

This is a monorepo containing both the frontend and backend for the Fingerprint application.

## Project Structure

```
.
├── packages/
│   ├── frontend/     # React + TypeScript frontend
│   └── backend/      # Go backend
├── package.json      # Root package.json
└── pnpm-workspace.yaml
```

## Setup

1. Install pnpm globally if you haven't already:
   ```bash
   npm install -g pnpm
   ```

2. Install dependencies:
   ```bash
   pnpm install
   ```

3. For the backend, install Go dependencies:
   ```bash
   cd packages/backend
   go mod tidy
   ```

## Development

### Frontend
```bash
pnpm --filter frontend dev
```

### Backend
```bash
cd packages/backend
go run main.go
```

## Building

### Frontend
```bash
pnpm --filter frontend build
```

### Backend
```bash
cd packages/backend
go build
``` 