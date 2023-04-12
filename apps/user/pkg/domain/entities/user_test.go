package entities

import (
	"testing"
)

func TestParseUserID(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    UserID
		wantErr bool
	}{
		{
			name:    "ValidUserID",
			input:   "123",
			want:    123,
			wantErr: false,
		},
		{
			name:    "InvalidUserID",
			input:   "invalid",
			want:    0,
			wantErr: true,
		},
		{
			name:    "EmptyUserID",
			input:   "",
			want:    0,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uid, err := ParseUserID(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseUserID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if uid != tt.want {
				t.Errorf("ParseUserID() = %v, want %v", uid, tt.want)
			}
		})
	}
}
