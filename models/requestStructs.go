package models

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	_ "github.com/lib/pq" // PostgreSQL driver
)

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
	Password    string    `json:"password" orm:"notnull"`
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
	Type        CarType `orm:"column(type);type(enum);default(suv)"`
	CreatedDate time.Time
	UpdateAt    time.Time
}

//<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<End Table Models>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>

type UserLoginRequest struct {
	Email    string `json:"email"`
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
	Email string `json:"email"`
}

type ResetUserPasswordOtp struct {
	Email   string `json:"email"`
	Otp     string `json:"otp"`
	NewPass string `json:"new_password"`
}

/// Car request structs

type GetNewCarRequest struct {
	CarName    string  `json:"car_name"`
	CarImage   string  `json:"car_imag"`
	ModifiedBy string  `json:"modified_by"`
	Model      string  `json:"model"`
	Type       CarType `json:"type"`
}

type GetcarRequest struct {
	Id int `json:"car_id"`
}
