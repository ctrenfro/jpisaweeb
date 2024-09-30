package items

type Item struct {
	Name        string
	ID          string
	Img         string
	Description string
	Price       string
}

func ReturnItems() []Item {

	return []Item{
		{
			Name:        "camera1",
			ID:          "1",
			Img:         "/public/assets/images/camera1.jpeg",
			Description: "this is a camera",
			Price:       "99.99",
		},
		{
			Name:        "camera2",
			ID:          "2",
			Img:         "/public/assets/images/camera2.jpeg",
			Description: "this is a camera",
			Price:       "49.99",
		},
		{
			Name:        "camera3",
			ID:          "3",
			Img:         "/public/assets/images/camera3.jpeg",
			Description: "this is a camera",
			Price:       "149.99",
		},
	}

}

func ReturnItem(itemID string) Item {

	itemList := ReturnItems()
	for _, item := range itemList {
		if itemID == item.ID {
			return Item{
				Name:        item.Name,
				ID:          item.ID,
				Img:         item.Img,
				Description: item.Description,
				Price:       item.Price,
			}
		}

	}

	return Item{}
}
