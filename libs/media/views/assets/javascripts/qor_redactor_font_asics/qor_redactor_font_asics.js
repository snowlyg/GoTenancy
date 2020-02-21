$R.add("plugin", "font", {
    translations: {
        en: {
            font: "Font",
            "font-clear": "Clear Font",
            "font-asics": "ASICS"
        }
    },
    init: function (app) {
        this.app = app;
        this.opts = app.opts;
        this.lang = app.lang;
        this.toolbar = app.toolbar;
        this.inline = app.inline;
    },
    // public
    start: function () {
        var dropdown = {
            asics: {
                title: this.lang.get("font-asics"),
                api: "plugin.font.asics"
            },
            clear: {
                title: this.lang.get("font-clear"),
                api: "plugin.font.clear"
            }
        };

        var $button = this.toolbar.addButton("font", {
            title: this.lang.get("font")
        });
        $button.setIcon('<i class="re-icon-fontfamily"></i>');
        $button.setDropdown(dropdown);
    },

    asics: function () {
        this.clear();

        var args = {
            tag: "span",
            class: `font-asics`,
            type: "toggle"
        };

        this.inline.format(args);
    },
    clear: function () {
        this.inline.remove({class: `font-asics`});
    }
});
