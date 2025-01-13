package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/selcukatav/chat-app/model"
)

// @Summary      Update user
// @Description  Update a user's information by their ID
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        id    path      string            true  "User ID"
// @Param        user  body      model.User        true  "User data"
// @Success      200   {object}  model.User
// @Failure      400   {object}  map[string]string
// @Router       /api/users/{id} [put]
func (h *Handler) UpdateUser(c echo.Context) error {
	id := c.Param("id")
	var user model.User

	if err := h.DB.First(&user, id).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"Message": "User not found",
		})
	}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"Message": "Invalid input",
		})
	}
	if err := h.DB.Model(&user).Updates(user).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"Message": "User could not be updated",
		})
	}

	return c.JSON(http.StatusOK, user)
}

// @Summary      Find a user
// @Description  Retrieve a user by their ID
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        id   path      string            true  "User ID"
// @Success      200  {object}  model.User
// @Failure      400  {object}  map[string]string
// @Router       /api/users/{id} [get]
func (h *Handler) GetUser(c echo.Context) error {
	id := c.Param("id")
	var user model.User

	if err := h.DB.First(&user, id).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"Message": "User not found",
		})
	}
	return c.JSON(http.StatusOK, user)
}

// @Summary      List all users
// @Description  Retrieve a list of all users
// @Tags         Users
// @Accept       json
// @Produce      json
// @Success      200  {array}   model.User
// @Failure      400  {object}  map[string]string
// @Router       /api/users [get]
func (h *Handler) ListUsers(c echo.Context) error {

	var users []model.User
	if err := h.DB.Find(&users).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"Message": "Users not found",
		})
	}

	return c.JSON(http.StatusOK, users)
}

// @Summary      Delete user
// @Description  Delete a user by their ID
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        id   path      string            true  "User ID"
// @Success      200  {object}  map[string]string
// @Failure      400  {object}  map[string]string
// @Router       /api/users/{id} [delete]
func (h *Handler) DeleteUser(c echo.Context) error {
	id := c.Param("id")
	var user model.User

	if err := h.DB.First(&user, id).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"Message": "User not found",
		})
	}
	if err := h.DB.Delete(&user).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"Message": "User could not be deleted",
		})
	}
	return c.JSON(http.StatusOK, map[string]string{
		"Message": "User deleted successfully!",
	})
}
