/*
 * datings.js - custom JS code dealing with the Datings
 *
 */
function createModalHeading(title) {

    var hdr = document.createElement('div');
    hdr.setAttribute('class', 'modal-header');

    var btn = document.createElement('button');
    btn.setAttribute('class', 'close');
    btn.setAttribute('data-dismiss', 'modal');

    var hdrspan1 = document.createElement('span');
    hdrspan1.setAttribute('aria-hidden', 'true');
    var x = document.createTextNode('x');
    hdrspan1.appendChild(x);
    var hdrspan2 = document.createElement('span');
    hdrspan2.setAttribute('class', 'sr-only');
    var close = document.createTextNode('Close');
    hdrspan2.appendChild(close);

    var title = document.createElement('h4');
    title.setAttribute('class', 'modal-title');
    title.setAttribute('id', 'view-dating-title');
    var titletext = document.createTextNode('View Dating');
    title.appendChild(titletext);

    btn.appendChild(hdrspan1);
    btn.appendChild(hdrspan2);

    hdr.appendChild(btn);
    hdr.appendChild(title);

/*
    <button type="button" class="close" data-dismiss="modal"><span aria-hidden="true">&times;</span><span class="sr-only">Close</span></button>
        <h4 class="modal-title" id="myModalLabel">Modal title</h4>
*/

    return hdr;
}

function createViewModalBody(data) {

    var id = data.Id;
    var dating = data.Dating;
    var descr = data.Description;
//   alert ('openViewWindow: id="' +id+'" ' + dating + ': ' + descr);  // DEBUG

    var body = document.createElement('div');
    body.setAttribute('class', 'modal-body');

    var dat = document.createElement('p');
    var dat_bold = document.createElement('b');
    var dat_title = document.createTextNode('Name:');
    dat_bold.appendChild(dat_title);
    var dat_cont = document.createTextNode(dating + ' [' + id + ']');
    dat.appendChild(dat_bold);
    dat.appendChild(dat_cont);

    var desc = document.createElement('p');
    var desc_title = document.createElement('b');
    var desc_lbl = document.createTextNode('Description:');
    desc_title.appendChild(desc_lbl);

    var desc_brk = document.createElement('br');
    var desc_cont = document.createTextNode(descr);
    desc.appendChild(desc_title);
    desc.appendChild(desc_brk);
    desc.appendChild(desc_cont);

    body.appendChild(dat);
    body.appendChild(desc);

    return body;
}

function createViewWindow(data) {

    var modal = document.createElement('div');
    modal.setAttribute('class', 'modal');
    modal.setAttribute('id', 'view-modal');
    modal.setAttribute('tabindex', '-1');
    modal.setAttribute('role', 'dialog');
    modal.setAttribute('aria-labelledby', '');
    modal.setAttribute('aria-hidden', 'true');
    modal.setAttribute('show', 'false');

    var dlg = document.createElement('div');
    dlg.setAttribute('class', 'modal-dialog');
    modal.appendChild(dlg);

    var cnt = document.createElement('div');
    cnt.setAttribute('class', 'modal-content');
    dlg.appendChild(cnt);

    var hdr = createModalHeading('View Dating');

    var body = createViewModalBody(data);

    cnt.appendChild(hdr);
    cnt.appendChild(body);

    return modal;
}

function openViewWindow(data) {

    // if input data is empty, do nothing... 
    if ((data === null) || (data === undefined)) { 
        alert("openViewWindow: No data received."); // DEBUG
        return; 
    }

    /*
    var parentdiv = document.getElementById('datings-table-div');
    var win = createViewWindow(data);
    parentdiv.appendChild(win);
    */
    $.get('web/templates/datings.mst', function() {

        var view = {
                id: data.Id, 
                dating: data.Dating, 
                description: data.Description};

        var tmplt = document.getElementById('view-dating-template').innerHtml;
        var rendered = Mustache.render(tmplt, view);
       // var parentdiv = document.getElementById('datings-table-div');
        $('#datings-table-div').html(rendered);
       // parentdiv.appendChild(win);
    });

    $('#viewModal').modal('show');

}

function openModifyWindow(data) {

    // if input data is empty, do nothing... 
    if ((data === null) || (data === undefined)) { 
        alert("openModifyWindow: No data received."); // DEBUG
        return; 
    }

    var id = data.Id;
    var dating = data.Dating;
    var descr = data.Description;

    alert ('openModifyWindow: id=' +id+' ' + dating + ': ' + descr);  // DEBUG
}


