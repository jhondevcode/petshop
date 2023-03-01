package routes

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/jhondevcode/petshop/dto"
	"github.com/jhondevcode/petshop/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func AuthRouter(engine *gin.Engine, db *gorm.DB) {
	router := engine.Group("/api/v1/auth")
	err := db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal(err)
	}

	router.POST("signup", func(c *gin.Context) { signupHandler(c, db) })
	router.POST("login", func(c *gin.Context) { loginHandler(c, db) })
	router.POST("logout", func(c *gin.Context) { logoutHandler(c, db) })
}

func signupHandler(c *gin.Context, db *gorm.DB) {
	var form dto.UserDTO
	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	if err := db.Where("email = ?", form.Email).First(&user).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "User already exists"})
		return
	}

	user, err := createUser(form, db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to create user"})
	} else {
		c.JSON(http.StatusCreated, gin.H{"user": user})
	}
}

func loginHandler(c *gin.Context, db *gorm.DB) {
	var loginForm dto.LoginDTO
	if err := c.ShouldBindJSON(&loginForm); err != nil {
		c.JSON(400, gin.H{"error": "invalid request"})
		return
	}

	if ok, user := verifyUser(loginForm.Email, loginForm.Password, db); !ok {
		c.JSON(401, gin.H{"error": "unauthorized"})
	} else {
		if token, err := createJWT(user); err != nil {
			c.JSON(500, gin.H{"error": "server error"})
		} else {
			c.JSON(200, gin.H{"token": token})
		}
	}
}

func logoutHandler(c *gin.Context, db *gorm.DB) {
	c.JSON(http.StatusOK, gin.H{"state": "logout"})
}

func createUser(form dto.UserDTO, db *gorm.DB) (models.User, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return models.User{}, err
	}

	costStr := os.Getenv("BCRYPT_COST")
	cost, err := strconv.Atoi(costStr)
	if err != nil {
		log.Fatalf("Error, loading bcrypt complexity: %v", err)
		cost = bcrypt.DefaultCost
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(form.Password), cost)
	if err != nil {
		return models.User{}, err
	}

	user := models.User{
		ID:        id,
		FirstName: form.FirstName,
		LastName:  form.LastName,
		Birthday:  form.Birthday,
		Gender:    form.Gender,
		Email:     form.Email,
		Password:  string(passwordHash),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err = db.Create(&user).Error
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func verifyUser(email, password string, db *gorm.DB) (bool, *models.User) {
	var user models.User
	result := db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return false, nil
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return false, nil
	}

	return true, &user
}

func createJWT(user *models.User) (string, error) {
	claims := jwt.MapClaims{
		"id":        user.ID,
		"email":     user.Email,
		"firstName": user.FirstName,
		"lastName":  user.LastName,
		"gender":    user.Gender,
		"birthday":  user.Birthday,
		"exp":       time.Now().Add(time.Hour * 1).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_KEY")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
