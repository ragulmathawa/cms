package migrations

import (
	"os"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(db dbx.Builder) error {

		// Get Configs from ENV

		envAppName := os.Getenv("APP_NAME")
		envAppUrl := os.Getenv("APP_URL")

		// add up queries...
		dao := daos.New(db)

		settings, _ := dao.FindSettings()
		if envAppName != "" {
			settings.Meta.AppName = envAppName
		}
		if envAppUrl != "" {
			settings.Meta.AppUrl = envAppUrl
		}

		return dao.SaveSettings(settings)
	}, nil)
}
