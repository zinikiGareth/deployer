package workspace

import (
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func WorkspaceConfigChanged(context *glsp.Context, params *protocol.DidChangeConfigurationParams) error {
	return nil
}
