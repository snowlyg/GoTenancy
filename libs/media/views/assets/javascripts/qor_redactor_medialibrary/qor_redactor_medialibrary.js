// Add media library button for redactor editor
// By Jason weng @theplant
//

$R.add("plugin", "medialibrary", {
    init: function (app) {
        this.app = app;
        this.opts = app.opts;
        this.lang = app.lang;
        this.inline = app.inline;
        this.toolbar = app.toolbar;
        this.func = new this.funcInit(this);
    },

    start: function () {
        var $button = this.toolbar.addButton("medialibrary", {
            title: "MediaLibrary",
            api: "plugin.medialibrary.addMedialibrary"
        });
        $button.setIcon('<i class="material-icons">photo_library</i>');
        this.buttonElement = $button.nodes[0];

        this.$currentTag = false;

        $(document).on(
            "reload.qor.bottomsheets",
            ".qor-bottomsheets__mediabox",
            this.func.initItem
        );
    },

    addMedialibrary: function () {
        this.$currentTag = this.app.selection.getCurrent();
        this.func.addMedialibrary();
    },

    funcInit: function (self) {
        var thisApp = self;
        var thisFunc = this;
        this.addMedialibrary = function () {
            var $element = $(thisApp.app.rootElement),
                data = {selectModal: "mediabox", maxItem: "1"},
                mediaboxUrl = $element.data().redactorSettings.medialibraryUrl,
                BottomSheets;

            thisFunc.BottomSheets = BottomSheets = $("body").data("qor.bottomsheets");
            data.url = mediaboxUrl;
            BottomSheets.open(data, function ($bottomsheets) {
                thisFunc.handleMediaLibrary($bottomsheets);
            });
        };

        this.handleMediaLibrary = function ($bottomsheets) {
            let options = {
                onSelect: thisFunc.handleResults, // render selected item after click item lists
                onSubmit: thisFunc.handleResults // render new items after new item form submitted
            };

            thisFunc.$bottomsheets = $bottomsheets;
            $bottomsheets
                .qorSelectCore(options)
                .addClass("qor-bottomsheets__mediabox");
            thisFunc.initItem();
        };

        this.initItem = function () {
            var $trs = $(".qor-bottomsheets").find("tbody tr"),
                $tr,
                $img;

            $trs.each(function () {
                $tr = $(this);
                $img = $tr.find(".qor-table--ml-slideout p img").first();
                $tr.find(".qor-table__actions").remove();
                if ($img.length) {
                    $tr
                        .find(".qor-table--medialibrary-item")
                        .css("background-image", "url(" + $img.prop("src") + ")");
                    $img.parent().remove();
                }
            });
        };

        this.handleResults = function (data) {
            var reVideo = /\.mp4$|\.m4p$|\.m4v$|\.m4v$|\.mov$|\.mpeg$|\.webm$|\.avi$|\.ogg$|\.ogv$/,
                mediaOption = data.MediaOption;

            if (
                data.SelectedType == "video_link" ||
                mediaOption.Video ||
                mediaOption.URL.match(reVideo)
            ) {
                thisFunc.insertVideo(data);
            } else {
                thisFunc.insertImage(data);
            }

            thisFunc.$bottomsheets.remove();
            if (!$(".qor-bottomsheets").is(":visible")) {
                $("body").removeClass("qor-bottomsheets-open");
            }
        };

        this.insertVideo = function (data) {
            const $editor = $(thisApp.app.editor.$editor.nodes[0]);
            const $rootElement = $(thisApp.app.rootElement);

            thisApp.opts.mediaContainerClass =
                typeof thisApp.opts.mediaContainerClass === "undefined"
                    ? "qor-video-container"
                    : thisApp.opts.mediaContainerClass;

            var htmlCode,
                videoLink,
                iframeStart,
                iframeEnd,
                videoType,
                $html,
                youkuID,
                callbackData = {},
                mediaContainerClass = thisApp.opts.mediaContainerClass,
                reUrlYoutube = thisApp.opts.regex.youtube,
                reUrlVimeo = thisApp.opts.regex.vimeo,
                reUrlYouku = /http?:\/\/(www\.)|(v\.)youku.com/,
                reUrlYoukuID = /(\/id_)(\w+)/,
                reVideo = /\.mp4$|\.m4p$|\.m4v$|\.m4v$|\.mov$|\.mpeg$|\.webm$|\.avi$|\.ogg$|\.ogv$/,
                randomString = (Math.random() + 1).toString(36).substring(7),
                videoIdentification = "qor-video-" + randomString,
                mediaOption = data.MediaOption,
                description = mediaOption.Description;

            iframeStart = `<figure class="${mediaContainerClass} video-scale"><iframe title="${description}" data-media-id="${data.ID ||
            data.primaryKey}" width="100%" height="380px" src="`;

            iframeEnd = '" frameborder="0" allowfullscreen="true"></iframe></figure>';

            if (data.SelectedType == "video_link") {
                videoLink = mediaOption.Video;

                if (videoLink.match(reUrlYoutube)) {
                    videoType = "youtube";
                    htmlCode = videoLink.replace(
                        reUrlYoutube,
                        iframeStart + "//www.youtube.com/embed/$1" + iframeEnd
                    );
                } else if (videoLink.match(reUrlVimeo)) {
                    videoType = "vimeo";
                    htmlCode = videoLink.replace(
                        reUrlVimeo,
                        iframeStart + "//player.vimeo.com/video/$2" + iframeEnd
                    );
                } else if (
                    videoLink.match(reUrlYouku) &&
                    reUrlYoukuID.test(videoLink)
                ) {
                    videoType = "youku";
                    youkuID = videoLink.match(reUrlYoukuID)[2];
                    htmlCode = `<div class="video-scale"><iframe width=100% height=400 data-media-id="${data.ID ||
                    data.primaryKey}" src="http://player.youku.com/embed/${youkuID}" frameborder=0 allowfullscreen="true"></iframe></div>`;
                } else {
                    videoType = "others";
                    htmlCode = `<div class="video-scale"><iframe data-media-id="${data.ID ||
                    data.primaryKey}" width=100% height=400 src="${videoLink}" frameborder=0 allowfullscreen="true"></iframe></div>`;
                }
            } else if (mediaOption.URL.match(reVideo)) {
                videoType = "uploadedVideo";
                htmlCode =
                    '<figure class="' +
                    mediaContainerClass +
                    '"><div class="video-scale" role="application"><video width="100%" title="' +
                    description +
                    '" aria-label="' +
                    description +
                    '" height="380px" controls="controls" aria-describedby="' +
                    videoIdentification +
                    '" tabindex="0"><source src="' +
                    mediaOption.URL +
                    '"></video></div></figure>';
            }

            if (!htmlCode) {
                return;
            }

            $html = $(htmlCode).addClass(videoIdentification);

            if (thisApp.$currentTag) {
                $(thisApp.$currentTag).after($html);
            } else {
                $editor.prepend($html);
            }

            // trigger insertedVideo.redactor event after inserted videos
            callbackData.type = videoType;
            callbackData.videoLink = videoLink || mediaOption.URL;
            callbackData.videoIdentification = videoIdentification;
            callbackData.description = description;
            callbackData.$editor = $editor;
            callbackData.$element = $rootElement;

            $rootElement.trigger("insertedVideo.redactor", [callbackData]);
        };

        this.insertImage = function (data) {
            const $editor = $(thisApp.app.editor.$editor.nodes[0]);
            const $rootElement = $(thisApp.app.rootElement);

            var src,
                $img = $("<img>"),
                $figure = $("<figure>"),
                mediaOption = data.MediaOption,
                callbackData = {};

            var imageFile = data.File && JSON.parse(data.File);

            src = mediaOption.URL.replace(/image\..+\./, "image.");

            $img.attr({
                src: src,
                alt: mediaOption.Description || imageFile.Description
            });
            $figure.append($img);

            if (thisApp.$currentTag) {
                $(thisApp.$currentTag).after($figure);
            } else {
                $editor.prepend($figure);
            }

            // set img editable
            $figure.addClass("redactor-component");
            $figure.attr({
                "data-redactor-type": "image",
                tabindex: "-1",
                contenteditable: false
            });

            // trigger insertedVideo.redactor event after inserted images
            callbackData.description = mediaOption.Description;
            callbackData.$img = $figure;
            callbackData.$editor = $editor;
            callbackData.$element = $rootElement;

            $rootElement.trigger("insertedImage.redactor", [callbackData]);
        };
    }
});
