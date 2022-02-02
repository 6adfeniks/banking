package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{"1001", "Kolya", "Minsk", "222160", "2003-05-22", "1"},
		{"1002", "Vanya", "Zhodino", "222167", "2005-07-07", "1"},
	}
	return CustomerRepositoryStub{customers: customers}
}
