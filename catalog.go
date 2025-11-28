package translate

type Catalog struct {
	Languages map[string]Translator
}

func NewEmptyCatalog() *Catalog {
	return &Catalog{Languages: make(map[string]Translator)}
}

func NewCatalog(defaultLang string, defaultTrans Translator) *Catalog {
	return &Catalog{Languages: map[string]Translator{
		defaultLang: defaultTrans,
	}}
}

func (c *Catalog) SupportedLanguages() []string {
	var languageCodes []string
	for lang := range c.Languages {
		languageCodes = append(languageCodes, lang)
	}
	return languageCodes
}

func (c *Catalog) AddLanguage(language string, translator Translator) {
	c.Languages[language] = translator
}

func (c *Catalog) GetTranslator(language string) Translator {
	return c.Languages[language]
}

func (c *Catalog) Quick(language string) QuickAdd {
	return QuickAdd{c: c, lang: language}
}

type QuickAdd struct {
	c    *Catalog
	lang string
}

func (q QuickAdd) Add(translator interface{}, err error) {
	if err == nil && translator != nil {
		if trans, ok := translator.(Translator); ok {
			q.c.Languages[q.lang] = trans
		}
	}
}
