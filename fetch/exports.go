package fetch

func Preprocessor(raw []byte, target interface{}, dougaType string) error {
	return preprocess(raw, target, dougaType)
}
