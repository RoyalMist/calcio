package middlewares

import (
	"testing"
	"time"

	"github.com/google/uuid"
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
				t.Errorf("HashPassword() error = %v, wantSignErr %v", err, tt.wantErr)
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

func TestSignVerifyToken(t *testing.T) {
	type args struct {
		userId   string
		claims   map[string]string
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
			name: "a valid token should be issued and validated with no claims",
			args: args{
				userId:   uuid.New().String(),
				claims:   make(map[string]string, 0),
				validity: 20 * time.Minute,
			},
			wait:          0,
			wantSignErr:   false,
			wantVerifyErr: false,
		},
		{
			name: "a valid token should be issued and validated with one claim",
			args: args{
				userId:   uuid.New().String(),
				claims:   map[string]string{"key": "value"},
				validity: 20 * time.Minute,
			},
			wait:          0,
			wantSignErr:   false,
			wantVerifyErr: false,
		},
		{
			name: "a valid token should be issued and validated with multiple claims",
			args: args{
				userId:   uuid.New().String(),
				claims:   map[string]string{"key1": "value1", "key2": "value2", "key3": "value3"},
				validity: 20 * time.Minute,
			},
			wait:          0,
			wantSignErr:   false,
			wantVerifyErr: false,
		},
		{
			name: "an expired token should not be validated",
			args: args{
				userId:   uuid.New().String(),
				claims:   map[string]string{"key1": "value1", "key2": "value2", "key3": "value3"},
				validity: 20 * time.Millisecond,
			},
			wait:          25 * time.Millisecond,
			wantSignErr:   false,
			wantVerifyErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, signErr := SignToken(tt.args.userId, tt.args.claims, tt.args.validity)
			if (signErr != nil) != tt.wantSignErr {
				t.Errorf("SignToken() error = %v, wantSignErr %v", signErr, tt.wantSignErr)
				return
			}

			if len(got) == 0 {
				t.Errorf("SignToken() got = %v, want non empty string", got)
			}

			time.Sleep(tt.wait)
			verifiedToken, verifyErr := VerifyToken(got)
			if (verifyErr != nil) != tt.wantVerifyErr {
				t.Errorf("VerifyToken() error = %v, wantVerifyErr %v", verifyErr, tt.wantVerifyErr)
				return
			}

			if verifyErr == nil {
				if verifiedToken.Subject != tt.args.userId {
					t.Errorf("VerifyToken() should retrive the provided subject %s, but found %s", tt.args.userId, verifiedToken.Subject)
				}

				for k, v := range tt.args.claims {
					gotValue := verifiedToken.Get(k)
					if v != gotValue {
						t.Errorf("VerifyToken() should give the value %s for key %s, but got %s", v, k, gotValue)
					}
				}
			}
		})
	}
}
