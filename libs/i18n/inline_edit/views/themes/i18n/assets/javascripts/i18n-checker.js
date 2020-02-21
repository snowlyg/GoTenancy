if (!window.loadedI18nAsset) {
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

    window.loadedI18nAsset = true;
    var prefix = document.currentScript.getAttribute("data-prefix");
    window.jQuery || loadjscssfile(prefix + "/assets/javascripts/vendors/jquery.min.js", "js");
    loadjscssfile(prefix + "/assets/javascripts/i18n-inline.js?theme=i18n", "js");
    loadjscssfile(prefix + "/assets/stylesheets/i18n-inline.css?theme=i18n", "css");
}
