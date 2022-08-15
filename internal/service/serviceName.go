package service

import (
	"context"
	"github.com/ZQCard/kratos-crud-layout/internal/biz"
	"github.com/go-kratos/kratos/v2/log"

	v1 "github.com/ZQCard/kratos-crud-layout/api/serviceName/v1"
)

type ServiceNameService struct {
	v1.UnimplementedServiceNameServer
	serviceNameUseCase *biz.ServiceNameUseCase
	log                *log.Helper
}

func NewServiceNameService(ServiceNameUseCase *biz.ServiceNameUseCase,
	logger log.Logger) *ServiceNameService {
	return &ServiceNameService{
		log:                log.NewHelper(log.With(logger, "module", "service/interface")),
		serviceNameUseCase: ServiceNameUseCase,
	}
}

func (s *ServiceNameService) CreateServiceName(ctx context.Context, req *v1.CreateServiceNameRequest) (*v1.ServiceNameInfoResponse, error) {
	bc := &biz.ServiceName{
		Name: req.Name,
	}
	serviceNameInfo, err := s.serviceNameUseCase.Create(ctx, bc)
	if err != nil {
		return nil, err
	}
	return &v1.ServiceNameInfoResponse{
		Id:        serviceNameInfo.Id,
		Name:      serviceNameInfo.Name,
		CreatedAt: serviceNameInfo.CreatedAt,
		UpdatedAt: serviceNameInfo.UpdatedAt,
		DeletedAt: serviceNameInfo.DeletedAt,
	}, nil
}
func (s *ServiceNameService) UpdateServiceName(ctx context.Context, req *v1.UpdateServiceNameRequest) (*v1.ServiceNameInfoResponse, error) {
	bc := &biz.ServiceName{
		Id:   req.Id,
		Name: req.Name,
	}
	serviceNameInfo, err := s.serviceNameUseCase.Update(ctx, bc)
	if err != nil {
		return nil, err
	}
	return &v1.ServiceNameInfoResponse{
		Id:        serviceNameInfo.Id,
		Name:      serviceNameInfo.Name,
		CreatedAt: serviceNameInfo.CreatedAt,
		UpdatedAt: serviceNameInfo.UpdatedAt,
		DeletedAt: serviceNameInfo.DeletedAt,
	}, nil
}

func (s *ServiceNameService) DeleteServiceName(ctx context.Context, req *v1.DeleteServiceNameRequest) (*v1.ServiceNameCheckResponse, error) {
	err := s.serviceNameUseCase.Delete(ctx, req.Id)
	success := true
	if err != nil {
		success = false
	}
	return &v1.ServiceNameCheckResponse{
		IsSuccess: success,
	}, err
}

func (s *ServiceNameService) RecoverServiceName(ctx context.Context, req *v1.RecoverServiceNameRequest) (*v1.ServiceNameCheckResponse, error) {
	err := s.serviceNameUseCase.Recover(ctx, req.Id)
	success := true
	if err != nil {
		success = false
	}
	return &v1.ServiceNameCheckResponse{
		IsSuccess: success,
	}, err
}

func (s *ServiceNameService) GetServiceName(ctx context.Context, req *v1.GetServiceNameRequest) (*v1.ServiceNameInfoResponse, error) {
	params := map[string]interface{}{}
	params["id"] = req.Id
	params["name"] = req.Name
	serviceNameInfo, err := s.serviceNameUseCase.Get(ctx, params)
	if err != nil {
		return nil, err
	}
	response := &v1.ServiceNameInfoResponse{
		Id:        serviceNameInfo.Id,
		Name:      serviceNameInfo.Name,
		CreatedAt: serviceNameInfo.CreatedAt,
		UpdatedAt: serviceNameInfo.UpdatedAt,
		DeletedAt: serviceNameInfo.DeletedAt,
	}
	return response, nil
}

func (s *ServiceNameService) ListServiceName(ctx context.Context, req *v1.ListServiceNameRequest) (*v1.ListServiceNameReply, error) {
	params := make(map[string]interface{})
	params["name"] = req.Name
	ServiceNameInfoList, count, err := s.serviceNameUseCase.List(ctx, req.PageNum, req.PageSize, params)
	if err != nil {
		return nil, err
	}
	response := &v1.ListServiceNameReply{}
	response.Total = count
	for _, v := range ServiceNameInfoList {
		response.List = append(response.List, &v1.ServiceNameInfoResponse{
			Id:        v.Id,
			Name:      v.Name,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
			DeletedAt: v.DeletedAt,
		})
	}
	return response, nil
}
