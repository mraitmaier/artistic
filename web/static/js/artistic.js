/*
 * artistic.js - custom JS code 
 *
 */

// submit form as POST to a given URL
function postForm(form_id, url) {
    var form = document.getElementById(form_id);
    form.setAttribute("action", url);
    form.setAttribute("method", "post");
    form.submit();
}

// submit a change password form
function changePwd(form_id, id) {
	
	var url = '/pwd/' + id;
	// alert("ID=-" + id + ", url=" + url");
	postForm(form_id, url); 
}

//  convert a slrlng into title case 
function toTitleCase(str) {
    return str.replace(/(?:^|\s)\w/g, function (match) {
        return match.toUpperCase();
    });
}

function validateUserForm() {

    var pwd = document.getElementById("password").value;
    var pwd2 = document.getElementById("password2").value;

    if (pwd !== pwd2) { 
        alert("Passwords do not match!");
        return false; 
    }

    return true;
}

// check 
function checkPasswords(pwd1, pwd2) {
    if (pwd1 === pwd2) { return true; }
    return false;
}

//
function validatePasswordChange() {

    var old = document.getElementById("oldpassword").value;
    var pwd = document.getElementById("newpassword").value;
    var pwd2 = document.getElementById("newpassword2").value;

    if (!checkPasswords(pwd, pwd)) { return false; }

    return true;
}

// send a HTTP request (GET, POST, DELETE, PUT...)
function sendRequest(method, url) {
    var req = new XMLHttpRequest();
    req.open(method, url, true);
    req.send(null);
}

// create an URL to send
function createURL(operation, item, id) {
    var s = "/" + item + "/" + operation + "/" + id;
    return s
}

// send a DELETE HTTP request for:
// - items are users, styles, artists, artworks, techniques....
// - commands (cmd) are "view", "insert", "delete" and "modify" 
// - id is the DB ID of the item
function sendDeleteReq(item, cmd, id ) {
    var url = createURL(cmd, item, id);
    sendRequest('DELETE', url);
}

// aux function for techniques, should be removed I guess...
function rerouteTechnique(method, cmd, id) {
    var url = createURL(cmd, "technique", id);
    sendRequest(method, url);
}

// reroute to URL using ordinary GET HTTP request, used as "onclick" callback
// - items are users, styles, artists, artworks, techniques....
// - commands (cmd) are "view", "insert", "delete", "modify", "changepwd" 
// - id is the DB ID of the item
// When command is to delete the item, first ask for permission.
function rerouteUsingGet(item, cmd, id) {
    if (cmd === "delete") {
        if (!(confirm("Do you really want to delete " + item + "?"))) { return; }
    }
    var url = "/" + item + "/" + cmd + "/" + id;
    window.location.href = url;
}

