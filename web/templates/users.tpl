{{define "users"}}
{{$role := .User.Role}}
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <title>Artistic - Administer Users</title>

    <!-- Bootstrap -->
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
    {{template "navbar" .}}
    <div class="container-fluid">
        <div class="row">
            <div class="col-md-2" id="menu">
                <h1 id="menu-header"></h1>
                {{template "accordion"}}
            </div>

            <div class="col-md-10" id="data-list">
                <h1 id="data-list-header">Users</h1>

        {{if eq $role "admin"}}
                <div id="new-user-btn">
                    <button type="button" class="btn btn-primary btn-sm" data-toggle="modal" data-target="#addUserModal">
                        <span class="glyphicon glyphicon-plus"></span> &nbsp; Add a New User
                    </button>
                </div>
        {{end}}
                <br />

        {{if .Users}}
                <table class="table table-striped table-hover small" id="user-list-table">
                    <thead>
                        <tr>
                            <th class="col-sm-1">#</th>
                            <th class="col-sm-1">Username</th>
                            <th class="col-sm-2">Name</th>
                            <th class="col-sm-1">Role</th>
                            <th class="col-sm-2">Email</th>
                            <th class="col-sm-1">Phone #</th>
                            <th class="col-sm-1">Disabled</th>
                            <th class="col-sm-1">MCP</th>
                            <th class="col-sm-2">Actions</th>
                        </tr>
                    </thead>

            		<tfoot>
                    	<tr class="bg-primary">
                    		<td colspan="9"><strong>{{.Num}} {{if eq .Num 1}} user {{else}} users {{end}} found.</strong></td>
                     	</tr>
                	</tfoot>

                    <tbody>
                        {{ $uname := .User.Username }}
                        {{range $index, $element := .Users}}
                            {{ $cnt := add $index 1 }}
                            {{if eq $element.Visible true}}
                        <tr id="user-row-{{$cnt}}">
                            <td>{{$cnt}}</td>
                            <td>{{$element.Username}}</td>
                            <td>{{$element.Fullname}}</td>
                            <td>{{totitle $element.Role}}</td>
                            <td>{{$element.Email}}</td>
                            <td>{{$element.Phone}}</td>
                            <td>{{if eq $element.Disabled true}}Yes{{else}}No{{end}}</td>
                            <td>{{if eq $element.MustChangePassword true}}Yes{{else}}No{{end}}</td>
                            <td>
								<span data-toggle="tooltip" data-placement="up" title="View details">
                            	<a data-toggle="modal" data-target="#viewUserModal"
                                       data-id="{{$element.ID.Hex}}"
                                       data-created="{{$element.Created}}"
                                       data-modified="{{$element.Modified}}"
                                       data-username="{{$element.User.Username}}"
                                       data-password="{{$element.User.Password}}"
                                       data-role="{{$element.User.Role}}"
                                       data-fullname="{{$element.User.Fullname}}"
                                       data-email="{{$element.User.Email}}"
                                       data-phone="{{$element.User.Phone}}"
                                       data-disabled="{{$element.User.Disabled}}"
                                       data-mcp="{{$element.User.MustChangePassword}}">
                                 <span class="glyphicon glyphicon-eye-open"></span>
                            	</a>
                            	</span>            
                            	&nbsp;&nbsp;
                                {{if and (ne $element.User.Username $uname) (eq $role "admin") }}
                            	<span data-toggle="tooltip" data-placement="up" title="Modify Details"> 
                            	<a data-toggle="modal" data-target="#modifyUserModal"
                                       data-id="{{$element.ID.Hex}}"  
                                       data-created="{{$element.Created}}" 
                                       data-modified="{{$element.Modified}}" 
                                       data-username="{{$element.User.Username}}"
                                       data-password="{{$element.User.Password}}"
                                       data-role="{{$element.User.Role}}"
                                       data-fullname="{{$element.User.Fullname}}"
                                       data-email="{{$element.User.Email}}"
                                       data-phone="{{$element.User.Phone}}"
                                       data-disabled="{{$element.User.Disabled}}"
                                       data-mcp="{{$element.User.MustChangePassword}}">
                                <span class="glyphicon glyphicon-edit"></span>
                            	</a>
                            	</span>       
                            	&nbsp;&nbsp;
                            	<span data-toggle="tooltip" data-placement="up" title="Remove"> 
                            	<a data-toggle="modal" data-target="#removeUserModal"
                                       data-id="{{$element.ID.Hex}}"  
                                       data-username="{{$element.User.Username}}"  
                                       data-fullname="{{$element.User.Fullname}}">
                                <span class="glyphicon glyphicon-remove"></span>
                            	</a>
                            	</span>       
                            	&nbsp;&nbsp;
                                {{end}}
                                {{if eq $role "admin"}}
                            	<span data-toggle="tooltip" data-placement="up" title="Change Password"> 
                            	<a data-toggle="modal" data-target="#changePasswdModal"
                                       data-id="{{$element.ID.Hex}}"  
                                       data-username="{{$element.User.Username}}"  
                                       data-fullname="{{$element.User.Fullname}}">
                                <span class="glyphicon glyphicon-random"></span>
                            	</a>
                            	</span>       
                                {{end}}
                            </td>
                        </tr>
                            {{end}}
                        {{end}}
                    </tbody>
                </table>

		<!-- Add modals -->
    {{template "view_user_modal"}}
{{if eq $role "admin"}}
    {{template "modify_user_modal"}}
    {{template "remove_user_modal"}}
    {{template "change_passwd_modal"}}
{{end}}
    <!-- end of modals definition -->        
        {{else}}
            <p>No users found.</p>
        {{end}}

            </div>
        </div> <!-- row -->
    </div> <!-- container fluid -->

