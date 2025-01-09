package auth

import (
	"context"
	"errors"
	"fmj/internal/email"
	"fmj/internal/models"
	"fmj/internal/utils"
	"fmt"
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Register(ctx context.Context, email, password string) error
	Login(email, password string) (*models.User, error)
	VerifyEmail(ctx context.Context, code uint32) error
	ForgotPassword(ctx context.Context, email string) error
	ResetPassword(ctx context.Context, token, newPassword string) error
}

type service struct {
	repo  Repository
	email email.Service
}

// ForgotPassword implements Service.
func (s *service) ForgotPassword(ctx context.Context, email string) error {
	_, err := s.repo.FindUserByEmail(email)
	if err != nil {
		log.Println("email not found")
		return err
	}
	// Generate reset token
	token, err := utils.GenerateCodes()
	if err != nil {
		log.Println("generate codes error: ", err)
		return errors.New("an error occurred, try again")
	}

	// Save token with expiry (24 hours)
	expiry := time.Now().Add(24 * time.Hour)
	if err := s.repo.SaveResetToken(ctx, email, token, expiry); err != nil {
		log.Println("save reset token error")
		return err
	}

	// send reset email
	return s.email.SendPasswordResetEmail(email, token)
}

// ResetPassword implements Service.
func (s *service) ResetPassword(ctx context.Context, token string, newPassword string) error {
	// Validate token and get user email
	email, err := s.repo.ValidateResetToken(ctx, token)
	if err != nil {
		return errors.New("invalid or expired reset token")
	}
	// Hash new password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	// update password
	return s.repo.UpdatePassword(ctx, email, string(hashedPassword))
}

func (s *service) Register(ctx context.Context, email, password string) error {
	// Check if user exists
	existing, _ := s.repo.FindUserByEmail(email)
	if existing != nil {
		return errors.New("email already registered")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	code, err := utils.GenerateCodes()
	if err != nil {
		log.Println("generate codes error: ", err)
		return errors.New("an error occurred, try again")
	}

	// Send verification email first
	fmt.Printf("sending verification email: %s\n", email)
	if err := s.email.SendVerificationEmail(email, code); err != nil {
		return fmt.Errorf("failed to send verification email: %w", err)
	}

	// Create user after email is sent successfully
	user := &models.User{
		Email:            email,
		Password:         string(hashedPassword),
		Verified:         false,
		VerificationCode: code,
	}

	fmt.Printf("creating new user: %s\n", user.Email)
	return s.repo.CreateUser(ctx, user)
}

func (s *service) Login(email, password string) (*models.User, error) {
	user, err := s.repo.FindUserByEmail(email)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	if !user.Verified {
		return nil, errors.New("email not verified")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("invalid credentials")
	}

	return user, nil
}

func (s *service) VerifyEmail(ctx context.Context, code uint32) error {
	if err := s.repo.VerifyUser(ctx, code); err != nil {
		return errors.New("invalid verification code")
	}
	return nil
}

func NewService(repo Repository, emailSvc email.Service) Service {
	return &service{
		repo:  repo,
		email: emailSvc,
	}
}
