package helpers

import (
	"reflect"
	"testing"
)

func TestTryParseUint(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name    string
		args    args
		want    uint64
		wantErr bool
	}{
		{
			name:    "should return an uint64 when the text is parseable to uint64.",
			args:    args{text: "1"},
			want:    1,
			wantErr: false,
		},
		{
			name:    "should return an error when the text is empty.",
			args:    args{text: ""},
			want:    0,
			wantErr: true,
		},
		{
			name:    "should return an error when the text is not parseable to uint64",
			args:    args{text: "asd"},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := TryParseUint(tt.args.text)
			if (err != nil) != tt.wantErr {
				t.Errorf("TryParseUint() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("TryParseUint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProcessUint64Numbers(t *testing.T) {
	type args struct {
		input []string
	}
	tests := []struct {
		name        string
		args        args
		wantNumbers []uint64
	}{
		{
			name:        "should return a slice of valid uint64 from input",
			args:        args{input: []string{"1"}},
			wantNumbers: []uint64{1},
		},
		{
			name:        "should return an empty slice if input is empty",
			args:        args{input: []string{}},
			wantNumbers: []uint64{},
		},
		{
			name:        "should return an empty slice if all of the input item is unparseable to uint64.",
			args:        args{input: []string{"asd", "ss2", "2ss"}},
			wantNumbers: []uint64{},
		},
		{
			name:        "should return a slice with all of the parseable to uint64 from input and ignores the unparseable",
			args:        args{input: []string{"1", "ss2", "4", "2ss"}},
			wantNumbers: []uint64{1, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotNumbers := ProcessUint64Numbers(tt.args.input); !reflect.DeepEqual(gotNumbers, tt.wantNumbers) {
				t.Errorf("ProcessUint64Numbers() = %v, want %v", gotNumbers, tt.wantNumbers)
			}
		})
	}
}
