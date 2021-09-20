package auth

import (
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func TestHashPassword(t *testing.T) {
	goodPassword := "myP@ssw0rD"

	type args struct {
		password string
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "valid input valid output", args: args{password: goodPassword}, wantErr: false},
		{name: "empty input leads to error", args: args{password: ""}, wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := HashPassword(tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("HashPassword() error = %v, wantErr %v", err, tt.wantErr)
				return
			} else if !tt.wantErr && bcrypt.CompareHashAndPassword([]byte(got), []byte(goodPassword)) != nil {
				t.Errorf("the given hash should be compared successfully with password %s", tt.args.password)
			}
		})
	}
}

func TestCheckPassword(t *testing.T) {
	goodPassword := "myP@ssw0rD"
	goodPasswordHash, _ := bcrypt.GenerateFromPassword([]byte(goodPassword), HashCost)
	type args struct {
		password string
		hash     string
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "wrong password against hash", args: args{password: "wrong", hash: string(goodPasswordHash)}, want: false},
		{name: "good password should be successfully verified against hash", args: args{password: goodPassword, hash: string(goodPasswordHash)}, want: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckPassword(tt.args.password, tt.args.hash); got != tt.want {
				t.Errorf("CheckPassword() = %v, want %v", got, tt.want)
			}
		})
	}
}
