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

func (p *RecordManager) AddRecord(recordName string, maxRow int, colProp DataStruct.Dictionary, isSave bool) (*Record) {
	if !p.records.Has(recordName) {
		record := NewRecord(p.guid, recordName, maxRow, colProp, isSave)
		p.records.Set(recordName, record)
		return record
	}
	return nil
}
func (p *RecordManager) GetRecord(record string) (*Record) {
	return p.records.Get(record).(*Record)
}
