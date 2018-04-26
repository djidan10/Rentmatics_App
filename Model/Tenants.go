package Model

type RentLogin struct {
	Username string
	Password string
}

type Loginsend struct {
	Role string
}

type Tenantsend struct {
	Tenantdetails        Tenant
	Homedetails          Home_single
	TenantPaymentdetails []Tenantpayment
	Executivedetail      Executivedetails
}

type Tenant struct {
	Tenant_id          int
	Homeid             int
	Executiveid        int
	Ownerid            int
	Loginid            string
	First_Name         string
	Last_Name          string
	Email_Id           string
	Contact            string
	Alernate_Contact   string
	DOB                string
	Tenant_Type        string
	Office_Address1    string
	Office_Address2    string
	Office_Area        string
	Office_City        string
	Office_Pin         string
	Office_Email       string
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

type Tenantpayment struct {
	Tenantpaymentid    int
	Tenant_id          int
	Home_id            int
	Owner_id           int
	Loginid            string
	Executiveid        int
	Lastmonth_paiddate string
	Duedate            string
	RentAmount         string
	Status             string
	PaymentDetails     string
	TransactionDetails string
}

type Complaints struct {
	Complaint_id          int
	Home_id               int
	Tenant_Id             int
	Complaint_raisedDate  string
	Raisedby              string
	Complaint_Description string
	Complaint_status      string
	Complaint_solvedDate  string
}
type Complaintsend struct {
	ComplaintsDetails Complaints
	Tenantdetails     Tenant
}

type Complaintsend1 struct {
	ComplaintsDetails []Complaints
	Tenantdetails     Tenant
}

type Vacate struct {
	Vacate_id        int
	Home_id          int
	Tenant_Id        int
	Vacate_priordate string
	Vacate_Date      string
	Vacate_reason    string
	Vacate_status    string
	ApproveDate      string
}

type Activity struct {
	Activity_id          int
	Home_id              int
	Tenant_Id            int
	Activity_Date        string
	Activity_Tittle      string
	Participation_count  int
	Activity_Description string
	Activity_Status      string
	ApproveDate          string
}

type Request struct {
	Request_id          int
	Home_id             int
	Tenant_Id           int
	Resquestername      string
	Request_Date        string
	Request_description string
	Raisedby            string
	Approve_Date        string
	Status              string
}

type Requestsend struct {
	Resquestdetails Request
	Tenantdetails   Tenant
}

type Complaintstatus struct {
	Complaintid string
	Status      string
	ApproveDate string
}

type Requeststatus struct {
	Requestid   string
	Status      string
	ApproveDate string
}
type ActivityUpdatestatus struct {
	Activityid  string
	Status      string
	ApproveDate string
}
type Vacatestatus struct {
	Vacateid    string
	Status      string
	ApproveDate string
}
type ResponseMessage struct {
	MessageId int
	Subject   string
	Homeid    int
	Message   string
}
