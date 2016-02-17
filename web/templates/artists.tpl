{{define "artists"}}
{{$role := .User.Role}}
{{$name := get_artist_type .Type}}
<!DOCTYPE html>
<html lang="en">
<head>
{{template "htmlhead" "Handle Artists"}}
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
        {{template "add-button" "Artist"}} <!-- XXX leave it alone! -->
        {{end}}
            <br />

            {{if .Artists}}
                <table class="table table-striped table-hover small" id="artist-list-table">

                <thead>
                  <tr>
                    <th class="col-sm-1">#</th>
                    <th class="col-sm-2">Name</th>
                    <th class="col-sm-2">RealName</th>
                    <th class="col-sm-2">Born</th>
                    <th class="col-sm-2">Died</th>
                    <th class="col-sm-2">Nationality</th>
                    <th class="col-sm-1">Actions</th>
                  </tr>
                </thead>

                <tfoot>
                    <tr class="bg-primary">
                    <td colspan="7"> 
                        <strong>
                        {{.Num}} {{if eq .Num 1}} {{tolower $name}} {{else}} {{tolower $name}}s {{end}} found.
                        </strong> 
                    </td>
                    </tr>
                </tfoot>

                <tbody>
                  {{range $index, $element := .Artists}}
                  {{ $cnt := add $index 1 }}
                  <tr id="artist-row-{{$cnt}}">
                    <td>{{$cnt}}</td>
                    <td>{{$element.Name}}</td>
                    <td>{{$element.RealName}}</td>
                    <td>{{$element.Born}}</td>
                    <td>{{$element.Died}}</td>
                    <td>{{$element.Nationality}}</td>
				    <td>
  						<span data-toggle="tooltip" data-placement="up" title="View details">
                            <a data-toggle="modal" data-target="#viewArtistModal"
                                       data-id="{{$element.ID.Hex}}"
                                       data-created="{{$element.Created}}"
                                       data-modified="{{$element.Modified}}"
                                       data-name="{{$element.Artist.Name}}"
                                       data-realname="{{$element.Artist.RealName}}"
                                       data-born="{{$element.Artist.Born}}"
                                       data-died="{{$element.Artist.Died}}"
                                       data-birthplace="{{$element.Artist.Birthplace}}"
                                       data-deathplace="{{$element.Artist.Deathplace}}"
                                       data-nationality="{{$element.Artist.Nationality}}"
                                       data-biography="{{$element.Artist.Biography}}"
                                       data-painter="{{$element.Artist.IsPainter}}"
                                       data-sculptor="{{$element.Artist.IsSculptor}}"
                                       data-architect="{{$element.Artist.IsArchitect}}"
                                       data-printmaker="{{$element.Artist.IsPrintmaker}}">
                                 <span class="glyphicon glyphicon-eye-open"></span>
                            </a>
                       </span>            
                            {{if ne $role "guest"}}
                       &nbsp;&nbsp;
                       <span data-toggle="tooltip" data-placement="up" title="Modify details"> 
                            <a data-toggle="modal" data-target="#modifyArtistModal"
                                       data-id="{{$element.ID.Hex}}"  
                                       data-created="{{$element.Created}}" 
                                       data-modified="{{$element.Modified}}" 
                                       data-first="{{$element.Artist.Name.First}}"
                                       data-middle="{{$element.Artist.Name.Middle}}"
                                       data-last="{{$element.Artist.Name.Last}}"
                                       data-realfirst="{{$element.Artist.RealName.First}}"
                                       data-realmiddle="{{$element.Artist.RealName.Middle}}"
                                       data-reallast="{{$element.Artist.RealName.Last}}"
                                       data-born="{{$element.Artist.Born}}"
                                       data-died="{{$element.Artist.Died}}"
                                       data-birthplace="{{$element.Artist.Birthplace}}"
                                       data-deathplace="{{$element.Artist.Deathplace}}"
                                       data-nationality="{{$element.Artist.Nationality}}"
                                       data-biography="{{$element.Artist.Biography}}"
                                       data-painter="{{$element.Artist.IsPainter}}"
                                       data-sculptor="{{$element.Artist.IsSculptor}}"
                                       data-architect="{{$element.Artist.IsArchitect}}"
                                       data-printmaker="{{$element.Artist.IsPrintmaker}}">
                                 <span class="glyphicon glyphicon-edit"></span>
                            </a>
                       </span>            
                       &nbsp;&nbsp;
                       <span data-toggle="tooltip" data-placement="up" title="Remove"> 
                            <a data-toggle="modal" data-target="#removeArtistModal"
                                       data-id="{{$element.ID.Hex}}"  
                                       data-name="{{$element.Artist.Name}}">
                                <span class="glyphicon glyphicon-remove"></span>
                            </a>
                       </span>       
                            {{end}}
                    </td>
                  </tr>
                  {{end}}
                </tbody>
                </table>
                <ul class="pagination pagination-sm" id="cases-pagination">
                    <li><a href="#">&laquo;</a></li>
                    <li><a href="#">1</a></li>
                    <li><a href="#">2</a></li>
                    <li><a href="#">3</a></li>
                    <li><a href="#">4</a></li>
                    <li><a href="#">&raquo;</a></li>
                </ul>

    <!-- add modals -->
    {{template "view_artist_modal"}}
{{if ne $role "guest"}}
    {{template "modify_artist_modal"}}
    {{template "remove_artist_modal"}}
{{end}}
    <!-- end of modals definition -->   

            {{else}}
                <p>No artists found.</p>
            {{end}}

            </div>
        </div> <!-- row -->
    </div> <!-- container fluid -->