function createRow(num, rowdata) {

    var row = document.createElement('tr');
    row.setAttribute('id', 'dating-row-' + num);

    var first = document.createElement('td');
    first.textContent = String(num);

    var second = document.createElement('td');
    second.textContent = rowdata.Dating;

    var third = document.createElement('td');
    third.textContent = rowdata.Description;

    var last = document.createElement('td');
    var id = 'dating-actions-' + num;
    last.setAttribute('id', id);

    createActionsCell(num, last, rowdata);

    row.appendChild(first);
    row.appendChild(second);
    row.appendChild(third);
    row.appendChild(last);

    return row;
}

function createActionsCell(num, elem, data) {

    var vspan = document.createElement('span');
    vspan.setAttribute('class', 'glyphicon glyphicon-eye-open');
    //vspan.class = 'glyphicon glyphicon-eye-open';

    var va = document.createElement('a');
    va.setAttribute('data-toggle', 'tooltip modal'); 
    va.setAttribute('data-placement','left');
    va.setAttribute('data-target','#view-modal');
    va.setAttribute('title','View dating details');
    va.setAttribute('href','');
    va.addEventListener('click', function() { openViewWindow(data) }, false); 
    va.appendChild(vspan);

    var mspan = document.createElement('span');
    mspan.setAttribute('class', 'glyphicon glyphicon-cog');
    //mspan.class = 'glyphicon glyphicon-cog';

    var ma = document.createElement('a');
    ma.setAttribute('data-toggle', 'tooltip'); 
    ma.setAttribute('data-placement','left');
    ma.setAttribute('title','Modify dating details');
    ma.setAttribute('href','');
    ma.addEventListener('click', function() { openModifyWindow(data) }, false);
    ma.appendChild(mspan);
        
    var txt = document.createTextNode(' ');

    elem.appendChild(va);
    elem.appendChild(txt);
    elem.appendChild(ma);
}

function createTableHeader() {
    var tblhdr = document.createElement('thead');
    var hdrrow = document.createElement('tr');
    var cell1 = document.createElement('th');
    cell1.textContent = '#';
    var cell2 = document.createElement('th');
    cell2.textContent = 'Dating';
    var cell3 = document.createElement('th');
    cell3.textContent = 'Description';
    var cell4 = document.createElement('th');
    cell4.textContent = 'Actions';
    hdrrow.appendChild(cell1);
    hdrrow.appendChild(cell2);
    hdrrow.appendChild(cell3);
    hdrrow.appendChild(cell4);
    tblhdr.appendChild(hdrrow);
    return tblhdr;
}

// Create a datings table dynamically. 
// The 'data' must be datings' data in JSON format.
function createDatingsTable(dat) {

    if ((dat.length === 0) || (dat === null) || (dat === undefined)) {
        var text = document.createTextElement('p');
        text.textContent = "No datings defined yet.";
        return text;        
    }

    /*
    // create the table
    var tbl = document.createElement('table');
    tbl.setAttribute('class', 'table table-stripped table-hover');
    tbl.setAttribute('id', 'datings-table');

    // create table body
    var body = document.createElement('tbody');
    for (var cnt = 0; cnt < data.length; cnt++) {
        var row = createRow(cnt+1, data[cnt]);
        body.appendChild(row);
    }

    // finally, assemble the complete table...
    tbl.appendChild(createTableHeader()); // this is table header...
    tbl.appendChild(body);
    */

    $.get('templates/mustache/datings.mst', function(dat) {

        var data = {datings: dat,

                    index: function() {
                        return ++window['INDEX']||(window['INDEX']=0);
                    }
   
                };

        var template = $("#datings-table").html();
        var tblhtml = Mustache.render(template, data);
        
        $("#datings-table-div").html(tblhtml); 
    })
        .fail( function(data) {
            alert('template loading failed.'); // DEBUG
            $("#datings-table-div").html('No datings defined currently.');
    });

    //return tbl;
}


function rerouteDating(id, cmd) {

    // check input parameters and correct it, if wrong
    if ( (id === null) || (id === undefined) ) {
        return; // if ID is empty, do nothing...
    }
    if ( (cmd === null) || (cmd === undefined) ) {
        cmd = ""; // correct the command
    }
    
    // create new URL
    var url = "/dating/" + id ;

    if ( (cmd === "view") || (cmd === "modify") ) {
        url = url + "/" + cmd;
    }

    // now redirect to new URL
    //window.location.replace(url);
    //window.location = url;
    //window.location.assign(url);
    window.location.href = url;
}
