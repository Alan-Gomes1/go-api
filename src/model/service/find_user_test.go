package service

import (
	"testing"

	"github.com/Alan-Gomes1/go-api/src/model"
	"github.com/Alan-Gomes1/go-api/src/tests/mocks"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/mock/gomock"
)

func TestFindUserByIDServices(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mocks.NewMockUserRepository(ctrl)
	service := NewUserDomainService(repository)
	t.Run("success", func(t *testing.T) {
		id := primitive.NewObjectID().Hex()
		userDomain := model.NewUserDomain("test@email.com", "test", "jhon", 18)
		userDomain.SetID(id)
		repository.EXPECT().FindUserByID(id).Return(userDomain, nil)
		user, err := service.FindUserByIDServices(id)

		assert.Nil(t, err)
		assert.Equal(t, userDomain, user)
	})
}
