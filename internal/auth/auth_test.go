package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	// First, we need to create some headers for testing
	header1 := http.Header{}
	token_str := "HelloWorld"
	header1.Set("Authorization", "ApiKey "+token_str)
	header2 := http.Header{}
	header3 := http.Header{}
	header3.Set("Authorization", "api_key "+token_str)

	tests := []struct {
		name      string
		header    http.Header
		wantToken string
		wantErr   bool
	}{
		{
			name:      "Correct APIKey",
			header:    header1,
			wantToken: token_str,
			wantErr:   false,
		},
		{
			name:      "Missing header",
			header:    header2,
			wantToken: "",
			wantErr:   true,
		},
		{
			name:      "Malformed authorization header",
			header:    header3,
			wantToken: "",
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotToken, err := GetAPIKey(tt.header)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAPIKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotToken != tt.wantToken {
				t.Errorf("GetAPIKey() gotToken = %v, want %v",
					gotToken, tt.wantToken)
			}
		})
	}
}
