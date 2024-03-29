package view

import (
	"github.com/mathews-r/golang/src/controller/model/response"
	"github.com/mathews-r/golang/src/model"
)

func ConvertDomainToResponse(
	userDomain model.UserDomainInterface,
) response.UserResponse {
	return response.UserResponse{
		ID:    userDomain.GetId(),
		Name:  userDomain.GetName(),
		Email: userDomain.GetEmail(),
		Age:   userDomain.GetAge(),
	}
}

func ConvertDomainToResponsePost(
	postDomain model.PostDomainInterface,
) response.PostResponse {
	return response.PostResponse{
		ID:        postDomain.GetPostId(),
		Title:     postDomain.GetTitle(),
		Content:   postDomain.GetContent(),
		Category:  postDomain.GetCategory(),
		UserId:    postDomain.GetUserId(),
		Published: postDomain.GetPublished(),
		Updated:   postDomain.GetUpdated(),
	}
}

func ConvertDomainToResponseCategory(
	categoryDomain model.CategoryDomainInterface,
) response.CategoryResponse {
	return response.CategoryResponse{
		ID:       categoryDomain.GetId(),
		Category: categoryDomain.GetCategory(),
	}
}
