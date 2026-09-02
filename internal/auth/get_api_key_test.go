package auth

import (
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_GetApiKey(t *testing.T) {

	tests := map[string]struct {
		input   http.Header
		want    string
		wantErr bool
	}{

		"malformed_header": {input: http.Header{"Authorization": {"Api_Key YOUR_API_KEY"}}, want: "YOUT_API_KEY", wantErr: true},
		"no_header":        {input: http.Header{}, want: "Your_juice", wantErr: true},
		"she_works":        {input: http.Header{"Authorization": {"ApiKey YOUR_API_KEY"}}, want: "YOUR_API_KEY", wantErr: false},
	}

	for name, tc := range tests {

		t.Run(name, func(t *testing.T) {
			got, err := GetAPIKey(tc.input)

			if tc.wantErr {
				if err == nil {
					t.Fatalf("expected error, got nil")
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			diff := cmp.Diff(tc.want, got)
			if diff != "" {
				t.Fatalf("GetApiKey mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
