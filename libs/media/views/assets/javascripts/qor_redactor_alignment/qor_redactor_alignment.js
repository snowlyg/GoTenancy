$R.add("plugin", "alignment", {
    translations: {
        en: {
            align: "Align",
            "align-left": "Align Left",
            "align-center": "Align Center",
            "align-right": "Align Right",
            "align-justify": "Align Justify"
        }
    },
    init: function (app) {
        this.app = app;
        this.opts = app.opts;
        this.lang = app.lang;
        this.block = app.block;
        this.selection = app.selection;
        this.toolbar = app.toolbar;
    },
    // public
    start: function () {
        var dropdown = {};

        dropdown.left = {
            title: this.lang.get("align-left"),
            api: "plugin.alignment.set",
            args: "left"
        };
        dropdown.center = {
            title: this.lang.get("align-center"),
            api: "plugin.alignment.set",
            args: "center"
        };
        dropdown.right = {
            title: this.lang.get("align-right"),
            api: "plugin.alignment.set",
            args: "right"
        };
        dropdown.justify = {
            title: this.lang.get("align-justify"),
            api: "plugin.alignment.set",
            args: "justify"
        };

        var $button = this.toolbar.addButton("alignment", {
            title: this.lang.get("align")
        });
        $button.setIcon('<i class="re-icon-alignment"></i>');
        $button.setDropdown(dropdown);
    },
    set: function (type) {
        var currentBlocks = this.selection.getBlocks();
        currentBlocks.forEach(function (currentBlock) {
            $R
                .dom(currentBlock)
                .removeClass(
                    `rd-text-left rd-text-right rd-text-center rd-text-justify`
                )
                .addClass(`rd-text-${type}`);
        });
    }
});
