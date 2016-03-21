{{define "htmlhead"}}
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <title>Artistic - {{.}}</title> <!-- this is page title -->

    <!-- Bootstrap -->
    <link href="/static/css/bootstrap.min.css" rel="stylesheet">
    <link href="/static/css/dataTablesdataTables.bootstrap.min.css" rel="stylesheet">
    <!-- HTML5 Shim and Respond.js IE8 support of HTML5 elements and media queries -->
    <!-- WARNING: Respond.js doesn't work if you view the page via file:// -->
    <!--[if lt IE 9]>
        <script src="https://oss.maxcdn.com/libs/html5shiv/3.7.0/html5shiv.js"></script>
        <script src="https://oss.maxcdn.com/libs/respond.js/1.4.2/respond.min.js"></script>
    <![endif]-->
    <!-- custom CSS, additional to bootstrap -->
    <link href="/static/css/custom.css" rel="stylesheet">
{{end}}

{{define "add-button"}}
        <div id="new-{{tolower .}}-btn">
            <button type="button" class="btn btn-primary btn-sm" 
                    data-toggle="modal" data-target="#add{{.}}Modal">
                <span class="glyphicon glyphicon-plus"></span> &nbsp; Add a New 
                {{if eq . "Print"}} Graphic {{.}} {{else}} {{.}} {{end}}
            </button>
        </div>
{{end}}

