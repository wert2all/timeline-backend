package validator

import "timeline/backend/graph/model"

func NewLimit(limit *model.Limit) int {
	result := 0
	if limit != nil && limit.To != nil {
		result = *limit.To
	}
	if result > 100 || result < 0 {
		result = 100
	}
	return result
}
