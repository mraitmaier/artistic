{{define "styles"}}
{{$role := .User.Role}}
<!DOCTYPE html>
<html lang="en">
<head>
{{template "htmlhead" "Handle Styles"}}
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
            <h1 id="data-list-header">Styles</h1>
        {{if ne $role "guest"}}
            <div id="new-req-btn">
            <button type="button" class="btn btn-primary btn-sm" data-toggle="modal" data-target="#addStyleModal">
                <span class="glyphicon glyphicon-plus"></span> &nbsp; Create a New Style
            </button>
        	</div>
        {{end}}
            <br />

		{{if .Styles}}
			<table class="table table-striped table-hover small" id="style-list-table">

				<thead>
					<tr>
						<th class="col-sm-1">#</th>
						<th class="col-sm-2">Style</th>
						<th class="col-sm-8">Description</th>
						<th class="col-sm-1 text-right">Actions</th>
					</tr>
				</thead>

               <tfoot>
                    <tr class="bg-primary">
                    <td colspan="4"> 
                        <strong>{{.Num}} {{if eq .Num 1}} style {{else}} styles {{end}} found.</strong> 
                    </td>
                    </tr>
                </tfoot>

				<tbody>
					{{range $index, $element := .Styles}}
					{{ $cnt := add $index 1 }}
					<tr class="art-single-row" id="style-row-{{$cnt}}">
						<td>{{$cnt}}</td>
						<td>{{printf "%s" $element.Name}}</td>
						<td>{{printf "%s" $element.Description}}</td>
						<td class="text-right">
  							<span data-toggle="tooltip" data-placement="up" title="View details">
                            <a data-toggle="modal" data-target="#viewStyleModal"
                                       data-id="{{$element.ID.Hex}}"
                                       data-created="{{$element.Created}}"
                                       data-modified="{{$element.Modified}}"
                                       data-name="{{$element.Style.Name}}"
                                       data-desc="{{$element.Style.Description}}">
                                 <span class="glyphicon glyphicon-eye-open"></span>
                            </a>
                            </span>            
                        {{if ne $role "guest"}}
                            &nbsp;&nbsp;
                            <span data-toggle="tooltip" data-placement="up" title="Modify details"> 
                            <a data-toggle="modal" data-target="#modifyStyleModal"
                                       data-id="{{$element.ID.Hex}}"  
                                       data-created="{{$element.Created}}" 
                                       data-modified="{{$element.Modified}}" 
                                       data-name="{{$element.Style.Name}}"
                                       data-desc="{{$element.Style.Description}}">
                                <span class="glyphicon glyphicon-edit"></span>
                            </a>
                            &nbsp;&nbsp;
                            <span data-toggle="tooltip" data-placement="up" title="Remove"> 
                            <a data-toggle="modal" data-target="#removeStyleModal"
                                       data-id="{{$element.ID.Hex}}"  
                                       data-name="{{$element.Style.Name}}">
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
    {{template "view_style_modal" .Styles}}
{{if ne $role "guest"}}
    {{template "modify_style_modal" .Styles}}
    {{template "remove_style_modal" .Styles}}
{{end}}
    <!-- end of modals definition -->          

		{{else}}
				<p>No styles found.</p>
		{{end}}

        </div>
     </div> <!-- row -->
    </div> <!-- container fluid -->

{{if ne $role "guest"}}
   {{template "add_style_modal"}}
{{end}}
    <!-- end of modals definition -->          

{{template "insert-js"}}
    <script>

    $('#viewStyleModal').on('show.bs.modal', function (event) {

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
        modal.find('.modal-title').text('The "' + name + '" Style Details');
        modal.find('.modal-body #hexid').val(id);
        modal.find('.modal-body #namev').val(name);
        modal.find('.modal-body #descriptionv').val(desc);
        modal.find('.modal-body #createdv').text(created);
        modal.find('.modal-body #modifiedv').text(modified);
    })

{{if ne $role "guest"}}
    $('#modifyStyleModal').on('show.bs.modal', function (event) {

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
        modal.find('.modal-title').text('Modify "' + name + '" Style Details');
        modal.find('.modal-body #hexid').val(id);
        modal.find('.modal-body #name').val(name);
        modal.find('.modal-body #description').val(desc);
        modal.find('.modal-body #created').val(created);
        modal.find('.modal-body #createdd').text(created);
        modal.find('.modal-body #modifiedd').text(modified);
    })

    // Handle the removals using modal pop-up 
   $('#removeStyleModal').on('show.bs.modal', function(event) {
    
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
        var url = '/style/' + id + '/delete';
        $('#removebtn').on('click', function(e) { 
            postForm('remove_style_form', url); 
            $('#removeStyleModal').modal('hide');
        });
   });

    // This should post a form to modify style
    var modifyStyle = function(form_id, id) {
        var url = '/style/' + id + '/put';
        postForm(form_id, url);
    }
{{end}}
 	</script>
  </body>
