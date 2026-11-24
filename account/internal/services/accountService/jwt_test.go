package accountService

import (
	"testing"
	"time"
	"github.com/golang-jwt/jwt/v5"
	"gotest.tools/v3/assert"
)

func TestGenerateJWTSuccess(t *testing.T) {
	secret := []byte("test-secret-key-2651698494998483")
	tokenString, err := GenerateJWT(15, "admin", secret, time.Hour)

	assert.NilError(t, err)
	assert.Assert(t, tokenString != "")

	// + разбор токена и проверка подписи тем же ключом
	data := &TokenData{}

	getSecret := func(token *jwt.Token) (interface{}, error) {
    	return secret, nil
	}
	token, err := jwt.ParseWithClaims(tokenString, data, getSecret, jwt.WithValidMethods([]string{"HS256"}))

	assert.NilError(t, err)
	assert.Assert(t, token.Valid)
	assert.Equal(t, data.Subject, "15")
	assert.Equal(t, data.Role, "admin")
}


func TestGenerateJWTInvalidTTL(t *testing.T) {
	secret := []byte("test-secret-key-2651698494998483")
	token, err := GenerateJWT(15, "admin", secret, 0)

	assert.Error(t, err, "срок действия JWT должен быть положительным")
	assert.Equal(t, token, "")
}

func TestGenerateJWTNegativeTTL(t *testing.T) {
	secret := []byte("test-secret-key-2651698494998483")
	token, err := GenerateJWT(15, "admin", secret, -time.Hour)

	assert.Error(t, err, "срок действия JWT должен быть положительным")
	assert.Equal(t, token, "")
}


func TestValidateJWTSuccess(t *testing.T) {
    secret := []byte("test-secret-key-2651698494998483")
    tokenString, err := GenerateJWT(15, "admin", secret, time.Hour)
    assert.NilError(t, err)
    data, err := ValidateJWT(tokenString, secret)

    assert.NilError(t, err)
    assert.Assert(t, data != nil)
    assert.Equal(t, data.Subject, "15")
    assert.Equal(t, data.Role, "admin")
}

func TestValidateJWTWrongSecret(t *testing.T) {
    secret := []byte("test-secret-key-2651698494998483")
    wrongSecret := []byte("another-secret-key-161911514791748818")

    tokenString, err := GenerateJWT(15, "admin", secret, time.Hour)
    assert.NilError(t, err)
    data, err := ValidateJWT(tokenString, wrongSecret)

    assert.Assert(t, err != nil)
    assert.Assert(t, data == nil)
}

func TestValidateJWTInvalidToken(t *testing.T) {
    secret := []byte("test-secret-key-2651698494998483")
    data, err := ValidateJWT("token26516984", secret)

    assert.Assert(t, err != nil)
    assert.Assert(t, data == nil)
}