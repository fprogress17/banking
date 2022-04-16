package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{Id: "10001", Name: "Abc", City: "New Delhi", Zipcode: "10001", DateofBirth: "2001", Status: "good"},
		{Id: "10002", Name: "Bbc", City: "New Delhi", Zipcode: "10001", DateofBirth: "2001", Status: "good"},
	}

	return CustomerRepositoryStub{customers}
}
