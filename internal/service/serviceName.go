package service

import (
	"context"
	"github.com/ZQCard/kratos-crud-layout/internal/biz"
	"github.com/go-kratos/kratos/v2/log"

	v1 "github.com/ZQCard/kratos-crud-layout/api/serviceName/v1"
)

type ServiceNameService struct {
	v1.UnimplementedServiceNameServer
	moduleUc *biz.ServiceNameUseCase
	log      *log.Helper
}

func NewServiceNameService(ServiceNameUseCase *biz.ServiceNameUseCase,
	logger log.Logger) *ServiceNameService {
	return &ServiceNameService{
		log:      log.NewHelper(log.With(logger, "module", "service/interface")),
		moduleUc: ServiceNameUseCase,
	}
}

func (s *ServiceNameService) CreateServiceName(ctx context.Context, req *v1.CreateServiceNameRequest) (*v1.ServiceNameInfoResponse, error) {
	bc := &biz.ServiceName{
		Id: req.Id,
	}
	ServiceNameInfo, err := s.moduleUc.Create(ctx, bc)
	if err != nil {
		return nil, err
	}
	return &v1.ServiceNameInfoResponse{
		Id: ServiceNameInfo.Id,
	}, nil
}
func (s *ServiceNameService) UpdateServiceName(ctx context.Context, req *v1.UpdateServiceNameRequest) (*v1.ServiceNameInfoResponse, error) {
	bc := &biz.ServiceName{
		Id: req.Id,
	}
	ServiceNameInfo, err := s.moduleUc.Update(ctx, bc)
	if err != nil {
		return nil, err
	}
	return &v1.ServiceNameInfoResponse{
		Id: ServiceNameInfo.Id,
	}, nil
}

func (s *ServiceNameService) DeleteServiceName(ctx context.Context, req *v1.DeleteServiceNameRequest) (*v1.ServiceNameCheckResponse, error) {
	err := s.moduleUc.Delete(ctx, req.Id)
	success := true
	if err != nil {
		success = false
	}
	return &v1.ServiceNameCheckResponse{
		IsSuccess: success,
	}, err
}

func (s *ServiceNameService) GetServiceName(ctx context.Context, req *v1.GetServiceNameRequest) (*v1.ServiceNameInfoResponse, error) {
	ServiceNameInfo, err := s.moduleUc.Get(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	response := &v1.ServiceNameInfoResponse{}
	response.Id = ServiceNameInfo.Id
	return response, nil
}

func (s *ServiceNameService) ListServiceName(ctx context.Context, req *v1.ListServiceNameRequest) (*v1.ListServiceNameReply, error) {
	if req.PageNum == 0 {
		req.PageNum = 1
	}
	if req.PageSize == 0 {
		req.PageSize = 10
	}
	ServiceNameInfoList, count, err := s.moduleUc.List(ctx, req.PageNum, req.PageSize)
	if err != nil {
		return nil, err
	}
	response := &v1.ListServiceNameReply{}
	response.Total = count
	for _, v := range ServiceNameInfoList {
		response.List = append(response.List, &v1.ServiceNameInfoResponse{
			Id: v.Id,
		})
	}
	return response, nil
}
