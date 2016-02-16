{{define  "accordion"}}
<div class="panel-group" id="main-menu" role="tablist" aria-multiselectable="true">

    <div class="panel panel-default">
        <div class="panel-heading" role="tab" id="artists-menu-header">
            <h4 class="panel-title">
                <a role="button" data-toggle="collapse" data-parent="#main-menu" href="#artists-collapse"  
                   aria-expanded="true" aria-controls="artists-collapse">Artists</a>
            </h4>
        </div>
    </div>
    <div class="panel-collapse collapse in" id="artists-collapse" role="tabpanel" aria-labelledby="artists-menu-header">
        <div class="panel-body small" >
                <p><a href="/painter"> Painters </a></p>
                <p><a href="/sculptor"> Sculptors </a></p>
                <p><a href="/printmaker"> Printmakers </a></p>
                <p><a href="/architect"> Architects </a></p>
        </div>
    </div>    

    <div class="panel panel-default">
        <div class="panel-heading" role="tab" id="artworks-menu-header">
            <h4 class="panel-title">
                <a role="button" data-toggle="collapse" data-parent="#main-menu" href="#artworks-collapse"   
                   aria-expanded="true" aria-controls="artworks-collapse">Artworks</a>
            </h4>
        </div>
    </div>
    <div class="panel-collapse collapse in" id="artworks-collapse" role="tabpanel" aria-labelledby="artworks-menu-header">
        <div class="panel-body small">
            <p><a href="/painting"> Paintings </a></p>
            <p><a href="/sculpture"> Sculptures </a></p>
            <p><a href="/print"> Graphic Prints </a></p>
            <p><a href="/building"> Buildings </a></p>
        </div>
    </div>

    <div class="panel panel-default">
        <div class="panel-heading" role="tab" id="other-menu-header">
            <h4 class="panel-title">
                <a role="button" data-toggle="collapse" data-parent="#main-menu" href="#other-collapse"
                   aria-expanded="tabpanel" aria-controls="other-collapse">Other</a>
            </h4>
        </div>
    </div>
    <div class="panel-collapse collapse" id="other-collapse" role="tabpanel" aria-labelledby="other-menu-header">
        <div class="panel-body small">
            <p><a href="/book">Books</a></p>
            <p><a href="/article">Articles</a></p>
            <p><a href="/dating">Datings</a></p>
            <p><a href="/technique">Techniques</a></p>
            <p><a href="/style">Styles</a></p>
        </div>
    </div>

</div>
{{end}}

