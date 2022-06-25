package locstrs

import (
	"bytes"
	"fmt"
	templatetext "text/template"
)

// Stores text string patterns by language codes
type LocText[ExecData any] struct {
	params *Parameters
	tmpls  map[string]*templatetext.Template
}

// Creates new LocText, with paramters and patters mapped by language codes
func NewLocText[ExecData any](params *Parameters, patterns map[string]string) (*LocText[ExecData], error) {
	if params == nil {

	}
	if len(patterns) == 0 {
		panic("patterns must contain at least one entry")
	}

	lt := &LocText[ExecData]{
		params: params,
		tmpls:  make(map[string]*templatetext.Template),
	}

	for languageCode, p := range patterns {
		var err error
		lt.tmpls[languageCode], err = templatetext.New("").Parse(p)
		if err != nil {
			return nil, fmt.Errorf("NewLocText `%v`: %w", languageCode, err)
		}
	}

	return lt, nil
}

// Works as NewLocText, but panic in case of errors
func MustLocText[ExecData any](params *Parameters, patterns map[string]string) *LocText[ExecData] {
	locText, err := NewLocText[ExecData](params, patterns)
	if err != nil {
		panic(fmt.Errorf("MustLocText: %w", err))
	}

	return locText
}

// Returns raw string passed in NewLocText by languageCode
func (lt *LocText[ExecData]) String(languageCode string) string {
	return getTemplateByLanguageCode(lt.params, lt.tmpls, languageCode).Tree.Root.String()
}

// Executes passed in NewLocText pattern by languageCode with passed data
func (lt *LocText[ExecData]) Execute(languageCode string, data ExecData) (string, error) {
	buf := bytes.NewBuffer(nil)
	err := getTemplateByLanguageCode(lt.params, lt.tmpls, languageCode).Execute(buf, data)
	if err != nil {
		return "", fmt.Errorf("Execute: %w", err)
	}

	return buf.String(), nil
}

// Works as Execute, but panics if error occures
func (lt *LocText[ExecData]) MustExecute(languageCode string, data ExecData) string {
	str, err := lt.Execute(languageCode, data)
	if err != nil {
		panic(fmt.Errorf("MustExecute: %w", err))
	}

	return str
}
