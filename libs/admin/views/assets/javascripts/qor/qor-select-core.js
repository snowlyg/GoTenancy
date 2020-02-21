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

    let FormData = window.FormData,
        QOR = window.QOR,
        NAMESPACE = 'qor.selectcore',
        EVENT_SELECTCORE_BEFORESEND = 'selectcoreBeforeSend.' + NAMESPACE,
        EVENT_ONSELECT = 'afterSelected.' + NAMESPACE,
        EVENT_ONSUBMIT = 'afterSubmitted.' + NAMESPACE,
        EVENT_CLICK = 'click.' + NAMESPACE,
        EVENT_SUBMIT = 'submit.' + NAMESPACE,
        CLASS_TABLE = 'table.qor-js-table tr',
        CLASS_FORM = 'form';

    function QorSelectCore(element, options) {
        this.$element = $(element);
        this.options = $.extend({}, QorSelectCore.DEFAULTS, $.isPlainObject(options) && options);
        this.init();
    }

    QorSelectCore.prototype = {
        constructor: QorSelectCore,

        init: function () {
            this.bind();
        },

        bind: function () {
            this.$element.on(EVENT_CLICK, CLASS_TABLE, this.processingData.bind(this)).on(EVENT_SUBMIT, CLASS_FORM, this.submit.bind(this));
        },

        unbind: function () {
            this.$element.off(EVENT_CLICK, CLASS_TABLE).off(EVENT_SUBMIT, CLASS_FORM);
        },

        processingData: function (e) {
            let $this = $(e.target).closest('tr'),
                $bottomsheets = $this.closest('.qor-bottomsheets'),
                data = {},
                url,
                options = this.options,
                onSelect = options.onSelect,
                loading = options.loading;

            data = $.extend({}, data, $this.data());
            data.$clickElement = $this;

            url = data.mediaLibraryUrl || data.url;

            if (loading && $.isFunction(loading)) {
                loading($bottomsheets);
            }

            if (url) {

                $.getJSON(url, function (json) {
                    json.MediaOption && (json.MediaOption = JSON.parse(json.MediaOption));
                    data = $.extend({}, json, data);
                    if (onSelect && $.isFunction(onSelect)) {
                        onSelect(data, e);
                        $(document).trigger(EVENT_ONSELECT);
                    }
                }).always(function () {
                    $bottomsheets.find('.qor-media-loading').remove();
                });

            } else {
                if (onSelect && $.isFunction(onSelect)) {
                    onSelect(data, e);
                    $(document).trigger(EVENT_ONSELECT);
                }
            }
            return false;
        },

        submit: function (e) {
            let form = e.target,
                $form = $(form),
                _this = this,
                $submit = $form.find(':submit'),
                data,
                $loading = $(QOR.$formLoading),
                onSubmit = this.options.onSubmit;

            $(document).trigger(EVENT_SELECTCORE_BEFORESEND);

            $form.find('.qor-fieldset--new').remove();

            if (FormData) {
                e.preventDefault();

                $.ajax($form.prop('action'), {
                    method: $form.prop('method'),
                    data: new FormData(form),
                    dataType: 'json',
                    processData: false,
                    contentType: false,
                    beforeSend: function () {
                        $('.qor-submit-loading').remove();
                        $loading.appendTo($submit.prop('disabled', true).closest('.qor-form__actions')).trigger('enable.qor.material');
                    },
                    success: function (json) {
                        json.MediaOption && (json.MediaOption = JSON.parse(json.MediaOption));
                        data = json;
                        data.primaryKey = data.ID;

                        $('.qor-error').remove();

                        if (onSubmit && $.isFunction(onSubmit)) {
                            onSubmit(data, e);
                            $(document).trigger(EVENT_ONSUBMIT);
                        } else {
                            _this.refresh();
                        }
                    },
                    error: function (err) {
                        QOR.handleAjaxError(err);
                    },
                    complete: function () {
                        $submit.prop('disabled', false);
                    }
                });
            }
        },

        refresh: function () {
            setTimeout(function () {
                window.location.reload();
            }, 350);
        },

        destroy: function () {
            this.unbind();
        }
    };

    QorSelectCore.plugin = function (options) {
        return this.each(function () {
            let $this = $(this),
                data = $this.data(NAMESPACE),
                fn;

            if (!data) {
                if (/destroy/.test(options)) {
                    return;
                }
                $this.data(NAMESPACE, (data = new QorSelectCore(this, options)));
            }

            if (typeof options === 'string' && $.isFunction((fn = data[options]))) {
                fn.apply(data);
            }
        });
    };

    $.fn.qorSelectCore = QorSelectCore.plugin;

    return QorSelectCore;
});
