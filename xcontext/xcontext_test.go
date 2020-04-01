package xcontext

import (
	"errors"
	"github.com/linyuanbin/fmt-conv/test"
	"github.com/linyuanbin/fmt-conv/xerror"
	"github.com/linyuanbin/fmt-conv/xtime"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMustHandleError(t *testing.T) {
	previousErr := errors.New("previous error")
	logger := test.DebugLogger()
	// context should not have an error.
	ctx, cancel := WithTimeout(logger, 5)
	defer cancel()
	err := MustHandleError(ctx, previousErr)
	assert.Equal(t, previousErr, err)
	// should panic.
	ctx, cancel = WithTimeout(logger, 5)
	defer cancel()
	assert.Panics(t, func() {
		MustHandleError(ctx, nil)
	})
	// context should timed out.
	ctx, cancel = WithTimeout(logger, 0.5)
	defer cancel()
	time.Sleep(xtime.Duration(1))
	err = MustHandleError(ctx, previousErr)
	xerr := test.AssertError(t, err)
	assert.Equal(t, xerror.TimeoutCode, xerror.Code(xerr))
	// context should have an error different
	// than context.DeadlineExceeded.
	ctx, cancel = WithTimeout(logger, 5)
	cancel()
	err = MustHandleError(ctx, previousErr)
	xerr = test.AssertError(t, err)
	assert.Equal(t, xerror.InternalCode, xerror.Code(xerr))
}
