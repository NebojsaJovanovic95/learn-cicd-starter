package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKeys(t *testing.T) {
	tests := []struct {
		Name          string
		Header        string
		ExpectedKey   string
		ExpectedError bool
	}{
		{
			Name:          "valid header",
			Header:        "ApiKey abc123",
			ExpectedKey:   "abc123",
			ExpectedError: false,
		},
		{
			Name:          "missing header",
			Header:        "",
			ExpectedError: true,
		},
		{
			Name:          "wrong prefix",
			Header:        "Bearer abc123",
			ExpectedError: true,
		},
		{
			Name:          "no key provided",
			Header:        "ApiKey",
			ExpectedError: true,
		},
		{
			Name:          "extra spaces",
			Header:        "ApiKey    abc123",
			ExpectedError: false,
			ExpectedKey:   "abc123",
		},
	}
	for _, test := range tests {

		t.Run(test.Name, func(t *testing.T) {
			headers := http.Header{}
			if test.Header != "" {
				headers.Set("Authorization", test.Header)
			}

			key, err := GetAPIKey(headers)

			if err == nil && test.ExpectedError {
				t.Fatal("expected error but got none")
			}

			if err != nil && !test.ExpectedError {
				t.Fatal("expected no error but got one")
			}

			if key != test.ExpectedKey {
				t.Fatalf("key %q while %q expected", key, test.ExpectedKey)
			}
		})
	}
}
