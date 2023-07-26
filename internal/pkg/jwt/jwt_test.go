package jwt

import (
	"testing"
)

func TestToken(t *testing.T) {
	type args struct {
		kvs map[string]any
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "case1",
			args: args{
				kvs: map[string]any{
					"userid": "123456",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Token(tt.args.kvs)
			if (err != nil) != tt.wantErr {
				t.Errorf("Token() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// if got != tt.want {
			// 	t.Errorf("Token() = %v, want %v", got, tt.want)
			// }
			if err != nil {
				t.Errorf("Token() err: %v", err)
			}
			t.Logf("got:%s", got)
		})
	}
}

func TestVerifyToken(t *testing.T) {
	type args struct {
		tokenStr string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "case1",
			args: args{
				tokenStr: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTA0NDQ3NzgsInVzZXJpZCI6IjEyMzQ1NiJ9.AOkmQjZPC2qPX52K_iHbCsED6iHz330zXUBEUOUXXhg",
			},
			want:    "123456",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := VerifyToken(tt.args.tokenStr)
			t.Logf("got: %v", got)
			if (err != nil) != tt.wantErr {
				t.Errorf("VerifyToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got["userid"] != tt.want {
				t.Errorf("VerifyToken() = %v, want %v", got, tt.want)
			}
		})
	}
}
