package message

import (
	_ "fmt"
	"testing"
)

func TestSendRequest_SetParam(t *testing.T) {
	type fields struct {
		Touser     string
		TemplateId string
		Url        string
		Data       map[string]map[string]string
	}
	type args struct {
		key string
		val string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := &SendRequest{
				Touser:     tt.fields.Touser,
				TemplateId: tt.fields.TemplateId,
				Url:        tt.fields.Url,
				Data:       tt.fields.Data,
			}
			handler.SetParam(tt.args.key, tt.args.val)
		})
	}
}

func TestWxSend_Send(t *testing.T) {
	type fields struct {
		Api         string
		AccessToekn string
		Request     *SendRequest
		Response    *SendResponse
	}
	type args struct {
		sreq *SendRequest
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := &WxSend{
				Api:         tt.fields.Api,
				AccessToekn: tt.fields.AccessToekn,
				Request:     tt.fields.Request,
				Response:    tt.fields.Response,
			}
			if got := handler.Send(tt.args.sreq); got != tt.want {
				t.Errorf("WxSend.Send() = %v, want %v", got, tt.want)
			}
		})
	}
}
