{{define "userprofile"}}
<!DOCTYPE html>
<html lang="en">

  <head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <title>Artistic - User Profile</title>

    <!-- Bootstrap -->
  <!--  <link href="css/bootstrap.min.css" rel="stylesheet"> -->
   <link href="/static/css/bootstrap.min.css" rel="stylesheet"> 

    <!-- HTML5 Shim and Respond.js IE8 support of HTML5 elements and media queries -->
    <!-- WARNING: Respond.js doesn't work if you view the page via file:// -->
    <!--[if lt IE 9]>
      <script src="https://oss.maxcdn.com/libs/html5shiv/3.7.0/html5shiv.js"></script>
      <script src="https://oss.maxcdn.com/libs/respond.js/1.4.2/respond.min.js"></script>
    <![endif]-->

    <!-- custom CSS, additional to CSS -->
    <link href="/static/css/custom.css" rel="stylesheet">
  </head>

  <body>
    {{template "navbar" .User.Username}}

    <div class="container-fluid">

    <div class="row">

        <div class="col-md-2" id="menu">
            <h1 id="menu-header"></h1>

            {{template "accordion"}}
        </div>

        <div class="col-md-10" id="data-list">
        {{if eq .Cmd "view"}} 
            <h1 id="data-list-header">View User Profile</h1>
            {{template "user-profile-view" .User}}
        {{else if eq .Cmd "modify"}}
            <h1 id="data-list-header">Modify User Profile</h1>
            {{template "user-profile-modify" .User}}
        {{else if eq .Cmd "changepwd"}} 
            <h1 id="data-list-header">Change User Password</h1>
            {{template "user-profile-change-passwd" .User}}
        {{else if eq .Cmd ""}} 
            <h1 id="data-list-header">View User Profile</h1>
            {{template "user-profile-view" .User}}
        {{end}}
        </div>

     </div> <!-- row -->

    </div> <!-- container fluid -->

