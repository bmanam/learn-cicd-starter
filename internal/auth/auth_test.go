package auth

import (
	"errors"
	"testing"
)

func TestGetAPIKey(t *testing.T) {

	// test set
	tests := []struct {
		name    string
		headers map[string][]string
		want    string
		wantErr error
	}{
		{
			name:    "no auth header",
			headers: map[string][]string{},
			want:    "",
			wantErr: ErrNoAuthHeaderIncluded,
		},
		{
			name:    "malformed auth header",
			headers: map[string][]string{"Authorization": {"ApiKey"}},
			want:    "",
			wantErr: errors.New("malformed authorization header"), //ErrMalformedAuthHeader, //errors.New("malformed authorization header"),
		},
		{
			name:    "valid auth header",
			headers: map[string][]string{"Authorization": {"ApiKey 123"}},
			want:    "123",
			wantErr: nil,
		},
	}

	// run tests
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAPIKey(tt.headers)
			if err != tt.wantErr {
				t.Errorf("GetAPIKey() - Got error:[%v] | Wanted error:[%v]", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetAPIKey() = %v, want %v", got, tt.want)
			}
		})
	}

}
