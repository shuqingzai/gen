// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"gorm.io/gen/tests/.expect/dal_test_relation/model"
)

func newCreditCard(db *gorm.DB, opts ...gen.DOOption) creditCard {
	_creditCard := creditCard{}

	_creditCard.creditCardDo.UseDB(db, opts...)
	_creditCard.creditCardDo.UseModel(&model.CreditCard{})

	tableName := _creditCard.creditCardDo.TableName()
	_creditCard.ALL = field.NewAsterisk(tableName)
	_creditCard.ID = field.NewInt64(tableName, "id")
	_creditCard.CreatedAt = field.NewTime(tableName, "created_at")
	_creditCard.UpdatedAt = field.NewTime(tableName, "updated_at")
	_creditCard.DeletedAt = field.NewField(tableName, "deleted_at")
	_creditCard.Number = field.NewString(tableName, "number")
	_creditCard.CustomerRefer = field.NewInt64(tableName, "customer_refer")
	_creditCard.BankID = field.NewInt64(tableName, "bank_id")

	_creditCard.fillFieldMap()

	return _creditCard
}

type creditCard struct {
	creditCardDo creditCardDo

	ALL           field.Asterisk
	ID            field.Int64
	CreatedAt     field.Time
	UpdatedAt     field.Time
	DeletedAt     field.Field
	Number        field.String
	CustomerRefer field.Int64
	BankID        field.Int64

	fieldMap map[string]field.Expr
}

func (c creditCard) Table(newTableName string) *creditCard {
	c.creditCardDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c creditCard) As(alias string) *creditCard {
	c.creditCardDo.DO = *(c.creditCardDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *creditCard) updateTableName(table string) *creditCard {
	c.ALL = field.NewAsterisk(table)
	c.ID = field.NewInt64(table, "id")
	c.CreatedAt = field.NewTime(table, "created_at")
	c.UpdatedAt = field.NewTime(table, "updated_at")
	c.DeletedAt = field.NewField(table, "deleted_at")
	c.Number = field.NewString(table, "number")
	c.CustomerRefer = field.NewInt64(table, "customer_refer")
	c.BankID = field.NewInt64(table, "bank_id")

	c.fillFieldMap()

	return c
}

func (c *creditCard) WithContext(ctx context.Context) *creditCardDo {
	return c.creditCardDo.WithContext(ctx)
}

func (c creditCard) TableName() string { return c.creditCardDo.TableName() }

func (c creditCard) Alias() string { return c.creditCardDo.Alias() }

func (c creditCard) Columns(cols ...field.Expr) gen.Columns { return c.creditCardDo.Columns(cols...) }

func (c *creditCard) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c *creditCard) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 7)
	c.fieldMap["id"] = c.ID
	c.fieldMap["created_at"] = c.CreatedAt
	c.fieldMap["updated_at"] = c.UpdatedAt
	c.fieldMap["deleted_at"] = c.DeletedAt
	c.fieldMap["number"] = c.Number
	c.fieldMap["customer_refer"] = c.CustomerRefer
	c.fieldMap["bank_id"] = c.BankID
}

func (c creditCard) clone(db *gorm.DB) creditCard {
	c.creditCardDo.ReplaceConnPool(db.Statement.ConnPool)
	return c
}

func (c creditCard) replaceDB(db *gorm.DB) creditCard {
	c.creditCardDo.ReplaceDB(db)
	return c
}

type creditCardDo struct{ gen.DO }

func (c creditCardDo) Debug() *creditCardDo {
	return c.withDO(c.DO.Debug())
}

func (c creditCardDo) WithContext(ctx context.Context) *creditCardDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c creditCardDo) ReadDB() *creditCardDo {
	return c.Clauses(dbresolver.Read)
}

func (c creditCardDo) WriteDB() *creditCardDo {
	return c.Clauses(dbresolver.Write)
}

func (c creditCardDo) Session(config *gorm.Session) *creditCardDo {
	return c.withDO(c.DO.Session(config))
}

func (c creditCardDo) Clauses(conds ...clause.Expression) *creditCardDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c creditCardDo) Returning(value interface{}, columns ...string) *creditCardDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c creditCardDo) Not(conds ...gen.Condition) *creditCardDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c creditCardDo) Or(conds ...gen.Condition) *creditCardDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c creditCardDo) Select(conds ...field.Expr) *creditCardDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c creditCardDo) Where(conds ...gen.Condition) *creditCardDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c creditCardDo) Order(conds ...field.Expr) *creditCardDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c creditCardDo) Distinct(cols ...field.Expr) *creditCardDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c creditCardDo) Omit(cols ...field.Expr) *creditCardDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c creditCardDo) Join(table schema.Tabler, on ...field.Expr) *creditCardDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c creditCardDo) LeftJoin(table schema.Tabler, on ...field.Expr) *creditCardDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c creditCardDo) RightJoin(table schema.Tabler, on ...field.Expr) *creditCardDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c creditCardDo) Group(cols ...field.Expr) *creditCardDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c creditCardDo) Having(conds ...gen.Condition) *creditCardDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c creditCardDo) Limit(limit int) *creditCardDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c creditCardDo) Offset(offset int) *creditCardDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c creditCardDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *creditCardDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c creditCardDo) Unscoped() *creditCardDo {
	return c.withDO(c.DO.Unscoped())
}

func (c creditCardDo) Create(values ...*model.CreditCard) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c creditCardDo) CreateInBatches(values []*model.CreditCard, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c creditCardDo) Save(values ...*model.CreditCard) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c creditCardDo) First() (*model.CreditCard, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreditCard), nil
	}
}

func (c creditCardDo) Take() (*model.CreditCard, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreditCard), nil
	}
}

func (c creditCardDo) Last() (*model.CreditCard, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreditCard), nil
	}
}

func (c creditCardDo) Find() ([]*model.CreditCard, error) {
	result, err := c.DO.Find()
	return result.([]*model.CreditCard), err
}

func (c creditCardDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.CreditCard, err error) {
	buf := make([]*model.CreditCard, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c creditCardDo) FindInBatches(result *[]*model.CreditCard, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c creditCardDo) Attrs(attrs ...field.AssignExpr) *creditCardDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c creditCardDo) Assign(attrs ...field.AssignExpr) *creditCardDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c creditCardDo) Joins(fields ...field.RelationField) *creditCardDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c creditCardDo) Preload(fields ...field.RelationField) *creditCardDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c creditCardDo) FirstOrInit() (*model.CreditCard, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreditCard), nil
	}
}

func (c creditCardDo) FirstOrCreate() (*model.CreditCard, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreditCard), nil
	}
}

func (c creditCardDo) FindByPage(offset int, limit int) (result []*model.CreditCard, count int64, err error) {
	result, err = c.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = c.Offset(-1).Limit(-1).Count()
	return
}

func (c creditCardDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c creditCardDo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c creditCardDo) Delete(models ...*model.CreditCard) (result gen.ResultInfo, err error) {
	return c.DO.Delete(models)
}

func (c *creditCardDo) withDO(do gen.Dao) *creditCardDo {
	c.DO = *do.(*gen.DO)
	return c
}
