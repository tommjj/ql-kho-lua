package auth

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/tommjj/ql-kho-lua/internal/config"
	"github.com/tommjj/ql-kho-lua/internal/core/domain"
)

func TestJWT(t *testing.T) {
	jwt := NewJWTTokenService(config.Auth{
		SecretKey: "0NnoPraIIutVDnY3RciMgpqsPtKpZS78ugSHML00u+k=",
		Duration:  time.Hour,
	})

	token, err := jwt.CreateToken(&domain.User{
		ID:    1,
		Name:  "fiammetta",
		Phone: "012456889",
		Email: "fiammetta@mail.com",
		Role:  domain.Root,
	})
	if err != nil {
		t.Fatal(err.Error())
	}

	t.Log(token)

	tokenPayload, err := jwt.VerifyToken(token)
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Log(tokenPayload)

	assert.Equal(t, domain.TokenPayload{
		ID:    1,
		Name:  "fiammetta",
		Email: "fiammetta@mail.com",
		Role:  domain.Root,
	}, *tokenPayload, "The two token pay load should be the same.")
}
