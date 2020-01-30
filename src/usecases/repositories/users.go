package repositories

type UserRepository interface {
	Select(*gorm.DB) *entites.User
}
