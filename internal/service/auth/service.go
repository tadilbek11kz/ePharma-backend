package auth

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/tadilbek11kz/ePharma-backend/internal/app/config"
	"github.com/tadilbek11kz/ePharma-backend/internal/app/store"
	"github.com/tadilbek11kz/ePharma-backend/internal/util/token"
	"golang.org/x/crypto/bcrypt"

	model "github.com/tadilbek11kz/ePharma-backend/pkg/user"
)

//go:generate mockery --name Service
type Service interface {
	CreateUser(ctx context.Context, req model.CreateUserRequest) (user model.User, err error)
	LoginUser(ctx context.Context, req model.LoginUserRequest) (jwt token.Token, err error)
	RefreshToken(ctx context.Context, cookie string) (jwt token.Token, err error)
}

type service struct {
	st *store.RepositoryStore
}

func New(st *store.RepositoryStore) (srv Service) {
	srv = &service{
		st: st,
	}
	srv = WithLogging(srv)
	return
}

func (s *service) CreateUser(ctx context.Context, req model.CreateUserRequest) (user model.User, err error) {
	return s.st.UserRepository.CreateUser(req)
}

func (s *service) LoginUser(ctx context.Context, req model.LoginUserRequest) (jwt token.Token, err error) {
	user, err := s.st.UserRepository.GetUser("email", req.Email)
	if err != nil {
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return
	}

	config := config.NewConfig()

	// Generate Tokens
	jwt.AccessToken, err = token.CreateToken(config.AccessTokenExpiresIn, user.ID, config.AccessTokenPrivateKey)
	if err != nil {
		return
	}
	jwt.RefreshToken, err = token.CreateToken(config.RefreshTokenExpiresIn, user.ID, config.RefreshTokenPrivateKey)
	if err != nil {
		return
	}

	jwt.AccessTokenMaxAge = config.AccessTokenMaxAge * 60
	jwt.RefreshTokenMaxAge = config.RefreshTokenMaxAge * 60

	return
}

func (s *service) RefreshToken(ctx context.Context, cookie string) (jwt token.Token, err error) {
	config := config.NewConfig()

	sub, err := token.ValidateToken(cookie, config.RefreshTokenPublicKey)
	if err != nil {
		return
	}

	user, err := s.st.UserRepository.GetUser("id", uuid.MustParse(fmt.Sprint(sub)))
	if err != nil {
		return
	}

	jwt.AccessToken, err = token.CreateToken(config.AccessTokenExpiresIn, user.ID, config.AccessTokenPrivateKey)
	if err != nil {
		return
	}

	jwt.AccessTokenMaxAge = config.AccessTokenMaxAge * 60

	return
}
