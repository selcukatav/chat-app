package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/selcukatav/chat-app/model"
	"gorm.io/gorm"
)

// @Summary      Get user's friends
// @Description  Retrieve all friends of a user by their ID
// @Tags         Friends
// @Accept       json
// @Produce      json
// @Param        id   path  string  true  "User ID"
// @Success      200  {array}  model.Friend
// @Failure      400  {object}  map[string]string
// @Router       /api/users/{id}/friends [get]
func (h *Handler) GetFriends(c echo.Context) error {
	id := c.Param("id")

	var friends []model.Friend

	if err := h.DB.Where("user_id = ?", id).Find(&friends).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"Message": "There are no friends to get",
		})
	}
	return c.JSON(http.StatusOK, friends)
}

// @Summary      Find friends by username
// @Description  Retrieve friends of a user by their username
// @Tags         Friends
// @Accept       json
// @Produce      json
// @Param        username   path  string  true  "Username"
// @Success      200  {array}  model.Friend
// @Failure      400  {object}  map[string]string
// @Router       /api/users/{username}/friends [get]
func (h *Handler) FindFriends(c echo.Context) error {
	username := c.Param("username")

	var friends []model.Friend

	if err := h.DB.Where("username = ?", username).Find(&friends).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"Message": "There are no friends to get",
		})
	}
	return c.JSON(http.StatusOK, friends)
}

// @Summary      Delete a friend
// @Description  Remove a friend by their Friend ID
// @Tags         Friends
// @Accept       json
// @Produce      json
// @Param        user_id   path  string  true  "User ID"
// @Param        friend_id   path  string  true  "Friend ID"
// @Success      200  {object}  map[string]string
// @Failure      400  {object}  map[string]string
// @Router       /api/friends [delete]
func (h *Handler) DeleteFriend(c echo.Context) error {
	
	var friend model.Friend
	if err := c.Bind(&friend); err != nil {
		return echo.ErrBadRequest
	}

	user_id, friend_id := friend.UserID, friend.FriendID
	if err := h.DB.Where("user_id=? and friend_id=?",user_id,friend_id).First(&friend).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"Message": "Friend not found",
		})
	}

	if err := h.DB.Delete(&friend).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"Message": "Friend could not be deleted",
		})
	}
	return c.JSON(http.StatusOK, map[string]string{
		"Message": "Friend deleted successfully!",
	})
}

// @Summary      Add a friend
// @Description  Add friends with each other
// @Tags         Friends
// @Accept       json
// @Produce      json
// @Param        user_id   path  string  true  "User ID"
// @Param        friend_id   path  string  true  "Friend ID"
// @Success      200  {object}  map[string]string
// @Failure      400  {object}  map[string]string
// @Router       /api/friends [post]
func (h *Handler) AddFriend(c echo.Context) error {

	var friend model.Friend

	if err := c.Bind(&friend); err != nil {
		return echo.ErrBadRequest
	}

	user_id, friend_id := friend.UserID, friend.FriendID
	result := h.DB.Where("user_id=? and friend_id=?", user_id, friend_id).First(&friend)
	if result.RowsAffected > 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"Message": "You are already friend with him/her!!",
		})
	} else if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"Message": "Something went wrong",
		})
	}

	if err := h.DB.Create(&friend).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"Message": "Error occured while adding friend",
		})
	}

	return c.JSON(http.StatusOK, friend)
}
