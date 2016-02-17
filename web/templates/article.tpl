{{define "articles"}}
{{$role := .User.Role}}
{{$name := totitle .Ptype}}
<!DOCTYPE html>
<html lang="en">
<head>
    {{template "htmlhead" "Handle Articles"}}
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

            {{if .Articles}}
                <table class="table table-striped table-hover small" id="article-list-table">

                <thead>
                  <tr>
                    <th class="col-sm-1">#</th>
                    <th class="col-sm-4">Title</th>
                    <th class="col-sm-3">Author(s)</th>
                    <th class="col-sm-2">Publication</th>
                    <th class="col-sm-1">Year</th>
                    <th class="col-sm-1">Actions</th>
                  </tr>
                </thead>

                <tfoot>
                    <tr class="bg-primary">
                    <td colspan="6"> 
                        <strong> {{.Num}} {{if eq .Num 1}} {{.Ptype}} {{else}} {{.Ptype}}s {{end}} found. </strong> 
                    </td>
                    </tr>
                </tfoot>

                <tbody>
                  {{range $index, $element := .Articles}}
                  {{ $cnt := add $index 1 }}
                  <tr id="article-row-{{$cnt}}">
                    <td>{{$cnt}}</td>
                    <td>{{$element.Title}}</td>
                    <td>{{$element.Authors}}</td>
                    <td>{{$element.Publication}}</td>
                    <td>{{$element.Year}}</td>
				    <td>
  						<span data-toggle="tooltip" data-placement="up" title="View details">
                            <a data-toggle="modal" data-target="#viewArticleModal"
                                       data-id="{{$element.ID.Hex}}"
                                       data-created="{{$element.Created}}"
                                       data-modified="{{$element.Modified}}"
                                       data-title="{{$element.Title}}"
                                       data-author="{{$element.Authors}}"
                                       data-publication="{{$element.Publication}}"
                                       data-volume="{{$element.Volume}}"
                                       data-issue="{{$element.Issue}}"
                                       data-publish="{{$element.Publisher}}"
                                       data-year="{{$element.Year}}"
                                       data-location="{{$element.Location}}"
                                       data-issn="{{$element.ISSN}}"
                                       data-link="{{$element.Link}}"
                                       data-keyword="{{$element.Keywords}}"
                                       data-notes="{{$element.Notes}}">
                                 <span class="glyphicon glyphicon-eye-open"></span>
                            </a>
                       </span>            
                       {{if ne $role "guest"}}
                       &nbsp;&nbsp;
                       <span data-toggle="tooltip" data-placement="up" title="Modify details"> 
                            <a data-toggle="modal" data-target="#modifyArticleModal"
                                       data-id="{{$element.ID.Hex}}"
                                       data-created="{{$element.Created}}"
                                       data-modified="{{$element.Modified}}"
                                       data-title="{{$element.Title}}"
                                       data-author="{{$element.Authors}}"
                                       data-publication="{{$element.Publication}}"
                                       data-volume="{{$element.Volume}}"
                                       data-issue="{{$element.Issue}}"
                                       data-publish="{{$element.Publisher}}"
                                       data-year="{{$element.Year}}"
                                       data-location="{{$element.Location}}"
                                       data-issn="{{$element.ISSN}}"
                                       data-link="{{$element.Link}}"
                                       data-keyword="{{$element.Keywords}}"
                                       data-notes="{{$element.Notes}}">
                                 <span class="glyphicon glyphicon-edit"></span>
                            </a>
                       &nbsp;&nbsp;
                       <span data-toggle="tooltip" data-placement="up" title="Remove"> 
                            <a data-toggle="modal" data-target="#removeArticleModal"
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
    {{template "view_article_modal"}}
{{if ne $role "guest"}}
    {{template "modify_article_modal"}}
    {{template "remove_article_modal" $name}}
{{end}}
    <!-- end of modals definition -->   

            {{else}}
                <p>No {{.Ptype}}s found.</p>
            {{end}}

            </div>
        </div> <!-- row -->
    </div> <!-- container fluid -->

{{if ne $role "guest"}}
    {{template "add_article_modal"}}
{{end}}

