package authservice

import (
	"context"
	"crypto/rand"
	"encoding/base64"

	"golang.org/x/crypto/bcrypt"
)

func (s *Service) Register(ctx context.Context, email, password string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return err
	}

	err = s.AuthRepo.CreateUser(ctx, email, string(hash))
	if err != nil {
		return err
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

	// TODO: Send verification email with the token
	
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
