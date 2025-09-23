package repository

import (
	"os"
	"testing"

	"github.com/Alan-Gomes1/go-api/src/model"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestUserRepository_CreateUser(t *testing.T) {
	databaseName := "user_database_test"
	collectionName := "user_collection_test"

	os.Setenv(MONGODB_USER_COLLECTION, collectionName)
	defer os.Clearenv()

	mtestDB := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mtestDB.Run("success", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 1},
			{Key: "n", Value: 1},
			{Key: "acknowledged", Value: true},
		})

		databaseMock := mt.Client.Database(databaseName)
		repo := NewUserRepository(databaseMock)
		domain := model.NewUserDomain("test@example.com", "test", "jhon", 18)
		userDomain, err := repo.CreateUser(domain)
		_, errId := primitive.ObjectIDFromHex(userDomain.GetID())

		assert.Nil(t, err)
		assert.Nil(t, errId)
		assert.EqualValues(t, userDomain.GetEmail(), domain.GetEmail())
		assert.EqualValues(t, userDomain.GetName(), domain.GetName())
		assert.EqualValues(t, userDomain.GetAge(), domain.GetAge())
		assert.EqualValues(t, userDomain.GetPassword(), domain.GetPassword())
	})
}
