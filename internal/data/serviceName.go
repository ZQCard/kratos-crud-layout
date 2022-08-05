package data

import (
	"context"
	"github.com/ZQCard/kratos-crud-layout/internal/biz"
	"github.com/ZQCard/kratos-crud-layout/internal/data/entity"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	"net/http"
)

type ServiceNameRepo struct {
	data *Data
	log  *log.Helper
}

func (b ServiceNameRepo) GetServiceNameByParams(params map[string]interface{}) (record entity.ServiceNameEntity, err error) {
	if len(params) == 0 {
		return entity.ServiceNameEntity{}, errors.New(http.StatusBadRequest, "MISSING_CONDITION", "缺少搜索条件")
	}
	conn := b.data.db.Model(&entity.ServiceNameEntity{})
	if id, ok := params["id"]; ok && id.(int64) != 0 {
		conn = conn.Where("id = ?", id)
	}
	if err = conn.First(&record).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.ServiceNameEntity{}, errors.New(http.StatusBadRequest, "RECORD_NOT_FOUND", biz.ErrRecordNotFound)
		}
		return record, errors.New(500, "SYSTEM_ERROR", err.Error())
	}
	return record, nil
}

func (b ServiceNameRepo) CreateServiceName(ctx context.Context, reqData *biz.ServiceName) (*biz.ServiceName, error) {
	modelTable := entity.ServiceNameEntity{}
	modelTable.Id = reqData.Id
	if err := b.data.db.Model(&modelTable).Create(&modelTable).Error; err != nil {
		return nil, errors.New(500, "SYSTEM_ERROR", err.Error())
	}
	response := ModelToResponse(modelTable)
	return &response, nil
}

func (b ServiceNameRepo) UpdateServiceName(ctx context.Context, reqData *biz.ServiceName) (*biz.ServiceName, error) {
	// 根据id查找记录
	record, err := b.GetServiceNameByParams(map[string]interface{}{
		"id": reqData.Id,
	})
	if err != nil {
		return nil, err
	}
	// 更新记录
	record.Id = reqData.Id
	if err := b.data.db.Model(&record).Where("id = ?", record.Id).Save(&record).Error; err != nil {
		return nil, errors.New(500, "SYSTEM_ERROR", err.Error())
	}
	// 返回数据
	response := ModelToResponse(record)
	return &response, nil
}

func (b ServiceNameRepo) GetServiceName(ctx context.Context, id int64) (*biz.ServiceName, error) {
	// 根据id查找记录
	record, err := b.GetServiceNameByParams(map[string]interface{}{
		"id": id,
	})
	if err != nil {
		return nil, err
	}
	// 返回数据
	response := ModelToResponse(record)
	return &response, nil
}

func (b ServiceNameRepo) ListServiceName(ctx context.Context, pageNum, pageSize int64) ([]*biz.ServiceName, int64, error) {
	list := []entity.ServiceNameEntity{}
	conn := b.data.db.Model(&entity.ServiceNameEntity{})
	err := conn.Scopes(entity.Paginate(pageNum, pageSize)).Find(&list).Error
	if err != nil {
		return nil, 0, errors.New(500, "SYSTEM_ERROR", err.Error())
	}

	count := int64(0)
	conn.Count(&count)
	rv := make([]*biz.ServiceName, 0, len(list))
	for _, record := range list {
		serviceName := ModelToResponse(record)
		rv = append(rv, &serviceName)
	}
	return rv, count, nil
}

func (b ServiceNameRepo) DeleteServiceName(ctx context.Context, id int64) error {
	// 根据id查找记录
	record, err := b.GetServiceNameByParams(map[string]interface{}{
		"id": id,
	})
	if err != nil {
		return err
	}
	return b.data.db.Model(&record).Where("id = ?", id).Delete(&record).Error
}

// ModelToResponse 转换 serviceName 表中所有字段的值
func ModelToResponse(serviceName entity.ServiceNameEntity) biz.ServiceName {
	administratorInfoRsp := biz.ServiceName{}
	administratorInfoRsp.Id = serviceName.Id
	return administratorInfoRsp
}

func NewServiceNameRepo(data *Data, logger log.Logger) biz.ServiceNameRepo {
	return &ServiceNameRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/administrator-service")),
	}
}
