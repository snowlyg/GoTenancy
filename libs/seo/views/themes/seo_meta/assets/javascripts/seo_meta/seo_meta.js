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

    let NAMESPACE = 'qor.seo',
        EVENT_ENABLE = 'enable.' + NAMESPACE,
        EVENT_DISABLE = 'disable.' + NAMESPACE,
        EVENT_CLICK = 'click.' + NAMESPACE,
        EVENT_BLUR = 'blur.' + NAMESPACE,
        CLASS_SUBMIT = '.qor-seo-submit',
        CLASS_ADD_TAGS_NAME = '.qor-seo-tag',
        CLASS_TAGS_INPUT_NAME = '.qor-seo__settings input[type="text"],.qor-seo__settings textarea[name]:visible',
        CLASS_DEFAULT_INPUT = '.qor-seo__defaults-input',
        CLASS_SETTINGS = '.qor-seo__settings';

    function QorSeo(element, options) {
        this.$element = $(element);
        this.options = $.extend({}, QorSeo.DEFAULTS, $.isPlainObject(options) && options);
        this.focusedInputID = '';
        this.init();
    }

    QorSeo.prototype = {
        constructor: QorSeo,

        init: function () {
            let $element = this.$element;

            this.$addTgas = $element.find(CLASS_ADD_TAGS_NAME);
            this.bind();
        },

        bind: function () {
            this.$element
                .on(EVENT_CLICK, CLASS_DEFAULT_INPUT, this.toggleDefault.bind(this))
                .on(EVENT_CLICK, CLASS_SUBMIT, this.submitSeo.bind(this))
                .on('click keyup', CLASS_TAGS_INPUT_NAME, this.tagInputsFocus.bind(this))
                .on(EVENT_BLUR, CLASS_TAGS_INPUT_NAME, this.tagInputsBlur.bind(this))
                .on(EVENT_CLICK, CLASS_ADD_TAGS_NAME, this.addTags.bind(this));
        },

        unbind: function () {
            this.$element
                .off(EVENT_CLICK)
                .off('click keyup')
                .off(EVENT_BLUR);
        },

        toggleDefault: function () {
            this.$element.find(CLASS_SETTINGS).toggle();
        },

        tagInputsFocus: function () {
            this.$addTgas.addClass('focus');
            var $focusedInput = $(document.activeElement);

            this.focusedInputID = $focusedInput.prop('name');
            this.focusedInputStart = $focusedInput[0].selectionStart;
            this.focusedInputEnd = $focusedInput[0].selectionEnd;
            this.focusedInputVal = $focusedInput.val();
        },

        tagInputsBlur: function () {
            this.$addTgas.removeClass('focus');
            this.$focusedInputID = false;
        },

        addTags: function (e) {
            if (!this.focusedInputID) {
                return;
            }

            var newVal = '';
            var startString = this.focusedInputVal.substring(0, this.focusedInputStart);
            var endString = this.focusedInputVal.substring(this.focusedInputEnd, this.focusedInputVal.length);
            var tagVal = "{{' + $(e.currentTarget).data('tagValue') + '}}";

            newVal = startString + tagVal + endString;

            this.$element
                .find('[name="' + this.focusedInputID + '"]')
                .val(newVal)
                .focus();

            this.focusedInputID = false;
        },

        submitSeo: function () {
            var $element = this.$element,
                $form = $element.find('.qor-form');

            // new FormData(form)
            // "GlobalSetting": { "SiteName" : "Qor", "BrandName" : "ThePlant" }

            $('.qor-seo-alert').hide();

            $.ajax({
                method: 'POST',
                url: $form.attr('action'),
                data: new FormData($form[0]),
                processData: false,
                contentType: false,
                dataType: 'json',
                success: function () {
                    window.onbeforeunload = null;
                    $.fn.qorSlideoutBeforeHide = null;
                    $('.qor-seo-alert.qor-alert--success').show();
                    setTimeout(function () {
                        $('.qor-alert--success').hide();
                    }, 3000);
                },
                error: function () {
                    $('.qor-seo-alert.qor-alert--error').show();
                }
            });
            return false;
        },

        destroy: function () {
            this.unbind();
            this.$element.removeData(NAMESPACE);
        }
    };

    QorSeo.DEFAULTS = {};

    QorSeo.plugin = function (options) {
        return this.each(function () {
            var $this = $(this);
            var data = $this.data(NAMESPACE);
            var fn;

            if (!data) {
                if (/destroy/.test(options)) {
                    return;
                }

                $this.data(NAMESPACE, (data = new QorSeo(this, options)));
            }

            if (typeof options === 'string' && $.isFunction((fn = data[options]))) {
                fn.apply(data);
            }
        });
    };

    $(function () {
        var selector = '[data-toggle="qor.seo"]',
            options = {};

        $(document)
            .on('click.qor.fixedAlert', '[data-dismiss="fixed-alert"]', function () {
                $(this)
                    .closest('.qor-alert')
                    .hide();
            })
            .on(EVENT_DISABLE, function (e) {
                QorSeo.plugin.call($(selector, e.target), 'destroy');
            })
            .on(EVENT_ENABLE, function (e) {
                QorSeo.plugin.call($(selector, e.target), options);
            })
            .triggerHandler(EVENT_ENABLE);
    });

    return QorSeo;
});
