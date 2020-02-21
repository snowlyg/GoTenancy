if (!window.loadedWidgetAsset) {
    window.loadjscssfile = function (filename, filetype) {
        var fileref;
        if (filetype == "js") {
            fileref = document.createElement('script');
            fileref.setAttribute("type", "text/javascript");
            fileref.setAttribute("src", filename);
        } else if (filetype == "css") {
            fileref = document.createElement("link");
            fileref.setAttribute("rel", "stylesheet");
            fileref.setAttribute("type", "text/css");
            fileref.setAttribute("href", filename);
        }
        if (typeof fileref != "undefined")
            document.getElementsByTagName("head")[0].appendChild(fileref);
    };

    window.loadedWidgetAsset = true;
    var prefix = document.currentScript.getAttribute("data-prefix");

    if (!window.jQuery) {
        loadjscssfile(prefix + "/assets/javascripts/vendors/jquery.min.js", "js");
    }

    loadjscssfile(prefix + "/assets/javascripts/widget_inline_edit.js?theme=widget", "js");
    loadjscssfile(prefix + "/assets/stylesheets/widget_inline_edit.css?theme=widget", "css");
}
