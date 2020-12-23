package mappers

import (
	"github.com/Masterminds/squirrel"
)

type UserMapper struct {

}

func (*UserMapper) GetUserListByPage(page, size int) *SqlMapper {
	return Mapper(squirrel.Select("*").From("t_user").Limit(uint64(size)).Offset(uint64(page*size - size)).ToSql())
}

func (*UserMapper) GetUserByID(id int) *SqlMapper {
	return Mapper(squirrel.Select("*").From("t_user").Where("u_id = ?", id).ToSql())
}

func (this *UserMapper) DeleteUserByID(id int) *SqlMapper {
	return Mapper(squirrel.Delete("t_user").Where("u_id = ?", id).ToSql())
}