</html>
{{end}}

{{define "add_style_modal"}}
<div class="modal fade" id="addStyleModal" tabindex="-1" role="dialog" aria-labelledby="addStyleModalLabel">
<div class="modal-dialog">
<div class="modal-content">

    <div class="modal-header">
    <div class="container-fluid">
        <div class="row">
            <h4 class="modal-title col-sm-8" id="addStyleModalLabel">Add a New Style</h4>
            <button type="button" class="btn btn-primary btn-sm col-sm-2" 
                    onclick="postForm('add_style_form', '/style'); $('#addStyleModal').modal('hide');"> Add </button>
            <button type="button" class="btn btn-default btn-sm col-sm-2" data-dismiss="modal">Cancel</button>
        </div> <!-- row -->
    </div> <!-- container-fluid -->
    </div> <!-- modal-header -->

    <div class="modal-body">
      <form id="add_style_form" class="form-horizontal" method="post">
            <div class="form-group form-group-sm">
                <label for="name" class="col-sm-2 control-label">Name</label>
                <div class="col-sm-10">
                  <input type="text" class="form-control" id="name" name="name" placeholder="Style Name" required> 
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

{{define "view_style_modal"}}
<div class="modal fade" id="viewStyleModal" tabindex="-1" role="dialog" aria-lebeleledby="ViewStyleModalLabel">
<div class="modal-dialog">
<div class="modal-content">

    <div class="modal-header">
    <div class="container-fluid">
        <div class="row">
            <h4 class="modal-title col-md-10" id="viewStyleModalLabel">Empty Style Details</h4>
            <button type="button" class="btn btn-default btn-sm col-md-2" data-dismiss="modal">Close</button>
        </div> <!-- row -->
    </div> <!-- container-fluid -->
    </div> <!-- modal-header -->

    <div class="modal-body">
    <div class="container-fluid">
        <div class="row">

			<form id="view_style_form" class="form-horizontal">
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

{{define "modify_style_modal"}}
<div class="modal fade" id="modifyStyleModal" tabindex="-1" role="dialog" aria-lebeleledby="modifyStyleModalLabel">
<div class="modal-dialog">
<div class="modal-content">

    <div class="modal-header">
    <div class="container-fluid">
        <div class="row">
            <h4 class="modal-title col-md-8" id="modifyStyleModalLabel">Empty Style Details</h4>
             <button type="button" class="btn btn-primary btn-sm col-sm-2" 
                     onclick="modifyStyle('modify_style_form',$('#hexid').val());$('#modifyStyleModal').modal('hide');">
                     Modify
             </button>
            <button type="button" class="btn btn-default btn-sm col-md-2" data-dismiss="modal">Cancel</button>
        </div> <!-- row -->
    </div> <!-- container-fluid -->
    </div> <!-- modal-header -->

    <div class="modal-body">
    <div class="container-fluid">
        <div class="row">
        <form id="modify_style_form" class="form-horizontal">
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

{{define "remove_style_modal"}}
<div class="modal fade" id="removeStyleModal" tabindex="-1" role="dialog" aria-labelledby="removeStyleModalLabel">
<div class="modal-dialog">
<div class="modal-content">

    <div class="modal-header">
    <div class="container-fluid">
        <div class="row">
            <h4 class="modal-title col-sm-8" id="removeStyleModalLabel">Remove Style</h4>
            <button type="button" class="btn btn-primary btn-sm col-sm-2" id="removebtn"> Remove </button>
            <button type="button" class="btn btn-default btn-sm col-sm-2" data-dismiss="modal"> Cancel </button>
        </div> <!-- row -->
    </div> <!-- container-fluid -->
    </div> <!-- modal-header -->

    <div class="modal-body">
    <p> Would you really like to remove the style '<span id="removename"></span>'?</p>
    <form method="post" id="remove_style_form">
        <input type="hidden" name="id" id="id" />
        <input type="hidden" name="name" id="name" />
    </form>
    </div>
</div>
</div>
</div>
{{end}}

