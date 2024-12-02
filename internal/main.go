//go:build js && wasm

package main

import (
	"fmt"
	"syscall/js"

	"github.com/google/yamlfmt"
	"github.com/google/yamlfmt/engine"
	"github.com/google/yamlfmt/formatters/basic"
)

// parseJsConfig parses the JavaScript config object into a basic.Config struct.
func parseJsConfig(jsConfig js.Value) (*basic.Config, error) {
	if jsConfig.Type() != js.TypeObject {
		return nil, fmt.Errorf("invalid config type: expected object")
	}

	config := &basic.Config{}
	properties := []string{
		"indent",
		"includeDocumentStart",
		"lineEnding",
		"maxLineLength",
		"retainLineBreaks",
		"retainLineBreaksSingle",
		"disallowAnchors",
		"scanFoldedAsLiteral",
		"indentlessArrays",
		"dropMergeTag",
		"padLineComments",
		"trimTrailingWhitespace",
		"eofNewline",
	}

	for _, property := range properties {
		value := jsConfig.Get(property)
		if value.IsUndefined() || value.IsNull() {
			continue
		}
		switch property {
		case "includeDocumentStart":
			config.IncludeDocumentStart = value.Bool()
		case "retainLineBreaks":
			config.RetainLineBreaks = value.Bool()
		case "retainLineBreaksSingle":
			config.RetainLineBreaksSingle = value.Bool()
		case "disallowAnchors":
			config.DisallowAnchors = value.Bool()
		case "scanFoldedAsLiteral":
			config.ScanFoldedAsLiteral = value.Bool()
		case "indentlessArrays":
			config.IndentlessArrays = value.Bool()
		case "dropMergeTag":
			config.DropMergeTag = value.Bool()
		case "trimTrailingWhitespace":
			config.TrimTrailingWhitespace = value.Bool()
		case "eofNewline":
			config.EOFNewline = value.Bool()
		case "indent":
			config.Indent = value.Int()
		case "maxLineLength":
			config.LineLength = value.Int()
		case "padLineComments":
			config.PadLineComments = value.Int()
		case "lineEnding":
			config.LineEnding = yamlfmt.LineBreakStyle(value.String())
		}
	}
	return config, nil
}

// newFormatter creates a new YAML formatter with the given config.
// This is taken from yamlfmt but its private
func newFormatter(config *basic.Config) yamlfmt.Formatter {
	return &basic.BasicFormatter{
		Config:       config,
		Features:     basic.ConfigureFeaturesFromConfig(config),
		YAMLFeatures: basic.ConfigureYAMLFeaturesFromConfig(config),
	}
}

// formatYAML formats the YAML input string based on the provided JavaScript config.
func formatYAML(buffer string, jsConfig js.Value) (string, error) {
	config, err := parseJsConfig(jsConfig)
	if err != nil {
		return "", err
	}
	formatter := newFormatter(config)

	eng := &engine.ConsecutiveEngine{
		LineSepCharacter: "\n",
		Formatter:        formatter,
		Quiet:            true,
		ContinueOnError:  true,
		OutputFormat:     "line",
	}

	output, err := eng.FormatContent([]byte(buffer))
	if err != nil {
		return "", err
	}
	return string(output), nil
}

// format is the JavaScript function exposed to format YAML.
func format(this js.Value, args []js.Value) interface{} {
	result := map[string]any{}
	if len(args) == 0 || len(args) > 2 {
		result["Error"] = "Invalid Type"
		return js.ValueOf(result)
	}
	if args[0].Type() != js.TypeString {
		result["Error"] = "First argument should be string"
		return js.ValueOf(result)
	}

	buffer := args[0].String()
	config := args[1]
	formattedYAML, err := formatYAML(buffer, config)
	if err != nil {
		result["Error"] = err.Error()
		return js.ValueOf(result)
	}
	result["Value"] = formattedYAML
	return js.ValueOf(result)
}

func main() {
	js.Global().Set("format", js.FuncOf(format))
	<-make(chan bool)
}
