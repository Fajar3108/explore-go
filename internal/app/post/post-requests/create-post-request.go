package post_requests

import "mime/multipart"

type CreatePostRequest struct {
	Title     string                `json:"title" validate:"required,min=5,max=255"`
	Body      string                `json:"body" validate:"required"`
	Thumbnail *multipart.FileHeader `json:"thumbnail" validate:"omitempty,image"`
}
