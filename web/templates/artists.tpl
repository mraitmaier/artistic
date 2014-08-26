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

            <p> {{template "artist-list" .Artists}}</p>

            <button type="button" class="btn btn-primary"
                    onclick="rerouteUsingGet('user', 'insert', '');">
            Add New Artist
            </button>
        </div>

    </div> <!-- row -->

    </div> <!-- container fluid -->
<!-- jQuery (necessary for Bootstrap's JavaScript plugins) -->
 <!--   <script 
        src="https://ajax.googleapis.com/ajax/libs/jquery/2.1.0/jquery.min.js">
        </script> -->
    <script  src="static/js/jquery.min.js"></script>

<!-- Include all compiled plugins, or include individual files as needed -->
    <script src="static/js/bootstrap.min.js"></script>
<!-- custom JS code -->
    <script src="static/js/artistic.js"></script>

    <script> </script>
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
                <a href="#" data-toggle="tooltip" data-placement="left" 
                            title="View artist details" id="view-artist-{{$id}}"
        onclick="return rerouteUsingGet('artist', 'view', {{$element.Id}});">
                    <span class="glyphicon glyphicon-eye-open"></span>
                </a>
                &nbsp;
                <a href="#" data-toggle="tooltip" data-placement="left" 
                            title="Edit artist" id="edit-artist-{{$id}}"
        onclick="return rerouteUsingGet('artist', 'modify', {{$element.Id}});">
                    <span class="glyphicon glyphicon-cog" ></span>
                </a>
                &nbsp;
                <a href="#" data-toggle="tooltip" data-placement="left" 
                            title="Delete artist" id="delete-artist-{{$id}}"
        onclick="return rerouteUsingGet('artist', 'delete', {{$element.Id}});">
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

