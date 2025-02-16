package helper

import "testing"

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

func TestJSONUnmarshal(t *testing.T) {
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
			if err := JSONUnmarshal(tt.args.data, tt.args.v); err != nil {
				t.Errorf("Unmarshal() error = %v", err)
			}
		})
	}
}

func TestJSONMarshal(t *testing.T) {
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
			if got := JSONMarshal(tt.args.data); got == "" {
				t.Errorf("JSONEncode() = %v, want not empty", got)
			}
		})
	}
}

func TestComparePassword(t *testing.T) {
	type args struct {
		hashedPassword string
		password       string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "success",
			args: args{
				hashedPassword: "$2y$12$eVBuBVaqnlGYbrvrEutaGeUUTuySf.j2B4b6OAg772Pm/opZ7nu0W",
				password:       "password",
			},
		},
		{
			name: "failed",
			args: args{
				hashedPassword: "$2y$12$eVBuBVaqnlGYbrvrEutaGeUUTuySf.j2B4b6OAg772Pm/opZ7nu0W",
				password:       "password1",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ComparePassword(tt.args.hashedPassword, tt.args.password); got != (tt.name == "success") {
				t.Errorf("ComparePassword() = %v, want %v", got, tt.name == "success")
			}
		})
	}
}

func TestArrayStringContains(t *testing.T) {
	type args struct {
		arr []string
		str string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "success",
			args: args{
				arr: []string{"john", "doe"},
				str: "john",
			},
		},
		{
			name: "failed",
			args: args{
				arr: []string{"john", "doe"},
				str: "john1",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ArrayStringContains(tt.args.arr, tt.args.str); got != (tt.name == "success") {
				t.Errorf("ArrayStringContains() = %v, want %v", got, tt.name == "success")
			}
		})
	}
}

func TestArrayInterfaceContains(t *testing.T) {
	type args struct {
		arr []interface{}
		str string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "success",
			args: args{
				arr: []interface{}{"john", "doe"},
				str: "john",
			},
		},
		{
			name: "failed",
			args: args{
				arr: []interface{}{"john", "doe"},
				str: "john1",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ArrayInterfaceContains(tt.args.arr, tt.args.str); got != (tt.name == "success") {
				t.Errorf("ArrayInterfaceContains() = %v, want %v", got, tt.name == "success")
			}
		})
	}
}

func TestIsUUID(t *testing.T) {
	type args struct {
		uuid string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "success",
			args: args{uuid: "f47ac10b-58cc-4372-a567-0e02b2c3d479"},
		},
		{
			name: "failed",
			args: args{uuid: "f47ac10b-58cc-4372-a567-0e02b2c3d4791"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsUUID(tt.args.uuid); got != (tt.name == "success") {
				t.Errorf("IsUUID() = %v, want %v", got, tt.name == "success")
			}
		})
	}
}

func TestEncryptAES256CBC(t *testing.T) {
	type args struct {
		data []byte
		key  []byte
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "success",
			args:    args{data: []byte("john doe"), key: []byte("12345678901234567890123456789012")},
			wantErr: false,
		},
		{
			name:    "failed",
			args:    args{data: []byte("john doe"), key: []byte("key")},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := EncryptAES256CBC(tt.args.data, tt.args.key); (err != nil) != tt.wantErr {
				t.Errorf("EncryptAES256CBC() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDecryptAES256CBC(t *testing.T) {
	type args struct {
		data []byte
		key  []byte
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "success",
			args:    args{data: []byte("john doe"), key: []byte("12345678901234567890123456789012")},
			wantErr: false,
		},
		{
			name:    "failed",
			args:    args{data: []byte("john doe"), key: []byte("key")},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			encryptedData, _ := EncryptAES256CBC([]byte("john doe"), []byte("12345678901234567890123456789012"))
			if _, err := DecryptAES256CBC(encryptedData, tt.args.key); (err != nil) != tt.wantErr {
				t.Errorf("DecryptAES256CBC() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestIsValidDeviceCategory(t *testing.T) {
	type args struct {
		category string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "success",
			args: args{category: "Web"},
		},
		{
			name: "failed",
			args: args{category: "desktop"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidDeviceCategory(tt.args.category); got != (tt.name == "success") {
				t.Errorf("IsValidDeviceCategory() = %v, want %v", got, tt.name == "success")
			}
		})
	}
}

func TestIsValidDeviceType(t *testing.T) {
	type args struct {
		deviceType string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "success",
			args: args{deviceType: "Desktop Browser"},
		},
		{
			name: "failed",
			args: args{deviceType: "desktop"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidDeviceType(tt.args.deviceType); got != (tt.name == "success") {
				t.Errorf("IsValidDeviceType() = %v, want %v", got, tt.name == "success")
			}
		})
	}
}
