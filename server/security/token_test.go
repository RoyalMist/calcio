package security

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func TestHashPassword(t *testing.T) {
	t.Parallel()
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
				t.Errorf("HashPassword() error = %v, wantSignErr %v", err, tt.wantErr)
				return
			} else if !tt.wantErr && bcrypt.CompareHashAndPassword([]byte(got), []byte(goodPassword)) != nil {
				t.Errorf("the given hash should be compared successfully with password %s", tt.args.password)
			}
		})
	}
}

func TestCheckPassword(t *testing.T) {
	t.Parallel()
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

func TestSignVerifyToken(t *testing.T) {
	t.Parallel()
	type args struct {
		claims   Claims
		validity time.Duration
	}

	tests := []struct {
		name          string
		args          args
		wait          time.Duration
		wantSignErr   bool
		wantVerifyErr bool
	}{
		{
			name: "a valid token should be issued and validated with one claim",
			args: args{
				claims: Claims{
					UserId: uuid.New().String(),
				},
				validity: 20 * time.Minute,
			},
			wait:          0,
			wantSignErr:   false,
			wantVerifyErr: false,
		},
		{
			name: "a valid token should be issued and validated with multiple claims",
			args: args{
				claims: Claims{
					UserId: uuid.New().String(),
				},
				validity: 20 * time.Minute,
			},
			wait:          0,
			wantSignErr:   false,
			wantVerifyErr: false,
		},
		{
			name: "an expired token should not be validated",
			args: args{
				claims: Claims{
					UserId: uuid.New().String(),
				},
				validity: 20 * time.Millisecond,
			},
			wait:          25 * time.Millisecond,
			wantSignErr:   false,
			wantVerifyErr: true,
		},
		{
			name: "a valid token should be issued and not validated with invalid userId",
			args: args{
				claims:   Claims{UserId: ""},
				validity: 20 * time.Minute,
			},
			wait:          0,
			wantSignErr:   false,
			wantVerifyErr: true,
		},
		{
			name: "a valid token should be issued and not validated with absent userId",
			args: args{
				claims:   Claims{},
				validity: 20 * time.Minute,
			},
			wait:          0,
			wantSignErr:   false,
			wantVerifyErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, signErr := SignToken(tt.args.claims, tt.args.validity)
			if (signErr != nil) != tt.wantSignErr {
				t.Errorf("SignToken() error = %v, wantSignErr %v", signErr, tt.wantSignErr)
				return
			}

			if len(got) == 0 {
				t.Errorf("SignToken() got = %v, want non empty string", got)
			}

			time.Sleep(tt.wait)
			verifiedToken, verifyErr := verifyToken(got)
			if (verifyErr != nil) != tt.wantVerifyErr {
				t.Errorf("verifyToken() error = %v, wantVerifyErr %v", verifyErr, tt.wantVerifyErr)
				return
			}

			if verifyErr == nil {
				if verifiedToken.UserId != tt.args.claims.UserId {
					t.Errorf("verifyToken() got userId %s, want %s", verifiedToken.UserId, tt.args.claims.UserId)
				}
			}
		})
	}
}
