package lsp

import (
	"github.com/tliron/glsp/server"
	"ziniki.org/deployer/lsp/internal/document"
	"ziniki.org/deployer/lsp/internal/workspace"

	protocol "github.com/tliron/glsp/protocol_3_16"

	_ "github.com/tliron/commonlog/simple"
)

func StdioServer() {
	name := "Ziniki Deployer Language Server"
	version := "0.0.2"

	handler := protocol.Handler{
		Initialized:                     initialized,
		Shutdown:                        shutdown,
		WorkspaceExecuteCommand:         workspace.ExecuteCommand,
		WorkspaceDidChangeConfiguration: workspace.WorkspaceConfigChanged,
		TextDocumentCompletion:          document.OfferCompletions,
	}
	handler.Initialize = initializeCommand(name, version, &handler)

	server := server.NewServer(&handler, name, true)

	server.RunStdio()
}
