package langserver

import (
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func TextDocumentDidChange(context *glsp.Context, params *protocol.DidChangeTextDocumentParams) error {
	if content, ok := getDocument(params.TextDocument.URI); ok {
		for _, change := range params.ContentChanges {
			if change_, ok := change.(protocol.TextDocumentContentChangeEvent); ok {
				startIndex, endIndex := change_.Range.IndexesIn(content)
				content = content[:startIndex] + change_.Text + content[endIndex:]
				log.Debugf("content:\n%s", content)
			} else if change_, ok := change.(protocol.TextDocumentContentChangeEventWhole); ok {
				content = change_.Text
			}
		}
		setDocument(params.TextDocument.URI, content)
		go validateDocumentState(params.TextDocument.URI, context.Notify)
	}
	return nil
}
