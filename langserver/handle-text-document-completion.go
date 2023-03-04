package langserver

import (
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func TextDocumentCompletion(context *glsp.Context, params *protocol.CompletionParams) (interface{}, error) {
	var completionItems []protocol.CompletionItem
	if documentState := getDocumentState(params.TextDocument.URI); documentState != nil {
		result, error := documentState.Context.Parse(documentState.Content, int(params.Position.Line), int(params.Position.Character))
		if error != nil {
			return nil, error
		}

		for _, label := range result {
			completionItems = append(completionItems, protocol.CompletionItem{
				Label: label,
			})
			log.Infof("############ completion!")

		}
	}

	return completionItems, nil
}
