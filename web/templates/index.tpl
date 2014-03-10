{{define "navbar"}}
<!-- <nav class="navbar navbar-inverse" role="navigation"> -->
<nav class="navbar navbar-inverse navbar-static-top" role="navigation">
  <div class="container-fluid">
    <!-- Brand and toggle get grouped for better mobile display -->
    <div class="navbar-header">
      <a class="navbar-brand" href="#">Artistic</a>
    </div>

    <!-- Collect the nav links, forms, and other content for toggling -->
    <div class="collapse navbar-collapse" id="bs-example-navbar-collapse-1">
      <ul class="nav navbar-nav">
        <li class="active"><a href="#">Link</a></li>
        <li><a href="#">Link</a></li>
      </ul>
      <form class="navbar-form navbar-left" role="search">
        <div class="form-group">
          <input type="text" class="form-control" placeholder="Search">
        </div>
        <button type="submit" class="btn btn-primary">Submit</button>
      </form>
      <ul class="nav navbar-nav navbar-right">
        <li>
        <a href="#"><span class="glyphicon glyphicon-star"></span></a>
        </li>
        <li>
        <a href="#"><span class="glyphicon glyphicon-cog"></span></a>
        </li>
        <li>
        <a href="#"><span class="glyphicon glyphicon-copyright-mark"></span></a>
        </li>
        <li>
        <a href="#"><span class="glyphicon glyphicon-user"></span></a>
        </li>
    <!--
        <li class="dropdown">
          <a href="#" class="dropdown-toggle" data-toggle="dropdown">
          View<b class="caret"></b>
          </a>
          <ul class="dropdown-menu">
            <li><a href="#">License</a></li>
            <li><a href="#">About</a></li>
          </ul>
        </li>
    -->
    <!--
        <li class="dropdown">
          <a href="#" class="dropdown-toggle" data-toggle="dropdown">
          Logged in as admin<b class="caret"></b>
          </a>
          <ul class="dropdown-menu">
            <li><a href="#">Profile</a></li>
            <li><a href="#">Logout</a></li>
          </ul>
        </li>
    -->
        <li><p class="navbar-text">Logged in as admin</p></li>
        <li>
            <a href="#"><span class="glyphicon glyphicon-log-out"></span></a>
        </li>
      </ul>
    </div><!-- /.navbar-collapse -->
  </div><!-- /.container-fluid -->
</nav>
{{end}}


{{define "index"}}
<!DOCTYPE html>
<html lang="en">

  <head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <title>Artistic - Main Page</title>

    <!-- Bootstrap -->
  <!--  <link href="css/bootstrap.min.css" rel="stylesheet"> -->
    <link href="static/css/bootstrap.min.css" rel="stylesheet">

    <!-- HTML5 Shim and Respond.js IE8 support of HTML5 elements and media queries -->
    <!-- WARNING: Respond.js doesn't work if you view the page via file:// -->
    <!--[if lt IE 9]>
      <script src="https://oss.maxcdn.com/libs/html5shiv/3.7.0/html5shiv.js"></script>
      <script src="https://oss.maxcdn.com/libs/respond.js/1.4.2/respond.min.js"></script>
    <![endif]-->
  </head>

  <body>
    {{template "navbar"}}

    <div class="container-fluid">

    <div class="row">

        <div class="col-md-2" id="menu">
            <h1>Hello</h1>

            <p>Hello, World!</p>
        </div>

        <div class="col-md-4" id="data-list">
            <h1>Hello, World!</h1>

            <p>Hello, World!</p>
        </div>

        <div class="col-md-6" id="data-details">
            <h1>Hello, World!</h1>

            <p>Hello, World!</p>
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
