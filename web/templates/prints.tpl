{{define "prints"}}
{{$role := .User.Role}}
{{$name := totitle .Ptype}}
<!DOCTYPE html>
<html lang="en">
<head>
{{template "htmlhead" "Handle Graphic Prints"}}
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
                <h1 id="data-list-header">Graphic {{$name}}s</h1>
        {{if ne $role "guest"}}
        {{template "add-button" $name}}
        {{end}}
                <br />

            {{if .Prints}}
                <table class="table table-striped table-hover small" id="print-list-table">

                <thead>
                  <tr>
                    <th class="col-sm-1">#</th>
                    <th class="col-sm-2">Title</th>
                    <th class="col-sm-2">Artist</th>
                    <th class="col-sm-1">Time</th>
                    <th class="col-sm-1">Technique</th>
                    <th class="col-sm-2">Size</th>
                    <th class="col-sm-2">Location</th>
                    <th class="col-sm-1">Actions</th>
                  </tr>
                </thead>

                <tfoot>
                    <tr class="bg-primary">
                    <td colspan="8"> 
                        <strong> {{.Num}} {{if eq .Num 1}} graphic print {{else}} graphic prints {{end}} found. </strong> 
                    </td>
                    </tr>
                </tfoot>

                <tbody>
                  {{range $index, $element := .Prints}}
                  {{ $cnt := add $index 1 }}
                  <tr id="print-row-{{$cnt}}">
                    <td>{{$cnt}}</td>
                    <td>{{$element.Work.Title}}</td>
                    <td>{{$element.Work.Artist}}</td>
                    <td>{{$element.Work.TimeOfCreation}}</td>
                    <td>{{$element.Work.Technique}}</td>
                    <td>{{$element.Work.Size}}</td>
                    <td>{{$element.Work.Location}}</td>
				    <td>
  						<span data-toggle="tooltip" data-placement="up" title="View details">
                            <a data-toggle="modal" data-target="#viewPrintModal"
                                       data-id="{{$element.ID.Hex}}"
                                       data-created="{{$element.Created}}"
                                       data-modified="{{$element.Modified}}"
                                       data-title="{{$element.Work.Title}}"
                                       data-artist="{{$element.Work.Artist}}"
                                       data-technique="{{$element.Work.Technique}}"
                                       data-artstyl="{{$element.Work.Style}}"
                                       data-size="{{$element.Work.Size}}"
                                       data-dating="{{$element.Work.Dating}}"
                                       data-creat="{{$element.Work.TimeOfCreation}}"
                                       data-motive="{{$element.Work.Motive}}"
                                       data-sign="{{$element.Work.Signature}}"
                                       data-place="{{$element.Work.Place}}"
                                       data-location="{{$element.Work.Location}}"
                                       data-prov="{{$element.Work.Provenance}}"
                                       data-cond="{{$element.Work.Condition}}"
                                       data-conddesc="{{$element.Work.ConditionDescription}}"
                                       data-desc="{{$element.Work.Description}}"
                                       data-exhibit="{{$element.Work.Exhibitions}}"
                                       data-sources="{{$element.Work.Sources}}"
                                       data-notes="{{$element.Work.Notes}}"
                                       data-pic="{{$element.Work.Picture}}">
                                 <span class="glyphicon glyphicon-eye-open"></span>
                            </a>
                       </span>            
                       {{if ne $role "guest"}}
                       &nbsp;&nbsp;
                       <span data-toggle="tooltip" data-placement="up" title="Modify details"> 
                            <a data-toggle="modal" data-target="#modifyPrintModal"
                                       data-id="{{$element.ID.Hex}}"
                                       data-created="{{$element.Created}}"
                                       data-modified="{{$element.Modified}}"
                                       data-title="{{$element.Work.Title}}"
                                       data-artist="{{$element.Work.Artist}}"
                                       data-technique="{{$element.Work.Technique}}"
                                       data-artstyl="{{$element.Work.Style}}"
                                       data-size="{{$element.Work.Size}}"
                                       data-dating="{{$element.Work.Dating}}"
                                       data-creat="{{$element.Work.TimeOfCreation}}"
                                       data-motive="{{$element.Work.Motive}}"
                                       data-sign="{{$element.Work.Signature}}"
                                       data-place="{{$element.Work.Place}}"
                                       data-location="{{$element.Work.Location}}"
                                       data-prov="{{$element.Work.Provenance}}"
                                       data-cond="{{$element.Work.Condition}}"
                                       data-conddesc="{{$element.Work.ConditionDescription}}"
                                       data-desc="{{$element.Work.Description}}"
                                       data-exhibit="{{$element.Work.Exhibitions}}"
                                       data-sources="{{$element.Work.Sources}}"
                                       data-notes="{{$element.Work.Notes}}"
                                       data-pic="{{$element.Work.Picture}}">
                                 <span class="glyphicon glyphicon-edit"></span>
                            </a>
                       </span>            
                       &nbsp;&nbsp;
                       <span data-toggle="tooltip" data-placement="up" title="Remove"> 
                            <a data-toggle="modal" data-target="#removePrintModal"
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
    {{template "view_print_modal"}}
{{if ne $role "guest"}}
    {{template "modify_print_modal"}}
    {{template "remove_print_modal" $name}}
{{end}}
    <!-- end of modals definition -->   

            {{else}}
                <p>No graphic prints found.</p>
            {{end}}

            </div>
        </div> <!-- row -->
    </div> <!-- container fluid -->
{{if ne $role "guest"}}
    {{template "add_print_modal"}}
{{end}}

