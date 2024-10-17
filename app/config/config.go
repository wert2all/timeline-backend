package config

type (
	Postgres struct {
		Port     int
		Host     string
		DB       string
		User     string
		Password string
	}

	Config struct {
		App struct {
			Cors struct {
				Debug         bool
				AllowedOrigin string
			}
		}

		Google struct {
			ClientID string
		}

		Postgres Postgres

		Sentry struct {
			Dsn string
		}
	}
)
