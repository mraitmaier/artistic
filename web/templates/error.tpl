
{{define "error404"}}
<!DOCTYPE html>
<html lang="en">

  <head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <title>Artistic - Page not Found</title>

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
    {{template "navbar" .Username}}

    <div class="container-fluid">

    <div class="row">

        <div class="col-md-2" id="menu">
            <h1 id="menu-header"></h1>
            {{template "accordion"}}
        </div>

        <div class="col-md-4" id="data-list">
            <h1 id="data-list-header">Error 404</h1>
            <p>This page has not been found!</p>
        </div>

        <div class="col-md-6" id="data-details">
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

    <script> </script>
  </body>
</html>
{{end}}

{{define "error"}}
<!DOCTYPE html>
<html lang="en">

  <head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <title>Artistic - Error</title>

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
    {{template "navbar" .Username}}

    <div class="container-fluid">

    <div class="row">

        <div class="col-md-2" id="menu">
            <h1 id="menu-header"></h1>
            {{template "accordion"}}
        </div>

        <div class="col-md-4" id="data-list">
            <h1 id="data-list-header">Error</h1>
            <p>Unknown error occured.</p>
        </div>

        <div class="col-md-6" id="data-details">
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

    <script> </script>
  </body>
</html>
{{end}}
