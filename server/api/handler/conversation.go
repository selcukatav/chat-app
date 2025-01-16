package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/selcukatav/chat-app/model"
)

// @Summary      Create a new conversation
// @Description  Create a new conversation with the given details
// @Tags         Conversations
// @Accept       json
// @Produce      json
// @Param        conversation  body      model.Conversation  true  "Conversation details"
// @Success      200           {object}  map[string]string
// @Failure      400           {object}  map[string]string
// @Failure      500           {object}  map[string]string
// @Router       /api/conversations [post]
func (h *Handler) CreateConversation(c echo.Context) error {
	var subject model.Conversation

	if err := c.Bind(&subject); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"Message": "Error occurred while taking the subject object!",
		})
	}
	if err := h.DB.Create(&subject).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"Message": "Error occurred while creating the conversation!",
		})
	}

	return c.JSON(http.StatusOK, subject)
}

// @Summary      List conversations
// @Description  Retrieve all conversations
// @Tags         Conversations
// @Accept       json
// @Produce      json
// @Param        user_id  path      string  true  "User ID"
// @Success      200      {array}   model.Conversation
// @Failure      400      {object}  map[string]string
// @Router       /api/conversations [get]
func (h *Handler) ListConversations(c echo.Context) error {
	var conversations []model.Conversation

	if err := h.DB.Find(&conversations).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"Message": "Error occurred while listing conversations",
		})
	}

	return c.JSON(http.StatusOK, conversations)
}

// @Summary      Add a participant to a conversation
// @Description  Add a user to an existing conversation
// @Tags         Conversations
// @Accept       json
// @Produce      json
// @Param        participant  body      model.ConversationParticipant  true  "Participant details"
// @Success      200          {object}  model.ConversationParticipant
// @Failure      400          {object}  map[string]string
// @Failure      500          {object}  map[string]string
// @Router       /api/conversations/participants [post]
func (h *Handler) AddConversationParticipant(c echo.Context) error {
	var conversationParticipant model.ConversationParticipant

	if err := c.Bind(&conversationParticipant); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"Message": "Error occurred while taking the subject object!",
		})
	}

	userID, conversationID := conversationParticipant.UserID, conversationParticipant.ConversationID

	result := h.DB.Where("user_id=? and conversation_id=?", userID, conversationID).First(&conversationParticipant)
	if result.RowsAffected > 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"Message": "User is already in the conversation!",
		})
	}

	if err := h.DB.Create(&conversationParticipant).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"Message": "Error occurred while adding user to conversation!",
		})
	}

	return c.JSON(http.StatusOK, conversationParticipant)
}

// @Summary      Remove a participant from a conversation
// @Description  Delete a user from an existing conversation
// @Tags         Conversations
// @Accept       json
// @Produce      json
// @Param        participant  body      model.ConversationParticipant  true  "Participant details"
// @Success      200          {object}  map[string]string
// @Failure      400          {object}  map[string]string
// @Failure      500          {object}  map[string]string
// @Router       /api/conversations/participants [delete]
func (h *Handler) DeleteConversationParticipant(c echo.Context) error {
	var conversationParticipant model.ConversationParticipant

	if err := c.Bind(&conversationParticipant); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"Message": "Error occurred while taking the subject object!",
		})
	}

	userID, conversationID := conversationParticipant.UserID, conversationParticipant.ConversationID

	if err := h.DB.Where("user_id=? and conversation_id=?", userID, conversationID).First(&conversationParticipant).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"Message": "Couldn't find the participant",
		})
	}

	if err := h.DB.Delete(&conversationParticipant).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"Message": "Error occurred while removing user from conversation!",
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"Message": "User successfully removed from conversation!",
	})
}

// @Summary      List conversation participants by conversation ID
// @Description  Retrieve all users associated with a specific conversation
// @Tags         Conversations
// @Accept       json
// @Produce      json
// @Param        conversation_id  path      string  true  "Conversation ID"
// @Success      200      {array}   model.Conversation
// @Failure      400      {object}  map[string]string
// @Router       /api/conversations/{conversation_id} [get]
func (h *Handler) ListConversationsParticipants(c echo.Context) error {
	var conservationID = c.Param("conversation_id")
	var conversationsParticipants []model.ConversationParticipant

	if err := h.DB.Where("conversation_id=?", conservationID).Find(&conversationsParticipants).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"Message": "Error occurred while listing conversations",
		})
	}

	return c.JSON(http.StatusOK, conversationsParticipants)
}

// @Summary      List conversations by user ID
// @Description  Retrieve all conversations associated with a specific user
// @Tags         Conversations
// @Accept       json
// @Produce      json
// @Param        user_id  path      string  true  "User ID"
// @Success      200      {array}   model.Conversation
// @Failure      400      {object}  map[string]string
// @Router       /api/conversations/{user_id} [get]
func (h *Handler) ListUserConversations(c echo.Context) error {
	var userID = c.Param("user_id")
	var conversations []model.ConversationParticipant

	if err := h.DB.Where("user_id=?", userID).Find(&conversations).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"Message": "Error occurred while listing conversations",
		})
	}

	return c.JSON(http.StatusOK, conversations)
}

func (h *Handler) ConversationRoom(c echo.Context) error {

	conn, err := upgrader.Upgrade(c.Response(), c.Request(), nil)

	if err != nil {
		fmt.Println("websocket connection is not established: ", err)
		return err
	}

	conversationID:=c.QueryParam("conversation_id")
	if conversationID==""{
		return c.JSON(http.StatusBadRequest,map[string]string{
			"Message":"conversation_id is required",
		})
	}
	if err:=h.ChatAction(conn, conversationID); err!=nil{
		fmt.Println("Chat action error: ",err)
		return err
	}
	return nil

}
