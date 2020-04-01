package printer

import (
	"github.com/linyuanbin/fmt-conv/conf"
	"github.com/linyuanbin/fmt-conv/test"
	"github.com/linyuanbin/fmt-conv/xerror"
	"github.com/linyuanbin/fmt-conv/xlog"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergePrinter(t *testing.T) {
	var (
		logger xlog.Logger = test.DebugLogger()
		config conf.Config = conf.DefaultConfig()
		fpaths []string    = test.MergeFpaths(t)
		opts   MergePrinterOptions
		dest   string
		p      Printer
		err    error
	)
	// default options.
	opts = DefaultMergePrinterOptions(config)
	p = NewMergePrinter(logger, fpaths, opts)
	dest = test.GenerateDestination()
	err = p.Print(dest)
	assert.Nil(t, err)
	err = os.RemoveAll(dest)
	assert.Nil(t, err)
	// should not be OK as context.Context
	// should timeout.
	opts = DefaultMergePrinterOptions(config)
	opts.WaitTimeout = 0.0
	p = NewMergePrinter(logger, fpaths, opts)
	dest = test.GenerateDestination()
	err = p.Print(dest)
	test.AssertError(t, err)
	assert.Equal(t, xerror.TimeoutCode, xerror.Code(err))
	err = os.RemoveAll(dest)
	assert.Nil(t, err)
}
