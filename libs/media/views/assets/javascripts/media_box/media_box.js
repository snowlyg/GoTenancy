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

    let $body = $('body'),
        $document = $(document),
        NAMESPACE = 'qor.medialibrary.select',
        PARENT_NAMESPACE = 'qor.bottomsheets',
        EVENT_CLICK = 'click.' + NAMESPACE,
        EVENT_ENABLE = 'enable.' + NAMESPACE,
        EVENT_DISABLE = 'disable.' + NAMESPACE,
        EVENT_RELOAD = 'reload.' + PARENT_NAMESPACE,
        CLASS_SELECT_ICON = '.qor-select__select-icon',
        CLASS_SELECT_HINT = '.qor-selectmany__hint',
        CLASS_PARENT = '.qor-field__mediabox',
        CLASS_LISTS = '.qor-field__mediabox-list',
        CLASS_ITEM = '.qor-field__mediabox-item',
        CLASS_LISTS_DATA = '.qor-field__mediabox-data',
        CLASS_SELECTED = 'is_selected',
        CLASS_DELETE = 'is_deleted',
        CLASS_CROPPER_OPTIONS = 'textarea.qor-file__options',
        CLASS_CROPPER_DELETE = '.qor-cropper__toggle--delete',
        CLASS_CROPPER_UNDO = '.qor-cropper__toggle--undo',
        CLASS_MEDIABOX = 'qor-bottomsheets__mediabox';

    function getYoutubeID(url) {
        var regExp = /^.*((youtu.be\/)|(v\/)|(\/u\/\w\/)|(embed\/)|(watch\?))\??v?=?([^#\&\?]*).*/;
        var match = url.match(regExp);
        if (match && match[7].length == 11) {
            return match[7];
        } else {
            return false;
        }
    }

    function QorMediaBox(element, options) {
        this.$element = $(element);
        this.options = $.extend({}, QorMediaBox.DEFAULTS, $.isPlainObject(options) && options);
        this.init();
    }

    QorMediaBox.prototype = {
        constructor: QorMediaBox,

        init: function () {
            var $element = this.$element;
            this.SELECT_MEDIABOX_UNDO_TEMPLATE = $element.find('[name="media-box-undo-delete"]').html();
            this.bind();
            this.initSelectedMedia();
        },

        bind: function () {
            $document.off(EVENT_RELOAD).on(EVENT_RELOAD, `.${CLASS_MEDIABOX}`, this.reloadData.bind(this));

            this.$element
                .off(EVENT_CLICK)
                .off('change.qor.cropper')
                .on(EVENT_CLICK, '[data-mediabox-url]', this.openBottomSheets.bind(this))
                .on(EVENT_CLICK, CLASS_CROPPER_DELETE, this.deleteSelected.bind(this))
                .on(EVENT_CLICK, CLASS_CROPPER_UNDO, this.undoDeleteSelected.bind(this))
                .on('change.qor.cropper', CLASS_CROPPER_OPTIONS, this.imageCrop.bind(this));
        },

        unbind: function () {
            $document.off(EVENT_RELOAD, `.${CLASS_MEDIABOX}`);

            this.$element.off(EVENT_CLICK).off('change.qor.cropper');
        },

        deleteSelected: function (e) {
            var $target = $(e.target),
                $selectFeild = $target.closest(CLASS_ITEM);

            $selectFeild
                .addClass(CLASS_DELETE)
                .append(this.SELECT_MEDIABOX_UNDO_TEMPLATE)
                .find('.qor-file__list')
                .hide();
            this.updateMediaLibraryData($target.closest(CLASS_LISTS));
            this.$element.find(CLASS_LISTS_DATA).data('isDeleted', true);

            return false;
        },

        undoDeleteSelected: function (e) {
            var $target = $(e.target),
                $selectFeild = $target.closest(CLASS_ITEM);

            $selectFeild
                .removeClass(CLASS_DELETE)
                .find('.qor-file__list')
                .show();
            this.updateMediaLibraryData($target.closest(CLASS_LISTS));
            $target.closest('.qor-fieldset__alert').remove();
            this.$element.find(CLASS_LISTS_DATA).data('isDeleted', false);

            return false;
        },

        imageCrop: function (e) {
            var $parent = $(e.target).closest(CLASS_ITEM);
            this.syncImageCrop($parent, this.resetImages);
        },

        openBottomSheets: function (e) {
            var $ele = $(e.target).closest('[data-mediabox-url]'),
                data = $ele.data(),
                $parent;

            if (data.isDisabled) {
                return;
            }

            this.BottomSheets = $body.data('qor.bottomsheets');
            this.bottomsheetsData = data;
            this.$parent = $parent = $ele.closest(CLASS_PARENT);
            this.$selectFeild = $parent.find(CLASS_LISTS);

            data.url = data.mediaboxUrl;

            // select many templates
            this.SELECT_MANY_SELECTED_ICON = $('[name="media-box-select-many-selected-icon"]').html();
            this.SELECT_MANY_HINT = $('[name="media-box-select-many-hint"]').html();

            this.TEMPLATE_IMAGE = $parent.find('[name="media-box-template"]').html();
            this.TEMPLATE_FILE = $parent.find('[name="media-box-file-template"]').html();
            this.TEMPLATE_UPLOADEDVIDEO = $parent.find('[name="media-box-uploadedvideo-template"]').html();
            this.TEMPLATE_VIDEOLINK = $parent.find('[name="media-box-videolink-template"]').html();
            this.SELECT_MEDIABOX_UNDO_TEMPLATE = $parent.find('[name="media-box-undo-delete"]').html();

            this.BottomSheets.open(data, this.handleSelectMany.bind(this));
        },

        initSelectedMedia: function () {
            var $element = this.$element,
                $selectedMedias = $element.find(CLASS_ITEM),
                $selectedMedia,
                selectedMediaData,
                mediaData = JSON.parse($element.find(CLASS_LISTS_DATA).val());

            if (mediaData) {
                for (var i = 0; i < mediaData.length; i++) {
                    $selectedMedia = $selectedMedias.filter('[data-primary-key="' + mediaData[i].ID + '"]');
                    selectedMediaData = $selectedMedia.data().description;
                    if (!selectedMediaData) {
                        $selectedMedia.data('description', mediaData[i].Description);
                    }
                }
            }
        },

        initMedia: function () {
            var $selectFeild = this.$selectFeild,
                $items = $selectFeild.find(CLASS_ITEM).not('.' + CLASS_DELETE),
                $trs = this.$bottomsheets.find('tbody tr'),
                _this = this,
                $tr,
                $img,
                key;

            $items.each(function () {
                key = $(this).data().primaryKey;
                $tr = $trs.filter('[data-primary-key="' + key + '"]').addClass(CLASS_SELECTED);
                _this.changeIcon($tr, true);
            });

            $trs.each(function () {
                $tr = $(this);
                $img = $tr.find('.qor-table--ml-slideout p img').first();
                $tr.find('.qor-table__actions').remove();
                if ($img.length) {
                    $tr.find('.qor-table--medialibrary-item').css('background-image', 'url(' + $img.prop('src') + ')');
                    $img.parent().remove();
                }
            });
            if (this.bottomsheetsData.maxItem != '1') {
                this.updateHint(this.getSelectedItemData());
            }
        },

        reloadData: function () {
            this.$selectFeild && this.initMedia();
        },

        renderHint: function (data) {
            return window.Mustache.render(this.SELECT_MANY_HINT, data);
        },

        getSelectedItemData: function ($ele) {
            let $selectFeild = $ele ? $ele : this.$selectFeild,
                $items = $selectFeild.find(CLASS_ITEM).not('.' + CLASS_DELETE),
                files = [];

            if ($items.length) {
                $items.each(function () {
                    let data = $(this).data();

                    files.push({
                        ID: data.primaryKey,
                        Url: data.originalUrl.replace(/.original.(\w+)$/, '.$1'),
                        Description: data.description,
                        FileName: data.fileName,
                        VideoLink: data.videolink
                    });
                });
            }

            return {
                files: files,
                selectedNum: files.length
            };
        },

        updateHint: function (data) {
            var template;

            $.extend(data, this.bottomsheetsData);
            template = this.renderHint(data);

            $(CLASS_SELECT_HINT).remove();
            this.$bottomsheets.find('.qor-page__body').before(template);
        },

        updateMediaLibraryData: function ($ele, data) {
            var $dataInput = $ele ? $ele.find(CLASS_LISTS_DATA) : this.$selectFeild.find(CLASS_LISTS_DATA),
                fileData = this.getSelectedItemData($ele);

            $dataInput
                .val(JSON.stringify(fileData.files))
                .data('mediaData', data)
                .trigger('changed.medialibrary', [data]);
        },

        changeIcon: function ($ele, isNew) {
            var $item = $ele.find('.qor-table--medialibrary-item'),
                $target = $item.length ? $item : $ele.find('td:first');

            $ele.find(CLASS_SELECT_ICON).remove();

            if (isNew) {
                if (isNew == 'one') {
                    $('.' + CLASS_MEDIABOX)
                        .find(CLASS_SELECT_ICON)
                        .remove();
                }
                $target.prepend(this.SELECT_MANY_SELECTED_ICON);
            }
        },

        syncImageCrop: function ($ele, callback) {
            let item = JSON.parse($ele.find(CLASS_CROPPER_OPTIONS).val()),
                $btn = $ele.closest("form").find(":submit"),
                url = $ele.data().mediaLibraryUrl,
                _this = this,
                syncData = {},
                sizes = ['Width', 'Height'],
                sizeResolutionData,
                sizeData,
                $imgs = $ele.find('img[data-size-name]');

            delete item.ID;
            delete item.Url;

            item.Sizes = {};

            $imgs.each(function () {
                sizeData = $(this).data();

                if (sizeData['sizeResolutionWidth'] || sizeData['sizeResolution']) {
                    item['Sizes'][sizeData.sizeName] = {};
                    for (let i = 0; i < sizes.length; i++) {
                        sizeResolutionData = sizeData['sizeResolution' + sizes[i]];
                        if (!sizeResolutionData) {
                            sizeResolutionData = sizeData['sizeResolution'][sizes[i]];
                        }
                        item['Sizes'][sizeData.sizeName][sizes[i]] = sizeResolutionData;
                    }
                }
            });

            syncData.MediaOption = JSON.stringify(item);

            this.addMediaLoading($ele);

            $.ajax({
                type: 'PUT',
                url: url,
                data: JSON.stringify(syncData),
                contentType: 'application/json',
                dataType: 'json',
                beforeSend: function () {
                    $btn.attr("disabled", true);
                },
                success: function (data) {
                    syncData.MediaOption = JSON.parse(data.MediaOption);
                    $ele.data().originalUrl = syncData.MediaOption.OriginalURL;
                    _this.updateMediaLibraryData($ele.closest(CLASS_LISTS), syncData);

                    if (callback && $.isFunction(callback)) {
                        callback(syncData, $ele);
                    }
                },
                complete: function () {
                    $btn.attr("disabled", false);
                    $ele.find('.qor-media-loading').remove();
                }
            });
        },

        showHiddenItem: function ($hiddenItem) {
            $hiddenItem
                .removeClass(CLASS_DELETE)
                .find('.qor-file__list')
                .show();
            $hiddenItem.find('.qor-fieldset__alert').remove();
        },

        removeItem: function (data) {
            let primaryKey = data.primaryKey;

            this.$selectFeild.find('[data-primary-key="' + primaryKey + '"]').remove();
            this.changeIcon(data.$clickElement);
        },

        compareCropSizes: function (data) {
            let cropOptions = data.MediaOption.CropOptions,
                needCropSizes = this.bottomsheetsData.cropSizes,
                sizesLen,
                cropOptionsKeys;

            if (!needCropSizes || data.SelectedType != 'image') {
                return false;
            }

            if (cropOptions === undefined) {
                return true;
            }

            needCropSizes = needCropSizes.split(',');
            sizesLen = needCropSizes.length - 1;

            if (window._.isObject(cropOptions)) {
                cropOptionsKeys = Object.keys(cropOptions);
            } else {
                return false;
            }

            if (cropOptionsKeys.length) {
                for (let i = 0; i < sizesLen; i++) {
                    if (cropOptionsKeys.indexOf(needCropSizes[i]) == -1) {
                        return true;
                    }
                }
            }

            return false;
        },

        addItem: function (data, isNewData) {
            let $template = $(window.Mustache.render(this.TEMPLATE_IMAGE, data)),
                $input = $template.find('.qor-file__input'),
                $item = $input.closest(CLASS_ITEM),
                $btn = $item.closest("form").find(":submit"),
                $hiddenItem = this.$selectFeild.find('[data-primary-key="' + data.primaryKey + '"]'),
                maxItem = this.bottomsheetsData.maxItem,
                selectedItem = this.getSelectedItemData().selectedNum,
                cropOptions = data.MediaOption.CropOptions,
                hasNewCropSize = this.compareCropSizes(data),
                selectedType = data.SelectedType,
                isSVG = /.svg$/.test(data.MediaOption.FileName),
                _this = this;

            $btn.attr("disabled", true);

            if (!isNewData) {
                if (maxItem == 1) {
                    this.changeIcon(data.$clickElement, 'one');
                } else {
                    this.changeIcon(data.$clickElement, true);
                }
            }

            if (maxItem && selectedItem >= maxItem) {
                if (maxItem == 1) {
                    this.$selectFeild.find(CLASS_ITEM).remove();
                } else {
                    window.alert(this.bottomsheetsData.maxItemHint);
                    return;
                }
            }

            if ($hiddenItem.length) {
                this.showHiddenItem($hiddenItem);
                if (maxItem == 1) {
                    setTimeout(function () {
                        _this.$bottomsheets.remove();
                        if (!$('.qor-bottomsheets').is(':visible')) {
                            $('body').removeClass('qor-bottomsheets-open');
                        }
                    }, 1000);
                }
                return;
            }

            if (maxItem == 1) {
                this.$selectFeild
                    .find(CLASS_ITEM)
                    .filter('.is_deleted')
                    .remove();
            }

            if (!isSVG) {
                if (selectedType === 'video') {
                    $template = $(window.Mustache.render(this.TEMPLATE_UPLOADEDVIDEO, data));
                } else if (selectedType === 'video_link') {
                    let externalVideoLink = getYoutubeID(data.MediaOption.Video);

                    data.VideoLink = externalVideoLink ? `//www.youtube.com/embed/${externalVideoLink}?rel=0&fs=0&modestbranding=1&disablekb=1` : data.MediaOption.Video;
                    $template = $(window.Mustache.render(this.TEMPLATE_VIDEOLINK, data));
                } else if (selectedType === 'file') {
                    $template = $(window.Mustache.render(this.TEMPLATE_FILE, data));
                }
            }

            $template.data({
                description: data.MediaOption.Description,
                mediaData: data,
                videolink: selectedType === 'video_link' ? data.MediaOption.Video : ''
            });

            if (isSVG) {
                $template
                    .addClass('is-svg')
                    .find('.qor-file__input')
                    .remove();
            }
            $template.appendTo(this.$selectFeild);

            // if image alread have CropOptions, replace original images as [big,middle, small] images.
            if (cropOptions && selectedType === 'image') {
                this.resetImages(data, $template);
            }

            // trigger cropper function for new item
            if (selectedType === 'image') {
                $template.find(CLASS_CROPPER_OPTIONS).val(JSON.stringify(data.MediaOption));
            }

            $template.trigger('enable');

            // if not have crop options or have crop options but have anothre size name to crop
            if (hasNewCropSize && $input.data('qor.cropper') && !isSVG) {
                $input.data('qor.cropper').load(data.MediaOption.URL, true, function () {
                    _this.syncImageCrop($item, _this.resetImages);
                });
            }

            this.$bottomsheets.find('.qor-media-loading').remove();

            if (isNewData || maxItem == 1) {
                setTimeout(function () {
                    _this.$bottomsheets.remove();
                    if (!$('.qor-bottomsheets').is(':visible')) {
                        $('body').removeClass('qor-bottomsheets-open');
                    }
                }, 150);
            }
        },

        resetImages: function (data, $template) {
            let cropOptions = data.MediaOption.CropOptions,
                keys = Object.keys(cropOptions),
                url = data.MediaOption.OriginalURL;

            if (!/original/.test(url)) {
                return;
            }

            for (let i = keys.length - 1; i >= 0; i--) {
                cropOptions[keys[i]]['URL'] = url.replace(/original/, keys[i]);
            }

            $template.find('img[data-size-name]').each(function () {
                let $this = $(this),
                    randomString = (Math.random() + 1).toString(36).substring(7),
                    sizeName = $this.data().sizeName;

                if (sizeName != 'original' && cropOptions[sizeName]) {
                    $this.prop('src', `${cropOptions[sizeName]['URL']}?${randomString}`);
                }
            });
        },

        handleSelectMany: function ($bottomsheets) {
            let options = {
                loading: this.addMediaLoading,
                onSelect: this.onSelectResults.bind(this),
                onSubmit: this.onSubmitResults.bind(this)
            };

            $bottomsheets.qorSelectCore(options).addClass(CLASS_MEDIABOX);
            this.$bottomsheets = $bottomsheets;
            this.initMedia();
        },

        onSelectResults: function (data) {
            this.handleResultsData(data);
        },

        onSubmitResults: function (data) {
            this.handleResultsData(data, true);
        },

        handleResultsData: function (data, isNewData) {
            var $element = data.$clickElement,
                isSelected;

            if (!data.mediaLibraryUrl && !isNewData) {
                data.mediaLibraryUrl = data.url;
            }

            if (isNewData) {
                data.mediaLibraryUrl = this.bottomsheetsData.mediaboxUrl.split('?')[0] + '/' + data.primaryKey;
                this.addItem(data, isNewData);
                this.updateDatas(data);
                return;
            }

            $element.toggleClass(CLASS_SELECTED);
            isSelected = $element.hasClass(CLASS_SELECTED);

            if (isSelected) {
                this.addItem(data);
            } else {
                this.removeItem(data);
            }
            this.updateDatas(data);
        },

        addMediaLoading: function ($ele) {
            $('<div class="qor-media-loading"><div class="mdl-spinner mdl-spinner--single-color mdl-js-spinner is-active"></div></div>')
                .appendTo($ele)
                .trigger('enable.qor.material');
        },

        updateDatas: function (data) {
            if (this.bottomsheetsData.maxItem != '1') {
                this.updateHint(this.getSelectedItemData());
            }
            this.updateMediaLibraryData(null, data);
        },

        destroy: function () {
            this.unbind();
            this.$element.removeData(NAMESPACE);
        }
    };

    QorMediaBox.plugin = function (options) {
        return this.each(function () {
            var $this = $(this);
            var data = $this.data(NAMESPACE);
            var fn;

            if (!data) {
                if (/destroy/.test(options)) {
                    return;
                }

                $this.data(NAMESPACE, (data = new QorMediaBox(this, options)));
            }

            if (typeof options === 'string' && $.isFunction((fn = data[options]))) {
                fn.apply(data);
            }
        });
    };

    $(function () {
        var selector = '[data-toggle="qor.mediabox"]';
        $(document)
            .on(EVENT_DISABLE, function (e) {
                QorMediaBox.plugin.call($(selector, e.target), 'destroy');
            })
            .on(EVENT_ENABLE, function (e) {
                QorMediaBox.plugin.call($(selector, e.target));
            })
            .triggerHandler(EVENT_ENABLE);
    });

    return QorMediaBox;
});
