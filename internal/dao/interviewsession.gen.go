// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dao

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"ELAB-registration-system-Backend/internal/model"
)

func newInterviewSession(db *gorm.DB, opts ...gen.DOOption) interviewSession {
	_interviewSession := interviewSession{}

	_interviewSession.interviewSessionDo.UseDB(db, opts...)
	_interviewSession.interviewSessionDo.UseModel(&model.InterviewSession{})

	tableName := _interviewSession.interviewSessionDo.TableName()
	_interviewSession.ALL = field.NewAsterisk(tableName)
	_interviewSession.ID = field.NewInt32(tableName, "id")
	_interviewSession.StartTime = field.NewTime(tableName, "start_time")
	_interviewSession.EndTime = field.NewTime(tableName, "end_time")
	_interviewSession.Location = field.NewString(tableName, "location")
	_interviewSession.Capacity = field.NewInt32(tableName, "capacity")
	_interviewSession.AppliedNum = field.NewInt32(tableName, "applied_num")
	_interviewSession.CreatedAt = field.NewTime(tableName, "created_at")
	_interviewSession.UpdatedAt = field.NewTime(tableName, "updated_at")

	_interviewSession.fillFieldMap()

	return _interviewSession
}

type interviewSession struct {
	interviewSessionDo interviewSessionDo

	ALL        field.Asterisk
	ID         field.Int32  // 主键
	StartTime  field.Time   // 面试开始时间
	EndTime    field.Time   // 面试结束时间
	Location   field.String // 面试地点
	Capacity   field.Int32  // 可参加人数
	AppliedNum field.Int32  // 已报名人数
	CreatedAt  field.Time   // 创建时间
	UpdatedAt  field.Time   // 最后修改时间

	fieldMap map[string]field.Expr
}

func (i interviewSession) Table(newTableName string) *interviewSession {
	i.interviewSessionDo.UseTable(newTableName)
	return i.updateTableName(newTableName)
}

func (i interviewSession) As(alias string) *interviewSession {
	i.interviewSessionDo.DO = *(i.interviewSessionDo.As(alias).(*gen.DO))
	return i.updateTableName(alias)
}

func (i *interviewSession) updateTableName(table string) *interviewSession {
	i.ALL = field.NewAsterisk(table)
	i.ID = field.NewInt32(table, "id")
	i.StartTime = field.NewTime(table, "start_time")
	i.EndTime = field.NewTime(table, "end_time")
	i.Location = field.NewString(table, "location")
	i.Capacity = field.NewInt32(table, "capacity")
	i.AppliedNum = field.NewInt32(table, "applied_num")
	i.CreatedAt = field.NewTime(table, "created_at")
	i.UpdatedAt = field.NewTime(table, "updated_at")

	i.fillFieldMap()

	return i
}

func (i *interviewSession) WithContext(ctx context.Context) *interviewSessionDo {
	return i.interviewSessionDo.WithContext(ctx)
}

func (i interviewSession) TableName() string { return i.interviewSessionDo.TableName() }

func (i interviewSession) Alias() string { return i.interviewSessionDo.Alias() }

func (i *interviewSession) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := i.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (i *interviewSession) fillFieldMap() {
	i.fieldMap = make(map[string]field.Expr, 8)
	i.fieldMap["id"] = i.ID
	i.fieldMap["start_time"] = i.StartTime
	i.fieldMap["end_time"] = i.EndTime
	i.fieldMap["location"] = i.Location
	i.fieldMap["capacity"] = i.Capacity
	i.fieldMap["applied_num"] = i.AppliedNum
	i.fieldMap["created_at"] = i.CreatedAt
	i.fieldMap["updated_at"] = i.UpdatedAt
}

func (i interviewSession) clone(db *gorm.DB) interviewSession {
	i.interviewSessionDo.ReplaceConnPool(db.Statement.ConnPool)
	return i
}

func (i interviewSession) replaceDB(db *gorm.DB) interviewSession {
	i.interviewSessionDo.ReplaceDB(db)
	return i
}

type interviewSessionDo struct{ gen.DO }

func (i interviewSessionDo) Debug() *interviewSessionDo {
	return i.withDO(i.DO.Debug())
}

func (i interviewSessionDo) WithContext(ctx context.Context) *interviewSessionDo {
	return i.withDO(i.DO.WithContext(ctx))
}

func (i interviewSessionDo) ReadDB() *interviewSessionDo {
	return i.Clauses(dbresolver.Read)
}

func (i interviewSessionDo) WriteDB() *interviewSessionDo {
	return i.Clauses(dbresolver.Write)
}

func (i interviewSessionDo) Session(config *gorm.Session) *interviewSessionDo {
	return i.withDO(i.DO.Session(config))
}

func (i interviewSessionDo) Clauses(conds ...clause.Expression) *interviewSessionDo {
	return i.withDO(i.DO.Clauses(conds...))
}

func (i interviewSessionDo) Returning(value interface{}, columns ...string) *interviewSessionDo {
	return i.withDO(i.DO.Returning(value, columns...))
}

func (i interviewSessionDo) Not(conds ...gen.Condition) *interviewSessionDo {
	return i.withDO(i.DO.Not(conds...))
}

func (i interviewSessionDo) Or(conds ...gen.Condition) *interviewSessionDo {
	return i.withDO(i.DO.Or(conds...))
}

