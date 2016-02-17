{{define "techniques"}}
{{$role := .User.Role}}
{{$name := totitle .Ptype}}
<!DOCTYPE html>
<html lang="en">
<head>
{{template "htmlhead" "Handle Techniques"}}
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


    {{if .Techniques}}
            <table class="table table-striped table-hover small" id="techniques-list-table">

                <thead>
                    <tr>
                        <th class="col-sm-1">#</th>
                        <th class="col-sm-2">Technique</th>
                        <th class="col-sm-8">Description</th>
                        <th class="col-sm-1 text-right">Actions</th>
                    </tr>
                </thead>

                <tfoot>
                    <tr class="bg-primary">
                    <td colspan="4"> 
                        <strong>{{.Num}} {{if eq .Num 1}} technique {{else}} techniques {{end}} found.</strong> 
                    </td>
                    </tr>
                </tfoot>

                <tbody>
                    {{range $index, $element := .Techniques}}
                    {{$cnt := add $index 1}}
                    <tr id="technique-row-{{$cnt}}">
                        <td>{{$cnt}}</td>
                        <td>{{$element.Name}}</td>
                        <td>{{$element.Description}}</td>
                        <td class="text-right">
                            <span data-toggle="tooltip" data-placement="up" title="View details">
                            <a data-toggle="modal" data-target="#viewTechniqueModal"
                                       data-id="{{$element.ID.Hex}}"
                                       data-created="{{$element.Created}}"
                                       data-modified="{{$element.Modified}}"
                                       data-name="{{$element.Technique.Name}}"
                                       data-desc="{{$element.Technique.Description}}">
                                 <span class="glyphicon glyphicon-eye-open"></span>
                            </a>
                            </span>            
                        {{if ne $role "guest"}}
                            &nbsp;&nbsp;
                            <span data-toggle="tooltip" data-placement="up" title="Modify details"> 
                            <a data-toggle="modal" data-target="#modifyTechniqueModal"
                                       data-id="{{$element.ID.Hex}}"  
                                       data-created="{{$element.Created}}" 
                                       data-modified="{{$element.Modified}}" 
                                       data-name="{{$element.Technique.Name}}"
                                       data-desc="{{$element.Technique.Description}}">
                                <span class="glyphicon glyphicon-edit"></span>
                            </a>
                            </span>       
                            &nbsp;&nbsp;
                            <span data-toggle="tooltip" data-placement="up" title="Remove"> 
                            <a data-toggle="modal" data-target="#removeTechniqueModal"
                                       data-id="{{$element.ID.Hex}}"  
                                       data-name="{{$element.Technique.Name}}">
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
    {{template "view_technique_modal"}}
{{if ne $role "guest"}}
    {{template "modify_technique_modal"}}
    {{template "remove_technique_modal" .}}
{{end}}
    <!-- end of modals definition -->                

    {{else}}
    <p>No techniques found.</p>
    {{end}}

        </div>
     </div> <!-- row -->
    </div> <!-- container fluid -->
{{if ne $role "guest"}}
    {{template "add_technique_modal"}}
{{end}}

{{template "insert-js"}}
    <script>

    $('#viewTechniqueModal').on('show.bs.modal', function (event) {

        var button = $(event.relatedTarget);     // Button that triggered the modal

        // Extract info from data-* requirement attribute
        var id = button.data('id');  
        var name = button.data('name');
		var desc = button.data('desc');
        var created = button.data('created');
        var modified = button.data('modified');
        // Update the modal's content. We'll use jQuery here, but you could use a data 
        // binding library or other methods instead.
        var modal = $(this)
        modal.find('.modal-title').text('The "' + name + '" Technique Details');
        modal.find('.modal-body #hexid').val(id);
        modal.find('.modal-body #namev').val(name);
        modal.find('.modal-body #descriptionv').val(desc);
        modal.find('.modal-body #createdv').text(created);
        modal.find('.modal-body #modifiedv').text(modified);
    })

{{if ne $role "guest"}}
    $('#modifyTechniqueModal').on('show.bs.modal', function (event) {

        var button = $(event.relatedTarget); // Button that triggered the modal

        // Extract info from data-* requirement attribute
        var id = button.data('id');  
        var name = button.data('name');
		var desc = button.data('desc');
        var created = button.data('created');
        var modified = button.data('modified');
        // Update the modal's content. We'll use jQuery here, but you could use a data binding library 
        // or other methods instead.
        var modal = $(this)
        modal.find('.modal-title').text('Modify "' + name + '" Technique Details');
        modal.find('.modal-body #hexid').val(id);
        modal.find('.modal-body #name').val(name);
        modal.find('.modal-body #description').val(desc);
        modal.find('.modal-body #created').val(created);
        modal.find('.modal-body #createdd').text(created);
        modal.find('.modal-body #modifiedd').text(modified);
    })

	// Handle the removals using modal pop-up 
   	$('#removeTechniqueModal').on('show.bs.modal', function(event) {
    
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
        var url = '/technique/' + id + '/delete';
        $('#removebtn').on('click', function(e) { 
            postForm('remove_technique_form', url); 
            $('#removeTechniqueModal').modal('hide');
        });
   	});

    // This should post  from to modify a technique
    var modifyTechnique = function(form_id, id) {
        var url = '/technique/' + id + '/put';
        postForm(form_id, url);
    }
{{end}}
	</script>

  </body>
</html>
{{end}}

