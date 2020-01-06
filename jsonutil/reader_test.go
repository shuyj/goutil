package jsonutil

import (
	"testing"
)


func TestReaderImpl(t *testing.T) {
	jsonstr := "{\"protocal\": \"rtmp\",\"start_time\": 1571627175,\"end_time\": 1572235575,\"status\": \"0\"," +
		"\"clist\": [0,0,0,0,0,1,2,3,4,5,6], \"slist\":[\"one\",\"two\",\"three\"]}"

	reader,err := NewJsonReader(jsonstr)

	if err != nil {
		t.Errorf("err is %v", err)
		return
	}
	if val, ok := reader.String("protocal"); !ok {
		t.Error(`TestReaderImpl_String - expected to have field "String"`)
	}else{
		t.Logf("protocal is %v", val)
	}

	if val, ok := reader.Int("end_time"); !ok {
		t.Error(`TestReaderImpl_Int - expected not to have field "Int"`)
	}else{
		t.Logf("end_time is %v", val)
	}

	if val, ok := reader.Float64("delay"); !ok {
		t.Error(`TestReaderImpl_Float64 - expected not to have field "Float64"`)
	}else{
		t.Logf("delay is %v", val)
	}
}


func TestReader_Array(t *testing.T) {
	jsonstr := "{\"protocal\": \"rtmp\",\"start_time\": 1571627175,\"end_time\": 1572235575,\"status\": \"0\"," +
		"\"clist\": [0,0,0,0,0,1,2,3,4,5,6], \"slist\":[\"one\",\"two\",\"three\"]}"

	reader,err := NewJsonReader(jsonstr)

	if err != nil {
		t.Errorf("err is %v", err)
		return
	}

	if val, ok := reader.IntArray("clist"); !ok {
		t.Error(`TestReaderImpl_CList - expected not to have field "Int"`)
	}else{
		t.Logf("clist is %v", val)
	}

	if val, ok := reader.StringArray("slist"); !ok {
		t.Error(`TestReaderImpl_SList - expected not to have field "String"`)
	}else{
		t.Logf("slist is %v", val)
	}
}
