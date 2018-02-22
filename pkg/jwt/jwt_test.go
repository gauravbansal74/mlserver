package jwt

import "testing"

var (
	email = "test@example.com"
	ID    = "1233213213213"
	key   = "unit-test"
)

func TestGetToken(t *testing.T) {
	type args struct {
		email string
		ID    string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Email-ID not empty",
			args: args{
				email: email,
				ID:    ID,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			token := GetToken(tt.args.email, tt.args.ID)
			if token == "" {
				t.Errorf("GetToken() token value can't be empty =  %v", token)
			}
			userData, err := GetData(token)
			if err != nil {
				t.Errorf("GetData() =  %v", err.Error())
			}
			if userData.Email != tt.args.email {
				t.Errorf("GetData() email = %v, want %v", userData.Email, tt.args.email)
			}
			if userData.ID != tt.args.ID {
				t.Errorf("GetData() ID= %v, want %v", userData.ID, tt.args.ID)
			}
		})
	}
}
func TestConfigInfo(t *testing.T) {
	ConfigInfo(Info{
		Key: key,
	})
	infoData := readConfig()
	if infoData.Key != key {
		t.Error("key value should be same as Key")
	}
}

func TestGetDataEmpty(t *testing.T) {
	_, err := GetData("")
	if err == nil {
		t.Error(err.Error())
	}
}
