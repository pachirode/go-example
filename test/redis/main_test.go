package main

import (
	"github.com/go-redis/redismock/v9"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLogin(t *testing.T) {
	rdb, mock := redismock.NewClientMock()

	mock.ExpectGet("sms:captcha:XXX").SetVal("123456")
	mock.ExpectGet("sms:captcha:Y").RedisNil()
	mock.ExpectGet("sms:captcha:Z").SetVal("123456")

	type args struct {
		mobile  string
		smsCode string
	}

	tests := []struct {
		name    string
		args    args
		want    string
		wantErr string
	}{{
		name: "login success",
		args: args{
			mobile:  "XXX",
			smsCode: "123456",
		},
		want: "123456",
	},
		{
			name: "invalid sms code or expired",
			args: args{
				mobile:  "Y",
				smsCode: "123456",
			},
			wantErr: "invalid sms code or expired",
		},
		{
			name: "invalid sms code",
			args: args{
				mobile:  "Z",
				smsCode: "123457",
			},
			wantErr: "invalid sms code",
		}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Login(tt.args.mobile, tt.args.smsCode, rdb)
			if tt.wantErr != "" {
				assert.Error(t, err)
				assert.Equal(t, tt.wantErr, err.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}

}
