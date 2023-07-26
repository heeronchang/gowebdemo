package main

import (
	"gowebdemo/configs/appone"
	"gowebdemo/internal/app/appone/routes"
	_ "gowebdemo/internal/pkg/config_zerolog"

	"github.com/rs/zerolog/log"
)

func main() {
	des := appone.GetDesc()
	log.Info().Msgf("app decs: %s", des)

	routes.StartWebAPI()
}