{{template "insert-js"}}

    <script>

    $('#viewArticleModal').on('show.bs.modal', function (event) {

        var button = $(event.relatedTarget);     // Button that triggered the modal
        // Extract info from data-* requirement attribute
        var id = button.data('id');  
        var lnk = button.data('link');

        // Update the modal's content. We'll use jQuery here, but you could use a data 
        // binding library or other methods instead.
        var modal = $(this)
        modal.find('.modal-title').text(button.data('title'));
        modal.find('.modal-body #authors').text(button.data('author'));
        modal.find('.modal-body #publication').text(button.data('publication'));
        modal.find('.modal-body #volume').text(button.data('volume'));
        modal.find('.modal-body #issue').text(button.data('issue'));
        modal.find('.modal-body #publisher').text(button.data('publish'));
        modal.find('.modal-body #location').text(button.data('location'));
        modal.find('.modal-body #issn').text(button.data('issn'));
        modal.find('.modal-body #year').text(button.data('year'));
        // if link is not empty, display it in a new tab; otherwise just hide the whole table row
        if (lnk !== "") {
            modal.find('.modal-body #link').attr('href', button.data('link'));
            modal.find('.modal-body #link').attr('target', '_blank');
            modal.find('.modal-body #link').css('text-decoration', 'none');
        } else {
            //modal.find('.modal-body #online-link').css('display', 'none');
            modal.find('.modal-body #online-link').hide();
        }
        modal.find('.modal-body #keywords').text(button.data('keyword'));
        //modal.find('.modal-body #notes').text(notes);
        modal.find('.modal-body #createdv').text(button.data('created'));
        modal.find('.modal-body #modifiedv').text(button.data('modified'));
    })

{{if ne $role "guest"}}
    $('#modifyArticleModal').on('show.bs.modal', function (event) {

        var button = $(event.relatedTarget);     // Button that triggered the modal
        // Extract info from data-* attribute
        var title = button.data('title');
        var created = button.data('created');

        // Update the modal's content. We'll use jQuery here, but you could use a data 
        // binding library or other methods instead.
        var modal = $(this)
        modal.find('.modal-title').text('Modify Article "' + title + '" Details');
        modal.find('.modal-body #hexid').val(button.data('id'));
        modal.find('.modal-body #title').val(title);
        modal.find('.modal-body #authors').val(button.data('author'));
        modal.find('.modal-body #publication').val(button.data('publication'));
        modal.find('.modal-body #volume').val(button.data('volume'));
        modal.find('.modal-body #issue').val(button.data('issue'));
        modal.find('.modal-body #publisher').val(button.data('publish'));
        modal.find('.modal-body #year').val(button.data('year'));
        modal.find('.modal-body #location').val(button.data('location'));
        modal.find('.modal-body #issn').val(button.data('issn'));
        modal.find('.modal-body #link').val(button.data('link'));
        modal.find('.modal-body #keywords').val(button.data('keyword'));
        //modal.find('.modal-body #notes').val(notes);
        modal.find('.modal-body #created').val(created);
        modal.find('.modal-body #createdm').text(created);
        modal.find('.modal-body #modifiedm').text(button.data('modified'));
    })

// Handle the removals using modal pop-up 
   $('#removeArticleModal').on('show.bs.modal', function(event) {
    
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
        var url = '/article/' + id + '/delete';
        $('#removebtn').on('click', function(e) { 
            postForm('remove_article_form', url); 
            $('#removeArticleModal').modal('hide');
        });
   });

    // This should post a form to modify article
    var modifyArticle = function(form_id, id) {
        var url = '/article/' + id + '/put';
        postForm(form_id, url);
    }
{{end}}
    </script>
</body>
</html>
{{end}}

