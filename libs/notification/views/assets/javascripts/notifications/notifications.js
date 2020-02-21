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

    let NAMESPACE = 'qor.notification',
        EVENT_ENABLE = 'enable.' + NAMESPACE,
        EVENT_DISABLE = 'disable.' + NAMESPACE,
        EVENT_CLICK = 'click.' + NAMESPACE,
        EVENT_UNDO = 'undo.qor.action',
        UNDO_TYPE = 'notification',
        UNDO_CONTAINER = '.qor-notifications__item',
        BUTTON_UNDO = '.qor-notifications__item-undo';

    function QorNotification(element, options) {
        this.$element = $(element);
        this.options = $.extend({}, QorNotification.DEFAULTS, $.isPlainObject(options) && options);
        this.init();
    }

    QorNotification.prototype = {
        constructor: QorNotification,

        init: function () {
            this.bind();
        },

        bind: function () {
            this.$element.on(EVENT_UNDO, $.proxy(this.undo, this));
            this.$element.on(EVENT_CLICK, '.qor-notification__load-more', this.load_more);
        },

        load_more: function (e) {
            let $target = $(e.target);

            $target.text($target.data('loading'));

            $.get($target.attr('href'), function (data) {
                var content = $(data).find('.qor-notifications');
                if ($('.qor-notifications--archived').length > 0) {
                    content.find('.qor-notifications--archived').remove();
                }
                $target.replaceWith(content.html());
            });
            return false;
        },

        undo: function (e, $actionButton, isUndo, data) {
            var actionData = $actionButton.data(),
                undoType = actionData.undoType,
                $undoContainer = $actionButton.closest(UNDO_CONTAINER),
                $item = $undoContainer.length ? $undoContainer : $actionButton.closest(BUTTON_UNDO).prev(UNDO_CONTAINER),
                $template;

            data.undoLabel = actionData.undoLabel;
            $template = $(window.Mustache.render(QorNotification.UNDO_HTML, data));
            !isUndo && $template.find('button').data(actionData);

            if (undoType == UNDO_TYPE) {
                $item.before(data.notification);
                !isUndo ? $item.after($template) : $item.next(BUTTON_UNDO).remove();
                $item.remove();
            }
        },

        unbind: function () {
            this.$element.off(EVENT_UNDO, this.undo).off(EVENT_CLICK);
        },

        destroy: function () {
            this.unbind();
            this.$element.removeData(NAMESPACE);
        }
    };

    QorNotification.DEFAULTS = {};
    QorNotification.UNDO_HTML = `<div class="qor-notifications__item-undo">
        <span>[[message]]</span>
        <button class="mdl-button mdl-button--colored is_undo" type="button" data-method="PUT" data-datatype="json" data-ajax-form="true" data-undo-type="notification">[[undoLabel]]</button>
        </div>`;

    QorNotification.plugin = function (options) {
        return this.each(function () {
            var $this = $(this);
            var data = $this.data(NAMESPACE);
            var fn;

            if (!data) {
                if (/destroy/.test(options)) {
                    return;
                }

                $this.data(NAMESPACE, (data = new QorNotification(this, options)));
            }

            if (typeof options === 'string' && $.isFunction((fn = data[options]))) {
                fn.apply(data);
            }
        });
    };

    $(function () {
        var selector = '.qor-notifications';

        $(document)
            .on(EVENT_DISABLE, function (e) {
                QorNotification.plugin.call($(selector, e.target), 'destroy');
            })
            .on(EVENT_ENABLE, function (e) {
                QorNotification.plugin.call($(selector, e.target));
            })
            .triggerHandler(EVENT_ENABLE);
    });

    return QorNotification;
});
