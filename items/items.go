package items

type Items struct {
	Name        string
	ID          string
	Img         string
	Description string
}

func ReturnItems() []Items {

	return []Items{
		{
			Name:        "camera1",
			ID:          "1",
			Img:         "./public/assets/images/camera1.jpeg",
			Description: "this is a camera",
		},
		{
			Name:        "camera2",
			ID:          "2",
			Img:         "public/assets/images/camera2.jpeg",
			Description: "this is a camera",
		},
		{
			Name:        "camera3",
			ID:          "3",
			Img:         "public/assets/images/camera3.jpeg",
			Description: "this is a camera",
		},
	}

}

func hello() {}
