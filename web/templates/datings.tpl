{{define "datings"}}
{{$role := .User.Role}}
<!DOCTYPE html>
<html lang="en">
<head>
{{template "htmlhead" "Handle Datings"}}
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
                <h1 id="data-list-header">Datings</h1>
                <!-- datings table is created dynamically by JS -->
                <div id="datings-table-div">

               {{if .Datings}}

                <table class="table table-striped table-hover small" id="dating-list-table">
                <thead>
                    <tr> <th class="col-sm-1">#</th> <th class="col-sm-1">Dating</th> 
                         <th class="col-sm-9">Description</th> <th class="col-sm-1 text-right">Actions</th> </tr>
                </thead>
<!--
                <tfoot>
                    <tr class="bg-primary">
                        <td colspan="4">
                             <strong>8 datings found.</strong> <!-- this is hardcoded, because it doesn't change. -->
                        </td>
                    </tr>
                </tfoot>
-->
                <tbody>
                    {{range $index, $element := .Datings}}
                    {{$num := add $index 1}}
                    <tr id="dating-row-{{$num}}">
                        <td>{{$num}}</td>
                        <td>{{$element.Dating.Dating}}</td>
                        <td>{{$element.Dating.Description}}</td>
                        <td class="text-right">
                            <span data-toggle="tooltip" data-placement="up" title="View details">
                            <a href="" data-toggle="modal" data-target="#viewDatingModal"
                                        data-id="{{$element.ID.Hex}}"
                                        data-created="{{$element.Created}}"
                                        data-modified="{{$element.Modified}}"
                                        data-dating="{{$element.Dating.Dating}}"
                                        data-desc="{{$element.Dating.Description}}">
                                <span class="glyphicon glyphicon-eye-open"></span>
                            </a>
                            </span>
                        {{if ne $role "guest"}}
                            &nbsp;&nbsp;
                            <span data-toggle="tooltip" data-placement="up" title="Modify details">
                            <a href="" data-toggle="modal" data-target="#modifyDatingModal"
                                        data-id="{{$element.ID.Hex}}"
                                        data-created="{{$element.Created}}"
                                        data-modified="{{$element.Modified}}"
                                        data-dating="{{$element.Dating.Dating}}"
                                        data-desc="{{$element.Dating.Description}}">
                                <span class="glyphicon glyphicon-edit"></span>
                            </a>
                            </span>
                        {{end}}
                        </td>
                    </tr>
                    {{end}}
                </tbody>

                </table>


    <!-- add modals -->
    {{template "view_dating_modal" .Datings}}
{{if ne $role "guest"}}
    {{template "modify_dating_modal" .Datings}}
{{end}}
    <!-- end of modals definition -->

                {{else}}
                    <p>No datings found.</p>
                {{end}}
           
            </div>
        </div> <!-- row -->
    </div> <!-- container fluid -->

{{template "insert-js"}}
    <script>
    // when page is ready...

        $('#viewDatingModal').on('show.bs.modal', function (event) {

            var btn = $(event.relatedTarget); // button that triggerd event
            // extract info from data-dating attribute
            //var hexid = btn.data('id');
            var dating = btn.data('dating');
            // Update the modal's content.
            var modal = $(this);
            modal.find('.modal-title').text('The "' + dating + '" Dating Details');
            modal.find('.modal-body #datingv').val(dating);
            modal.find('.modal-body #descriptionv').val(btn.data('desc'));
            modal.find('.modal-body #createdv').text(btn.data('created'));
            modal.find('.modal-body #modifiedv').text(btn.data('modified'));
        });

{{if ne $role "guest"}}
        $('#modifyDatingModal').on('show.bs.modal', function (event) {

            var btn = $(event.relatedTarget); // button that triggerd event
            // extract info from data-dating attribute
            // var hexid = btn.data('id');
            var dating = btn.data('dating');
            var created = btn.data('created');
            // Update the modal's content.
            var modal = $(this);
            modal.find('.modal-title').text('Modify Dating "' + dating + '"');
            modal.find('.modal-body #hexid').val(btn.data('hexid'));
            modal.find('.modal-body #dating').val(dating);
            modal.find('.modal-body #description').val(btn.data('desc'));
            modal.find('.modal-body #created').val(created);   // hidden val
            modal.find('.modal-body #createdm').text(created); // only display
            modal.find('.modal-body #modifiedm').text(btn.data('modified')); //only display
        });

	// This should post form (PUT method) to modify a dating
	var modifyDating = function(form_id, id) {
    	var url = '/dating/' + id + '/put';
        //alert("ID=" + id); //DEBUG
    	postForm(form_id, url);
	}
{{end}}
    </script>

</body>
</html>
{{end}}

