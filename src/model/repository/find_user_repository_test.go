package repository

import (
	"fmt"
	"os"
	"testing"

	"github.com/Alan-Gomes1/go-api/src/model/repository/entity"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestUserRepository_FindUserByEmail(t *testing.T) {
	databaseName := "user_database_test"
	collectionName := "user_collection_test"

	os.Setenv(MONGODB_USER_COLLECTION, collectionName)
	defer os.Clearenv()

	mtestDB := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	mtestDB.Run("success", func(mt *mtest.T) {
		userEntity := entity.UserEntity{
			ID:       primitive.NewObjectID(),
			Email:    "test@example.com",
			Password: "test",
			Name:     "jhon",
			Age:      18,
		}
		mt.AddMockResponses(
			mtest.CreateCursorResponse(
				1,
				fmt.Sprintf("%s.%s", databaseName, collectionName),
				mtest.FirstBatch,
				convertEntityToBson(userEntity),
			),
		)

		databaseMock := mt.Client.Database(databaseName)
		repo := NewUserRepository(databaseMock)
		userDomain, err := repo.FindUserByEmail(userEntity.Email)

		assert.Nil(t, err)
		assert.EqualValues(t, userDomain.GetID(), userEntity.ID.Hex())
		assert.EqualValues(t, userDomain.GetEmail(), userEntity.Email)
		assert.EqualValues(t, userDomain.GetName(), userEntity.Name)
		assert.EqualValues(t, userDomain.GetAge(), userEntity.Age)
		assert.EqualValues(t, userDomain.GetPassword(), userEntity.Password)
	})
}

func convertEntityToBson(userEntity entity.UserEntity) bson.D {
	return bson.D{
		{Key: "_id", Value: userEntity.ID},
		{Key: "email", Value: userEntity.Email},
		{Key: "password", Value: userEntity.Password},
		{Key: "name", Value: userEntity.Name},
		{Key: "age", Value: userEntity.Age},
	}
}
