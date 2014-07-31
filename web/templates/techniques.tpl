{{define "techniques"}}
<!DOCTYPE html>
<html lang="en">

  <head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <title>Artistic - Techniques</title>

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
            <h1 id="data-list-header">Techniques</h1>

            {{template "tech-list" .Techniques}}

            <button type="button" class="btn btn-primary"
                    onclick="rerouteUsingGet('technique', 'create', '');">
            Add New Technique
            </button>
        </div>

     </div> <!-- row -->

    </div> <!-- container fluid -->
<!-- jQuery (necessary for Bootstrap's JavaScript plugins) -->
 <!--   <script 
        src="https://ajax.googleapis.com/ajax/libs/jquery/2.1.0/jquery.min.js">
    </script> -->
    <script  src="/static/js/jquery.min.js"></script>

<!-- Include all compiled plugins (below), or include individual files as needed -->
    <script src="/static/js/bootstrap.min.js"></script>
    <script src="/static/js/artistic.js"></script>
    <script> </script>
  </body>
</html>
{{end}}

{{define "tech-list"}}
    {{if .}}
    <table class="table table-striped table-hover" id="techniques-list-table">

    <thead>
        <tr>
            <th>#</th>
            <th>Technique</th>
            <th>Description</th>
            <th>Actions</th>
        </tr>
    </thead>

    <tbody>
        {{range $index, $element := .}}
        {{$id := add $index 1}}
        <tr id="technique-row-{{$id}}">
            <td>{{$id}}</td>
            <td>{{$element.Name}}</td>
            <td>{{$element.Description}}</td>
            <td>
                <a data-toggle="tooltip" data-placement="left"
                   title="View technique details" id="view-technique"
          onclick="rerouteUsingGet('technique', 'view', {{$element.Id}});">
                    <span class="glyphicon glyphicon-eye-open"></span>
                </a>
                &nbsp;
                <a data-toggle="tooltip" data-placement="left"
                   title="Modify technique" id="modify-technique"
          onclick="rerouteUsingGet('technique', 'modify', {{$element.Id}});">
<!--         onclick="rerouteTechnique('PUT', 'modify', {{$element.Id}});"-->
                    <span class="glyphicon glyphicon-cog" ></span>
                </a>
                &nbsp;
                <a data-toggle="tooltip" data-placement="left"
                   title="Delete technique" id="delete-technique"
          onclick="rerouteUsingGet('technique', 'delete', {{$element.Id}});">
       <!--  onclick="rerouteTechnique('DELETE', 'delete', {{$element.Id}});"-->
                    <span class="glyphicon glyphicon-trash"></span>
                </a>
            </td>
        </tr>
        {{end}}
    </tbody>

    </table>
    {{else}}
    <p>There are no techniques defined yet.</p>
    {{end}}

{{end}}

{{define "technique"}}
<!DOCTYPE html>
<html lang="en">

  <head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <title>Artistic - Technique</title>

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
            <h1 id="data-list-header">View Technique</h1>
            {{template "single-technique-view" .Technique}}
        {{else if eq .Cmd "modify"}}
            <h1 id="data-list-header">Modify Technique</h1>
            {{template "single-technique-modify" .Technique}}
        {{else if eq .Cmd "create"}}
            <h1 id="data-list-header">Create New Technique</h1>
            <p>Please enter the data to create a new technique.</p>
            {{template "technique-create"}}
        {{else if eq .Cmd ""}} 
            <h1 id="data-list-header">View Technique</h1>
            {{template "single-technique-view" .Technique}}
        {{end}}
        </div>

     </div> <!-- row -->

    </div> <!-- container fluid -->

<!-- jQuery (necessary for Bootstrap's JavaScript plugins) -->
 <!--   <script 
        src="https://ajax.googleapis.com/ajax/libs/jquery/2.1.0/jquery.min.js">
    </script> -->
    <script  src="/static/js/jquery.min.js"></script>

<!-- Include all compiled plugins (below), or include individual files as needed -->
    <script src="/static/js/bootstrap.min.js"></script>
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

{{define "single-technique-view"}}
<div id="view-technique-table-div">
    <table id="view-technique-table" class="table table-hover">
    <tbody>
        <tr> <td>Name</td><td>{{.Name}}</td> </tr>
        <tr> <td>Description</td><td>{{.Description}}</td> </tr>
    </tbody>
    </table>
</div>
{{end}}

{{define "single-technique-modify"}}
<div id="modify-technique-table-div">

    <form class="form-vertical" role="form" method="post"
                                id="technique-modify-form">
    <fieldset>

    <div class="form-group"> 
        <label for="technique-name" class="col-md-2 control-label">Name</label>
        <div class="col-md-10">
        <input type="text" class="form-control" id="technique-name"
               name="technique-name" value="{{.Name}}"></input>
        </div>
    </div>
    <div class="form-group"> 
        <label for="technique-description" class="col-md-2 control-label">
        Description</label>
        <div class="col-md-10">
        <textarea type="text" class="form-control"  rows="5"
        id="technique-description" name="technique-description">
        {{.Description}}
        </textarea>
        </div>
    </div>
    <div class="form-group">
        <button class="btn btn-primary" type="submit" 
                id="technique-submit">Modify</button>
    </div>

    </fieldset>
    </form>

</div>
{{end}}

{{define "technique-create"}}
    <div id="create-technique-form-div">
    <form class="form-vertical" role="form" method="post"
                id="create-technique-form" action="/technique/create/">
        <fieldset>
        <div class="form-group">
            <label for="technique-name" 
                   class="col-md-2 control-label">Name</label>
            <div class="col-md-10">
            <input type="text" class="form-control" id="technique-name"
                    name="technique-name" value="{{.Name}}" required></input>
            </div>
        </div>
        
        <div class="form-group">
            <label for="technique-description" 
                   class="col-md-2 control-label">Description</label>
            <div class="col-md-10">
            <textarea type="text" class="form-control" rows="10"
                      name="technique-description"
                      id="technique-description">{{.Description}}</textarea>
                      </div>
        </div>

        <div class="form-group">
            <button class="btn btn-primary" type="submit"
                    id="technique-submit">Create</button>
            <button class="btn btn-default" type="reset">Clear</button>
        </div>
        </fieldset>
    </form>
    </div>
{{end}}


