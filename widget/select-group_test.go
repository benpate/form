package widget

import (
	"encoding/json"
	"html"
	"regexp"
	"strings"
	"testing"

	"github.com/benpate/form"
	"github.com/benpate/rosetta/mapof"
	"github.com/benpate/rosetta/schema"
	"github.com/stretchr/testify/require"
)

// TestSelectGroup_ChildValueStaysInsideItsString confirms that a stored child value holding quotes
// and backslashes cannot end the hyperscript string it is written into
func TestSelectGroup_ChildValueStaysInsideItsString(t *testing.T) {

	UseAll()

	s := schema.New(schema.Object{
		Properties: schema.ElementMap{
			"color": schema.String{},
			"other": schema.String{},
		},
	})

	f := form.New(s, form.Element{Type: "select-group", Path: "color", Options: mapof.Template{"provider": "test", "children": "other"}})

	for _, childValue := range []string{`plain`, `x', y:'evil`, `ends with a backslash\`, `\'); call alert(1); log('`} {

		value := mapof.Any{"color": "ONE", "other": childValue}

		result, err := f.Editor(&value, testLookupProvider{})
		require.NoError(t, err)

		script := selectGroupScript(t, result)

		// The value literal decodes back to the stored value, and nothing follows it but the closing paren
		start := strings.Index(script, "value:'")
		require.NotEqual(t, -1, start, "script: %s", script)

		decoded, rest := readHyperscriptString(t, script[start+len("value:"):])
		require.Equal(t, childValue, decoded, "script: %s", script)
		require.Equal(t, ")", rest, "script: %s", script)
	}
}

// TestSelectGroup_OptionsAreValidJSON confirms that the options literal still decodes to the
// lookup codes, so the browser can parse it
func TestSelectGroup_OptionsAreValidJSON(t *testing.T) {

	UseAll()

	s := schema.New(schema.Object{
		Properties: schema.ElementMap{
			"color": schema.String{},
			"other": schema.String{},
		},
	})

	f := form.New(s, form.Element{Type: "select-group", Path: "color", Options: mapof.Template{"provider": "test", "children": "other"}})
	value := mapof.Any{"color": "ONE", "other": "TWO"}

	result, err := f.Editor(&value, testLookupProvider{})
	require.NoError(t, err)

	script := selectGroupScript(t, result)
	start := strings.Index(script, "options:'")
	require.NotEqual(t, -1, start, "script: %s", script)

	decoded, _ := readHyperscriptString(t, script[start+len("options:"):])

	var lookupCodes []form.LookupCode
	require.NoError(t, json.Unmarshal([]byte(decoded), &lookupCodes))
	require.Len(t, lookupCodes, 5)
}

// selectGroupScript returns the decoded data-script attribute that installs the SelectGroup behavior
func selectGroupScript(t *testing.T, markup string) string {

	match := regexp.MustCompile(`data-script="(install SelectGroup[^"]*)"`).FindStringSubmatch(markup)
	require.Len(t, match, 2, "markup: %s", markup)

	return html.UnescapeString(match[1])
}

// readHyperscriptString decodes the single-quoted literal at the start of source the way the
// hyperscript tokenizer does, and returns it with whatever follows the closing quote
func readHyperscriptString(t *testing.T, source string) (string, string) {

	require.True(t, strings.HasPrefix(source, "'"), "source: %s", source)

	var result strings.Builder

	for index := 1; index < len(source); index++ {

		switch source[index] {

		case '\\':
			index++
			require.Less(t, index, len(source), "source ends mid-escape: %s", source)
			result.WriteByte(source[index])

		case '\'':
			return result.String(), source[index+1:]

		default:
			result.WriteByte(source[index])
		}
	}

	require.Fail(t, "unterminated string", "source: %s", source)
	return "", ""
}
