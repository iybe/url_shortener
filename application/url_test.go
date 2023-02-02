package application

import (
	"testing"

	"github.com/gin-gonic/gin"
)

func TestUrlApp_ShortenUrl(t *testing.T) {
	type fields struct {
		SizeUrlShortened int
	}
	type args struct {
		url string
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantSize int
		wantErr  error
	}{
		{
			name: "success",
			fields: fields{
				SizeUrlShortened: 6,
			},
			args:     args{},
			wantSize: 6,
			wantErr:  nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UrlApp{
				SizeUrlShortened: tt.fields.SizeUrlShortened,
			}
			got, err := u.ShortenUrl(&gin.Context{}, tt.args.url)
			if err != tt.wantErr {
				t.Errorf("UrlApp.ShortenUrl() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) != tt.wantSize {
				t.Errorf("UrlApp.ShortenUrl() = %v, want %v", got, tt.wantSize)
			}
		})
	}
}
