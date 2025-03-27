package gormutil

import (
	"strings"

	"gorm.io/gorm"
)

// PaginateScope 用于构建分页查询的 Scope，限制页码和页大小
func PaginateScope(page, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		// 页码小于等于0时默认第1页
		if page <= 0 {
			page = 1
		}
		// 限制每页条数，超过100时取100，小于等于0时默认10
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}
		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

// BuildQuery 根据传入的 filters 构建查询条件，支持关键词查询和其他字段精确匹配
func BuildQuery(db *gorm.DB, filters map[string]interface{}, keywordFields []string) *gorm.DB {
	// 如果存在 keywords，则使用 KeywordScope
	if keywords, ok := filters["keywords"].(string); ok && keywords != "" {
		db = db.Scopes(KeywordScope(keywords, keywordFields))
	}
	// 处理其他条件（精确匹配）
	for key, value := range filters {
		if key == "keywords" {
			continue
		}
		// 过滤掉 nil 或空值
		if value == nil || value == "" {
			continue
		}

		db = db.Where(key+" = ?", value)
	}
	return db
}

// KeywordScope 封装关键词查询的条件，将所有指定字段用 OR 拼接
func KeywordScope(keyword string, fields []string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if keyword == "" || len(fields) == 0 {
			return db
		}
		// 构造查询条件
		var queryBuilder strings.Builder
		args := make([]interface{}, 0, len(fields))
		for i, field := range fields {
			if i > 0 {
				queryBuilder.WriteString(" OR ")
			}
			queryBuilder.WriteString(field + " LIKE ?")
			args = append(args, "%"+keyword+"%")
		}
		return db.Where(queryBuilder.String(), args...)
	}
}
