package consulsd

import (
	"github.com/davyxu/cellmesh/discovery"
	"github.com/hashicorp/consul/api"
)

// check包括：node的check和service check
func isMeshServiceHealth(entry *api.ServiceEntry) bool {

	for _, check := range entry.Checks {
		if check.ServiceID == entry.Service.ID &&
			check.Output == makeHealthWords(entry.Service.ID) {
			return true
		}
	}

	return false
}

func consulSvcToService(s *api.ServiceEntry) *discovery.ServiceDesc {

	return &discovery.ServiceDesc{
		Name: s.Service.Service,
		ID:   s.Service.ID,
		Host: s.Service.Address,
		Port: s.Service.Port,
		Tags: s.Service.Tags,
	}
}

func existsInServiceList(svclist []*discovery.ServiceDesc, id string) bool {
	for _, svc := range svclist {

		if svc.ID == id {
			return true
		}
	}

	return false
}