{{if eq $role "admin"}}
   {{template "add_user_modal"}}
{{end}}

    <script  src="/static/js/jquery.min.js"></script>
    <!-- Include all compiled plugins, or include individual files as needed -->
    <script src="/static/js/bootstrap.min.js"></script>
    <!-- custom JS code -->
    <script src="/static/js/artistic.js"></script>
	<script>

        $('#viewUserModal').on('show.bs.modal', function (event) {

            var btn = $(event.relatedTarget); // button that tr;iggerd event
            // extract info from data-dating attribute
            var hexid = btn.data('id');
            var username = btn.data('username');
            var password = btn.data('password');
            var role = btn.data('role');
            var fullname = btn.data('fullname');
            var email = btn.data('email');
            var phone = btn.data('phone');
            var disabled = btn.data('disabled');
            var mustchange = btn.data('mcp');
            var created = btn.data('created');
            var modified = btn.data('modified');

            // Update the modal's content.
            var modal = $(this);
            modal.find('.modal-title').text('View User Details');
            modal.find('.modal-body #hexid').val(hexid);
            modal.find('.modal-body #username').text(username);
            modal.find('.modal-body #password').text(password);
            modal.find('.modal-body #urole').text(toTitleCase(role));
            modal.find('.modal-body #fullname').text(fullname);
            modal.find('.modal-body #email').text(email);
            modal.find('.modal-body #phone').text(phone);
            if (disabled === "true") {
                modal.find('.modal-body #disabled').text('Yes');
            } else {
                modal.find('.modal-body #disabled').text('No');
            }
            if (mustchange === "true") {
                modal.find('.modal-body #mustchange').text('Yes');
            } else {
                modal.find('.modal-body #mustchange').text('No');
            }
            modal.find('.modal-body #created').text(created);  
            modal.find('.modal-body #modified').text(modified);
    	});

