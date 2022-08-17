package data

import (
	"context"
	"github.com/ZQCard/kratos-crud-layout/internal/biz"
	"github.com/ZQCard/kratos-crud-layout/internal/data/entity"
	"github.com/ZQCard/kratos-crud-layout/third_party/timeSugar"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	"net/http"
)

type ServiceNameRepo struct {
	data *Data
	log  *log.Helper
}

// searchParam 搜索条件
func (s ServiceNameRepo) searchParam(params map[string]interface{}) *gorm.DB {
	conn := s.data.db.Model(&entity.ServiceNameEntity{})
	if id, ok := params["id"]; ok && id.(int64) != 0 {
		conn = conn.Where("id = ?", id)
	}
	if name, ok := params["name"]; ok && name.(string) != "" {
		conn = conn.Where("name LIKE ?", "%"+name.(string)+"%")
	}
	// 开始时间
	if start, ok := params["created_at_start"]; ok && start.(string) != "" {
		conn = conn.Where("created_at >= ", start.(string))
	}
	// 结束时间
	if end, ok := params["created_at_end"]; ok && end.(string) != "" {
		conn = conn.Where("created_at <= ", end.(string))
	}
	// 已删除
	if isDeleted, ok := params["is_deleted"]; ok && isDeleted.(string) == entity.ServiceNameDeleted {
		conn = conn.Scopes(entity.HasDelete())
	}
	// 未删除
	if isDeleted, ok := params["is_deleted"]; ok && isDeleted.(string) == entity.ServiceNameUnDeleted {
		conn = conn.Scopes(entity.UnDelete())
	}
	return conn
}

func (s ServiceNameRepo) GetServiceNameByParams(params map[string]interface{}) (record entity.ServiceNameEntity, err error) {
	if len(params) == 0 {
		return entity.ServiceNameEntity{}, errors.New(http.StatusBadRequest, "MISSING_CONDITION", "缺少搜索条件")
	}
	conn := s.searchParam(params)
	if err = conn.First(&record).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.ServiceNameEntity{}, errors.New(http.StatusBadRequest, "RECORD_NOT_FOUND", biz.ErrRecordNotFound)
		}
		return record, errors.New(http.StatusInternalServerError, "SYSTEM_ERROR", err.Error())
	}
	return record, nil
}

func (s ServiceNameRepo) CreateServiceName(ctx context.Context, reqData *biz.ServiceName) (*biz.ServiceName, error) {
	modelTable := entity.ServiceNameEntity{
		Id:        0,
		Name:      reqData.Name,
		CreatedAt: timeSugar.GetCurrentTime(),
		UpdatedAt: timeSugar.GetCurrentTime(),
		DeletedAt: "",
	}

	modelTable.Id = reqData.Id
	if err := s.data.db.Model(&modelTable).Create(&modelTable).Error; err != nil {
		return nil, errors.New(http.StatusInternalServerError, "SYSTEM_ERROR", err.Error())
	}
	response := ModelToResponse(modelTable)
	return &response, nil
}

func (s ServiceNameRepo) UpdateServiceName(ctx context.Context, reqData *biz.ServiceName) (*biz.ServiceName, error) {
	// 根据id查找记录
	record, err := s.GetServiceNameByParams(map[string]interface{}{
		"id": reqData.Id,
	})
	if err != nil {
		return nil, err
	}
	// 更新字段
	record.Name = reqData.Name
	if err := s.data.db.Model(&record).Where("id = ?", record.Id).Save(&record).Error; err != nil {
		return nil, errors.New(http.StatusInternalServerError, "SYSTEM_ERROR", err.Error())
	}
	// 返回数据
	response := ModelToResponse(record)
	return &response, nil
}

func (s ServiceNameRepo) GetServiceName(ctx context.Context, params map[string]interface{}) (*biz.ServiceName, error) {
	// 根据id查找记录
	record, err := s.GetServiceNameByParams(params)
	if err != nil {
		return nil, err
	}
	// 返回数据
	response := ModelToResponse(record)
	return &response, nil
}

func (s ServiceNameRepo) ListServiceName(ctx context.Context, pageNum, pageSize int64, params map[string]interface{}) ([]*biz.ServiceName, int64, error) {
	list := []entity.ServiceNameEntity{}
	conn := s.searchParam(params)
	err := conn.Scopes(entity.Paginate(pageNum, pageSize)).Find(&list).Error
	if err != nil {
		return nil, 0, errors.New(http.StatusInternalServerError, "SYSTEM_ERROR", err.Error())
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

func (s ServiceNameRepo) DeleteServiceName(ctx context.Context, id int64) error {
	// 根据id查找记录
	record, err := s.GetServiceNameByParams(map[string]interface{}{
		"id": id,
	})
	if err != nil {
		return err
	}
	if id != record.Id {
		return errors.New(http.StatusBadRequest, "RECORD_NOT_FOUND", biz.ErrRecordNotFound)
	}
	return s.data.db.Model(&record).Where("id = ?", id).UpdateColumn("deleted_at", timeSugar.GetCurrentYMDHIS()).Error
}

func (s ServiceNameRepo) RecoverServiceName(ctx context.Context, id int64) error {
	if id == 0 {
		return errors.New(http.StatusBadRequest, "MISSING_CONDITION", "缺少搜索条件")
	}
	err := s.data.db.Model(entity.ServiceNameEntity{}).Where("id = ?", id).UpdateColumn("deleted_at", "").Error
	if err != nil {
		return errors.New(http.StatusInternalServerError, "SYSTEM_ERROR", err.Error())
	}
	return nil
}

// ModelToResponse 转换 serviceName 表中所有字段的值
func ModelToResponse(serviceName entity.ServiceNameEntity) biz.ServiceName {
	administratorInfoRsp := biz.ServiceName{
		Id:        serviceName.Id,
		Name:      serviceName.Name,
		CreatedAt: timeSugar.FormatYMDHIS(serviceName.CreatedAt),
		UpdatedAt: timeSugar.FormatYMDHIS(serviceName.UpdatedAt),
		DeletedAt: serviceName.DeletedAt,
	}
	return administratorInfoRsp
}

func NewServiceNameRepo(data *Data, logger log.Logger) biz.ServiceNameRepo {
	return &ServiceNameRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/administrator-service")),
	}
}
