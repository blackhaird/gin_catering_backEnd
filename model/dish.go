package model

type Dish struct {
	DishNo    int     `json:"dish_no"`
	DishName  string  `json:"dish_name"`
	DishPrice float64 `json:"dish_price"`
	DishTaste string  `json:"dish_taste"`
}

type PostDish struct {
	PostId string `json:"post_id"`
	Dish   []Dish `json:"data"`
}
