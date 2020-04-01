package printer

import (
	"github.com/linyuanbin/fmt-conv/xlog"
)

// NewURLPrinter returns a Printer which
// is able to convert a URL to PDF.
func NewURLPrinter(logger xlog.Logger, url string, opts ChromePrinterOptions) Printer {
	return chromePrinter{
		logger: logger,
		url:    url,
		opts:   opts,
	}
}
