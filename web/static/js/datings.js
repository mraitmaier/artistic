/*
 * datings.js - custom JS code dealing with the Datings
 *
 */

function createViewWindow(data) {
}

function openViewWindow(data) {

    // if input data is empty, do nothing... 
    if ((data === null) || (data === undefined)) { 
        alert("openViewWindow: No data received."); // DEBUG
        return; 
    }

    var id = data.Id;
    var dating = data.Dating;
    var descr = data.Description;

    alert ('openViewWindow: id="' +id+'" ' + dating + ': ' + descr);  // DEBUG

    var win = createViewWindow(data);

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
    va.setAttribute('data-toggle', 'tooltip'); 
    va.setAttribute('data-placement','left');
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
function createDatingsTable(data) {

    if ((data.length === 0) || (data === null) || (data === undefined)) {
        var text = document.createTextElement('p');
        text.textContent = "No datings defined yet.";
        return text;        
    }

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
    return tbl;
}

