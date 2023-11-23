// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"CarCrudDemo/controllers"
	"CarCrudDemo/middleware"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/login",
			beego.NSRouter("/", &middleware.MiddlewareController{}, "post:Login"),
		),
		beego.NSNamespace("/user",
			// beego.NSBefore(middleware.JWTMiddleware),
			beego.NSRouter("/forgot_pass", &controllers.UserController{}, "post:SendOtp"),
			beego.NSRouter("/reset_pass_otp", &controllers.UserController{}, "post:VerifyOtpResetpassword"),
			beego.NSRouter("/users", &controllers.UserController{}, "get:GetAllUser"),
			beego.NSRouter("/register", &controllers.UserController{}, "post:RegisterNewUser"),
			beego.NSRouter("/update", &controllers.UserController{}, "post:UpdateUser"),
			beego.NSRouter("/reset_pass", &controllers.UserController{}, "post:ResetPassword"),
		),
		beego.NSNamespace("/car",
			// beego.NSBefore(middleware.JWTMiddleware),
			beego.NSRouter("/", &controllers.CarController{}, "post:GetSingleCar"),
			beego.NSRouter("/cars", &controllers.CarController{}, "get:GetAllCars"),
			beego.NSRouter("/create", &controllers.CarController{}, "post:AddNewCar"),
			beego.NSRouter("/update", &controllers.CarController{}, "put:UpdateCar"),
			beego.NSRouter("/delete", &controllers.CarController{}, "delete:DeleteCar"),
		),
	)
	beego.AddNamespace(ns)
}
