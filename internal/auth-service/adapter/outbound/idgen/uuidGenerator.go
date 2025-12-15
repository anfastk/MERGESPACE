package idgen

import (
	"github.com/anfastk/MERGESPACE/internal/auth-service/application/port/outbound"
	"github.com/google/uuid"
)

type UUIDGenerator struct{}

var _ outbound.IDGenerator = (*UUIDGenerator)(nil)

func NewUUIDGenerator() outbound.IDGenerator {
	return &UUIDGenerator{}
}

func (g *UUIDGenerator) NewID() string {
	return uuid.NewString()
}
