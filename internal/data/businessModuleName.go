package data

import (
	"context"
	"errors"
	"github.com/ZQCard/kratos-crud-layout/internal/biz"
	"github.com/ZQCard/kratos-crud-layout/internal/data/model"
	"github.com/go-kratos/kratos/v2/log"
)

type businessModuleNameRepo struct {
	data *Data
	log  *log.Helper
}

func (b businessModuleNameRepo) GetBusinessModuleNameByParams(params map[string]interface{}) (record model.BusinessModuleNameTable, err error) {
	if len(params) == 0 {
		return model.BusinessModuleNameTable{}, errors.New("查询条件不得为空")
	}
	conn := b.data.db.Model(&model.BusinessModuleNameTable{})
	if id, ok := params["id"]; ok && id.(int64) != 0 {
		conn = conn.Where("id = ?", id)
	}
	err = conn.First(&record).Error
	return record, err
}

func (b businessModuleNameRepo) CreateBusinessModuleName(ctx context.Context, reqData *biz.BusinessModuleName) (*biz.BusinessModuleName, error) {
	response := &biz.BusinessModuleName{}
	modelTable := model.BusinessModuleNameTable{}
	modelTable.Id = reqData.Id
	if err := b.data.db.Model(&modelTable).Create(&modelTable).Error; err != nil {
		return nil, err
	}
	return response, nil
}

func (b businessModuleNameRepo) UpdateBusinessModuleName(ctx context.Context, reqData *biz.BusinessModuleName) (*biz.BusinessModuleName, error) {
	// 根据id查找记录
	record, err := b.GetBusinessModuleNameByParams(map[string]interface{}{
		"id": reqData.Id,
	})
	if err != nil {
		return nil, err
	}
	// 更新记录
	record.Id = reqData.Id
	if err := b.data.db.Model(&record).Where("id = ?", record.Id).Save(&record).Error; err != nil {
		return nil, err
	}
	// 返回数据
	response := &biz.BusinessModuleName{}
	response.Id = record.Id
	return response, nil
}

func (b businessModuleNameRepo) GetBusinessModuleName(ctx context.Context, id int64) (*biz.BusinessModuleName, error) {
	// 根据id查找记录
	record, err := b.GetBusinessModuleNameByParams(map[string]interface{}{
		"id": id,
	})
	if err != nil {
		return nil, err
	}
	// 返回数据
	response := &biz.BusinessModuleName{}
	response.Id = record.Id
	return response, nil
}

func (b businessModuleNameRepo) ListBusinessModuleName(ctx context.Context, pageNum, pageSize int64) ([]*biz.BusinessModuleName, int64, error) {
	list := []model.BusinessModuleNameTable{}
	conn := b.data.db.Model(&model.BusinessModuleNameTable{})
	err := conn.Offset(int(pageNum)).Limit(int(pageSize)).Find(list).Error
	if err != nil {
		return nil, 0, err
	}
	count := int64(0)
	conn.Count(&count)
	rv := make([]*biz.BusinessModuleName, 0, len(list))
	for _, record := range list {
		rv = append(rv, &biz.BusinessModuleName{
			Id: record.Id,
		})
	}
	return rv, count, nil
}

func (b businessModuleNameRepo) DeleteBusinessModuleName(ctx context.Context, id int64) error {
	// 根据id查找记录
	record, err := b.GetBusinessModuleNameByParams(map[string]interface{}{
		"id": id,
	})
	if err != nil {
		return err
	}
	return b.data.db.Model(&record).Where("id = ?", id).Delete(&record).Error
}

func NewBusinessModuleNameRepo(data *Data, logger log.Logger) biz.BusinessModuleNameRepo {
	return &businessModuleNameRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/administrator-service")),
	}
}
