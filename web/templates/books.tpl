{{define "books"}}
{{$role := .User.Role}}
{{$name := totitle .Ptype}}
<!DOCTYPE html>
<html lang="en">
<head>
{{template "htmlhead" "Handle Books"}}
</head>

<body>
    {{template "navbar" .}}
    <div class="container-fluid">
        <div class="row">
            <div class="col-md-2" id="menu">
                <h1 id="menu-header"></h1>
                {{template "accordion"}}
            </div>

            <div class="col-md-10" id="data-list">
                <h1 id="data-list-header">{{$name}}s</h1>

    {{if ne $role "guest"}}
    {{template "add-button" $name}}
    {{end}}
                <br />

            {{if .Books}}
                <table class="table table-striped table-hover small" id="book-list-table">

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
                        <strong> {{.Num}} {{if eq .Num 1}} book {{else}} books {{end}} found. </strong> 
                    </td>
                    </tr>
                </tfoot>

                <tbody>
                  {{range $index, $element := .Books}}
                  {{ $cnt := add $index 1 }}
                  <tr id="book-row-{{$cnt}}">
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
    {{template "remove_book_modal" $name}}
{{end}}
    <!-- end of modals definition -->   

            {{else}}
                <p>No books found.</p>
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
{{template "remove-modal" .}}
{{end}}
