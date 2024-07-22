package services

import (
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	postInternal "gogram/internal/app/post"
	"testing"
)

var postServiceMock = &PostServiceMock{
	Mock: mock.Mock{},
}

func TestPostService_Find(t *testing.T) {
	postId2 := &postInternal.Post{
		ID: 2,
	}

	errorNotFound := fiber.NewError(fiber.StatusNotFound, "Post id not found")

	postServiceMock.Mock.On("Find", "1").Return(nil, errorNotFound)
	postServiceMock.Mock.On("Find", "2").Return(postId2, nil)

	t.Run("Post ID Not Found", func(t *testing.T) {
		post, err := postServiceMock.Find("1")

		assert.Nil(t, post)
		assert.NotNil(t, err)
		assert.Equal(t, errorNotFound, err)
	})

	t.Run("Post ID Found", func(t *testing.T) {
		post, err := postServiceMock.Find("2")

		assert.NotNil(t, post)
		assert.Nil(t, err)
		assert.Equal(t, postId2.ID, post.ID)
	})
}
