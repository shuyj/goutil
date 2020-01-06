package common

import (
	"testing"
	"time"
)

type TestConfig struct {
	Appname string `bson:"appname" json:"appname"`
	Appver  string `bson:"appver" json:"appver"`
	Applt   string `bson:"applt" json:"applt"`
	Gray    string `bson:"gray" json:"gray"`
	Body    struct{
		Codec string `bson:"codec" json:"codec"`
		Types []string `bson:"types" json:"types"`
		Widths []int `bson:"widths" json:"widths"`
		NewLive struct{
			Enable int `bson:"enable" json:"enable"`
		} `bson:"new_live" json:"new_live"`
	} `bson:"body" json:"body"`
	CreateAt time.Time `bson:"create_at" json:"create_at"`
	ModifyAt time.Time `bson:"modify_at" json:"modify_at"`
	Versions []string `bson:"versions" json:"versions"`
}

func TestMapToStructJsonTagName(t *testing.T) {

	mobj := map[string]interface{}{"appname":"test1", "gray":"ABC", "body":map[string]interface{}{"codec":"h264", "types":[]string{"mp4","flv","hls"}, "widths":[]int{1,2,3,4}, "new_live":map[string]interface{}{"enable":1}}, "versions":[]string{"111","222","333"}}

	target := &TestConfig{}

	err := MapToStructJsonTagName(mobj, target)

	t.Logf("err = %v result  = %+v", err, target)

}

func BenchmarkMapToStructJsonTagName(b *testing.B) {
	mobj := map[string]interface{}{"appname":"test1", "gray":"ABC", "body":map[string]interface{}{"codec":"h264", "types":[]string{"mp4","flv","hls"}, "widths":[]int{1,2,3,4}, "new_live":map[string]interface{}{"enable":1}}, "versions":[]string{"111","222","333"}}

	target := &TestConfig{}

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next(){
			err := MapToStructJsonTagName(mobj, target)
			if err != nil {
				b.Logf("result  = %v", err)
			}
		}
	})


}

func TestMapToStructJsonTagNameWithJson(t *testing.T) {
	mobj := map[string]interface{}{"appname":"test1", "gray":"ABC", "body":map[string]interface{}{"codec":"h264", "types":[]string{"mp4","flv","hls"}, "widths":[]int{1,2,3,4}, "new_live":map[string]interface{}{"enable":1}}, "versions":[]string{"111","222","333"}}

	target := &TestConfig{}

	err := MapToStructJsonTagNameWithJson(mobj, target)

	t.Logf("err = %v result  = %+v", err, target)
}

func BenchmarkMapToStructJsonTagNameWithJson(b *testing.B) {
	mobj := map[string]interface{}{"appname":"test1", "gray":"ABC", "body":map[string]interface{}{"codec":"h264", "types":[]string{"mp4","flv","hls"}, "widths":[]int{1,2,3,4}, "new_live":map[string]interface{}{"enable":1}}, "versions":[]string{"111","222","333"}}

	target := &TestConfig{}

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next(){
			err := MapToStructJsonTagNameWithJson(mobj, target)
			if err != nil {
				b.Logf("result  = %v", err)
			}
		}
	})

}


func TestStructJsonTagNameToMap(t *testing.T) {
	mobj := &TestConfig{Appname:"test1", Applt:"android"}
	mobj.Body.Codec = "h264"

	mmap := StructJsonTagNameToMap(*mobj)

	t.Logf("map result = %v", mmap)

}

func BenchmarkStructJsonTagNameToMap(b *testing.B) {
	mobj := &TestConfig{Appname:"test1", Applt:"android"}
	mobj.Body.Codec = "h264"

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			mmap := StructJsonTagNameToMap(*mobj)
			if mmap == nil {
				b.Logf("result is nil")
			}
		}
	})
}

func TestStructJsonTagNameToMapWithJson(t *testing.T) {
	mobj := &TestConfig{Appname:"test1", Applt:"android"}
	mobj.Body.Codec = "h264"

	mmap := StructJsonTagNameToMapWithJson(mobj)

	t.Logf("map result = %v", mmap)
}

func BenchmarkStructJsonTagNameToMapWithJson(b *testing.B) {
	mobj := &TestConfig{Appname:"test1", Applt:"android"}
	mobj.Body.Codec = "h264"

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			mmap := StructJsonTagNameToMapWithJson(mobj)
			if mmap == nil {
				b.Logf("result is nil")
			}
		}
	})
}




