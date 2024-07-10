package jsonutil

func (j *JsonGo) GetItem(key string) (interface{}, error) {
	return j.GetItemByKeys(getkeys(key)...)
}

func (j *JsonGo) GetJson(key string) *JsonGo {
	return j.GetJsonByKeys(getkeys(key)...)
}

func (j *JsonGo) GetArray(key string) []interface{} {
	return j.GetArrayByKeys(getkeys(key)...)
}
func (j *JsonGo) GetMap(key string) map[string]interface{} {
	return j.GetMapByKeys(getkeys(key)...)
}

func (j *JsonGo) GetArrayInt(key string) []int {
	return j.GetArrayIntByKeys(getkeys(key)...)
}
func (j *JsonGo) GetStr(key string) string {
	return j.GetStrByKeys(getkeys(key)...)
}
func (j *JsonGo) GetNum(key string) (float64, error) {
	return j.GetNumByKeys(getkeys(key)...)
}
func (j *JsonGo) GetBool(key string) (bool, error) {
	return j.GetBoolByKeys(getkeys(key)...)
}
