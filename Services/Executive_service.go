package Services

import (
	Db "Rentmatics_App/Common/DB/Mysql"
	Model "Rentmatics_App/Model"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/smtp"
	"os"
	"path/filepath"
	"strconv"
)

//Get All Executive Details
func GetExecutive(w http.ResponseWriter, r *http.Request) {
	Data := Db.GetAllExecutive()
	Senddata, err := json.Marshal(Data)
	if err != nil {
		log.Error("Error on Fetching All Executive Details", err)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Orgin", "*")
	w.Write(Senddata)

}

type Executivesingle struct {
	Executiveid int
}

//Get Particular Executive Details
func GetsingleExecutive(w http.ResponseWriter, r *http.Request) {

	var GetExecutive Executivesingle
	err := json.NewDecoder(r.Body).Decode(&GetExecutive)
	if err != nil {
		log.Error("Error on Fetching All Executive Details", err)
	}
	Data := Db.GetSingleExecutive_Db(GetExecutive.Executiveid)
	Senddata, err := json.Marshal(Data)
	if err != nil {
		log.Error("Cannot mashal executive value from databse", err)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Orgin", "*")
	w.Write(Senddata)

}

//Auth For login executive
func Auth(w http.ResponseWriter, r *http.Request) {
	fmt.Println("inside Rent")

	var UserData Model.Login
	err := json.NewDecoder(r.Body).Decode(&UserData)
	if err != nil {
		fmt.Println("err", err)
	}

	UserAuth := Db.Checklogin_DB(UserData)
	if UserData.Password == UserAuth.Password {
		Senddata, err := json.Marshal(UserAuth)

		if err != nil {
			panic(err)
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Access-Control-Allow-Orgin", "*")

		w.Write(Senddata)

	} else {
		w.Write([]byte("Login Failed"))

	}
}

//Executive Upload
func InsertHomeDetails(w http.ResponseWriter, r *http.Request) {
	fmt.Println("inside GetImage")
	var ImageurlFinal []string
	var Imageurl string

	var Userdata Model.HomeInsert

	r.ParseMultipartForm(32 << 20)
	r.ParseForm()

	//  Userdata.Housename = r.Form["housename"][0]
	Userdata.Housename = r.Form["houseownername"][0]
	Userdata.Houseownername = r.Form["housename"][0]
	Userdata.ExecutiveID = r.Form["executiveid"][0]
	Userdata.Ownerid = r.Form["ownerid"][0]
	Userdata.Adress1 = r.Form["address1"][0]
	Userdata.Adress2 = r.Form["address2"][0]
	Userdata.City = r.Form["city"][0]
	Userdata.District = r.Form["district"][0]
	Userdata.State = r.Form["state"][0]
	Userdata.Country = r.Form["country"][0]
	Userdata.Pin = r.Form["pin"][0]
	Userdata.Phone_number = r.Form["phonenumber"][0]
	Userdata.Email = r.Form["email"][0]
	Userdata.Secutity_deposit = r.Form["securitydeposit"][0]
	Userdata.Month_rent = r.Form["monthlyrent"][0]
	Userdata.Bed_rent = r.Form["bedrent"][0]
	Userdata.Bhk_rent = r.Form["bhkrent"][0]
	Userdata.Bed = r.Form["bed"][0]
	Userdata.Bhk = r.Form["bhk"][0]
	Userdata.Availble_Bed = r.Form["availbed"][0]
	Userdata.Avail_room = r.Form["availbhk"][0]
	Userdata.Booked_bed = r.Form["bookbed"][0]
	Userdata.Booked_bhk = r.Form["bookbhk"][0]
	Userdata.Tenant_type = r.Form["tenant"][0]
	Userdata.Furnish_type = r.Form["furnishtype"][0]
	Userdata.Booking_type = r.Form["booking"][0]
	Userdata.House_type = r.Form["housetype"][0]
	Userdata.Distance = r.Form["distance"][0]
	Userdata.Latitude = r.Form["lalitude"][0]
	Userdata.Longitude = r.Form["longitude"][0]
	Userdata.Listing = r.Form["listing"][0]
	Userdata.Squarefeet = r.Form["sqft"][0]
	Userdata.Description = r.Form["description"][0]
	Userdata.Parking = r.Form["parking"][0]
	Userdata.Facing = r.Form["facing"][0]
	Userdata.Totalfloors = r.Form["tnf"][0]
	Userdata.Liked = "false"
	fmt.Println(Userdata.Housename, Userdata.Houseownername, Userdata.Adress1, Userdata.Adress2, Userdata.City)
	fmt.Println(Userdata.District, Userdata.State)

	//Userdata.Amenities = r.Form["amenities"][0]
	Amin := r.Form["groupamen"]
	fmt.Println(Amin)
	var Amenities string
	for k, v := range Amin {
		if k < len(Amin)-1 {
			Amenities = Amenities + v + ","

		} else {
			Amenities = Amenities + v
		}
	}
	fmt.Println(Amenities)
	Userdata.Amenities = Amenities

	//Insert image in Database
	//Create Directory
	newpath := filepath.Join("Houseimage/", Userdata.Housename, "/")
	os.MkdirAll(newpath, os.ModePerm)
	for i := 0; i <= 5; i++ {

		file, handler, err := r.FormFile("uploadfile" + strconv.Itoa(i))
		if err != nil {
			fmt.Println("data error image", err)
			return
		}

		defer file.Close()

		Imageurl = "Houseimage/" + Userdata.Housename + "/" + handler.Filename

		f, err := os.OpenFile(Imageurl, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()

		io.Copy(f, file)
		fmt.Println("Imageurl", Imageurl)
		ImageurlFinal = append(ImageurlFinal, "http://localhost:8083/"+Imageurl)

	}

	fmt.Println("imageurl", ImageurlFinal)

	SendEmail(Userdata, ImageurlFinal)
	Db.Inserthome_DB(Userdata, ImageurlFinal)

}

//Once Executive upload data Mail go to Rentmatics owner
func SendEmail(Userdata Model.HomeInsert, ImageurlFinal []string) {
	var ToString = "sandhiyabalakrishnan6@gmail.cpm"

	auth := smtp.PlainAuth("", "rentmatics@gmail.com", "RENTMATICS2017", "smtp.gmail.com")
	var receiptent = "To:" + ToString + "\r\n"
	to := []string{ToString}
	msg := []byte(receiptent + "Subject: RENTMATICS NOTIFICATION!\n" + "Executive Sucessfully uploaded one home details please login to admin and see details!\r\n")
	log.Info(msg)
	err := smtp.SendMail("smtp.gmail.com:587", auth, "rentmatics@gmail.com", to, msg)

	if err != nil {
		log.Error("smtp Error please check executive mail part", err)
	}
	log.Info("Succesfully Mailed")
}

func Insertexecutive(w http.ResponseWriter, r *http.Request) {

	var InsExecutive Model.Executivedetails
	err := json.NewDecoder(r.Body).Decode(&InsExecutive)
	if err != nil {
		log.Error("Error: Tenant vacate ", err)
	}
	//Insert Vacate Details
	Db.InsertExecutive_Db(InsExecutive)

}
