package access_token

import (
	"reflect"
	"testing"
)

func TestAccessToken_Get(t *testing.T) {
	type fields struct {
		Api            string
		AppId          string
		Secret         string
		Response       Response
		IsCache        bool
		CacheDirectory string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := &AccessToken{
				Api:            tt.fields.Api,
				AppId:          tt.fields.AppId,
				Secret:         tt.fields.Secret,
				Response:       tt.fields.Response,
				IsCache:        tt.fields.IsCache,
				CacheDirectory: tt.fields.CacheDirectory,
			}
			if got := handler.Get(); got != tt.want {
				t.Errorf("AccessToken.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAccessToken_getFromServe(t *testing.T) {
	type fields struct {
		Api            string
		AppId          string
		Secret         string
		Response       Response
		IsCache        bool
		CacheDirectory string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := &AccessToken{
				Api:            tt.fields.Api,
				AppId:          tt.fields.AppId,
				Secret:         tt.fields.Secret,
				Response:       tt.fields.Response,
				IsCache:        tt.fields.IsCache,
				CacheDirectory: tt.fields.CacheDirectory,
			}
			if got := handler.getFromServe(); got != tt.want {
				t.Errorf("AccessToken.getFromServe() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_generateFilename(t *testing.T) {
	type args struct {
		key string
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
			if got := generateFilename(tt.args.key); got != tt.want {
				t.Errorf("generateFilename() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_readFile(t *testing.T) {
	type args struct {
		saveDirecoty string
		filename     string
	}
	tests := []struct {
		name  string
		args  args
		want  bool
		want1 Response
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := readFile(tt.args.saveDirecoty, tt.args.filename)
			if got != tt.want {
				t.Errorf("readFile() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("readFile() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_saveFile(t *testing.T) {
	type args struct {
		saveDirecoty string
		filename     string
		res          Response
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := saveFile(tt.args.saveDirecoty, tt.args.filename, tt.args.res); got != tt.want {
				t.Errorf("saveFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
