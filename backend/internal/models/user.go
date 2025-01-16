package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID               primitive.ObjectID `bson:"_id"`
	BrandName        string             `bson:"brand_name"`
	Email            string             `bson:"email"`
	Password         string             `bson:"password"`
	Verified         bool               `bson:"verified"`
	VerificationCode uint32             `bson:"verification_code,omitempty"`
	ResetToken       string             `bson:"reset_token,omitempty"`
	ResetTokenExpiry time.Time          `bson:"reset_token_expires,omitempty"`
	CreatedAt        time.Time          `bson:"created_at"`
	UpdatedAt        time.Time          `bson:"updated_at"`
}

type Creator struct {
	ID                   primitive.ObjectID `bson:"_id"`
	User                 User               `bson:"user"`
	DisplayName          string             `bson:"display_name"`
	Bio                  string             `bson:"bio,omitempty"`
	Avatar               string             `bson:"avatar,omitempty"`
	BannerImage          string             `bson:"banner_image,omitempty"`
	PayStackAccountID    string             `bson:"paystack_account_id,omitempty"`
	FlutterWaveAccountID string             `bson:"flutterwave_account_id,omitempty"`
	Facebook             string             `bson:"facebook,omitempty"`
	LinkedIn             string             `bson:"linked_in,omitempty"`
	Instagram            string             `bson:"instagram,omitempty"`
	Twitter              string             `bson:"twitter,omitempty"`
	Website              string             `bson:"website,omitempty"`
	Youtube              string             `bson:"youtube,omitempty"`
	CreatedAt            time.Time          `bson:"created_at"`
	UpdatedAt            time.Time          `bson:"updated_at"`
}
