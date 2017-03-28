package pay

import "testing"

func TestWxPay_initPost(t *testing.T) {
	type fields struct {
		url           string
		Account       WxPayAccount
		PostRequest   WxPostRequest
		PostResponse  WxPostResponse
		QueryRequest  WxQueryRequest
		QueryResponse WxQueryResponse
	}
	tests := []struct {
		name   string
		fields fields
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wp := &WxPay{
				url:           tt.fields.url,
				Account:       tt.fields.Account,
				PostRequest:   tt.fields.PostRequest,
				PostResponse:  tt.fields.PostResponse,
				QueryRequest:  tt.fields.QueryRequest,
				QueryResponse: tt.fields.QueryResponse,
			}
			wp.initPost()
		})
	}
}

func TestWxPay_Post(t *testing.T) {
	type fields struct {
		url           string
		Account       WxPayAccount
		PostRequest   WxPostRequest
		PostResponse  WxPostResponse
		QueryRequest  WxQueryRequest
		QueryResponse WxQueryResponse
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
			wp := &WxPay{
				url:           tt.fields.url,
				Account:       tt.fields.Account,
				PostRequest:   tt.fields.PostRequest,
				PostResponse:  tt.fields.PostResponse,
				QueryRequest:  tt.fields.QueryRequest,
				QueryResponse: tt.fields.QueryResponse,
			}
			if got := wp.Post(); got != tt.want {
				t.Errorf("WxPay.Post() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWxPay_initQuery(t *testing.T) {
	type fields struct {
		url           string
		Account       WxPayAccount
		PostRequest   WxPostRequest
		PostResponse  WxPostResponse
		QueryRequest  WxQueryRequest
		QueryResponse WxQueryResponse
	}
	tests := []struct {
		name   string
		fields fields
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wp := &WxPay{
				url:           tt.fields.url,
				Account:       tt.fields.Account,
				PostRequest:   tt.fields.PostRequest,
				PostResponse:  tt.fields.PostResponse,
				QueryRequest:  tt.fields.QueryRequest,
				QueryResponse: tt.fields.QueryResponse,
			}
			wp.initQuery()
		})
	}
}

func TestWxPay_Query(t *testing.T) {
	type fields struct {
		url           string
		Account       WxPayAccount
		PostRequest   WxPostRequest
		PostResponse  WxPostResponse
		QueryRequest  WxQueryRequest
		QueryResponse WxQueryResponse
	}
	tests := []struct {
		name   string
		fields fields
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wp := &WxPay{
				url:           tt.fields.url,
				Account:       tt.fields.Account,
				PostRequest:   tt.fields.PostRequest,
				PostResponse:  tt.fields.PostResponse,
				QueryRequest:  tt.fields.QueryRequest,
				QueryResponse: tt.fields.QueryResponse,
			}
			wp.Query()
		})
	}
}

func TestGenerateSign(t *testing.T) {
	type args struct {
		params []map[string]string
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
			if got := GenerateSign(tt.args.params); got != tt.want {
				t.Errorf("GenerateSign() = %v, want %v", got, tt.want)
			}
		})
	}
}
