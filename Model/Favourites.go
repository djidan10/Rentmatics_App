package Model

type Favourites struct {
	Favid  int
	UserId int
	Homeid int
}

type GetFavourites struct {
	UserId int
	Homeid []int
}
