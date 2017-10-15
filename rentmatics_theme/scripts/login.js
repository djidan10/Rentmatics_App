
var loginidvalue =""
  loginidvalue = localStorage.getItem("LoginID");

//Login
function Login() {
    var loginuser = document.getElementById('uname').value;
    var loginpass = document.getElementById('upass').value;
    var Rentmatics = {
        "Username": loginuser,
        "Password": loginpass,


    };


    $.ajax({
        type: "POST",
        url: "http://localhost:8083/Login",
        contentType: "application/json; charset=utf-8",
        data: JSON.stringify(Rentmatics),

        success: function(responsee) {
            var searchout = JSON.parse(responsee);
            if (searchout.Status == "Success") {
                loginidd = searchout.Loginid;
                $("#login").html("MyAccount");
                document.getElementById("loginhiden").style.display = "none";
                document.getElementById("account").style.visibility = "visible";
                   window.history.back();
            } else {
                alert("Incorect Username and Password")

            } }
    });
}





// Signup-Register the Rentmatics user
function SignupRent() {
    var signuser = document.getElementById('Susername').value;
    var signpass = document.getElementById('cpassword').value;
    var signemail = document.getElementById('Semail').value;
    var Rentmatics = {
        "Id ": "",
        "Username": signuser,
        "Password": signpass,
        "Loginid": signemail,
        "Logintype": "Rentmatics"

    };


    $.ajax({
        type: "POST",
        url: "http://localhost:8083/User",
        contentType: "application/json; charset=utf-8",
        data: JSON.stringify(Rentmatics),

        success: function(responsee) {
            var searchout = JSON.parse(responsee);
            var uservalue = searchout.Username;
            loginidd = searchout.Loginid;
            $("#login").html(loginidd);
            
            document.getElementById("loginhiden").style.display = "none";
            document.getElementById("account").style.visibility = "visible";
              window.history.back();
        }
    });
}



//Redirect to wishlist page

function Getwish() {
    alert("inside wish " + loginidd);
    url = "wishlist.html";
    document.location.href = url;

}




   



//Post Data to Next page
function Gethomes() {

   // alert(loginidvalue);
var url =""
    var cities = document.getElementById('tags').value;
  
 
  var partsOfStr = cities.split(',');
   
   if (partsOfStr  == null || partsOfStr  == ""){
     url = 'listings-half-map-grid-compact.html' 
   }else{
    
      url = 'listings-half-map-grid-compact.html?Id=' + partsOfStr[0];
   localStorage.setItem("Getcity", partsOfStr[0]);

   }
  

     document.location.href = url;
   
}

function Payalert(){
    alert("Pay Myrent Comminng Soon!...please wait")
}




//For Google User-Signin
var auth2; // The Sign-In object.
var googleUser; // The current user.
var appStart = function() {
    gapi.load('auth2', initSigninV2);
};
var initSigninV2 = function() {
    auth2 = gapi.auth2.init({
        client_id: 'CLIENT_ID.apps.googleusercontent.com',
        scope: 'profile'
    });


    auth2.isSignedIn.listen(signinChanged);
    auth2.currentUser.listen(userChanged);

    if (auth2.isSignedIn.get() == true) {
        auth2.signIn();
    }
    refreshValues();
};


var signinChanged = function(val) {
    console.log('Signin state changed to ', val);

};

var userChanged = function(user) {
    console.log('User now: ', user);
    googleUser = user;
    updateGoogleUser();

};


var updateGoogleUser = function() {
    if (googleUser) {
        console.log("got google user");
        console.log(googleUser)

    } else {
        console.log("did nit get google user");

    }
};


var refreshValues = function() {
    if (auth2) {
        console.log('Refreshing values...');

        googleUser = auth2.currentUser.get();
        console.log(googleUser)
        updateGoogleUser();
    }
}

// function GsignOut() {
//     var auth2 = gapi.auth2.getAuthInstance();
//     auth2.signOut().then(function() {
//         console.log('User signed out.');

//     });
// }



var delete_cookie = function(name) {
    document.cookie = name + '=;expires=Thu, 01 Jan 1970 00:00:01 GMT;';
};
var email = "";

function onSignIn(googleUser) {
    
    var profile = googleUser.getBasicProfile();
    console.log('ID: ' + profile.getId()); // Do not send to your backend! Use an ID token instead.
    console.log('Name: ' + profile.getName());
    console.log('Image URL: ' + profile.getImageUrl());
    console.log('Email: ' + profile.getEmail());

    // This is null if the 'email' scope is not present.

    var Id = profile.getId(); // Do not send to your backend! Use an ID token instead.
    var name = profile.getName();
    // var iamprofile.getImageUrl());
    email = profile.getEmail();



    var Gmail = {
        "Id ": Id,
        "Username": name,
        "Password": "",
        "Loginid": email,
        "Logintype": "Gmail"

    };


    $.ajax({
        type: "POST",
        url: "http://localhost:8083/User",
        contentType: "application/json; charset=utf-8",
        data: JSON.stringify(Gmail),

        success: function(responsee) {
            var searchout = JSON.parse(responsee);



            console.log("response");
            console.log(searchout)
            var uservalue = searchout.Username;

            // $("#Userdetails").html(uservalue);

            $("#login").html(uservalue);



            document.getElementById("loginhide").style.display = "none";
            document.getElementById("signuphide").style.display = "none";
            document.getElementById("Userdetails").style.visibility = "visible";
              window.history.back();



        }

    });
}

function signOut() {
    var auth2 = gapi.auth2.getAuthInstance();
    auth2.signOut().then(function() {
        console.log('User signed out.');
    });
}















