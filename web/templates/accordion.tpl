{{define  "accordion"}}
<div class="panel-group" id="main-menu" role="tablist" aria-multiselectable="true">

    <div class="panel panel-default" id="artists-menu">
        <div class="panel-heading" role="tab" id="artists-menu-header">
            <h4 class="panel-title">
                <a href="#artists-collapse" data-parent="#main-menu" data-toggle="collapse" 
                   class="accordion-toggle" aria-expanded="true" aria-controls="#artists-collapse">Artists</a>
            </h4>
        </div>
    </div>
    <div class="panel-collapse collapse in" id="artists-collapse">
        <div class="panel-body small" role="tabpanel" aria-labelledby="#artists-menu-header">
                <p><a href="/painter">Painters</a></p>
                <p><a href="/sculptor">Sculptors</a></p>
                <p><a href="/printmaker">Printmakers</a></p>
                <p><a href="/ceramicist">Ceramicists</a></p>
                <p><a href="/architect">Architects</a></p>
        </div>
    </div>    

    <div class="panel panel-default" id="artworks-menu">
        <div class="panel-heading" role="tab" id="artworks-menu-header">
            <h4 class="panel-title">
                <a href="#artworks-collapse" data-parent="#main-menu" data-toggle="collapse" 
                   class="accordion-toggle" aria-expanded="true" aria-controls="#artworks-collapse">Artworks</a>
            </h4>
        </div>
    </div>
    <div class="panel-collapse collapse in" id="artworks-collapse">
        <div class="panel-body small" role="tabpanel" aria-labelledby="#artworks-menu-header">
            <p><a href="/painting">Paintings</a></p>
            <p><a href="/sculpture">Sculptures</a></p>
            <p><a href="/print">Graphic Prints </a></p>
            <p>Ceramics</p>
            <p>Buildings</p>
        </div>
    </div>

    <div class="panel panel-default" id="other-menu">
        <div class="panel-heading" role="tab" id="other-menu-header">
            <h4 class="panel-title">
                <a href="#other-collapse" data-parent="#main-menu" data-toggle="collapse" 
                   class="accordion-toggle" aria-expanded="tabpanel" aria-controls="#other-collapse">Other</a>
            </h4>
        </div>
    </div>
    <div class="panel-collapse collapse in" id="other-collapse">
        <div class="panel-body small" role="tabpanel" aria-labelledby="#other-menu-header">
            <p><a href="/dating">Datings</a></p>
            <p><a href="/technique">Techniques</a></p>
            <p><a href="/style">Styles</a></p>
        </div>
    </div>

</div>
{{end}}

