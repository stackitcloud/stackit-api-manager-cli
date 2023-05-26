package util

import "testing"

func TestEncodeBase64File(t *testing.T) {
	type args struct {
		file string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				file: "./test_data/test.json",
			},
			want:    "eyAiaGVsbG8iOiAiZHVkZXR0ZSIgfQo=",
			wantErr: false,
		},
		{
			name: "file not found",
			args: args{
				file: "./test_data/test-not-found.json",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := EncodeBase64File(tt.args.file)
			if (err != nil) != tt.wantErr {
				t.Errorf("EncodeBase64File() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("EncodeBase64File() = %v, want %v", got, tt.want)
			}
		})
	}
}
