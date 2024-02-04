package gen

import (
	"gorm.io/gen/internal/generate"
	"gorm.io/gen/internal/model"
)

// getModelConfigByTableName 通过表名获取模型配置
func (g *Generator) getModelConfigByTableName(tableName string, opts ...ModelOpt) *model.Config {
	modelName := g.db.Config.NamingStrategy.SchemaName(tableName)
	return g.genModelConfig(tableName, modelName, opts)
}

// Except 排除一些模型生成
//
// 需要在调用 Execute 之前调用
func (g *Generator) Except(tableName ...string) *Generator {
	if len(tableName) == 0 {
		return g
	}

	for _, name := range tableName {
		conf := g.getModelConfigByTableName(name)
		_, structName, _ := conf.GetNames()
		delete(g.models, structName)
		delete(g.Data, structName)
	}
	return g
}

// Only 仅生成一些模型
//
// 需要在调用 Execute 之前调用
func (g *Generator) Only(tableName ...string) *Generator {
	if len(tableName) == 0 {
		return g
	}
	models := make(map[string]*generate.QueryStructMeta)
	datas := make(map[string]*genInfo)
	for _, name := range tableName {
		conf := g.getModelConfigByTableName(name)
		_, structName, _ := conf.GetNames()
		if _, ok := g.models[structName]; ok {
			models[structName] = g.models[structName]
		}
		if _, ok := g.Data[structName]; ok {
			datas[structName] = g.Data[structName]
		}
	}
	g.models = models
	g.Data = datas
	return g
}

// QueryStructMetas 获取所有模型配置
func (g *Generator) QueryStructMetas() []*generate.QueryStructMeta {
	models := make([]*generate.QueryStructMeta, 0, len(g.models))
	for _, m := range g.models {
		models = append(models, m)
	}
	return models
}
