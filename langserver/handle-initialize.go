package langserver

import (
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
	"github.com/tliron/kutil/version"
)

var clientCapabilities *protocol.ClientCapabilities

func Initialize(context *glsp.Context, params *protocol.InitializeParams) (interface{}, error) {
	clientCapabilities = &params.Capabilities

	if params.Trace != nil {
		protocol.SetTraceValue(*params.Trace)
	}

	serverCapabilities := Handler.CreateServerCapabilities()
	serverCapabilities.TextDocumentSync = protocol.TextDocumentSyncKindIncremental
	serverCapabilities.CompletionProvider = &protocol.CompletionOptions{}

	return &protocol.InitializeResult{
		Capabilities: serverCapabilities,
		ServerInfo: &protocol.InitializeResultServerInfo{
			Name:    toolName,
			Version: &version.GitVersion,
		},
	}, nil
}

func Initialized(context *glsp.Context, params *protocol.InitializedParams) error {
	return nil
}
