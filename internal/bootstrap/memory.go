package bootstrap

import (
	"fmt"

	"github.com/Seasky89/go-report-api/internal/external"
	"github.com/Seasky89/go-report-api/internal/models"
	"github.com/Seasky89/go-report-api/internal/store"
)

func InitMemoryStore() store.DataStore {
	users, err := external.GetJSON[models.User]("https://jsonplaceholder.typicode.com/users")
	if err != nil {
		fmt.Println("failed to load users: ", err)
		users = []models.User{}
	}

	posts, err := external.GetJSON[models.Post]("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		fmt.Println("failed to load posts: ", err)
		posts = []models.Post{}
	}

	return store.NewMemoryStore(users, posts)
}
