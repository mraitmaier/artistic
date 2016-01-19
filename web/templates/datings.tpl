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
    {{template "navbar" .User.Fullname}}

    <div class="container-fluid">
        <div class="row">
            <div class="col-md-2" id="menu">
                <h1 id="menu-header"></h1>
                {{template "accordion"}}
            </div>

            <div class="col-md-10" id="data-list">
                <h1 id="data-list-header">Datings</h1>
                <!-- datings table is created dynamically by JS -->
                <div id="datings-table-div">
                {{template "dating-list" .Datings}}
            </div>
        </div> <!-- row -->
    </div> <!-- container fluid -->

    <!-- jQuery (necessary for Bootstrap's JavaScript plugins) -->
    <!--   <script src="https://ajax.googleapis.com/ajax/libs/jquery/2.1.0/jquery.min.js"></script> -->
    <script  src="/static/js/jquery.min.js"></script>

    <!-- Include all compiled plugins (below), or include individual files as needed -->
    <script src="/static/js/bootstrap.min.js"></script>
    <script src="/static/js/artistic.js"></script>
    <script>

    // when page is ready...
    $(document).ready( function() {

        // parse the JSON data...
        {{/*//var data = JSON.parse({{.Json}}); */}}

        // create a table
        //$("#datings-table-div").append(createDatingsTable(data));
        //$("#datings-table-div").append(createDatingsTable(data));
        //createDatingsTable(data);

    });

    </script>

</body>
</html>
{{end}}

{{define "dating-list"}}
    {{if .}}
    <table class="table table-striped table-hover" id="dating-list-table">
    <thead>
        <tr> <th>#</th> <th>Dating</th> <th>Description</th> <th>Actions</th> </tr>
    </thead>

    <tbody>
        {{range $index, $element := .}}
        {{$id := add $index 1}}
        <tr id="dating-row-{{$id}}">
            <td>{{$id}}</td>
            <td>{{$element.Dating.Dating}}</td>
            <td>{{$element.Dating.Description}}</td>
            <td>
                <a href="#" data-toggle="tooltip" data-placement="left" title="View dating details" id="view-dating"
                                         onclick="rerouteUsingGet('dating', 'view', {{$element.Id}});">
                    <span class="glyphicon glyphicon-eye-open"></span>
                </a>
                &nbsp;
                <a href="#" data-toggle="tooltip" data-placement="left" title="Edit dating" id="edit-dating"
                                         onclick="rerouteUsingGet('dating', 'modify', {{$element.Id}});">
                    <span class="glyphicon glyphicon-pencil" ></span>
                </a>
            </td>
        </tr>
        {{end}}
    </tbody>

    </table>
    {{else}}
        <p>The are no datings defined yet.</p>
    {{end}}
{{end}}

{{define "dating"}}
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <title>Artistic - Dating</title>

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
    {{template "navbar" .User.Fullname}}

    <div class="container-fluid">
        <div class="row">
            <div class="col-md-2" id="menu">
                <h1 id="menu-header"></h1>
                {{template "accordion"}}
            </div>

            <div class="col-md-10" id="data-list">
        {{if eq .Cmd "view"}} 
                <h1 id="data-list-header">View Dating</h1>
                {{template "single-dating-view" .Dating}}
        {{else if eq .Cmd "modify"}}
                <h1 id="data-list-header">Modify Dating</h1>
                {{template "single-dating-modify" .Dating}}
        {{else if eq .Cmd ""}} 
                <h1 id="data-list-header">View Dating</h1>
                {{template "single-dating-view" .Dating}}
        {{end}}
            </div>
        </div> <!-- row -->
    </div> <!-- container fluid -->

    <!-- jQuery (necessary for Bootstrap's JavaScript plugins) -->
    <!--   <script src="https://ajax.googleapis.com/ajax/libs/jquery/2.1.0/jquery.min.js"></script> -->
    <script  src="/static/js/jquery.min.js"></script>

    <!-- Include all compiled plugins (below), or include individual files as needed -->
    <script src="/static/js/bootstrap.min.js"></script>
    <script src="/static/js/datings.js"></script>
    <script>

    // when page is ready...
    $(document).ready( function() {

        // parse the JSON data...
        {{/*var data = JSON.parse({{.Json}}); */}}

        // create a table
        //$("#datings-table-div").append(createDatingsTable(data));
        //$("#datings-table-div").append(createDatingsTable(data));
        //createDatingsTable(data);

    });

    </script>
</body>
</html>
{{end}}

{{define "single-dating-view"}}
<div id="view-dating-table-div" class="container-fluid">

    <div class="row">
        <table id="view-dating-table" class="table table-hover">
        <tbody>
            <tr> <td class="col-md-2">Name</td> <td class="col-md-10">{{.Dating.Dating}}</td> </tr>
            <tr> <td class="col-md-2">Description</td> <td class="col-md-10">{{.Dating.Description}}</td> </tr>
            <!--
            <tr> <td class="col-md-2">Created</td><td class="col-md-10">{printf "%q" .Dating.Created.String}</td></tr> 
            <tr> <td class="col-md-2">Last Modified</td> <td class="col-md-10">{printf "%q" .Dating.Modified.String}</td> </tr>
            -->
        </tbody>
        </table>
    </div><!-- row -->

    <div class="row">&nbsp;</row> <!-- empty row -->

    <div class="row">
        <div class="col-md-1 col-md-offset-7">
            <a type="button" class="btn btn-primary" href="/datings">
                <span class="glyphicon glyphicon-arrow-left"></span>&nbsp;&nbsp;Back
            </a>
        </div>
    </div> <!-- row -->

</div> <!-- container-fluid -->
{{end}}

{{define "single-dating-modify"}}
<div id="modify-dating-table-div" class="container-fluid">
<form class="form-vertical" role="form" method="post" id="dating-modify-form">
    <fieldset>

    <div class="row">
    <div class="form-group"> 
        <label for="dating-name" class="col-md-2 control-label">Name</label>
        <div class="col-md-6">
            <input type="text" class="form-control" id="dating-name" name="dating-name" value="{{.Dating.Dating}}" readonly></input>
        </div>
    </div>
    </div>

    <div class="row">
    <div class="form-group"> 
        <label for="dating-description" class="col-md-2 control-label"> Description</label>
        <div class="col-md-6">
            <textarea type="text" class="form-control"  rows="5" id="dating-description" name="dating-description">
            {{.Dating.Description}}
            </textarea>
        </div>
    </div>
    </div>

<!--
    <div class="row">
    <div class="form-group"> 
        <label class="col-md-2 control-label" for="created">Created</label>
        <div class="col-md-4">
            <input type="text" class="form-control" id="created" name="created" value="{.Dating.Created.String}" readonly>
            </input>
        </div>
        <label class="col-md-2 control-label" for="modified">Last Modified</label>
        <div class="col-md-4">
            <input type="text" class="form-control" id="modified" name="modified" value="{.Dating.Modified.String}" readonly>
            </input>
        </div>
    </div>
    </div>
-->

    <div class="row">&nbsp;</row> <!-- empty row -->

    <div class="row">
    <div class="form-group">
        <div class="col-md-4">
            <button class="btn btn-primary" type="submit" id="dating-submit">Modify</button>
        </div>
        <div class="col-md-1 col-md-offset-3">
            <a type="button" class="btn btn-primary" href="/datings">
                <span class="glyphicon glyphicon-arrow-left"></span>&nbsp;&nbsp;Back
            </a>
        </div>
    </div>
    </div>

    </fieldset>
</form>
</div> <!-- container-fluid -->
{{end}}
