{{define  "accordion"}}
<div class="panel-group" id="main-menu">

    <div class="panel panel-default" id="artists-menu">
        <div class="panel-heading" id="artists-menu-header">
            <h4 class="panel-title">
                <a href="#artists-collapse" data-parent="#main-menu"
                   data-toggle="collapse" class="accordion-toggle">Artists</a>
            </h4>
        </div>
    </div>
    <div class="panel-collapse collapse in" id="artists-collapse">
        <div class="panel-body">
                <p><a href="/painters">Painters</a></p>
                <p><a href="/sculptors">Sculptors</a></p>
                <p><a href="/printmakers">Printmakers</a></p>
                <p><a href="/architects">Architects</a></p>
                <p><a href="/ceramicists">Ceramicists</a></p>
        </div>
    </div>    

    <div class="panel panel-default" id="artworks-menu">
        <div class="panel-heading">
            <a href="#artworks-collapse" data-parent="#main-menu" 
               data-toggle="collapse" class="accordion-toggle">Artworks</a>
        </div>
    </div>
    <div class="panel-collapse collapse in" id="artworks-collapse">
        <div class="panel-body">
            <p>Paintings</p>
            <p>Sculptures</p>
            <p>Graphic Prints</p>
            <p>Buildings</p>
            <p>Ceramics</p>
        </div>
    </div>

    <div class="panel panel-default" id="other-menu">
        <div class="panel-heading">
            <a href="#other-collapse" data-parent="#main-menu" 
               data-toggle="collapse" class="accordion-toggle">Other</a>
        </div>
    </div>
    <div class="panel-collapse collapse in" id="other-collapse">
        <div class="panel-body">
            <p><a href="/dating">Datings</a></p>
            <p><a href="/technique">Techniques</a></p>
            <p><a href="/style">Styles</a></p>
        </div>
    </div>


</div>
{{end}}

