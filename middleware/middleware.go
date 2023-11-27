package middleware

import (
	"CarCrudDemo/helpers"
	"CarCrudDemo/models"
	"net/http"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte(beego.AppConfig.String("JWT_SEC_KEY"))

type MiddlewareController struct {
	beego.Controller
}

func (c *MiddlewareController) Login() {
	var user models.UserLoginRequest
	err := helpers.RequestBody(c.Ctx, &user)
	if err != nil {
		helpers.ApiFailure(c.Ctx, err.Error(), 1001)
		return
	}
	HashPassWord, err := models.GetUserByEmail(user.Email)
	if err != nil {
		helpers.ApiFailure(c.Ctx, err.Error(), 1001)

		return
	}
	if HashPassWord.Password == "" {
		helpers.ApiFailure(c.Ctx, "please enter valid Username Or Password ", 1001)
		return
	}
	err = helpers.VerifyHashedData(HashPassWord.Password, user.Password)
	if err != nil {
		helpers.ApiFailure(c.Ctx, err.Error(), 1001)

		return
	}
	userData, _ := models.LoginUser(user.Email, HashPassWord.Password)
	if userData.Email == "" && userData.FirstName == "" {
		helpers.ApiFailure(c.Ctx, "Unauthorized User", 5001)
		return
	}
	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &models.JwtClaim{Email: userData.Email, ID: int(userData.Id), StandardClaims: jwt.StandardClaims{
		ExpiresAt: expirationTime.Unix(),
	}}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		helpers.ApiFailure(c.Ctx, err.Error(), 5001)
		return
	}
	data := map[string]interface{}{"User_Data": token.Claims, "Tokan": tokenString}
	helpers.ApiSuccess(c.Ctx, data, 5000)
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
