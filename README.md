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

You can install TailwindCSS either using the standalone CLI tool or by using 
npm. Just ensure it's in your path.

```bash
go install github.com/cosmtrek/air@latest # install air
go install github.com/a-h/templ/cmd/templ@latest # install templ
go install github.com/pressly/goose/v3/cmd/goose@latest # install goose
```

3. Set up DB and Environment

If you want to use `.env`, perform the following command:

```bash
cp .env.example .env
```

Ensure that these variables are correct for your environment.

You need to set the following environment variables:

```
GOOSE_DRIVER=postgres
GOOSE_DBSTRING=$DATABASE_URL
GOOSE_MIGRATION_DIR=./db/migrations/
```

The `DATABASE_URL` is the connection string to the postgres database. I 
highly recommend using [direnv](https://github.com/direnv/direnv?tab=readme-ov-file)
if you simply want to define these in your `.env` file.
You can very simply add the following to your `.envrc` to achieve these results.
```
dotenv
```

Direnv will now read your `.env` file for variables.

4. Create database (docker)

```bash
./start-database.sh
```

Or start it however you want. Just ensure you've got the correct connection 
string in your environment variables.

5. Run migrations

```bash
goose up
```

6. Run the project

```bash
air # or go run cmd/server/main.go
```

# Getting Started (Docker Compose) (Unsupported for now)

1. Configure environment variables

```bash
cp .env.example .env
```

Open `.env` in your editor and configure to your liking.

2. Compose

```bash
docker compose up
```

