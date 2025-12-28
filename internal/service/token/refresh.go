package tokenservice

import (
	"context"
	"fmt"
	"time"

	"github.com/AdityaTaggar05/annora-auth/internal/model"
)

func (s *Service) Refresh(ctx context.Context, oldToken string) (model.TokenPair, error) {
	tokens := model.TokenPair{}

	rt, err := s.TokenRepo.GetRefreshToken(ctx, oldToken)
	if err != nil || rt.Revoked || rt.ExpiresAt.Before(time.Now()) {
		return tokens, fmt.Errorf("Unauthorized Request")
	}

	err = s.TokenRepo.RevokeRefreshToken(ctx, oldToken)
	if err != nil {
		return tokens, err
	}

	refreshToken, err := model.GenerateRefreshToken(rt.UserID, s.Config.RefreshTTL)
	if err != nil {
		return tokens, err
	}
	tokens.RefreshToken = refreshToken.Token

	_ = s.TokenRepo.CreateRefreshToken(ctx, rt.UserID, tokens.RefreshToken, time.Now().Add(s.Config.RefreshTTL))
	tokens.AccessToken, err = model.GenerateJWT(rt.UserID, s.SigningKey, s.Config.AccessTTL)
	if err != nil {
		return tokens, err
	}

	return tokens, nil
}