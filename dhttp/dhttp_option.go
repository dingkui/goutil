package dhttp

type optionType byte

const (
	optionParam    optionType = iota // URL parameter
	optionHeader                     // HTTP header
	optionArg                        // http调用时的参数
	optionUrlParam                   // URL 地址参数
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
	option := ops.find(key, ot)
	if option == nil {
		option = &Option{key, value, ot}
		ops.options = append(ops.options, option)
	} else {
		option.value = value
	}
}

func (ops *Options) find(key string, op optionType) *Option {
	for _, option := range ops.options {
		if option.key == key && option.ot == op {
			return option
		}
	}
	return nil
}

func (ops *Options) isOptionSet(key string, op optionType) (bool, interface{}, error) {
	return ops.find(key, op) != nil, nil, nil
}

func (ops *Options) deleteOption(strKey string, op optionType) {
	var outOption []*Option
	for _, option := range ops.options {
		if option.key != strKey || option.ot != op {
			outOption = append(outOption, option)
		}
	}
	ops.options = outOption
}
