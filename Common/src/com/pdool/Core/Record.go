package Core

import (
	"container/list"
	"com/pdool/DataStruct"
)

type Record struct {
	guid       GUID   //	对象id
	recordName string //	表名
	maxRow     int
	colProp    DataStruct.Dictionary //	列名 ：列类型
	isSave       bool        //	是否保存

	rows       *DataStruct.LinkedList
	callbacks  *list.List
	oldVar     *DataStruct.LinkedList
	newVar     *DataStruct.LinkedList
	opType     int
}

func NewRecord(guid GUID, recordName string, maxRow int, colProp DataStruct.Dictionary,isSave bool) (*Record) {
	r := new(Record)
	r.guid = guid
	r.recordName = recordName
	r.maxRow = maxRow
	r.colProp = colProp
	r.rows = new(DataStruct.LinkedList)
	r.oldVar = new(DataStruct.LinkedList)
	r.newVar = new(DataStruct.LinkedList)
	r.callbacks = list.New()
	r.isSave = isSave

	r.opType = RECORD_CREATE
	r.doCB(0, 0)
	return r
}

//	获得有多少行
func (r *Record) GetRowNum() int {
	return r.rows.Size()
}

//	获得有多少列
func (r *Record) GetColNum() int {
	return r.colProp.Size()
}

//	获得最多有多少行
func (r *Record) GetMaxRowNum() int {
	return r.maxRow
}

//	获得列的类型
func (r *Record) GetColumnType(column int) int {
	return r.colProp.Get(column).(int)
}

//	获得某一列的值
func (r *Record) GetValue(column int) interface{} {
	return r.rows.Get(column)
}

//	添加回调
func (r *Record) AddCallBack(cb IRecordHandler) {
	if r.callbacks == nil {
		r.callbacks = list.New()
	}
	r.callbacks.PushBack(cb)
}

//	设置第row行第col列的值
func (r *Record) SetValue(row int, col int, value interface{}) {
	if row >= 0 && row < r.maxRow {
		rowValue := r.rows.Get(row).(*DataStruct.LinkedList)
		old := rowValue.Get(col)
		if old != value {
			rowValue.Set(col, value)
			if r.callbacks != nil {
				r.oldVar.Clear()
				r.newVar.Clear()

				r.opType = RECORD_UPDATE
				r.newVar.Append(value)
				r.oldVar.Append(old)

				r.doCB(row, col)
			}
		}
	}
}

//	添加一行
func (r *Record) AddRow(rowValue *RecordRow) int {
	if r.rows.Size()+1 > r.maxRow {
		return -1
	}
	r.rows.Append(rowValue)
	return r.rows.Size()

}

//	删除一行
func (r *Record) DelRow(row int) bool {
	if r.rows.Size() < row {
		return false
	}

	i := r.rows.RemoveAt(row)
	if r.callbacks != nil {
		r.oldVar.Clear()
		r.newVar.Clear()

		r.opType = RECORD_DEL
		r.oldVar.Append(i)

		r.doCB(row, 0)
	}

	return true
}

//
func (r *Record) doCB(row int, col int) {
	for cb := r.callbacks.Front(); cb != nil; cb = cb.Next() {
		handler := cb.Value.(IRecordHandler)
		handler.RecordHandle(r.guid, r.recordName, r.opType, row, col, *r.oldVar, *r.newVar)
	}
}
