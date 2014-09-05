{{define "main-carousel"}}
<div id="main-carousel" class="carousel slide" data-ride="carousel"
                        data-interval="5000">

    <!-- Indicators -->
    <!-- indicators are disabled here... leave it as a comment as an example.
    <ol class="carousel-indicators">
        <li data-target="#main-carousel" data-slide-to="0" class="active"></li>
        <li data-target="#main-carousel" data-slide-to="1"></li>
        <li data-target="#main-carousel" data-slide-to="2"></li>
    </ol>
    -->

    <!-- Wrapper for slides -->
    <div class="carousel-inner">
    
        <div class="item active">
            <img src="static/images/azbe_zamorka.jpg" alt="Anton Ažbe: Zamorka">
            <div class="carousel-caption">
                <p><strong>Anton Ažbe: Zamorka</strong></p>
                <p>Description</p>
            </div>
        </div>

         <div class="item">
            <img src="static/images/grohar_cvetoca_jablana.jpg" alt="Ivan Grohar: Cvetoča jablana">
            <div class="carousel-caption">
                <p><strong>Ivan Grohar: Cvetoča jablana</strong></p>
                <p>Description</p>
            </div>
        </div>

         <div class="item">
            <img src="static/images/kobilca_kofetarca.jpg" alt="Ivana Kobilca: Kofetarica">
            <div class="carousel-caption">
                <p><strong>Ivana Kobilca: Kofetarica</strong></p>
                <p>Description</p>
            </div>
        </div>
        
         <div class="item">
            <img src="static/images/stroj_luiza_pesjakova.jpg" alt="Mihael Stroj: Luiza Pesjakova">
            <div class="carousel-caption">
                <p><strong>Mihael Stroj: Luiza Pesjakova</strong></p>
                <p>Description</p>
            </div>
        </div>

         <div class="item">
            <img src="static/images/jakopic_zima.jpg" alt="Rihard Jakipoč: Zima">
            <div class="carousel-caption">
                <p><strong>Rihard Jakopič:Zima</strong></p>
                <p>Description</p>
            </div>
        </div>

        <div class="item">
            <img src="static/images/kunl_ribji_trg.jpg" alt="Pavel Künl: Ribji trg">
            <div class="carousel-caption">
                <p><strong>Pavel Künl: Ribji trg</strong></p>
                <p>Description</p>
            </div>
        </div>

        <div class="item">
            <img src="static/images/langus_zena.jpg" alt="Matevž Langus: Slikarjeva žena">
            <div class="carousel-caption">
                <p><strong>Matevž Langus: Slikarjeva žena</strong></p>
                <p>Description</p>
            </div>
        </div>

        <div class="item">
            <img src="static/images/petkovsek_tihozitje.jpg" alt="Jožef Petkovšek: Tihožitje">
            <div class="carousel-caption">
                <p><strong>Jožef Petkovšek: Tihožitje</strong></p>
                <p>Description</p>
            </div>
        </div>

        <div class="item">
            <img src="static/images/santel_rdecelasa_deklica.jpg" alt="Henrika Šantel: Rdečelasa deklica">
            <div class="carousel-caption">
                <p><strong>Henrika Šantel: Rdečelasa deklica</strong></p>
                <p>Description</p>
            </div>
        </div>

        <div class="item">
            <img src="static/images/sternen_rdeci_parasol.jpg" alt="Matej Sternen: Rdeči parasol">
            <div class="carousel-caption">
                <p><strong>Matej Sternen: Rdeči parasol</strong></p>
                <p>Description</p>
            </div>
        </div>

        <div class="item">
            <img src="static/images/subic_pred_lovom.jpg" alt="Jurij Šubic: Pred lovom">
            <div class="carousel-caption">
                <p><strong>Jurij Šubic: Pred lovom</strong></p>
                <p>Description</p>
            </div>
        </div>

        <div class="item">
            <img src="static/images/tominc_dama.jpg" alt="Jožef Tominc: Dama s kamelijo">
            <div class="carousel-caption">
                <p><strong>Jožef Tominc: Dama s kamelijo</strong></p>
                <p>Description</p>
            </div>
        </div>

        <div class="item">
            <img src="static/images/vavpotic_plesalka.jpg" alt="Ivan Vavpotič: Rut kot plesalka">
            <div class="carousel-caption">
                <p><strong>Ivan Vavpotič: Rut kot plesalka</strong></p>
                <p>Description</p>
            </div>
        </div>

        <div class="item">
            <img src="static/images/vesel_portret_zene.jpg" alt="Ferdo Vesel: Portret žene">
            <div class="carousel-caption">
                <p><strong>Ferdo Vesel: Portret žene</strong></p>
                <p>Description</p>
            </div>
        </div>

    </div>

    <!-- Controls -->
    <a class="left carousel-control" href="#main-carousel" data-slide="prev">
        <span class="glyphicon glyphicon-chevron-left"></span>
    </a>
    <a class="right carousel-control" href="#main-carousel" data-slide="next">
        <span class="glyphicon glyphicon-chevron-right"></span>
    </a>

</div>
{{end}}
