package models

type Language string

const (
	English Language = "English"
	Persian Language = "Persian"
	Unknown Language = "unknown"
)

type LanguageCode string

const (
	EnglishCode LanguageCode = "en"
	PersianCode LanguageCode = "fa"
	UnknownCode LanguageCode = "unknown"
)

func LanguageFromStringCode(languageCodeRaw string) LanguageCode {
	switch LanguageCode(languageCodeRaw) {
	case EnglishCode:
		return EnglishCode
	case PersianCode:
		return PersianCode
	default:
		return UnknownCode
	}
}

func LanguageCodeFromLanguage(language Language) LanguageCode {
	switch language {
	case English:
		return EnglishCode
	case Persian:
		return PersianCode
	default:
		return UnknownCode
	}
}