{{define "add_article_modal"}}
<div class="modal fade" id="addArticleModal" tabindex="-1" role="dialog" aria-labelledby="addArticleModalLabel">
<div class="modal-dialog modal-lg">
<div class="modal-content">

    <div class="modal-header">
    <div class="container-fluid">
        <div class="row">
            <h3 class="modal-title col-sm-8" id="addArticleModalLabel">Add a New Article</h3>
            <button type="button" class="btn btn-primary btn-sm col-sm-2" 
                    onclick="postForm('add-article-form', '/article'); $('#addArticleModal').modal('hide');"> Add </button>
            <button type="button" class="btn btn-default btn-sm col-sm-2" data-dismiss="modal">Cancel</button>
        </div> <!-- row -->
    </div> <!-- container-fluid -->
    </div> <!-- modal-header -->

    <div class="modal-body">
    <form class="form-horizontal" role="form" method="post" id="add-article-form">
    <fieldset>

        <div class="form-group form-group-sm has-error">
            <label for="title" class="col-sm-3 control-label">Title</label>
            <div class="col-sm-9">
                <input type="text" class="form-control" id="title" name="title" placeholder="Article title" required />
            </div>
        </div> <!-- form-group -->

        <div class="form-group form-group-sm">
            <label for="authors" class="col-sm-3 control-label">Author(s)</label>
            <div class="col-sm-9">
                <input type="text" class="form-control" id="authors" name="authors" placeholder="Authors" />
            </div>
        </div> <!-- form-group -->

        <div class="form-group form-group-sm">
            <label for="publication" class="col-sm-3 control-label">Publication</label>
            <div class="col-sm-9">
                <input type="text" class="form-control" id="publication" name="publication" placeholder="Publication" />
            </div>
        </div> <!-- form-group -->

        <div class="form-group form-group-sm">
            <label for="volume" class="col-sm-3 control-label">Volume</label>
            <div class="col-sm-3">
                <input type="text" class="form-control" id="volume" name="volume" placeholder="Volume" />
            </div>
            <label for="issue" class="col-sm-3 control-label">Issue</label>
            <div class="col-sm-3">
                <input type="text" class="form-control" id="issue" name="issue" placeholder="Issue" />
            </div>
        </div> <!-- form-group -->

        <div class="form-group form-group-sm">
            <label for="year" class="col-sm-3 control-label">Year</label>
            <div class="col-sm-3">
                <input type="text" class="form-control" id="year" name="year" placeholder="Year" />
            </div>
            <label for="issn" class="col-sm-3 control-label">ISSN</label>
            <div class="col-sm-3">
                <input type="text" class="form-control" id="issn" name="issn" placeholder="ISSN number" />
            </div>
        </div> <!-- form-group -->

        <div class="form-group form-group-sm">
            <label for="publisher" class="col-sm-3 control-label">Publisher</label>
            <div class="col-sm-3">
                <input type="text" class="form-control" id="publisher" name="publisher" placeholder="Publisher" />
            </div>
            <label for="location" class="col-sm-3 control-label"> Location </label>
            <div class="col-sm-3">
                <input type="text" class="form-control" id="location" name="location" placeholder="Location" />
            </div>
        </div> <!-- form-group -->

        <div class="form-group form-group-sm">
            <label for="link" class="col-sm-3 control-label"> Link </label>
            <div class="col-sm-9">
                <input type="text" class="form-control" id="link" name="link" placeholder="Link" />
            </div>
        </div> <!-- form-group -->

        <div class="form-group form-group-sm">
            <label for="keywords" class="col-sm-3 control-label"> Keywords </label>
            <div class="col-sm-9">
                <input type="text" class="form-control" id="keywords" name="keywords" />
            </div>
        </div> <!-- form-group -->

        <!-- TODO:  notes -->

    </fieldset>
    </form>

    </div> <!-- modal-body -->

</div>
</div>
</div>
{{end}}

{{define "view_article_modal"}}
<div class="modal fade" id="viewArticleModal" tabindex="-1" role="dialog" aria-lebeleledby="ViewArticleModalLabel">
<div class="modal-dialog modal-lg">
<div class="modal-content">

    <div class="modal-header">
    <div class="container-fluid">
        <div class="row">
            <h3 class="modal-title col-md-10" id="viewArticleModalLabel"></h3>
            <button type="button" class="btn btn-default btn-sm col-md-2" data-dismiss="modal">Close</button>
        </div> <!-- row -->
    </div> <!-- container-fluid -->
    </div> <!-- modal-header -->

    <div class="modal-body">
        <div class="container-fluid" id="view-article-table-div">

        <div class="row">
             <table id="view-article-table" class="table table-hover small">
             <tbody>
                 <tr> <td class="col-sm-3"><strong>Author(s)<strong/></td>
                      <td class="col-sm-9"><span id="authors"></span></td> </tr>
                 <tr> <td class="col-sm-3"><strong>Publication<strong/></td>
                      <td class="col-sm-9"><span id="publication"></span></td> </tr>
                 <tr> <td class="col-sm-3"><strong>Volume<strong/></td>
                      <td class="col-sm-9"><span id="volume"></span></td> </tr>
                 <tr> <td class="col-sm-3"><strong>Issue<strong/></td>
                      <td class="col-sm-9"><span id="issue"></span></td> </tr>
                 <tr> <td class="col-sm-3"><strong>Year<strong/></td>
                      <td class="col-sm-9"><span id="year"></span></td> </tr>
                 <tr> <td class="col-sm-3"><strong>Publisher<strong/></td>
                      <td class="col-sm-9"><span id="publisher"></span></td> </tr>
                 <tr> <td class="col-sm-3"><strong>Location<strong/></td>
                      <td class="col-sm-9"><span id="location"></span></td> </tr>
                 <tr> <td class="col-sm-3"><strong>ISSN Number<strong/></td>
                      <td class="col-sm-9"><span id="issn"></span></td> </tr>
                 <tr> <td class="col-sm-3"><strong>Keywords<strong/></td>
                      <td class="col-sm-9"><span id="keywords"></span></td> </tr>
                 <tr id="online-link"> <td class="col-sm-3"><strong><strong/>Link</td>
                      <td class="col-sm-9"><a id="link">View Online</a></td></tr>
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

