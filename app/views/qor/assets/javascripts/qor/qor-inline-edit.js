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

    const NAMESPACE = 'qor.inlineEdit',
        EVENT_ENABLE = 'enable.' + NAMESPACE,
        EVENT_DISABLE = 'disable.' + NAMESPACE,
        EVENT_CLICK = 'click.' + NAMESPACE,
        EVENT_MOUSEENTER = 'mouseenter.' + NAMESPACE,
        EVENT_MOUSELEAVE = 'mouseleave.' + NAMESPACE,
        CLASS_FIELD = '.qor-field',
        CLASS_FIELD_SHOW = '.qor-field__show',
        CLASS_FIELD_SHOW_INNER = '.qor-field__show-inner',
        CLASS_EDIT = '.qor-inlineedit__edit',
        CLASS_SAVE = '.qor-inlineedit__save',
        CLASS_BUTTONS = '.qor-inlineedit__buttons',
        CLASS_CANCEL = '.qor-inlineedit__cancel',
        CLASS_CONTAINER = 'qor-inlineedit__field';

    function QorInlineEdit(element, options) {
        this.$element = $(element);
        this.options = $.extend({}, QorInlineEdit.DEFAULTS, $.isPlainObject(options) && options);
        this.init();
    }

    function getJsonData(names, data) {
        let key,
            value = data[names[0].slice(1)];

        if (names.length > 1) {
            for (let i = 1; i < names.length; i++) {
                key = names[i].slice(1);
                value = $.isArray(value) ? value[0][key] : value[key];
            }
        }

        return value;
    }

    QorInlineEdit.prototype = {
        constructor: QorInlineEdit,

        init: function () {
            let $element = this.$element,
                saveButton = $element.data('button-save'),
                cancelButton = $element.data('button-cancel');

            this.TEMPLATE_SAVE = `<div class="qor-inlineedit__buttons">
                                        <button class="mdl-button mdl-button--colored mdl-js-button qor-button--small qor-inlineedit__cancel" type="button">${cancelButton}</button>
                                        <button class="mdl-button mdl-button--colored mdl-js-button qor-button--small qor-inlineedit__save" type="button">${saveButton}</button>
                                      </div>`;
            this.bind();
        },

        bind: function () {
            this.$element
                .on(EVENT_MOUSEENTER, CLASS_FIELD_SHOW, this.showEditButton)
                .on(EVENT_MOUSELEAVE, CLASS_FIELD_SHOW, this.hideEditButton)
                .on(EVENT_CLICK, CLASS_CANCEL, this.hideEdit)
                .on(EVENT_CLICK, CLASS_SAVE, this.saveEdit)
                .on(EVENT_CLICK, CLASS_EDIT, this.showEdit.bind(this));
        },

        unbind: function () {
            this.$element
                .off(EVENT_MOUSEENTER)
                .off(EVENT_MOUSELEAVE)
                .off(EVENT_CLICK);
        },

        showEditButton: function (e) {
            let $edit = $(QorInlineEdit.TEMPLATE_EDIT),
                $field = $(e.target).closest(CLASS_FIELD);

            if ($field.find('input:disabled, textarea:disabled,select:disabled').length) {
                return false;
            }

            $edit.appendTo($(this));
        },

        hideEditButton: function () {
            $('.qor-inlineedit__edit').remove();
        },

        showEdit: function (e) {
            let $parent = $(e.target)
                    .closest(CLASS_EDIT)
                    .hide()
                    .closest(CLASS_FIELD)
                    .addClass(CLASS_CONTAINER),
                $save = $(this.TEMPLATE_SAVE);

            $save.appendTo($parent);
        },

        hideEdit: function () {
            let $parent = $(this)
                .closest(CLASS_FIELD)
                .removeClass(CLASS_CONTAINER);
            $parent.find(CLASS_BUTTONS).remove();
        },

        saveEdit: function () {
            let $btn = $(this),
                $parent = $btn.closest(CLASS_FIELD),
                $form = $btn.closest('form'),
                $hiddenInput = $parent.closest('.qor-fieldset').find('input.qor-hidden__primary_key[type="hidden"]'),
                $input = $parent.find('input[name*="QorResource"],textarea[name*="QorResource"],select[name*="QorResource"]'),
                names = $input.length && $input.prop('name').match(/\.\w+/g),
                inputData = $input.serialize();

            if ($hiddenInput.length) {
                inputData = `${inputData}&${$hiddenInput.serialize()}`;
            }

            if (names.length) {
                $.ajax($form.prop('action'), {
                    method: $form.prop('method'),
                    data: inputData,
                    dataType: 'json',
                    beforeSend: function () {
                        $btn.prop('disabled', true);
                    },
                    success: function (data) {
                        let newValue = getJsonData(names, data),
                            $show = $parent.removeClass(CLASS_CONTAINER).find(CLASS_FIELD_SHOW),
                            $inner = $show.find(CLASS_FIELD_SHOW_INNER);

                        if ($inner.length) {
                            $inner.html(newValue);
                        } else {
                            $show.html(newValue);
                        }

                        $parent.find(CLASS_BUTTONS).remove();
                        $btn.prop('disabled', false);
                    },
                    error: function (err) {
                        window.QOR.handleAjaxError(err);
                        $btn.prop('disabled', false);
                    }
                });
            }
        },

        destroy: function () {
            this.unbind();
            this.$element.removeData(NAMESPACE);
        }
    };

    QorInlineEdit.DEFAULTS = {};

    QorInlineEdit.TEMPLATE_EDIT = `<button class="mdl-button mdl-js-button mdl-button--icon mdl-button--colored qor-inlineedit__edit" type="button"><i class="material-icons">mode_edit</i></button>`;

    QorInlineEdit.plugin = function (options) {
        return this.each(function () {
            var $this = $(this);
            var data = $this.data(NAMESPACE);
            var fn;

            if (!data) {
                $this.data(NAMESPACE, (data = new QorInlineEdit(this, options)));
            }

            if (typeof options === 'string' && $.isFunction((fn = data[options]))) {
                fn.call(data);
            }
        });
    };

    $(function () {
        let selector = '[data-toggle="qor.inlineEdit"]',
            options = {};

        $(document)
            .on(EVENT_DISABLE, function (e) {
                QorInlineEdit.plugin.call($(selector, e.target), 'destroy');
            })
            .on(EVENT_ENABLE, function (e) {
                QorInlineEdit.plugin.call($(selector, e.target), options);
            })
            .triggerHandler(EVENT_ENABLE);
    });

    return QorInlineEdit;
});
