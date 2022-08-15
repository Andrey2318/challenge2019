package container

import (
	"challenge2019/internal/domain/capacity"
	"challenge2019/internal/domain/partner"
	capacityRepository "challenge2019/internal/infrastructure/store/csv/capacity"
	partnerRepository "challenge2019/internal/infrastructure/store/csv/partner"
	"github.com/patrickmn/go-cache"
)

var database *cache.Cache

func Database() *cache.Cache {
	if database != nil {
		return database
	}

	database = cache.New(cache.NoExpiration, 0)

	return database
}

var partnerRep partner.Repository

func PartnerRepository() partner.Repository {
	if partnerRep != nil {
		return partnerRep
	}

	partnerRep = partnerRepository.New(Database())

	return partnerRep
}

var capacityRep capacity.Repository

func CapacityRepository() capacity.Repository {
	if capacityRep != nil {
		return capacityRep
	}

	capacityRep = capacityRepository.New(Database())

	return capacityRep
}
