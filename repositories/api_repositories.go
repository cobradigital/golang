package repositories

var User UserRepository

func Init() {
	Config()
	User = initUserRepository()
}
