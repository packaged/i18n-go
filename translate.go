package translate

type Translator interface {
	Translate(key string) string
	TranslateWith(key string, args map[string]interface{}) string
}
