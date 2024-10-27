package environment

const (
	EnvDev  = "development"
	EnvLive = "production"
)

type Environment struct {
	Env string
}

var envObs Environment

func InitEnvironment(env string) {
	envObs.Env = env
}

func GetEnv() string {
	if envObs.Env == "" {
		return EnvDev
	}

	return envObs.Env
}
