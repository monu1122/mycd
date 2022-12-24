package mycd

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"
)

func TestMyCd_PrintDir(t *testing.T) {
	type fields struct {
		CurrentDir string
		NewDir     string
	}
	tests := []struct {
		name   string
		fields fields
		result string
	}{
		{"testcase_id_1", fields{"/", "abc"}, "/abc"},
		{"testcase_id_2", fields{"/abc/def", "ghi"}, "/abc/def/ghi"},
		{"testcase_id_3", fields{"/abc/def", ".."}, "/abc"},
		{"testcase_id_4", fields{"/abc/def", "/abc"}, "/abc"},
		{"testcase_id_5", fields{"/abc/def", "/abc/klm"}, "/abc/klm"},
		{"testcase_id_6", fields{"/abc/def", "../.."}, "/"},
		{"testcase_id_7", fields{"/abc/def", "../../.."}, "/"},
		{"testcase_id_8", fields{"/abc/def", "."}, "/abc/def"},
		{"testcase_id_9", fields{"/abc/def", "..klm"}, "..klm: No such file or directory"},
		{"testcase_id_10", fields{"/abc/def", "//////"}, "/"},
		{"testcase_id_11", fields{"/abc/def", "......"}, "......: No such file or directory"},
		{"testcase_id_12", fields{"/abc/def", "../gh///../klm/."}, "/abc/klm"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bkp := os.Stdout
			re, wr, _ := os.Pipe()
			os.Stdout = wr

			cd := NewMyCd(tt.fields.CurrentDir, tt.fields.NewDir)
			cd.PrintDir()
			outC := make(chan string)
			go func() {
				var buf bytes.Buffer
				io.Copy(&buf, re)
				outC <- buf.String()
			}()
			wr.Close()
			os.Stdout = bkp
			res := <-outC
			res = res[:len(res)-1]
			if res != tt.result {
				t.Fail()
			}
		})
	}
}
