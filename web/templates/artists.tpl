{{define "artists"}}
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <title>Artistic - View Artists</title>

    <!-- Bootstrap -->
    <link href="/static/css/bootstrap.min.css" rel="stylesheet">
    <!-- HTML5 Shim and Respond.js IE8 support of HTML5 elements and media queries -->
    <!-- WARNING: Respond.js doesn't work if you view the page via file:// -->
    <!--[if lt IE 9]>
        <script src="https://oss.maxcdn.com/libs/html5shiv/3.7.0/html5shiv.js"></script>
        <script src="https://oss.maxcdn.com/libs/respond.js/1.4.2/respond.min.js"></script>
    <![endif]-->
    <!-- custom CSS, additional to bootstrap -->
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
                <h1 id="data-list-header">{{get_artist_type .Type}}s</h1>
                {{template "artist-list" .Artists}}
                <button type="button" class="btn btn-primary" onclick="rerouteUsingGet('artist', 'insert', '');">
                Add New Artist</button>
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
          <a href="#" data-toggle="tooltip" data-placement="left" title="Modify artist data" id="edit-artist-{{$id}}"
                      onclick="rerouteUsingGet('artist', 'modify', {{$element.Id}});">
              <span class="glyphicon glyphicon-pencil"></span>
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
    <title>Artistic - Artist Administration</title>

    <!-- Bootstrap -->
    <link href="/static/css/bootstrap.min.css" rel="stylesheet">
    <!-- HTML5 Shim and Respond.js IE8 support of HTML5 elements and media queries -->
    <!-- WARNING: Respond.js doesn't work if you view the page via file:// -->
    <!--[if lt IE 9]>
      <script src="https://oss.maxcdn.com/libs/html5shiv/3.7.0/html5shiv.js"></script>
      <script src="https://oss.maxcdn.com/libs/respond.js/1.4.2/respond.min.js"></script>
    <![endif]-->
    <!-- custom CSS, additional to bootstrap CSS -->
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

            <div class="col-md-10" id="data-view">
    {{if eq .Cmd "view"}}
            {{template "artist-view" .Artist}}
    {{else if eq .Cmd "modify"}}
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
    <script src="/static/js/jquery.min.js"></script>
    <!-- Include all compiled plugins, or include individual files as needed -->
    <script src="/static/js/bootstrap.min.js"></script>
    <script src="/static/js/artistic.js"></script>
</body>
</html>
{{end}}

{{define "artist-create"}}
<div class="container-fluid" id="create-artist-form-div">
<form class="form-horizontal" role="form" method="post" id="create-artist-form" action="/artist/insert/">
    <fieldset>

    <div class="form-group has-success">
        <div class="col-md-2 control-label"><b>Name</b></div>
        <label for="first" class="col-md-1 control-label">First</label>
        <div class="col-md-2">
            <input type="text" class="form-control" id="first" name="first" value="" placeholder="first name" required />
        </div>
        
        <label for="middle" class="col-md-1 control-label">Middle</label>
        <div class="col-md-2">
            <input type="text" class="form-control" id="middle" name="middle" value="" placeholder="middle" />
        </div>
        
        <label for="last" class="col-md-1 control-label">Last</label>
        <div class="col-md-3">
            <input type="text" class="form-control" id="last" name="last" value="" placeholder="last" />
        </div>
    </div> <!-- form-group -->

    <div class="form-group">
        <div class="col-md-2 control-label"><b>Real Name</b></div>
        <label for="realfirst" class="col-md-1 control-label">First</label>
        <div class="col-md-2">
            <input type="text" class="form-control" id="realfirst" name="realfirst" value="" placeholder="first name" />
        </div>
        
        <label for="realmiddle" class="col-md-1 control-label">Middle</label>
        <div class="col-md-2">
            <input type="text" class="form-control" id="realmiddle" name="realmiddle" value="" placeholder="middle" />
        </div>
        
        <label for="reallast" class="col-md-1 control-label">Last</label>
        <div class="col-md-3">
            <input type="text" class="form-control" id="reallast" name="reallast" value="" placeholder="last" />
        </div>
    </div> <!-- form-group -->

    <div class="form-group">
        <label for="born" class="col-md-2 control-label">Born</label>
        <div class="col-md-3">
            <input type="date" class="form-control" id="born" name="born" value="" />
        </div>
        <label for="died" class="col-md-1 control-label">Died</label>
        <div class="col-md-3">
            <input type="date" class="form-control" id="died" name="died" value="" />
        </div>
    </div> <!-- form-group -->

    <div class="form-group">
        <label for="nationality" class="col-md-2 control-label">Nationality</label>
        <div class="col-md-3">
            <input type="text" class="form-control" id="nationality" name="nationality" value="" placeholder="nationality" />
        </div>
    </div> <!-- form-group -->

    <div class="form-group">
        <label for="biography" class="control-label">Biography</label>
        <textarea class="form-control" id="biography" name="biography" rows="5">Biography goes here... </textarea>
    </div> <!-- form-group -->

    <!-- TODO -->

    <div class="form-group form-inline">
        &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
        <div class="checkbox">
            <label><input type="checkbox" id="painter" name="painter" value="yes">&nbsp;&nbsp;Painter</label>
        </div> <!-- checkbox -->
        &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
        <div class="checkbox">
            <label><input type="checkbox" id="sculptor" name="sculptor" value="yes">&nbsp;&nbsp;Sculptor</label>
        </div> <!-- checkbox -->
        &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
        <div class="checkbox">
            <label><input type="checkbox" id="printmaker" name="printmaker" value="yes">&nbsp;&nbsp;Printmaker</label>
        </div> <!-- checkbox -->
        &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
        <div class="checkbox">
            <label><input type="checkbox" id="ceramicist" name="ceramicist" value="yes">&nbsp;&nbsp;Ceramicist</label>
        </div> <!-- checkbox -->
        &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
        <div class="checkbox">
            <label><input type="checkbox" id="architect" name="architect" value="yes">&nbsp;&nbsp;Architect</label>
        </div> <!-- checkbox -->
   </div> <!-- form-group -->

    <div class="form-group">
        <div class="col-md-4">
            <button class="btn btn-primary" type="submit" id="artist-submit">Create</button>
            <button class="btn btn-default" type="reset">Clear</button>
        </div>
        <div class="col-md-1 col-md-offset-3">
            <a class="btn btn-primary" href="/artists">
                <span class="glyphicon glyphicon-arrow-left"></span>&nbsp;&nbsp;Back
            </a>
        </div>
    </div> <!-- form-group -->

    </fieldset>
