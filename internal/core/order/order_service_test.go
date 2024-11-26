package order_test

import (
	"errors"
	"github.com/google/uuid"
	"github.com/pangolin-do-golang/tech-challenge-order-api/internal/core/order"
	"github.com/pangolin-do-golang/tech-challenge-order-api/internal/errutil"
	"github.com/pangolin-do-golang/tech-challenge-order-api/mocks"
	"github.com/stretchr/testify/mock"
	"reflect"
	"testing"
	"time"
)

func TestService_Get(t *testing.T) {
	var (
		id = uuid.MustParse("9a29c14a-72d3-4e15-8adc-a8ed9c845fe2")
		o  = &order.Order{
			ID: uuid.MustParse("9a29c14a-72d3-4e15-8adc-a8ed9c845fe2"),
		}
	)

	type fields struct {
		genOrderRepository func() order.IOrderRepository
	}
	type args struct {
		id uuid.UUID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *order.Order
		wantErr bool
	}{
		{
			name: "returns error from order repository",
			fields: fields{
				func() order.IOrderRepository {
					m := new(mocks.IOrderRepository)
					m.On("Get", id).Return(nil, errors.New("error"))
					return m
				},
			},
			args: args{
				id: id,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "returns error for order not found from order repository",
			fields: fields{
				func() order.IOrderRepository {
					m := new(mocks.IOrderRepository)
					m.On("Get", id).Return(nil, errutil.ErrRecordNotFound)
					return m
				},
			},
			args: args{
				id: id,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "returns order from repository",
			fields: fields{
				func() order.IOrderRepository {
					m := new(mocks.IOrderRepository)
					m.On("Get", id).Return(o, nil)
					return m
				},
			},
			args: args{
				id: id,
			},
			want:    o,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &order.Service{
				OrderRepository: tt.fields.genOrderRepository(),
			}
			got, err := s.Get(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_GetAll(t *testing.T) {
	orders := []order.Order{
		{
			ID:          uuid.UUID{},
			CreatedAt:   time.Time{},
			ClientID:    uuid.UUID{},
			TotalAmount: 0,
			Status:      "",
		},
		{
			ID:          uuid.UUID{},
			CreatedAt:   time.Time{},
			ClientID:    uuid.UUID{},
			TotalAmount: 0,
			Status:      "",
		},
	}

	type fields struct {
		genOrderRepository func() order.IOrderRepository
	}
	tests := []struct {
		name    string
		fields  fields
		want    []order.Order
		wantErr bool
	}{
		{
			name: "returns error from repository",
			fields: fields{
				func() order.IOrderRepository {
					m := new(mocks.IOrderRepository)
					m.On("GetAll").Return(nil, errors.New("error"))
					return m
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "returns result from repository",
			fields: fields{
				func() order.IOrderRepository {
					m := new(mocks.IOrderRepository)
					m.On("GetAll").Return(orders, nil)
					return m
				},
			},
			want:    orders,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &order.Service{
				OrderRepository: tt.fields.genOrderRepository(),
			}
			got, err := s.GetAll()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAll() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_Update(t *testing.T) {
	var (
		o = &order.Order{
			ID: uuid.MustParse("9a29c14a-72d3-4e15-8adc-a8ed9c845fe2"),
		}
	)

	type fields struct {
		genOrderRepository func() order.IOrderRepository
	}
	type args struct {
		order *order.Order
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *order.Order
		wantErr bool
	}{
		{
			name: "returns error from repository get method",
			fields: fields{
				func() order.IOrderRepository {
					m := new(mocks.IOrderRepository)
					m.On("Get", o.ID).Return(nil, errors.New("error"))
					return m
				},
			},
			args: args{
				order: o,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "blocks transition from status created to paid",
			fields: fields{
				func() order.IOrderRepository {
					m := new(mocks.IOrderRepository)
					m.On("Get", mock.Anything).Return(&order.Order{
						Status: order.StatusCreated,
					}, nil)
					return m
				},
			},
			args: args{
				order: &order.Order{
					Status: order.StatusPaid,
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "blocks transition from status created to paid",
			fields: fields{
				func() order.IOrderRepository {
					m := new(mocks.IOrderRepository)
					m.On("Get", mock.Anything).Return(&order.Order{
						Status: order.StatusCreated,
					}, nil)
					return m
				},
			},
			args: args{
				order: &order.Order{
					Status: order.StatusPaid,
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "accepts transition from status created to pending",
			fields: fields{
				func() order.IOrderRepository {
					m := new(mocks.IOrderRepository)
					m.On("Get", mock.Anything).Return(&order.Order{
						Status: order.StatusCreated,
					}, nil)
					m.On("Update", mock.Anything).Return(nil)
					return m
				},
			},
			args: args{
				order: &order.Order{
					Status: order.StatusPending,
				},
			},
			want: &order.Order{
				Status: order.StatusPending,
			},
			wantErr: false,
		},
		{
			name: "blocks transition from status created to paid",
			fields: fields{
				func() order.IOrderRepository {
					m := new(mocks.IOrderRepository)
					m.On("Get", mock.Anything).Return(&order.Order{
						Status: order.StatusPending,
					}, nil)
					return m
				},
			},
			args: args{
				order: &order.Order{
					Status: order.StatusCreated,
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "accepts transition from status pending to paid",
			fields: fields{
				func() order.IOrderRepository {
					m := new(mocks.IOrderRepository)
					m.On("Get", mock.Anything).Return(&order.Order{
						Status: order.StatusPending,
					}, nil)
					m.On("Update", mock.Anything).Return(nil)
					return m
				},
			},
			args: args{
				order: &order.Order{
					Status: order.StatusPaid,
				},
			},
			want: &order.Order{
				Status: order.StatusPaid,
			},
			wantErr: false,
		},
		{
			name: "accepts transition from status pending to declined",
			fields: fields{
				func() order.IOrderRepository {
					m := new(mocks.IOrderRepository)
					m.On("Get", mock.Anything).Return(&order.Order{
						Status: order.StatusPending,
					}, nil)
					m.On("Update", mock.Anything).Return(nil)
					return m
				},
			},
			args: args{
				order: &order.Order{
					Status: order.StatusDeclined,
				},
			},
			want: &order.Order{
				Status: order.StatusDeclined,
			},
			wantErr: false,
		},
		{
			name: "blocks transition from status paid to preparing",
			fields: fields{
				func() order.IOrderRepository {
					m := new(mocks.IOrderRepository)
					m.On("Get", mock.Anything).Return(&order.Order{
						Status: order.StatusPaid,
					}, nil)
					return m
				},
			},
			args: args{
				order: &order.Order{
					Status: order.StatusCreated,
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "accepts transition from status paid to preparing",
			fields: fields{
				func() order.IOrderRepository {
					m := new(mocks.IOrderRepository)
					m.On("Get", mock.Anything).Return(&order.Order{
						Status: order.StatusPaid,
					}, nil)
					m.On("Update", mock.Anything).Return(nil)
					return m
				},
			},
			args: args{
				order: &order.Order{
					Status: order.StatusPreparing,
				},
			},
			want: &order.Order{
				Status: order.StatusPreparing,
			},
			wantErr: false,
		},
		{
			name: "accepts transition from status preparing to finished",
			fields: fields{
				func() order.IOrderRepository {
					m := new(mocks.IOrderRepository)
					m.On("Get", mock.Anything).Return(&order.Order{
						Status: order.StatusPaid,
					}, nil)
					m.On("Update", mock.Anything).Return(nil)
					return m
				},
			},
			args: args{
				order: &order.Order{
					Status: order.StatusPreparing,
				},
			},
			want: &order.Order{
				Status: order.StatusPreparing,
			},
			wantErr: false,
		},
		{
			name: "blocks transition from status canceled to any other",
			fields: fields{
				func() order.IOrderRepository {
					m := new(mocks.IOrderRepository)
					m.On("Get", mock.Anything).Return(&order.Order{
						Status: order.StatusCanceled,
					}, nil)
					return m
				},
			},
			args: args{
				order: &order.Order{
					Status: order.StatusCreated,
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "blocks transition from status preparing to any paid",
			fields: fields{
				func() order.IOrderRepository {
					m := new(mocks.IOrderRepository)
					m.On("Get", mock.Anything).Return(&order.Order{
						Status: order.StatusPreparing,
					}, nil)
					return m
				},
			},
			args: args{
				order: &order.Order{
					Status: order.StatusPaid,
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "blocks transition for invalid status",
			fields: fields{
				func() order.IOrderRepository {
					m := new(mocks.IOrderRepository)
					m.On("Get", mock.Anything).Return(&order.Order{
						Status: order.StatusPreparing,
					}, nil)
					return m
				},
			},
			args: args{
				order: &order.Order{
					Status: "anyother",
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "returns error from order update",
			fields: fields{
				func() order.IOrderRepository {
					m := new(mocks.IOrderRepository)
					m.On("Get", mock.Anything).Return(&order.Order{
						Status: order.StatusPaid,
					}, nil)
					m.On("Update", mock.Anything).Return(errors.New("error"))
					return m
				},
			},
			args: args{
				order: &order.Order{
					Status: order.StatusPreparing,
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "sends order to preparing if status is paid",
			fields: fields{
				func() order.IOrderRepository {
					m := new(mocks.IOrderRepository)
					m.On("Get", mock.Anything).Return(&order.Order{
						Status: order.StatusPending,
					}, nil)
					m.On("Update", &order.Order{
						Status: order.StatusPaid,
					}).Return(nil)

					m.On("Update", &order.Order{
						Status: order.StatusPreparing,
					}).Return(nil)

					return m
				},
			},
			args: args{
				order: &order.Order{
					Status: order.StatusPaid,
				},
			},
			want: &order.Order{
				Status: order.StatusPaid,
			},
			wantErr: false,
		},
		{
			name: "returns error from second order update",
			fields: fields{
				func() order.IOrderRepository {
					m := new(mocks.IOrderRepository)
					m.On("Get", mock.Anything).Return(&order.Order{
						Status: order.StatusPending,
					}, nil)
					m.On("Update", &order.Order{
						Status: order.StatusPaid,
					}).Return(nil)

					m.On("Update", &order.Order{
						Status: order.StatusPreparing,
					}).Return(errors.New("error"))

					return m
				},
			},
			args: args{
				order: &order.Order{
					Status: order.StatusPaid,
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &order.Service{
				OrderRepository: tt.fields.genOrderRepository(),
			}
			got, err := s.Update(tt.args.order)
			if (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Update() got = %v, want %v", got, tt.want)
			}
		})
	}
}
