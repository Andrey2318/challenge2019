package config

import "challenge2019/internal/env"

func GetLogLevel() string {
	return env.MustEnv("LOG_LEVEL")
}

func GetInputFile() string {
	return env.MustEnv("INPUT_FILE")
}

func GetPartnersFileName() string {
	return env.MustEnv("PARTNERS_FILE_NAME")
}

func GetCapacitiesFileName() string {
	return env.MustEnv("CAPACITIES_FILE_NAME")
}