{{if ne $role "guest"}}
{{template "add_artist_modal"}}
{{end}}

{{template "insert-js"}}
    <script>

    $('#viewArtistModal').on('show.bs.modal', function (event) {

        var button = $(event.relatedTarget);     // Button that triggered the modal
        // Extract info from data-* requirement attribute
        //var id = button.data('id');  
 		var painter = button.data('painter');
 		var sculptor = button.data('sculptor');
 		var architect = button.data('architect');
 		var ceramic = button.data('ceramicist');
 		var print = button.data('printmaker');

        // Update the modal's content. We'll use jQuery here, but you could use a data 
        // binding library or other methods instead.
        var modal = $(this);
        modal.find('.modal-title').text(button.data('name'));
        //modal.find('.modal-body #hexid').val(id);
        modal.find('.modal-body #realnamev').text(button.data('realname'));
        modal.find('.modal-body #diedv').text(button.data('died'));
        modal.find('.modal-body #bornv').text(button.data('born'));
        modal.find('.modal-body #birthplacev').text(button.data('birthplace'));
        modal.find('.modal-body #deathplacev').text(button.data('deathplace'));
        modal.find('.modal-body #nationalityv').text(button.data('nationality'));
        modal.find('.modal-body #biographyv').text(button.data('biography'));
        modal.find('.modal-body #createdv').text(button.data('created'));
        modal.find('.modal-body #modifiedv').text(button.data('modified'));

        var roles = ""
        if (painter) { 
            roles += " Painter" 
        } 
        if (sculptor) { 
            roles += " Sculptor"
        } 
        if (print) { 
            roles += " Printmaker" 
        } 
        if (architect) { 
            roles += " Architect" 
        } 
        modal.find('.modal-body #rolesv').text(roles);
    })

