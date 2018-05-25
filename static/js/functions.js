function localDataFields() {
    var form = $('#dest-extra-data')
    form.append("<div class =\"row\">\n" +
        "       <div class=\"col-md-12\">\n" +
        "           <div class=\"form-group\">\n" +
        "               <label for=\"destPath\">Destination Path</label>\n" +
        "               <input type=\"text\" class=\"form-control\" id=\"destPath\" aria-describedby=\"destPathHelp\" placeholder=\"Destination Path\">\n" +
        "               <small id=\"destPathHelp\" class=\"form-text text-muted\">Please choose your detination path.</small>\n" +
        "           </div>\n" +
        "       </div>\n" +
        "    </div>")
}