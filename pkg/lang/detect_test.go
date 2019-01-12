package lang

import (
	"reflect"
	"testing"
)

func Test_detector_Detect(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		d       *detector
		args    args
		want    *Spec
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &detector{}
			got, err := d.Detect(tt.args.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("detector.Detect(%v) error = %v, wantErr %v", tt.args.filename, err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("detector.Detect(%v) = %v, want %v", tt.args.filename, got, tt.want)
			}
		})
	}
}
