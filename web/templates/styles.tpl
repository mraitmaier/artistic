
{{define "styles"}}
<!DOCTYPE html>
<html lang="en">

  <head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <title>Artistic - Styles</title>

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
            <h1 id="data-list-header">Styles</h1>

<!--            <p>Styles</p> -->
            {{template "style-list" .Styles}}
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

{{define "style-list"}}
    {{if .}}
    <table class="table table-striped table-hover">

    <thead>
        <tr>
            <th>#</th>
            <th>Style</th>
            <th>Description</th>
            <th>Actions</th>
        </tr>
    </thead>

    <tbody>
        {{range $index, $element := .}}
        <tr>
            <td>{{add $index 1}}</td>
            <td>{{printf "%s" $element.Name}}</td>
            <td>{{printf "%s" $element.Description}}</td>
            <td>
                <a href="#" data-toggle="tooltip" data-placement="left"
                            title="View style details" id="view-style">
                    <span class="glyphicon glyphicon-eye-open"></span>
                </a>
                &nbsp;
                <a href="#" data-toggle="tooltip" data-placement="left"
                            title="Edit style" id="edit-style">
                    <span class="glyphicon glyphicon-cog" ></span>
                </a>
                &nbsp;
                <a href="#" data-toggle="tooltip" data-placement="left"
                            title="Delete style" id="delete-style">
                    <span class="glyphicon glyphicon-trash"></span>
                </a>
            </td>
        </tr>
        {{end}}
    </tbody>

    </table>
    {{else}}
    <p>There are no styles defined yet.</p>
    {{end}}
{{end}}
