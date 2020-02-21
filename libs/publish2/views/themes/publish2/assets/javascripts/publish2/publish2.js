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
    let $document = $(document),
        NAMESPACE = 'qor.publish2',
        EVENT_ENABLE = 'enable.' + NAMESPACE,
        EVENT_DISABLE = 'disable.' + NAMESPACE,
        EVENT_CLICK = 'click.' + NAMESPACE,
        EVENT_CHANGE = 'change.' + NAMESPACE,
        EVENT_SELECTONE_SELECTED = 'qor.selectone.selected qor.selectone.unselected',
        EVENT_REPLICATOR_ADDED = 'added.qor.replicator.publish2',
        // sharedable version input name, please change this if adjust name in template !!
        // <input name="QorResource.ColorVariations[0].SizeVariations[0].ShareableVersion" />
        NAME_SHAREABLEVERSION = 'ShareableVersion',
        CLASS_VERSION_LINK = '.qor-publish2__version',
        VERSION_LIST = 'qor-table__inner-list',
        VERSION_BLOCK = 'qor-table__inner-block',
        CLASS_VERSION_LIST = '.' + VERSION_LIST,
        CLASS_VERSION_BLOCK = '.' + VERSION_BLOCK,
        CLASS_EVENT_ID = '.qor-pulish2__eventid',
        CLASS_EVENT_INPUT = '.qor-pulish2__eventid-input',
        CLASS_PUBLISH_ACTION = '.qor-pulish2__action',
        CLASS_PUBLISH_ACTION_START = '.qor-pulish2__action-start',
        CLASS_PUBLISH_ACTION_END = '.qor-pulish2__action-end',
        CLASS_PUBLISH_ACTION_INPUT = '.qor-pulish2__action-input',
        CLASS_PICKER_BUTTON = '.qor-action__picker-button',
        CLASS_MEDIALIBRARY_TR = '>tbody>tr',
        IS_MEDIALIBRARY = 'qor-table--medialibrary',
        IS_SHOW_VERSION = 'is-showing';

    function QorPublish2(element, options) {
        this.$element = $(element);
        this.options = $.extend({}, QorPublish2.DEFAULTS, $.isPlainObject(options) && options);
        this.init();
    }

    QorPublish2.prototype = {
        constructor: QorPublish2,

        init: function () {
            this.actionType = this.options.element;
            this.bind();
            this.initActionTemplate();
        },

        bind: function () {
            $document
                .on(EVENT_CLICK, CLASS_VERSION_LINK, this.loadPublishVersion.bind(this))
                .on(EVENT_CHANGE, CLASS_PUBLISH_ACTION_INPUT, this.action.bind(this))
                .on(EVENT_SELECTONE_SELECTED, CLASS_EVENT_ID, this.eventidChanged.bind(this))
                .on(EVENT_REPLICATOR_ADDED, this.replicatorAdded.bind(this));
        },

        unbind: function () {
            $document
                .off(EVENT_CLICK, CLASS_VERSION_LINK)
                .off(EVENT_CHANGE, CLASS_PUBLISH_ACTION_INPUT)
                .off(EVENT_SELECTONE_SELECTED, CLASS_EVENT_ID)
                .off(EVENT_REPLICATOR_ADDED);
        },

        initActionTemplate: function () {
            if (!$(CLASS_PUBLISH_ACTION).closest('.qor-slideout').length) {
                $(CLASS_PUBLISH_ACTION)
                    .prependTo($('.mdl-layout__content .qor-page__body form').first())
                    .show();
            }
            QorPublish2.initSharedVersion();
        },

        replicatorAdded: function (e, $element) {
            QorPublish2.generateSharedVersionLabel($element);
        },

        action: function (e) {
            var $element = $(e.target),
                isCheckbox = $element.prop('type') == 'checkbox',
                currentValue = $element.val(),
                $target = $(this.actionType[$element.data().actionType]),
                $checkboxLabel = $target.closest('label');

            if (!$target.length) {
                return;
            }

            if (isCheckbox) {
                $target.prop('checked', $element.is(':checked'));
                $element.is(':checked') ? $checkboxLabel.addClass('is-checked') : $checkboxLabel.removeClass('is-checked');
            } else {
                $target.val(currentValue);
            }
        },

        eventidChanged: function (e, data) {
            if (data) {
                $(CLASS_EVENT_INPUT).val(data.primaryKey);
            } else {
                $(CLASS_EVENT_INPUT).val('');
            }

            this.updateDate(data, e.target);
            $(CLASS_EVENT_INPUT).trigger('change');
        },

        updateDate: function (data, element) {
            var $parent = $(element).closest(CLASS_PUBLISH_ACTION),
                $start = $parent.find(CLASS_PUBLISH_ACTION_START),
                $end = $parent.find(CLASS_PUBLISH_ACTION_END),
                $button = $parent.find(CLASS_PICKER_BUTTON),
                $input = $button.parent().find('input');

            if (data) {
                $start.val(data.ScheduledStartAt);
                $end.val(data.ScheduledEndAt);
                $button.hide();
                $input.attr('disabled', true);
            } else {
                $button.show();
                $input
                    .attr('disabled', false)
                    .closest('.is-disabled')
                    .removeClass('is-disabled');
            }

            $start.trigger('change');
            $end.trigger('change');
        },

        loadPublishVersion: function (e) {
            var $target = $(e.target).parent('a'),
                url = $target.data().versionUrl,
                $table = $target.closest('table'),
                $tr = $target.closest('tr'),
                colspan = $tr.find('td').length,
                isMediaLibrary = $table.hasClass(IS_MEDIALIBRARY),
                $list,
                $newRow = $('<tr class="' + VERSION_LIST + '"><td colspan="' + colspan + '"></td></tr>'),
                $version = $('<div class="' + VERSION_BLOCK + '"><div style="text-align: center;"><div class="mdl-spinner mdl-js-spinner is-active"></div></div></div>');

            if ($tr.hasClass(IS_SHOW_VERSION)) {
                $(CLASS_VERSION_LIST).remove();
                $table.find('tr').removeClass(IS_SHOW_VERSION);
                return false;
            }

            $(CLASS_VERSION_LIST).remove();
            $('table tr').removeClass(IS_SHOW_VERSION);

            $tr.addClass(IS_SHOW_VERSION);

            if (isMediaLibrary) {
                var $trs = $table.find(CLASS_MEDIALIBRARY_TR),
                    columnNum = parseInt($table.width() / 217),
                    currentNum = $trs.index($tr) + 1,
                    rows = Math.ceil($trs.length / columnNum),
                    currentRow = Math.ceil(currentNum / columnNum);

                $tr = $($trs.get(columnNum * currentRow - 1));
                if (!$tr.length) {
                    $tr = $trs.last();
                }
                $newRow = $('<tr class="' + VERSION_LIST + '"><td></td></tr>');
                if (rows > 1) {
                    $newRow.width(217 * columnNum - 16);
                }
            }

            $tr.after($newRow);
            $list = $(CLASS_VERSION_LIST).find('td');

            $version.appendTo($list).trigger('enable');

            url &&
            $.get(url, function (html) {
                $(CLASS_VERSION_BLOCK)
                    .html(html)
                    .trigger('enable');
            });

            return false;
        },

        destroy: function () {
            this.unbind();
            this.$element.removeData(NAMESPACE);
        }
    };

    QorPublish2.generateSharedVersionLabel = function ($element) {
        let sharedVersion = $('[name="shared-version-checkbox"]').html(),
            $inputs = $('input[name$="' + NAME_SHAREABLEVERSION + '"]'),
            data = {},
            randomString;

        if ($element) {
            $inputs = $element.find('input[name$="' + NAME_SHAREABLEVERSION + '"]');
        }

        $inputs.each(function () {
            let $input = $(this),
                $field = $input.closest('.qor-fieldset'),
                $template;

            if ($field.hasClass('qor-fieldset--new')) {
                return;
            }

            randomString = (Math.random() + 1).toString(36).substring(7);
            data.id = [NAME_SHAREABLEVERSION, randomString].join('_');
            $template = $(window.Mustache.render(sharedVersion, data));

            $template.find('input').on(EVENT_CLICK, function () {
                $(this).is(':checked') ? $input.val('true') : $input.val('');
            });

            if ($input.val() == 'true') {
                $template.find('input').prop('checked', true);
            }

            $template.prependTo($field).trigger('enable');
            $input.closest('.qor-field').hide();
        });
    };

    QorPublish2.initSharedVersion = function () {
        if (!$(CLASS_PUBLISH_ACTION).length) {
            return;
        }

        QorPublish2.generateSharedVersionLabel();
    };

    $.fn.qorSliderAfterShow = $.fn.qorSliderAfterShow || {};

    $.fn.qorSliderAfterShow.initSharedVersion = QorPublish2.initSharedVersion;

    $.fn.qorSliderAfterShow.initPublishForm = function () {
        let $action = $(CLASS_PUBLISH_ACTION),
            $types = $action.find('[data-action-type]'),
            $slideoutForm = $('.qor-slideout__body form.qor-form').first(),
            $bottomsheetForm = $('.qor-bottomsheets__body form'),
            isInBottomsheets = $action.closest('.qor-bottomsheets').length,
            isInSlideout = $action.closest('.qor-slideout').length,
            $parent,
            element = QorPublish2.ELEMENT;

        // move publsh2 actions into slideout form tag
        if ($action.length && !$slideoutForm.data("takeover-publish")) {
            if ($bottomsheetForm.length && isInBottomsheets) {
                $action.prependTo($bottomsheetForm.first());
            } else if ($slideoutForm.length && isInSlideout) {
                $action.prependTo($slideoutForm.first());
            }
        }

        if (!$action.length || !$types.length) {
            return;
        }

        if (isInSlideout) {
            $parent = $('.qor-slideout');
        } else if (isInBottomsheets) {
            $parent = $('.qor-bottomsheets');
        }

        $types.each(function () {
            var $this = $(this);
            $parent
                .find(element[$this.data('actionType')])
                .closest('.qor-form-section')
                .hide();
        });

        $(CLASS_PUBLISH_ACTION_INPUT).trigger(EVENT_CHANGE);
    };

    QorPublish2.DEFAULTS = {};

    QorPublish2.ELEMENT = {
        scheduledstart: '[name="QorResource.ScheduledStartAt"]',
        scheduledend: '[name="QorResource.ScheduledEndAt"]',
        publishready: '[name="QorResource.PublishReady"]',
        versionname: '[name="QorResource.VersionName"]',
        eventid: '[name="QorResource.ScheduledEventID"]'
    };

    QorPublish2.plugin = function (options) {
        return this.each(function () {
            var $this = $(this);
            var data = $this.data(NAMESPACE);
            var fn;

            if (!data) {
                if (/destroy/.test(options)) {
                    return;
                }

                $this.data(NAMESPACE, (data = new QorPublish2(this, options)));
            }

            if (typeof options === 'string' && $.isFunction((fn = data[options]))) {
                fn.apply(data);
            }
        });
    };

    $(function () {
        var selector = '.qor-theme-publish2';
        var options = {};

        options['element'] = QorPublish2.ELEMENT;

        $(document)
            .on(EVENT_DISABLE, function (e) {
                QorPublish2.plugin.call($(selector, e.target), 'destroy');
            })
            .on(EVENT_ENABLE, function (e) {
                QorPublish2.plugin.call($(selector, e.target), options);
            })
            .triggerHandler(EVENT_ENABLE);
    });

    return QorPublish2;
});
