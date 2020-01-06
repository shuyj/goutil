package jsonutil

import (
	"strings"
	"testing"
)

func TestNewWriter(t *testing.T) {
	writer := NewJsonWriter()

	writer.AddField("appname", "www")
	writer.AddField("appver", "V1.0")
	writer.AddField("applt", "android")
	writer.AddField("type", "mp4")
	writer.AddField("ratio", "10")

	jsonstr, err := writer.BuildJson()

	t.Logf("BuildJson err=%v json=%s", err, jsonstr)
}

func TestNewWriterArray(t *testing.T) {
	writer := NewJsonWriter()

	writer.AddField("appname", []string{"news","test"})
	//writer.AddField("appver", "V1.0")
	writer.AddField("applt", []string{"android","ios"})
	//writer.AddField("type", "mp4")
	writer.AddField("ratio", "10")

	jsonstr, err := writer.BuildJson()

	t.Logf("BuildJson err=%v json=%s", err, jsonstr)
}



func TestAppverCheck(t *testing.T){

	vermap := map[string]string{
		"v1.0.9":"v1.0.1",
		"V11219.0910.01":"V11219.0911.01",
		"v7.02.2019.10.21.01.lcs":"v7.02.2019.10.22.18.lcs",
		"v1.0.1":"v1.0.1.0",
		"v1.0.0":"newsapp",
	}

	for ver1,ver2 := range vermap {
		t.Logf("Compare %s>%s = %v", ver1, ver2, strings.Compare(ver1, ver2))
	}
}

