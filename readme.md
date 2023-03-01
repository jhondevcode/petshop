# Petshop Service

A REST api writted in golang using gin and gorm

## 1. Installation

```bash
$ go get -u github.com/gin-gonic/gin
$ go get -u github.com/gin-contrib/cors
$ go get -u github.com/joho/godotenv
$ go get -u gorm.io/gorm
$ go get -u gorm.io/driver/mysql
```
## 2. Settings
Before running the project, you must create the .env variables file where you will set the following environment variables for the project to work correctly:

| Key             | Example           |
| --------------- | ----------------- |
| **DB_HOST**     | *127.0.0.1*       |
| **DB_PORT**     | *3306*            |
| **DB_USER**     | *admin*           |
| **DB_PASSWORD** | *my_awesome_pswd* |
| **DB_NAME**     | *petshop*         |
| **BCRYPT_COST** | *10*              |
| **JWT_KEY**     | claims_key        |
