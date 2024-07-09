package httputil

import "errors"

type optionType string

const (
	optionParam    optionType = "HTTPParameter" // URL parameter
	optionHTTP     optionType = "HTTPHeader"    // HTTP header
	optionArg      optionType = "FuncArgument"  // Function argument
	optionUrlParam optionType = "UrlParameter"  // URL parameter
)

type Option struct {
	key   string
	value interface{}
	ot    optionType
}

type Options struct {
	options []*Option
}

func (ops *Options) addOption(key string, value interface{}, ot optionType) {
	option := ops.Find(key)
	if option == nil {
		option = &Option{key, value, ot}
		ops.options = append(ops.options, option)
	} else {
		option.value = value
	}
}

func (ops *Options) AddHeader(key string, value string) {
	ops.addOption(key, value, optionHTTP)
}
func (ops *Options) AddParam(key string, value interface{}) {
	ops.addOption(key, value, optionParam)
}
func (ops *Options) AddUrlParam(key string, value interface{}) {
	ops.addOption(key, value, optionUrlParam)
}
func (ops *Options) AddArg(key string, value interface{}) {
	ops.addOption(key, value, optionArg)
}

func (ops *Options) FillHeaders(headers map[string]string) error {
	for _, option := range ops.options {
		if option.ot == optionHTTP {
			headers[option.key] = option.value.(string)
		}
	}
	return nil
}

func (ops *Options) GetRawParams() (map[string]interface{}, error) {
	paramsm := map[string]interface{}{}
	for _, option := range ops.options {
		if option.ot == optionParam {
			paramsm[option.key] = option.value
		}
	}
	return paramsm, nil
}
func (ops *Options) GetUrlParams() (map[string]interface{}, error) {
	paramsm := map[string]interface{}{}
	for _, option := range ops.options {
		if option.ot == optionUrlParam {
			paramsm[option.key] = option.value
		}
	}
	return paramsm, nil
}
func (ops *Options) GetArgs() (map[string]interface{}, error) {
	paramsm := map[string]interface{}{}
	for _, option := range ops.options {
		if option.ot == optionArg {
			paramsm[option.key] = option.value
		}
	}
	return paramsm, nil
}
func (ops *Options) GetArg(key string) interface{} {
	for _, option := range ops.options {
		if option.ot == optionArg && option.key == key {
			return option.value
		}
	}
	return nil
}

func (ops *Options) Find(key string) *Option {
	for _, option := range ops.options {
		if option.key == key {
			return option
		}
	}
	return nil
}

func (ops *Options) FindOption(key string, defaultVal interface{}) (interface{}, error) {
	find := ops.Find(key)
	if find != nil {
		return find, nil
	}
	if defaultVal == nil {
		return nil, errors.New("not found")
	}
	return defaultVal, nil
}

func (ops *Options) GetOption(param string, defaultVal interface{}) (interface{}, error) {
	option := ops.Find(param)
	if option != nil {
		return option.value, nil
	}
	return defaultVal, nil
}

func (ops *Options) IsOptionSet(key string) (bool, interface{}, error) {
	return ops.Find(key) != nil, nil, nil
}

func (ops *Options) DeleteOption(strKey string) []Option {
	var outOption []Option
	for _, option := range ops.options {
		if option.key != strKey {
			outOption = append(outOption, *option)
		}
	}
	return outOption
}
