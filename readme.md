# Petshop Service

A REST api writted in golang using gin and gorm

## 1. Installation

```bash
go get -u github.com/gin-gonic/gin
go get -u github.com/gin-contrib/cors
go get -u github.com/joho/godotenv
go get -u gorm.io/gorm
go get -u gorm.io/driver/mysql
```

## 2. Settings

Before running the project, you must create the .env variables file where you will set the following environment variables for the project to work correctly:

| Key             | Example           |
| --------------- | ----------------- |
| **DB_HOST**     | _127.0.0.1_       |
| **DB_PORT**     | _3306_            |
| **DB_USER**     | _admin_           |
| **DB_PASSWORD** | _my_awesome_pswd_ |
| **DB_NAME**     | _petshop_         |
| **BCRYPT_COST** | _10_              |
| **JWT_KEY**     | claims_key        |
