package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{Id: "1001", Name: "Daniel", City: "Guarulhos", Zipcode: "08162310", DateofBirth: "2000-01-04", Status: "1"},
		{Id: "1002", Name: "Pedro", City: "Buenos Aires", Zipcode: "02162310", DateofBirth: "2000-03-04", Status: "1"},
	}

	return CustomerRepositoryStub{
		customers: customers,
	}
}
