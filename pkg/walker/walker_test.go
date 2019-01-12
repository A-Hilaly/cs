package walker

import (
	"reflect"
	"testing"

	"github.com/a-hilaly/cs/pkg/lang"
	"github.com/a-hilaly/cs/pkg/parser"
	"github.com/a-hilaly/cs/pkg/stats"
)

func TestNew(t *testing.T) {
	type args struct {
		c *Config
	}
	tests := []struct {
		name string
		args args
		want *walker
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New(%v) = %v, want %v", tt.args.c, got, tt.want)
			}
		})
	}
}

func Test_walker_Walk(t *testing.T) {
	type fields struct {
		config *Config
		ld     lang.Detector
		parser parser.Parser
	}
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *stats.Total
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &walker{
				config: tt.fields.config,
				ld:     tt.fields.ld,
				parser: tt.fields.parser,
			}
			got, err := w.Walk(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("walker.Walk(%v) error = %v, wantErr %v", tt.args.path, err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("walker.Walk(%v) = %v, want %v", tt.args.path, got, tt.want)
			}
		})
	}
}

func Test_walker_simpleWalk(t *testing.T) {
	type fields struct {
		config *Config
		ld     lang.Detector
		parser parser.Parser
	}
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *stats.Total
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &walker{
				config: tt.fields.config,
				ld:     tt.fields.ld,
				parser: tt.fields.parser,
			}
			got, err := w.simpleWalk(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("walker.simpleWalk(%v) error = %v, wantErr %v", tt.args.path, err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("walker.simpleWalk(%v) = %v, want %v", tt.args.path, got, tt.want)
			}
		})
	}
}

func Test_walker_concurrentWalk(t *testing.T) {
	type fields struct {
		config *Config
		ld     lang.Detector
		parser parser.Parser
	}
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *stats.Total
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &walker{
				config: tt.fields.config,
				ld:     tt.fields.ld,
				parser: tt.fields.parser,
			}
			got, err := w.concurrentWalk(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("walker.concurrentWalk(%v) error = %v, wantErr %v", tt.args.path, err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("walker.concurrentWalk(%v) = %v, want %v", tt.args.path, got, tt.want)
			}
		})
	}
}
