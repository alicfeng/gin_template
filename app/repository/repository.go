package repository

import (
	"gorm.io/gorm"
)

type Repository struct {
}

// Find 查询数据集合 */
func (Repository *Repository) Find(dest interface{}, conditions ...interface{}) *gorm.DB {

	return _mysql.Find(dest, conditions)
}

// ModelUpdate 根据模型条件修改数据 */
func (Repository *Repository) ModelUpdate(model interface{}, column string, value interface{}) *gorm.DB {

	return _mysql.Model(model).Update(column, value)
}

// ConditionUpdate 根据条件修改数据 */
func (Repository *Repository) ConditionUpdate(column string, value interface{}, query interface{}, args ...interface{}) *gorm.DB {

	return _mysql.Where(query, args).Update(column, value)
}

// First 查询一条数据 */
func (Repository *Repository) First(dest interface{}, conditions ...interface{}) *gorm.DB {

	return _mysql.First(dest, conditions)
}

// Delete 删除数据 */
func (Repository *Repository) Delete(value interface{}, conditions ...interface{}) *gorm.DB {
	return _mysql.Delete(value, conditions)
}