{{if eq $role "admin"}}
       $('#modifyUserModal').on('show.bs.modal', function (event) {

            var btn = $(event.relatedTarget); // button that triggerd event
            // extract info from data-dating attribute
            var hexid = btn.data('id');
            var username = btn.data('username');
            var password = btn.data('password');
            var role = btn.data('role');
            var fullname = btn.data('fullname');
            var email = btn.data('email');
            var phone = btn.data('phone');
            var disabled = btn.data('disabled');
            var mustchange = btn.data('mcp');
            var created = btn.data('created');
            var modified = btn.data('modified');

            // Update the modal's content.
            var modal = $(this);
            modal.find('.modal-title').text('Modify User Details');
            modal.find('.modal-body #hexid').val(hexid);
            modal.find('.modal-body #username').val(username);
            modal.find('.modal-body #password').val(password);
            modal.find('.modal-body #urole').val(role);
            modal.find('.modal-body #fullname').val(fullname);
            modal.find('.modal-body #email').val(email);
            modal.find('.modal-body #phone').val(phone);
            if (disabled === "true") {
                modal.find('.modal-body #disabled').val('yes');
            } else {
                modal.find('.modal-body #disabled').val('no');
            }
            if (mustchange === "true") {
                modal.find('.modal-body #mustchange').val('yes');
            } else {
                modal.find('.modal-body #mustchange').val('no');
            }
            modal.find('.modal-body #created').val(created);   // hidden val
            modal.find('.modal-body #createdd').text(created); // only display
            modal.find('.modal-body #modifiedd').text(modified); //only display

            modal.find('.modal-body #ch_pwd_btn').attr('data-id', hexid);
            modal.find('.modal-body #ch_pwd_btn').attr('data-username', username);
            modal.find('.modal-body #ch_pwd_btn').attr('data-fullname', fullname);
    	});

        //
        $('#removeUserModal').on('show.bs.modal', function (event) {

            var btn = $(event.relatedTarget); // button that triggerd event
            // extract info from data-dating attribute
            var hexid = btn.data('id');
            var username = btn.data('username');
            var fullname = btn.data('fullname');
            // Update the modal's content.
            var modal = $(this);
            modal.find('.modal-body #hexid').val(hexid);
            modal.find('.modal-body #removename').text(fullname +' (' + username + ')');

        	// Let's define the 'remove' button onclick() callback... 
        	var url = '/user/' + hexid + '/delete';
        	$('#removebtn').on('click', function(e) { 
            	postForm('remove_user_form', url); 
            	$('#removeUserModal').modal('hide');
        	});
    	});

        //
        $('#changePasswdModal').on('show.bs.modal', function (event) {

            var btn = $(event.relatedTarget); // button that triggerd event
            // extract info from data-dating attribute
            var hexid = btn.data('id');
            var username = btn.data('username');
            var fullname = btn.data('fullname');
            // Update the modal's content.
            var modal = $(this);
            modal.find('.modal-body #pwdid').val(hexid);
            modal.find('.modal-body #name').text(fullname +' (' + username + ')');
    	});

		// This should post form to modify a profile
		var modifyUser = function(form_id, id) {
    		var url = '/user/' + id + '/put';
            //   	alert("ID=" + id); //DEBUG
    		postForm(form_id, url);
		}
{{end}}
    	</script>
</body>
</html>
{{end}}

