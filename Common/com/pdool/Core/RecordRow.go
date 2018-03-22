package Core

type RecordRow struct {
	values map[int]interface{}
}

func NewDataRow() *RecordRow {
	v := &RecordRow{values: make(map[int]interface{})}
	return v
}
func (row *RecordRow) SetValue(index int, value interface{}) {
	row.values[index] = value
}
func (row *RecordRow) GetInt(index int) int {
	i := row.values[index]
	return i.(int)
}
func (row *RecordRow) GetFloat(index int) float32 {
	i := row.values[index]
	return i.(float32)
}
func (row *RecordRow) GetString(index int) string {
	i := row.values[index]
	return i.(string)
}
func (row *RecordRow) GetObj(index int) GUID {
	i := row.values[index]
	return i.(GUID)
}
func  (row *RecordRow) GetValue(index int) interface{} {
	i := row.values[index]
	return i
}