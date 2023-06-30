package dhttp

func (ops *Options) AddParam(key string, value interface{}) {
	ops.addOption(key, value, optionParam)
}
func (ops *Options) GetParams() map[string]interface{} {
	paramsm := map[string]interface{}{}
	for _, option := range ops.options {
		if option.ot == optionParam {
			paramsm[option.key] = option.value
		}
	}
	return paramsm
}
func (ops *Options) GetParam(key string) interface{} {
	for _, option := range ops.options {
		if option.ot == optionParam && option.key == key {
			return option.value
		}
	}
	return ""
}
