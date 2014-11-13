{{define "users"}}
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <title>Artistic - Administer Users</title>

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
                <h1 id="data-list-header">Users</h1>
                {{template "user-list" .Users}}
                <button type="button" class="btn btn-primary" onclick="rerouteUsingGet('user', 'insert', '');">
                Add New User
                </button>
            </div>
        </div> <!-- row -->
    </div> <!-- container fluid -->

    <!-- jQuery (necessary for Bootstrap's JavaScript plugins) -->
    <!--   <script src="https://ajax.googleapis.com/ajax/libs/jquery/2.1.0/jquery.min.js"></script> -->
    <script  src="/static/js/jquery.min.js"></script>

    <!-- Include all compiled plugins, or include individual files as needed -->
    <script src="/static/js/bootstrap.min.js"></script>

    <!-- custom JS code -->
    <script src="/static/js/artistic.js"></script>
</body>
</html>
{{end}}

{{define "user-list"}}
<table class="table table-striped table-hover" id="user-list-table">
    <thead>
        <tr>
            <th>#</th>
            <th>Username</th>
            <th>Name</th>
            <th>Role</th>
            <th>Email</th>
            <th>Actions</th>
        </tr>
    </thead>

    <tbody>
        {{range $index, $element := .}}
        {{ $id := add $index 1 }}
        <tr id="user-row-{{$id}}">
            <td>{{$id}}</td>
            <td>{{printf "%s" $element.Username}}</td>
            <td>{{printf "%s" $element.Name}}</td>
            <td>{{printf "%s" $element.Role}}</td>
            <td>{{printf "%s" $element.Email}}</td>
            <td>
                <a href="#" data-toggle="tooltip" data-placement="left" title="View user details" id="view-user-{{$id}}"
                           onclick="return rerouteUsingGet('user', 'view', {{$element.Id}});">
                    <span class="glyphicon glyphicon-eye-open"></span>
                </a>
                &nbsp;
                <a href="#" data-toggle="tooltip" data-placement="left" title="Edit user" id="edit-user-{{$id}}"
                            onclick="return rerouteUsingGet('user', 'modify', {{$element.Id}});">
                    <span class="glyphicon glyphicon-cog" ></span>
                </a>
                &nbsp;
                <a href="#" data-toggle="tooltip" data-placement="left" title="Delete user" id="delete-user-{{$id}}"
                            onclick="return rerouteUsingGet('user', 'delete', {{$element.Id}});">
                    <span class="glyphicon glyphicon-trash"></span>
                </a>
                &nbsp;
                <a href="#" data-toggle="tooltip" data-placement="left" title="Change user password" id="pwd-{{$id}}"
                            onclick="return rerouteUsingGet('user', 'changepwd', {{$element.Id}});">
                    <span class="glyphicon glyphicon-random"></span>
                </a>
            </td>
        </tr>
        {{end}}
    </tbody>
</table>
{{end}}

{{define "user"}}
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <title>Artistic - User</title>

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
                <h1 id="data-list-header">View User</h1>
            {{template "single-user-view" .UserProfile}}
        {{else if eq .Cmd "modify"}}
                <h1 id="data-list-header">Modify User</h1>
            {{template "single-user-modify" .UserProfile}}
        {{else if eq .Cmd "insert"}}
                <h1 id="data-list-header">Create New User</h1>
                <p>Please enter the data to create a new user.</p>
            {{template "user-create"}}
        {{else if eq .Cmd "changepwd"}} 
                <h1 id="data-list-header">Change User Password</h1>
            {{template "change-passwd" .UserProfile}}
        {{else if eq .Cmd ""}} 
                <h1 id="data-list-header">View User</h1>
            {{template "single-user-view" .UserProfile}}
        {{end}}
            </div>
        </div> <!-- row -->
    </div> <!-- container fluid -->

    <!-- jQuery (necessary for Bootstrap's JavaScript plugins) -->
    <!--<script src="https://ajax.googleapis.com/ajax/libs/jquery/2.1.0/jquery.min.js"> </script> -->
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

{{define "single-user-view"}}
<div class="container-fluid" id="view-user-table-div">
    <div class="row">
        <table id="view-user-table" class="table table-hover">
        <tbody>
            <tr> <td class="col-md-2"><b>Username<b/></td> <td class="col-md-10">{{.Username}}</td> </tr>
            <tr> <td class="col-md-2"><b>Password</b></td> <td class="col-md-10">{{.Password}}</td> </tr>
            <tr> <td class="col-md-2"><b>Role</b></td><td class="col-md-10">{{.Role}}</td> </tr>
            <tr> <td class="col-md-2"><b>Full Name</b></td><td class="col-md-10">{{.Name}}</td> </tr>
            <tr> <td class="col-md-2"><b>Email Address</b></td> <td class="col-md-10">{{.Email}}</td> </tr>
        </tbody>
        </table>
    </div>

    <div class="row"> <div class="col-md-1 col-md-offset-7">
        <a type="button" class="btn btn-primary" href="/users">
            <span class="glyphicon glyphicon-arrow-left"></span>&nbsp;&nbsp;Back
        </a>
    </div>
</div> <!-- container-fluid -->
{{end}}

{{define "single-user-modify"}}
<div class="container-fluid" id="modify-user-form-div">
    <form class="form-vertical" role="form" method="post" id="user-modify-form">
    <fieldset>

    <div class="row">
    <div class="form-group"> 
            <label for="username" class="col-md-2 control-label">Username</label>
        <div class="col-md-6">
            <input type="text" class="form-control" id="username" name="username" value="{{.Username}}" readonly></input>
        </div>
    </div>
    </div>

    <div class="row">
    <div class="form-group"> 
            <label for="password" class="col-md-2 control-label">Password</label> 
        <div class="col-md-6">
            <input type="text" class="form-control" id="password" name="password" value="{{.Password}}" readonly></input>
        </div>
    </div>
    </div>

    <div class="row">
    <div class="form-group"> 
        <label for="role" class="col-md-2 control-label">User Role</label>
        <div class="col-md-6">
            <select class="form-control" name="role" id="role">
        {{ $roles := allowedroles }}
        {{ $current := .Role}}
        {{range $role := $roles}}
            {{if eq $role $current}} 
                <option value="{{$role}}" selected>{{totitle $role}}</option>
            {{else}}
                <option value="{{$role}}">{{totitle $role}}</option>
            {{end}}
        {{end}}
            </select>
        </div>
    </div>
    </div>

    <div class="row">
    <div class="form-group"> 
        <label for="fullname" class="col-md-2 control-label">Full Name</label>
        <div class="col-md-6">
            <input type="text" class="form-control" id="fullname" name="fullname" value="{{.Name}}"></input>
        </div>
    </div>
    </div>

    <div class="row">
    <div class="form-group"> 
        <label for="email" class="col-md-2 control-label">E-mail Address</label>
        <div class="col-md-6">
            <input type="email" class="form-control" id="email" name="email" value="{{.Email}}"></input>
        </div>
    </div>
    </div>

    <div class="row">&nbsp;<!-- empty row --></div>

    <div class="row">
    <div class="form-group">
        <div class="col-md-4">
            <button class="btn btn-primary" type="submit" id="user-submit">Modify User</button>

            <button class="btn btn-danger" type="button" onclick="return rerouteUsingGet('user', 'changepwd', {{.Id}});"
                    id="user-changepwd">Change Password</button>
        </div>

        <div class="col-md-1 col-md-offset-3">
            <a type="button" class="btn btn-primary" href="/users">
                <span class="glyphicon glyphicon-arrow-left"></span>&nbsp;&nbsp;Back
            </a>
        </div>
    </div> <!-- form-group -->
    </div> <!-- row -->

    </fieldset>
    </form>
</div> <!-- container-fluid -->
{{end}}

{{define "user-create"}}
<div class="container-fluid" id="create-user-form-div">
    <form class="form-vertical" role="form" method="post" onsubmit="validateUserForm();" 
                                id="create-user-form" action="/user/insert/">
        <fieldset>

        <div class="row">
        <div class="form-group">
            <label for="username" class="col-md-2 control-label">Username</label>
            <div class="col-md-6">
                <input type="text" class="form-control" id="username" name="username" value="" placeholder="username" required>
                </input>
            </div>
        </div>
        </div>
        
        <div class="row">
        <div class="form-group">
            <label for="password" class="col-md-2 control-label">Password</label>
            <div class="col-md-6">
                <input type="password" class="form-control" id="password" name="password" value="" placeholder="password" required>
                </input>
            </div>
        </div>
        </div>

        <div class="row">
        <div class="form-group">
            <label for="password2" class="col-md-2 control-label">Retype Password</label>
            <div class="col-md-6">
                <input type="password" class="form-control" id="password2" name="password2" value="" 
                                       placeholder="retype password" required> </input>
            </div>
        </div>
        </div>

        <div class="row">
        <div class="form-group">
            <label for="role" class="col-md-2 control-label">User Role</label>
            <div class="col-md-6">
                <select class="form-control" name="role" id="role" required>
            {{ $roles := allowedroles }}
            {{range $role := $roles}}
                {{if eq $role "user"}}
                    <option value="{{$role}}" selected>{{totitle $role}}</option>
                {{else}}
                    <option value="{{$role}}">{{totitle $role}}</option>
                {{end}}
            {{end}}
                </select>
            </div>
        </div>
        </div>

        <div class="row">
        <div class="form-group">
            <label for="fullname" class="col-md-2 control-label">Full Name</label>
            <div class="col-md-6">
                <input type="text" class="form-control" name="fullname" id="fullname" placeholder="enter full name"></input>
            </div>
        </div>
        </div>

        <div class="row">
        <div class="form-group">
            <label for="email" class="col-md-2 control-label">Email Address</label>
            <div class="col-md-6">
                <input type="email" class="form-control" id="email" name="email" placeholder="e-mail address"></input>
            </div>
        </div>
        </div>

        <div class="row">&nbsp;<!-- empty row --></div>

        <div class="row">
        <div class="form-group">
            <div class="col-md-4">
                <button class="btn btn-primary" type="submit" id="user-submit">Create</button>
                <button class="btn btn-default" type="reset">Clear</button>
            </div>

            <div class="col-md-1 col-md-offset-3">
                <a type="button" class="btn btn-primary" href="/users">
                    <span class="glyphicon glyphicon-arrow-left"></span>&nbsp;&nbsp;Back
                </a>
            </div>
        </div>
        </div>

        </fieldset>
    </form>
</div> <!-- container-fluid -->
{{end}}

{{define "change-passwd"}}
<div class="container-fluid" id="change-pwd-form-div">
    <form class="form-vertical" role="form" id="change-pwd-form" method="post" onsubmit="return validatePasswordChange();"> 
    <!--
    <form class="form-vertical" role="form" method="post" onsubmit="validatePasswordChange();"  
                                            action="/user/changepwd/{{urlquery .Id}}"> 
    -->
    <fieldset>

    <div class="row">
    <div class="form-group">
        <label for="oldpassword" class="col-md-2 control-label">Old Password</label>
        <div class="col-md-6">
            <input type="password" class="form-control" id="oldpassword" name="oldpassword" value="" 
                                   placeholder="enter old password" required></input>
        </div>
    </div>
    </div>
 
    <div class="row">
    <div class="form-group">
        <label for="newpassword" class="col-md-2 control-label">New Password</label>
        <div class="col-md-6">
            <input type="password" class="form-control" id="newpassword" name="newpassword" value="" 
                                   placeholder="enter new password" required></input>
        </div>
    </div>
    </div>

    <div class="row">
    <div class="form-group">
        <label for="newpassword2" class="col-md-2 control-label">Retype New Password</label>
        <div class="col-md-6">
            <input type="password" class="form-control" id="newpassword2" name="newpassword2" value="" 
                                   placeholder="retype new password" required></input>
        </div>
    </div>
    </div>

    <div class="row">&nbsp;</div>

    <div class="row">
    <div class="form-group">
        <div class="col-md-4">
            <button class="btn btn-primary" type="button" id="pwd-submit"
                    onclick="return postForm('change-pwd-form', createURL('changepwd', 'user', {{.Id}}));">
            Change Password</button>
            <!-- <button class="btn btn-primary" type="submit" id="pwd-submit">Change Password</button> -->
            <button class="btn btn-default" type="reset">Clear</button>
        </div>
        <div class="col-md-1 col-md-offset-3">
            <a type="button" class="btn btn-primary" href="/users">
                <span class="glyphicon glyphicon-arrow-left"></span>&nbsp;&nbsp;Back
            </a>
        </div>
    </div>
    </div>

    </fieldset>
    </form>
</div>
{{end}}
