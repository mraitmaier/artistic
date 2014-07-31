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

// send a DELETE HTTP request for:
// - items are users, styles, artists, artworks, techniques....
// - commands (cmd) are "view", "create", "delete" and "modify" 
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
// - commands (cmd) are "view", "create", "delete" and "modify" 
// - id is the DB ID of the item
function rerouteUsingGet(item, cmd, id) {
    var url = "/" + item + "/" + cmd + "/" + id;
    window.location.href = url;
}

