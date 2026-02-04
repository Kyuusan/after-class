package task

import (
	"strconv"
	"tasklybe/internal/dto"
	"tasklybe/internal/validation"
	"github.com/gofiber/fiber/v2"
)


// GetTasks godoc
// @Summary Get tasks
// @Description Get all tasks
// @Tags tasks
// @Produce json
// @Success 200 {array} Task
// @Router /task [get]
// @Param page query int true "Page number"
// @Param limit query int true "Limit per page"
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




// GetTask godoc
// @Summary Get task
// @Description Get task by id
// @Tags tasks
// @Produce json
// @Success 200 {object} Task
// @Router /task/{id} [get]
// @Param id path string true "Task ID"
func HandleGetDetailTask(c *fiber.Ctx) error {
	id := c.Params("id")

	task, err := GetDetailTask(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(dto.ResponseWrapper[Task]{
			Data:    nil,
			Success: false,
			Message: err.Error(),
		})
	}

	return c.JSON(dto.ResponseWrapper[Task]{
		Data:    task,
		Success: true,
		Message: "Success! Detail task successfully fetched",
	})
}


// CreateTask godoc
// @Summary Create task
// @Description Create a new task
// @Tags tasks
// @Produce json
// @Success 200 {object} Task
// @Router /task [post]
// @Param req body CreateTaskRequest true "Create Task Request"
func HandleCreateTask(c *fiber.Ctx) error {

  // Disini kita deklarasi variable req yang menampung data request user
	var req CreateTaskRequest
	
	// Kemudian variable req divalidasi apakah request nya sesuai atau tidak
	if err := validation.BindAndValidate(c, &req); err != nil {
		// Jika tidak kita return sebagai error validation
		return c.Status(fiber.StatusBadRequest).JSON(dto.ResponseWrapper[Task]{
			Data:    nil,
			Success: false,
			Message: "Failed! Validation error.",
			Error:   validation.FormatValidationError(err),
		})
	}

	// Memanggil service CreateTask
	task, err := CreateTask(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ResponseWrapper[Task]{
			Data:    nil,
			Success: false,
			Message: "Failed! Something went wrong.",
		})
	}

	return c.JSON(dto.ResponseWrapper[Task]{
		Data:    task,
		Success: true,
		Message: "Success! task created.",
	})
}

// EditTask godoc
// @Summary Edit task
// @Description Edit task by id
// @Tags tasks
// @Produce json
// @Success 200 {object} Task
// @Router /task/{id} [put]
// @Param id path string true "Task ID"
// @Param req body EditTaskRequest true "Edit Task Request"
func HandleEditTask(c *fiber.Ctx) error {

	/** 
	Ambil value id dari request parameter
	Contoh: http://localhost:3000/task/[xxx-xxxx-xxxx-xxxx] <- ini adalah request param
	**/
	id := c.Params("id")

	// Lakukan validasi request
	var req EditTaskRequest
	if err := validation.BindAndValidate(c, &req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ResponseWrapper[Task]{
			Data:    nil,
			Success: false,
			Message: "Failed! Validation error.",
			Error:   validation.FormatValidationError(err),
		})
	}

	// Panggil service EditTask
	task, err := EditTask(id, req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ResponseWrapper[Task]{
			Data:    nil,
			Success: false,
			Message: "Failed! Something went wrong.",
		})
	}

	return c.JSON(dto.ResponseWrapper[Task]{
		Data:    task,
		Success: true,
		Message: "Success! task updated.",
	})
}

// DeleteTask godoc
// @Summary Delete task
// @Description Delete task by id
// @Tags tasks
// @Produce json
// @Success 200 {object} Task
// @Router /task/{id} [delete]
// @Param id path string true "Task ID"
func HandleDeleteTask(c *fiber.Ctx) error {
	// Ambil id task dari url param
	id := c.Params("id")
	
	// Panggil fungsi DeleteTask, sambil di cek apakah ada error, jika ada return 500
	if err := DeleteTask(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ResponseWrapper[Task]{
			Data:    nil,
			Success: false,
			Message: "Failed! Something went wrong.",
		})
	}

	return c.JSON(dto.ResponseWrapper[Task]{
		Success: true,
		Message: "Success! task deleted.",
	})
}