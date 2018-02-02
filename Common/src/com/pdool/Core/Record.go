package Core

import "container/list"

type Record struct {
	guid GUID
	recordName string
	maxRow int
	varTypes RecordRow
	colName RecordRow
	rows map[int]RecordRow
	callbacks *list.List
	oldVar *RecordRow
	newVar *RecordRow
}
func NewRecord(guid GUID,recordName string,maxRow int, varTypes RecordRow, colName RecordRow){
	r := new(Record)
	r.guid = guid
	r.recordName = recordName
	r.maxRow = maxRow
	r.varTypes = varTypes
	r.colName = colName
	r.rows = make(map[int]RecordRow)
	r.oldVar = new(RecordRow)
	r.newVar = new(RecordRow)
	r.callbacks = list.New()
}

func (r *Record)GetRowNum()int{
	return len(r.rows)
}
func (r *Record)GetColNum()int{
	return len(r.varTypes.values)
}
func (r *Record)GetMaxRowNum()int{
	return r.maxRow
}
func (r *Record)GetColumnType(column int)int{
	return r.varTypes.values[column].(int)
}
