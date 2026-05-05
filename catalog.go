package i18n

import "golang.org/x/text/language"

type Catalog struct {
	Languages map[string]Translator
}

func NewEmptyCatalog() *Catalog {
	return &Catalog{Languages: make(map[string]Translator)}
}

func NewCatalog(defaultLang language.Tag, defaultTrans Translator) *Catalog {
	return &Catalog{Languages: map[string]Translator{
		defaultLang.String(): defaultTrans,
	}}
}

func (c *Catalog) SupportedLanguages() []string {
	var languageCodes []string
	for lang := range c.Languages {
		languageCodes = append(languageCodes, lang)
	}
	return languageCodes
}

func (c *Catalog) AddLanguage(language language.Tag, translator Translator) {
	c.Languages[language.String()] = translator
}

func (c *Catalog) GetTranslator(language language.Tag) Translator {
	return c.Languages[language.String()]
}

// TranslateList translates each key for the given language and joins the results as a localized list.
func (c *Catalog) TranslateList(language language.Tag, keys []string) string {
	return TranslatedList(language, c.GetTranslator(language), keys)
}

func (c *Catalog) Quick(language language.Tag) QuickAdd {
	return QuickAdd{c: c, lang: language}
}

type QuickAdd struct {
	c    *Catalog
	lang language.Tag
}

func (q QuickAdd) Add(translator interface{}, err error) {
	if err == nil && translator != nil {
		if trans, ok := translator.(Translator); ok {
			q.c.Languages[q.lang.String()] = trans
		}
	}
}