{{define "insert-js"}}
    <!-- jQuery (necessary for Bootstrap's JavaScript plugins) -->
    <!--   <script src="https://ajax.googleapis.com/ajax/libs/jquery/2.1.0/jquery.min.js"></script> -->
    <script  src="/static/js/jquery.min.js"></script>
    <script  src="/static/js/jquery.dataTables.min.js"></script>
    <script  src="/static/js/jquery.validate.min.js"></script>
    <!-- Include all compiled plugins, or include individual files as needed -->
    <script src="/static/js/bootstrap.min.js"></script>
    <script src="/static/js/dataTables.bootstrap.min.js"></script>
    <!-- custom JS code -->
    <script src="/static/js/artistic.js"></script>
{{end}}

{{define  "accordion"}}
<div class="panel-group" id="main-menu" role="tablist" aria-multiselectable="true">

    <div class="panel panel-default">
        <div class="panel-heading" role="tab" id="artists-menu-header">
            <h4 class="panel-title">
                <a role="button" data-toggle="collapse" data-parent="#main-menu" href="#artists-collapse"  
                   aria-expanded="true" aria-controls="artists-collapse">Artists</a>
            </h4>
        </div>
    </div>
    <div class="panel-collapse collapse in" id="artists-collapse" role="tabpanel" aria-labelledby="artists-menu-header">
        <div class="panel-body small" >
                <p><a href="/painter"> Painters </a></p>
                <p><a href="/sculptor"> Sculptors </a></p>
                <p><a href="/printmaker"> Printmakers </a></p>
                <p><a href="/architect"> Architects </a></p>
        </div>
    </div>    

    <div class="panel panel-default">
        <div class="panel-heading" role="tab" id="artworks-menu-header">
            <h4 class="panel-title">
                <a role="button" data-toggle="collapse" data-parent="#main-menu" href="#artworks-collapse"   
                   aria-expanded="true" aria-controls="artworks-collapse">Artworks</a>
            </h4>
        </div>
    </div>
    <div class="panel-collapse collapse in" id="artworks-collapse" role="tabpanel" aria-labelledby="artworks-menu-header">
        <div class="panel-body small">
            <p><a href="/painting"> Paintings </a></p>
            <p><a href="/sculpture"> Sculptures </a></p>
            <p><a href="/print"> Graphic Prints </a></p>
            <p><a href="/building"> Buildings </a></p>
        </div>
    </div>

    <div class="panel panel-default">
        <div class="panel-heading" role="tab" id="other-menu-header">
            <h4 class="panel-title">
                <a role="button" data-toggle="collapse" data-parent="#main-menu" href="#other-collapse"
                   aria-expanded="tabpanel" aria-controls="other-collapse">Other</a>
            </h4>
        </div>
    </div>
    <div class="panel-collapse collapse" id="other-collapse" role="tabpanel" aria-labelledby="other-menu-header">
        <div class="panel-body small">
            <p><a href="/book">Books</a></p>
            <p><a href="/article">Articles</a></p>
            <p><a href="/dating">Datings</a></p>
            <p><a href="/technique">Techniques</a></p>
            <p><a href="/style">Styles</a></p>
        </div>
    </div>

</div>
{{end}}

{{define "navbar"}}
    <!--<nav class="navbar navbar-inverse" role="navigation"> -->
    <nav class="navbar navbar-default" role="navigation"> 
        <div class="container-fluid">
            <!-- Brand and toggle get grouped for better mobile display -->
            <div class="navbar-header">
                <a class="navbar-brand" href="/index" data-toggle="tooltip" data-placement="left" title="Home">Artistic
                <!--         <span class="glyphicon glyphicon-home"></span> -->
            </a>
            </div>

            <!-- Collect the nav links, forms, and other content for toggling -->
            <div class="collapse navbar-collapse" id="bs-example-navbar-collapse-1">

                <form class="navbar-form navbar-left" role="search" method="post" action="/search">
                    <div class="form-group">
                        <input type="hidden" name="search-type" value="{{.Ptype}}">
                        <input type="text" class="form-control" name="search-string" placeholder="Search">
                    </div>
                    <button type="submit" class="btn btn-primary">Submit</button>
                </form>
                
                <ul class="nav navbar-nav navbar-right">

                {{if eq .User.Role "admin"}}
                    <li class="dropdown">
                        <a href="#" class="dropdown-toggle" data-toggle="dropdown">Admin<b class="caret"></b> </a>
                        <ul class="dropdown-menu">
                            <li><a href="/user">Users</a></li>
                            <!--<li><a href="/log">Log</a></li> -->
                        </ul>
                    </li>
               {{end}}
 
                    <li>
                        <a href="#" data-toggle="tooltip" data-placement="left" title="About">
                            <span class="glyphicon glyphicon-star"></span>
                        </a>
                    </li>

                    <li>
                        <a href="#" data-toggle="tooltip" data-placement="left" title="Settings">
                            <span class="glyphicon glyphicon-cog" ></span>
                        </a>
                    </li>
        
                    <li>
                        <a href="/license" data-toggle="tooltip" data-placement="left" title="License">
                            <span class="glyphicon glyphicon-copyright-mark"></span>
                        </a>
                    </li>

                    <li>
                        <a href="/profile" data-toggle="tooltip" data-placement="left" title="User profile">
                            <span class="glyphicon glyphicon-user"></span>
                        </a>
                    </li>

                    <li><p class="navbar-text">Signed in as {{.User.Fullname}}</p></li>

                    <li>
                        <a href="/logout" data-toggle="tooltip" data-placement="left" title="Sign out">
                            <span class="glyphicon glyphicon-log-out"></span>
                        </a>
                    </li>
                </ul>
            </div><!-- /.navbar-collapse -->
        </div><!-- /.container-fluid -->
    </nav>
{{end}}

{{define "remove-modal"}}
<div class="modal fade" id="remove{{.}}Modal" tabindex="-1" role="dialog" aria-labelledby="remove{{.}}ModalLabel">
<div class="modal-dialog">
<div class="modal-content">

    <div class="modal-header">
    <div class="container-fluid">
        <div class="row">
            <h3 class="modal-title col-sm-8" id="remove{{.}}ModalLabel">
            Remove {{if eq . "Print"}} Graphic {{.}} {{else}} {{.}} {{end}}</h3>
            <button type="button" class="btn btn-primary btn-sm col-sm-2" id="removebtn"> Remove </button>
            <button type="button" class="btn btn-default btn-sm col-sm-2" data-dismiss="modal"> Cancel </button>
        </div> <!-- row -->
    </div> <!-- container-fluid -->
    </div> <!-- modal-header -->

    <div class="modal-body">
    <p> Would you really like to remove the {{if eq . "Print"}} graphic {{tolower .}} ' {{else}} {{tolower .}} '{{end}}
            <span id="removename"></span>'?</p>
    <form method="post" id="remove_{{tolower .}}_form">
        <input type="hidden" name="id" id="id" />
    </form>
    </div>
</div>
</div>
</div>
{{end}}

{{define "created-modified-modify"}}
            <div class="form-group form-group-sm small">
                <input type="hidden" id="created" name="created"></input>
                <div class="col-sm-3 text-right"><strong>Created</strong></div>
                <div id="createdm" name="createdm" class="col-sm-3 text-left">Error</div>
                <div class="col-sm-3 text-right"><strong>Last Modified</strong></div>
                <div id="modifiedm" name="modifiedm" class="col-sm-3 text-left">Error</div>
            </div>
{{end}}
