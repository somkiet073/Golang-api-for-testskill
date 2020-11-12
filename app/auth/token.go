package auth

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	config "github.com/somkiet073/Golang-api-for-testskill/app/utils/configs"
)

var c config.Env
var configs, errC = c.LoadApp()

/** exported CreateToken
*	@param uint32 userID
*	@return string, error
 */
func CreateToken(userID uint32) (string, error) {

	// check error load config
	if errC != nil {
		return "Error Load config.", errC
	}

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userID"] = userID
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	return token.SignedString([]byte(configs.APISecret))

}

func extractToken(r *http.Request) string {
	// เขียนรับ query string ด้วย net/http
	// keys := r.URL.Query() // รับค่า Query string url
	// token := keys.Get("token") // get ค่า จาก query string "token"

	token := mux.Vars(r)["token"]
	return token
}

func parseTokenString(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
	}
	return []byte(configs.APISecret), nil
}

// ExtractTokenID = extractTokenID
func ExtractTokenID(r *http.Request) (uint32, error) {
	tokenString := extractToken(r)
	token, err := jwt.Parse(tokenString, parseTokenString)

	if err != nil {
		return 0, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		uid, err := strconv.ParseUint(fmt.Sprintf("%.0f", claims["user_id"]), 10, 32)
		if err != nil {
			return 0, err
		}
		return uint32(uid), nil
	}

	return 0, nil

}
