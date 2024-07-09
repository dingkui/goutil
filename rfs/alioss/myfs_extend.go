package oss

func HandleOptions(headers map[string]string, options []Option) error {
	return handleOptions(headers, options)
}
