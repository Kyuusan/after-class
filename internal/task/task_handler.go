package task

import (
	"strconv"
	"tasklybe/internal/dto"
	"github.com/gofiber/fiber/v2"
)

func HandleGetTask(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))

	task, totalData, err := GetTasks(page, limit)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ResponseWrapper[[]Task]{
			Data:    nil,
			Success: false,
			Message: "Internal Server Error",
		})
	}

	totalPages := 0
	if limit > 0 {
		totalPages = int((totalData+int64(limit))-1) / limit

	}

	return c.JSON(dto.ResponseWrapper[[]Task]{
		Data:    task,
		Success: true,
		Message: "Succes! task su ccessfully fetched",
		Pagination: &dto.PaginationResponse{
			Page:      page,
			Limit:     limit,
			Total:     totalData,
			TotalPage: totalPages,
		},
	})
}

func HandleGetDetailTask( c *fiber.Ctx) error {
	id := c.Params("id")
	task, err := GetDetailTask(id)
	return c.Status(fiber.StatusInternalServerError).JSON(dto.ResponseWrapper[Task]{
		Data: nil,
		Success:  false,
		Message : err.Error(),
	})
	return c.JSON(dto.ResponseWrapper[Task]{
		Data: task,
		Success: true,
		Message: "Succes! Detail task successfully fetched",
	})

	
}