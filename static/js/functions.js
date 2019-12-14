function localDataFields() {
    var form = $('#dest-extra-data')
    form.append("<div class =\"row\">\n" +
        "       <div class=\"col-md-12\">\n" +
        "           <div class=\"form-group\">\n" +
        "               <label for=\"path\">Repository Path</label>\n" +
        "               <input type=\"text\" class=\"form-control\" id=\"path\" aria-describedby=\"pathHelp\" placeholder=\"/path/to/backup\" required>\n" +
        "           </div>\n" +
        "       </div>\n" +
        "    </div>")
}

function sftpDataFields() {
    var form = $('#dest-extra-data')
    form.append("<div class =\"row\">\n" +
        "       <div class=\"col-md-12\">\n" +
        "           <div class=\"form-group\">\n" +
        "               <label for=\"server\">Repository Server</label>\n" +
        "               <input type=\"text\" class=\"form-control\" id=\"server\" aria-describedby=\"pathHelp\" placeholder=\"host\" required>\n" +
        "           </div>\n" +
        "       </div>\n" +
        "    </div>").append("<div class =\"row\">\n" +
        "       <div class=\"col-md-12\">\n" +
        "           <div class=\"form-group\">\n" +
        "               <label for=\"user\">Repository User</label>\n" +
        "               <input type=\"text\" class=\"form-control\" id=\"user\" aria-describedby=\"pathHelp\" placeholder=\"user_name\" required>\n" +
        "           </div>\n" +
        "       </div>\n" +
        "    </div>").append("<div class =\"row\">\n" +
        "       <div class=\"col-md-12\">\n" +
        "           <div class=\"form-group\">\n" +
        "               <label for=\"path\">Repository Path</label>\n" +
        "               <input type=\"text\" class=\"form-control\" id=\"path\" aria-describedby=\"pathHelp\" placeholder=\"/path/to/backup\" required>\n" +
        "           </div>\n" +
        "       </div>\n" +
        "    </div>")
}

function bbDataFields() {
    var form = $('#dest-extra-data')
    form.append("<div class =\"row\">\n" +
        "       <div class=\"col-md-12\">\n" +
        "           <div class=\"form-group\">\n" +
        "               <label for=\"B2_ACCOUNT_ID\">B2 Account ID</label>\n" +
        "               <input type=\"text\" class=\"form-control\" id=\"B2_ACCOUNT_ID\" aria-describedby=\"pathHelp\" placeholder=\"B2 Account ID\" required>\n" +
        "           </div>\n" +
        "       </div>\n" +
        "    </div>").append("<div class =\"row\">\n" +
        "       <div class=\"col-md-12\">\n" +
        "           <div class=\"form-group\">\n" +
        "               <label for=\"B2_ACCOUNT_KEY\">B2 Account Key</label>\n" +
        "               <input type=\"text\" class=\"form-control\" id=\"B2_ACCOUNT_KEY\" aria-describedby=\"pathHelp\" placeholder=\"B2 Account Key\" required>\n" +
        "           </div>\n" +
        "       </div>\n" +
        "    </div>").append("<div class =\"row\">\n" +
        "       <div class=\"col-md-12\">\n" +
        "           <div class=\"form-group\">\n" +
        "               <label for=\"bucket_name\">Bucket Name</label>\n" +
        "               <input type=\"text\" class=\"form-control\" id=\"bucket_name\" aria-describedby=\"pathHelp\" placeholder=\"Bucket Name\" required>\n" +
        "           </div>\n" +
        "       </div>\n" +
        "    </div>").append("<div class =\"row\">\n" +
        "       <div class=\"col-md-12\">\n" +
        "           <div class=\"form-group\">\n" +
        "               <label for=\"path\">Repository Path</label>\n" +
        "               <input type=\"text\" class=\"form-control\" id=\"path\" aria-describedby=\"pathHelp\" placeholder=\"/path/to/backup\" required>\n" +
        "           </div>\n" +
        "       </div>\n" +
        "    </div>")
}

function s3DataFields() {
    var form = $('#dest-extra-data')
    form.append("<div class =\"row\">\n" +
        "       <div class=\"col-md-12\">\n" +
        "           <div class=\"form-group\">\n" +
        "               <label for=\"AWS_ACCESS_KEY_ID\">AWS Access Key</label>\n" +
        "               <input type=\"text\" class=\"form-control\" id=\"AWS_ACCESS_KEY_ID\" aria-describedby=\"pathHelp\" placeholder=\"AWS Access Key\" required>\n" +
        "           </div>\n" +
        "       </div>\n" +
        "    </div>").append("<div class =\"row\">\n" +
        "       <div class=\"col-md-12\">\n" +
        "           <div class=\"form-group\">\n" +
        "               <label for=\"AWS_SECRET_ACCESS_KEY\">AWS Secret Key</label>\n" +
        "               <input type=\"text\" class=\"form-control\" id=\"AWS_SECRET_ACCESS_KEY\" aria-describedby=\"pathHelp\" placeholder=\"AWS Secret Key\" required>\n" +
        "           </div>\n" +
        "       </div>\n" +
        "    </div>").append("<div class =\"row\">\n" +
        "       <div class=\"col-md-12\">\n" +
        "           <div class=\"form-group\">\n" +
        "               <label for=\"bucket_name\">Bucket Name</label>\n" +
        "               <input type=\"text\" class=\"form-control\" id=\"bucket_name\" aria-describedby=\"pathHelp\" placeholder=\"Bucket Name\" required>\n" +
        "           </div>\n" +
        "       </div>\n" +
        "    </div>").append("<div class =\"row\">\n" +
        "       <div class=\"col-md-12\">\n" +
        "           <div class=\"form-group\">\n" +
        "               <label for=\"path\">Repository Path</label>\n" +
        "               <input type=\"text\" class=\"form-control\" id=\"path\" aria-describedby=\"pathHelp\" placeholder=\"/path/to/backup\" required>\n" +
        "           </div>\n" +
        "       </div>\n" +
        "    </div>")
}