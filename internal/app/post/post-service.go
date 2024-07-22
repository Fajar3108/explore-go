package post

import (
	"database/sql"
	"github.com/gofiber/fiber/v2"
	"gogram/internal/app/auth"
	post_requests2 "gogram/internal/app/post/post-requests"
	"gogram/internal/app/user"
	"gogram/internal/database"
	file_storage "gogram/pkg/file-storage"
	"gorm.io/gorm"
)

type PostService struct {
	db *gorm.DB
}

func NewPostService() *PostService {
	return &PostService{
		db: database.InitDB(),
	}
}

func (ps *PostService) GetAll(page, limit int) (posts *[]Post, err error) {
	offset := (page - 1) * limit

	result := ps.db.Preload("User").Order("id desc").Offset(offset).Limit(limit).Find(&posts)

	if result.Error != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, result.Error.Error())
	}

	return posts, nil
}

func (ps *PostService) Store(ctx *fiber.Ctx, request *post_requests2.CreatePostRequest, usr *user.User) (post *Post, err error) {
	post = &Post{
		Title: request.Title,
		Body:  request.Body,
		User:  *usr,
	}

	if request.Thumbnail != nil {
		filePath, err := file_storage.Store(ctx, request.Thumbnail, "thumbnails")

		if err != nil {
			return nil, err
		}

		post.Thumbnail = sql.NullString{
			String: filePath,
			Valid:  true,
		}
	}

	result := ps.db.Create(post)

	if result.Error != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, result.Error.Error())
	}

	return post, nil
}

func (ps *PostService) Find(id string) (post *Post, err error) {
	result := ps.db.Preload("User").Find(&post, id)

	if result.Error != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, result.Error.Error())
	}

	if post.ID <= 0 {
		return nil, fiber.NewError(fiber.StatusNotFound, "Post id not found")
	}

	return post, nil
}

func (ps *PostService) Update(ctx *fiber.Ctx, id string, request *post_requests2.EditPostRequest) (post *Post, err error) {
	post, err = ps.Find(id)

	if err != nil {
		return nil, err
	}

	userClaims, err := auth.User(ctx)

	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	if post.UserID != userClaims.User.ID {
		return nil, fiber.NewError(fiber.StatusUnauthorized, "Unauthorized")
	}

	post.Title = request.Title
	post.Body = request.Body

	if request.Thumbnail != nil {
		if post.Thumbnail.Valid {
			if err = file_storage.Remove(post.Thumbnail.String); err != nil {
				return nil, err
			}
		}

		filePath, err := file_storage.Store(ctx, request.Thumbnail, "thumbnails")

		if err != nil {
			return nil, err
		}

		post.Thumbnail = sql.NullString{
			String: filePath,
			Valid:  true,
		}
	}

	result := ps.db.Preload("User").Save(post)

	if result.Error != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, result.Error.Error())
	}

	return post, nil
}

func (ps *PostService) Delete(id string) (err error) {
	post, err := ps.Find(id)

	if err != nil {
		return err
	}

	if post.Thumbnail.Valid {
		if err = file_storage.Remove(post.Thumbnail.String); err != nil {
			return err
		}
	}

	result := ps.db.Delete(post)

	if result.Error != nil {
		return fiber.NewError(fiber.StatusInternalServerError, result.Error.Error())
	}

	return nil
}
