package cache

import "testing"

func TestConnectRedis(t *testing.T) {
	type args struct {
		host     string
		password string
		dbi      int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test redis connection",
			args: args{
				host: "localhost:6379",
				dbi:  0,
			},
		},
		{
			name: "test redis connection panic",
			args: args{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if err, ok := recover().(error); ok && err != nil {
					t.Errorf(err.Error())
				}
			}()
			ConnectRedis(tt.args.host, tt.args.password, tt.args.dbi)
		})
	}
}
