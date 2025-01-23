package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserRole string
type UserStatus string
type PaymentMethod string

const (
	UserRoleUser       UserRole = "user"
	UserRoleAdmin      UserRole = "admin"
	UserRoleSuperAdmin UserRole = "super_admin"
)

const (
	UserStatusActive    UserStatus = "active"
	UserStatusSuspended UserStatus = "suspended"
)

const (
	PaymentMethodMTN         PaymentMethod = "MTN"
	PaymentMethodStripe      PaymentMethod = "Stripe"
	PaymentMethodFlutterWave PaymentMethod = "FlutterWave"
	PaymentMethodPaystack    PaymentMethod = "Paystack"
)

type Analytics struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	PageVisits      int                `bson:"page_visits,omitempty" json:"page_visits"`
	UniqueVisitors  int                `bson:"unique_visitors,omitempty" json:"unique_visitors"`
	VisitorIPs      []string           `bson:"visitor_ips,omitempty" json:"-"`
	SupportersCount int                `bson:"supporters_count,omitempty" json:"supporters_count"`
}

type Supporters struct {
	ID               primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Name             string             `json:"name"`
	Email            string             `json:"email"`
	ContributionType string             `json:"contribution_type"`
}

type User struct {
	ID                   primitive.ObjectID `json:"_id" bson:"_id"`
	Username             string             `bson:"username" json:"username" binding:"required"`
	FullName             string             `bson:"full_name" json:"full_name" binding:"required"`
	Website              string             `bson:"website,omitempty" json:"website"`
	Biography            string             `bson:"biography,omitempty" json:"biography"`
	Country              string             `bson:"country" json:"country" binding:"required"`
	BannerImage          string             `bson:"banner_image,omitempty" json:"banner_image"`
	PaymentMethod        []string           `bson:"payment_method" json:"payment_method" binding:"required"`
	PayStackAccountID    string             `bson:"paystack_account_id,omitempty" json:"-"`
	FlutterWaveAccountID string             `bson:"flutterwave_account_id,omitempty" json:"-"`
	StripeAccountID      string             `bson:"stripe_account_id,omitempty" json:"-"`
	Facebook             string             `bson:"facebook,omitempty" json:"facebook"`
	LinkedIn             string             `bson:"linked_in,omitempty" json:"linked_in"`
	Instagram            string             `bson:"instagram,omitempty" json:"instagram"`
	Twitter              string             `bson:"twitter,omitempty" json:"twitter"`
	Github               string             `bson:"github,omitempty" json:"github"`
	Threads              string             `bson:"threads,omitempty" json:"threads"`
	TikTok               string             `bson:"tiktok,omitempty" json:"tiktok"`
	Supporters           []Supporters       `json:"supporters" bson:"supporters,omitempty"`
	Analytics            Analytics          `bson:"analytics,omitempty" json:"analytics"`
	Email                string             `bson:"email" json:"email" binding:"required"`
	Password             string             `bson:"password" json:"password" binding:"required"`
	Verified             bool               `bson:"verified" json:"verified" binding:"required"`
	VerificationCode     string             `bson:"verification_code,omitempty" json:"-"`
	ResetToken           string             `bson:"reset_token,omitempty" json:"-"`
	ResetTokenExpiry     time.Time          `bson:"reset_token_expires,omitempty" json:"-"`
	Status               UserStatus         `bson:"status" json:"status" binding:"required"`
	Role                 UserRole           `bson:"role" json:"role" binding:"required"`
	CreatedAt            time.Time          `bson:"created_at" json:"created_at" binding:"required"`
	UpdatedAt            time.Time          `bson:"updated_at" json:"updated_at" binding:"required"`
}

type RegisterInputs struct {
	Email string `json:"email" binding:"required"`
	Pass  string `json:"pass" binding:"required"`
}

type VerifyEmailInputs struct {
	Code string `json:"code" binding:"required"`
}

type SetupUserInputs struct {
	Email         string   `bson:"email" json:"email" binding:"required"`
	Username      string   `json:"username" binding:"required"`
	Name          string   `json:"name" binding:"required"`
	Bio           string   `json:"bio" binding:"required"`
	Country       string   `json:"country" binding:"required"`
	PaymentMethod []string `json:"payment_method" binding:"required"`
}
