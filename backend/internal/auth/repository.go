package auth

import (
	"context"
	"fmj/internal/models"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"

	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	CreateUser(ctx context.Context, user *models.User) error
	FindUserByEmail(email string) (*models.User, error)
	UpdateUser(ctx context.Context, user *models.User) error
	VerifyUser(ctx context.Context, code uint32) error
	FindUserByUsername(username string) (*models.User, error)
	SaveResetToken(ctx context.Context, email string, token uint32, expiry time.Time) error
	ValidateResetToken(ctx context.Context, token string) (string, error)
	UpdatePassword(ctx context.Context, email, hashedPassword string) error
}

type repository struct {
	db  *mongo.Database
	ctx context.Context
}

// FindUserByUsername implements Repository.
func (r *repository) FindUserByUsername(username string) (*models.User, error) {
	var user models.User
	err := r.db.Collection("users").FindOne(
		r.ctx,
		bson.M{"username": username},
	).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// UpdatePassword implements Repository.
func (r *repository) UpdatePassword(ctx context.Context, email string, hashedPassword string) error {
	_, err := r.db.Collection("users").UpdateOne(
		ctx,
		bson.M{"email": email},
		bson.M{"$set": bson.M{
			"password":            hashedPassword,
			"reset_token":         "",
			"reset_token_expires": nil,
			"updated_at":          time.Now(),
		}},
	)
	return err
}

// ValidateResetToken implements Repository.
func (r *repository) ValidateResetToken(ctx context.Context, token string) (string, error) {
	var user *models.User
	err := r.db.Collection("users").FindOne(ctx, bson.M{
		"reset_token": token,
		"reset_token_expires": bson.M{
			"$gt": time.Now(),
		},
	}).Decode(&user)
	if err != nil {
		return "", err
	}
	return user.Email, nil
}

func (r *repository) SaveResetToken(ctx context.Context, email string, token uint32, expiry time.Time) error {
	_, err := r.db.Collection("users").UpdateOne(
		ctx,
		bson.M{"email": email},
		bson.M{"$set": bson.M{
			"reset_token":         token,
			"reset_token_expires": expiry,
			"updated_at":          time.Now(),
		}},
	)
	return err
}

func (r *repository) CreateUser(ctx context.Context, user *models.User) error {
	user.ID = primitive.NewObjectID()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	_, err := r.db.Collection("users").InsertOne(ctx, user)
	return err
}

func (r *repository) FindUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.db.Collection("users").FindOne(r.ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, mongo.ErrNoDocuments
		}
		return nil, fmt.Errorf("database error: %w", err)
	}
	return &user, nil
}

func (r *repository) UpdateUser(ctx context.Context, user *models.User) error {
	user.UpdatedAt = time.Now()
	_, err := r.db.Collection("users").UpdateOne(
		ctx,
		bson.M{"_id": user.ID},
		bson.M{"$set": user},
	)
	return err
}

func (r *repository) VerifyUser(ctx context.Context, code uint32) error {
	codeStr := fmt.Sprintf("%d", code)
	result, err := r.db.Collection("users").UpdateOne(
		ctx,
		bson.M{"verification_code": codeStr},
		bson.M{"$set": bson.M{"verified": true, "verification_code": ""}},
	)
	if err != nil {
		return err
	}
	if result.MatchedCount == 0 {
		return errors.New("verification code not found")
	}
	return nil
}

func NewRepository(db *mongo.Database, ctx context.Context) Repository {
	return &repository{db, ctx}
}
