package sha256

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"hash"
)

const (
	path            = "sha256/"
	secretKeyHeader = "secret_key"
)

type Sha256 struct {
}

func NewSha256(f *fiber.App) Sha256 {
	s := Sha256{}

	f.Post(path, s.Sign)

	return s
}

func newHash(secret string) hash.Hash {
	return hmac.New(sha256.New, []byte(secret))
}

func Sign(secret string, data []byte) string {
	h := newHash(secret)

	h.Write(data)

	return hex.EncodeToString(h.Sum(nil))
}

func (s Sha256) Sign(ctx *fiber.Ctx) error {
	secret := ctx.Get(secretKeyHeader)
	if secret == "" {
		return fmt.Errorf("secret key header is empty")
	}

	body := ctx.Body()
	if body == nil {
		return fmt.Errorf("body is empty")
	}

	resp := Sign(secret, body)

	return ctx.SendString(resp)
}
