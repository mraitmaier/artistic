{{define "buildings"}}
{{$role := .User.Role}}
<!DOCTYPE html>
<html lang="en">
<head>
{{template "htmlhead" "Handle Buildings"}}
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
                <h1 id="data-list-header">Buildings</h1>

    {{if ne $role "guest"}}
                <div id="new-building-btn">
                <button type="button" class="btn btn-primary btn-sm" data-toggle="modal" data-target="#addBuildingModal">
                <span class="glyphicon glyphicon-plus"></span> &nbsp; Add a New Building
                </button>
        	    </div>
    {{end}}
                <br />

            {{if .Buildings}}
                <table class="table table-striped table-hover small" id="building-list-table">

                <thead>
                  <tr>
                    <th class="col-sm-1">#</th>
                    <th class="col-sm-3">Name</th>
                    <th class="col-sm-3">Architect(s)</th>
                    <th class="col-sm-1">Time</th>
                    <th class="col-sm-3">Location</th>
                    <th class="col-sm-1">Actions</th>
                  </tr>
                </thead>

                <tfoot>
                    <tr class="bg-primary">
                    <td colspan="6"> 
                        <strong> {{.Num}} {{if eq .Num 1}} building {{else}} buildings {{end}} found. </strong> 
                    </td>
                    </tr>
                </tfoot>

                <tbody>
                  {{range $index, $element := .Buildings}}
                  {{ $cnt := add $index 1 }}
                  <tr id="building-row-{{$cnt}}">
                    <td>{{$cnt}}</td>
                    <td>{{$element.Work.Title}}</td>
                    <td>{{$element.Work.Artist}}</td>
                    <td>{{$element.Work.TimeOfCreation}}</td>
                    <td>{{$element.Work.Location}}</td>
				    <td>
  						<span data-toggle="tooltip" data-placement="up" title="View details">
                            <a data-toggle="modal" data-target="#viewBuildingModal"
                                       data-id="{{$element.ID.Hex}}"
                                       data-created="{{$element.Created}}"
                                       data-modified="{{$element.Modified}}"
                                       data-title="{{$element.Work.Title}}"
                                       data-artist="{{$element.Work.Artist}}"
                                       data-artstyl="{{$element.Work.Style}}"
                                       data-size="{{$element.Work.Size}}"
                                       data-creat="{{$element.Work.TimeOfCreation}}"
                                       data-location="{{$element.Work.Location}}"
                                       data-cond="{{$element.Work.Condition}}"
                                       data-conddesc="{{$element.Work.ConditionDescription}}"
                                       data-desc="{{$element.Work.Description}}"
                                       data-sources="{{$element.Work.Sources}}"
                                       data-notes="{{$element.Work.Notes}}"
                                       data-pic="{{$element.Work.Picture}}">
                                 <span class="glyphicon glyphicon-eye-open"></span>
                            </a>
                       </span>            
                       {{if ne $role "guest"}}
                       &nbsp;&nbsp;
                       <span data-toggle="tooltip" data-placement="up" title="Modify details"> 
                            <a data-toggle="modal" data-target="#modifyBuildingModal"
                                       data-id="{{$element.ID.Hex}}"
                                       data-created="{{$element.Created}}"
                                       data-modified="{{$element.Modified}}"
                                       data-title="{{$element.Work.Title}}"
                                       data-artist="{{$element.Work.Artist}}"
                                       data-artstyl="{{$element.Work.Style}}"
                                       data-size="{{$element.Work.Size}}"
                                       data-creat="{{$element.Work.TimeOfCreation}}"
                                       data-location="{{$element.Work.Location}}"
                                       data-cond="{{$element.Work.Condition}}"
                                       data-conddesc="{{$element.Work.ConditionDescription}}"
                                       data-desc="{{$element.Work.Description}}"
                                       data-sources="{{$element.Work.Sources}}"
                                       data-notes="{{$element.Work.Notes}}"
                                       data-pic="{{$element.Work.Picture}}">
                                 <span class="glyphicon glyphicon-edit"></span>
                            </a>
                       &nbsp;&nbsp;
                       <span data-toggle="tooltip" data-placement="up" title="Remove"> 
                            <a data-toggle="modal" data-target="#removeBuildingModal"
                                       data-id="{{$element.ID.Hex}}"  
                                       data-title="{{$element.Work.Title}}">
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
    {{template "view_building_modal"}}
{{if ne $role "guest"}}
    {{template "modify_building_modal"}}
    {{template "remove_building_modal"}}
{{end}}
    <!-- end of modals definition -->   

            {{else}}
                <p>No buildings found.</p>
            {{end}}

            </div>
        </div> <!-- row -->
    </div> <!-- container fluid -->

