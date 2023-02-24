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

func TestCheckTimeout(t *testing.T) {
	var timeout int64 = 1200
	checkTime := DateTimeToTimestamp("2023-02-23 13:04:05")
	t.Logf("check time -> %d, timeout -> %d", checkTime, timeout)
	ret := CheckTimeout(checkTime, timeout)
	assert.Equal(t, ret, true, "assert timeout failure")
}
