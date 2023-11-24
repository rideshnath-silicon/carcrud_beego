package models

import (
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/dgrijalva/jwt-go"
	_ "github.com/lib/pq" // PostgreSQL driver
)

func init() {
	orm.RegisterDriver("postgres", orm.DRPostgres)
	orm.RegisterDataBase("default", "postgres", "user=postgres password=root dbname=mydb sslmode=disable")
	orm.RegisterModel(new(Users), new(Car), new(HomeSetting))
	orm.RunSyncdb("default", false, true)
}

// >>>>>>>>>>>>Models For tables Start from Here <<<<<<<<<<<<<<<<<<<<<<
type Users struct {
	Id          uint      `json:"user_id" orm:"pk;auto"`
	FirstName   string    `json:"first_name" orm:"null"`
	LastName    string    `json:"last_name" orm:"null"`
	Email       string    `json:"email" orm:"unique"`
	PhoneNumber string    `json:"phone_number" orm:"null"`
	Country     string    `json:"country"`
	Role        string    `json:"role"`
	Age         int       `json:"age" orm:"size(3)"`
	Password    string    `json:"password"`
	Otp         string    `orm:"null"`
	Verified    string    `orm:"null"`
	CreatedAt   time.Time `orm:"null"`
	UpdatedAt   time.Time `orm:"null"`
	DeletedAt   time.Time `orm:"null"`
}

type CarType string

const (
	Sedan     CarType = "sedan"
	Hatchback CarType = "hatchback"
	SUV       CarType = "SUV"
)

type Car struct {
	Id          uint `json:"car_id" orm:"pk;auto"`
	CarName     string
	CarImage    string `orm:"null"`
	ModifiedBy  string
	Model       string
	Type        CarType   `orm:"column(car_type);type(enum)"`
	CreatedDate time.Time `orm:"null"`
	UpdateAt    time.Time `orm:"null"`
}

type HomeSetting struct {
	Id        uint `orm:"pk;auto"`
	Section   string
	Type      string
	Key       string
	Value     string
	CreatedAt time.Time `orm:"null"`
	UpdateAt  time.Time `orm:"null"`
}

//<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<End Table Models>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>

type UserLoginRequest struct {
	Email    string `json:"username"`
	Password string `json:"password"`
}

type NewUserRequest struct {
	FirstName string `json:"first_name" orm:"null"`
	LastName  string `json:"last_name" orm:"null"`
	Email     string `json:"email" orm:"unique"`
	Country   string `json:"country"`
	Role      string `json:"role"`
	Age       int    `json:"age" orm:"size(3)"`
	Password  string `json:"password" orm:"notnull"`
}
type UpdateUserRequest struct {
	Id        uint   `json:"user_id"`
	FirstName string `json:"first_name" orm:"null"`
	LastName  string `json:"last_name" orm:"null"`
	Email     string `json:"email" orm:"unique"`
	Country   string `json:"country"`
	Role      string `json:"role"`
	Age       int    `json:"age" orm:"size(3)"`
}

type ResetUserPassword struct {
	CurrentPass string `json:"current_password"`
	NewPass     string `json:"new_password"`
	ConfirmPass string `json:"confirm_password"`
}

type JwtClaim struct {
	Email string
	ID    int
	jwt.StandardClaims
}

type UserDetailsRequest struct {
	Id        uint   `json:"user_id"`
	FirstName string `json:"first_name" `
	LastName  string `json:"last_name" `
	Email     string `json:"email"`
	Age       int    `json:"age"`
	Country   string `json:"country"`
}

type SendOtpData struct {
	Username string `json:"username"`
}

type ResetUserPasswordOtp struct {
	Email   string `json:"email"`
	Otp     string `json:"otp"`
	NewPass string `json:"new_password"`
}

type VerifyEmailOTPRequest struct {
	Username string `json:"username"`
	Otp      string `json:"otp"`
}

/// Car request structs

type GetNewCarRequest struct {
	CarName    string  `json:"car_name" form:"car_name"`
	CarImage   string  `json:"car_imag" form:"file"`
	ModifiedBy string  `json:"modified_by" form:"modified_by"`
	Model      string  `json:"model" form:"model"`
	Type       CarType `json:"type" form:"type"`
}

type UpdateCarRequest struct {
	Id         uint    `json:"car_id" form:"car_id"`
	CarName    string  `json:"car_name" form:"car_name"`
	CarImage   string  `json:"car_imag" form:"file"`
	ModifiedBy string  `json:"modified_by" form:"modified_by"`
	Model      string  `json:"model" form:"model"`
	Type       CarType `json:"type" form:"type"`
}

type GetcarRequest struct {
	Id uint `json:"car_id"`
}
type OutgoingCallerID struct {
	PhoneNumber string `json:"phone_number"`
}

type GetCarLike struct {
	Search string `json:"search"`
}

type CarDetailsRequest struct {
	CarName    string  `json:"car_name"`
	CarImage   string  `json:"car_imag"`
	ModifiedBy string  `json:"modified_by"`
	Model      string  `json:"model"`
	Type       CarType `json:"type"`
}

// Home Setting reuests

type InserNewHomeSettingRequest struct {
	Section string `json:"section" form:"section"`
	Type    string `json:"type" form:"type"`
	Key     string `json:"key" form:"key"`
	Value   string `json:"value" form:"value"`
}

type UpdateHomeSetingRequest struct {
	Id      uint   `json:"home_seting_id"`
	Section string `json:"section"`
	Type    string `json:"type"`
	Key     string `json:"key"`
	Value   string `json:"value"`
}
