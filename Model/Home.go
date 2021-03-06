package Model

type Home struct {
	Id               int
	Housename        string
	Executive_id     int
	Ownerid          int
	Adress1          string
	Adress2          string
	City             string
	District         string
	State            string
	Country          string
	Pin              string
	Phone_number     string
	Month_rent       float64
	Bed_rent         float64
	Bhk_rent         float64
	Tenant_type      string
	Booking_type     string
	House_type       string
	Bhk              int
	Bed              int
	Avail_bed        int
	Avail_room       int
	Booked_bed       int
	Booked_bhk       int
	Distance         string
	Furnish_type     string
	Secutity_deposit float64
	Listing          string
	Amenities        string
	Description      string
	Latitude         float64
	Longitude        float64
	Liked            bool
	Squarefeet       string
	Totalfloors      string
	Facing           string
	Parking          string
}

type HomeInsert struct {
	//Id               int
	Housename        string
	Houseownername   string
	ExecutiveID      string
	Ownerid          string
	Adress1          string
	Adress2          string
	City             string
	District         string
	State            string
	Country          string
	Pin              string
	Phone_number     string
	Email            string
	Secutity_deposit string
	Month_rent       string
	Bed_rent         string
	Bhk_rent         string
	Bhk              string
	Bed              string
	Availble_Bed     string
	Avail_room       string
	Booked_bed       string
	Booked_bhk       string
	Tenant_type      string
	Booking_type     string
	House_type       string
	Distance         string
	Furnish_type     string
	Listing          string
	Amenities        string
	Description      string
	Latitude         string
	Longitude        string
	Liked            string
	Squarefeet       string
	Totalfloors      string
	Facing           string
	Parking          string
}

type Home_images struct {
	Home_id     int64
	Picture_id  int
	Picture_url string
}

type RentSend struct {
	//Cityname       string
	RentFullStruct Home
	RentFullimages []Home_images
}

type RentSendFav struct {
	RentFullStruct Home
	//RentCity       string
	RentFullimages []Home_images
}
type RentSendall struct {
	Allhomes []RentSend
}
type Amenties struct {
	Amenties string
	Url      string
}

type Home_single struct {
	Cityname     string
	Home_Data    Home
	Home_images  []Home_images
	Amenties     []Amenties
	Liked        bool
	Ownerdetails Owner
}
type Filter struct {
	F_Loginid        string
	F_Monthrent_Min  int
	F_Monthrent_Max  int
	F_Tenant_type    string
	F_Bookingtype    string
	F_Housetype      string
	F_Furnished_type string
	F_Distance       string
	F_Bhk            int
	F_Bed            int
}
type APIError struct {
	Error struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}
