package Model

type Executive struct {
	Homestruct Home
	images     []string
}

type LoginData struct {
	Id       string
	Username string
	Password string
	Role     string
}

type Login struct {
	Username string
	Password string
}

type Executivedetails struct {
	Executive_id       int
	First_Name         string
	Last_Name          string
	Email_Id           string
	Contact            string
	Alernate_Contact   string
	DOB                string
	Permanent_Address1 string
	Permanent_Address2 string
	Permanent_Area     string
	Permanent_City     string
	Permanent_Pin      string
	Executive_img      string
	Pan_Card           string
	Aadhar_Card        string
	Voter_Card         string
}

type Executivesend struct {
	ExecutiveData Executivedetails
	HomeData      []RentSendTenant
}

type RentSendTenant struct {
	//Cityname       string
	RentFullStruct Home
	RentFullimages []Home_images
	Tenantdetails  Tenant
}
