package main

import (
	"github.com/tliron/commonlog"
	"ziniki.org/deployer/lsp/internal/lsp"

	_ "github.com/tliron/commonlog/simple"
)

func main() {
	commonlog.Configure(2, nil)
	lsp.StdioServer()
}
