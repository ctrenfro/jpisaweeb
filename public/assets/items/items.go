package items

type Items struct {
	Name string
	ID   string
	Img  string
}

func (*Items) returnItems() []Items {

	items := []Items{
		{
			Name: "camera1",
			ID:   "1",
			Img:  "../images/camera1.jpeg",
		},
		{
			Name: "camera2",
			ID:   "2",
			Img:  "camera2.jpeg",
		},
		{
			Name: "camera3",
			ID:   "3",
			Img:  "camera3.jpeg",
		},
	}

	return items
}
