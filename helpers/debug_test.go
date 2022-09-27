package helpers

import "testing"

func TestIsDebug(t *testing.T) {
	type EnvMap map[string]string
	tests := []struct {
		name    string
		withEnv EnvMap
		want    bool
	}{
		{
			name:    "should disable debug mode if there is no DEBUG envvars.",
			withEnv: EnvMap{},
			want:    false,
		},
		{
			name: "should disable debug mode if DEBUG envvars set to anything other than 1 or true.",
			withEnv: EnvMap{
				"DEBUG": "whatever",
			},
			want: false,
		},
		{
			name: "should enable debug mode if DEBUG envvars set to 1.",
			withEnv: EnvMap{
				"DEBUG": "1",
			},
			want: true,
		},
		{
			name: "should enable debug mode if DEBUG envvars set to true.",
			withEnv: EnvMap{
				"DEBUG": "true",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		for name, value := range tt.withEnv {
			t.Setenv(name, value)
		}
		t.Run(tt.name, func(t *testing.T) {
			if got := IsDebug(); got != tt.want {
				t.Errorf("IsDebug() = %v, want %v", got, tt.want)
			}
		})
	}
}