{{if ne $role "guest"}}
    $('#modifyArtistModal').on('show.bs.modal', function (event) {

        var button = $(event.relatedTarget); // Button that triggered the modal
        // Extract info from data-* requirement attribute
        //var id = button.data('id');  
        var first = button.data('first');
        var middle = button.data('middle');
        var last = button.data('last');
 		var painter= button.data('painter');
 		var sculptor = button.data('sculptor');
 		var architect = button.data('architect');
 		var ceramic = button.data('ceramicist');
 		var print = button.data('printmaker');
        var created = button.data('created');

        // Update the modal's content. We'll use jQuery here, but you could use a data binding library 
        // or other methods instead.
        var modal = $(this);
        modal.find('.modal-title').text('Modify "' + first + ' ' + middle + ' ' + last + '" Details');
        modal.find('.modal-body #hexid').val(button.data('id'));
        modal.find('.modal-body #firstm').val(first);
        modal.find('.modal-body #middlem').val(middle);
        modal.find('.modal-body #lastm').val(last);
        modal.find('.modal-body #realfirstm').val(button.data('realfirst'));
        modal.find('.modal-body #realmiddlem').val(button.data('realmiddle'));
        modal.find('.modal-body #reallastm').val(button.data('reallast'));
        modal.find('.modal-body #diedm').val(button.data('died'));
        modal.find('.modal-body #bornm').val(button.data('born'));
        modal.find('.modal-body #birthplacem').val(button.data('birthplace'));
        modal.find('.modal-body #deathplacem').val(button.data('deathplace'));
        modal.find('.modal-body #nationalitym').val(button.data('nationality'));
        modal.find('.modal-body #biographym').val(button.data('biography'));
        if (painter) { 
            modal.find('.modal-body #painterm').prop('checked', true); 
        }
        if (sculptor) { 
            modal.find('.modal-body #sculptorm').prop('checked', true); 
        }
        if (architect) { 
            modal.find('.modal-body #architectm').prop('checked', true); 
        }
        if (ceramic) { 
            modal.find('.modal-body #ceramicistm').prop('checked', true); 
        }
        if (print) { 
            modal.find('.modal-body #printmakerm').prop('checked', true); 
        }
        modal.find('.modal-body #created').val(created);
        modal.find('.modal-body #createdm').text(created);
        modal.find('.modal-body #modifiedm').text(button.data('modified'));
    })

// Handle the removals using modal pop-up 
   $('#removeArtistModal').on('show.bs.modal', function(event) {
    
        var button = $(event.relatedTarget);
        var id = button.data('id');
        var name = button.data('name');

        // Update the modal's content. We'll use jQuery here, but you could use a data binding library 
        // or other methods instead.
        var modal = $(this);
        modal.find('.modal-body #removename').text(name);
        modal.find('.modal-body #hexid').val(id);
        modal.find('.modal-body #name').val(name);

        // Let's define the 'remove' button onclick() callback... 
        var url = '/artist/' + id + '/delete';
        $('#removebtn').on('click', function(e) { 
            postForm('remove_artist_form', url); 
            $('#removeArtistModal').modal('hide');
        });
   });

    // This should post a form to modify artist
    var modifyArtist = function(form_id, id) {
        var url = '/artist/' + id + '/put';
        postForm(form_id, url);
    }
{{end}}
    </script>
</body>
</html>
{{end}}

