// Media display mode for redactor editor
// By Jason weng @theplant
// // product.Meta(&admin.Meta{Name: "Description", Config: &admin.RichEditorConfig{Plugins: []admin.RedactorPlugin{
// 	{Name: "mediadisplaymode", Source: "/admin/assets/javascripts/qor_redactor_media_display_mode.js"},
// },
// 	Settings: map[string]interface{}{
// 		"mediaDisplayModeOptions": "Left|Right|Container Width|Full Width",
// 	},
// }})
// rendered classname:
// rd-display-left, rd-display-right, rd-display-containerwidth, rd-display-fullwidth

$R.add("plugin", "mediadisplaymode", {
    langs: {
        en: {
            mediaDisplayMode: "Media Display Mode",
            mediaDisplayModeButton: "Display mode"
        }
    },
    getEditterButton: function () {
        return `<a href="#"><span id="redactor-image-displaymode" data-redactor="verified" contenteditable="false">mediaDisplayModeButton</span></a>`;
    },
    init: function (app) {
        this.app = app;
        this.opts = app.opts;
        this.lang = app.lang;
        this.block = app.block;
        this.toolbar = app.toolbar;

        this.mediadisplaymode = {};
    },
    start: function () {
        const $editor = $(this.app.editor.$editor.nodes[0]);

        let mediaDisplayModeOptions = this.opts.mediaDisplayModeOptions || "",
            $modes = "",
            $imgs = $editor.find("img");

        if (
            this.opts.type === "pre" ||
            this.opts.type === "inline" ||
            mediaDisplayModeOptions === ""
        ) {
            return;
        }

        mediaDisplayModeOptions.split(/\||,/).forEach(mode => {
            let value = mode.toLowerCase().replace(/\s/g, "");
            $modes = `${$modes}<option value="${value}">${mode}</option>`;
        });

        this.mediadisplaymode.$modes = $modes;

        $editor.on(
            "click.redactor-mediadisplaymode touchstart.redactor-mediadisplaymode",
            $imgs,
            this.setImageEditter.bind(this)
        );
    },

    insertButton: function ($ele) {
        let self = this;
        if ($ele.is("img")) {
            let $button = $(self.getEditterButton()).prependTo(
                $(".redactor-context-toolbar.open")
            );

            $button.on("click.redactor-mediadisplaymode", function () {
                self.show.bind(self)($button);
            });
        }
    },

    removeButton: function (e) {
        if (!$(e.target).closest("#redactor-image-box").length) {
            $("#redactor-image-displaymode").remove();
            $(document).off("click.redactor-mediadisplaymode");
            $("#modal-media-display-mode").closest("#redactor-modal").remove();
        }
    },

    setImageEditter: function (e) {
        let self = this,
            $image = $(e.target),
            $imageTag = $image.closest("figure");

        this.mediadisplaymode.$imageTag = $imageTag;
        this.mediadisplaymode.$image = $image;
        this.mediadisplaymode.currentDisplaymode = this.getDisplayMode(
            $imageTag.attr("class")
        );

        $(document).on("click.redactor-mediadisplaymode", this.removeButton);

        setTimeout(function () {
            self.insertButton($image);
        }, 10);
    },

    getDisplayMode: function (className) {
        let mode = className ? className.match(/rd-display-\w+/) : null;

        if (mode) {
            mode = mode[0].replace("rd-display-", "");
        }

        return mode;
    },
    modals: {
        mediaModal: `<div class="redactor-modal-tab redactor-group">
    <div id="redactor-image-preview" class="redactor-modal-tab-side"></div>
    <div class="redactor-modal-tab-area" id="redactor-modal-displaymode">
        <section>
            <select id="modal-media-display-mode">
                <option value="0">Please select display mode</option>
            </select>
        </section>
    </div></div>`
    },
    show: function ($button) {
        var options = {
            name: "mediaModal",
            title: "Media Display Mode",
            handle: "save", // optional, command which will be fired on enter pressed
            // optional object
            commands: {
                save: {title: "Save"},
                cancel: {title: "Cancel"}
            }
        };
        this.app.api("module.modal.build", options);
    },

    onmodal: {
        mediaModal: {
            open: function ($modal, $form) {
                let currentDisplaymode = this.mediadisplaymode.currentDisplaymode;

                $("#modal-media-display-mode").append(this.mediadisplaymode.$modes);

                $("#modal-media-display-mode").val(
                    currentDisplaymode ? currentDisplaymode : 0
                );
                $("#redactor-image-preview").html(
                    `<img src="${this.mediadisplaymode.$image.prop(
                        "src"
                    )}" style="max-width: 100%; opacity: 1">`
                );
            },
            save: function ($modal, $form) {
                let displayMode = $("#modal-media-display-mode").val(),
                    $imageTag = this.mediadisplaymode.$imageTag;

                // remove all rd-display-* className
                $imageTag.removeClass(function (index, className) {
                    return (className.match(/rd-display-\w+/g) || []).join(" ");
                });

                if (displayMode != 0) {
                    $imageTag.addClass(`rd-display-${displayMode}`);
                }
                this.app.api("module.modal.close");
            }
        }
    }
});
