package parser

import (
	"errors"

	"github.com/lithammer/fuzzysearch/fuzzy"
)

type PositionType int64

const (
	Key PositionType = iota
	Value
)

type PositionParser struct {
	PositionType PositionType
	Path         string
	Value        string
}

var keywords []string = []string{
	"manifest_version",
	"name",
	"version",
	"action",
	"default_icon",
	"16",
	"24",
	"32",
	"default_title",
	"default_popup",
	"default_locale",
	"description",
	"icons",
	"author",
	"automation",
	"background",
	"chrome_setting_overrides",
	"chrome_url_overrides",
	"commands",
	"content_scripts",
	"matches",
	"css",
	"js",
	"content_security_policy",
	"cross_origin_embedder_policy",
	"cross_origin_opener_policy",
	"declarative_net_request",
	"devtool_page",
	"event_rules",
	"export",
	"externally_connectable",
	"file_browser_handlers",
	"file_system_provider_capabilities",
	"homepage_url",
	"host_permissions",
	"import",
	"incognito",
	"input_components",
	"key",
	"minimum_chrome_version",
	"nacl_modules",
	"oauth2",
	"omnibox",
	"optional_host_permissions",
	"optional_permissions",
	"options_page",
	"options_ui",
	"permissions",
	"replacement_web_app",
	"requirements",
	"sandbox",
	"short_name",
	"storage",
	"tts_engine",
	"update_url",
	"version_name",
	"web_accessible_resources",
}

func (self *Context) Parse(content string, line int, character int) ([]string, error) {
	var result []string

	positionParser, err := getPositionParser(content, line, character)
	if err != nil {
		return nil, err
	}

	if positionParser.PositionType == Value {
		return result, nil
	}

	value := positionParser.Value

	if value[0] == ',' {
		result = append(result, "\"")
	} else if value[0] == '{' {
		result = append(result, "\"")
	} else if value[0] == '"' {
		trimValue := value[1:]

		result = fuzzy.Find(trimValue, keywords)
	}

	return result, nil
}

func getPositionParser(content string, line int, character int) (*PositionParser, error) {
	countNewLine := 0
	currentPosition := 0
	isFound := false
	for index, char := range content {
		if char == '\n' {
			countNewLine++
		}

		if countNewLine == line {
			isFound = true
			currentPosition = index + character
			break
		}
	}

	if !isFound {
		return nil, errors.New("Not found position")
	}

	positionParser := PositionParser{}
	isFoundPositionType := false
	isFoundComma := false

	for i := currentPosition; i >= 0; i-- {
		if !isFoundPositionType && content[i] == '[' {
			isFoundPositionType = true
			positionParser.PositionType = Value
			break
		}

		if !isFoundPositionType && (content[i] == '{' || content[i] == ']') {
			isFoundPositionType = true
			positionParser.PositionType = Key
			break
		}

		if content[i] == ',' {
			isFoundComma = true
		}

		if !isFoundPositionType && content[i] == ':' {
			isFoundPositionType = true

			if isFoundComma {
				positionParser.PositionType = Key
			} else {
				positionParser.PositionType = Value
			}

			break
		}
	}

	startValue := currentPosition
	for i := currentPosition; i >= 0; i-- {
		if content[i] == '[' || content[i] == '{' || content[i] == '}' || content[i] == ']' || content[i] == ',' || content[i] == '"' {
			startValue = i
			break
		}
	}
	positionParser.Value = content[startValue:currentPosition]

	return &positionParser, nil
}
