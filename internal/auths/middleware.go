package auths

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte(os.Getenv("JWT_SECRET"))

type Claims struct {
	FirstName string `json:"firstName"`
	ID        int    `json:"id"`
	Mobile    string `json:"mobile"`
	jwt.RegisteredClaims
}

func CreateJSONToken(id int, firstName string, mobile string) (string, error) {
	claims := &Claims{
		ID:        id,
		FirstName: firstName,
		Mobile:    mobile,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt: jwt.NewNumericDate(time.Now()),
			Issuer:   "inventory-hub",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)

	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyJwtToken(c *gin.Context) {

	tokenString, err := c.Cookie("token")

	if err != nil {
		fmt.Println("Token is missing")
		c.Redirect(http.StatusSeeOther, "/login")
		c.Abort()
		return
	}

	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		fmt.Println("Token is invlaid")
		c.Redirect(http.StatusSeeOther, "/login")
		c.Abort()
		return
	}
	fmt.Println("token is valid ", token)
	c.Next()
}