{{define "add_artist_modal"}}
<div class="modal fade" id="addArtistModal" tabindex="-1" role="dialog" aria-labelledby="addArtistModalLabel">
<div class="modal-dialog modal-lg">
<div class="modal-content">

    <div class="modal-header">
    <div class="container-fluid">
        <div class="row">
            <h3 class="modal-title col-sm-8" id="addArtistModalLabel">Add a New Artist</h3>
            <button type="button" class="btn btn-primary btn-sm col-sm-2" 
                    onclick="postForm('add-artist-form', '/artist'); $('#addArtistModal').modal('hide');"> Add </button>
            <button type="button" class="btn btn-default btn-sm col-sm-2" data-dismiss="modal">Cancel</button>
        </div> <!-- row -->
    </div> <!-- container-fluid -->
    </div> <!-- modal-header -->

    <div class="modal-body">
    <form class="form-horizontal" role="form" method="post" id="add-artist-form">
    <fieldset>

    <div class="form-group form-group-sm has-success">
        <div class="col-sm-1 control-label"> <strong>Name</strong> </div>
        <label for="first" class="col-sm-1 control-label">First</label>
        <div class="col-sm-3">
            <input type="text" class="form-control" id="first" name="first" placeholder="first name" required />
        </div>
        <label for="middle" class="col-sm-1 control-label">Middle</label>
        <div class="col-sm-2">
            <input type="text" class="form-control" id="middle" name="middle" placeholder="middle" />
        </div>
        <label for="last" class="col-sm-1 control-label">Last</label>
        <div class="col-sm-3">
            <input type="text" class="form-control" id="last" name="last" placeholder="last name" />
        </div>
    </div> <!-- form-group -->

    <div class="form-group form-group-sm">
        <div class="col-sm-1 control-label"><strong>Real Name</strong></div>
        <label for="realfirst" class="col-sm-1 control-label">First</label>
        <div class="col-sm-3">
            <input type="text" class="form-control" id="realfirst" name="realfirst" value="" placeholder="first name" />
        </div>
        <label for="realmiddle" class="col-sm-1 control-label">Middle</label>
        <div class="col-sm-2">
            <input type="text" class="form-control" id="realmiddle" name="realmiddle" value="" placeholder="middle" />
        </div>
        <label for="reallast" class="col-sm-1 control-label">Last</label>
        <div class="col-sm-3">
            <input type="text" class="form-control" id="reallast" name="reallast" value="" placeholder="last" />
        </div>
    </div> <!-- form-group -->

    <div class="form-group form-group-sm">
        <label for="born" class="col-sm-2 control-label">Born</label>
        <div class="col-sm-4">
            <input type="date" class="form-control" id="born" name="born" />
        </div>
        <div class="col-sm-6">
            <input type="text" class="form-control" id="birthplace" name="birthplace" placeholder="Birthplace..." />
        </div>
    </div> <!-- form-group -->
    <div class="form-group form-group-sm">
        <label for="died" class="col-sm-2 control-label">Died</label>
        <div class="col-sm-4">
            <input type="date" class="form-control" id="died" name="died" value="" />
        </div>
        <div class="col-sm-6">
            <input type="text" class="form-control" id="deathplace" name="deathplace" placeholder="Place of death..." />
        </div>
    </div> <!-- form-group -->

    <div class="form-group form-group-sm">
        <label for="nationality" class="col-sm-2 control-label">Nationality</label>
        <div class="col-sm-10">
            <input type="text" class="form-control" id="nationality" name="nationality" value="" placeholder="nationality" />
        </div>
    </div> <!-- form-group -->

    <div class="form-group form-group-sm">
        <label for="biography" class="col-sm-2 control-label"> Biography </label>
    </div> <!-- form-group -->
    <div class="form-group form-group-sm">
        <textarea class="col-sm-10 form-control" id="biography" name="biography" rows="6">Biography goes here... </textarea>
    </div> <!-- form-group -->

    <div class="form-group form-group-sm form-inline">
        <div class="checkbox-inline col-sm-2">
            <label class="control-label pull-right"><input type="checkbox" name="painter" value="yes"> Painter </label>
        </div> <!-- checkbox -->
        <div class="checkbox-inline col-sm-2"> 
            <label class="control-label pull-right"><input type="checkbox" name="sculptor" value="yes"> Sculptor </label>
        </div> <!-- checkbox -->
        <div class="checkbox-inline col-sm-2">
            <label class="control-label pull-right"><input type="checkbox" name="printmaker" value="yes"> Printmaker </label>
        </div> <!-- checkbox -->
        <div class="checkbox-inline col-sm-2">
            <label class="control-label pull-right"><input type="checkbox" name="architect" value="yes"> Architect </label>
        </div> <!-- checkbox -->
    </div> <!-- form-group -->

<!--
    <div class="form-group form-group-sm">
        <div class="checkbox-inline col-sm-2">
            <label class="control-label pull-right"><input type="checkbox" name="writer" value="yes"> Writer </label>
        </div>
        <div class="checkbox-inline col-sm-2">
            <label class="control-label pull-right"><input type="checkbox" name="poet" value="yes"> Poet </label>
        </div> 
    </div> 
-->

    <!-- TODO -->

       </fieldset>
    </form>

    </div> <!-- modal-body -->
</div>
</div>
</div>
{{end}}

