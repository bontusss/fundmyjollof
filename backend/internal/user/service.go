package user

import (
	"context"
	"fmj/internal/auth"
	"fmj/internal/email"
	"fmj/internal/models"
)

type Service interface {
	SetupUserProfile(ctx context.Context, username, name, bio, country string, paymentMethod []string) (*models.User, error)
	ConfirmUserProfileSetup(ctx context.Context, username string) (bool, error)
}

type service struct {
	repo  auth.Repository
	email email.Service
}

func (s *service) SetupUserProfile(ctx context.Context, username, name, bio, country string, paymentMethod []string) (*models.User, error) {
	//_, err := s.repo.FindUserByEmail(email)
	//if err != nil {
	//	return nil, errors.New("user not found")
	//}
	user := &models.User{
		Username:      username,
		FullName:      name,
		Biography:     bio,
		Country:       country,
		PaymentMethod: paymentMethod,
	}
	if err := s.repo.UpdateUser(ctx, user); err != nil {
		return nil, err
	}
	return user, nil
}

func (s *service) ConfirmUserProfileSetup(ctx context.Context, username string) (bool, error) {
	return true, nil
}

func NewService(repo auth.Repository, email email.Service) Service {
	return &service{repo: repo, email: email}
}
