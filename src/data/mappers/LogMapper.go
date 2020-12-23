package mappers

import (
	"ginSkill/src/models/LogModel"
	"github.com/Masterminds/squirrel"
)

type LogMapper struct {
}

func (*LogMapper) AddLog(log *LogModel.LogImpl) *SqlMapper {
	return Mapper(squirrel.Insert(log.TableName()).Columns("log_name", "log_time").Values(log.Name, log.UpdateTime).ToSql())
}