{{define "view_artist_modal"}}
<div class="modal fade" id="viewArtistModal" tabindex="-1" role="dialog" aria-lebeleledby="ViewArtistModalLabel">
<div class="modal-dialog modal-lg">
<div class="modal-content">

    <div class="modal-header">
    <div class="container-fluid">
        <div class="row">
            <h3 class="modal-title col-md-10" id="viewArtistModalLabel"><span id="namev"></span></h3>
            <button type="button" class="btn btn-default btn-sm col-md-2" data-dismiss="modal">Close</button>
        </div> <!-- row -->
    </div> <!-- container-fluid -->
    </div> <!-- modal-header -->

    <div class="modal-body">
        <div class="container-fluid" id="view-artist-table-div">

        <div class="row">
            <strong><span id="rolesv"></span></strong>
        </div> <!-- row -->

        <div class="row">
        &nbsp; <!-- empty row -->
        </div> <!-- row -->

        <div class="row">
             <table id="view-artist-table" class="table table-hover small">
             <tbody>
                 <tr> <td class="col-sm-3"><strong>Real Name</strong></td>
                      <td class="col-sm-9"><span id="realnamev"></span></td> </tr>
                 <tr> <td class="col-sm-3"><strong>Born</strong></td> 
                      <td class="col-sm-9"><span id="bornv"></span>&emsp;
                                           <span id="birthplacev"></span></td> </tr>
                 <tr> <td class="col-sm-3"><strong>Died</strong></td> 
                      <td class="col-sm-9"><span id="diedv"></span>&emsp;
                                           <span id="deathplacev"></span></td> </tr>
                 <tr> <td class="col-sm-3"><strong>Nationality</strong></td> 
                      <td id="nationalityv" class="col-sm-9"></td> </tr>
                 <tr> <td class="col-sm-3"><strong>Biography</strong></td> 
                      <td id="biographyv" class="col-sm-9"></td> </tr>
                 <tr> <td class="col-sm-3"><strong>Created</strong></td> 
                      <td id="createdv" class="col-sm-9"></td> </tr>
                 <tr> <td class="col-sm-3"><strong>Last Modified</strong></td> 
                      <td id="modifiedv" class="col-sm-9"></td> </tr>
             </tbody>
             </table>
        </div> <!-- row -->
        </div> <!-- container-fluid -->
    </div> <!-- modal-body -->

</div>
</div>
</div>
{{end}}

