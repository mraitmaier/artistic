/*
 * techniques.js - custom JS code dealing with the Techniques
 *
 */

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

function sendDeleteReq(item, cmd, id ) {

    var url = createURL(cmd, item, id);
    sendRequest('DELETE', url);
}

function rerouteTechnique(method, cmd, id) {

    var url = createURL(cmd, "technique", id);
    sendRequest(method, url);
}

function rerouteUsingGet(item, cmd, id) {
    var url = "/" + item + "/" + cmd + "/" + id;
    window.location.href = url;
}

/*
function rerouteTechnique(id, cmd) {

    // check input parameters and correct it, if wrong
    if ( (id === null) || (id === undefined) ) {
        return; // if ID is empty, do nothing...
    }
    if ( (cmd === null) || (cmd === undefined) || (cmd === "")) {
        cmd = "view"; // correct the command
    }
    
    // display a warning pop-up when dealing with deleting items
    if (cmd === "delete") {
        if (!confirm("Do You really want to delete a Technique?")) {
            return;
        }
    }
    // create new URL and redirect to it
    var url = "";
    if ( id === "" ) {
        url = "/technique/" + cmd + "/";
    } else {
        url = "/technique/" + cmd + "/" + id ;
    }
    url = "/technique/" + cmd + "/" + id;
    window.location.href = url;
}
*/
