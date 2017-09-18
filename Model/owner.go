package Model

type Owner struct {
	Owner_id           int
	Homeid             int
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
	Executiveid   int
	Tenantid      int
	Homeid        int
	Loginid       string
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
