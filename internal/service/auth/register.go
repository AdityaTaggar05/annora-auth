package authservice

import (
	"context"
	"crypto/rand"
	"encoding/base64"

	"golang.org/x/crypto/bcrypt"
)

func (s *Service) Register(ctx context.Context, email, password string) error {
	if !isValidEmail(email) {
		return ErrInvalidEmailFormat
	}

	if !isValidPassword(password) {
		return ErrInvalidPasswordFormat
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return err
	}

	err = s.AuthRepo.CreateUser(ctx, email, string(hash))
	if err != nil {
		return ErrUserAlreadyExists
	}

	user, err := s.AuthRepo.GetUserByEmail(ctx, email)
	if err != nil {
		return err
	}

	token, err := generateEmailVerificationToken()
	if err != nil {
		return err
	}

	key := "email_verify:" + token

	err = s.TokenRepo.CreateEmailToken(ctx, key, user.ID, s.EmailTokenTTL)
	if err != nil {
		return err
	}

	go s.Mailer.SendVerificationEmail(email, token)

	return nil
}

func generateEmailVerificationToken() (string, error) {
	b := make([]byte, 32)

	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	return base64.URLEncoding.EncodeToString(b), nil
}
