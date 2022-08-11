package service

import (
	"github.com/dannielss/banking/domain"
	"github.com/dannielss/banking/dto"
	"github.com/dannielss/banking/errs"
)

type CustomerService interface {
	GetAllCustomers(status string) ([]dto.CustomerResponse, *errs.AppError)
	GetCustomer(string) (*dto.CustomerResponse, *errs.AppError)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomers(status string) ([]dto.CustomerResponse, *errs.AppError) {
	if status == "active" {
		status = "1"
	} else if status == "inactive" {
		status = "0"
	} else {
		status = ""
	}

	c, err := s.repo.FindAll(status)

	if err != nil {
		return nil, err
	}

	response := make([]dto.CustomerResponse, len(c))

	for i, customer := range c {
		response[i] = customer.ToDto()
	}

	return response, nil
}

func (s DefaultCustomerService) GetCustomer(id string) (*dto.CustomerResponse, *errs.AppError) {
	c, err := s.repo.ById(id)

	if err != nil {
		return nil, err
	}

	response := c.ToDto()

	return &response, nil
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{
		repo: repository,
	}
}
