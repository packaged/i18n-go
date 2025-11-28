package translate

type Catalog struct {
	Languages map[string]Translator
}

func NewCatalog() *Catalog {
	return &Catalog{Languages: make(map[string]Translator)}
}

func (c *Catalog) AddLanguage(language string, translator Translator) {
	c.Languages[language] = translator
}

func (c *Catalog) GetTranslator(language string) Translator {
	return c.Languages[language]
}
