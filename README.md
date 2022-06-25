# Locstrs

Strings localisation library in Go.

Uses [text/template](https://pkg.go.dev/text/template) and [html/template](https://pkg.go.dev/html/template) for placing variable parts in strings.

Supports fallback settings (see examplae below)

Documentation: https://pkg.go.dev/github.com/nickname76/locstrs

*Please, **star** this repository, if you found this library useful.*

## Example usage

```Go
package main

import (
	"time"

	"github.com/nickname76/locstrs"
)

func main() {
	locstrsParams := &locstrs.Parameters{
		// Default fallback language
		FallbackLanguageCode: "en",
		SpecificFallbackLanguageCodes: map[string]string{
			// Fallback belarussian to russian
			"be": "ru",
		},
	}

	welcomePatterns := map[string]string{
		"en": `Welcome!`,
		"ru": `Добро пожаловать!`,
	}

	locWelcome := locstrs.MustLocText[*struct{}](locstrsParams, welcomePatterns)

	curTimePatterns := map[string]string{
		"en": `<p>Current time is {{.}}</p>`,
		"ru": `<p>Текущее время: {{.}}</p>`,
	}

	locCurTime := locstrs.MustLocHTML[time.Time](locstrsParams, curTimePatterns)

	println(locWelcome.String("be"))
	println(locCurTime.MustExecute("en", time.Now()))
}

```
