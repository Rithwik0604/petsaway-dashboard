package auth

import (
	"encoding/base64"
	"log"
	"os"
	"petsaway/internal/database"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/securecookie"
	"golang.org/x/crypto/bcrypt"
)

var SessionDuration = time.Hour * 24 * 30 // 1 month
var s *securecookie.SecureCookie
var debugMode bool

type Token struct {
	UserId  string    `json:"user_id"`
	Expires time.Time `json:"expires"`
}

func SetupAuth() {
	secretStr := os.Getenv("SESSION_SECRET")
	hashKey, err := base64.StdEncoding.DecodeString(secretStr)
	if err != nil || len(hashKey) < 32 {
		log.Fatal("Invalid SESSION_SECRET: must be 32-byte base64 string")
	}
	s = securecookie.New(hashKey, nil)

	debugMode, err = strconv.ParseBool(os.Getenv("DEBUG"))
	if err != nil {
		panic(err)
	}
}

func InvalidateToken(c *gin.Context) {
	c.SetCookie("token", "", -1, "/", "", !debugMode, true)
}

func SetCookie(c *gin.Context, token string) {
	c.SetCookie("token", token, int(SessionDuration), "/", "", !debugMode, true)
}

func CreateToken(userId string) (string, error) {
	tokenData := Token{
		UserId:  userId,
		Expires: time.Now().Add(SessionDuration),
	}

	// Encode the struct into a signed string
	return s.Encode("token", tokenData)
}

func TokenIsValid(cookieValue string) (string, bool) {
	var decodedToken Token
	err := s.Decode("token", cookieValue, &decodedToken)
	if err != nil {
		return "", false
	}

	// Check if the current time is after the expiration time
	if time.Now().After(decodedToken.Expires) {
		return "", false
	}

	return decodedToken.UserId, true
}

// Verifies user's password. Returns nil for success and an error for failure
func VerifyPassword(username string, password string) (database.User, error) {
	user, err := database.DB.GetUserByName(database.Qctx, username)
	if err != nil {
		return database.User{}, err
	}

	// if user found match passwords
	return user, bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
}
