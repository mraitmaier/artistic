{{define "htmlhead"}}
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <title>Artistic - {{.}}</title> <!-- this is page title -->

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
{{end}}

{{define "add-button"}}
        <div id="new-book-btn">
            <button type="button" class="btn btn-primary btn-sm" 
                    data-toggle="modal" data-target="{{printf "#add%sModal"  .}}">
                <span class="glyphicon glyphicon-plus"></span> &nbsp; Add a New {{.}}
            </button>
        </div>
{{end}}

{{define "insert-js"}}
    <!-- jQuery (necessary for Bootstrap's JavaScript plugins) -->
    <!--   <script src="https://ajax.googleapis.com/ajax/libs/jquery/2.1.0/jquery.min.js"></script> -->
    <script  src="/static/js/jquery.min.js"></script>
    <!-- Include all compiled plugins, or include individual files as needed -->
    <script src="/static/js/bootstrap.min.js"></script>
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

{{define "general"}}
{{$role := .User.Role}}
<!DOCTYPE html>
<html lang="en">
{{template "htmlhead" .PageTitle}}
<body>
    {{template "navbar" .}}
    <div class="container-fluid">
        <div class="row">
            <div class="col-md-2" id="menu">
                <h1 id="menu-header"></h1>
                {{template "accordion"}}
            </div>

            <div class="col-md-10" id="data-list">
                <h1 id="data-list-header">{{.ItemName}}s</h1> <!-- like "Book-s" or "Style-s" -->

    {{if ne $role "guest"}}
        {{template "add-button" .ItemName}}
    {{end}}
        <br>

            {{if .Items}}
                <table class="table table-striped table-hover small" id="{{tolower .ItemName}}-list-table">

                <thead>
                  <tr>
                    <th class="col-sm-1">#</th>
                    <th class="col-sm-4">Title</th>
                    <th class="col-sm-3">Author(s)</th>
                    <th class="col-sm-1">Year</th>
                    <th class="col-sm-2">Publisher</th>
                    <th class="col-sm-1">Actions</th>
                  </tr>
                </thead>

                <tfoot>
                    <tr class="bg-primary">
                    <td colspan="6"> 
                        <strong> 
                        {{.Num}} {{if eq .Num 1}} {{tolower .ItemName}} {{else}} {{tolower .ItemName}}s {{end}} found. 
                        </strong> 
                    </td>
                    </tr>
                </tfoot>

                <tbody>
                  {{range $index, $element := .Items}}
                  {{ $cnt := add $index 1 }}
                  <tr id="{{tolower .ItemName}}-row-{{$cnt}}"> <!-- id like "book-row-1" -->
                    <td>{{$cnt}}</td>
                    <td>{{$element.Title}}</td>
                    <td>{{$element.Authors}}</td>
                    <td>{{$element.Year}}</td>
                    <td>{{$element.Publisher}}</td>
				    <td>
  						<span data-toggle="tooltip" data-placement="up" title="View details">
                            <a data-toggle="modal" data-target="#viewBookModal"
                                       data-id="{{$element.ID.Hex}}"
                                       data-created="{{$element.Created}}"
                                       data-modified="{{$element.Modified}}"
                                       data-title="{{$element.Title}}"
                                       data-author="{{$element.Authors}}"
                                       data-edition="{{$element.Edition}}"
                                       data-publish="{{$element.Publisher}}"
                                       data-year="{{$element.Year}}"
                                       data-location="{{$element.Location}}"
                                       data-isbn="{{$element.ISBN}}"
                                       data-keyword="{{$element.Keywords}}"
                                       data-notes="{{$element.Notes}}"
                                       data-front="{{$element.Front}}">
                                 <span class="glyphicon glyphicon-eye-open"></span>
                            </a>
                       </span>            
                       {{if ne $role "guest"}}
                       &nbsp;&nbsp;
                       <span data-toggle="tooltip" data-placement="up" title="Modify details"> 
                            <a data-toggle="modal" data-target="#modifyBookModal"
                                       data-id="{{$element.ID.Hex}}"
                                       data-created="{{$element.Created}}"
                                       data-modified="{{$element.Modified}}"
                                       data-title="{{$element.Title}}"
                                       data-author="{{$element.Authors}}"
                                       data-edition="{{$element.Edition}}"
                                       data-publish="{{$element.Publisher}}"
                                       data-year="{{$element.Year}}"
                                       data-location="{{$element.Location}}"
                                       data-isbn="{{$element.ISBN}}"
                                       data-keyword="{{$element.Keywords}}"
                                       data-notes="{{$element.Notes}}"
                                       data-front="{{$element.Front}}">
                                 <span class="glyphicon glyphicon-edit"></span>
                            </a>
                       &nbsp;&nbsp;
                       <span data-toggle="tooltip" data-placement="up" title="Remove"> 
                            <a data-toggle="modal" data-target="#removeBookModal"
                                       data-id="{{$element.ID.Hex}}"  
                                       data-title="{{$element.Title}}">
                                <span class="glyphicon glyphicon-remove"></span>
                            </a>
                       </span>       
                        {{end}}
                    </td>
                  </tr>
                  {{end}}
                </tbody>
                </table>

    <!-- add modals -->
    {{template "view_book_modal"}}
{{if ne $role "guest"}}
    {{template "modify_book_modal"}}
    {{template "remove_book_modal"}}
{{end}}
    <!-- end of modals definition -->   

            {{else}}
                <p>No {{tolower .ItemName}}s found.</p>
            {{end}}

            </div>
        </div> <!-- row -->
    </div> <!-- container fluid -->

{{if ne $role "guest"}}
    {{template "add_book_modal"}}
{{end}}

{{template "insert-js"}}

    <script>

    $('#viewBookModal').on('show.bs.modal', function (event) {

        var button = $(event.relatedTarget);     // Button that triggered the modal
        // Extract info from data-* requirement attribute
        var id = button.data('id');  
        var year = button.data('year');

        // Update the modal's content. We'll use jQuery here, but you could use a data 
        // binding library or other methods instead.
        var modal = $(this)
        modal.find('.modal-title').text(button.data('title') + ' (' + year + ')');
        modal.find('.modal-body #authors').text(button.data('author'));
        modal.find('.modal-body #edition').text(button.data('edition'));
        modal.find('.modal-body #publisher').text(button.data('publish'));
        modal.find('.modal-body #location').text(button.data('location'));
        modal.find('.modal-body #isbn').text(button.data('isbn'));
        modal.find('.modal-body #year').text(button.data('year'));
        modal.find('.modal-body #keywords').text(button.data('keyword'));
        //modal.find('.modal-body #notes').text(notes);
        //modal.find('.modal-body #front').text(fronf);
        modal.find('.modal-body #createdv').text(button.data('created'));
        modal.find('.modal-body #modifiedv').text(button.data('modified'));
    })

{{if ne $role "guest"}}
    $('#modifyBookModal').on('show.bs.modal', function (event) {

        var button = $(event.relatedTarget);     // Button that triggered the modal
        // Extract info from data-* attribute
        var title = button.data('title');
        var created = button.data('created');

        // Update the modal's content. We'll use jQuery here, but you could use a data 
        // binding library or other methods instead.
        var modal = $(this)
        modal.find('.modal-title').text('Modify Book "' + title + '" Details');
        modal.find('.modal-body #hexid').val(button.data('id'));
        modal.find('.modal-body #title').val(title);
        modal.find('.modal-body #authors').val(button.data('author'));
        modal.find('.modal-body #edition').val(button.data('edition'));
        modal.find('.modal-body #publisher').val(button.data('publish'));
        modal.find('.modal-body #year').val(button.data('year'));
        modal.find('.modal-body #location').val(button.data('location'));
        modal.find('.modal-body #isbn').val(button.data('isbn'));
        modal.find('.modal-body #keywords').val(button.data('keyword'));
        //modal.find('.modal-body #notes').val(notes);
        //modal.find('.modal-body #front').val(front);
        modal.find('.modal-body #created').val(created);
        modal.find('.modal-body #createdm').text(created);
        modal.find('.modal-body #modifiedm').text(button.data('modified'));
    })

// Handle the removals using modal pop-up 
   $('#removeBookModal').on('show.bs.modal', function(event) {
    
        var button = $(event.relatedTarget);
        var id = button.data('id');
        var title = button.data('title');
        // Update the modal's content. We'll use jQuery here, but you could use a data binding library 
        // or other methods instead.
        var modal = $(this);
        modal.find('.modal-body #removename').text(title);
        modal.find('.modal-body #hexid').val(id);
        modal.find('.modal-body #title').val(title);

        // Let's define the 'remove' button onclick() callback... 
        var url = '/book/' + id + '/delete';
        $('#removebtn').on('click', function(e) { 
            postForm('remove_book_form', url); 
            $('#removeBookModal').modal('hide');
        });
   });

    // This should post a form to modify book
    var modifyBook = function(form_id, id) {
        var url = '/book/' + id + '/put';
        postForm(form_id, url);
    }
{{end}}
    </script>
</body>
</html>
{{end}}

{{define "add_book_modal"}}
<div class="modal fade" id="addBookModal" tabindex="-1" role="dialog" aria-labelledby="addBookModalLabel">
<div class="modal-dialog modal-lg">
<div class="modal-content">

    <div class="modal-header">
    <div class="container-fluid">
        <div class="row">
            <h3 class="modal-title col-sm-8" id="addBookModalLabel">Add a New Book</h3>
            <button type="button" class="btn btn-primary btn-sm col-sm-2" 
                    onclick="postForm('add-book-form', '/book'); $('#addBookModal').modal('hide');"> Add </button>
            <button type="button" class="btn btn-default btn-sm col-sm-2" data-dismiss="modal">Cancel</button>
        </div> <!-- row -->
    </div> <!-- container-fluid -->
    </div> <!-- modal-header -->

    <div class="modal-body">
    <form class="form-horizontal" role="form" method="post" id="add-book-form">
    <fieldset>

        <div class="form-group form-group-sm has-error">
            <label for="title" class="col-sm-3 control-label">Title</label>
            <div class="col-sm-9">
                <input type="text" class="form-control" id="title" name="title" placeholder="Book title" required />
            </div>
        </div> <!-- form-group -->

        <div class="form-group form-group-sm">
            <label for="authors" class="col-sm-3 control-label">Author(s)</label>
            <div class="col-sm-9">
                <input type="text" class="form-control" id="authors" name="authors" placeholder="Authors" />
            </div>
        </div> <!-- form-group -->

        <div class="form-group form-group-sm">
            <label for="edition" class="col-sm-3 control-label">Edition</label>
            <div class="col-sm-3">
                <input type="text" class="form-control" id="edition" name="edition" placeholder="Edition" />
            </div>
            <label for="isbn" class="col-sm-3 control-label">ISBN</label>
            <div class="col-sm-3">
                <input type="text" class="form-control" id="isbn" name="isbn" placeholder="ISBN number" />
            </div>
        </div> <!-- form-group -->

        <div class="form-group form-group-sm">
            <label for="publisher" class="col-sm-3 control-label">Publisher</label>
            <div class="col-sm-9">
                <input type="text" class="form-control" id="publisher" name="publisher" placeholder="Publisher" />
            </div>
        </div> <!-- form-group -->

        <div class="form-group form-group-sm">
            <label for="year" class="col-sm-3 control-label">Year</label>
            <div class="col-sm-3">
                <input type="text" class="form-control" id="year" name="year" placeholder="Year" />
            </div>
            <label for="location" class="col-sm-3 control-label"> Location </label>
            <div class="col-sm-3">
                <input type="text" class="form-control" id="location" name="location" placeholder="Location" />
            </div>
        </div> <!-- form-group -->

        <div class="form-group form-group-sm">
            <label for="keywords" class="col-sm-3 control-label"> Keywords </label>
            <div class="col-sm-9">
                <input type="text" class="form-control" id="keywords" name="keywords" />
            </div>
        </div> <!-- form-group -->

        <!-- TODO:  notes, front -->

    </fieldset>
    </form>

    </div> <!-- modal-body -->

</div>
</div>
</div>
{{end}}

{{define "view_book_modal"}}
<div class="modal fade" id="viewBookModal" tabindex="-1" role="dialog" aria-lebeleledby="ViewBookModalLabel">
<div class="modal-dialog modal-lg">
<div class="modal-content">

    <div class="modal-header">
    <div class="container-fluid">
        <div class="row">
            <h3 class="modal-title col-md-10" id="viewBookModalLabel"></h3>
            <button type="button" class="btn btn-default btn-sm col-md-2" data-dismiss="modal">Close</button>
        </div> <!-- row -->
    </div> <!-- container-fluid -->
    </div> <!-- modal-header -->

    <div class="modal-body">
        <div class="container-fluid" id="view-book-table-div">

        <div class="row">
             <table id="view-user-table" class="table table-hover small">
             <tbody>
                 <tr> <td class="col-sm-3"><strong>Author(s)<strong/></td>
                      <td class="col-sm-9"><span id="authors"></span></td> </tr>
                 <tr> <td class="col-sm-3"><strong>Edition<strong/></td>
                      <td class="col-sm-9"><span id="edition"></span></td> </tr>
                 <tr> <td class="col-sm-3"><strong>Publisher<strong/></td>
                      <td class="col-sm-9"><span id="publisher"></span></td> </tr>
                 <tr> <td class="col-sm-3"><strong>Year<strong/></td>
                      <td class="col-sm-9"><span id="year"></span></td> </tr>
                 <tr> <td class="col-sm-3"><strong>Location<strong/></td>
                      <td class="col-sm-9"><span id="location"></span></td> </tr>
                 <tr> <td class="col-sm-3"><strong>ISBN Number<strong/></td>
                      <td class="col-sm-9"><span id="isbn"></span></td> </tr>
                 <tr> <td class="col-sm-3"><strong>Keywords<strong/></td>
                      <td class="col-sm-9"><span id="keywords"></span></td> </tr>
                 <tr> <td class="col-sm-3"><strong>Created</strong></td> 
                      <td class="col-sm-9" id="createdv"></td> </tr> 
                 <tr> <td class="col-sm-3"><strong>Last Modified</strong></td> 
                      <td class="col-sm-9" id="modifiedv"></td> </tr>
             </tbody>
             </table>
        </div> <!-- row -->
        </div> <!-- container-fluid -->
    </div> <!-- modal-body -->

</div>
</div>
</div>
{{end}}

{{define "modify_book_modal"}}
<div class="modal fade" id="modifyBookModal" tabindex="-1" role="dialog" aria-lebeleledby="modifyBookModalLabel">
<div class="modal-dialog modal-lg">
<div class="modal-content" />

    <div class="modal-header">
    <div class="container-fluid">
        <div class="row">
            <h3 class="modal-title col-md-8" id="modifyBookModalLabel">Empty Book Details</h3>
             <button type="button" class="btn btn-primary btn-sm col-sm-2" 
                     onclick="modifyBook('modify-book-form', $('#hexid').val()); $('#modifyBookModal').modal('hide');">
                     Modify
             </button>
            <button type="button" class="btn btn-default btn-sm col-md-2" data-dismiss="modal">Cancel</button>
        </div> <!-- row -->
    </div> <!-- container-fluid -->
    </div> <!-- modal-header -->

    <div class="modal-body">
    <div class="container-fluid">
        <div class="row">

        <form id="modify-book-form" class="form-horizontal">

        <fieldset>
            <input type="hidden" id="hexid" name="hexid" />

            <div class="form-group form-group-sm has-error">
                <label for="title" class="col-sm-3 control-label">Title</label>
                <div class="col-sm-9">
                    <input type="text" class="form-control" id="title" name="title" required />
                </div>
            </div> <!-- form-group -->

            <div class="form-group form-group-sm">
                <label for="authors" class="col-sm-3 control-label">Author(s)</label>
                <div class="col-sm-9">
                    <input type="text" class="form-control" id="authors" name="authors" />
                </div>
            </div> <!-- form-group -->

            <div class="form-group form-group-sm">
                <label for="edition" class="col-sm-3 control-label">Edition</label>
                <div class="col-sm-3">
                    <input type="text" class="form-control" id="edition" name="edition" />
                </div>
                <label for="isbn" class="col-sm-3 control-label">ISBN</label>
                <div class="col-sm-3">
                    <input type="text" class="form-control" id="isbn" name="isbn" />
                </div>
            </div> <!-- form-group -->

            <div class="form-group form-group-sm">
                <label for="publisher" class="col-sm-3 control-label">Publisher</label>
                <div class="col-sm-9">
                    <input type="text" class="form-control" id="publisher" name="publisher" />
                </div>
            </div> <!-- form-group -->

            <div class="form-group form-group-sm">
                <label for="year" class="col-sm-3 control-label">Year</label>
                <div class="col-sm-3">
                    <input type="text" class="form-control" id="year" name="year" />
                </div>
                <label for="location" class="col-sm-3 control-label"> Location </label>
                <div class="col-sm-3">
                    <input type="text" class="form-control" id="location" name="location" />
                </div>
            </div> <!-- form-group -->

            <div class="form-group form-group-sm">
                <label for="keywords" class="col-sm-3 control-label"> Keywords </label>
                <div class="col-sm-9">
                    <input type="text" class="form-control" id="keywords" name="keywords" />
                </div>
            </div> <!-- form-group -->

            <!-- TODO:  notes, front -->

            <div class="form-group form-group-sm small">
                <input type="hidden" id="created" name="created"></input>
                <div class="col-sm-3 text-right"><strong>Created</strong></div>
                <div id="createdm" name="createdm" class="col-sm-3 text-left">Error</div>
                <div class="col-sm-3 text-right"><strong>Last Modified</strong></div>
                <div id="modifiedm" name="modifiedm" class="col-sm-3 text-left">Error</div>
            </div>

        </fieldset>
        </form>
        </div> <!-- row -->
    </div> <!-- container-fluid -->
    </div> <!-- modal-body -->

</div>
</div>
</div>
{{end}}

{{define "remove_book_modal"}}
<div class="modal fade" id="removeBookModal" tabindex="-1" role="dialog" aria-labelledby="removeBookModalLabel">
<div class="modal-dialog">
<div class="modal-content">

    <div class="modal-header">
    <div class="container-fluid">
        <div class="row">
            <h3 class="modal-title col-sm-8" id="removeBookModalLabel">Remove Book</h3>
            <button type="button" class="btn btn-primary btn-sm col-sm-2" id="removebtn"> Remove </button>
            <button type="button" class="btn btn-default btn-sm col-sm-2" data-dismiss="modal"> Cancel </button>
        </div> <!-- row -->
    </div> <!-- container-fluid -->
    </div> <!-- modal-header -->

    <div class="modal-body">
    <p> Would you really like to remove the book '<span id="removename"></span>'?</p>
    <form method="post" id="remove_book_form">
        <input type="hidden" name="id" id="id" />
        <input type="hidden" name="name" id="name" />
    </form>
    </div>
</div>
</div>
</div>
{{end}}


