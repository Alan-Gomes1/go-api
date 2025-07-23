package view

import (
	"github.com/Alan-Gomes1/go-api/src/controller/model/response"
	"github.com/Alan-Gomes1/go-api/src/model"
)

func ConvertDomainToResponse(domain model.UserDomainInterface) response.UserResponse {
	return response.UserResponse{
		ID:    domain.GetID(),
		Email: domain.GetEmail(),
		Name:  domain.GetName(),
		Age:   domain.GetAge(),
	}
}