</form>
</div> <!-- container-fluid -->
{{end}}

{{define "artist-view"}}
<div class="container-fluid" id="view-artist-table-div">

    <div class="row">
        <h1 id="data-view-header">{{.Name}}</h1>
     <!--   <div class="col-md-2 control-label"><b>Name</b></div>
        <label for="name" class="col-md-1 control-label">Name</label>
        <div class="col-md-2">
            <input type="text" class="form-control" id="name" name="name" value="{{.Name}}" readonly />
        </div> -->
    </div>    

    <div class="row">
        <h3><span class="col-md-2">{{.Born}}</span><span class="col-md1">-</span><span class="col-md-2">{{.Died}}</span></h3>
    </div>    

    <div class="row">&nbsp;<!--empty row --></div>    

    <div class="row"> <b>{{.RealName}}</b> </div>    

    <div class="row">
         {{if eq .IsPainter true}}<div class="col-md-1"><b>Painter&nbsp;&nbsp;</b></div>{{end}}
         {{if eq .IsSculptor true}}<div class="col-md-1"><b>Sculptor&nbsp;&nbsp;</b></div>{{end}}
         {{if eq .IsPrintmaker true}}<div class="col-md-1"><b>Printmaker&nbsp;&nbsp;</b></div>{{end}}
         {{if eq .IsCeramicist true}}<div class="col-md-1"><b>Ceramicist&nbsp;&nbsp;</b></div>{{end}}
         {{if eq .IsArchitect true}}<div class="col-md-1"><b>Architect&nbsp;&nbsp;</b></div>{{end}}
    </div>    

    <div class="row">
        <p><div class="col-md-2"><b>Nationality</b></div><div class="col-md-10">{{.Nationality}}</div></p> 
    </div>    

    <div class="row">
        <div class="control-label"><b>Biography</b></div><br />
        <textarea class="form-control" id="biography" rows="10" readonly>{{.Biography}}</textarea>
      <!--  <div class="col-md-2"><b>Biography</b></div><br />
        <div class="col-md-6"><textarea rows="5" readonly>{{.Biography}}</textarea></div> -->
    </div>    

    <div class="row"> <div class="col-md-1 col-md-offset-9">
        <a type="button" class="btn btn-primary" href="/artists">
            <span class="glyphicon glyphicon-arrow-left"></span>&nbsp;&nbsp;Back
        </a>
    </div>
</div> <!-- container-fluid -->
{{end}}

{{define "artist-modify"}}
<div class="container-fluid" id="modify-artist-form-div">
    <div class="row">
        <h1 id="data-view-header">Modify Artist</h1>
    </div>

