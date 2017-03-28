package transfer

import (
	"crypto/tls"
	"reflect"
	"testing"
)

func TestTransfer_Post(t *testing.T) {
	type fields struct {
		Api           string
		Pem           Pem
		Request       *PostRequest
		Response      *PostResponse
		QueryRequest  *QueryRequest
		QueryResponse *QueryResponse
	}
	type args struct {
		pem  Pem
		requ *PostRequest
		resp *PostResponse
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
			handler := &Transfer{
				Api:           tt.fields.Api,
				Pem:           tt.fields.Pem,
				Request:       tt.fields.Request,
				Response:      tt.fields.Response,
				QueryRequest:  tt.fields.QueryRequest,
				QueryResponse: tt.fields.QueryResponse,
			}
			if got := handler.Post(tt.args.pem, tt.args.requ, tt.args.resp); got != tt.want {
				t.Errorf("Transfer.Post() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_loadPEM(t *testing.T) {
	type args struct {
		CertFilePath string
		KeyFilePath  string
		CAFilePath   string
	}
	tests := []struct {
		name string
		args args
		want *tls.Config
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := loadPEM(tt.args.CertFilePath, tt.args.KeyFilePath, tt.args.CAFilePath); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("loadPEM() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTransfer_Query(t *testing.T) {
	type fields struct {
		Api           string
		Pem           Pem
		Request       *PostRequest
		Response      *PostResponse
		QueryRequest  *QueryRequest
		QueryResponse *QueryResponse
	}
	tests := []struct {
		name   string
		fields fields
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := &Transfer{
				Api:           tt.fields.Api,
				Pem:           tt.fields.Pem,
				Request:       tt.fields.Request,
				Response:      tt.fields.Response,
				QueryRequest:  tt.fields.QueryRequest,
				QueryResponse: tt.fields.QueryResponse,
			}
			handler.Query()
		})
	}
}

func Test_generateSign(t *testing.T) {
	type args struct {
		signParam []map[string]string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateSign(tt.args.signParam); got != tt.want {
				t.Errorf("generateSign() = %v, want %v", got, tt.want)
			}
		})
	}
}
