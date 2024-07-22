package post

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"gogram/internal/app/auth"
	post_requests2 "gogram/internal/app/post/post-requests"
	"gogram/pkg/helpers"
	"gogram/pkg/pagination"
	"gogram/pkg/validation"
)

type PostHandler struct {
	service *PostService
}

func NewPostHandler() *PostHandler {
	return &PostHandler{
		service: NewPostService(),
	}
}

func (ph *PostHandler) Index(ctx *fiber.Ctx) error {
	page := ctx.QueryInt("page", 1)
	limit := ctx.QueryInt("limit", 10)

	posts, err := ph.service.GetAll(page, limit)

	if err != nil {
		return err
	}

	res := helpers.NewResponseHelper(
		fiber.StatusOK,
		"Get all posts success",
		posts,
		pagination.NewPagination(page, limit),
	)

	return ctx.Status(res.Code).JSON(res)
}

func (ph *PostHandler) Create(ctx *fiber.Ctx) error {
	req := new(post_requests2.CreatePostRequest)

	if err := validation.Validate[post_requests2.CreatePostRequest](ctx, req); err != nil {
		return err
	}

	if thumbnail, err := ctx.FormFile("thumbnail"); err != nil {
		req.Thumbnail = nil
	} else {
		req.Thumbnail = thumbnail
	}

	userClaims, err := auth.User(ctx)

	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	post, err := ph.service.Store(ctx, req, &userClaims.User)

	if err != nil {
		return err
	}

	res := helpers.NewResponseHelper(
		fiber.StatusCreated,
		"Create post success",
		post,
		nil,
	)

	return ctx.Status(res.Code).JSON(res)
}

func (ph *PostHandler) Show(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	post, err := ph.service.Find(id)

	if err != nil {
		return err
	}

	res := helpers.NewResponseHelper(
		fiber.StatusOK,
		"Get post success",
		post,
		nil,
	)

	return ctx.Status(res.Code).JSON(res)
}

func (ph *PostHandler) Update(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	req := new(post_requests2.EditPostRequest)

	if err := validation.Validate[post_requests2.EditPostRequest](ctx, req); err != nil {
		return err
	}

	if thumbnail, err := ctx.FormFile("thumbnail"); err != nil {
		req.Thumbnail = nil
	} else {
		req.Thumbnail = thumbnail
	}

	post, err := ph.service.Update(ctx, id, req)

	if err != nil {
		return err
	}

	res := helpers.NewResponseHelper(
		fiber.StatusOK,
		"Update post success",
		post,
		nil,
	)

	return ctx.Status(res.Code).JSON(res)
}

func (ph *PostHandler) Delete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	err := ph.service.Delete(id)

	if err != nil {
		return err
	}

	res := helpers.NewResponseHelper(
		fiber.StatusOK,
		"Delete post success",
		nil,
		nil,
	)

	return ctx.Status(res.Code).JSON(res)
}

func (ph *PostHandler) GetFromJSONPlaceholder(ctx *fiber.Ctx) error {
	agent := fiber.Get("https://jsonplaceholder.typicode.com/posts")

	statusCode, body, errs := agent.Bytes()

	if len(errs) > 0 {
		return fiber.NewError(fiber.StatusInternalServerError, errs[0].Error())
	}

	// karena response dari json placeholder dalam bentuk array/slice bukan object
	var data []fiber.Map

	err := json.Unmarshal(body, &data)

	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	res := helpers.NewResponseHelper(
		statusCode,
		"Get posts from JSONPlaceholder success",
		data,
		nil,
	)

	return ctx.Status(res.Code).JSON(res)
}
