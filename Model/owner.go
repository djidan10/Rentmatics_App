package Model

type Owner struct {
	Owner_id           int
	Homeid             int
	First_Name         string
	Last_Name          string
	Rusername          string
	Email_Id           string
	Contact            string
	Alernate_Contact   string
	DOB                string
	Permanent_Address1 string
	Permanent_Address2 string
	Permanent_Area     string
	Permanent_City     string
	Permanent_Pin      string
	Tenant_img         string
	Pan_Card           string
	Aadhar_Card        string
	Voter_Card         string
	Aggrement          string
}

type Ownersend struct {
	OwnerData Owner
	HomeData  []RentSend
}

type Ownercomplaints struct {
	Ownercom_id   int
	Ownerid       int
	Homeid        int
	ComplaintDate string
	Message       string
	Status        string
	Approvedate   string
}

type RentSendOwner struct {
	RentFullStruct Home
	RentFullimages []Home_images
	Ownerdetails   Ownersend
}

type OwnertoExecutive struct {
	Name       string
	Email      string
	Phone      string
	Tenanttype string
	Address    string
}

type OwnerRequest struct {
	OwnerRequestid int
	OwnerId        int
	Ownername      string
	Owneremail     string
	Ownerhomeid    int
	Description    string
	Status         string
	RequestDate    string
	ApproveDate    string
}

type OwnerVisit struct {
	Ownername   string
	Owneremail  string
	Ownerhomeid int
	VisitDate   string
	Description string
	Status      string
}
type Ownermessage struct {
	Fromemail   string
	Toemail     string
	Subject     string
	Homeid      int
	Towhom      string
	Description string
}
