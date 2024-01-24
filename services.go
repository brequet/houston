package main

import (
	"log"

	"golang.org/x/sys/windows/svc/mgr"
)

type Service struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      uint32 `json:"status"`
	ProcessId   uint32 `json:"processId"`
}

func GetAllServices() ([]Service, error) {
	m, err := mgr.Connect()
	if err != nil {
		log.Println("1")
		return nil, err
	}
	defer m.Disconnect()

	serviceNames, err := m.ListServices()
	if err != nil {
		log.Println("2")
		return nil, err
	}

	var result []Service
	log.Println("Services -- ", serviceNames)
	for _, serviceName := range serviceNames {
		log.Println("Service ", serviceName)
		var status, processId uint32
		var description string
		service, err := m.OpenService(serviceName)
		if err != nil {
			log.Println("3")
			// return nil, err
		} else {
			serviceStatus, err := service.Query()
			if err != nil {
				log.Println("4")
				// return nil, err
			} else {
				status = uint32(serviceStatus.State)
				processId = serviceStatus.ProcessId
			}

			serviceConfig, err := service.Config()
			if err != nil {
				log.Println("5")
				// return nil, err
			} else {
				description = serviceConfig.Description
			}
		}

		result = append(result, Service{
			Name:        serviceName,
			Description: description,
			Status:      status,
			ProcessId:   processId,
		})
	}

	return result, nil
}
