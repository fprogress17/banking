package service

import "github.com/fprogress17/banking/domain"

type CustomerService interface {
  GetAllCustomer() ([]domain.Customer, error)
}

type DefaultCustomerService struct {
  repo domain.CustomerRepository
}

func (s DefaultCustomerService) GellAllCustomer() ([]domain.Customer, error) {
  return s.repo.FindAll()
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
  return DefaultCustomerService(repository)
}
