{{define "error404"}}
<!DOCTYPE html>
<html lang="en">
<head>
{{template "htmlhead" "Page Not Found"}}
</head>
<body>
    {{template "navbar" .}}
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
{{template "insert-js"}}
</body>
</html>
{{end}}

{{define "error"}}
<!DOCTYPE html>
<html lang="en">
<head>
{{template "htmlhead" "Unknown Error"}}
</head>
<body>
    {{template "navbar" .}}

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
{{template "insert-js"}}
</body>
</html>
{{end}}
