package mappers

import (
	"ginSkill/src/dbs"
	"github.com/jinzhu/gorm"
)

type SqlMapper struct {
	Sql  string
	Args []interface{}
	db   *gorm.DB
}

func NewSqlMapper(sql string, args []interface{}) *SqlMapper {
	return &SqlMapper{Sql: sql, Args: args}
}

//转化sqlMapper
func Mapper(sql string, args []interface{}, err error) *SqlMapper {
	if err != nil {
		panic(err)
	}
	return NewSqlMapper(sql, args)
}

func (this *SqlMapper) setDB(db *gorm.DB) {
	this.db = db
}

func (this *SqlMapper) Query() *gorm.DB {
	//if this.db != nil {
	//	this.db.Raw(this.Sql, this.Args)
	//}
	return dbs.Orm.Raw(this.Sql, this.Args)
}

func (this *SqlMapper) Exec() *gorm.DB {
	//if this.db != nil {
	//	this.db.Exec(this.Sql, this.Args)
	//}
	return dbs.Orm.Exec(this.Sql, this.Args)
}

type SqlMappers []*SqlMapper

func Mappers(SqlMappers ...*SqlMapper) SqlMappers {
	return SqlMappers
}

func (this SqlMappers) apply(tx *gorm.DB) {
	for _, sql := range this {
		sql.setDB(tx)
	}
}

func (this SqlMappers) Exec(f func() error) error {
	return dbs.Orm.Transaction(func(tx *gorm.DB) error {
		this.apply(tx)
		return f()
	})
}
