package di

type Postgres struct {
	Port     string `yaml:"port"`
	Host     string `yaml:"host"`
	Db       string `yaml:"db"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

type Config struct {
	App struct {
		Cors struct {
			Debug         bool   `yaml:"debug"`
			AllowedOrigin string `yaml:"allowedOrigin"`
		} `yaml:"cors"`
	} `yaml:"app"`

	Google struct {
		ClientId string `yaml:"clientId"`
	} `yaml:"google"`

	Postgres Postgres `yaml:"postgres"`

	Sentry struct {
		Dsn string `yaml:"dsn"`
	} `yaml:"sentry"`
}