{{if ne $role "guest"}}
    {{template "add_building_modal"}}
{{end}}
    <!-- jQuery (necessary for Bootstrap's JavaScript plugins) -->
    <!--   <script src="https://ajax.googleapis.com/ajax/libs/jquery/2.1.0/jquery.min.js"></script> -->
    <script  src="/static/js/jquery.min.js"></script>
    <!-- Include all compiled plugins, or include individual files as needed -->
    <script src="/static/js/bootstrap.min.js"></script>
    <!-- custom JS code -->
    <script src="/static/js/artistic.js"></script>
    <script>

    $('#viewBuildingModal').on('show.bs.modal', function (event) {

        var button = $(event.relatedTarget);     // Button that triggered the modal
        // Extract info from data-* requirement attribute
        var id = button.data('id');  
        var creat = button.data('creat');
        //var exhibit = button.data('exhibit');
        //var sources = button.data('sources');
        //var notes =  button.data('notes');
        //var pic = button.data('pic');

        // Update the modal's content. We'll use jQuery here, but you could use a data 
        // binding library or other methods instead.
        var modal = $(this)
        modal.find('.modal-title').text(button.data('title') + ' (' + creat + ')');
        modal.find('.modal-body #artist').text(button.data('artist'));
        modal.find('.modal-body #artstyle').text(button.data('artstyl'));
        modal.find('.modal-body #size').text(button.data('size'));
        modal.find('.modal-body #location').text(button.data('location'));
        modal.find('.modal-body #condition').text(button.data('cond'));
        modal.find('.modal-body #conddescription').text(button.data('conddesc'));
        modal.find('.modal-body #description').text(button.data('desc'));
        //modal.find('.modal-body #exbitions').text(exhibit);
        //modal.find('.modal-body #sources').text(sources);
        //modal.find('.modal-body #notes').text(notes);
        //modal.find('.modal-body #picture').text(pic);
        modal.find('.modal-body #createdv').text(button.data('created'));
        modal.find('.modal-body #modifiedv').text(button.data('modified'));
    })

{{if ne $role "guest"}}
    $('#modifyBuildingModal').on('show.bs.modal', function (event) {

        var button = $(event.relatedTarget);     // Button that triggered the modal
        // Extract info from data-* requirement attribute
        var title = button.data('title');
        var created = button.data('created');
        //var exhibit = button.data('exhibit');
        //var sources = button.data('sources');
        //var notes =  button.data('notes');
        //var pic = button.data('pic');

        // Update the modal's content. We'll use jQuery here, but you could use a data 
        // binding library or other methods instead.
        var modal = $(this)
        modal.find('.modal-title').text('Modify Building "' + title + '" Details');
        modal.find('.modal-body #hexid').val(button.data('id'));
        modal.find('.modal-body #title').val(title);
        modal.find('.modal-body #timecreat').val(button.data('creat'));
        modal.find('.modal-body #artist').val(button.data('artist'));
        modal.find('.modal-body #artstyle').val(button.data('artstyl'));
        modal.find('.modal-body #size').val(button.data('size'));
        modal.find('.modal-body #location').val(button.data('location'));
        modal.find('.modal-body #condition').val(button.data('cond'));
        modal.find('.modal-body #conddescription').val(button.data('conddesc'));
        modal.find('.modal-body #description').val(button.data('desc'));
        //modal.find('.modal-body #exbitions').val(exhibit);
        //modal.find('.modal-body #sources').val(sources);
        //modal.find('.modal-body #notes').val(notes);
        //modal.find('.modal-body #picture').val(pic);
        modal.find('.modal-body #created').val(created);
        modal.find('.modal-body #createdm').text(created);
        modal.find('.modal-body #modifiedm').text(button.data('modified'));
    })

// Handle the removals using modal pop-up 
   $('#removeBuildingModal').on('show.bs.modal', function(event) {
    
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
        var url = '/building/' + id + '/delete';
        $('#removebtn').on('click', function(e) { 
            postForm('remove_building_form', url); 
            $('#removeBuildingModal').modal('hide');
        });
   });

    // This should post a form to modify building
    var modifyBuilding = function(form_id, id) {
        var url = '/building/' + id + '/put';
        postForm(form_id, url);
    }
{{end}}
    </script>
</body>
</html>
{{end}}

{{define "add_building_modal"}}
<div class="modal fade" id="addBuildingModal" tabindex="-1" role="dialog" aria-labelledby="addBuildingModalLabel">
<div class="modal-dialog modal-lg">
<div class="modal-content">

    <div class="modal-header">
    <div class="container-fluid">
        <div class="row">
            <h3 class="modal-title col-sm-8" id="addBuildingModalLabel">Add a New Building</h3>
            <button type="button" class="btn btn-primary btn-sm col-sm-2" 
                    onclick="postForm('add-building-form', '/building'); $('#addBuildingModal').modal('hide');"> Add </button>
            <button type="button" class="btn btn-default btn-sm col-sm-2" data-dismiss="modal">Cancel</button>
        </div> <!-- row -->
    </div> <!-- container-fluid -->
    </div> <!-- modal-header -->

    <div class="modal-body">
    <form class="form-horizontal" role="form" method="post" id="add-building-form">
    <fieldset>

        <div class="form-group form-group-sm has-success">
            <label for="title" class="col-sm-3 control-label">Name</label>
            <div class="col-sm-9">
                <input type="text" class="form-control" id="title" name="title" placeholder="Building name" required />
            </div>
        </div> <!-- form-group -->

        <div class="form-group form-group-sm">
            <label for="artist" class="col-sm-3 control-label">Architect(s)</label>
            <div class="col-sm-9">
                <input type="text" class="form-control" id="artist" name="artist" />
            </div>
        </div> <!-- form-group -->

        <div class="form-group form-group-sm">
            <label for="timecreat" class="col-sm-3 control-label">Time Of Creation</label>
            <div class="col-sm-3">
                <input type="text" class="form-control" id="timecreat" name="timecreat" placeholder="Year ..." />
            </div>
            <label for="artstyle" class="col-sm-3 control-label"> Art Style </label>
            <div class="col-sm-3">
                <select class="form-control" id="artstyle" name="artstyle">
            {{ $styles := get_styles }}
            {{range $s := $styles}}
                    <option value="{{$s}}">{{$s}}</option>
            {{end}}
                </select>
            </div>
        </div> <!-- form-group -->

        <div class="form-group form-group-sm">
            <label for="size" class="col-sm-3 control-label"> Dimensions </label>
            <div class="col-sm-9">
                <input type="text" class="form-control" id="size" name="size" placeholder="Dimensions" />
            </div>
        </div> <!-- form-group -->

        <div class="form-group form-group-sm">
            <label for="location" class="col-sm-3 control-label"> Location </label>
            <div class="col-sm-9">
                <input type="text" class="form-control" id="location" name="location" placeholder="Location ..." />
            </div>
        </div> <!-- form-group -->

        <div class="form-group form-group-sm">
            <label for="condition" class="col-sm-3 control-label"> Condition </label>
            <div class="col-sm-9">
                <input type="text" class="form-control" id="condition" name="condition" placeholder="Condition ..." />
            </div>
        </div> <!-- form-group -->

        <div class="form-group form-group-sm">
            <label for="conddescription" class="col-sm-3 control-label"> Condition Description </label>
            <div class="col-sm-9">
                <textarea class="form-control" id="conddescription" name="conddescription" rows="2"> </textarea>
            </div>
        </div> <!-- form-group -->

        <div class="form-group form-group-sm">
            <label for="description" class="col-sm-3 control-label"> Description </label>
            <div class="col-sm-9">
                <textarea class="form-control" id="description" name="description" rows="5"> </textarea>
            </div>
        </div> <!-- form-group -->

        <!-- TODO:  notes, picture, sources -->

    </fieldset>
    </form>

    </div> <!-- modal-body -->

</div>
</div>
</div>
{{end}}

{{define "view_building_modal"}}
<div class="modal fade" id="viewBuildingModal" tabindex="-1" role="dialog" aria-lebeleledby="ViewBuildingModalLabel">
<div class="modal-dialog modal-lg">
<div class="modal-content">

    <div class="modal-header">
    <div class="container-fluid">
        <div class="row">
            <h3 class="modal-title col-md-10" id="viewBuildingModalLabel"></h3>
            <button type="button" class="btn btn-default btn-sm col-md-2" data-dismiss="modal">Close</button>
        </div> <!-- row -->
    </div> <!-- container-fluid -->
    </div> <!-- modal-header -->

    <div class="modal-body">
        <div class="container-fluid" id="view-building-table-div">

        <div class="row">
             <table id="view-user-table" class="table table-hover small">
             <tbody>
                 <tr> <td class="col-sm-3"><strong>Architect(s)<strong/></td>
                      <td class="col-sm-9"><span id="artist"></span></td> </tr>
                 <tr> <td class="col-sm-3"><strong>Dimensions<strong/></td>
                      <td class="col-sm-9"><span id="size"></span></td> </tr>
                 <tr> <td class="col-sm-3"><strong>Art Style</strong></td> 
                      <td class="col-sm-9" id="artstyle"></td>      </tr>
                 <tr> <td class="col-sm-3"><strong>Location<strong/></td>
                      <td class="col-sm-9"><span id="location"></span></td> </tr>
                 <tr> <td class="col-sm-3"><strong>Condition<strong/></td>
                      <td class="col-sm-9"><span id="condition"></span></td> </tr>
                 <tr> <td class="col-sm-3"><strong>Condition Description<strong/></td>
                      <td class="col-sm-9"><span id="conddescription"></span></td> </tr>
                 <tr> <td class="col-sm-3"><strong>Description<strong/></td>
                      <td class="col-sm-9"><span id="description"></span></td> </tr>
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

{{define "modify_building_modal"}}
<div class="modal fade" id="modifyBuildingModal" tabindex="-1" role="dialog" aria-lebeleledby="modifyBuildingModalLabel">
<div class="modal-dialog modal-lg">
<div class="modal-content">

    <div class="modal-header">
    <div class="container-fluid">
        <div class="row">
            <h3 class="modal-title col-md-8" id="modifyBuildingModalLabel">Empty Building Details</h3>
             <button type="button" class="btn btn-primary btn-sm col-sm-2" 
                     onclick="modifyBuilding('modify-building-form',$('#hexid').val());$('#modifyBuildingModal').modal('hide');">
                     Modify
             </button>
            <button type="button" class="btn btn-default btn-sm col-md-2" data-dismiss="modal">Cancel</button>
        </div> <!-- row -->
    </div> <!-- container-fluid -->
    </div> <!-- modal-header -->

    <div class="modal-body">
    <div class="container-fluid">
        <div class="row">

        <form id="modify-building-form" class="form-horizontal">

        <fieldset>
            <input type="hidden" id="hexid" name="hexid" />

            <div class="form-group form-group-sm has-success">
                <label for="title" class="col-sm-3 control-label">Name</label>
                <div class="col-sm-9">
                    <input type="text" class="form-control" id="title" name="title" required />
                </div>
            </div> <!-- form-group -->

            <div class="form-group form-group-sm">
                <label for="artist" class="col-sm-3 control-label">Architect(s)</label>
                <div class="col-sm-9">
                    <input type="text" class="form-control" id="artist" name="artist" />
                </div>
            </div> <!-- form-group -->

            <div class="form-group form-group-sm">
                <label for="timecreat" class="col-sm-3 control-label">Time Of Creation</label>
                <div class="col-sm-3">
                    <input type="text" class="form-control" id="timecreat" name="timecreat" placeholder="Year ..." />
                </div>
                <label for="artstyle" class="col-sm-3 control-label"> Art Style </label>
                <div class="col-sm-3">
                    <select class="form-control" id="artstyle" name="artstyle">
                {{ $styles := get_styles }}
                {{range $s := $styles}}
                        <option value="{{$s}}">{{$s}}</option>
                {{end}}
                    </select>
                </div>
            </div> <!-- form-group -->

            <div class="form-group form-group-sm">
                <label for="size" class="col-sm-3 control-label"> Dimensions </label>
                <div class="col-sm-9">
                    <input type="text" class="form-control" id="size" name="size" placeholder="Size ..." />
                </div>
            </div> <!-- form-group -->

            <div class="form-group form-group-sm">
                <label for="location" class="col-sm-3 control-label"> Location </label>
                <div class="col-sm-9">
                    <input type="text" class="form-control" id="location" name="location" placeholder="Location ..." />
                </div>
            </div> <!-- form-group -->

            <div class="form-group form-group-sm">
                <label for="condition" class="col-sm-3 control-label"> Condition </label>
                <div class="col-sm-9">
                    <input type="text" class="form-control" id="condition" name="condition" placeholder="Condition ..." />
                </div>
            </div> <!-- form-group -->

            <div class="form-group form-group-sm">
                <label for="conddescription" class="col-sm-3 control-label"> Condition Description </label>
                <div class="col-sm-9">
                    <textarea class="form-control" id="conddescription" name="conddescription" rows="2"> </textarea>
                </div>
            </div> <!-- form-group -->

            <div class="form-group form-group-sm">
                <label for="description" class="col-sm-3 control-label"> Description </label>
                <div class="col-sm-9">
                    <textarea class="form-control" id="description" name="description" rows="5"> </textarea>
                </div>
            </div> <!-- form-group -->

            <!-- TODO:  notes, picture, sources -->

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

{{define "remove_building_modal"}}
<div class="modal fade" id="removeBuildingModal" tabindex="-1" role="dialog" aria-labelledby="removeBuildingModalLabel">
<div class="modal-dialog">
<div class="modal-content">

    <div class="modal-header">
    <div class="container-fluid">
        <div class="row">
            <h3 class="modal-title col-sm-8" id="removeBuildingModalLabel">Remove Building</h3>
            <button type="button" class="btn btn-primary btn-sm col-sm-2" id="removebtn"> Remove </button>
            <button type="button" class="btn btn-default btn-sm col-sm-2" data-dismiss="modal"> Cancel </button>
        </div> <!-- row -->
    </div> <!-- container-fluid -->
    </div> <!-- modal-header -->

    <div class="modal-body">
    <p> Would you really like to remove the building '<span id="removename"></span>'?</p>
    <form method="post" id="remove_building_form">
        <input type="hidden" name="id" id="id" />
        <input type="hidden" name="name" id="name" />
    </form>
    </div>
</div>
</div>
</div>
{{end}}
