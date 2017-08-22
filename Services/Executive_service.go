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
func InsertHomeDetails(w http.ResponseWriter, r *http.Request) {
	fmt.Println("inside GetImage")
	var ImageurlFinal []string
	var Imageurl string
	var Userdata Model.HomeInsert

	r.ParseMultipartForm(32 << 20)
	r.ParseForm()

	Userdata.Housename = r.Form["housename"][0]
	Userdata.Adress1 = r.Form["address1"][0]
	Userdata.Adress2 = r.Form["address2"][0]
	fmt.Println("********addess2********", Userdata.Adress2)
	Userdata.City = r.Form["city"][0]
	Userdata.Pin = r.Form["pin"][0]

	Userdata.State = r.Form["state"][0]
	Userdata.Country = "India"

	Userdata.Phone_number = r.Form["phonenumber"][0]
	Userdata.Month_rent = r.Form["monthrent"][0]
	Userdata.Tenant_type = r.Form["gettenant"][0]

	fmt.Println("***************", Userdata.Tenant_type)
	Userdata.Booking_type = r.Form["getbook"][0]
	Userdata.House_type = r.Form["gethouse"][0]
	Userdata.Furnish_type = r.Form["getfurnished"][0]
	Userdata.Distance = r.Form["getdistance"][0]
	fmt.Println("********Distance********", Userdata.Distance)

	Userdata.Bhk = r.Form["bhk"][0]
	Userdata.Bed = r.Form["bed"][0]
	Userdata.Secutity_deposit = r.Form["security"][0]
	Userdata.Listing = r.Form["listing"][0]
	//Userdata.Amenities = r.Form["amenities"][0]
	Amin := r.Form["groupsolo"]
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

	Userdata.Description = r.Form["description"][0]

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
		fmt.Println("handler", handler)
		defer file.Close()

		Imageurl = "Houseimage/" + Userdata.Housename + "/" + handler.Filename

		f, err := os.OpenFile(Imageurl, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()

		io.Copy(f, file)

	}
	ImageurlFinal = append(ImageurlFinal, "http://localhost:8083/"+Imageurl)
	fmt.Println("imageurl", ImageurlFinal)

	SendEmail(Userdata, ImageurlFinal)

	Db.Inserthome_DB(Userdata, ImageurlFinal)

}

/*
type HomeInsert struct {
	Housename        string
	Adress1          string
	Adress2          string
	Cityname         string
	State            string
	Country          string
	Pin              string
	Phone_number     string
	Month_rent       string
	Tenant_type      string
	Booking_type     string
	House_type       string
	Bhk              string
	Bed              string
	Distance         string
	Furnish_type     string
	Secutity_deposit string
	Listing          string
	Amenities        string
	Description      string
}
*/
func SendEmail(Userdata Model.HomeInsert, ImageurlFinal []string) {
	var ToString = "sandhiyabalakrishnan6@gmail.cpm"

	auth := smtp.PlainAuth("", "rentmatics@gmail.com", "RENTMATICS2017", "smtp.gmail.com")
	var receiptent = "To:" + ToString + "\r\n"
	to := []string{ToString}
	msg := []byte(receiptent + "Subject: RENTMATICS NOTIFICATION!\n" + "Phone Number : " + Userdata.Phone_number + "\n" + "Monthly Rent : " + Userdata.Month_rent + "\r\n" + "Tenent Type : " + Userdata.Tenant_type + "\r\n" + "Booking Type : " + Userdata.Booking_type + "\r\n" + "House Type : " + Userdata.House_type + "\r\n" + "BHK : " + Userdata.Bhk + "\r\n" + "No of Beds : " + Userdata.Bed + "\r\n" + "Distance : " + Userdata.Distance + "\r\n" + "Furnish Type : " + Userdata.Furnish_type + "\r\n" + "Security Deposit : " + Userdata.Secutity_deposit + "\r\n" + "Listing : " + Userdata.Listing + "\r\n" + "Amenities : " + Userdata.Amenities + "\r\n" + "Description : " + Userdata.Description + "\r\n" + "Currently We are working under this area,Rentmatics comre to you as soon!\r\n")

	fmt.Println("msg body", msg)
	//			+"Monthly Rent : "+Userdata.Month_rent+"\r\n"
	//			+"Tenent Type : "+Userdata.Tenant_type+"\r\n"
	//			+"Booking Type : "+Userdata.Booking_type+"\r\n"
	//			+"House Type : "+Userdata.House_type+"\r\n"
	//			+"BHK : "+Userdata.Bhk+"\r\n"
	//			+"No of Beds : "+Userdata.Bed+"\r\n"
	//			+"Distance : "+Userdata.Distance+"\r\n"+"Furnish Type : "+Userdata.Furnish_type+"\r\n"+"Security Deposit : "+Userdata.Secutity_deposit+"\r\n"
	//
	//
	//			+"Listing : "+Userdata.Listing+"\r\n"+"Amenities : "+Userdata.Amenities+"\r\n"+"Description : "+Userdata.Description+"\r\n"+"Currently We are working under this area,Rentmatics comre to you as soon!\r\n")
	//
	//

	//

	//	"House Name : "+Userdata.Housename+"\n"+"Address : "+Userdata.Adress1+" "+Userdata.Adress2+" "+Userdata.Cityname+" "+Userdata.State+" "+Userdata.Country+" "+Userdata.Pin+" "+"\n"

	err := smtp.SendMail("smtp.gmail.com:587", auth, "rentmatics@gmail.com", to, msg)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("finished mailing")
}
