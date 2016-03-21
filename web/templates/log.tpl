{{define "log"}}
{{$num := length .Contents}}
<!DOCTYPE html>
<html lang="en">
<head>
{{template "htmlhead" "Handle Log"}}
</head>
<body>
    {{template "navbar" .}}

    <div class="container-fluid">

    <div class="row">
        <div class="col-md-2" id="menu">
            <h1 id="menu-header"></h1>
            {{template "accordion"}}
        </div>

        <div class="container-fluid" id=""> 
           <div class="col-md-10" id="data-list">
           <h1>Log</h1>
           <p>Log administation</p>

           <p>
           <button type="button" class="btn btn-primary" data-toggle="modal" data-target="#clearLogModal">Clear Log</button>
           </p>
           <p>Displaying {{$num}} log message(s).</p>


           <!-- modal pop-up definition -->
           <div class="modal fade" id="clearLogModal" 
                                   tabindex="-1" role="dialog" aria-labelledby="clearLogModal" aria-hidden="true">
                <div class="modal-dialog">
                    <div class="modal-content">
                <!--
                        <div class="modal-header">
                            <button type="button" class="close" data-dismiss="modal">
                                <span aria-hidden="true">&times;</span><span class="sr-only">Close</span>
                            </button>    
                            <h4 class="modal-title">Clear log?</h4>
                        </div>
                -->

                        <div class="modal-body">
                            <p>Do you really want to clear the current log?</p>
                        </div>

                        <div class="modal-footer">
                            <button type="button" class="btn btn-default" onclick="clearlog();">Clear</button>
                            <button type="button" class="btn" data-dismiss="modal">Cancel</button>
                        </div>
                    </div>
                </div>

           </div> <!-- modal definition -->                        

           <p>
           <div class="panel panel-primary" id="log-contents-panel">
                <div class="panel-heading" id="log-contents-panel-header">
                     <h3 class="panel-title">Log Contents</h3>
                </div>
                <div class="panel-body" id="log-contents-panel-body">
                    <small>
                    {{range $element := .Contents}}
                        <span class="">&gt;&nbsp; {{printf $element}}</span><br />
                    {{end}}
                    </small>
                    <div id="log-pagination" onLoad="createPagination()"> </div>
                </div>
           </div>
           </p>

        </div>
    </div> <!-- row -->
    </div> <!-- container fluid -->

{{template "insert-js"}}
<script>

        // basic vars
        var perPage = {{.PerPage}};
        var numOfMessages = {{ $num }};

        // calculate the number of pages needed
        function numOfPages(messages, perpage) {

            var pages = messages / perpage;
            if ((messages % perpage) === 0) {
                pages += 1;
            }

            return pages;
        }

        // dynamically create pagination 
        function createPagination() {

            var pages = numOfPages(numOfMessages, perPage);

            var div = document.getElementById('log-pagination');

            var ul = document.createElement('ul');
            var pagattr = document.createAttribute('class');
            pagattr.value = 'pagination pagination-sm';
            ul.setAttributeNode(pagattr);

            var laquo = document.createElement('li');
            var a_laquo = document.createElement('a');
            a_laquo.innerHtml = '&laquo;';
            laquo.appendChild(a_laquo);

            var lt = document.createElement('li');
            var a_lt = document.createElement('a');
            a_lt.innerHtml = '&lt;';
            lt.appendChild(a_lt);

            ul.appendChild(laquo);
            ul.appendChild(lt);

            for (var cnt = 1; cnt <= pages; cnt++) {
                var page = document.createElement('li');
                var a_page = document.createElement('a');
                a_page.innerHtml = string(cnt);
                page.appendChld(a_page);
                ul.appendChild(page);
            }
/*
                <ul class="pagination pagination-sm">
                        <li><a href="#">&laquo;</a></li>
                        <li><a href="#">&lt;</a></li>
                        <li class="active"><a href="#">1</a></li>
                        <li><a href="#">2</a></li>
                        <li><a href="#">3</a></li>
                        <li><a href="#">4</a></li>
                        <li><a href="#">5</a></li>
                        <li><a href="#">&gt;</a></li>
                        <li><a href="#">&raquo;</a></li>
                     <ul/>
*/
            var gt = document.createElement('li');
            var a_gt = document.createElement('a');
            a_gt.innerHtml = '&gt;';
            gt.appendChild(a_gt);

            var raquo = document.createElement('li');
            var a_raquo = document.createElement('a');
            a_raquo.innerHtml = '&raquo;';
            raquo.appendChild(a_raquo);

            ul.appendChild(gt);
            ul.appendChild(raquo);
            
            div.appendChild(ul); // finally append the <ul>
        }

        // display the appropriate messages
        function handlePagination(page, first, last) {

            var first = (page - 1) * perPage + 1;
            var last =  (last > numOfMessages) ? numOfMessages : page * perPage;

            for (var cnt = 1; cnt <= numOfMessages; cnt++) {
            }
        }

        // handle click on 'Clear log' button on modal: send POST request to itself and hide the modal
        function clearlog() {

            $('#clearLogModal').modal('hide');
            sendRequest('post', '/log');
        }
 
    </script>
  </body>
</html>
{{end}}
