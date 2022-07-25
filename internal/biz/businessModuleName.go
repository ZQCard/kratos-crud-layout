package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type BusinessModuleName struct {
	Id int64
}

// 模块接口
type BusinessModuleNameRepo interface {
	CreateBusinessModuleName(ctx context.Context, reqData *BusinessModuleName) (*BusinessModuleName, error)
	UpdateBusinessModuleName(ctx context.Context, reqData *BusinessModuleName) (*BusinessModuleName, error)
	GetBusinessModuleName(ctx context.Context, id int64) (*BusinessModuleName, error)
	ListBusinessModuleName(ctx context.Context, pageNum, pageSize int64) ([]*BusinessModuleName, int64, error)
	DeleteBusinessModuleName(ctx context.Context, id int64) error
}

type BusinessModuleNameUseCase struct {
	repo BusinessModuleNameRepo
	log  *log.Helper
}

func NewBusinessModuleNameUseCase(repo BusinessModuleNameRepo, logger log.Logger) *BusinessModuleNameUseCase {
	return &BusinessModuleNameUseCase{repo: repo, log: log.NewHelper(log.With(logger, "module", "usecase/beer"))}
}

func (uc *BusinessModuleNameUseCase) Create(ctx context.Context, data *BusinessModuleName) (*BusinessModuleName, error) {
	return uc.repo.CreateBusinessModuleName(ctx, data)
}

func (uc *BusinessModuleNameUseCase) Delete(ctx context.Context, id int64) error {
	return uc.repo.DeleteBusinessModuleName(ctx, id)
}

func (uc *BusinessModuleNameUseCase) Update(ctx context.Context, data *BusinessModuleName) (*BusinessModuleName, error) {
	return uc.repo.UpdateBusinessModuleName(ctx, data)
}

func (uc *BusinessModuleNameUseCase) Get(ctx context.Context, id int64) (*BusinessModuleName, error) {
	return uc.repo.GetBusinessModuleName(ctx, id)
}

func (uc *BusinessModuleNameUseCase) List(ctx context.Context, pageNum, pageSize int64) ([]*BusinessModuleName, int64, error) {
	return uc.repo.ListBusinessModuleName(ctx, pageNum, pageSize)
}
