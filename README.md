# SmlLnk (Small Link)

SmlLnk is a URL shortener. I wrote this for fun because I missed Go

# Getting Started (Host)

1. Clone the project

```bash
git clone https://github.com/jackjohn7/smllnk.git
```

2. Install deps ([air](https://github.com/cosmtrek/air),
[goose](https://github.com/pressly/goose),
[templ](https://templ.guide))

```bash
go install github.com/cosmtrek/air@latest # install air
go install github.com/a-h/templ/cmd/templ@latest # install templ
go install github.com/pressly/goose/v3/cmd/goose@latest # install goose
```

3. Create DB Utility

```bash
go build -o dbu cmd/db/main.go
```

4. Configure environment variables

```bash
cp .env.example .env
```

Then configure the variables to your want or need.

5. Create database (docker)

```bash
dbu up

# dbu down # to kill db
```

6. Run migrations

```bash
dbu migrate
```

7. Run the project

```bash
air # or go run cmd/server/main.go
```

# Getting Started (Docker Compose)

1. Configure environment variables

```bash
cp .env.example .env
```

Open `.env` in your editor and configure to your liking.

2. Compose

```bash
docker compose up
```

