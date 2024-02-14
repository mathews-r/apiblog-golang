package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mathews-r/golang/src/configs/logger"
	"github.com/mathews-r/golang/src/configs/rest_err"
	"github.com/mathews-r/golang/src/configs/validation"
	"github.com/mathews-r/golang/src/controller/model/request"
	"github.com/mathews-r/golang/src/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (pc *postControllerInterface) UpdatePost(c *gin.Context) {
	var postRequest request.PostUpdateRequest

	if err := c.ShouldBindJSON(&postRequest); err != nil {
		restErr := validation.ValidateUserError(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	postId := c.Param("postId")
	if _, err := primitive.ObjectIDFromHex(postId); err != nil {
		errRest := rest_err.NewBadRequestErr("Invalid postId, must be hex")
		c.JSON(errRest.Code, errRest)
	}

	domain := model.UpdatePostDomain(
		postRequest.Title,
		postRequest.Content,
	)

	err := pc.service.UpdatePost(postId, domain)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	logger.Info("Post updated successfully")
	c.Status(http.StatusOK)
}