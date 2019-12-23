package jwtpassport

import (
	"git.championtek.com.tw/go/passport"
	"git.championtek.com.tw/go/responses"
	"github.com/dgrijalva/jwt-go"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
)

type Passport struct {
	cfg *passport.Config
}

func New(cfg *passport.Config) context.Handler {
	p := &Passport{}
	p.cfg = cfg

	_ = passport.Psp.Init(cfg)

	return p.Serve
}

func (p *Passport) Serve(ctx context.Context) {
	tokenRaw, err := passport.Psp.RetrieveTokenFromHeader(ctx)
	if err != nil {
		rs := responses.Responses{
			HeaderCode: iris.StatusBadRequest,
			Status:     responses.JWTTokenNotExist,
			Message:    err,
			Data:       nil,
		}
		rs.JSONBody()
	}

	// validate token or refresh the new token
	token, _, err := passport.Psp.ValidateToken(tokenRaw, true)
	if token == nil && err != nil {
		rs := responses.Responses{
			HeaderCode: iris.StatusBadRequest,
			Status:     responses.JWTParseTokenError,
			Message:    err,
			Data:       nil,
		}
		rs.JSONBody()
	}

	ctx.Header("Authorization", "bearer " + token.Raw)
	ctx.Next()
}

func (p *Passport) GenerateToken() (*jwt.Token, error)  {
	return passport.Psp.GenerateToken()
}