{{define "add_user_modal"}}
<div class="modal fade" id="addUserModal" tabindex="-1" role="dialog" aria-labelledby="addUserModalLabel">
<div class="modal-dialog">
<div class="modal-content">

    <div class="modal-header">
    <div class="container-fluid">
        <div class="row">
            <h3 class="modal-title col-sm-8" id="addUserModalLabel">Add a New User</h3>
            <button type="button" class="btn btn-primary btn-sm col-sm-2" 
                    onclick="postForm('add_user_form', '/user'); $('#addUserModal').modal('hide');"> Add </button>
            <button type="button" class="btn btn-default btn-sm col-sm-2" data-dismiss="modal"> Cancel </button>
        </div> <!-- row -->
    </div> <!-- container-fluid -->
    </div> <!-- modal-header -->

    <div class="modal-body">
      <form id="add_user_form" class="form-horizontal" method="post">
        <div class="form-group form-group-sm">
            <label for="username" class="col-sm-4 control-label">Username</label>
            <div class="col-sm-8">
                <input type="text" class="form-control" id="username" name="username" placeholder="Username" autofocus required>
                </input>
            </div>
        </div>
        <div class="form-group form-group-sm">
            <label for="password" class="col-sm-4 control-label">Password</label>
            <div class="col-sm-8">
                <input type="password" class="form-control" id="password" name="password" placeholder="Password" required>
                </input>
            </div>
        </div>
        <div class="form-group form-group-sm">
            <label for="password2" class="col-sm-4 control-label">Retype Password</label>
            <div class="col-sm-8">
                <input type="password" class="form-control" id="password2" name="password2" 
                                       placeholder="Retype password" required> </input>
            </div>
        </div>
        <div class="form-group form-group-sm">
            <label for="role" class="col-sm-4 control-label">User Role</label>
            <div class="col-sm-8">
                <select class="form-control" name="urole" id="urole" required>
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
        <div class="form-group form-group-sm">
            <label for="fullname" class="col-sm-4 control-label">Full Name</label>
            <div class="col-sm-8">
                <input type="text" class="form-control" name="fullname" id="fullname" placeholder="Enter full name"></input>
            </div>
        </div>
        <div class="form-group form-group-sm">
            <label for="email" class="col-sm-4 control-label">Email Address</label>
            <div class="col-sm-8">
                <input type="email" class="form-control" id="email" name="email" placeholder="Email address"></input>
            </div>
        </div>
        <div class="form-group form-group-sm">
            <label for="phone" class="col-sm-4 control-label">Phone Number</label>
            <div class="col-sm-8">
                <input type="tel" class="form-control" id="phone" name="phone" placeholder="phone number"></input>
            </div>
        </div>
        <div class="form-group form-group-sm">
            <label for="disabled" class="col-sm-4 control-label">Disabled</label>
            <div class="col-sm-8">
                <select class="form-control" id="disabled" name="disabled" required>
                    <option value="yes"> Yes </option>
                    <option value="no"> No </option>
                </select>
            </div>
        </div>
        <div class="form-group form-group-sm">
            <label for="change" class="col-sm-4 control-label">Must Change Password</label>
            <div class="col-sm-8">
                <select class="form-control" id="change" name="change" required>
                    <option value="yes"> Yes </option>
                    <option value="no"> No </option>
                </select>
            </div>
        </div>

      </form>
    </div> <!-- modal-body -->
</div>
</div>
</div>
{{end}}

{{define "view_user_modal"}}
<div class="modal fade" id="viewUserModal" tabindex="-1" role="dialog" aria-lebeleledby="ViewUserModalLabel">
<div class="modal-dialog modal-lg">
<div class="modal-content">

    <div class="modal-header">
    <div class="container-fluid">
        <div class="row">
            <h3 class="modal-title col-md-10" id="viewUserModalLabel">Empty User Details</h3>
            <button type="button" class="btn btn-default btn-sm col-md-2" data-dismiss="modal">Close</button>
        </div> <!-- row -->
    </div> <!-- container-fluid -->
    </div> <!-- modal-header -->

    <div class="modal-body">
	<div class="container-fluid" id="view-user-table-div">
        <div class="row">
             <table id="view-user-table" class="table table-hover small">
             <tbody>
                 <tr> <td class="col-sm-3"><b>Username<b/></td>
                      <td class="col-sm-9"><span id="username"></span></td> </tr>
                 <tr> <td class="col-sm-3"><b>Password</b></td> 
                      <td id="password" class="col-sm-9"></td> </tr>
                 <tr> <td class="col-sm-3"><b>Role</b></td> 
                      <td id="urole" class="col-sm-9"></td> </tr>
                 <tr> <td class="col-sm-3"><b>Full Name</b></td> 
                      <td id="fullname" class="col-sm-9"></td> </tr>
                 <tr> <td class="col-sm-3"><b>Email Address</b></td> 
                      <td id="email" class="col-sm-9"></td> </tr>
                 <tr> <td class="col-sm-3"><b>Phone Number</b></td> 
                      <td id="phone" class="col-sm-9"></td> </tr>
                 <tr> <td class="col-sm-3"><b>Disabled</b></td> 
                      <td id="disabled" class="col-sm-9"></td> </tr>
                 <tr> <td class="col-sm-3"><b>Must Change Password</b></td> 
                      <td id="mustchange" class="col-sm-9"></td> </tr>
                 <tr> <td class="col-sm-3"><b>Created</b></td> 
                      <td id="created" class="col-sm-9"></td> </tr>
                 <tr> <td class="col-sm-3"><b>Last Modified</b></td> 
                      <td id="modified" class="col-sm-9"></td> </tr>
             </tbody>
             </table>
      	</div> <!-- row -->
	</div> <!-- container-fluid -->
   	</div> <!-- modal-body -->

