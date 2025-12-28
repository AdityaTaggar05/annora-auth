package authrepo

import "context"

func (r *AuthRepository) MarkEmailVerified(ctx context.Context, userID string) error {
	_, err := r.DB.Exec(
		ctx,
		`UPDATE users SET email_verified=true WHERE id=$1`,
		userID,
	)

	return err
}