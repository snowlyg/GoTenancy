(function (factory) {
    if (typeof define === 'function' && define.amd) {
        // AMD. Register as anonymous module.
        define(['jquery'], factory);
    } else if (typeof exports === 'object') {
        // Node / CommonJS
        factory(require('jquery'));
    } else {
        // Browser globals.
        factory(jQuery);
    }
})(function ($) {
    'use strict';

    let _ = window._,
        NAMESPACE = 'qor.bannereditor',
        EVENT_ENABLE = 'enable.' + NAMESPACE,
        EVENT_DISABLE = 'disable.' + NAMESPACE,
        EVENT_CLICK = 'click.' + NAMESPACE,
        EVENT_CHANGE = 'change.' + NAMESPACE,
        EVENT_DBCLICK = 'dblclick.' + NAMESPACE,
        EVENT_DRAGSTOP = 'dragstop.' + NAMESPACE,
        EVENT_RESIZESTOP = 'resizestop.' + NAMESPACE,
        EVENT_DRAG = 'drag.' + NAMESPACE,
        EVENT_SHOWN = 'shown.qor.modal',
        EVENT_HIDDEN = 'hidden.qor.modal',
        CLASS_DRAGGABLE = '.qor-bannereditor__draggable',
        CLASS_MEDIABOX = 'qor-bottomsheets__mediabox',
        CLASS_TOOLBAR_BUTTON = '.qor-bannereditor__button',
        CLASS_BANNEREDITOR_VALUE = '.qor-bannereditor__value',
        CLASS_BANNEREDITOR_IMAGE = '.qor-bannereditor__toolbar-image',
        CLASS_BANNEREDITOR_DRAGGING = 'qor-bannereditor__dragging',
        CLASS_BANNEREDITOR_CONTENT = '.qor-bannereditor__contents',
        CLASS_CANVAS = '.qor-bannereditor__canvas',
        CLASS_PLATFORM_TRIGGER = '.qor-bannereditor__platform-trigger',
        CLASS_PLATFORM_PANEL = '.qor-bannereditor__platform-panel',
        CLASS_CONTAINER = '.qor-bannereditor__container',
        CLASS_DEVICE_SELECTOR = '.qor-bannereditor__device',
        CLASS_DEVICE_TOOLBAR = '.qor-bannereditor__device-toolbar',
        CLASS_DEVICE_MODE = 'qor-bannereditor__device-mode',
        CLASS_TOP = 'qor-bannereditor__draggable-top',
        CLASS_LEFT = 'qor-bannereditor__draggable-left',
        CLASS_BG_IMAGE = '.qor-bannereditor-image',
        CLASS_CROP_WRAP = '.qor-cropper__wrapper',
        CLASS_NEED_REMOVE = '.qor-bannereditor__button-inline,.ui-resizable-handle,.qor-bannereditor__draggable-coordinate,.ui-draggable-handle,.ui-resizable',
        CLASS_ACTIVE = 'is-active';

    function getImgSize(url, callback) {
        let img = new Image();

        img.onload = function () {
            if ($.isFunction(callback)) {
                callback(this.naturalWidth || this.width, this.naturalHeight || this.height);
            }
        };
        img.src = url;
    }

    function getDeviceSize(size, sym) {
        let arr = size.split(sym || 'x');
        return {
            width: parseInt(arr[0]),
            height: parseInt(arr[1])
        };
    }

    function getObject(value, platformName) {
        let bannerValue = _.where(value, {Name: platformName})[0];
        return bannerValue ? bannerValue : false;
    }

    function generateRandomString() {
        return (Math.random() + 1).toString(36).substring(7);
    }

    function replaceText(str, data) {
        if (typeof str === 'string') {
            if (typeof data === 'object') {
                $.each(data, function (key, val) {
                    str = str.replace('$[' + String(key).toLowerCase() + ']', val);
                });
            }
        }

        return str;
    }

    function QorBannerEditor(element, options) {
        this.$element = $(element);
        this.options = $.extend({}, QorBannerEditor.DEFAULTS, $.isPlainObject(options) && options);
        this.init();
    }

    QorBannerEditor.prototype = {
        constructor: QorBannerEditor,

        init: function () {
            let $element = this.$element,
                $container = $element.closest(CLASS_CONTAINER),
                $textarea = $container.find(CLASS_BANNEREDITOR_VALUE),
                config = {},
                _this = this,
                $canvas = $element.find(CLASS_CANVAS),
                platformName = $element.data('platform-name'),
                configure = $textarea.data('configure'),
                bannerValues = JSON.parse($textarea.val()),
                platforms = configure.Platforms,
                html,
                $iframe = $('<iframe width="100%" height="300px" />'),
                bannerSizes = getObject(platforms, platformName),
                currentBannerValue = getObject(bannerValues, platformName) || {Value: ''};

            html = $(`<div class="qor-bannereditor__canvas">${decodeURIComponent(currentBannerValue.Value)}</div>`);

            config.toolbar = configure.Elements;
            config.editURL = configure.EditURL;
            config.externalStylePath = configure.ExternalStylePath;
            config.Platforms = platforms;

            this.config = config;
            this.$textarea = $textarea;
            this.$container = $container;
            this.$content = $element.find(CLASS_BANNEREDITOR_CONTENT);

            this.initWidth = bannerSizes.Width || '';
            this.initHeight = bannerSizes.Height || '';

            $element
                .find('.qor-bannereditor__toolbar--size')
                .show()
                .find('span')
                .html(`${this.initWidth || '100%'} X ${this.initHeight || '100%'}`);

            $canvas.html($iframe).removeClass('qor-bannereditor__canvas');

            this.$iframe = $iframe;
            this.platformName = platformName;

            if ($('.qor-slideout').is(':visible')) {
                // for sliderout
                $iframe.ready(function () {
                    setTimeout(function () {
                        _this.initIframe(html);
                    }, 500);
                });
            } else {
                // for single page
                $iframe.on('load', function () {
                    setTimeout(function () {
                        _this.initIframe(html);
                    }, 500);
                });
            }
        },

        initIframe: function (html) {
            let $ele = this.$iframe.contents(),
                $head = $ele.find('head'),
                $bannerHtml,
                $canvas,
                externalStylePath = this.config.externalStylePath,
                defaultCSS = this.$element.closest(CLASS_CONTAINER).data('prefix') + '/assets/stylesheets/banner_editor_iframe.css?theme=banner_editor',
                linkTemplate = function (url) {
                    return `<link rel="stylesheet" type="text/css" href="${url}">`;
                };

            $head.append(linkTemplate(defaultCSS));

            // load banner editor external style
            if (externalStylePath && externalStylePath.length > 0) {
                for (let i = 0; i < externalStylePath.length; i++) {
                    $head.append(linkTemplate(externalStylePath[i]));
                }
            }

            $ele.find('body').html(html);
            $canvas = $ele.find(CLASS_CANVAS);

            $bannerHtml = $ele.find('.qor-bannereditor__html');

            if (!$bannerHtml.length) {
                $bannerHtml = $('<div class="qor-bannereditor__html" style="position: relative; height: 100%;" />').appendTo($canvas);
            }

            this.$canvas = $canvas;
            this.$bannerHtml = $bannerHtml;

            this.initBannerEditor();
            this.bind();
            if (this.platformName === 'Mobile') {
                this.$element.find(CLASS_DEVICE_SELECTOR).trigger('change');
                this.$canvas.addClass(CLASS_DEVICE_MODE);
            }
        },

        initBannerEditor: function () {
            let $toolbar,
                $element = this.$element,
                $canvas = this.$canvas,
                $iframe = this.$iframe,
                replaceTexts = {},
                eleData = $element.data(),
                canvasWidth = this.initWidth || this.$bannerHtml.data('image-width'),
                canvasHeight = this.initHeight || this.$bannerHtml.data('image-height'),
                isInBottomsheet = $element.closest('.qor-bottomsheets').length,
                isInSlideout = $('.qor-slideout').is(':visible'),
                hasFullClass = $('.qor-slideout').hasClass('qor-slideout__fullscreen');

            this.config.toolbar.forEach(function (obj) {
                obj.id = generateRandomString();
            });
            $toolbar = $(window.Mustache.render(QorBannerEditor.toolbar, this.config));
            $toolbar.appendTo($element.find('.qor-bannereditor__toolbar-btns'));
            this.$popover = $(QorBannerEditor.popover).appendTo('body');

            replaceTexts = {
                title: eleData.cropperTitle,
                ok: eleData.cropperOk,
                cancel: eleData.cropperCancel
            };

            this.$cropModal = $(replaceText(QorBannerEditor.modal, replaceTexts)).appendTo('body');

            $element.closest('.qor-fieldset').addClass('qor-fieldset-bannereditor');
            this.resetToolbarTooltips();
            $element.find('.qor-bannereditor__toolbar').trigger('enable');

            if (isInSlideout && !isInBottomsheet && !hasFullClass) {
                $('.qor-slideout__fullscreen').click();
            }

            if (isInBottomsheet) {
                $element.closest('.qor-bottomsheets').addClass('qor-bottomsheets__fullscreen');
            }

            $canvas.width(canvasWidth).height(canvasHeight);
            $iframe.height(canvasHeight);

            $element.find('.qor-bannereditor__contents').show();
        },

        resetToolbarTooltips: function () {
            let $button = this.$element.find('.qor-bannereditor__toolbar--ml'),
                $innerButtons = $button.find(' > button'),
                $innerTip = $button.find('.mdl-tooltip'),
                $all = $button.find(' > button, .mdl-tooltip'),
                randomString = generateRandomString();

            $all.removeAttr('data-upgraded');
            $innerButtons.attr('id', randomString);
            $innerTip.attr('data-mdl-for', randomString);
        },

        bind: function () {
            let $canvas = this.$canvas;

            this.$container.on(EVENT_CLICK, CLASS_PLATFORM_TRIGGER, this.switchPlatform.bind(this));

            this.$cropModal.on(EVENT_SHOWN, this.initCrop.bind(this)).on(EVENT_HIDDEN, this.stopCrop.bind(this));

            this.$element
                .on(EVENT_CLICK, CLASS_TOOLBAR_BUTTON, this.addElements.bind(this))
                .on(EVENT_CLICK, '.qor-bannereditor__toolbar-clear', this.clearAllElements.bind(this))
                .on(EVENT_CLICK, CLASS_BANNEREDITOR_IMAGE, this.openBottomSheets.bind(this))
                .on(EVENT_CHANGE, CLASS_DEVICE_SELECTOR, this.switchDevice.bind(this));

            $canvas
                .on(EVENT_CLICK, '.qor-bannereditor__editimage', this.replaceBackgroundImage.bind(this))
                .on(EVENT_CLICK, '.qor-bannereditor__deleteimage', this.deleteBackgroundImage.bind(this))
                .on(EVENT_CLICK, '.qor-bannereditor__cropimage', this.cropBackgroundImage.bind(this))
                .on(EVENT_CLICK, CLASS_BG_IMAGE, this.editBackground.bind(this))
                .on(EVENT_CLICK, CLASS_DRAGGABLE, this.handleInlineEdit.bind(this))
                .on(EVENT_DBCLICK, CLASS_DRAGGABLE, this.showInlineEdit.bind(this))
                .on(EVENT_CLICK, '.qor-bannereditor__button-inline button', this.showEdit.bind(this))
                .on(EVENT_DRAGSTOP, CLASS_DRAGGABLE, this.handleDragStop.bind(this))
                .on(EVENT_RESIZESTOP, CLASS_DRAGGABLE, this.handleResizeStop.bind(this))
                .on(EVENT_DRAG, CLASS_DRAGGABLE, this.handleDrag.bind(this));

            $canvas
                .find(CLASS_DRAGGABLE)
                .draggable(this.options.draggable)
                .resizable(this.options.resizable);

            $(document).on(EVENT_CLICK, this.hideElement.bind(this));
        },

        unbind: function () {
            let $canvas = this.$canvas;
            this.$container.off(EVENT_CLICK);
            this.$element.off(EVENT_CLICK).off(EVENT_CHANGE);
            $canvas
                .off(EVENT_CLICK)
                .off(EVENT_DBCLICK)
                .off(EVENT_DRAGSTOP)
                .off(EVENT_RESIZESTOP)
                .off(EVENT_DRAG);
            $canvas
                .find(CLASS_DRAGGABLE)
                .draggable('destroy')
                .resizable('destroy');
            $(document).off(EVENT_CLICK, this.hideElement);
        },

        clearAllElements: function (e) {
            let $target = $(e.target),
                _this = this,
                message = {
                    confirm: $target.data('hint-message')
                };

            window.QOR.qorConfirm(message, function (confirm) {
                if (confirm) {
                    _this.$bannerHtml.html('');
                    _this.setValue();
                }
            });
        },

        switchPlatform: function (e) {
            let $container = this.$container,
                $target = $(e.target),
                id = $target.attr('name'),
                $content = $container.find(id),
                $element = $content.find('.qor-bannereditor');

            $container.find(CLASS_PLATFORM_TRIGGER).removeClass(CLASS_ACTIVE);
            $container.find(CLASS_PLATFORM_PANEL).hide();
            $content.show();
            $target.addClass(CLASS_ACTIVE);

            if (!$element.data(NAMESPACE)) {
                $content.trigger('enable');
            }

            return false;
        },

        resetDevice: function () {
            let initWidth = this.initWidth || 'auto',
                initHeight = this.initHeight || 300;

            this.$content.css('width', 'auto');
            this.$iframe.css({
                width: '100%',
                height: initHeight
            });
            this.$canvas.css({
                width: initWidth,
                height: initHeight
            });
        },

        switchDevice: function (e) {
            this.resetBannerEditorSize($(e.target).val());
        },

        resetBannerEditorSize: function (size) {
            let $element = this.$element,
                deviceSize = getDeviceSize(size),
                deviceWidth = deviceSize.width,
                deviceHeight = deviceSize.height;

            this.$iframe.css({
                width: deviceWidth,
                height: deviceHeight
            });
            this.$canvas.css({
                width: deviceWidth,
                height: this.initHeight || deviceHeight
            });

            this.$content.width(deviceWidth);
            $element.find(CLASS_DEVICE_TOOLBAR).width(deviceWidth);
        },

        showCoordinate: function () {
            let $target = this.$canvas.find('.qor-bannereditor__dragging'),
                position = {};

            (position.left = parseInt($target.attr('data-position-left'), 10)), (position.top = parseInt($target.attr('data-position-top'), 10));

            this.$canvas.find('.qor-bannereditor__draggable-coordinate').remove();
            $target.append(window.Mustache.render(QorBannerEditor.dragCoordinate, position));
            $target.find('.qor-bannereditor__button-inline').hide();
            this.setValue();
        },

        initMedia: function ($bottomsheets) {
            let $trs = $bottomsheets.find('tbody tr'),
                $tr,
                $img;

            $trs.each(function () {
                $tr = $(this);
                $img = $tr.find('.qor-table--ml-slideout p img').first();
                $tr.find('.qor-table__actions').remove();

                if ($img.length) {
                    $tr.find('.qor-table--medialibrary-item').css('background-image', 'url(' + $img.prop('src') + ')');
                    $img.parent().remove();
                }
            });
        },

        replaceBackgroundImage: function () {
            this.$element.find(CLASS_BANNEREDITOR_IMAGE).click();
            this.clearElements();
        },

        deleteBackgroundImage: function () {
            this.$canvas.find(CLASS_BG_IMAGE).remove();
            this.clearElements();
            this.setValue();
        },

        editBackground: function (e) {
            let $target = $(e.target).closest(CLASS_BG_IMAGE),
                $editHTML = $(QorBannerEditor.imageEdit);

            this.clearElements();

            if ($target.data('medialibrary-url')) {
                $editHTML.prepend('<button class="mdl-button mdl-button--icon qor-bannereditor__cropimage" type="button"><i class="material-icons">crop</i></button>');
            }

            $target.append($editHTML);
        },

        cropBackgroundImage: function (e) {
            let $target = $(e.target).closest(CLASS_BG_IMAGE),
                data = $target.data(),
                src = $target.find('img').attr('src'),
                originalUrl = src.replace(/file.bannereditor\_\w+/, 'file.original');

            this.cropImageData = {
                $target: $target,
                url: src,
                originalUrl: originalUrl,
                sizeName: data.sizeName,
                medialibraryUrl: data.medialibraryUrl
            };

            this.$cropModal.qorModal('show');
        },

        initCrop: function () {
            let src = this.cropImageData.originalUrl,
                $img = $(`<img src="${src}">`);

            this.$cropModal.find(CLASS_CROP_WRAP).html($img);
            this.$cropImage = $img;
            getImgSize(src, this.startCrop.bind(this));
        },

        startCrop: function (naturalWidth, naturalHeight) {
            let initWidth = this.initWidth || 0,
                initHeight = this.initHeight || 0,
                hasSize = initWidth && initHeight, //have width and height for banner editor setting
                sizeAspectRatio = hasSize ? initWidth / initHeight : NaN,
                $cropModal = this.$cropModal,
                $cropImage = this.$cropImage,
                cropImageData = this.cropImageData,
                $bannerHtml = this.$bannerHtml,
                $canvas = this.$canvas,
                $iframe = this.$iframe,
                $loading = $('<div class="qor-modal-loading"><div class="mdl-spinner mdl-spinner--single-color mdl-js-spinner is-active"></div></div>'),
                initCropOption = cropImageData.$target.data('crop-options'),
                sizeName = cropImageData.sizeName,
                _this = this,
                cropData = {
                    Sizes: {},
                    CropOptions: {},
                    URL: cropImageData.originalUrl.replace(/file\.original\./, 'file.'),
                    Crop: true
                },
                newCropData = {
                    x: Math.round((naturalWidth - initWidth) / 2),
                    y: Math.round((naturalHeight - initHeight) / 2),
                    width: Math.round(initWidth || naturalWidth),
                    height: Math.round(initHeight || naturalHeight)
                },
                emulateCropData = initCropOption ? initCropOption : newCropData;

            $cropImage.cropper({
                aspectRatio: sizeAspectRatio,
                background: false,
                data: emulateCropData,
                movable: false,
                zoomable: false,
                scalable: false,
                rotatable: false,
                autoCropArea: 1,

                ready: function () {
                    $cropModal.find('.qor-cropper__save').one(EVENT_CLICK, function () {
                        let syncData = {},
                            cropOption = $cropImage.cropper('getData', true),
                            cropOptionWidth = cropOption.width,
                            cropOptionHeight = cropOption.height;

                        if (!hasSize) {
                            cropData['Sizes'][sizeName] = {
                                Width: cropOptionWidth,
                                Height: cropOptionHeight
                            };
                        }

                        cropData['CropOptions'][sizeName] = cropOption;
                        syncData.MediaOption = JSON.stringify(cropData);

                        $loading.appendTo($cropModal.find(CLASS_CROP_WRAP)).trigger('enable.qor.material');

                        $.ajax(cropImageData.medialibraryUrl, {
                            type: 'PUT',
                            contentType: 'application/json',
                            data: JSON.stringify(syncData),
                            dataType: 'json',
                            success: function (data) {
                                let mediaOption = JSON.parse(data.MediaOption),
                                    newName = `file.${sizeName}.`,
                                    newURL = `${mediaOption.OriginalURL.replace(/file\.original\./, newName)}?${$.now()}`;

                                cropImageData.$target.attr('data-crop-options', JSON.stringify(cropOption)).html(`<img src='${newURL}'>`);
                                if (!hasSize) {
                                    $bannerHtml.attr({
                                        'data-image-width': cropOptionWidth,
                                        'data-image-height': cropOptionHeight
                                    });
                                    $canvas.css({
                                        width: cropOptionWidth,
                                        height: cropOptionHeight
                                    });
                                    $iframe.height(cropOptionHeight);
                                }

                                _this.setValue();
                                $cropModal.qorModal('hide');
                                $cropModal.find('.qor-modal-loading').remove();
                            }
                        });
                    });
                }
            });
        },

        stopCrop: function () {
            this.$cropModal
                .trigger('disable.qor.material')
                .find(CLASS_CROP_WRAP + ' > img')
                .cropper('destroy')
                .remove();
        },

        hideElement: function (e) {
            if (!$(e.target).closest('.qor-bannereditor__contents').length) {
                this.clearElements();
            }
        },

        clearElements: function () {
            let $canvas = this.$canvas;

            $canvas.find('.qor-bannereditor__button-inline,.qor-bannereditor__draggable-coordinate').remove();
            $canvas.find(CLASS_DRAGGABLE).removeClass('qor-bannereditor__dragging');
        },

        openBottomSheets: function (e) {
            var $ele = $(e.target).closest(CLASS_BANNEREDITOR_IMAGE),
                data = $ele.data();

            this.BottomSheets = $('body').data('qor.bottomsheets');

            data.url = data.bannerMediaUrl;
            this.bannerMediaUrl = data.bannerMediaUrl;
            this.BottomSheets.open(data, this.handleBannerImage.bind(this));

            return false;
        },

        handleBannerImage: function ($bottomsheets) {
            let options = {
                loading: this.addMediaLoading,
                onSelect: this.addBannerImage.bind(this),
                onSubmit: this.addBannerImage.bind(this)
            };

            $bottomsheets.qorSelectCore(options).addClass(CLASS_MEDIABOX);
            this.$bottomsheets = $bottomsheets;
            this.initMedia($bottomsheets);
        },

        addBannerImage: function (data) {
            let MediaOption = data.MediaOption,
                $ele = data.$clickElement,
                $urlEle = $ele && $ele.find('[data-heading="BannerEditorUrl"]'),
                $loading = $('<div class="qor-bannereditor-laoding"><div class="mdl-spinner mdl-spinner--single-color mdl-js-spinner is-active"></div></div>'),
                initWidth = this.initWidth,
                initHeight = this.initHeight,
                _this = this,
                imgUrl,
                mediaLibraryUrl = data.mediaLibraryUrl,
                sizeName = `bannereditor_${(Math.random() + 1).toString(36).substring(7)}`,
                syncData = {
                    MediaOption: {
                        Sizes: {},
                        Crop: true
                    }
                };

            if (!mediaLibraryUrl) {
                mediaLibraryUrl = `${this.bannerMediaUrl}/${data.primaryKey}`;
            }

            if (initWidth || initHeight) {
                let croppedSizeName = this.hasCroppedSize(data);

                if (croppedSizeName) {
                    this.insertImage(MediaOption.OriginalURL.replace(/\.original/, `.${croppedSizeName}`), mediaLibraryUrl, croppedSizeName);
                    return false;
                }

                let item = syncData.MediaOption.Sizes,
                    $content = this.$content,
                    $bannerHtml = this.$bannerHtml;

                item[sizeName] = {};
                item[sizeName]['Width'] = initWidth || 0;
                item[sizeName]['Height'] = initHeight || 0;

                syncData.MediaOption = JSON.stringify(syncData.MediaOption);
                this.$bottomsheets.remove();

                $.ajax({
                    type: 'PUT',
                    url: mediaLibraryUrl,
                    data: JSON.stringify(syncData),
                    contentType: 'application/json',
                    dataType: 'json',
                    beforeSend: function () {
                        $bannerHtml.find(CLASS_BG_IMAGE).remove();
                        $loading.prependTo($content).trigger('enable.qor.material');
                    },
                    success: function (data) {
                        imgUrl = JSON.parse(data.MediaOption).OriginalURL.replace(/file\.original\./, `file.${sizeName}.`);
                        $content.find('.qor-bannereditor-laoding').remove();
                        _this.insertImage(imgUrl, mediaLibraryUrl, sizeName, false);
                    }
                });
            } else {
                if (MediaOption.URL) {
                    imgUrl = MediaOption.URL;
                } else if ($urlEle) {
                    imgUrl = $urlEle.text();
                } else {
                    imgUrl = JSON.parse(data.File).Url;
                }
                this.insertImage(imgUrl, mediaLibraryUrl, sizeName);
            }

            return false;
        },

        insertImage: function (url, mediaLibraryUrl, sizeName, removed) {
            let $img = $(`<span class="qor-bannereditor-image" data-size-name="${sizeName || ''}" data-medialibrary-url="${mediaLibraryUrl}"><img src="${url}" /></span>`);

            if (!removed) {
                this.$bannerHtml.find(CLASS_BG_IMAGE).remove();
                this.$bottomsheets.remove();
            }

            $img.prependTo(this.$bannerHtml);

            if (!(this.initWidth && this.initHeight)) {
                this.resetBoxSize(url);
            }

            if (!$('.qor-bottomsheets').is(':visible')) {
                $('body').removeClass('qor-bottomsheets-open');
            }

            this.setValue();
        },

        hasCroppedSize: function (data) {
            let sizes = data.MediaOption.Sizes,
                keys = sizes ? Object.keys(sizes) : [],
                initWidth = this.initWidth,
                initHeight = this.initHeight;

            if (!sizes || !keys.length) {
                return false;
            }

            for (let i = 0; i < keys.length; i++) {
                let key = keys[i];
                if (/^bannereditor_/.test(key)) {
                    let obj = sizes[key];
                    if (initWidth === obj.Width && initHeight === obj.Height) {
                        return key;
                    }
                }
            }

            return false;
        },

        addMediaLoading: function ($ele) {
            $('<div class="qor-media-loading"><div class="mdl-spinner mdl-spinner--single-color mdl-js-spinner is-active"></div></div>')
                .appendTo($ele)
                .trigger('enable.qor.material');
        },

        resetBoxSize: function (url) {
            let $canvas = this.$canvas,
                $iframe = this.$iframe,
                initWidth = this.initWidth,
                initHeight = this.initHeight,
                $bannerHtml = this.$bannerHtml,
                _this = this;

            getImgSize(url, function (width, height) {
                width = initWidth || width;
                height = initHeight || height;

                $canvas.width(width).height(height);
                $iframe.height(height);

                $bannerHtml.attr({'data-image-width': width, 'data-image-height': height});
                _this.setValue();
            });
        },

        handleInlineEdit: function (e) {
            let $target = $(e.target).closest(CLASS_DRAGGABLE),
                $canvas = this.$canvas;

            $canvas.find('.qor-bannereditor__button-inline, .qor-bannereditor__draggable-coordinate').remove();
            $canvas.find(CLASS_DRAGGABLE).removeClass(CLASS_BANNEREDITOR_DRAGGING);
            $target.addClass(CLASS_BANNEREDITOR_DRAGGING).append(QorBannerEditor.inlineEdit);

            return false;
        },

        ajaxForm: function (url, title) {
            let $popover = this.$popover,
                _this = this;

            $.ajax(url, {
                method: 'GET',
                dataType: 'html',
                success: function (html) {
                    let $content = $(html).find('.qor-form-container'),
                        popupTitle =
                            title ||
                            $(html)
                                .find('.mdl-layout-title')
                                .html();

                    $content
                        .find('.qor-button--cancel')
                        .attr('data-dismiss', 'modal')
                        .removeAttr('href');
                    $popover.find('.qor-bannereditor__title').html(popupTitle);
                    $popover.find('.qor-bannereditor__content').html($content.html());
                    $popover.trigger('enable').qorModal('show');

                    $popover.off(EVENT_CLICK).on(EVENT_CLICK, '.qor-bannereditor__content button[type="submit"]', _this.renderElement.bind(_this));
                }
            });
        },

        showEdit: function (e) {
            let $target = $(e.target).closest('button'),
                type = $target.data('edit-type'),
                $element = $target.closest(CLASS_DRAGGABLE),
                data = $element.data();

            switch (type) {
                case 'edit':
                    this.showEditForm(data, $element);
                    break;

                case 'delete':
                    this.deleteElement($element);
                    break;

                case 'left':
                case 'center':
                case 'right':
                    this.alignHorizontally($element, type);
                    break;

                case 'top':
                case 'middle':
                case 'bottom':
                    this.alignVertically($element, type);
                    break;
            }

            e.stopPropagation();
            return false;
        },

        alignHorizontally: function ($element, type) {
            $element.attr('align-horizontally', type);
            this.alignElement($element, type);
        },

        alignVertically: function ($element, type) {
            $element.attr('align-vertically', type);
            this.alignElement($element, type);
        },

        alignElement: function ($element, type) {
            let options = this.options,
                horizontally = $element.attr('align-horizontally'),
                vertically = $element.attr('align-vertically'),
                positionWidth = 300 - $element.width(),
                css = options[type];

            if (vertically === 'top') {
                $element.addClass(CLASS_TOP);
            } else if ($element.hasClass(CLASS_TOP)) {
                $element.removeClass(CLASS_TOP);
            }

            if (horizontally === 'left' && positionWidth > 0) {
                $element.addClass(CLASS_LEFT);
            } else if ($element.hasClass(CLASS_LEFT)) {
                $element.removeClass(CLASS_LEFT);
            }

            if (horizontally === 'center' && vertically === 'middle') {
                css = options.centermiddle;
            } else if (horizontally && vertically) {
                css = $.extend({}, options[horizontally], options[vertically]);
            }

            $element
                .css('transform', '')
                .css(css)
                .attr('data-position-left', parseInt($element.css('left')))
                .attr('data-position-top', parseInt($element.css('top')));

            this.setValue();
        },

        showInlineEdit: function (e) {
            let $ele = $(e.target).closest(CLASS_DRAGGABLE);

            this.showEditForm($ele.data(), $ele);
        },

        showEditForm: function (data, $element) {
            let url = this.config.editURL.replace(/:id/, data.editId);

            this.$editElement = $element;
            this.ajaxForm(url, null, true, $element);
        },

        deleteElement: function ($element) {
            $element.remove();
            this.setValue();
        },

        handleDrag: function (event, ui) {
            let $target = ui.helper,
                positionWidth = 300 - $target.width();

            ui.position.left = parseInt(ui.position.left, 10);
            ui.position.top = parseInt(ui.position.top, 10);

            $target.css({right: 'auto', bottom: 'auto'});

            if ($target.css('transform')) {
                $target.css('transform', '');
            }

            if ($target.css('bottom') === '0px' || $target.css('right') === '0px') {
                $target.css({bottom: 'auto', right: 'auto'});
            }

            if ($target.attr('align-vertically') || $target.attr('align-horizontally')) {
                $target.removeAttr('align-vertically').removeAttr('align-horizontally');
            }

            if (ui.position.top < 40) {
                $target.addClass(CLASS_TOP);
            } else if ($target.hasClass(CLASS_TOP)) {
                $target.removeClass(CLASS_TOP);
            }

            if (positionWidth > 0 && ui.position.left < positionWidth) {
                $target.addClass(CLASS_LEFT);
            } else if ($target.hasClass(CLASS_LEFT)) {
                $target.removeClass(CLASS_LEFT);
            }

            this.$canvas.find('.qor-bannereditor__draggable-coordinate').remove();
            $target.addClass(CLASS_BANNEREDITOR_DRAGGING).append(window.Mustache.render(QorBannerEditor.dragCoordinate, ui.position));
            $target.find('.qor-bannereditor__button-inline').hide();
        },

        handleDragStop: function (event, ui) {
            let cWidth = this.$canvas.width(),
                cHeight = this.$canvas.height(),
                helper = ui.helper,
                helperLeft = ui.position.left / cWidth * 100,
                helperTop = ui.position.top / cHeight * 100,
                css = {
                    left: helperLeft + '%',
                    top: helperTop + '%'
                };

            if (helperLeft > 55) {
                css.right = (cWidth - ui.position.left - helper.outerWidth()) / cWidth * 100 + '%';
                css.left = 'auto';
            }

            if (helperTop > 55) {
                css.bottom = (cHeight - ui.position.top - helper.outerHeight()) / cHeight * 100 + '%';
                css.top = 'auto';
            }

            helper.css(css).attr({
                'data-position-left': ui.position.left,
                'data-position-top': ui.position.top
            });

            if (!helper.find('.qor-bannereditor__button-inline').length) {
                helper.removeClass(CLASS_BANNEREDITOR_DRAGGING);
            }

            ui.helper.find('.qor-bannereditor__button-inline').show();
            ui.helper.find('.qor-bannereditor__draggable-coordinate').remove();
            this.setValue();
        },

        handleResizeStop: function (event, ui) {
            let cWidth = this.$canvas.width(),
                helperWidth = ui.size.width / cWidth * 100 + '%';

            ui.helper.css('width', helperWidth);
            this.setValue();
        },

        renderElement: function (e) {
            let $form = $(e.target).closest('form'),
                url = $form.prop('action'),
                method = $form.prop('method'),
                _this = this,
                $bannerHtml = this.$bannerHtml,
                $popover = this.$popover,
                $editElement = this.$editElement,
                defaultCSS = {
                    position: 'absolute',
                    left: '10%',
                    top: '10%'
                },
                options = this.options;

            if (!$form.length) {
                return;
            }

            // remove replicator template data
            $form.find('.qor-fieldset--new').remove();

            $.ajax(url, {
                method: method,
                dataType: 'json',
                data: new FormData($form[0]),
                processData: false,
                contentType: false,
                success: function (data) {
                    if (!data.Template) {
                        return;
                    }

                    if ($editElement) {
                        $editElement.html(data.Template).resizable(options.resizable);
                        _this.$editElement = null;
                    } else {
                        $(`<span class="${CLASS_DRAGGABLE.slice(1)}">${data.Template}</span>`)
                            .css(defaultCSS)
                            .attr('data-edit-id', data.ID)
                            .appendTo($bannerHtml)
                            .draggable(options.draggable)
                            .resizable(options.resizable);
                    }
                    $popover.qorModal('hide');
                    _this.setValue();
                    _this.$popover.off(EVENT_CLICK);
                },
                error: function (xhr, textStatus, errorThrown) {
                    if (xhr.status === 422 && xhr.responseJSON.errors[0]) {
                        _this.$popover.find('form').before(window.Mustache.render(QorBannerEditor.errorMessage, {message: xhr.responseJSON.errors[0]}));
                    } else {
                        window.alert([textStatus, errorThrown].join(': '));
                    }
                }
            });

            return false;
        },

        addElements: function (e) {
            let $target = $(e.target).closest('button'),
                url = $target.data('banner-url'),
                title = $target.data('title');

            this.ajaxForm(url, title);
        },

        setValue: function () {
            let $html = this.$canvas.clone(),
                $textarea = this.$textarea,
                newValue,
                bannerValues = JSON.parse($textarea.val()),
                platformName = this.platformName;

            $html.find(CLASS_DRAGGABLE).removeClass('ui-draggable-handle ui-resizable ui-draggable-dragging qor-bannereditor__dragging');
            $html.find(CLASS_NEED_REMOVE).remove();

            if (this.$bannerHtml.is(':empty')) {
                newValue = '';
            } else {
                newValue = encodeURIComponent($html.html());
            }

            if (getObject(bannerValues, platformName)) {
                bannerValues.splice(
                    _.findIndex(bannerValues, function (obj) {
                        return obj.Name === platformName;
                    }),
                    1
                );
            }

            bannerValues.push({
                Name: platformName,
                Value: newValue
            });

            $textarea.val(JSON.stringify(bannerValues));
        },

        destroy: function () {
            this.unbind();
            this.$element.removeData(NAMESPACE);
        }
    };

    QorBannerEditor.DEFAULTS = {
        draggable: {
            addClasses: false,
            distance: 10,
            snap: true,
            containment: 'parent',
            scroll: false
        },
        resizable: {
            handles: 'e'
        },
        left: {
            left: 0,
            right: 'auto'
        },
        center: {
            left: '50%',
            right: 'auto',
            transform: 'translateX(-50%)'
        },
        right: {
            left: 'auto',
            right: 0
        },
        top: {
            top: 0,
            bottom: 'auto'
        },
        middle: {
            top: '50%',
            bottom: 'auto',
            transform: 'translateY(-50%)'
        },
        bottom: {
            top: 'auto',
            bottom: 0
        },
        centermiddle: {
            top: '50%',
            left: '50%',
            bottom: 'auto',
            right: 'auto',
            transform: 'translate(-50%,-50%)'
        }
    };

    QorBannerEditor.toolbar = `[[#toolbar]]
                                    <button class="mdl-button mdl-button--colored qor-bannereditor__button" data-banner-url="[[CreateURL]]" id="[[id]]" data-title="[[Label]]" type="button">
                                        [[#Icon]][[& Icon]][[/Icon]]
                                        [[^Icon]]
                                            [[Label]]
                                        [[/Icon]]
                                    </button>
                                    [[#Icon]]
                                        <span class="mdl-tooltip" data-mdl-for="[[id]]">
                                            [[Label]]
                                        </span>
                                    [[/Icon]]
                                [[/toolbar]]`;

    QorBannerEditor.dragCoordinate = `<div class="qor-bannereditor__draggable-coordinate"><span>x :<em>[[left]]</em></span><span>y :<em>[[top]]</em></span></div>`;

    QorBannerEditor.errorMessage = `<ul class="qor-error">
                                        <li>
                                            <label for="">
                                                <i class="material-icons">error</i>
                                                <span>[[message]]</span>
                                            </label>
                                        </li>
                                    </ul>`;

    QorBannerEditor.inlineEdit = `<div class="qor-bannereditor__button-inline">
                                    <button class="mdl-button mdl-button--icon" data-edit-type="left" type="button"><i class="material-icons">format_align_left</i></button>
                                    <button class="mdl-button mdl-button--icon" data-edit-type="center" type="button"><i class="material-icons">format_align_center</i></button>
                                    <button class="mdl-button mdl-button--icon" data-edit-type="right" type="button"><i class="material-icons">format_align_right</i></button>
                                    <hr />
                                    <button class="mdl-button mdl-button--icon" data-edit-type="top" type="button"><i class="material-icons">vertical_align_top</i></button>
                                    <button class="mdl-button mdl-button--icon" data-edit-type="middle" type="button"><i class="material-icons">vertical_align_center</i></button>
                                    <button class="mdl-button mdl-button--icon" data-edit-type="bottom" type="button"><i class="material-icons">vertical_align_bottom</i></button>
                                    <hr />
                                    <button class="mdl-button mdl-button--icon" data-edit-type="edit" type="button"><i class="material-icons">mode_edit</i></button>
                                    <button class="mdl-button mdl-button--icon" data-edit-type="delete" type="button"><i class="material-icons">delete_forever</i></button>
                                  </div>`;

    QorBannerEditor.imageEdit = `<div class="qor-bannereditor__button-inline qor-bannereditor__button-bg">
                                        <button class="mdl-button mdl-button--icon qor-bannereditor__editimage" type="button"><i class="material-icons">mode_edit</i></button>
                                        <button class="mdl-button mdl-button--icon qor-bannereditor__deleteimage" type="button"><i class="material-icons">delete_forever</i></button>
                                    </div>`;

    QorBannerEditor.popover = `<div class="qor-modal fade qor-bannereditor__form" tabindex="-1" role="dialog" aria-hidden="true">
                                  <div class="mdl-card mdl-shadow--2dp" role="document">
                                    <div class="mdl-card__title">
                                        <h2 class="mdl-card__title-text qor-bannereditor__title"></h2>
                                    </div>
                                    <div class="mdl-card__supporting-text qor-bannereditor__content"></div>
                                  </div>
                                </div>`;

    QorBannerEditor.modal = `<div class="qor-modal fade qor-bannereditor-modal__cropimage" tabindex="-1" role="dialog" aria-hidden="true">
                                <div class="mdl-card mdl-shadow--2dp" role="document">
                                    <div class="mdl-card__title">
                                        <h2 class="mdl-card__title-text">$[title]</h2>
                                    </div>
                                    <div class="mdl-card__supporting-text">
                                        <div class="qor-cropper__wrapper"><div class="mdl-spinner mdl-spinner--single-color mdl-js-spinner is-active"></div></div>
                                    </div>
                                    <div class="mdl-card__actions mdl-card--border">
                                        <a class="mdl-button mdl-button--colored mdl-button--raised qor-cropper__save">$[ok]</a>
                                        <a class="mdl-button mdl-button--colored" data-dismiss="modal">$[cancel]</a>
                                    </div>
                                    <div class="mdl-card__menu">
                                        <button class="mdl-button mdl-button--icon" data-dismiss="modal" aria-label="close">
                                            <i class="material-icons">close</i>
                                        </button>
                                    </div>
                                </div>
                            </div>`;

    QorBannerEditor.plugin = function (options) {
        return this.each(function () {
            let $this = $(this),
                data = $this.data(NAMESPACE),
                fn;

            if (!data) {
                if (/destroy/.test(options)) {
                    return;
                }
                $this.data(NAMESPACE, (data = new QorBannerEditor(this, options)));
            }

            if (typeof options === 'string' && $.isFunction((fn = data[options]))) {
                fn.apply(data);
            }
        });
    };

    $(function () {
        let selector = '[data-toggle="qor.bannereditor"]:visible'; // if element is hide, return fasle. (e.g., replicator)

        $(document)
            .on(EVENT_DISABLE, function (e) {
                QorBannerEditor.plugin.call($(selector, e.target), 'destroy');
            })
            .on(EVENT_ENABLE, function (e) {
                QorBannerEditor.plugin.call($(selector, e.target));
            })
            .triggerHandler(EVENT_ENABLE);
    });

    return QorBannerEditor;
});
