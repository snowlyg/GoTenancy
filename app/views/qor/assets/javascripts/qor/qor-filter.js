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

    let location = window.location,
        NAMESPACE = 'qor.filter',
        EVENT_FILTER_CHANGE = 'filterChanged.' + NAMESPACE,
        EVENT_ENABLE = 'enable.' + NAMESPACE,
        EVENT_DISABLE = 'disable.' + NAMESPACE,
        EVENT_CLICK = 'click.' + NAMESPACE,
        EVENT_CHANGE = 'change.' + NAMESPACE,
        CLASS_IS_ACTIVE = 'is-active',
        CLASS_BOTTOMSHEETS = '.qor-bottomsheets';

    function encodeSearch(data, detached) {
        var search = decodeURI(location.search);
        var per_page = location.search.match(/per_page=\d+/);
        var params;

        search = search.replace(/per_page=\d+/g, '').replace(/page=\d+/, 'page=1');


        if (per_page && per_page.length) {
            search = search + "&" + per_page[0];
        }


        if ($.isArray(data)) {
            params = decodeSearch(search);

            $.each(data, function (i, param) {
                i = $.inArray(param, params);

                if (i === -1) {
                    params.push(param);
                } else if (detached) {
                    params.splice(i, 1);
                }
            });

            search = '?' + params.join('&');
        }

        return search;
    }

    function decodeSearch(search) {
        var data = [];

        if (search && search.indexOf('?') > -1) {
            search = search.replace(/\+/g, ' ').split('?')[1];

            if (search && search.indexOf('#') > -1) {
                search = search.split('#')[0];
            }

            if (search) {
                // search = search.toLowerCase();
                data = $.map(search.split('&'), function (n) {
                    var param = [];
                    var value;

                    n = n.split('=');
                    value = n[1];
                    param.push(n[0]);

                    if (value) {
                        value = $.trim(decodeURIComponent(value));

                        if (value) {
                            param.push(value);
                        }
                    }

                    return param.join('=');
                });
            }
        }

        return data;
    }

    function QorFilter(element, options) {
        this.$element = $(element);
        this.options = $.extend({}, QorFilter.DEFAULTS, $.isPlainObject(options) && options);
        this.init();
    }

    QorFilter.prototype = {
        constructor: QorFilter,

        init: function () {
            // this.parse();
            this.bind();
        },

        bind: function () {
            var options = this.options;

            this.$element.on(EVENT_CLICK, options.label, $.proxy(this.toggle, this)).on(EVENT_CHANGE, options.group, $.proxy(this.toggle, this));
        },

        unbind: function () {
            this.$element.off(EVENT_CLICK, this.toggle).off(EVENT_CHANGE, this.toggle);
        },

        toggle: function (e) {
            let $target = $(e.currentTarget),
                data = [],
                params,
                param,
                search,
                name,
                value,
                index,
                matched,
                paramName;

            if ($target.is('select')) {
                params = decodeSearch(decodeURI(location.search));

                paramName = name = $target.attr('name');
                value = $target.val();
                param = [name];

                if (value) {
                    param.push(value);
                }

                param = param.join('=');

                if (value) {
                    data.push(param);
                }

                $target.children().each(function () {
                    var $this = $(this);
                    var param = [name];
                    var value = $.trim($this.prop('value'));

                    if (value) {
                        param.push(value);
                    }

                    param = param.join('=');
                    index = $.inArray(param, params);

                    if (index > -1) {
                        matched = param;
                        return false;
                    }
                });

                if (matched) {
                    data.push(matched);
                    search = encodeSearch(data, true);
                } else {
                    search = encodeSearch(data);
                }
            } else if ($target.is('a')) {
                e.preventDefault();
                paramName = $target.data().paramName;
                data = decodeSearch($target.attr('href'));
                if ($target.hasClass(CLASS_IS_ACTIVE)) {
                    search = encodeSearch(data, true); // set `true` to detach
                } else {
                    search = encodeSearch(data);
                }
            }

            if (this.$element.closest(CLASS_BOTTOMSHEETS).length) {
                $(CLASS_BOTTOMSHEETS).trigger(EVENT_FILTER_CHANGE, [search, paramName]);
            } else {
                location.search = search;
            }
        },

        destroy: function () {
            this.unbind();
            this.$element.removeData(NAMESPACE);
        }
    };

    QorFilter.DEFAULTS = {
        label: false,
        group: false
    };

    QorFilter.plugin = function (options) {
        return this.each(function () {
            var $this = $(this);
            var data = $this.data(NAMESPACE);
            var fn;

            if (!data) {
                if (/destroy/.test(options)) {
                    return;
                }

                $this.data(NAMESPACE, (data = new QorFilter(this, options)));
            }

            if (typeof options === 'string' && $.isFunction((fn = data[options]))) {
                fn.apply(data);
            }
        });
    };

    $(function () {
        var selector = '[data-toggle="qor.filter"]';
        var options = {
            label: 'a',
            group: 'select'
        };

        $(document)
            .on(EVENT_DISABLE, function (e) {
                QorFilter.plugin.call($(selector, e.target), 'destroy');
            })
            .on(EVENT_ENABLE, function (e) {
                QorFilter.plugin.call($(selector, e.target), options);
            })
            .triggerHandler(EVENT_ENABLE);
    });

    return QorFilter;
});
