package main

import (
	"fmt"

	"github.com/Jett65/questline/APIs/gameAPI/internal/database"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (apiCfg *apiconfig) handlerCreateTask(c *fiber.Ctx) error {
	task := new(Task)

	err := c.BodyParser(task)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, fmt.Sprintf("::::Failed to parse body::: %e", err))
	}

	fmt.Println(task.Game_id)

	gameTask, err := apiCfg.DB.CreateTask(c.Context(), database.CreateTaskParams{
		ID:          uuid.New(),
		Description: ToNullString(task.Description),
		GameID:      task.Game_id,
	})
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, fmt.Sprintf("::::Failed to Create Task:::: %e", err))
	}

	payload, err := databaseTaskToTask(gameTask)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, fmt.Sprintf("::::Failed to Re-Convert Task:::: %e", err))
	}

	return c.JSON(payload)
}

func (apiCfg *apiconfig) handlerGetAllTasks(c *fiber.Ctx) error {
	tasks, err := apiCfg.DB.GetTask(c.Context())
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, fmt.Sprintf("::::Failed to Fetch Tasks:::: %e", err))
	}

	payload, err := databasetasksToTasks(tasks)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, fmt.Sprintf("::::Failed to Convert Tasks:::: %e", err))
	}

	fmt.Println(fiber.StatusAccepted)

	return c.JSON(payload)
}

func (apiCfg *apiconfig) handlerGetTaskById(c *fiber.Ctx) error {
	task_id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, fmt.Sprintf("::::Not a valid ID:::: %e", err))
	}

	task, err := apiCfg.DB.GetTaskById(c.Context(), task_id)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, fmt.Sprintf("::::Failed to Fetch Task:::: %e", err))
	}

	payload, err := databaseTaskToTask(task)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, fmt.Sprintf("::::Failed to Convert Task:::: %e", err))
	}

	fmt.Println(fiber.StatusAccepted)

	return c.JSON(payload)
}

func (apiCfg *apiconfig) handlerUpdateTaskById(c *fiber.Ctx) error {
	task_id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, fmt.Sprintf("::::Not a valid ID:::: %e", err))
	}

	task := new(Task)

	err = c.BodyParser(task)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, fmt.Sprintf("::::Could not Parse Task to Update:::: %e", err))
	}

	Uptask, err := apiCfg.DB.UpdateTask(c.Context(), database.UpdateTaskParams{
		ID:          task_id,
		Description: ToNullString(task.Description),
		GameID:      task.Game_id,
	})
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, fmt.Sprintf("::::Could not Update Task:::: %e", err))
	}

	payload, err := databaseTaskToTask(Uptask)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, fmt.Sprintf("::::Could not Convert Task:::: %e", err))
	}
	return c.JSON(payload)
}

func (apiCfg *apiconfig) handlerDeleteTaskById(c *fiber.Ctx) error {
	task_id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, fmt.Sprintf("::::Not a valid ID:::: %e", err))
	}

	err = apiCfg.DB.DeleteTask(c.Context(), task_id)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, fmt.Sprintf("::::Failed to Delete Task By That ID:::: %e", err))
	}

	return nil
}

func (apiCfg *apiconfig) handlerGetTasksByCatalogGameId(c *fiber.Ctx) error {
	game_id, err := uuid.Parse(c.Params("gameid"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, fmt.Sprintf("::::Not a valid ID:::: %e", err))
	}

	tasks, err := apiCfg.DB.GetTaskByGameId(c.Context(), game_id)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, fmt.Sprintf("::::Failed to Fetch Tasks by Game:::: %e", err))
	}

	payload, err := databasetasksToTasks(tasks)

	return c.JSON(payload)
}
