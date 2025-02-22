package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tommjj/ql-kho-lua/internal/core/ports"
)

func TestImplements(t *testing.T) {
	assert.Implements(t, (*ports.IRiceService)(nil), new(riceService))
}