</div>
</div>
</div>
{{end}}

{{define "modify_user_modal"}}
<div class="modal fade" id="modifyUserModal" tabindex="-1" role="dialog" aria-lebeleledby="modifyUserModalLabel">
<div class="modal-dialog modal-lg">
<div class="modal-content">

    <div class="modal-header">
    <div class="container-fluid">
        <div class="row">
            <h3 class="modal-title col-md-8" id="modifyUserModalLabel">Empty User Details</h3>
             <button type="button" class="btn btn-primary btn-sm col-sm-2" 
                     onclick="modifyUser('modify_user_form',$('#hexid').val());$('#modifyUSerModal').modal('hide');">
                     Modify
             </button>
            <button type="button" class="btn btn-default btn-sm col-md-2" data-dismiss="modal">Cancel</button>
        </div> <!-- row -->
    </div> <!-- container-fluid -->
    </div> <!-- modal-header -->

    <div class="modal-body">
    <div class="container-fluid">
        <div class="row">

        <form id="modify_user_form" class="form-horizontal">
            <input type="hidden" id="hexid" name="hexid" />

      	<div class="form-group form-group-sm">
            <label for="username" class="col-md-2 control-label">Username</label>
        	<div class="col-md-4">
            	<input type="text" class="form-control" id="username" name="username" readonly></input>
        	</div>
        	<label for="fullname" class="col-md-2 control-label">Full Name</label>
        	<div class="col-md-4">
            <input type="text" class="form-control" id="fullname" name="fullname"></input>
        	</div>
    	</div>

      	<div class="form-group form-group-sm">
            <label for="password" class="col-md-2 control-label">Password</label> 
        	<div class="col-md-10">
            <input type="text" class="form-control" id="password" name="password" readonly></input>
            </div>
    	</div>

      	<div class="form-group form-group-sm">
        	<label for="urole" class="col-md-2 control-label">User Role</label>
        	<div class="col-md-10">
            <select class="form-control" name="urole" id="urole" required>
        {{ $roles := allowedroles }}
        {{range $role := $roles}}
                <option value="{{$role}}">{{totitle $role}}</option>
        {{end}}
            </select>
        	</div>
    	</div>

      	<div class="form-group form-group-sm">
        	<label for="email" class="col-md-2 control-label">E-mail Address</label>
        	<div class="col-md-4">
            <input type="email" class="form-control" id="email" name="email"></input>
        	</div>
        	<label for="phone" class="col-md-2 control-label">Phone Number</label>
        	<div class="col-md-4">
            <input type="tel" class="form-control" id="phone" name="phone"></input>
        	</div>
    	</div>

      	<div class="form-group form-group-sm">
        	<label for="disabled" class="col-md-2 control-label">Disabled</label>
        	<div class="col-md-4">
             <select class="form-control" id="disabled" name="disabled" required>
                 <option value="yes">Yes</option>
                 <option value="no">No</option>
             </select>
        	</div>
        	<label for="change" class="col-md-2 control-label">Change Password</label>
        	<div class="col-md-4">
            <select class="form-control" id="mustchange" name="mustchange" required>
                <option value="yes">Yes</option>
                <option value="no">No</option>
            </select>
        	</div>
    	</div>
   		<div class="form-group form-group-sm small">
            <input type="hidden" id="created" name="created" />
            <div class="col-sm-2 text-right"><strong>Created:</strong></div>
            <div id="createdd" name="createdd" class="col-sm-4 text-left">Error</div>
            <div class="col-sm-2 text-right"><strong>Modified:</strong></div>
            <div id="modifiedd" name="modifiedd" class="col-sm-4 text-left">Error</div>
        </div>

        </div> <!-- row -->
    </div> <!-- container-fluid -->
    </form>
    </div> <!-- modal-body -->

