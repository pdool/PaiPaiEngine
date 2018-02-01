package Data

import "log"

type Data struct {
	value interface{}
}

func NewData(value interface{}) *Data {
	d := new(Data)
	d.value = value
	return d
}

func (data *Data) SetData(value interface{}) {
	data.value = value
}

/**
* 从other拷贝数据，类型可以不同
* @param other
*/
func (data *Data) GetInt() int {
	value, ok := data.value.(int)
	if !ok {
		log.Fatal("wrong type!!")
		return 0
	}
	return value
}
func (data *Data) GetFloat() float32 {
	value, ok := data.value.(float32)
	if !ok {
		log.Fatal("wrong type!!")
		return 0
	}
	return value
}
func (data *Data) GetString() string {
	value, ok := data.value.(string)
	if !ok {
		log.Fatal("wrong type!!")
		return ""
	}
	return value
}
func (data *Data) GetObject() interface{} {
	value, ok := data.value.(Data)
	if !ok {
		log.Fatal("wrong type!!")
		return ""
	}
	return value
}
