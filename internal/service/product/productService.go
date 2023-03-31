package product

import "fmt"

type productService interface {
	Get(productID int64) (*Product, error)
	List(cursor, limit int64) (*[]Product, error)
	Create(product Product) (int64, error)
	Update(productID int64, product Product) error
	Remove(productID int64) (bool, error)
}

type Service struct {
	products []Product
}

func NewService() *Service {
	return &Service{
		products: allProducts,
	}
}

func (s *Service) Get(productID int64) (*Product, error) {
	if productID < 0 || productID > int64(len(s.products)-1) {
		return nil, fmt.Errorf("incorrect index")
	}
	result := s.products[productID]
	return &result, nil
}

func (s *Service) List(cursor, limit int64) (*[]Product, error) {
	if len(s.products) == 0 {
		return nil, fmt.Errorf("len = 0")
	}
	return &s.products, nil
}

func (s *Service) Create(product *Product) (int64, error) {
	s.products = append(s.products, *product)
	return int64(len(s.products) - 1), nil
}

func (s *Service) Update(productID int64, product *Product) error {
	if productID < 0 || productID > int64(len(s.products)-1) {
		return fmt.Errorf("incorrect index")
	}
	s.products[productID] = *product
	return nil
}

func (s *Service) Remove(productID int64) (bool, error) {
	if productID < 0 || productID > int64(len(s.products)-1) {
		return false, fmt.Errorf("incorrect index")
	}

	s.products = append(s.products[:productID], s.products[productID+1:]...)
	return true, nil
}