</div>
</div>
</div>
{{end}}

{{define "remove_user_modal"}}
<div class="modal fade" id="removeUserModal" tabindex="-1" role="dialog" aria-labelledby="removeUserModalLabel">
<div class="modal-dialog">
<div class="modal-content">
    <div class="modal-header">
    <div class="container-fluid">
        <div class="row">
            <h3 class="modal-title col-sm-8" id="removeUserModalLabel">Remove User</h3>
            <button type="button" class="btn btn-primary btn-sm col-sm-2" id="removebtn"> Remove </button>
            <button type="button" class="btn btn-default btn-sm col-sm-2" data-dismiss="modal"> Cancel </button>
        </div> <!-- row -->
    </div> <!-- container-fluid -->
    </div> <!-- modal-header -->

    <div class="modal-body">
    <p> Would you really like to remove the user <strong><span id="removename"></span></strong>?</p>
    <form method="post" id="remove_user_form">
        <input type="hidden" name="id" id="id" />
        <input type="hidden" name="fullname" id="fullname" />
        <input type="hidden" name="username" id="username" />
    </form>
    </div>
</div>
</div>
</div>
{{end}}

{{define "change_passwd_modal"}}
<div class="modal fade" id="changePasswdModal" tabindex="-1" role="dialog" aria-lebeleledby="changePasswdModalLabel">
<div class="modal-dialog">
<div class="modal-content">

    <div class="modal-header">
    <div class="container-fluid">
        <div class="row">
            <h3 class="modal-title col-md-8" id="changePwdModalLabel">Change Password</h3>
             <button type="button" class="btn btn-danger btn-sm col-sm-2" 
                     onclick="changePwd('change_pwd_form', $('#pwdid').val()); $('#changePasswdModal').modal('hide');"> Change 
             </button>
            <button type="button" class="btn btn-default btn-sm col-md-2" data-dismiss="modal">Cancel</button>
        </div> <!-- row -->
    </div> <!-- container-fluid -->
    </div> <!-- modal-header -->

    <div class="modal-body">
    <div class="container-fluid">

    <p>Changing password for <strong><span id="name"></span></strong></p>
    <br />

    <form class="form-horizontal" method="post" id="change_pwd_form" onsubmit="return validatrPasswordChange(this);"> 
            <input type="hidden" id="pwdid" name="pwdid"></input>
            <input type="hidden" name="prev" value="user"></input>

        <div class="row">
        <div class="form-group form-group-sm">
            <label for="oldpassword" class="col-sm-4 control-label">Old Password</label>
            <div class="col-sm-8">
                <input type="password" class="form-control" id="oldpassword" name="oldpassword"
                                       placeholder="enter old password" required autofocus></input>
            </div>
        </div>
 
        <div class="form-group form-group-sm">
            <label for="newpassword" class="col-sm-4 control-label">New Password</label>
            <div class="col-sm-8">
                <input type="password" class="form-control" id="newpassword" name="newpassword"
                                       placeholder="enter new password" required></input>
            </div>
        </div>

        <div class="form-group form-group-sm">
            <label for="newpassword2" class="col-sm-4 control-label">Retype New Password</label>
            <div class="col-sm-8">
                <input type="password" class="form-control" id="newpassword2" name="newpassword2"
                                       placeholder="retype new password" required></input>
            </div>
        </div>
        </div> <!-- row -->

    </form>

    </div>
    </div> <!-- modal-body -->

</div>
</div>
</div>
{{end}}

