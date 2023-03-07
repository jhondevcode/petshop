# Petshop Service

A REST api writted in golang using gin and gorm

## 1. Installation

```bash
go mod download
```

## 2. Settings

Before running the project, you must create the .env variables file where you will set the following environment variables for the project to work correctly:

| Key             | Example           | Priority   |
| --------------- | ----------------- | ---------- |
| **DB_HOST**     | _127.0.0.1_       | _required_ |
| **DB_PORT**     | _3306_            | _required_ |
| **DB_USER**     | _admin_           | _required_ |
| **DB_PASSWORD** | _my_awesome_pswd_ | _required_ |
| **DB_NAME**     | _petshop_         | _required_ |
| **BCRYPT_COST** | _10_              | _optional_ |
| **JWT_KEY**     | claims_key        | _required_ |
| **SERVER_PORT** | 8080              | _optional_ |

## 3. Execution

```bash
go run main.go
```

## 4. Build

```bash
go build -o petshop
```

## Running using docker compose

If you have runtimes for the application, you can use the following options in conjunction with docker compose:

- For production:

```bash
docker-compose --env-file .env.prod -f docker-compose.prod.yaml up
```

- For development:

```bash
docker-compose --env-file .env.dev -f docker-compose.dev.yaml up
```

- For testing:

```bash
docker-compose --env-file .env.test -f docker-compose.test.yaml up
```

Remember: to be able to execute the previous commands you must have an .env file and a docker-compose.yaml with respect to the environment where you are going to execute it.
