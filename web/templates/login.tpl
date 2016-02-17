{{define "login"}}
<!DOCTYPE html>
<html lang="en">
<head>
{{template "htmlhead" "Login"}}
</head>
<body>
    <div class="container">
        <div class="row">
            <div class="col-md-2"></div>

            <div class="col-md-8">
                <div class="well well-sm">
                    <h1 class="text-center">Artistic</h1>
                    <p class="text-center text-info"> Artistic is a web application to manage art history resources.</p>
                </div> <!-- well -->
            </div> <!-- col-md-8 -->
            
            <div class="col-md-2"></div>
        </div> <!-- row -->

        <div class="row">
            <div class="col-md-4"></div>

            <div class="col-md-4">
                <form class="form-signin" role="form" id="signin_form" method="post">
                    <!--<h2 class="form-signin-heading">Please sign in</h2>-->
                    <input type="text" class="form-control" name="username" placeholder="Username" required autofocus>
                    <input type="password" class="form-control" name="password" placeholder="Password" required>
                    <button class="btn btn-lg btn-primary btn-block" id="signin_button" type="submit">Sign in</button>
                </form>
            </div>

            <div class="col-md-4"></div>
        </div> <!-- row -->
    </div> <!-- /container -->

{{template "insert-js"}}
</body>
</html>
{{end}}
