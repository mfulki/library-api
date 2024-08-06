package env

type JwtConfig struct {
	UserKey  string
	AdminKey string
}

func (e *JwtConfig) Load() {
	e.UserKey = getEnv("JWT_USER_SECRET_KEY")
	e.AdminKey = getEnv("JWT_ADMIN_SECRET_KEY")

}
