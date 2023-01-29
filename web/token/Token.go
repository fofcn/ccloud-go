package token

import (
	"errors"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
)

type AuthTokenService interface {
	CreateToken(TokenPayload) string
	GetToken(string) (*TokenPayload, error)
}

type memauthtokenservice struct {
	tokenstore TokenStore
}

var tokenservice *memauthtokenservice
var once sync.Once

func NewAuthTokenService() AuthTokenService {
	once.Do(func() {
		tokenservice = &memauthtokenservice{
			tokenstore: NewMemTokenStore(),
		}
	})
	return tokenservice
}

func (mem memauthtokenservice) CreateToken(payload TokenPayload) string {
	uuid := uuid.New()
	token := uuid.String()
	mem.tokenstore.StoreToken(token, payload)
	return token
}

func (mem memauthtokenservice) GetToken(token string) (*TokenPayload, error) {
	payload, err := mem.tokenstore.RreadToken(token)
	if err != nil {
		return nil, err
	}

	return payload, nil
}

type TokenPayload struct {
	Expire  time.Time
	Payload map[string]string
}

type TokenStore interface {
	StoreToken(token string, payload TokenPayload)
	RreadToken(token string) (*TokenPayload, error)
}

type memtokenstore struct {
	tokentable map[string]TokenPayload
}

func NewMemTokenStore() TokenStore {
	return &memtokenstore{
		tokentable: make(map[string]TokenPayload),
	}
}

func (store memtokenstore) StoreToken(token string, payload TokenPayload) {
	store.tokentable[token] = payload
}

func (store memtokenstore) RreadToken(token string) (*TokenPayload, error) {
	payload, exists := store.tokentable[token]
	if !exists {
		return nil, errors.New(strings.Join([]string{"token", token, "not exists"}, " "))
	}

	return &payload, nil
}
