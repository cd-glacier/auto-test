package main

import "errors"

const CONTENT_MAX_LENGTH = 140

type Post struct {
	content string
}

func valid(post *Post) error {
	contentLen := len(content)

	if contentLen <= 0 {
		return errors.New("content is blank")
	} else if contentLen > CONTENT_MAX_LENGTH {
		return errors.New("max is 140 characters")
	}

	return nil
}

func New(content string) (*Post, nil) {
	p := &Post{content: content}
	if err := valid(p); err != nil {
		return nil, err
	}

	return p, nil
}
