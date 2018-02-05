package Core

func GetInt(v interface{}) int {
	return v.(int)
}
func GetFloat(v interface{}) float32 {
	return v.(float32)
}
func GetString(v interface{}) string {
	return v.(string)
}
func GetObj(v interface{}) GUID {
	return v.(GUID)
}