{{define "modify_artist_modal"}}
<div class="modal fade" id="modifyArtistModal" tabindex="-1" role="dialog" aria-lebeleledby="modifyArtistModalLabel">
<div class="modal-dialog modal-lg">
<div class="modal-content">

    <div class="modal-header">
    <div class="container-fluid">
        <div class="row">
            <h3 class="modal-title col-md-8" id="modifyArtistModalLabel">Empty Artist Details</h3>
             <button type="button" class="btn btn-primary btn-sm col-sm-2" 
                     onclick="modifyArtist('modify_artist_form', $('#hexid').val()); $('#modifyArtistModal').modal('hide');">
                     Modify
             </button>
            <button type="button" class="btn btn-default btn-sm col-md-2" data-dismiss="modal">Cancel</button>
        </div> <!-- row -->
    </div> <!-- container-fluid -->
    </div> <!-- modal-header -->

    <div class="modal-body">
    <div class="container-fluid">
        <div class="row">

        <form id="modify_artist_form" class="form-horizontal">
            <fieldset>
                <input type="hidden" id="hexid" name="hexid" />
                <div class="form-group form-group-sm">
                    <div class="col-sm-1 control-label"> <strong>Name</strong> </div>
                    <label for="first" class="col-sm-1 control-label">First</label>
                    <div class="col-sm-3">
                        <input type="text" class="form-control" id="firstm" name="first" required />
                    </div>
                    <label for="middle" class="col-sm-1 control-label">Middle</label>
                    <div class="col-sm-2">
                        <input type="text" class="form-control" id="middlem" name="middle" />
                    </div>
                    <label for="last" class="col-sm-1 control-label">Last</label>
                    <div class="col-sm-3">
                        <input type="text" class="form-control" id="lastm" name="last" />
                    </div>
                </div> <!-- form-group -->

                <div class="form-group form-group-sm">
                    <div class="col-sm-1 control-label"><strong>Real Name</strong></div>
                    <label for="realfirst" class="col-sm-1 control-label">First</label>
                    <div class="col-sm-3">
                        <input type="text" class="form-control" id="realfirstm" name="realfirst" />
                    </div>
                    <label for="realmiddle" class="col-sm-1 control-label">Middle</label>
                    <div class="col-sm-2">
                        <input type="text" class="form-control" id="realmiddlem" name="realmiddle" />
                    </div>
                    <label for="reallast" class="col-sm-1 control-label">Last</label>
                    <div class="col-sm-3">
                        <input type="text" class="form-control" id="reallastm" name="reallast" />
                    </div>
                </div> <!-- form-group -->

                <div class="form-group form-group-sm">
                    <label for="born" class="col-sm-2 control-label">Born</label>
                    <div class="col-sm-4">
                        <input type="date" class="form-control" id="bornm" name="born" />
                    </div>
                    <div class="col-sm-6">
                        <input type="text" class="form-control" id="birthplacem" name="birthplace" />
                    </div>
                </div> <!-- form-group -->
                <div class="form-group form-group-sm">
                    <label for="died" class="col-sm-2 control-label">Died</label>
                    <div class="col-sm-4">
                        <input type="date" class="form-control" id="diedm" name="died" />
                    </div>
                    <div class="col-sm-6">
                        <input type="text" class="form-control" id="deathplacem" name="deathplace" />
                    </div>
                </div> <!-- form-group -->

                <div class="form-group form-group-sm">
                    <label for="nationality" class="col-sm-2 control-label">Nationality</label>
                    <div class="col-sm-10">
                        <input type="text" class="form-control" id="nationalitym" name="nationality" />
                    </div>
                </div> <!-- form-group -->

                <div class="form-group form-group-sm">
                    <label for="biography" class="col-sm-2 control-label"> Biography </label>
                </div> <!-- form-group -->
                <div class="form-group form-group-sm">
                    <textarea class="col-sm-10 form-control" id="biographym" name="biography" rows="6"></textarea>
                </div> <!-- form-group -->

                <div class="form-group form-group-sm form-inline">
                    <div class="checkbox-inline col-sm-2">
                        <label class="control-label pull-right">
                            <input type="checkbox" id="painterm" name="painter" value="yes"/> Painter 
                        </label>
                    </div> <!-- checkbox -->
                    <div class="checkbox-inline col-sm-2"> 
                        <label class="control-label pull-right">
                            <input type="checkbox" id="sculptorm" name="sculptor" value="yes"/> Sculptor 
                        </label>
                    </div> <!-- checkbox -->
                    <div class="checkbox-inline col-sm-2">
                        <label class="control-label pull-right">
                            <input type="checkbox" id="printmakerm" name="printmaker" value="yes"/> Printmaker 
                        </label>
                    </div> <!-- checkbox -->
                    <div class="checkbox-inline col-sm-2">
                        <label class="control-label pull-right">
                            <input type="checkbox" id="architectm" name="architect" value="yes" /> Architect 
                        </label>
                    </div> <!-- checkbox -->
                </div> <!-- form-group -->

<!--
    <div class="form-group form-group-sm">
        <div class="checkbox-inline col-sm-2">
            <label class="control-label pull-right"><input type="checkbox" name="writer"> Writer </label>
        </div>
        <div class="checkbox-inline col-sm-2">
            <label class="control-label pull-right"><input type="checkbox" name="poet"> Poet </label>
        </div> 
    </div> 
-->

    <!-- TODO -->
        <div class="form-group form-group-sm small">
            <input type="hidden" id="created" name="created"></input>
            <div class="col-sm-2 text-right"><strong>Created:</strong></div>
            <div id="createdm" name="createdm" class="col-sm-4 text-left">Error</div>
            <div class="col-sm-2 text-right"><strong>Modified:</strong></div>
            <div id="modifiedm" name="modifiedm" class="col-sm-4 text-left">Error</div>
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

{{define "remove_artist_modal"}}
{{template "remove-modal" "Artist"}}
{{end}}
