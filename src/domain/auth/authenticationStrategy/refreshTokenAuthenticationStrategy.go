package authenticationStrategy

import (
	"errors"
	"go-uaa/src/domain/auth/refreshToken"
	"go-uaa/src/domain/user"
)

type RefreshTokenAuthenticationStrategy struct {
	userRepository           user.UserRepository
	refreshTokenDeserializer refreshToken.RefreshTokenDeserializer
}

func (strategy *RefreshTokenAuthenticationStrategy) Authenticate(request *AuthenticationStrategyRequest) (*user.User, error) {
	refreshToken, err := strategy.getRefreshToken(request.RefreshToken)
	if err != nil {
		return nil, err
	}

	user, err := strategy.userRepository.FindByUsername(refreshToken.Sub)
	if err != nil {
		return nil, err
	}
	if len(user.RefreshToken) == 0 {
		return nil, StrategyAuthenticationError{
			Username: request.Username,
			Strategy: "refresh_token",
			Message:  "Refresh token not found",
		}
	}

	if user.RefreshToken != refreshToken.Id {
		if err := strategy.invalidateUserRefreshToken(user); err != nil {
			return nil, err
		}
		return nil, StrategyAuthenticationError{
			Username: request.Username,
			Strategy: "refresh_token",
			Message:  "Refresh token is not valid",
		}
	}
	return user, nil
}

func (strategy *RefreshTokenAuthenticationStrategy) ValidateStrategyRequest(request *AuthenticationStrategyRequest) error {
	if request.RefreshToken == "" {
		return errors.New("Missing refresh token for refresh_token authentication")
	}
	return nil
}

func (strategy *RefreshTokenAuthenticationStrategy) getRefreshToken(serializedRefreshToken string) (*refreshToken.RefreshToken, error) {
	refreshToken, err := strategy.refreshTokenDeserializer.Deserialize(serializedRefreshToken)
	if err != nil {
		return nil, err
	}
	return refreshToken, nil
}

func (strategy *RefreshTokenAuthenticationStrategy) invalidateUserRefreshToken(user *user.User) error {
	user.RefreshToken = ""
	return strategy.userRepository.Save(*user)
}

func NewRefreshTokenAuthenticationStrategy(userRepository user.UserRepository, refreshTokenDeserializer refreshToken.RefreshTokenDeserializer) *RefreshTokenAuthenticationStrategy {
	strategy := RefreshTokenAuthenticationStrategy{
		userRepository:           userRepository,
		refreshTokenDeserializer: refreshTokenDeserializer,
	}
	return &strategy
}
