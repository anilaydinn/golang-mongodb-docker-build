package main

type API struct {
	service *Service
}

func NewAPI(service *Service) API {
	return API{
		service: service,
	}
}

type DataDTO struct {
	Text string `json:"text"`
}
