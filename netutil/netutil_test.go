package netutil_test

import (
	"testing"

	"github.com/gookit/goutil/netutil"
	"github.com/gookit/goutil/testutil/assert"
)

func TestInternalIP(t *testing.T) {
	assert.NotEmpty(t, netutil.InternalIP())
}
