
{{define "users"}}
<!DOCTYPE html>
<html lang="en">

  <head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <title>Artistic - Administrate users</title>

    <!-- Bootstrap -->
  <!--  <link href="css/bootstrap.min.css" rel="stylesheet"> -->
    <link href="static/css/bootstrap.min.css" rel="stylesheet">

    <!-- HTML5 Shim and Respond.js IE8 support of HTML5 elements and media queries -->
    <!-- WARNING: Respond.js doesn't work if you view the page via file:// -->
    <!--[if lt IE 9]>
      <script src="https://oss.maxcdn.com/libs/html5shiv/3.7.0/html5shiv.js"></script>
      <script src="https://oss.maxcdn.com/libs/respond.js/1.4.2/respond.min.js"></script>
    <![endif]-->

    <!-- custom CSS, additional to CSS -->
    <link href="static/css/custom.css" rel="stylesheet">
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

            <p> {{template "user-list" .Users}}</p>
        </div>

    </div> <!-- row -->

    </div> <!-- container fluid -->
<!-- jQuery (necessary for Bootstrap's JavaScript plugins) -->
 <!--   <script 
        src="https://ajax.googleapis.com/ajax/libs/jquery/2.1.0/jquery.min.js">
    </script> -->
    <script  src="static/js/jquery.min.js"></script>

<!-- Include all compiled plugins (below), or include individual files as needed -->
    <script src="static/js/bootstrap.min.js"></script>

    <script> </script>
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
            <!--
                <a href="#" data-toggle="tooltip" data-placement="left" 
                            title="View user details" id="view-user-{{$id}}">
                    <span class="glyphicon glyphicon-eye-open"></span>
            -->
                <a data-toggle="modal" data-target="#my-modal" 
                            title="View user details" id="view-user-{{$id}}">
                    <span class="glyphicon glyphicon-eye-open"></span>
                </a>
                &nbsp;
                <a href="#" data-toggle="tooltip" data-placement="left" 
                            title="Edit user" id="edir-user-{{$id}}">
                    <span class="glyphicon glyphicon-cog" ></span>
                </a>
                &nbsp;
                <a href="#" data-toggle="tooltip" data-placement="left" 
                            title="Delete user" id="delete-user-{{$id}}">
                    <span class="glyphicon glyphicon-trash"></span>
                </a>
            </td>
        </tr>
        {{end}}
    </tbody>
    </table>
{{end}}

{{define "view-user-details"}}
<div class="modal hide fade" id="view-user-modal">
  <div class="modal-dialog">
    <div class="modal-content">

      <div class="modal-header">
        <button type="button" class="close" data-dismiss="modal" 
                              aria-hidden="true">x</button>
        <h4 class="modal-title">View {{.Username}} Details</h4>
      </div>

      <div class="modal-body">
        <p>Username: {{.Username}}</p>
        <p>Full Name: {{.Name}}</p>
        <p>Role: {{.Role}}</p>
        <p>Email Address: {{.Email}}</p>
      </div>

      <div class="modal-footer">
        <button type="button" class="btn btn-default" data-dismiss="modal">
        Close
        </button>
      <!-- 
      <button type="button" class="btn btn-primary">Save changes</button>
      -->
      </div> <!-- class="modal-content" -->
    </div>
  </div>
</div>
{{end}}
