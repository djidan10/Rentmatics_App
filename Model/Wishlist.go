package Model

type Wishlist struct {
	Id      int
	Homeid  int
	Loginid string
}

type WishlistDelete struct {
	Homeid  int
	Loginid string
}
