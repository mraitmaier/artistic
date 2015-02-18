
{{define "log"}}
    {{ $num := length .Contents }}
<!DOCTYPE html>
<html lang="en">

  <head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <title>Artistic - View Log</title>

    <!-- Bootstrap -->
    <!--  <link href="css/bootstrap.min.css" rel="stylesheet"> -->
    <link href="/static/css/bootstrap.min.css" rel="stylesheet">

    <!-- HTML5 Shim and Respond.js IE8 support of HTML5 elements and media queries -->
    <!-- WARNING: Respond.js doesn't work if you view the page via file:// -->
    <!--[if lt IE 9]>
        <script src="https://oss.maxcdn.com/libs/html5shiv/3.7.0/html5shiv.js"></script>
        <script src="https://oss.maxcdn.com/libs/respond.js/1.4.2/respond.min.js"></script>
    <![endif]-->

    <!-- custom CSS, additional to CSS -->
    <link href="/static/css/custom.css" rel="stylesheet">

    <script>
           
    </script>

  </head>

  <body>
    {{template "navbar" .User.Username}}

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
<!-- jQuery (necessary for Bootstrap's JavaScript plugins) -->
 <!--   <script 
        src="https://ajax.googleapis.com/ajax/libs/jquery/2.1.0/jquery.min.js">
    </script> -->
    <script  src="/static/js/jquery.min.js"></script>

<!-- Include all compiled plugins (below), or include individual files as needed -->
    <script src="/static/js/bootstrap.min.js"></script>

    <script src="/static/js/artistic.js"></script>
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
