package order

import (
	"errors"
	"github.com/google/uuid"
	"github.com/pangolin-do-golang/tech-challenge-order-api/internal/errutil"
)

type Service struct {
	OrderRepository IOrderRepository
}

func NewOrderService(repo IOrderRepository) IOrderService {
	return &Service{
		OrderRepository: repo,
	}
}

func (s *Service) Get(id uuid.UUID) (*Order, error) {
	o, err := s.OrderRepository.Get(id)
	if err != nil {
		if errors.Is(err, errutil.ErrRecordNotFound) {
			return nil, errutil.NewBusinessError(err, "order not found")
		}

		return nil, err
	}

	return o, nil
}

func (s *Service) GetAll() ([]Order, error) {
	return s.OrderRepository.GetAll()
}

func (s *Service) Update(order *Order) (*Order, error) {
	o, err := s.OrderRepository.Get(order.ID)
	if err != nil {
		return nil, errutil.NewBusinessError(err, "order not found")
	}

	if err := o.ValidateStatusTransition(order.Status); err != nil {
		return nil, errutil.NewBusinessError(err, err.Error())
	}

	o.Status = order.Status
	err = s.OrderRepository.Update(o)
	if err != nil {
		return nil, err
	}
	oldOrder := *o

	// "simula" o período de uma tarefa async/em segundo plano pegar o
	// pedido "pago" e mudar o status para "preparando"1
	// dessa forma o usuário recebe o status "PAID"
	if o.Status == StatusPaid {
		o.Status = StatusPreparing
		if err := s.OrderRepository.Update(o); err != nil {
			return nil, err
		}
	}

	return &oldOrder, nil
}

func (s *Service) Create(clientID uuid.UUID) (*Order, error) {
	// TODO: implementar pegada aos microserviços
	/*
			c, err := s.CartService.GetFullCart(clientID)
		if err != nil {
			return nil, err
		}

		if len(c.Products) == 0 {
			return nil, fmt.Errorf("empty cart")
		}

		order := &Order{
			ClientID: clientID,
			Status:   StatusCreated,
		}

		o, err := s.OrderRepository.Create(order)
		if err != nil {
			return nil, err
		}

		var total float64
		for _, p := range c.Products {
			stockProduct, err := s.ProductService.GetByID(p.ProductID)
			if err != nil {
				return nil, err
			}

			productTotal := stockProduct.Price * float64(p.Quantity)

			orderProduct := &Product{
				ClientID:  clientID,
				ProductID: p.ProductID,
				Quantity:  p.Quantity,
				Comments:  p.Comments,
				Total:     productTotal,
			}

			err = s.OrderProductRepository.Create(context.Background(), o.ID, orderProduct)
			if err != nil {
				return nil, err
			}

			total += productTotal
		}

		o.TotalAmount = total
		o.Status = StatusPending
		err = s.OrderRepository.Update(o)
		if err != nil {
			return nil, err
		}

		if err = s.CartService.Cleanup(clientID); err != nil {
			return nil, err
		}

		return o, nil
	*/
	return &Order{}, nil
}
