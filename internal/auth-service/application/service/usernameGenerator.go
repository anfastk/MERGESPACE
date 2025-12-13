package service

import (
	"context"
	"fmt"
	"regexp"
	"strings"

	"github.com/anfastk/MERGESPACE/internal/auth-service/application/port/outbound"
	"github.com/anfastk/MERGESPACE/internal/auth-service/domain/errs"
)

type UsernameGenerator struct {
	userRepo outbound.UserRepository
}

func NewUsernameGenerator(userRepo outbound.UserRepository) *UsernameGenerator {
	return &UsernameGenerator{userRepo: userRepo}
}

func (g *UsernameGenerator) Generate(ctx context.Context, firstName string, lastName string) (string, error) {

	base := normalize(firstName) + "." + normalize(lastName)

	for i := 0; i < 10; i++ {
		username := base
		if i > 0 {
			username = fmt.Sprintf("%s%d", base, i)
		}

		existing, err := g.userRepo.FindByUsername(ctx, username)
		if err != nil {
			return "", err
		}
		if existing == nil {
			return username, nil
		}
	}

	return "", errs.ErrUsernameGenerationFailed
}

func normalize(s string) string {
	s = strings.ToLower(strings.TrimSpace(s))
	reg := regexp.MustCompile(`[^a-z0-9]`)
	return reg.ReplaceAllString(s, "")
}
