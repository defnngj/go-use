package timeconversion

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDateTimeToTimestamp(t *testing.T) {

	ret := DateTimeToTimestamp("2023-02-23 13:04:05")
	t.Logf("timestamp -> %d", ret)
	assert.Equal(t, ret, int64(1677128645), "assert timestamp failure")
}
