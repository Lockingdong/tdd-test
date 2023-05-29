package post

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPostPresenter_GetAuthor(t *testing.T) {
	type fields struct {
		Post            *Post
		PresenterConfig *PresenterConfig
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "GetAuthor",
			fields: fields{
				Post: NewPost("test_title", "test_content", "test_author"),
				PresenterConfig: NewPresenterConfig(
					NewPost("test_title", "test_content", "test_author"),
					false,
					"",
				),
			},
			want: "test_author",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := NewPostPresenter(tt.fields.Post, tt.fields.PresenterConfig)
			assert.Equal(t, tt.want, p.GetAuthor())
		})
	}
}

func TestPostPresenter_GetAnonymousAuthor(t *testing.T) {
	type fields struct {
		Post            *Post
		PresenterConfig *PresenterConfig
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "GetAuthor",
			fields: fields{
				Post: NewPost("test_title", "test_content", "test_author"),
				PresenterConfig: NewPresenterConfig(
					NewPost("test_title", "test_content", "test_author"),
					true,
					"",
				),
			},
			want: "匿名作者",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := NewPostPresenter(tt.fields.Post, tt.fields.PresenterConfig)
			assert.Equal(t, tt.want, p.GetAuthor())
		})
	}
}

func TestPostPresenter_GetCustomTitle(t *testing.T) {
	type fields struct {
		Post            *Post
		CustomTitle     string
		PresenterConfig *PresenterConfig
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "GetAuthor",
			fields: fields{
				Post: NewPost("test_title", "test_content", "test_author"),
				PresenterConfig: NewPresenterConfig(
					NewPost("test_title", "test_content", "test_author"),
					false,
					"custom_title",
				),
			},
			want: "custom_title",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := NewPostPresenter(tt.fields.Post, tt.fields.PresenterConfig)
			assert.Equal(t, tt.want, p.GetTitle())
		})
	}
}
