package auth

import (
	"crypto/ed25519"
	"testing"

	"github.com/vk-rv/pvx"
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

	auth := New(nil)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := auth.HashPassword(tt.args.password)
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

	auth := New(nil)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := auth.CheckPassword(tt.args.password, tt.args.hash); got != tt.want {
				t.Errorf("CheckPassword() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAuth_SignToken(t *testing.T) {
	publicKey, privateKey, _ := ed25519.GenerateKey(nil)

	type fields struct {
		pv4       *pvx.ProtoV4Public
		publicKey *pvx.AsymPublicKey
		secretKey *pvx.AsymSecretKey
	}

	type args struct {
		claims Claims
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "should retrieve provided claims when validate the token",
			fields: fields{
				pv4:       pvx.NewPV4Public(),
				publicKey: pvx.NewAsymmetricPublicKey(publicKey, pvx.Version4),
				secretKey: pvx.NewAsymmetricSecretKey(privateKey, pvx.Version4),
			},
			args:    args{claims: Claims{UserId: "xxx-111"}},
			want:    "xxx-111",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := Auth{
				pv4:       tt.fields.pv4,
				publicKey: tt.fields.publicKey,
				secretKey: tt.fields.secretKey,
			}

			got, err := a.SignToken(tt.args.claims)
			if (err != nil) != tt.wantErr {
				t.Errorf("SignToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			var claims Claims
			_ = a.pv4.Verify(got, a.publicKey).ScanClaims(&claims)
			if claims.UserId != tt.want {
				t.Errorf("SignToken() got = %v, want %v", claims.UserId, tt.want)
			}
		})
	}
}
