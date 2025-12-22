package auth

import (
	"context"

	"github.com/AdityaTaggar05/connectify-auth/internal/config"
	"github.com/AdityaTaggar05/connectify-auth/internal/utils"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	UserRepo      *UserRepository
	RefreshRepo *RefreshTokenRepository
	Config config.Config
}

func NewService(DB *pgxpool.Pool, cfg config.Config) *Service {
	return &Service{
		UserRepo: &UserRepository{DB: DB},
		RefreshRepo: &RefreshTokenRepository{DB: DB},
		Config: cfg,
	}
}

func (s *Service) Register(ctx context.Context, email, password string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return err
	}
	return s.UserRepo.CreateUser(ctx, email, string(hash))
}

func (s *Service) Login(ctx context.Context, email, password string) (string, error) {
	user, err := s.UserRepo.GetUserByEmail(ctx, email)
	if err != nil {
		return "", err
	}
	
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", err
	}

	return utils.GenerateJWT(user.ID, s.Config.JWT_SECRET, s.Config.JWT_EXP)
}