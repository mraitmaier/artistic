{{define "license"}}
<!DOCTYPE html>
<html lang="en">
<head>
{{template "htmlhead" "License Information"}}
</head>
<body>
    {{template "navbar" .}}

    <div class="container-fluid">
    <div class="row">

        <div class="col-md-2" id="menu">
            <h1 id="menu-header"></h1>

            {{template "accordion"}}
        </div>

<!--
        <div class="col-md-4" id="data-list">
            <h1 id="data-list-header">Data list</h1>

            <p>List</p>
        </div>

        <div class="col-md-6" id="data-details">
            <h1 id="data-details-header">Details</h1>

            </div>
    -->

        <div class="col-md-10" id="data-list">
<h1>BSD Simplified License</h1>
        
<p><b>Copyright 2014 by Miran Raitmaier. All Rights Reserved.</b></p>
<p>Redistribution and use in source and binary forms, with or without 
modification, are permitted provided that the following conditions are met:
</p>
<p>
1. Redistributions of source code must retain the above copyright notice, this
list of conditions and the following disclaimer.
</p>
<p>
2. Redistributions in binary form must reproduce the above copyright notice,
this list of conditions and the following disclaimer in the documentation
and/or other materials provided with the distribution.
</p>
<p>
3. The name of the author may not be used to endorse or promote products
derived from this software without specific prior written permission.
</p>
<p>
THIS SOFTWARE IS PROVIDED BY MIRAN RAITMAIER "AS IS" AND ANY EXPRESS OR IMPLIED
WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF
MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO
EVENT SHALL THE AUTHOR BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL,
EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT
OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS
INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN
CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING
IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY
OF SUCH DAMAGE.
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

    <script> </script>
  </body>
</html>
{{end}}
