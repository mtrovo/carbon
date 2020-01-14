package lang

import "fmt"

var (
	locales map[string]map[string]string
)

func init()  {
	locales = make(map[string]map[string]string)
	locales["en"] = en
}

// LoadLocaleText will load the translations from the locale json files according to the locale
func LoadLocaleText(l string) (map[string]string, error) {
	locale, ok := locales[l]
	if !ok {
		return nil, fmt.Errorf("locale data not available: %s", l)
	}

	return locale, nil
}
