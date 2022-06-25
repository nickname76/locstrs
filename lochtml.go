package locstrs

import (
	"bytes"
	"fmt"
	templatehtml "html/template"
)

// Stores html string patterns by language codes
type LocHTML[ExecData any] struct {
	params *Parameters
	tmpls  map[string]*templatehtml.Template
}

// Creates new LocHTML, with paramters and patters mapped by language codes
func NewLocHTML[ExecData any](params *Parameters, patterns map[string]string) (*LocHTML[ExecData], error) {
	if len(patterns) == 0 {
		panic("patterns must contain at least one entry")
	}

	lh := &LocHTML[ExecData]{
		params: params,
		tmpls:  make(map[string]*templatehtml.Template),
	}

	for languageCode, p := range patterns {
		var err error
		lh.tmpls[languageCode], err = templatehtml.New("").Parse(p)
		if err != nil {
			return nil, fmt.Errorf("NewLocaleHTML `%v`: %w", languageCode, err)
		}
	}

	return lh, nil
}

// Works as NewLocHTML, but panic in case of errors
func MustLocHTML[ExecData any](params *Parameters, patterns map[string]string) *LocHTML[ExecData] {
	locHTML, err := NewLocHTML[ExecData](params, patterns)
	if err != nil {
		panic(fmt.Errorf("MustLocHTML: %w", err))
	}

	return locHTML
}

// Returns raw string passed in NewLocHTML by languageCode
func (lh *LocHTML[ExecData]) String(languageCode string) string {
	return getTemplateByLanguageCode(lh.params, lh.tmpls, languageCode).Tree.Root.String()
}

// Executes passed in NewLocHTML pattern by languageCode with passed data
func (lh *LocHTML[ExecData]) Execute(languageCode string, data ExecData) (string, error) {
	buf := bytes.NewBuffer(nil)
	err := getTemplateByLanguageCode(lh.params, lh.tmpls, languageCode).Execute(buf, data)
	if err != nil {
		return "", fmt.Errorf("Execute: %w", err)
	}

	return buf.String(), nil
}

// Works as Execute, but panics if error occures
func (lh *LocHTML[ExecData]) MustExecute(languageCode string, data ExecData) string {
	str, err := lh.Execute(languageCode, data)
	if err != nil {
		panic(fmt.Errorf("MustExecute: %w", err))
	}

	return str
}
