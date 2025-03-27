package util

import (
	"strings"

	"github.com/gin-gonic/gin"
)

// ParseFilters 从查询字符串中解析 filters 参数
func ParseFilters(c *gin.Context) (map[string]interface{}, error) {
	filters := make(map[string]interface{})

	// 遍历所有查询参数
	for key, values := range c.Request.URL.Query() {
		if strings.HasPrefix(key, "filters[") && strings.HasSuffix(key, "]") {
			// 提取出过滤条件的键
			filterKey := strings.TrimSuffix(strings.TrimPrefix(key, "filters["), "]")

			// 检查值是否为空，如果为空则不添加
			if len(values) > 0 && values[0] != "" {
				filters[filterKey] = values[0]
			}
		}
	}
	return filters, nil
}
