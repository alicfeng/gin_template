package validator

import (
	"github.com/alicfeng/gin_template/app/helper"
	"github.com/alicfeng/gin_template/app/repository"
	"github.com/go-playground/validator/v10"
	"strings"
)

// Exist 校验数据是否存在*/
func Exist(fieldLevel validator.FieldLevel) bool {
	params := strings.Split(fieldLevel.Param(), ".")
	if 2 != len(params) {
		panic("validator exist configuration error")
	}

	var rowsAffected int64

	repository.GetMySQL().
		Table(params[0]).
		Where(params[1]+" = ?", helper.DynamicValue(fieldLevel.Field())).
		Where("ISNULL(deleted_at)").
		Select(params[1]).
		Count(&rowsAffected)

	if 0 == rowsAffected {
		return false
	}

	return true
}

// Unique 校验数据是否唯一 */
func Unique(fieldLevel validator.FieldLevel) bool {
	params := strings.Split(fieldLevel.Param(), ".")
	if 2 > len(params) {
		panic("validator exist configuration error")
	}

	var rowsAffected int64
	var builder = repository.GetMySQL().Table(params[0])

	// 修改时过滤本身元祖数据 第三位 index is 2
	if 3 <= len(params) {
		builder.Where("id != ?", fieldLevel.Parent().FieldByName("Id").Int())
	}

	builder.Where(params[1]+" = ?", helper.DynamicValue(fieldLevel.Field())).
		Where("ISNULL(deleted_at)").
		Select(params[1]).
		Count(&rowsAffected)

	if 0 != rowsAffected {
		return false
	}

	return true
}
