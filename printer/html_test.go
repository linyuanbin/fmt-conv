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

func TestHTMLPrinter(t *testing.T) {
	var (
		logger xlog.Logger = test.DebugLogger()
		config conf.Config = conf.DefaultConfig()
		fpath  string      = test.HTMLFpaths(t)[0]
		opts   ChromePrinterOptions
		dest   string
		p      Printer
		err    error
	)
	// default options.
	opts = DefaultChromePrinterOptions(config)
	p = NewHTMLPrinter(logger, fpath, opts)
	dest = test.GenerateDestination()
	err = p.Print(dest)
	assert.Nil(t, err)
	err = os.RemoveAll(dest)
	assert.Nil(t, err)
	// options with a wait delay.
	opts = DefaultChromePrinterOptions(config)
	opts.WaitDelay = 0.5
	p = NewHTMLPrinter(logger, fpath, opts)
	dest = test.GenerateDestination()
	err = p.Print(dest)
	assert.Nil(t, err)
	err = os.RemoveAll(dest)
	assert.Nil(t, err)
	// options with a page ranges.
	opts = DefaultChromePrinterOptions(config)
	opts.PageRanges = "1"
	p = NewHTMLPrinter(logger, fpath, opts)
	dest = test.GenerateDestination()
	err = p.Print(dest)
	assert.Nil(t, err)
	err = os.RemoveAll(dest)
	assert.Nil(t, err)
	// should not be OK as options have
	// a wrong page ranges.
	opts = DefaultChromePrinterOptions(config)
	opts.PageRanges = "foo"
	p = NewHTMLPrinter(logger, fpath, opts)
	dest = test.GenerateDestination()
	err = p.Print(dest)
	test.AssertError(t, err)
	assert.Equal(t, xerror.InvalidCode, xerror.Code(err))
	err = os.RemoveAll(dest)
	assert.Nil(t, err)
	// should not be OK as context.Context
	// should timeout.
	opts = DefaultChromePrinterOptions(config)
	opts.WaitTimeout = 0.0
	p = NewHTMLPrinter(logger, fpath, opts)
	dest = test.GenerateDestination()
	err = p.Print(dest)
	test.AssertError(t, err)
	assert.Equal(t, xerror.TimeoutCode, xerror.Code(err))
	err = os.RemoveAll(dest)
	assert.Nil(t, err)
}
