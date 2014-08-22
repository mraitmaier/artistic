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
            <h1 id="data-list-header">Styles</h1>

            {{template "style-list" .Styles}}
            <button type="button" class="btn btn-primary"
                    onclick="rerouteUsingGet('style', 'insert', '');">
            Add New Style
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

{{define "style-list"}}
    {{if .}}
    <table class="table table-striped table-hover" id="style-list-table">

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
        {{ $id := add $index 1 }}
        <tr id="style-row-{{$id}}">
            <td>{{$id}}</td>
            <td>{{printf "%s" $element.Name}}</td>
            <td>{{printf "%s" $element.Description}}</td>
            <td>
                <a href="#" data-toggle="tooltip" data-placement="left"
                            title="View style details" id="view-style"
                   onclick="rerouteUsingGet('style', 'view', {{$element.Id}});">
                    <span class="glyphicon glyphicon-eye-open"></span>
                </a>
                &nbsp;
                <a href="#" data-toggle="tooltip" data-placement="left"
                            title="Edit style" id="edit-style"
                 onclick="rerouteUsingGet('style', 'modify', {{$element.Id}});">
                    <span class="glyphicon glyphicon-cog" ></span>
                </a>
                &nbsp;
                <a href="#" data-toggle="tooltip" data-placement="left"
                            title="Delete style" id="delete-style"
                 onclick="rerouteUsingGet('style', 'delete', {{$element.Id}});">
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

{{define "style"}}
<!DOCTYPE html>
<html lang="en">

  <head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <title>Artistic - Style</title>

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
            <h1 id="data-list-header">View Style</h1>
            {{template "single-style-view" .Style}}
        {{else if eq .Cmd "modify"}}
            <h1 id="data-list-header">Modify Style</h1>
            {{template "single-style-modify" .Style}}
        {{else if eq .Cmd "insert"}}
            <h1 id="data-list-header">Create New Style</h1>
            <p>Please enter the data to create a new style.</p>
            {{template "style-create"}}
        {{else if eq .Cmd ""}} 
            <h1 id="data-list-header">View Style</h1>
            {{template "single-style-view" .Style}}
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

{{define "single-style-view"}}
<div id="view-style-table-div" class="container-fluid">
    <div class="row">
    <table id="view-style-table" class="table table-hover">
    <tbody>
        <tr> 
        <td class="col-md-2">Name</td>
        <td class="col-md-10">{{.Name}}</td> 
        </tr>
        <tr> 
        <td class="col-md-2">Description</td>
        <td class="col-md-10">{{.Description}}</td> 
        </tr>
    </tbody>
    </table>
    </div>
    <div class="row">
        <div class="col-md-1 col-md-offset-7">
        <a type="button" class="btn btn-primary" href="/styles">
        <span class="glyphicon glyphicon-arrow-left"></span>&nbsp;&nbsp;Back
        </a>
        </div>
    </div>
</div>
{{end}}

{{define "single-style-modify"}}
<div id="modify-style-table-div" class="container-fluid">
    <form class="form-vertical" role="form" method="post"
                                id="style-modify-form">
    <fieldset>

    <div class="row">
    <div class="form-group"> 
        <label for="style-name" class="col-md-2 control-label">Name</label>
        <div class="col-md-6">
        <input type="text" class="form-control" id="style-name"
               name="style-name" value="{{.Name}}"></input>
        </div>
    </div>
    </div>

    <div class="row">
    <div class="form-group"> 
        <label for="style-description" class="col-md-2 control-label">
        Description</label>
        <div class="col-md-6">
        <textarea type="text" class="form-control"  rows="5"
        id="style-description" name="style-description">
        {{.Description}}
        </textarea>
        </div>
    </div>
    </div>

    <div class="row">&nbsp;</div> <!-- empty row -->

    <div class="row">
    <div class="form-group">
        <div class="col-md-2">
            <button class="btn btn-primary" type="submit" 
                    id="style-submit">Modify Style</button>
        </div>
        <div class="col-md-1 col-md-offset-5">
            <a type="button" class="btn btn-primary" href="/styles">
            <span class="glyphicon glyphicon-arrow-left"></span>&nbsp;&nbsp;Back
            </a>
         </div>
    </div>
    </div>

    </fieldset>
    </form>
</div>
{{end}}

{{define "style-create"}}
    <div id="create-style-form-div" class="container-fluid">
    <form class="form-vertical" role="form" method="post"
                id="create-style-form" action="/style/insert/">
        <fieldset>

        <div class="row">
        <div class="form-group">
            <label for="style-name" 
                   class="col-md-2 control-label">Name</label>
            <div class="col-md-6">
            <input type="text" class="form-control" id="style-name"
                    name="style-name" value="{{.Name}}" required></input>
            </div>
        </div>
        </div>
        
        <div class="row">
        <div class="form-group">
            <label for="style-description" 
                   class="col-md-2 control-label">Description</label>
            <div class="col-md-6">
            <textarea type="text" class="form-control" rows="10"
                      name="style-description"
                      id="style-description">{{.Description}}</textarea>
                      </div>
        </div>
        </div>

        <div class="row">&nbsp;</div> <!-- empty row -->

        <div class="row">
        <div class="form-group">
            <div class="col-md-2">
                <button class="btn btn-primary" type="submit"
                        id="style-submit">Create</button>
                <button class="btn btn-default" type="reset">Clear</button>
            </div>
            <div class="col-md-1 col-md-offset-5">
                <a type="button" class="btn btn-primary" href="/styles">
                <span class="glyphicon glyphicon-arrow-left"></span>
                &nbsp;&nbsp;Back
                </a>
            </div>
        </div>
        </div>

        </fieldset>
    </form>
    </div>
{{end}}

