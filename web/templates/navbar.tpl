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

