package helpers

import "testing"

func TestMaxLenOnSlice(t *testing.T) {
	type args struct {
		slice *[]string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "should return the maximum length from slice of strings",
			args: args{
				slice: &[]string{
					"asdasd",
					"asd",
					"asdasd",
					"asdasdasd",
				},
			},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxLenOnSlice(tt.args.slice); got != tt.want {
				t.Errorf("MaxLenOnSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringInSlice(t *testing.T) {
	type args struct {
		text  string
		slice []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "should return true if there is a string from text argument within the slice ",
			args: args{
				text: "asd",
				slice: []string{
					"bad",
					"das",
					"asd",
				},
			},
			want: true,
		},
		{
			name: "should return false if slice is empty.",
			args: args{
				text:  "word",
				slice: []string{},
			},
			want: false,
		},
		{
			name: "should return false if there is no string from text argument within the slice.",
			args: args{
				text: "word",
				slice: []string{
					"bad",
					"das",
					"asd",
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringInSlice(tt.args.text, tt.args.slice); got != tt.want {
				t.Errorf("StringInSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}
