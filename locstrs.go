package locstrs

import (
	templatehtml "html/template"
	templatetext "text/template"
)

type Parameters struct {
	FallbackLanguageCode          string
	SpecificFallbackLanguageCodes map[string]string
}

func getTemplateByLanguageCode[TemplateType templatetext.Template | templatehtml.Template](params *Parameters, templs map[string]*TemplateType, languageCode string) *TemplateType {
	t := templs[languageCode]
	if t != nil {
		return t
	}

	if params.SpecificFallbackLanguageCodes != nil {
		lc, ok := params.SpecificFallbackLanguageCodes[languageCode]
		if ok {
			t = templs[lc]
			if t != nil {
				return t
			}
		}
	}

	t = templs[params.FallbackLanguageCode]
	if t != nil {
		return t
	}

	for _, v := range templs {
		t = v
		break
	}

	return t
}
