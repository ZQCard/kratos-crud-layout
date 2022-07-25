package service

import (
	"context"
	"github.com/ZQCard/kratos-crud-layout/internal/biz"
	"github.com/go-kratos/kratos/v2/log"

	v1 "github.com/ZQCard/kratos-crud-layout/api/yourServiceName/v1"
)

type BusinessModuleNameService struct {
	v1.UnimplementedBusinessModuleNameServer
	moduleUc *biz.BusinessModuleNameUseCase
	log      *log.Helper
}

func NewBusinessModuleNameService(businessModuleNameUseCase *biz.BusinessModuleNameUseCase,
	logger log.Logger) *BusinessModuleNameService {
	return &BusinessModuleNameService{
		log:      log.NewHelper(log.With(logger, "module", "service/interface")),
		moduleUc: businessModuleNameUseCase,
	}
}

func (s *BusinessModuleNameService) CreateBusinessModuleName(ctx context.Context, req *v1.CreateBusinessModuleNameRequest) (*v1.BusinessModuleNameInfoResponse, error) {
	bc := &biz.BusinessModuleName{
		Id: req.Id,
	}
	businessModuleNameInfo, err := s.moduleUc.Create(ctx, bc)
	return &v1.BusinessModuleNameInfoResponse{
		Id: businessModuleNameInfo.Id,
	}, err
}
func (s *BusinessModuleNameService) UpdateBusinessModuleName(ctx context.Context, req *v1.UpdateBusinessModuleNameRequest) (*v1.BusinessModuleNameInfoResponse, error) {
	bc := &biz.BusinessModuleName{
		Id: req.Id,
	}
	businessModuleNameInfo, err := s.moduleUc.Update(ctx, bc)
	return &v1.BusinessModuleNameInfoResponse{
		Id: businessModuleNameInfo.Id,
	}, err
}
func (s *BusinessModuleNameService) DeleteBusinessModuleName(ctx context.Context, req *v1.DeleteBusinessModuleNameRequest) (*v1.BusinessModuleNameCheckResponse, error) {
	err := s.moduleUc.Delete(ctx, req.Id)
	success := true
	if err != nil {
		success = false
	}
	return &v1.BusinessModuleNameCheckResponse{
		IsSuccess: success,
	}, err
}
func (s *BusinessModuleNameService) GetBusinessModuleName(ctx context.Context, req *v1.GetBusinessModuleNameRequest) (*v1.BusinessModuleNameInfoResponse, error) {
	businessModuleNameInfo, err := s.moduleUc.Get(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	response := &v1.BusinessModuleNameInfoResponse{}
	response.Id = businessModuleNameInfo.Id
	return response, nil
}
func (s *BusinessModuleNameService) ListBusinessModuleName(ctx context.Context, req *v1.ListBusinessModuleNameRequest) (*v1.ListBusinessModuleNameReply, error) {
	businessModuleNameInfoList, count, err := s.moduleUc.List(ctx, req.PageNum, req.PageSize)
	if err != nil {
		return nil, err
	}
	response := &v1.ListBusinessModuleNameReply{}
	response.Total = count
	for _, v := range businessModuleNameInfoList {
		response.List = append(response.List, &v1.BusinessModuleNameInfoResponse{
			Id: v.Id,
		})
	}
	return response, nil
}
