package creator

import (
	"context"
	"fmj/internal/auth"
	"fmj/internal/email"
	"fmj/internal/models"
)

type Service interface {
	SetupUserProfile(
		ctx context.Context,
		email,
		username,
		name,
		bio,
		country string,
		paymentMethod []string) (*models.User, error)
	ConfirmUserProfileSetup(ctx context.Context, username string) (bool, error)
	FindCreatorByUsername(username string) (*models.User, error)
	UpdateAnalytics(username, visitorIP string) error
}

type service struct {
	repo  auth.Repository
	email email.Service
}

func (s *service) SetupUserProfile(
	ctx context.Context,
	email,
	username,
	name,
	bio,
	country string,
	paymentMethod []string) (*models.User, error) {
	user, err := s.repo.FindUserByEmail(email)
	if err != nil {
		return nil, err
	}
	user.Email = email
	user.Username = username
	user.FullName = name
	user.Biography = bio
	user.Country = country
	user.PaymentMethod = paymentMethod

	if err := s.repo.UpdateUser(ctx, user); err != nil {
		return nil, err
	}
	return user, nil
}

func (s *service) ConfirmUserProfileSetup(ctx context.Context, username string) (bool, error) {
	return true, nil
}

func (s *service) FindCreatorByUsername(username string) (*models.User, error) {
	user, err := s.repo.FindUserByUsername(username)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *service) UpdateAnalytics(username, visitorIP string) error {
	return s.repo.UpdateAnalytics(username, visitorIP)
}

func NewService(repo auth.Repository, email email.Service) Service {
	return &service{repo: repo, email: email}
}
