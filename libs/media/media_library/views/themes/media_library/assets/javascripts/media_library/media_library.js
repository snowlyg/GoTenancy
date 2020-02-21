(function (factory) {
    if (typeof define === "function" && define.amd) {
        // AMD. Register as anonymous module.
        define(["jquery"], factory);
    } else if (typeof exports === "object") {
        // Node / CommonJS
        factory(require("jquery"));
    } else {
        // Browser globals.
        factory(jQuery);
    }
})(function ($) {
    "use strict";

    var NAMESPACE = "qor.medialibrary.action";
    var EVENT_ENABLE = "enable." + NAMESPACE;
    var EVENT_DISABLE = "disable." + NAMESPACE;
    var EVENT_KEYUP = "keyup." + NAMESPACE;
    var EVENT_SWITCHED = "switched.qor.tabbar.radio";
    var EVENT_SWITCHED_TARGET = '[data-toggle="qor.tab.radio"]';
    var EVENT_BOTTOMSHEETS_RELOAD = "reload.qor.bottomsheets bottomsheetLoaded.qor.bottomsheets";
    var CLASS_MEDIA_DATA = '[name="QorResource.SelectedType"]';
    var CLASS_VIDEO_TAB = '[data-tab-source="video_link"]';
    var CLASS_VIDEO = ".qor-video__link";
    var CLASS_VIDEO_TABLE = ".qor-medialibrary__video-link";
    var CLASS_UPLOAD_VIDEO_TABLE = ".qor-medialibrary__video";
    var CLASS_IMAGE_DESC = ".qor-medialibrary__desc";
    var CLASS_FILE_OPTION = ".qor-file__options";
    var CLASS_MEDIABOX = ".qor-bottomsheets__mediabox";
    var CLASS_MEDIA_OPTION = 'input[name="QorResource.MediaOption"]';

    function getYoutubeID(url) {
        var regExp = /^.*((youtu.be\/)|(v\/)|(\/u\/\w\/)|(embed\/)|(watch\?))\??v?=?([^#\&\?]*).*/;
        var match = url.match(regExp);
        if (match && match[7].length == 11) {
            return match[7];
        } else {
            return false;
        }
    }

    function getYoukuID(url) {
        /******

         // <iframe height=498 width=510 src='http://player.youku.com/embed/XMTM1NzQ0NTQ4' frameborder=0 'allowfullscreen'></iframe>
         // http://player.youku.com/player.php/sid/XMTM1NzQ0NTQ4/v.swf
         // <embed src='http://player.youku.com/player.php/sid/XMTM1NzQ0NTQ4/v.swf' allowFullScreen='true' quality='high' width='480' height='400' align='middle' allowScriptAccess='always' type='application/x-shockwave-flash'></embed>
         // http://v.youku.com/v_show/id_XMTc4MjU2NTk4OA.html

         *****/

        var regExp = /(\/id_)(\w+)/,
            regYouku = /http?:\/\/(www\.)|(v\.)youku.com/,
            match = url.match(regExp);
        if (regYouku.test(url) && match && match[2]) {
            return match[2];
        } else {
            return false;
        }
    }

    function QorMedialibraryAction(element, options) {
        this.$element = $(element);
        this.options = $.extend(
            {},
            QorMedialibraryAction.DEFAULTS,
            $.isPlainObject(options) && options
        );
        this.init();
    }

    QorMedialibraryAction.prototype = {
        constructor: QorMedialibraryAction,

        init: function () {
            this.bind();
            this.initMedia();
            $.fn.qorSliderAfterShow.renderMediaVideo();
        },

        bind: function () {
            $(document)
                .on(
                    EVENT_SWITCHED,
                    EVENT_SWITCHED_TARGET,
                    this.resetMediaData.bind(this)
                )
                .on(EVENT_KEYUP, CLASS_VIDEO, this.setVideo.bind(this))
                .on(EVENT_KEYUP, CLASS_IMAGE_DESC, this.setImageDesc.bind(this))
                .on(
                    EVENT_BOTTOMSHEETS_RELOAD,
                    CLASS_MEDIABOX,
                    this.initMedia.bind(this, "bottomsheet")
                );
        },

        unbind: function () {
            $(document)
                .off(
                    EVENT_SWITCHED,
                    EVENT_SWITCHED_TARGET,
                    this.resetMediaData.bind(this)
                )
                .off(EVENT_KEYUP, CLASS_VIDEO, this.setVideo.bind(this))
                .off(EVENT_KEYUP, CLASS_IMAGE_DESC, this.setImageDesc.bind(this));
        },

        setMediaData: function ($form, value) {
            var $fileOption = $form.find(CLASS_FILE_OPTION),
                $MediaOption = $form.find(CLASS_MEDIA_OPTION);

            $fileOption.val(JSON.stringify(value));
            $MediaOption.val(JSON.stringify(value));
        },

        setImageDesc: function (e) {
            var $input = $(e.target),
                $form = $input.closest("form"),
                $fileOption,
                fileOption;

            $fileOption = $form.find(CLASS_FILE_OPTION);
            fileOption = JSON.parse($fileOption.val());
            fileOption.Description = $input.val();

            this.setMediaData($form, fileOption);
        },

        initMedia: function (bottomsheet) {
            var $uploadVideo = $(CLASS_UPLOAD_VIDEO_TABLE),
                $linkedvideo = $(CLASS_VIDEO_TABLE);

            if (bottomsheet) {
                $uploadVideo = $(CLASS_MEDIABOX).find(CLASS_UPLOAD_VIDEO_TABLE);
                $linkedvideo = $(CLASS_MEDIABOX).find(CLASS_VIDEO_TABLE);
            }

            $(CLASS_MEDIABOX)
                .find(".qor-table--medialibrary-file")
                .each(function () {
                    $(this)
                        .closest(".mdl-card__supporting-text")
                        .addClass("qor-table--files");
                });

            if (!$uploadVideo.length && !$linkedvideo.length) {
                return;
            }

            $uploadVideo.each(function () {
                var $this = $(this),
                    url = $this.data("videolink"),
                    videoType =
                        url &&
                        url.match(
                            /\.mp4$|\.m4p$|\.m4v$|\.m4v$|\.mov$|\.mpeg$|\.webm$|\.avi$|\.ogg$|\.ogv$/
                        );

                if (videoType) {
                    $this
                        .parent()
                        .addClass("qor-table--video qor-table--video-internal")
                        .html(
                            '<video width=100% height=100% controls><source src="' +
                            url +
                            '"></video>'
                        );
                }
            });

            $linkedvideo.each(function () {
                var $this = $(this),
                    url = $this.data("videolink"),
                    youtubeID = getYoutubeID(url),
                    youkuID = getYoukuID(url);

                if (youtubeID) {
                    $this
                        .parent()
                        .addClass("qor-table--video qor-table--video-external")
                        .html(
                            '<iframe width="100%" height="100%" src="//www.youtube.com/embed/' +
                            youtubeID +
                            '?rel=0" frameborder="0" allowtransparency="true" allowfullscreen="true"></iframe>'
                        );
                } else if (youkuID) {
                    $this
                        .parent()
                        .addClass("qor-table--video qor-table--video-external")
                        .html(
                            '<iframe width=100% height=100% src="http://player.youku.com/embed/' +
                            youkuID +
                            '" frameborder=0 allowtransparency="true" allowfullscreen="true"></iframe>'
                        );
                } else {
                    $this
                        .parent()
                        .addClass("qor-table--video qor-table--video-external")
                        .html(
                            '<iframe width=100% height=100% src="' +
                            url +
                            '" frameborder=0 allowtransparency="true" allowfullscreen="true"></iframe>'
                        );
                }
            });
        },

        setVideo: function (event) {
            var $input = $(event.target),
                $parent = $input.closest("[data-tab-source]"),
                $form = $input.closest("form"),
                $fileOption = $form.find(CLASS_FILE_OPTION),
                fileOption = JSON.parse($fileOption.val()),
                url = $input.val(),
                $iframe = $parent.find("iframe"),
                youtubeID = getYoutubeID(url),
                youkuID = getYoukuID(url);

            fileOption.SelectedType = "video_link";
            fileOption.Video = url;
            fileOption.Description = $parent.find(CLASS_IMAGE_DESC).val();

            this.setMediaData($form, fileOption);

            $parent.find("iframe").remove();

            if (youtubeID || youkuID) {
                $iframe.length && $iframe.remove();
                if (youtubeID) {
                    $parent.append(
                        '<iframe width="100%" height="400" src="//www.youtube.com/embed/' +
                        youtubeID +
                        '?rel=0" frameborder="0" allowtransparency="true" allowfullscreen="true"></iframe>'
                    );
                }
                if (youkuID) {
                    $parent.append(
                        '<iframe width=100% height=400 src="http://player.youku.com/embed/' +
                        youkuID +
                        '" frameborder=0 allowtransparency="true" allowfullscreen="true"></iframe>'
                    );
                }
            } else {
                $parent.append(
                    '<iframe width=100% height=400 src="' +
                    url +
                    '" frameborder=0  allowtransparency="true" allowfullscreen="true"></iframe>'
                );
            }
        },

        resetMediaData: function (e, element, type) {
            var $element = $(element),
                $form = $element.closest("form"),
                $fileOption = $element.find(CLASS_FILE_OPTION),
                $alert = $element.find(CLASS_VIDEO_TAB).find(".qor-fieldset__alert"),
                fileOption = JSON.parse($fileOption.val());

            fileOption.SelectedType = type;

            if (type == "video_link") {
                fileOption.Video = $element.find(CLASS_VIDEO).val();
                $alert.length && $alert.remove();
            }

            fileOption.Description = $('[data-tab-source="' + type + '"]')
                .find(CLASS_IMAGE_DESC)
                .val();
            $(CLASS_MEDIA_DATA).val(type);

            this.setMediaData($form, fileOption);
        },

        destroy: function () {
            this.unbind();
        }
    };

    QorMedialibraryAction.DEFAULTS = {};

    $.fn.qorSliderAfterShow = $.fn.qorSliderAfterShow || {};
    $.fn.qorSliderAfterShow.renderMediaVideo = function () {
        var $render = $(CLASS_VIDEO_TAB),
            $desc = $(CLASS_IMAGE_DESC),
            url = $render.length && $render.data().videourl,
            youtubeID = url && getYoutubeID(url),
            youkuID = url && getYoukuID(url);

        $desc.length && $desc.val($desc.data().imageInfo.Description);

        if ($render.length && url) {
            if (youtubeID) {
                $render.append(
                    '<iframe width="100%" height="400" src="//www.youtube.com/embed/' +
                    youtubeID +
                    '?rel=0&fs=0&modestbranding=1&disablekb=1" frameborder="0" allowtransparency="true" allowfullscreen="true"></iframe>'
                );
            } else if (youkuID) {
                $render.append(
                    '<iframe width=100% height=400 src="http://player.youku.com/embed/' +
                    youkuID +
                    '" frameborder=0 allowtransparency="true" allowfullscreen="true"></iframe>'
                );
            } else {
                $render.append(
                    '<iframe width=100% height=400 src="' +
                    url +
                    '" frameborder=0 allowtransparency="true" allowfullscreen="true"></iframe>'
                );
            }
        }
    };

    QorMedialibraryAction.plugin = function (options) {
        return this.each(function () {
            var $this = $(this);
            var data = $this.data(NAMESPACE);
            var fn;

            if (!data) {
                if (/destroy/.test(options)) {
                    return;
                }

                $this.data(
                    NAMESPACE,
                    (data = new QorMedialibraryAction(this, options))
                );
            }

            if (typeof options === "string" && $.isFunction((fn = data[options]))) {
                fn.apply(data);
            }
        });
    };

    $(function () {
        var selector = ".qor-theme-media_library";

        $(document)
            .on(EVENT_DISABLE, function (e) {
                QorMedialibraryAction.plugin.call($(selector, e.target), "destroy");
            })
            .on(EVENT_ENABLE, function (e) {
                QorMedialibraryAction.plugin.call($(selector, e.target));
            })
            .triggerHandler(EVENT_ENABLE);
    });

    return QorMedialibraryAction;
});
