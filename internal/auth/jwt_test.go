package auth

import (
	"fmt"
	"testing"
)

func TestJWT(t *testing.T) {
	type args struct {
		email  string
		name   string
		getenv func(string) string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test generateJWT",
			args: args{
				email: "me@guerra.io",
				name:  "guerra",
				getenv: func(string) string {
					return "secret"
				},
			},
			wantErr: false,
		},
		{
			name: "Test no mail",
			args: args{
				email: "",
				name:  "guerra",
				getenv: func(string) string {
					return "secret"
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotTokenString, err := GenerateJWT(tt.args.email, tt.args.name, tt.args.getenv)
			if (err != nil) != tt.wantErr {
				t.Errorf("generateJWT() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			fmt.Println(gotTokenString)

			claims, err := validateToken(gotTokenString, tt.args.getenv)
			if tt.wantErr {
				if err == nil {
					t.Errorf("validateToken() error = %v, wantErr %v", err, tt.wantErr)
				}

				return
			}

			if claims.Name != tt.args.name {
				t.Errorf("validateToken() claims.Name = %v, want %v", claims.Name, tt.args.name)
			}

			if claims.Email != tt.args.email {
				t.Errorf("validateToken() claims.Email = %v, want %v", claims.Email, tt.args.email)
			}
		})
	}
}
