package middleware

import (
	"CarCrudDemo/helpers"
	"CarCrudDemo/models"
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte(os.Getenv("JWT_SEC_KEY"))

type MiddlewareController struct {
	beego.Controller
}

func (c *MiddlewareController) Login() {
	var user models.UserLoginRequest
	body := c.Ctx.Input.RequestBody
	err := json.Unmarshal(body, &user)
	if err != nil {
		c.Data["json"] = err.Error()
		c.ServeJSON()
		return
	}
	HashPassWord, err := models.GetUserByEmail(user.Email)
	if err != nil {
		c.Data["json"] = err.Error()
		c.ServeJSON()
		return
	}
	err = helpers.VerifyHashedData(HashPassWord.Password, user.Password)
	if err != nil {
		c.Data["json"] = err.Error()
		c.ServeJSON()
		return
	}
	userData, _ := models.LoginUser(user.Email, HashPassWord.Password)
	if userData.Email == "" && userData.FirstName == "" {
		c.Data["json"] = "unauthorized User"
		c.ServeJSON()
		return
	}
	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &models.JwtClaim{Email: userData.Email, ID: int(userData.Id), StandardClaims: jwt.StandardClaims{
		ExpiresAt: expirationTime.Unix(),
	}}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		c.Data["json"] = err.Error()
		c.ServeJSON()
		return
	}
	data := map[string]interface{}{"User_Data": token.Claims, "Tokan": tokenString}
	c.Data["json"] = data
	c.ServeJSON()
}

func JWTMiddleware(ctx *context.Context) {
	tokenString := ctx.Input.Header("Authorization")
	if tokenString == "" {
		ctx.Output.SetStatus(http.StatusUnauthorized)
		ctx.Output.JSON(map[string]string{"error": "Unauthorized"}, true, false)
		return
	}
	tokenString = tokenString[7:]
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil || !token.Valid {
		ctx.Output.SetStatus(http.StatusUnauthorized)
		ctx.Output.JSON(map[string]string{"error": "Invalid token"}, true, false)
		return
	}
	ctx.Input.SetData("user", token.Claims.(jwt.MapClaims))
}
