package authservice

import "context"

func (s *Service) VerifyEmail(ctx context.Context, token string) error {
	key := "email_verify:" + token
	
	userID, err := s.TokenRepo.VerifyEmailToken(ctx, key)
	if err != nil {
		return err
	}

	return s.AuthRepo.MarkEmailVerified(ctx, userID)
}
