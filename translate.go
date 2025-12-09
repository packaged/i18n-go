package i18n

type Translator interface {
	Translate(key string) string
	TranslateWith(key string, args map[string]interface{}) string
	TranslatePlural(key string, number int64, args map[string]interface{}) string
}
