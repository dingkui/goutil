package native

import (
	"fmt"
	"strings"
)

type NameSet struct {
	names map[string]bool
}

func NewNameSet() *NameSet {
	return &NameSet{
		names: make(map[string]bool),
	}
}

func (self *NameSet) UniqueName(name string) string {
	_, has := self.names[name]
	if has {
		root := name
		idx := 1

		for {
			name = fmt.Sprintf("%s.%03d", root, idx)
			_, has = self.names[name]

			if !has {
				break
			}
			idx += 1
		}
	}
	self.names[name] = true

	return name
}

type FileNameSet struct {
	root_last_suffix map[string]int
}

func NewFileNameSet() *FileNameSet {
	return &FileNameSet{root_last_suffix: make(map[string]int)}
}

const _max_file_name_root_length = 80

func _shorten_file_name_root(file_name_root string, max_length int) string {
	if len(file_name_root) > max_length {
		return file_name_root[0:max_length-1] + "_"
	}
	return file_name_root
}

func (self *FileNameSet) Unique_file_name(file_name string) string {
	file_name = strings.ReplaceAll(file_name, " ", "")
	file_name = strings.ReplaceAll(file_name, "#", "")
	root, ext := FileUtil.Splitext(file_name)
	lower_root := strings.ToLower(root)
	last_suffix, has := self.root_last_suffix[lower_root]
	if has {
		suffix := last_suffix + 1
		new_root := ""
		lower_new_root := ""
		for {
			suffix_str := fmt.Sprintf(".%03d", suffix)
			new_root = _shorten_file_name_root(root, _max_file_name_root_length-len(suffix_str)) + suffix_str
			lower_new_root = strings.ToLower(new_root)
			_, has := self.root_last_suffix[lower_new_root]
			if !has {
				break
			}
			suffix += 1
		}

		self.root_last_suffix[lower_root] = suffix
		self.root_last_suffix[lower_new_root] = 0
		return new_root + ext
	}

	self.root_last_suffix[lower_root] = 0
	return file_name
}
