package Core

import (
	"com/pdool/DataStruct"
)

type RecordManager struct {
	guid    GUID
	records DataStruct.Dictionary
}

func NewRecordManager(guid GUID) *RecordManager {
	rMgr := new(RecordManager)
	rMgr.guid = guid
	rMgr.records = DataStruct.Dictionary{}
	return rMgr
}

func (p *RecordManager) AddRecord(recordName string, maxRow int, varTypes RecordRow, colName RecordRow) (*Record) {
	if !p.records.Has(recordName) {
		record := NewRecord(p.guid, recordName, maxRow, varTypes, colName)
		p.records.Set(recordName, record)
		return record
	}
	return nil
}
func (p *RecordManager) GetRecord(record string) (*Record) {
	return p.records.Get(record).(*Record)
}
func (p *RecordManager) GetRecordInt(recordName string, row int, col int) int {
	rec := p.records.Get(recordName)
	if rec != nil {
		return rec.(*Record).GetValue(col).(int)
	}
	return 0
}

func (rMgr *RecordManager) AddCallBack(recordName string, cb IRecordHandler) bool {
	rec := rMgr.records.Get(recordName).(*Record)
	if rec != nil {
		rec.AddCallBack(cb)
		return true
	}
	return false
}
