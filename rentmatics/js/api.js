// <script type="text/javascript">

// // Cookies
//    var loginidd="";
//    var cooki=$.cookie('RentmaticsCookie');
//    $(document).ready(function(){
//       if(cooki==""||cooki==null||cooki==undefined||cooki=="NODATA"){
//             document.getElementById("Userdetails").style.visibility = "hidden"; 
//             document.getElementById("loginhide").style.visibility = "visible"; 
//             document.getElementById("signuphide").style.visibility = "visible"; 
//       }else{
             	
//           document.getElementById("loginhide").style.display = "none"; 
//           document.getElementById("signuphide").style.display = "none"; 
//           $("#login").html(cooki);
//           document.getElementById("Userdetails").style.visibility = "visible"; 
//           document.getElementById("wishlist").style.visibility = "visible"; 
//           document.getElementById("logout").style.visibility = "visible";
//           $("#wish").html("Wishlist");
//           }
//          });
         

// //Login with Rentmatics user
// function Login(){
//    var loginuser=document.getElementById('uname').value;
//    var loginpass=document.getElementById('upass').value;
//    var Rentmatics ={
//          	"Username": loginuser,
//          	"Password":loginpass,
// 	 };
               
//  $.ajax({
//         type: "POST",
//         url: "http://localhost:8083/Login",
//         contentType:"application/json; charset=utf-8",
//         data: JSON.stringify(Rentmatics),
         
//            success: function(responsee){
//            var searchout = JSON.parse(responsee);
//  if (searchout.Status == "Success"){
// 			loginidd=searchout.Loginid;
//             $("#login").html(loginidd);
//             document.getElementById("loginhide").style.display = "none"; 
//             document.getElementById("signuphide").style.display = "none"; 
//             document.getElementById("Userdetails").style.visibility = "visible"; 
//  }else{
// 			alert("Incorect Username and Password")
//   }}
// });
// }
      
//  //Signup   For Rentmatics User   
// function SignupRent(){
//  	var signuser=document.getElementById('Susername').value;
//     var signpass=document.getElementById('cpassword').value;
//     var signemail=document.getElementById('Semail').value;
//     var Rentmatics ={
//          	"Id ": "",      
//          	"Username": signuser,
//          	"Password":signpass,
//          	"Loginid":  signemail, 
//          	"Logintype": "Rentmatics"
              
//               };
               
         
//              $.ajax({
//                type: "POST",
//                url: "http://localhost:8083/User",
//                contentType:"application/json; charset=utf-8",
//                data: JSON.stringify(Rentmatics),
         
//             success: function(responsee){
//            var searchout = JSON.parse(responsee);
//            var uservalue=searchout.Username;
//            loginidd=searchout.Loginid;
         
          
//            $("#login").html(loginidd);
//           document.getElementById("loginhide").style.display = "none"; 
//           document.getElementById("signuphide").style.display = "none"; 
//           document.getElementById("Userdetails").style.visibility = "visible"; 
         
//         }
//           });
//          }


//  // Redirect to wishlist page  
// function Getwish(){
//             url = 'wishlist.html?loginid='+loginidd;
//          	document.location.href = url;
         				
//           }   
         

//    //Google Login Code
         
// 					var auth2; // The Sign-In object.
//  					var googleUser; // The current user.
// 					var appStart = function() {
//          			gapi.load('auth2', initSigninV2);
//          			};
//          			var initSigninV2 = function() {
//          			  auth2 = gapi.auth2.init({
//          			      client_id: 'CLIENT_ID.apps.googleusercontent.com',
//          			      scope: 'profile'
//          			  });
         
         			  
//          			  auth2.isSignedIn.listen(signinChanged);
//         			  auth2.currentUser.listen(userChanged);
         
//          			  if (auth2.isSignedIn.get() == true) {
//          			    auth2.signIn();
//          			  }
// 					  refreshValues();
//          			};
         
         			
//          			var signinChanged = function (val) {
//          			  console.log('Signin state changed to ', val);
         			  
//          			};
         
//          			var userChanged = function (user) {
//          			  console.log('User now: ', user);
//          			  googleUser = user;
//          			  updateGoogleUser();
         			 
//          			};
         
         			
//          			var updateGoogleUser = function () {
//          			  if (googleUser) {
//          			  	console.log("got google user");
//          			  	console.log(googleUser)
         			    
