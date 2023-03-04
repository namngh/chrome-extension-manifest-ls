package langserver

import (
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func TextDocumentDidOpen(context *glsp.Context, params *protocol.DidOpenTextDocumentParams) error {
	go Configure(context, &params.TextDocument.URI)
	protocol.Trace(context, protocol.MessageTypeInfo, "hi!!!")

	setDocument(params.TextDocument.URI, params.TextDocument.Text)
	go validateDocumentState(params.TextDocument.URI, context.Notify)
	return nil
}
