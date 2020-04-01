package printer

import (
	"fmt"
	"github.com/linyuanbin/fmt-conv/xlog"
	"github.com/mafredri/cdp/protocol/emulation"
)

// NewHTMLPrinter returns a Printer which
// is able to convert an HTML file to PDF.
func NewHTMLPrinter(logger xlog.Logger, fpath string, opts ChromePrinterOptions, devArgs *emulation.SetDeviceMetricsOverrideArgs) Printer {
	URL := fmt.Sprintf("file://%s", fpath)
	return chromePrinter{
		logger:  logger,
		url:     URL,
		opts:    opts,
		devArgs: devArgs,
	}
}
