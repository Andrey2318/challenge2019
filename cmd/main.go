package main

import (
	"challenge2019/cmd/cli"
	"challenge2019/internal/config"
	"challenge2019/internal/container"
	"challenge2019/internal/env"
	"challenge2019/internal/infrastructure/services/parser/csv"
	"challenge2019/internal/logger"
	"context"
	"os"
	"time"
)

func init() {
	time.Local, _ = time.LoadLocation("UTC")
	env.LoadEnvs()
	a, err := csv.ParsePartners(config.GetPartnersFileName())
	if err != nil {
		logger.Logger().WithError(err).Fatal()
	}
	if err := container.PartnerRepository().Save(context.Background(), a...); err != nil {
		logger.Logger().WithError(err).Fatal()
	}

	b, err := csv.ParseCapacities(config.GetCapacitiesFileName())
	if err != nil {
		logger.Logger().WithError(err).Fatal()
	}
	if err := container.CapacityRepository().Save(context.Background(), b...); err != nil {
		logger.Logger().WithError(err).Fatal()
	}
}

func main() {
	err := cli.App().Run(os.Args)
	if err != nil {
		logger.Logger().WithError(err).Fatal()
	}
}
