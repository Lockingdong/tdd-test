package post

import "github.com/google/uuid"

type Post struct {
	ID      string
	Title   string
	Content string
	Author  string
}

func NewPost(title, content, author string) *Post {
	return &Post{
		ID:      uuid.NewString(),
		Title:   title,
		Content: content,
		Author:  author,
	}
}

type PresenterConfig struct {
	Post        *Post
	Anonymous   bool
	CustomTitle string
}

func NewPresenterConfig(
	post *Post,
	anonymous bool,
	customTitle string,
) *PresenterConfig {
	return &PresenterConfig{
		Post:        post,
		Anonymous:   anonymous,
		CustomTitle: customTitle,
	}
}

func (c *PresenterConfig) GetAuthor() string {
	if c.Anonymous {
		return "匿名作者"
	}
	return c.Post.Author
}

func (c *PresenterConfig) GetTitle() string {
	if c.CustomTitle != "" {
		return c.CustomTitle
	}
	return c.Post.Title
}

// PostPresenter is a struct that holds a Post
type PostPresenter struct {
	Post   *Post
	Config *PresenterConfig
}

func NewPostPresenter(post *Post, config *PresenterConfig) *PostPresenter {
	return &PostPresenter{
		Post:   post,
		Config: config,
	}
}

func (p *PostPresenter) GetAuthor() string {
	return p.Config.GetAuthor()
}

func (p *PostPresenter) GetTitle() string {
	return p.Config.GetTitle()
}

// AnonymousPostPresenter is a struct that holds a Post
// type AnonymousPostPresenter struct {
// 	*PostPresenter
// }

// func NewAnonymousPostPresenter(post *Post) *AnonymousPostPresenter {
// 	return &AnonymousPostPresenter{
// 		PostPresenter: NewPostPresenter(post),
// 	}
// }

// func (p *AnonymousPostPresenter) GetAuthor() string {
// 	return "匿名作者"
// }

// type CustomTitlePostPresenter struct {
// 	*PostPresenter
// 	CustomTitle string
// }

// func NewCustomTitlePostPresenter(
// 	post *Post,
// 	customTitle string,
// ) *CustomTitlePostPresenter {
// 	return &CustomTitlePostPresenter{
// 		PostPresenter: NewPostPresenter(post, nil),
// 		CustomTitle:   customTitle,
// 	}
// }

// func (p *CustomTitlePostPresenter) GetTitle() string {
// 	return p.CustomTitle
// }