//          			  } else {
//          			  	console.log("did nit get google user");
         			  
//          			  }
//          			};
         
         			
//          			var refreshValues = function() {
//          			  if (auth2){
//          			    console.log('Refreshing values...');
         
//          			    googleUser = auth2.currentUser.get();
//          				console.log(googleUser)
//          				updateGoogleUser();
//          			  }
//          			}
//          			var email="";


// //Onsignin Google user
//   function onSignIn(googleUser) {
//          					alert("inside google")
//          			  var profile = googleUser.getBasicProfile();
//          			  console.log('ID: ' + profile.getId()); // Do not send to your backend! Use an ID token instead.
//          			  console.log('Name: ' + profile.getName());
//          			  console.log('Image URL: ' + profile.getImageUrl());
//          			  console.log('Email: ' + profile.getEmail()); 
         
//          			  // This is null if the 'email' scope is not present.
         			
//           				var Id = profile.getId(); // Do not send to your backend! Use an ID token instead.
//          			 var name = profile.getName();
//          			// var iamprofile.getImageUrl());
//          			   email = profile.getEmail();
         
         
         
//          var Gmail ={
//          	"Id ": Id,      
//          	"Username": name,
//          	"Password":  "",
//          	"Loginid":  email, 
//          	"Logintype": "Gmail"
              
//               };
               
         
//              $.ajax({
//                type: "POST",
//                url: "http://localhost:8083/User",
//                contentType:"application/json; charset=utf-8",
//                data: JSON.stringify(Gmail),
         
//             success: function(responsee){
//             	var searchout = JSON.parse(responsee);
//           		var uservalue=searchout.Username;
//          $("#login").html(uservalue);
//          document.getElementById("loginhide").style.display = "none"; 
//           document.getElementById("signuphide").style.display = "none"; 
//           document.getElementById("Userdetails").style.visibility = "visible"; 
         
         
//           }
//           });
//          }
         
//     //Signout     
//      function signOut() {
//              var auth2 = gapi.auth2.getAuthInstance();
//              auth2.signOut().then(function () {
//          	 console.log('User signed out.');
//          			    });
//          			  }
         		
//            window.fbAsyncInit = function() {
//              FB.init({
//                appId      : '837076093106264',
//                xfbml      : true,
//                version    : 'v2.9'
//              });
//              FB.AppEvents.logPageView();
//            };
         
//            (function(d, s, id){
//               var js, fjs = d.getElementsByTagName(s)[0];
//               if (d.getElementById(id)) {return;}
//               js = d.createElement(s); js.id = id;
//               //js.src = "//connect.facebook.net/en_US/sdk.js";
//               js.src = "//connect.facebook.net/en_US/sdk.js#xfbml=1&version=v2.9&appId=1243386559017640";
//               fjs.parentNode.insertBefore(js, fjs);
//             }(document, 'script', 'facebook-jssdk'));
         
//           // This is called with the results from from FB.getLoginStatus().
//            function statusChangeCallback(response) {
//              console.log('statusChangeCallback');
//              console.log(response);
            
//              if (response.status === 'connected') {
//                // Logged into your app and Facebook.
//                testAPI();
//              } else {
             
               
//              }
//            }
    
//     //CheckLoginstate     
//  function checkLoginState() {
//              FB.getLoginStatus(function(response) {
//                statusChangeCallback(response);
//              });
//            }
         
         
//            function testAPI() {
//              console.log('Welcome!  Fetching your information.... ');
//              FB.api('/me', function(response) {
//                console.log('Successful login for: ' + response.name);
//                console.log(response);
//                document.getElementById('status').innerHTML =
//                  'Thanks for logging in, ' + response.name + '!';
//              });
//            }




//  //Logout          
         
//        function logout(){
         
//          	$.ajax({
//            type: "GET",
//            url: "http://localhost:8083/Logout",
         
//          success: function(responsee){
//           document.getElementById("Userdetails").style.visibility = "hidden"; 
//          document.getElementById("wishlist").style.visibility = "hidden"; 
//          document.getElementById("logout").style.visibility = "hidden"; 
//          document.getElementById("loginhide").style.visibility = "visible";
//          document.getElementById("signuphide").style.visibility = "visible";  
//          }
         
//          });
         	
//          }
         
         
          
//       </script>