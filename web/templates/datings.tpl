
{{define "datings"}}
<!DOCTYPE html>
<html lang="en">

  <head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <title>Artistic - Datings</title>

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
            <h1 id="data-list-header">Datings</h1>

<!--            <p>Datings</p> -->
            {{template "dating-list" .Datings}}
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

{{define "dating-list"}}
    {{if .}}
    <table class="table table-striped table-hover" id="dating-list-table">

    <thead>
        <tr>
            <th>#</th>
            <th>Dating</th>
            <th>Description</th>
            <th>Action</th>
        </tr>
    </thead>

    <tbody>
        {{range $index, $element := .}}
        {{$id := add $index 1}}
        <tr id="dating-row-{{$id}}">
            <td>{{$id}}</td>
            <td>{{printf "%s" $element.Dating}}</td>
            <td>{{printf "%s" $element.Description}}</td>
            <td>
                <a href="#" data-toggle="tooltip" data-placement="left"
                            title="View dating details" id="view-dating">
                    <span class="glyphicon glyphicon-eye-open"></span>
                </a>
                &nbsp;
                <a href="#" data-toggle="tooltip" data-placement="left"
                            title="Edit dating" id="edit-dating">
                    <span class="glyphicon glyphicon-cog" ></span>
                </a>
<!--
                &nbsp;
                <a href="#" data-toggle="tooltip" data-placement="left"
                            title="Delete dating" id="delete-dating">
                    <span class="glyphicon glyphicon-trash"></span>
                </a>
-->
            </td>
        </tr>
        {{end}}
    </tbody>

    </table>
    {{else}}
        <p>The are no datings defined yet.</p>
    {{end}}
{{end}}
