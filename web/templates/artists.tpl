{{define "artists"}}
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <title>Artistic - View Artists</title>

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
                <h1 id="data-list-header">{{get_artist_type .Type}}s</h1>
                {{template "artist-list" .Artists}}
                <button type="button" class="btn btn-primary" onclick="rerouteUsingGet('artist', 'insert', '');">
                Add New Artist</button>
            </div>
        </div> <!-- row -->
    </div> <!-- container fluid -->

    <!-- jQuery (necessary for Bootstrap's JavaScript plugins) -->
    <!--   <script src="https://ajax.googleapis.com/ajax/libs/jquery/2.1.0/jquery.min.js"></script> -->
    <script  src="static/js/jquery.min.js"></script>
    <!-- Include all compiled plugins, or include individual files as needed -->
    <script src="static/js/bootstrap.min.js"></script>
    <!-- custom JS code -->
    <script src="static/js/artistic.js"></script>
</body>
</html>
{{end}}

{{define "artist-list"}}
{{if .}}
    <table class="table table-striped table-hover" id="artist-list-table">
    <thead>
      <tr>
        <th>#</th>
        <th>Name</th>
        <th>RealName</th>
        <th>Born</th>
        <th>Died</th>
        <th>Nationality</th>
        <th>Actions</th>
      </tr>
    </thead>

    <tbody>
      {{range $index, $element := .}}
      {{ $id := add $index 1 }}
      <tr id="artist-row-{{$id}}">
        <td>{{$id}}</td>
        <td>{{printf "%s" $element.Name}}</td>
        <td>{{printf "%s" $element.RealName}}</td>
        <td>{{printf "%s" $element.Born}}</td>
        <td>{{printf "%s" $element.Died}}</td>
        <td>{{printf "%s" $element.Nationality}}</td>
        <td>
          <a href="#" data-toggle="tooltip" data-placement="left" title="View artist details" id="view-artist-{{$id}}"
                      onclick="rerouteUsingGet('artist', 'view', {{$element.Id}});">
              <span class="glyphicon glyphicon-eye-open"></span>
          </a>
          &nbsp;
          <a href="#" data-toggle="tooltip" data-placement="left" title="Edit artist" id="edit-artist-{{$id}}"
                      onclick="rerouteUsingGet('artist', 'modify', {{$element.Id}});">
              <span class="glyphicon glyphicon-cog"></span>
          </a>
          &nbsp;
          <a href="#" data-toggle="tooltip" data-placement="left" title="Delete artist" id="delete-artist-{{$id}}"
                      onclick="rerouteUsingGet('artist', 'delete', {{$element.Id}});">
              <span class="glyphicon glyphicon-trash"></span>
          </a>
        </td>
      </tr>
      {{end}}
    </tbody>
    </table>
{{else}}
    <p>No artists were found.</p>
{{end}}
{{end}}

{{define "artist"}}
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <title>Artistic - Artist</title>

    <!-- Bootstrap -->
    <link href="static/css/bootstrap.min.css" rel="stylesheet">

    <!-- HTML5 Shim and Respond.js IE8 support of HTML5 elements and media queries -->
    <!-- WARNING: Respond.js doesn't work if you view the page via file:// -->
    <!--[if lt IE 9]>
      <script src="https://oss.maxcdn.com/libs/html5shiv/3.7.0/html5shiv.js"></script>
      <script src="https://oss.maxcdn.com/libs/respond.js/1.4.2/respond.min.js"></script>
    <![endif]-->

    <!-- custom CSS, additional to bootstrap CSS -->
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

        <div class="col-md-10" id="data-view">
    {{if eq .Cmd "view"}}
        <h1 id="data-view-header">{{.Artist.Name}}</h1>
        {{template "artist-view" .Artist}}
    {{else if eq .Cmd "modify"}}
        <h1 id="data-view-header">Modify {{.Artist.Name}}</h1>
        {{template "artist-modify" .Artist}}
    {{else if eq .Cmd "insert"}}
        <h1 id="data-view-header">Create New Artist</h1>
        <p>Please enter the data to create a new artist.</p>
        {{template "artist-create"}}
    {{else if eq .Cmd ""}}
        <h1 id="data-view-header">{{.Artist.Name}}</h1>
        {{template "artist-view" .Artist}}
    {{end}}
        </div>

      </div> <!-- row -->
    </div> <!-- container-fluid -->

    <!-- jQuery (necessary for Bootstrap's JavaScript plugins) -->
    <!--   <script src="https://ajax.googleapis.com/ajax/libs/jquery/2.1.0/jquery.min.js"> </script> -->
    <script src="static/js/jquery.min.js"></script>

    <!-- Include all compiled plugins, or include individual files as needed -->
    <script src="static/js/bootstrap.min.js"></script>
    <!-- custom JS code -->
    <script src="static/js/artistic.js"></script>

</body>
</html>
{{end}}

{{define "artist-create"}}
<div class="container-fluid" id="create-artist-form-div">
<form class="form-vertical" role="form" method="post" id="create-artist-form" action="/artist/insert/">
    <fieldset>

    <div class="row">
    <div class="form-group">
        <label for="name" class="col-md-2 control-label">Name</label>
        <div class="col-md-6">
            <input type="text" class="form-control" id="artist-name" name="artist-name" value="" placeholder="type name" required>
            </input>
        </div>
    </div> <!-- form-group -->
    </div> <!-- row -->

    <div class="row">
    <div class="form-group">
        <label for="born" class="col-md-2 control-label">Born</label>
        <div class="col-md-6">
            <input type="date" class="form-control" id="born" name="born" value="" placeholder="date of birth"></input>
        </div>
    </div> <!-- form-group -->
    </div> <!-- row -->

    <div class="row">
    <div class="form-group">
        <label for="died" class="col-md-2 control-label">Died</label>
        <div class="col-md-6">
            <input type="date" class="form-control" id="died" name="died" value="" placeholder="date of death"></input>
        </div>
    </div> <!-- form-group -->
    </div> <!-- row -->





    <div class="row">&nbsp;</div><!-- empty row -->

    <div class="row">
    <div class="form-group">
        <div class="col-md-2">
            <button class="btn btn-primary" type="submit" id="artist-submit">Create</button>
            <button class="btn btn-default" type="reset">Clear</button>
        </div>

        <div class="col-md-1 col-md-offset-5">
            <a type="button" class="btn btn-primary" href="/artists">
                <span class="glyphicon glyphicon-arrow-left"></span>&nbsp;&nbsp;Back
            </a>
        </div>
    </div> <!-- form-group -->
    </div> <!-- row -->

    </fieldset>
</form>
</div> <!-- container-fluid -->
{{end}}

{{define "artist-view"}}
<p>View artist mockup page.</p>
{{end}}

{{define "artist-modify"}}
<p>Modify artist mockup page.</p>
{{end}}
