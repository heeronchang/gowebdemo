package myresty

import (
	"testing"

	"github.com/go-resty/resty/v2"
)

func TestGetClient(t *testing.T) {
	tests := []struct {
		name string
		want *resty.Client
	}{
		{
			name: "myresty",
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// if got := GetClient(); !reflect.DeepEqual(got, tt.want) {
			// 	t.Errorf("GetClient() = %v, want %v", got, tt.want)
			// }

			if c := GetClient(); c != nil {
				res, err := c.R().Get("https://api.github.com/")
				if err != nil {
					t.Error(err)
				}

				t.Log(res)
			}
		})
	}
}