<form class="form-horizontal" role="form" method="post" id="modify-artist-form">
    <fieldset>

    <div class="row form-group has-success">
        <div class="col-md-2 control-label"><b>Name</b></div>
        <label for="first" class="col-md-1 control-label">First</label>
        <div class="col-md-2">
            <input type="text" class="form-control" id="first" name="first" value="{{.Name.First}}" 
                                                                placeholder="first name" required />
        </div>
        
        <label for="middle" class="col-md-1 control-label">Middle</label>
        <div class="col-md-2">
            <input type="text" class="form-control" id="middle" name="middle" value="{{.Name.Middle}}" placeholder="middle" />
        </div>
        
        <label for="last" class="col-md-1 control-label">Last</label>
        <div class="col-md-3">
            <input type="text" class="form-control" id="last" name="last" value="{{.Name.Last}}" placeholder="last" />
        </div>
    </div> <!-- form-group -->

    <div class="row form-group">
        <div class="col-md-2 control-label"><b>Real Name</b></div>
        <label for="realfirst" class="col-md-1 control-label">First</label>
        <div class="col-md-2">
            <input type="text" class="form-control" id="realfirst" name="realfirst" value="{{.RealName.First}}" />
        </div> 
        <label for="realmiddle" class="col-md-1 control-label">Middle</label>
        <div class="col-md-2">
            <input type="text" class="form-control" id="realmiddle" name="realmiddle" value="{{.RealName.Middle}}" />
        </div>
        
        <label for="reallast" class="col-md-1 control-label">Last</label>
        <div class="col-md-3">
            <input type="text" class="form-control" id="reallast" name="reallast" value="{{.RealName.Last}}" />
        </div>
    </div> <!-- form-group -->

    <div class="row form-group">
        <label for="born" class="col-md-2 control-label">Born</label>
        <div class="col-md-3">
            <input type="date" class="form-control" id="born" name="born" value="{{.Born}}" />
        </div>
        <label for="died" class="col-md-1 control-label">Died</label>
        <div class="col-md-3">
            <input type="date" class="form-control" id="died" name="died" value="{{.Died}}" />
        </div>
    </div> <!-- form-group -->

    <div class="row form-group">
        <label for="nationality" class="col-md-2 control-label">Nationality</label>
        <div class="col-md-3">
            <input type="text" class="form-control" id="nationality" name="nationality" value="{{.Nationality}}" />
        </div>
    </div> <!-- form-group -->

    <div class="row form-group">
        <label for="biography" class="control-label">Biography</label>
        <textarea class="form-control" id="biography" name="biography" rows="5">{{.Biography}}</textarea>
    </div> <!-- form-group -->

    <!-- TODO -->

    <div class="row form-group form-inline">
        &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
        <div class="col-md-2">
        <div class="checkbox">
            <label>
            {{if eq .IsPainter true}}
            <input type="checkbox" id="painter" name="painter" value="yes" checked>
            {{else}}
            <input type="checkbox" id="painter" name="painter" value="yes">
            {{end}}
            &nbsp;&nbsp;Painter</label>
        </div> <!-- checkbox -->
        </div>
        <div class="col-md-2">
        
        <div class="checkbox">
            <label class="text-right">
            {{if eq .IsSculptor true}}
            <input type="checkbox" id="sculptor" name="sculptor" value="yes" checked>
            {{else}}
            <input type="checkbox" id="sculptor" name="sculptor" value="yes">
            {{end}}
            &nbsp;&nbsp;Sculptor
            </label>
        </div> <!-- checkbox -->
        </div>
        <div class="col-md-2">
        
        <div class="checkbox">
            <label class="text-right">
            {{if eq .IsPrintmaker true}}
            <input type="checkbox" id="printmaker" name="printmaker" value="yes" checked>
            {{else}}
            <input type="checkbox" id="printmaker" name="printmaker" value="yes">
            {{end}}
            &nbsp;&nbsp;Printmaker
            </label>
        </div> <!-- checkbox -->
        </div>
        <div class="col-md-2">
        
        <div class="checkbox">
            <label>
            {{if eq .IsCeramicist true}}
            <input type="checkbox" id="ceramicist" name="ceramicist" value="yes" checked>
            {{else}}
            <input type="checkbox" id="ceramicist" name="ceramicist" value="yes">
            &nbsp;&nbsp;Ceramicist
            {{end}}
            </label>
        </div> <!-- checkbox -->
        </div>
        <div class="col-md-2">
        
        <div class="checkbox">
            <label>
            {{if eq .IsArchitect true}}
            <input type="checkbox" id="architect" name="architect" value="yes" checked>
            {{else}}
            <input type="checkbox" id="architect" name="architect" value="yes">
            {{end}}
            &nbsp;&nbsp;Architect
            </label>
        </div> <!-- checkbox -->
        </div>
   </div> <!-- form-group -->

    <div class="row form-group">
        <div class="col-md-4">
            <button class="btn btn-primary" type="submit" id="artist-submit">Modify</button>
            <button class="btn btn-default" type="reset">Clear</button>
        </div>
        <div class="col-md-1 col-md-offset-3">
            <a class="btn btn-primary" href="/artists">
                <span class="glyphicon glyphicon-arrow-left"></span>&nbsp;&nbsp;Back
            </a>
        </div>
    </div> <!-- form-group -->

    </fieldset>
</form>
</div> <!-- container-fluid -->

{{end}}
