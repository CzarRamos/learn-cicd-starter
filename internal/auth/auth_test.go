package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {

	tests := map[string]struct {
		name  string
		input http.Header
		want  string
	}{
		"No Header":              {"no header", http.Header{}, ""},
		"Empty Header":           {"empty header", http.Header{"Authorization": []string{""}}, ""},
		"Wrong Schema":           {"wrong schema", http.Header{"Authorization": []string{"Bearer xyz"}}, ""},
		"Missing Key":            {"missing key", http.Header{"Authorization": []string{"ApiKey"}}, ""},
		"Key is only Whitespace": {"key is only whitespace", http.Header{"Authorization": []string{"ApiKey "}}, ""},
		"Valid Key":              {"should be valid key", http.Header{"Authorization": []string{"ApiKey apikey123"}}, "apikey123"},
		"Wrong Case":             {"wrong case format", http.Header{"Authorization": []string{"apikey apikey123"}}, ""},
		"Leading Spaces":         {"contains leading spaces", http.Header{"Authorization": []string{"  ApiKey apikey123"}}, ""},
	}

	for name, testCase := range tests {
		got, _ := GetAPIKey(testCase.input)
		if got != testCase.want {
			t.Fatalf("%s: expected: %v, got: %v", name, testCase.want, got)
		}
	}
}