func (i interviewSessionDo) Select(conds ...field.Expr) *interviewSessionDo {
	return i.withDO(i.DO.Select(conds...))
}

func (i interviewSessionDo) Where(conds ...gen.Condition) *interviewSessionDo {
	return i.withDO(i.DO.Where(conds...))
}

func (i interviewSessionDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *interviewSessionDo {
	return i.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (i interviewSessionDo) Order(conds ...field.Expr) *interviewSessionDo {
	return i.withDO(i.DO.Order(conds...))
}

func (i interviewSessionDo) Distinct(cols ...field.Expr) *interviewSessionDo {
	return i.withDO(i.DO.Distinct(cols...))
}

func (i interviewSessionDo) Omit(cols ...field.Expr) *interviewSessionDo {
	return i.withDO(i.DO.Omit(cols...))
}

func (i interviewSessionDo) Join(table schema.Tabler, on ...field.Expr) *interviewSessionDo {
	return i.withDO(i.DO.Join(table, on...))
}

func (i interviewSessionDo) LeftJoin(table schema.Tabler, on ...field.Expr) *interviewSessionDo {
	return i.withDO(i.DO.LeftJoin(table, on...))
}

func (i interviewSessionDo) RightJoin(table schema.Tabler, on ...field.Expr) *interviewSessionDo {
	return i.withDO(i.DO.RightJoin(table, on...))
}

func (i interviewSessionDo) Group(cols ...field.Expr) *interviewSessionDo {
	return i.withDO(i.DO.Group(cols...))
}

func (i interviewSessionDo) Having(conds ...gen.Condition) *interviewSessionDo {
	return i.withDO(i.DO.Having(conds...))
}

func (i interviewSessionDo) Limit(limit int) *interviewSessionDo {
	return i.withDO(i.DO.Limit(limit))
}

func (i interviewSessionDo) Offset(offset int) *interviewSessionDo {
	return i.withDO(i.DO.Offset(offset))
}

func (i interviewSessionDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *interviewSessionDo {
	return i.withDO(i.DO.Scopes(funcs...))
}

func (i interviewSessionDo) Unscoped() *interviewSessionDo {
	return i.withDO(i.DO.Unscoped())
}

func (i interviewSessionDo) Create(values ...*model.InterviewSession) error {
	if len(values) == 0 {
		return nil
	}
	return i.DO.Create(values)
}

func (i interviewSessionDo) CreateInBatches(values []*model.InterviewSession, batchSize int) error {
	return i.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (i interviewSessionDo) Save(values ...*model.InterviewSession) error {
	if len(values) == 0 {
		return nil
	}
	return i.DO.Save(values)
}

func (i interviewSessionDo) First() (*model.InterviewSession, error) {
	if result, err := i.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.InterviewSession), nil
	}
}

func (i interviewSessionDo) Take() (*model.InterviewSession, error) {
	if result, err := i.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.InterviewSession), nil
	}
}

func (i interviewSessionDo) Last() (*model.InterviewSession, error) {
	if result, err := i.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.InterviewSession), nil
	}
}

func (i interviewSessionDo) Find() ([]*model.InterviewSession, error) {
	result, err := i.DO.Find()
	return result.([]*model.InterviewSession), err
}

func (i interviewSessionDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.InterviewSession, err error) {
	buf := make([]*model.InterviewSession, 0, batchSize)
	err = i.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (i interviewSessionDo) FindInBatches(result *[]*model.InterviewSession, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return i.DO.FindInBatches(result, batchSize, fc)
}

func (i interviewSessionDo) Attrs(attrs ...field.AssignExpr) *interviewSessionDo {
	return i.withDO(i.DO.Attrs(attrs...))
}

func (i interviewSessionDo) Assign(attrs ...field.AssignExpr) *interviewSessionDo {
	return i.withDO(i.DO.Assign(attrs...))
}

func (i interviewSessionDo) Joins(fields ...field.RelationField) *interviewSessionDo {
	for _, _f := range fields {
		i = *i.withDO(i.DO.Joins(_f))
	}
	return &i
}

func (i interviewSessionDo) Preload(fields ...field.RelationField) *interviewSessionDo {
	for _, _f := range fields {
		i = *i.withDO(i.DO.Preload(_f))
	}
	return &i
}

func (i interviewSessionDo) FirstOrInit() (*model.InterviewSession, error) {
	if result, err := i.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.InterviewSession), nil
	}
}

func (i interviewSessionDo) FirstOrCreate() (*model.InterviewSession, error) {
	if result, err := i.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.InterviewSession), nil
	}
}

func (i interviewSessionDo) FindByPage(offset int, limit int) (result []*model.InterviewSession, count int64, err error) {
	result, err = i.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = i.Offset(-1).Limit(-1).Count()
	return
}

func (i interviewSessionDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = i.Count()
	if err != nil {
		return
	}

	err = i.Offset(offset).Limit(limit).Scan(result)
	return
}

func (i interviewSessionDo) Scan(result interface{}) (err error) {
	return i.DO.Scan(result)
}

func (i interviewSessionDo) Delete(models ...*model.InterviewSession) (result gen.ResultInfo, err error) {
	return i.DO.Delete(models)
}

func (i *interviewSessionDo) withDO(do gen.Dao) *interviewSessionDo {
	i.DO = *do.(*gen.DO)
	return i
}