{{define "dating-list"}}
    {{if .}}
    <table class="table table-striped table-hover small" id="dating-list-table">
    <thead>
        <tr> <th class="col-sm-1">#</th> <th class="col-sm-1">Dating</th> 
             <th class="col-sm-8">Description</th> <th class="col-sm-2">Actions</th> </tr>
    </thead>

    <tfoot>
        <tr class="bg-primary">
            <td colspan="8" class="text-right"> <!-- this is hardcoded, because it doesn't change. -->
                 <strong>8 datings found.</strong> <!-- this is hardcoded, because it doesn't change. -->
            </td>
        </tr>
    </tfoot>

    <tbody>
        {{range $index, $element := .}}
        {{$id := add $index 1}}
        <tr id="dating-row-{{$id}}">
            <td>{{$id}}</td>
            <td>{{$element.Dating}}</td>
            <td>{{$element.Description}}</td>
            <td class="text-right">
                <span data-toggle="tooltip" data-placement="up" title="View details">
                <a href="" data-toggle="modal" data-target="#viewDatingModal"
                            data-id="{{$element.ID.Hex}}"
                            data-created="{{$element.Created}}"
                            data-modified="{{$element.Modified}}"
                            data-dating="{{$element.Dating}}"
                            data-desc="{{$element.Description}}">
                    <span class="glyphicon glyphicon-eye-open"></span>
                </a>
                </span>
                &nbsp;&nbsp;
                <span data-toggle="tooltip" data-placement="up" title="Modify details">
                <a href="" data-toggle="modal" data-target="#modifyDatingModal"
                            data-id="{{$element.ID.Hex}}"
                            data-created="{{$element.Created}}"
                            data-modified="{{$element.Modified}}"
                            data-dating="{{$element.Dating}}"
                            data-desc="{{$element.Description}}">
                    <span class="glyphicon glyphicon-edit"></span>
                </a>
                </span>

            </td>
        </tr>
        {{end}}
    </tbody>

    </table>
    {{else}}
        <p>There are no datings defined yet.</p>
    {{end}}

{{end}}

{{define "view_dating_modal"}}
<div class="modal fade" id="viewDatingModal" tabindex="-1" role="dialog" aria-lebeleledby="ViewDatingModalLabel">
<div class="modal-dialog">
<div class="modal-content">

    <div class="modal-header">
    <div class="container-fluid">
        <div class="row">
            <h4 class="modal-title col-md-10" id="viewDatingModalLabel">Empty Dating Details</h4>
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
            <label for="datingv" class="col-sm-3 control-label">Dating</label>
            <div class="col-sm-9">
                <input type="text" class="form-control" id="datingv" name="datingv" readonly></input>
            </div>
        </div>
        <div class="form-group form-group-sm">
            <label for="descriptionv" class="col-sm-3 control-label">Description</label>
            <div class="col-sm-offset-9"></div>
			<div class="col-sm-12">
                <textarea class="form-control" rows="5" id="descriptionv" name="descriptionv" readonly></textarea>
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
            <div id="createdv" name="createdv" class="col-sm-4 text-left">Error</div>
            <div class="col-sm-2 text-right"><strong>Modified:</strong></div>
            <div id="modifiedv" name="modifiedv" class="col-sm-4 text-left">Error</div>
        </div>
    </form>

        </div>
    </div>
    </div> <!-- modal-body -->

</div>
</div>
</div>
{{end}}

{{define "modify_dating_modal"}}
<div class="modal fade" id="modifyDatingModal" tabindex="-1" role="dialog" aria-lebeleledby="modifyDatingModalLabel">
<div class="modal-dialog">
<div class="modal-content">

    <div class="modal-header">
    <div class="container-fluid">
        <div class="row">
            <h4 class="modal-title col-md-8" id="modifyDatingModalLabel">Empty Dating Details</h4>
             <button type="button" class="btn btn-primary btn-sm col-sm-2" 
                     onclick="modifyDating('modify_dating_form', $('#hexid').val()); $('#modifyDatingModal').modal('hide');">Modify
             </button>
            <button type="button" class="btn btn-default btn-sm col-md-2" data-dismiss="modal">Cancel</button>
        </div> <!-- row -->
    </div> <!-- container-fluid -->
    </div> <!-- modal-header -->

    <div class="modal-body">
    <div class="container-fluid">
        <div class="row">
        <form id="modify_dating_form" class="form-horizontal">
            <input type="hidden" id="hexid" name="hexid" />
        <div class="form-group form-group-sm">
            <label for="dating" class="col-sm-3 control-label">Dating</label>
            <div class="col-sm-9">
                <input type="text" class="form-control" id="dating" name="dating" readonly></input>
            </div>
        </div>
        <div class="form-group form-group-sm">
            <label for="description" class="col-sm-3 control-label">Description</label>
            <div class="col-sm-offset-9"></div>
			<div class="col-sm-12">
                <textarea class="form-control" rows="5" id="description" name="description"></textarea>
            </div>
        </div>
        {{template "created-modified-modify"}}
   	    </form>
        </div> <!-- row -->
    </div> <!-- container-fluid -->
    </div> <!-- modal-body -->

</div>
</div>
</div>
{{end}}