{{define "add_technique_modal"}}
<div class="modal fade" id="addTechniqueModal" tabindex="-1" role="dialog" aria-labelledby="addTechniqueModalLabel">
<div class="modal-dialog">
<div class="modal-content">

    <div class="modal-header">
    <div class="container-fluid">
        <div class="row">
            <h4 class="modal-title col-sm-8" id="addTechniqueModalLabel">Add a New Technique</h4>
            <button type="button" class="btn btn-primary btn-sm col-sm-2" 
                    onclick="postForm('add_technique_form', '/technique'); $('#addTechnniqueModal').modal('hide');"> Add
			</button>
            <button type="button" class="btn btn-default btn-sm col-sm-2" data-dismiss="modal">Cancel</button>
        </div> <!-- row -->
    </div> <!-- container-fluid -->
    </div> <!-- modal-header -->

    <div class="modal-body">
      <form id="add_technique_form" class="form-horizontal" method="post">
            <div class="form-group form-group-sm">
                <label for="name" class="col-sm-2 control-label">Name</label>
                <div class="col-sm-10">
                  <input type="text" class="form-control" id="name" name="name" placeholder="Technique Name" required> 
                </div>
            </div>
            <div class="form-group form-group-sm">
                <label for="description" class="col-sm-2 control-label">Description</label>
                <div class="col-sm-offset-10">&nbsp;</div>
                <div class="col-sm-12">
                <textarea class="form-control" rows="15" id="description" name="description"></textarea>
                </div>
            </div>
      </form>
    </div>
</div>
</div>
</div>
{{end}}

{{define "view_technique_modal"}}
<div class="modal fade" id="viewTechniqueModal" tabindex="-1" role="dialog" aria-lebeleledby="ViewTechniqueModalLabel">
<div class="modal-dialog">
<div class="modal-content">

    <div class="modal-header">
    <div class="container-fluid">
        <div class="row">
            <h4 class="modal-title col-md-10" id="viewTechniqueModalLabel">Empty Dating Details</h4>
            <button type="button" class="btn btn-default btn-sm col-md-2" data-dismiss="modal">Close</button>
        </div> <!-- row -->
    </div> <!-- container-fluid -->
    </div> <!-- modal-header -->

    <div class="modal-body">
    <div class="container-fluid">
        <div class="row">

    <form id="view_dating_form" class="form-horizontal">
         <!--  <input type="hidden" id="hexid" name="hexid"></input> -->
        <div class="form-group form-group-sm">
            <label for="namev" class="col-sm-3 control-label">Name</label>
            <div class="col-sm-9">
                <input type="text" class="form-control" id="namev" name="namev" readonly></input>
            </div>
        </div>
        <div class="form-group form-group-sm">
            <label for="descriptionv" class="col-sm-3 control-label">Description</label>
            <div class="col-sm-offset-9"></div>
            <div class="col-sm-12">
                <textarea class="form-control" rows="15" id="descriptionv" name="descriptionv" readonly></textarea>
            </div>
        </div>
        <!--
        <div class="form-group form-group-sm">
            <label for="hexid" class="col-sm-3 control-label">Hex ID</label>
            <div class="col-sm-9">
                <input type="text" class="form-control" id="hexid" name="hexid" readonly></input>
            </div>
        </div>
        -->
        <div class="form-group form-group-sm small">
            <div class="col-sm-2 text-right"><strong>Created:</strong></div>
            <div id="createdv" class="col-sm-4 text-left">Error</div>
            <div class="col-sm-2 text-right"><strong>Modified:</strong></div>
            <div id="modifiedv" class="col-sm-4 text-left">Error</div>
        </div>
    </form>

        </div>
    </div>
    </div> <!-- modal-body -->

</div>
</div>
</div>
{{end}}

{{define "modify_technique_modal"}}
<div class="modal fade" id="modifyTechniqueModal" tabindex="-1" role="dialog" aria-lebeleledby="modifyTechniqueModalLabel">
<div class="modal-dialog">
<div class="modal-content">

    <div class="modal-header">
    <div class="container-fluid">
        <div class="row">
            <h4 class="modal-title col-md-8" id="modifyTechniqueModalLabel">Empty Technique Details</h4>
             <button type="button" class="btn btn-primary btn-sm col-sm-2" 
                     onclick="modifyTechnique('modify_technique_form',$('#hexid').val());$('#modifyTechniqueModal').modal('hide');">
                     Modify
             </button>
            <button type="button" class="btn btn-default btn-sm col-md-2" data-dismiss="modal">Cancel</button>
        </div> <!-- row -->
    </div> <!-- container-fluid -->
    </div> <!-- modal-header -->

    <div class="modal-body">
    <div class="container-fluid">
        <div class="row">
        <form id="modify_technique_form" class="form-horizontal">
            <input type="hidden" id="hexid" name="hexid" />
        <div class="form-group form-group-sm">
            <label for="name" class="col-sm-3 control-label">Name</label>
            <div class="col-sm-9">
                <input type="text" class="form-control" id="name" name="name" required></input>
            </div>
        </div>
        <div class="form-group form-group-sm">
            <label for="description" class="col-sm-3 control-label">Description</label>
            <div class="col-sm-offset-9"></div>
            <div class="col-sm-12">
                <textarea class="form-control" rows="15" id="description" name="description"></textarea>
            </div>
        </div>
   <div class="form-group form-group-sm small">
            <input type="hidden" id="created" name="created" />
            <div class="col-sm-2 text-right"><strong>Created:</strong></div>
            <div id="createdd" name="createdd" class="col-sm-4 text-left">Error</div>
            <div class="col-sm-2 text-right"><strong>Modified:</strong></div>
            <div id="modifiedd" name="modifiedd" class="col-sm-4 text-left">Error</div>
        </div>

        </div> <!-- row -->
    </div> <!-- container-fluid -->
    </form>
    </div> <!-- modal-body -->

</div>
</div>
</div>
{{end}}

{{define "remove_technique_modal"}}
{{template "remove-modal" "Technique"}}
{{end}}