{{template "insert-js"}}
<script>

    $('#viewPrintModal').on('show.bs.modal', function (event) {

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
        modal.find('.modal-body #technique').text(button.data('technique'));
        modal.find('.modal-body #size').text(button.data('size'));
        modal.find('.modal-body #dating').text(button.data('dating'));
        modal.find('.modal-body #motive').text(button.data('motive'));
        modal.find('.modal-body #signature').text(button.data('sign'));
        modal.find('.modal-body #place').text(button.data('place'));
        modal.find('.modal-body #location').text(button.data('location'));
        modal.find('.modal-body #provenance').text(button.data('prov'));
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
    $('#modifyPrintModal').on('show.bs.modal', function (event) {

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
        modal.find('.modal-title').text('Modify Graphic Print "' + title + '" Details');
        modal.find('.modal-body #hexid').val(button.data('id'));
        modal.find('.modal-body #title').val(title);
        modal.find('.modal-body #timecreat').val(button.data('creat'));
        modal.find('.modal-body #artist').val(button.data('artist'));
        modal.find('.modal-body #artstyle').val(button.data('artstyl'));
        modal.find('.modal-body #technique').val(button.data('technique'));
        modal.find('.modal-body #size').val(button.data('size'));
        modal.find('.modal-body #dating').val(button.data('dating'));
        modal.find('.modal-body #motive').val(button.data('motive'));
        modal.find('.modal-body #signature').val(button.data('sign'));
        modal.find('.modal-body #place').val(button.data('place'));
        modal.find('.modal-body #location').val(button.data('location'));
        modal.find('.modal-body #provenance').val(button.data('prov'));
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
   $('#removePrintModal').on('show.bs.modal', function(event) {
    
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
        var url = '/print/' + id + '/delete';
        $('#removebtn').on('click', function(e) { 
            postForm('remove_print_form', url); 
            $('#removePrintModal').modal('hide');
        });
   });

    // This should post a form to modify print
    var modifyPrint = function(form_id, id) {
        var url = '/print/' + id + '/put';
        postForm(form_id, url);
    }
{{end}}
</script>
</body>
</html>
{{end}}

{{define "add_print_modal"}}
<div class="modal fade" id="addPrintModal" tabindex="-1" role="dialog" aria-labelledby="addPrintModalLabel">
<div class="modal-dialog modal-lg">
<div class="modal-content">

    <div class="modal-header">
    <div class="container-fluid">
        <div class="row">
            <h3 class="modal-title col-sm-8" id="addPrintModalLabel">Add a New Graphic Print</h3>
            <button type="button" class="btn btn-primary btn-sm col-sm-2" 
                    onclick="postForm('add-print-form', '/print'); $('#addPrintModal').modal('hide');"> Add </button>
            <button type="button" class="btn btn-default btn-sm col-sm-2" data-dismiss="modal">Cancel</button>
        </div> <!-- row -->
    </div> <!-- container-fluid -->
    </div> <!-- modal-header -->

    <div class="modal-body">
    <form class="form-horizontal" role="form" method="post" id="add-print-form">
    <fieldset>

        <div class="form-group form-group-sm has-success">
            <label for="title" class="col-sm-3 control-label">Title</label>
            <div class="col-sm-9">
                <input type="text" class="form-control" id="title" name="title" placeholder="Print title" required autofocus />
            </div>
        </div> <!-- form-group -->

        <div class="form-group form-group-sm">
            <label for="artist" class="col-sm-3 control-label">Artist</label>
            <div class="col-sm-9">
                <input type="text" class="form-control" id="artist" name="artist" />
            </div>
        </div> <!-- form-group -->

        <div class="form-group form-group-sm">
            <label for="technique" class="col-sm-3 control-label">Technique</label>
            <div class="col-sm-3">
                <select class="form-control" id="technique" name="technique">
            {{ $techs := get_techniques }}
            {{range $t := $techs}}
                    <option value="{{$t}}">{{$t}}</option>
            {{end}}
                </select>
 
            </div>
            <label for="timecreat" class="col-sm-3 control-label">Time Of Creation</label>
            <div class="col-sm-3">
                <input type="text" class="form-control" id="timecreat" name="timecreat" placeholder="Year ..." />
            </div>
        </div> <!-- form-group -->

        <div class="form-group form-group-sm">
            <label for="artstyle" class="col-sm-3 control-label"> Style </label>
            <div class="col-sm-3">
                <select class="form-control" id="artstyle" name="artstyle">
            {{ $styles := get_styles }}
            {{range $s := $styles}}
                    <option value="{{$s}}">{{$s}}</option>
            {{end}}
                </select>
            </div>
            <label for="dating" class="col-sm-3 control-label"> Dating </label>
            <div class="col-sm-3">
                <select class="form-control" id="dating" name="dating">
            {{ $datings := get_datings }}
            {{range $d := $datings}}
                    <option value="{{$d}}">{{$d}}</option>
            {{end}}
                </select>   
            </div>
        </div> <!-- form-group -->

        <div class="form-group form-group-sm">
            <label for="size" class="col-sm-3 control-label"> Size </label>
            <div class="col-sm-9">
                <input type="text" class="form-control" id="size" name="size" placeholder="Size ..." />
            </div>
        </div> <!-- form-group -->

        <div class="form-group form-group-sm">
            <label for="motive" class="col-sm-3 control-label"> Motive </label>
            <div class="col-sm-9">
                <input type="text" class="form-control" id="motive" name="motive" placeholder="Motive ..." />
            </div>
        </div> <!-- form-group -->

        <div class="form-group form-group-sm">
            <label for="signature" class="col-sm-3 control-label"> Signature </label>
            <div class="col-sm-9">
                <textarea class="form-control" id="signature" name="signature" rows="2"> </textarea>
            </div>
        </div> <!-- form-group -->

        <div class="form-group form-group-sm">
            <label for="place" class="col-sm-3 control-label"> Place </label>
            <div class="col-sm-9">
                <input type="text" class="form-control" id="place" name="place" placeholder="Place ..." />
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

        <div class="form-group form-group-sm">
            <label for="provenance" class="col-sm-3 control-label"> Provenance </label>
            <div class="col-sm-9">
                <textarea class="form-control" id="provenance" name="provenance" rows="5"> </textarea>
            </div>
        </div> <!-- form-group -->

        <!-- TODO:  notes, exhibitions, picture, sources -->

    </fieldset>
    </form>

    </div> <!-- modal-body -->
<!--
    <div class="modal-footer">
    <div class="container-fluid">
        <div class="row text-right">
            <button type="button" class="btn btn-primary btn-sm col-sm-2" 
                    onclick="postForm('add-print-form', '/print'); $('#addPrintModal').modal('hide');"> Add </button>
            <button type="button" class="btn btn-default btn-sm col-sm-2" data-dismiss="modal">Cancel</button>
        </div> <!- row ->
    </div> <!- container-fluid >
    </div> <!- modal-footer ->
-->

</div>
</div>
</div>
{{end}}

{{define "view_print_modal"}}
<div class="modal fade" id="viewPrintModal" tabindex="-1" role="dialog" aria-lebeleledby="ViewPrintModalLabel">
<div class="modal-dialog modal-lg">
<div class="modal-content">

    <div class="modal-header">
    <div class="container-fluid">
        <div class="row">
            <h3 class="modal-title col-md-10" id="viewPrintModalLabel"></h3>
            <button type="button" class="btn btn-default btn-sm col-md-2" data-dismiss="modal">Close</button>
        </div> <!-- row -->
    </div> <!-- container-fluid -->
    </div> <!-- modal-header -->

    <div class="modal-body">
        <div class="container-fluid" id="view-print-table-div">

        <div class="row">
             <table id="view-print-table" class="table table-hover small">
             <tbody>
                 <tr> <td class="col-sm-3"><strong>Artist<strong/></td>
                      <td class="col-sm-9" colspan="3"><span id="artist"></span></td> </tr>
                 <tr> <td class="col-sm-3"><strong>Size<strong/></td>
                      <td class="col-sm-9" colspan="3"><span id="size"></span></td> </tr>
                 <tr> <td class="col-sm-3"><strong>Style</strong></td> 
                      <td class="col-sm-3" id="artstyle"></td>      
                      <td class="col-sm-3"><strong>Dating</strong></td> 
                      <td class="col-sm-3" id="dating"></td> </tr>
                 <tr> <td class="col-sm-3"><strong>Technique<strong/></td>
                      <td class="col-sm-9" colspan="3"><span id="technique"></span></td> </tr>
                 <tr> <td class="col-sm-3"><strong>Signature<strong/></td>
                      <td class="col-sm-9" colspan="3"><span id="signature"></span></td> </tr>
                 <tr> <td class="col-sm-3"><strong>Place<strong/></td>
                      <td class="col-sm-9" colspan="3"><span id="place"></span></td> </tr>
                 <tr> <td class="col-sm-3"><strong>Location<strong/></td>
                      <td class="col-sm-9" colspan="3"><span id="location"></span></td> </tr>
                 <tr> <td class="col-sm-3"><strong>Condition<strong/></td>
                      <td class="col-sm-9" colspan="3"><span id="condition"></span></td> </tr>
                 <tr> <td class="col-sm-3"><strong>Condition Description<strong/></td>
                      <td class="col-sm-9" colspan="3"><span id="conddescription"></span></td> </tr>
                 <tr> <td class="col-sm-3"><strong>Description<strong/></td>
                      <td class="col-sm-9" colspan="3"><span id="description"></span></td> </tr>
                 <tr> <td class="col-sm-3"><strong>Provenance<strong/></td>
                      <td class="col-sm-9"colspan="3"><span id="provenance"></span></td> </tr>
                 <tr> <td class="col-sm-3"><strong>Created</strong></td> 
                      <td class="col-sm-3" id="createdv"></td> 
                      <td class="col-sm-3"><strong>Last Modified</strong></td> 
                      <td class="col-sm-3" id="modifiedv"></td> </tr>
             </tbody>
             </table>
        </div> <!-- row -->
        </div> <!-- container-fluid -->
    </div> <!-- modal-body -->

</div>
</div>
</div>
{{end}}

{{define "modify_print_modal"}}
<div class="modal fade" id="modifyPrintModal" tabindex="-1" role="dialog" aria-lebeleledby="modifyPrintModalLabel">
<div class="modal-dialog modal-lg">
<div class="modal-content">

    <div class="modal-header">
    <div class="container-fluid">
        <div class="row">
            <h3 class="modal-title col-md-8" id="modifyPrintModalLabel">Empty Print Details</h3>
             <button type="button" class="btn btn-primary btn-sm col-sm-2" 
                     onclick="modifyPrint('modify-print-form',$('#hexid').val());$('#modifyPrintModal').modal('hide');">
                     Modify
             </button>
            <button type="button" class="btn btn-default btn-sm col-md-2" data-dismiss="modal">Cancel</button>
        </div> <!-- row -->
    </div> <!-- container-fluid -->
    </div> <!-- modal-header -->

    <div class="modal-body">
    <div class="container-fluid">
        <div class="row">

        <form id="modify-print-form" class="form-horizontal">

        <fieldset>
            <input type="hidden" id="hexid" name="hexid" />

            <div class="form-group form-group-sm has-success">
                <label for="title" class="col-sm-3 control-label">Title</label>
                <div class="col-sm-9">
                    <input type="text" class="form-control" id="title" name="title" placeholder="Print title" required />
                </div>
            </div> <!-- form-group -->

            <div class="form-group form-group-sm">
                <label for="artist" class="col-sm-3 control-label">Artist</label>
                <div class="col-sm-9">
                    <input type="text" class="form-control" id="artist" name="artist" />
                </div>
            </div> <!-- form-group -->

            <div class="form-group form-group-sm">
                <label for="technique" class="col-sm-3 control-label">Technique</label>
                <div class="col-sm-3">
                    <select class="form-control" id="technique" name="technique">
                {{ $techs := get_techniques }}
                {{range $t := $techs}}
                        <option value="{{$t}}">{{$t}}</option>
                {{end}}
                    </select>
     
                </div>
                <label for="timecreat" class="col-sm-3 control-label">Time Of Creation</label>
                <div class="col-sm-3">
                    <input type="text" class="form-control" id="timecreat" name="timecreat" placeholder="Year ..." />
                </div>
            </div> <!-- form-group -->

            <div class="form-group form-group-sm">
                <label for="artstyle" class="col-sm-3 control-label"> Style </label>
                <div class="col-sm-3">
                    <select class="form-control" id="artstyle" name="artstyle">
                {{ $styles := get_styles }}
                {{range $s := $styles}}
                        <option value="{{$s}}">{{$s}}</option>
                {{end}}
                    </select>
                </div>
                <label for="dating" class="col-sm-3 control-label"> Dating </label>
                <div class="col-sm-3">
                    <select class="form-control" id="dating" name="dating">
                {{ $datings := get_datings }}
                {{range $d := $datings}}
                        <option value="{{$d}}">{{$d}}</option>
                {{end}}
                    </select>   
                </div>
            </div> <!-- form-group -->

            <div class="form-group form-group-sm">
                <label for="size" class="col-sm-3 control-label"> Size </label>
                <div class="col-sm-9">
                    <input type="text" class="form-control" id="size" name="size" placeholder="Size ..." />
                </div>
            </div> <!-- form-group -->

            <div class="form-group form-group-sm">
                <label for="motive" class="col-sm-3 control-label"> Motive </label>
                <div class="col-sm-9">
                    <input type="text" class="form-control" id="motive" name="motive" placeholder="Motive ..." />
                </div>
            </div> <!-- form-group -->

            <div class="form-group form-group-sm">
                <label for="signature" class="col-sm-3 control-label"> Signature </label>
                <div class="col-sm-9">
                    <textarea class="form-control" id="signature" name="signature" rows="2"> </textarea>
                </div>
            </div> <!-- form-group -->

            <div class="form-group form-group-sm">
                <label for="place" class="col-sm-3 control-label"> Place </label>
                <div class="col-sm-9">
                    <input type="text" class="form-control" id="place" name="place" placeholder="Place ..." />
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

            <div class="form-group form-group-sm">
                <label for="provenance" class="col-sm-3 control-label"> Provenance </label>
                <div class="col-sm-9">
                    <textarea class="form-control" id="provenance" name="provenance" rows="5"> </textarea>
                </div>
            </div> <!-- form-group -->

            <!-- TODO:  notes, exhibitions, picture, sources -->

            {{template "created-modified-modify"}}
        </fieldset>
        </form>
        </div> <!-- row -->
    </div> <!-- container-fluid -->
    </div> <!-- modal-body -->

</div>
</div>
</div>
{{end}}

{{define "remove_print_modal"}}
{{template "remove-modal" .}}
{{end}}