<!-- jQuery (necessary for Bootstrap's JavaScript plugins) -->
<!--<script 
        src="https://ajax.googleapis.com/ajax/libs/jquery/2.1.0/jquery.min.js">
    </script> -->
    <script  src="/static/js/jquery.min.js"></script>

<!-- Include all compiled plugins, or include individual files as needed -->
    <script src="/static/js/bootstrap.min.js"></script>

<!-- custom JS code -->
    <script src="/static/js/artistic.js"></script>
    <script>

    // when page is ready...
    $(document).ready( function() {

        // parse the JSON data...
        {{/*var data = JSON.parse({{.Json}}); */}}

    });

    </script>

  </body>
</html>
{{end}}

{{define "user-profile-view"}}
<div class="container-fluid" id="view-user-table-div">
    <div class="row">
    <table id="view-user-table" class="table table-hover">
    <tbody>
        <tr> <td class="col-md-2"><b>Username<b/></td>
             <td class="col-md-10">{{.Username}}</td>   </tr>
        <tr> <td class="col-md-2"><b>Password</b></td>
             <td class="col-md-10">{{.Password}}</td>   </tr>
        <tr> <td class="col-md-2"><b>Role</b></td>
             <td class="col-md-10">{{.Role}}</td>       </tr>
        <tr> <td class="col-md-2"><b>Full Name</b></td>
             <td class="col-md-10">{{.Name}}</td>       </tr>
        <tr> <td class="col-md-2"><b>Email Address</b></td>
             <td class="col-md-10">{{.Email}}</td>      </tr>
    </tbody>
    </table>
    </div>

    <div class="row">
        <div class="col-md-4">

    <a type="button" class="btn btn-primary" 
       onclick="return rerouteUsingGet('userprofile', 'modify', {{.Id}});">
       Modify Profile</a>
            
    <a type="button" class="btn btn-danger" 
       onclick="return rerouteUsingGet('userprofile', 'changepwd', {{.Id}});">
    Change Password </a>
        </div>
    </div>
</div>
{{end}}

{{define "user-profile-modify"}}
<div class="container-fluid" id="modify-user-form-div">
    <form class="form-vertical" role="form" method="post"
                                id="user-modify-form">
    <fieldset>

    <div class="row">
    <div class="form-group"> 
        <label for="username" class="col-md-2 control-label">Username</label>
        <div class="col-md-6">
        <input type="text" class="form-control" id="username"
               name="username" value="{{.Username}}" readonly></input>
        </div>
    </div>
    </div>

    <div class="row">
    <div class="form-group"> 
        <label for="password" class="col-md-2 control-label">Password</label> 
        <div class="col-md-6">
        <input type="text" class="form-control" id="password"   
                           name="password" value="{{.Password}}" readonly>
        </input>
        </div>
    </div>
    </div>

    <div class="row">
    <div class="form-group"> 
        <label for="role" class="col-md-2 control-label">User Role</label>
        <div class="col-md-6">
        <input type="text" class="form-control" id="role"   
                           name="role" value="{{.Role}}" readonly>
        </input>
        </div>
    </div>
    </div>

    <div class="row">
    <div class="form-group"> 
        <label for="fullname" class="col-md-2 control-label">Full Name</label>
        <div class="col-md-6">
        <input type="text" class="form-control" id="fullname"
               name="fullname" value="{{.Name}}"></input>
        </div>
    </div>
    </div>

    <div class="row">
    <div class="form-group"> 
        <label for="email" class="col-md-2 control-label">E-mail Address</label>
        <div class="col-md-6">
        <input type="email" class="form-control" id="email"
               name="email" value="{{.Email}}"></input>
        </div>
    </div>
    </div>

    <div class="row">&nbsp;<!-- empty row --></div>

    <div class="row">
    <div class="form-group">
        <div class="col-md-4">
        <button class="btn btn-primary" type="submit" 
                id="user-submit">Modify User</button>

        <button class="btn btn-danger" type="button" 
          onclick="return rerouteUsingGet('userprofile', 'changepwd', {{.Id}});"
                id="user-changepwd">Change Password</button>
        </div>

        <div class="col-md-1 col-md-offset-3">
    <a type="button" class="btn btn-primary" href="/userprofile">
        <span class="glyphicon glyphicon-arrow-left"></span>&nbsp;&nbsp;Back
    </a>
        </div>
    </div> <!-- form-group -->
    </div> <!-- row -->

    </fieldset>
    </form>
</div> <!-- container-fluid -->
{{end}}

{{define "user-profile-change-passwd"}}
<div class="container-fluid" id="change-pwd-form-div">
    <form class="form-vertical" role="form" method="post"
                onsubmit="validatePasswordChange();"
                id="change-pwd-form"> 
        <fieldset>

    <div class="row">
        <div class="form-group">
            <label for="oldpassword" 
                   class="col-md-2 control-label">Old Password</label>
            <div class="col-md-6">
            <input type="password" class="form-control" id="oldpassword"
                   name="oldpassword" value="" 
                   placeholder="enter old password" required>
            </input>
            </div>
        </div>
    </div>
 
    <div class="row">
        <div class="form-group">
            <label for="newpassword" 
                   class="col-md-2 control-label">New Password</label>
            <div class="col-md-6">
            <input type="password" class="form-control" id="newpassword"
                   name="newpassword" value="" 
                   placeholder="enter new password" required>
            </input>
            </div>
        </div>
    </div>

    <div class="row">
        <div class="form-group">
            <label for="newpassword2" 
                   class="col-md-2 control-label">Retype New Password</label>
            <div class="col-md-6">
            <input type="password" class="form-control" id="newpassword2"
                   name="newpassword2" value="" 
                   placeholder="retype new password" required>
            </input>
            </div>
        </div>
    </div>

    <div class="row">&nbsp;</div>

    <div class="row">
        <div class="form-group">
            <div class="col-md-4">
            <button class="btn btn-primary" type="button" 
                    onclick="return postForm('change-pwd-form', 
                                    createURL('changepwd', 'userprofile', {{.Id}}));"
                    id="pwd-submit">Change Password</button>
            <!--
            <button class="btn btn-primary" type="submit" 
                    id="pwd-submit">Change Password</button>
             -->
            <button class="btn btn-default" type="reset">Clear</button>
            </div>
            <div class="col-md-1 col-md-offset-3">
            <a type="button" class="btn btn-primary" href="/userprofile">
        <span class="glyphicon glyphicon-arrow-left"></span>&nbsp;&nbsp;Back
            </a>
            </div>
        </div>
    </div>

        </fieldset>
    </form>
</div>
{{end}}
