package container

import "challenge2019/internal/application/traffic"

var trafficApp *traffic.UseCase

func TrafficApplication() *traffic.UseCase {
	if trafficApp != nil {
		return trafficApp
	}

	trafficApp = traffic.New(PartnerRepository(), CapacityRepository())

	return trafficApp
}
