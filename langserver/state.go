package langserver

import (
	"sync"

	"github.com/namngh/chrome-extension-manifest-ls/parser"
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
	urlpkg "github.com/tliron/kutil/url"
)

var documentStates sync.Map // protocol.DocumentUri to DocumentState

func getDocumentState(documentUri protocol.DocumentUri) *DocumentState {
	if documentState, ok := documentStates.Load(documentUri); ok {
		return documentState.(*DocumentState)
	} else {
		return nil
	}
}

func validateDocumentState(documentUri protocol.DocumentUri, notify glsp.NotifyFunc) *DocumentState {
	documentState, created := _getOrCreateDocumentState(documentUri)

	if created {
		go notify(protocol.ServerTextDocumentPublishDiagnostics, &protocol.PublishDiagnosticsParams{
			URI: documentUri,
		})
	}

	return documentState
}

func deleteDocumentState(documentUri protocol.DocumentUri) {
	documentStates.Delete(documentUri)
}

func resetDocumentStates() {
	documentStates.Range(func(protocolUri interface{}, documentState interface{}) bool {
		documentStates.Delete(protocolUri)
		urlpkg.DeregisterInternalURL(documentUriToInternalPath(protocolUri.(protocol.DocumentUri)))
		return true
	})
}

func _getOrCreateDocumentState(documentUri protocol.DocumentUri) (*DocumentState, bool) {
	if documentState, ok := documentStates.Load(documentUri); ok {
		return documentState.(*DocumentState), false
	} else {
		documentState := NewDocumentState(documentUri)
		if existing, loaded := documentStates.LoadOrStore(documentUri, documentState); loaded {
			return existing.(*DocumentState), false
		} else {
			return documentState, true
		}
	}
}

type DocumentState struct {
	Content string
	Context *parser.Context

	DocumentURI protocol.DocumentUri
}

func NewDocumentState(documentUri protocol.DocumentUri) *DocumentState {
	self := DocumentState{DocumentURI: documentUri}

	urlContext := urlpkg.NewContext()
	defer urlContext.Release()

	self.Content, _ = getDocument(documentUri)
	self.Context = parser.NewContext()

	return &self
}
