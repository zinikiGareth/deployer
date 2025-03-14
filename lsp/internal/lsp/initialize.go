package lsp

import (
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func initializeCommand(name, version string, handler *protocol.Handler) protocol.InitializeFunc {
	return func(context *glsp.Context, params *protocol.InitializeParams) (any, error) {
		capabilities := handler.CreateServerCapabilities()

		capabilities.CompletionProvider = &protocol.CompletionOptions{}

		return protocol.InitializeResult{
			Capabilities: capabilities,
			ServerInfo: &protocol.InitializeResultServerInfo{
				Name:    name,
				Version: &version,
			},
		}, nil
	}
}

func initialized(context *glsp.Context, params *protocol.InitializedParams) error {
	return nil
}
