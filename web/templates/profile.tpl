{{define "profile"}}
{{$role := .User.Role }}
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
    {{template "navbar" .}}

    <div class="container-fluid">
        <div class="row">
            <div class="col-md-2" id="menu">
                <h1 id="menu-header"></h1>
                {{template "accordion"}}
            </div>

            <div class="col-md-10" id="data-list">
            <h1 id="data-list-header">View My Profile</h1>

            <p>
            <div class="btn-group">
            <button type="button" class="btn btn-primary btn-sm" data-toggle="modal" data-target="#modProfileModal"
								  	data-id="{{.User.ID.Hex}}"
									data-username="{{.User.Username}}"
									data-password="{{.User.Password}}"
									data-role="{{.User.Role}}"
									data-fullname="{{.User.Fullname}}"
									data-email="{{.User.Email}}"
									data-phone="{{.User.Phone}}"
									data-disabled="{{.User.Disabled}}"
									data-mustchange="{{.User.MustChangePassword}}"
									data-created="{{.User.Created}}"
									data-modified="{{.User.Modified}}">
                <span class="glyphicon glyphicon-edit"></span> &nbsp; Modify Profile
            </button>
            <button type="button" class="btn btn-danger btn-sm" data-toggle="modal" data-target="#changePwdModal"
								  	data-id="{{.User.ID.Hex}}"
									data-username="{{.User.Username}}"
									data-created="{{.User.Created}}"
									data-modified="{{.User.Modified}}">
                <span class="glyphicon glyphicon-random"></span> &nbsp; Change Password
            </button>
            </div> <!-- btn-group -->
            </p>

            <div class="container-fluid" id="view-user-table-div">
                    <div class="row">
                        <table id="view-user-table" class="table table-hover small">
                        <tbody>
                        <tr> <td class="col-sm-2"><b>Username<b/></td><td class="col-sm-10">{{.User.Username}}</td> </tr>
                        <tr> <td class="col-sm-2"><b>Password</b></td> <td class="col-sm-10">{{.User.Password}}</td> </tr>
                        <tr> <td class="col-sm-2"><b>Role</b></td> <td class="col-sm-10">{{totitle .User.Role}}</td> </tr>
                        <tr> <td class="col-sm-2"><b>Full Name</b></td> <td class="col-sm-10">{{.User.Fullname}}</td> </tr>
                        <tr> <td class="col-sm-2"><b>Email Address</b></td> <td class="col-sm-10">{{.User.Email}}</td> </tr>
                        <tr> <td class="col-sm-2"><b>Phone #</b></td> <td class="col-sm-10">{{.User.Phone}}</td> </tr>
                        <tr> <td class="col-sm-2"><b>Disabled</b></td> 
                             <td class="col-sm-10">{{if eq .User.Disabled true}}Yes{{else}}No{{end}}</td> </tr>
                        <tr> <td class="col-sm-2"><b>Must Change Password</b></td> 
                             <td class="col-sm-10">{{if eq .User.MustChangePassword true}}Yes{{else}}No{{end}}</td> </tr>
                        <tr> <td class="col-sm-2"><b>Created</b></td> <td class="col-sm-10">{{.User.Created}}</td> </tr>
                        <tr> <td class="col-sm-2"><b>Last Modified</b></td> <td class="col-sm-10">{{.User.Modified}}</td> </tr>
                        </tbody>
                        </table>
            	</div> <!-- row -->
			</div> <!-- container-fluid -->

        </div> <!-- row -->
    </div> <!-- container fluid -->

   <!-- add modals -->
    {{template "change_pwd_modal" .User}}
    {{template "modify_profile_modal" .User}}
    <!-- end of modals definition -->

    <!-- jQuery (necessary for Bootstrap's JavaScript plugins) -->
    <!--<script src="https://ajax.googleapis.com/ajax/libs/jquery/2.1.0/jquery.min.js"></script> -->
    <script  src="/static/js/jquery.min.js"></script>

    <!-- Include all compiled plugins, or include individual files as needed -->
    <script src="/static/js/bootstrap.min.js"></script>

    <!-- custom JS code -->
    <script src="/static/js/artistic.js"></script>
    <script>

       $('#modProfileModal').on('show.bs.modal', function (event) {

            var btn = $(event.relatedTarget); // button that triggerd event
            // extract info from data-dating attribute
            var hexid = btn.data('id');
            var username = btn.data('username');
            var password = btn.data('password');
            var urole = btn.data('role');
            var fullname = btn.data('fullname');
            var email = btn.data('email');
            var phone = btn.data('phone');
            var disabled = btn.data('disabled');
            var mustchange = btn.data('mustchange');
            var created = btn.data('created');
            var modified = btn.data('modified');

            // Update the modal's content.
            var modal = $(this);
            modal.find('.modal-title').text('Modify Your Profile');
            modal.find('.modal-body #hexid').val(hexid);
            modal.find('.modal-body #username').val(username);
            modal.find('.modal-body #password').val(password);
            modal.find('.modal-body #urole').val(urole);
            modal.find('.modal-body #fullname').val(fullname);
            modal.find('.modal-body #email').val(email);
            modal.find('.modal-body #phone').val(phone);
            modal.find('.modal-body #disabled').val(disabled);
            modal.find('.modal-body #mustchange').val(mustchange);
            modal.find('.modal-body #created').val(created);   // hidden val
            modal.find('.modal-body #createdd').text(created); // only display
            modal.find('.modal-body #modifiedd').text(modified); //only display
    	});

        $('changePwdModal').on('show.bs.modal', function (event) {

            var btn = $(event.relatedTarget); // button that triggerd event
            // extract info from data-dating attribute
            var hexid = btn.data('id');
            var username = btn.data('username');
            //var fullname = btn.data('fullname');

            // Update the modal's content.
            var modal = $(this);
            modal.find('.modal-body #hexid').val(hexid);
            modal.find('.modal-body #username').val(username);
            //modal.find('.modal-body #fullname').val(fullname);
        });

		// This should post form (PUT method) to modify a profile
		var modifyProfile = function(form_id, id) {
    		var url = '/profile/' + id
            //    	alert("ID=" + id); //DEBUG
    		postForm(form_id, url);
		}

    </script>
</body>
</html>
{{end}}

{{define "modify_profile_modal"}}
<div class="modal fade" id="modProfileModal" tabindex="-1" role="dialog" aria-lebeleledby="modProdileModalLabel">
<div class="modal-dialog">
<div class="modal-content">

    <div class="modal-header">
    <div class="container-fluid">
        <div class="row">
            <h4 class="modal-title col-md-8" id="modProfileModalLabel">Empty Profile Details</h4>
             <button type="button" class="btn btn-primary btn-sm col-sm-2" 
                     onclick="modifyProfile('modify_profile_form', $('#hexid').val()); $('#modProfileModal').modal('hide');">
					Modify
             </button>
            <button type="button" class="btn btn-default btn-sm col-md-2" data-dismiss="modal">Cancel</button>
        </div> <!-- row -->
    </div> <!-- container-fluid -->
    </div> <!-- modal-header -->

    <div class="modal-body">
    <div class="container-fluid">
        <div class="row">
        <form id="modify_profile_form" class="form-horizontal">
            <input type="hidden" id="hexid" name="hexid" />
        	<input type="hidden" id="password" name="password" />
        	<input type="hidden" id="urole" name="urole" />
        	<input type="hidden" id="disabled" name="disabled" />
        	<input type="hidden" id="mustchange" name="mustchange" />

        <div class="form-group form-group-sm">
            <label for="username" class="col-sm-3 control-label">Username</label>
            <div class="col-sm-9">
                <input type="text" class="form-control" id="username" name="username" readonly></input>
            </div>
        </div>
        <div class="form-group form-group-sm">
            <label for="fullname" class="col-sm-3 control-label">Full Name</label>
            <div class="col-sm-9">
                <input type="text" class="form-control" id="fullname" name="fullname"></input>
            </div>
        </div>
        <div class="form-group form-group-sm">
            <label for="email" class="col-sm-3 control-label">Email Address</label>
            <div class="col-sm-9">
                <input type="email" class="form-control" id="email" name="email"></input>
            </div>
        </div>
        <div class="form-group form-group-sm">
            <label for="phone" class="col-sm-3 control-label">Phone Number</label>
            <div class="col-sm-9">
                <input type="phone" class="form-control" id="phone" name="phone"></input>
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
Your 
{{define "change_pwd_modal"}}
<div class="modal fade" id="changePwdModal" tabindex="-1" role="dialog" aria-lebeleledby="changePwdModalLabel">
<div class="modal-dialog">
<div class="modal-content">

    <div class="modal-header">
    <div class="container-fluid">
        <div class="row">
            <h4 class="modal-title col-md-8" id="changePwdModalLabel">Change Your Password</h4>
             <button type="button" class="btn btn-danger btn-sm col-sm-2" 
                     onclick="changePwd('change_pwd_form', $('#hexid').val()); $('#changePwdModal').modal('hide');"> Change 
             </button>
            <button type="button" class="btn btn-default btn-sm col-md-2" data-dismiss="modal">Cancel</button>
        </div> <!-- row -->
    </div> <!-- container-fluid -->
    </div> <!-- modal-header -->

    <div class="modal-body">
    <div class="container-fluid">
    <form class="form-horizontal" method="post" onsubmit="validatePasswordChange(this);" id="change_pwd_form"> 
            <input type="hidden" id="hexid" name="hexid" value="{{.ID.Hex}}" />
            <input type="hidden" id="prev" name="prev" value="profile" />

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
