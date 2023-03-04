package langserver

import (
	protocol "github.com/tliron/glsp/protocol_3_16"
)

var Handler protocol.Handler

func init() {
	// General Messages
	Handler.Initialize = Initialize
	Handler.Initialized = Initialized
	Handler.Shutdown = Shutdown
	Handler.LogTrace = LogTrace
	Handler.SetTrace = SetTrace

	// Text Document Synchronization
	Handler.TextDocumentDidOpen = TextDocumentDidOpen
	Handler.TextDocumentDidChange = TextDocumentDidChange
	Handler.TextDocumentDidSave = TextDocumentDidSave
	Handler.TextDocumentDidClose = TextDocumentDidClose

	// Language Features
	Handler.TextDocumentCompletion = TextDocumentCompletion
}
