package services

import (
	"github.com/stretchr/testify/mock"
	postInternal "gogram/internal/app/post"
)

type PostServiceMock struct {
	Mock mock.Mock
}

func (psm *PostServiceMock) Find(id string) (post *postInternal.Post, err error) {
	arguments := psm.Mock.Called(id)

	if arguments.Get(0) != nil {
		return arguments.Get(0).(*postInternal.Post), nil
	}

	return nil, arguments.Error(1)
}
