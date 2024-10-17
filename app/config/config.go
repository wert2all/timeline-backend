package config

type (
	Postgres struct {
		Port     int
		Host     string
		DB       string
		User     string
		Password string
	}

	Cors struct {
		Debug         bool
		AllowedOrigin string
	}

	App struct {
		Cors Cors
	}

	Google struct {
		ClientID string
	}

	Sentry struct {
		Dsn string
	}

	Config struct {
		App      App
		Postgres Postgres

		Google Google
		Sentry Sentry
	}
)
