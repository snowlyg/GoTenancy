// Load style plugin for redactor editor
// By Jason weng @theplant
//
//
// Plugin name: loadstyle
// required options:
// loadStyleNamespace, loadStyleLink
//
/* ********************************* Usage in qor-example:
product.Meta(&admin.Meta{Name: "Description", Config: &admin.RichEditorConfig{AssetManager: assetManager,
    Plugins: []admin.RedactorPlugin{
        {Name: "loadstyle", Source: "/admin/assets/javascripts/qor_redactor_loadstyle.js"},
    },
    Settings: map[string]interface{}{
        "loadStyleNamespace": "yourNamespace",
        "loadStyleLink": "http://your_stylesheets_file_path",
    }
}})
********************************* */

$R.add("plugin", "loadstyle", {
    init: function (app) {
        this.app = app;
        this.opts = app.opts;
    },
    start: function () {
        this.loadStyle();
    },

    loadStyle: function () {
        var $editor = $(this.app.editor.$editor.nodes[0]);

        // loadStyleNamespace, loadStyleLink are required options
        // namespace is needed, but loadStyleLink is not
        if (typeof this.opts.loadStyleNamespace === "undefined") {
            window.alert(
                "please define loadStyleNamespace setting in your config file!"
            );
        } else {
            // add namespace class into editor
            $editor.addClass(this.opts.loadStyleNamespace);
        }

        if (typeof this.opts.loadStyleLink === "undefined") {
            if (typeof this.opts.loadStyleNamespace === "undefined") {
                window.alert(
                    "please define loadStyleLink setting in your config file!"
                );
            }
        } else {
            var ss = document.createElement("link");

            // insert stylesheet
            ss.type = "text/css";
            ss.rel = "stylesheet";
            ss.href = this.opts.loadStyleLink;
            document.getElementsByTagName("head")[0].appendChild(ss);
        }
    }
});
