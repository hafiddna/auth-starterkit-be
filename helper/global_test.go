package helper

import "testing"

func TestUnmarshal(t *testing.T) {
	type args struct {
		data []byte
		v    interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "success",
			args: args{
				data: []byte(`{"name":"john doe"}`),
				v:    &struct{ Name string }{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Unmarshal(tt.args.data, tt.args.v); err != nil {
				t.Errorf("Unmarshal() error = %v", err)
			}
		})
	}
}

func TestRandomString(t *testing.T) {
	type args struct {
		length int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "success",
			args: args{length: 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RandomString(tt.args.length); len(got) != tt.args.length {
				t.Errorf("RandomString() = %v, want %v", got, tt.args.length)
			}
		})
	}
}

func TestRandomInt(t *testing.T) {
	type args struct {
		min int
		max int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "success",
			args: args{min: 0, max: 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RandomInt(tt.args.min, tt.args.max); got < tt.args.min || got > tt.args.max {
				t.Errorf("RandomInt() = %v, want between %v and %v", got, tt.args.min, tt.args.max)
			}
		})
	}
}

func TestJSONEncode(t *testing.T) {
	type args struct {
		data interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "success",
			args: args{data: struct{ Name string }{Name: "john doe"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := JSONEncode(tt.args.data); got == "" {
				t.Errorf("JSONEncode() = %v, want not empty", got)
			}
		})
	}
}
