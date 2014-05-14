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
                <p>Painters</p>
                <p>Sculptors</p>
                <p>Printmakes</p>
                <p>Architects</p>
                <p>Ceramicists</p>
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
    <div class="panel-collapse collapse" id="other-collapse">
        <div class="panel-body">
            <p><a href="datings">Datings</a></p>
            <p><a href="techniques">Techniques</a></p>
            <p><a href="styles">Styles</a></p>
        </div>
    </div>


</div>
{{end}}

