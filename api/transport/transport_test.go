package transport

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTransportDo(t *testing.T) {
	transport := New()
	ctx := transport.Do()
	assert.NotNil(t, ctx)
}
