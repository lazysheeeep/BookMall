package serializer

import "BookMall/model"

type CategoryVO struct {
	Category string `json:"category"`
	Text     string `json:"text"`
}

func BuildCategory(category model.Category) CategoryVO {
	return CategoryVO{
		Category: category.Category,
		Text:     category.Text,
	}
}

func BuildCategories(categories []model.Category) (items []CategoryVO) {
	for _, item := range categories {
		tmp := BuildCategory(item)
		items = append(items, tmp)
	}
	return
}
