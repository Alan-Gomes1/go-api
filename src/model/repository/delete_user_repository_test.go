package repository

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestUserRepository_DeleteUser(t *testing.T) {
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
		err := repo.DeleteUser("abc123")

		assert.Nil(t, err)
	})

	mtestDB.Run("error", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{{Key: "ok", Value: 0}})

		databaseMock := mt.Client.Database(databaseName)
		repo := NewUserRepository(databaseMock)
		err := repo.DeleteUser("abc123")

		assert.NotNil(t, err)
	})
}
