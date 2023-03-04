package langserver

import (
	protocol "github.com/tliron/glsp/protocol_3_16"
	urlpkg "github.com/tliron/kutil/url"
)

const INTERNAL_PATH_PREFIX = "language-server:"

func documentUriToInternalPath(uri protocol.DocumentUri) string {
	return INTERNAL_PATH_PREFIX + string(uri)
}

func setDocument(documentUri protocol.DocumentUri, content string) {
	urlpkg.UpdateInternalURL(documentUriToInternalPath(documentUri), content)
	deleteDocumentState(documentUri)
}

func getDocument(documentUri protocol.DocumentUri) (string, bool) {
	if url, err := urlpkg.NewValidInternalURL(documentUriToInternalPath(documentUri), nil); err == nil {
		return url.Content, true
	} else {
		return "", false
	}
}

func deleteDocument(documentUri protocol.DocumentUri) {
	urlpkg.DeregisterInternalURL(documentUriToInternalPath(documentUri))
	deleteDocumentState(documentUri)
}