{{define "modify_article_modal"}}
<div class="modal fade" id="modifyArticleModal" tabindex="-1" role="dialog" aria-lebeleledby="modifyArticleModalLabel">
<div class="modal-dialog modal-lg">
<div class="modal-content" />

    <div class="modal-header">
    <div class="container-fluid">
        <div class="row">
            <h3 class="modal-title col-md-8" id="modifyArticleModalLabel">Empty Article Details</h3>
             <button type="button" class="btn btn-primary btn-sm col-sm-2" 
                     onclick="modifyArticle('modify-article-form', $('#hexid').val()); $('#modifyArticleModal').modal('hide');">
                     Modify
             </button>
            <button type="button" class="btn btn-default btn-sm col-md-2" data-dismiss="modal">Cancel</button>
        </div> <!-- row -->
    </div> <!-- container-fluid -->
    </div> <!-- modal-header -->

    <div class="modal-body">
    <div class="container-fluid">
        <div class="row">

        <form id="modify-article-form" class="form-horizontal">

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
                <label for="publication" class="col-sm-3 control-label">Publication</label>
                <div class="col-sm-9">
                    <input type="text" class="form-control" id="publication" name="publication" placeholder="Publication" />
                </div>
            </div> <!-- form-group -->

            <div class="form-group form-group-sm">
                <label for="volume" class="col-sm-3 control-label">Volume</label>
                <div class="col-sm-3">
                    <input type="text" class="form-control" id="volume" name="volume" placeholder="Volume" />
                </div>
                <label for="issue" class="col-sm-3 control-label">Issue</label>
                <div class="col-sm-3">
                    <input type="text" class="form-control" id="issue" name="issue" placeholder="Issue" />
                </div>
            </div> <!-- form-group -->

            <div class="form-group form-group-sm">
                <label for="year" class="col-sm-3 control-label">Year</label>
                <div class="col-sm-3">
                    <input type="text" class="form-control" id="year" name="year" placeholder="Year" />
                </div>
                <label for="issn" class="col-sm-3 control-label">ISSN</label>
                <div class="col-sm-3">
                    <input type="text" class="form-control" id="issn" name="issn" placeholder="ISSN number" />
                </div>
            </div> <!-- form-group -->

            <div class="form-group form-group-sm">
                <label for="publisher" class="col-sm-3 control-label">Publisher</label>
                <div class="col-sm-3">
                    <input type="text" class="form-control" id="publisher" name="publisher" placeholder="Publisher" />
                </div>
                <label for="location" class="col-sm-3 control-label"> Location </label>
                <div class="col-sm-3">
                    <input type="text" class="form-control" id="location" name="location" placeholder="Location" />
                </div>
            </div> <!-- form-group -->

            <div class="form-group form-group-sm">
                <label for="link" class="col-sm-3 control-label"> Link </label>
                <div class="col-sm-9">
                    <input type="text" class="form-control" id="link" name="link" placeholder="Link" />
                </div>
            </div> <!-- form-group -->

            <div class="form-group form-group-sm">
                <label for="keywords" class="col-sm-3 control-label"> Keywords </label>
                <div class="col-sm-9">
                    <input type="text" class="form-control" id="keywords" name="keywords" />
                </div>
            </div> <!-- form-group -->

            <!-- TODO:  notes -->

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

{{define "remove_article_modal"}}
{{template "remove-modal" .}}
{{end}}
