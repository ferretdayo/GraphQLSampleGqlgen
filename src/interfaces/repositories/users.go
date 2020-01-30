package repositories

import "github.com/GraphQLSample/src/entities"

type UserRepository struct {
}

func (repository *UserRepository) Select(db *gorm.DB) *entities.User {
	user := &entities.User{
		ID:       134,
		NickName: "AAAA",
		Old:      25,
	}
	return user
}
