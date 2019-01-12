package parser

import (
	"reflect"
	"testing"

	"github.com/a-hilaly/cs/pkg/lang"
	"github.com/a-hilaly/cs/pkg/stats"
)

func TestParseFileBytes(t *testing.T) {
	type args struct {
		b  []byte
		ls *lang.Spec
	}
	tests := []struct {
		name    string
		args    args
		want    *stats.File
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseFileBytes(tt.args.b, tt.args.ls)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseFileBytes(%v, %v) error = %v, wantErr %v", tt.args.b, tt.args.ls, err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseFileBytes(%v, %v) = %v, want %v", tt.args.b, tt.args.ls, got, tt.want)
			}
		})
	}
}
