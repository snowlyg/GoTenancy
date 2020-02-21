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
        NAMESPACE = 'qor.sorting',
        EVENT_ENABLE = 'enable.' + NAMESPACE,
        EVENT_DISABLE = 'disable.' + NAMESPACE,
        EVENT_CHANGE = 'change.' + NAMESPACE,
        EVENT_CLICK = 'click.' + NAMESPACE,
        EVENT_MOUSE_DOWN = 'mousedown.' + NAMESPACE,
        EVENT_MOUSE_UP = 'mouseup.' + NAMESPACE,
        EVENT_DRAG_START = 'dragstart.' + NAMESPACE,
        EVENT_DRAG_END = 'dragend.' + NAMESPACE,
        EVENT_DRAG_OVER = 'dragover.' + NAMESPACE,
        EVENT_DROP = 'drop.' + NAMESPACE,
        CLASS_SORTING = 'qor-sorting',
        CLASS_HIGHLIGHT = 'qor-sorting__highlight',
        CLASS_HOVER = 'qor-sorting__hover',
        SELECTOR_TR = 'tbody > tr';

    function QorSorter(element, options) {
        this.$element = $(element);
        this.options = $.extend({}, QorSorter.DEFAULTS, $.isPlainObject(options) && options);
        this.$source = null;
        this.ascending = false;
        this.orderType = 0; // 0 -> mix, 1 -> ascending order, -1 -> descending order
        this.startY = 0;
        this.init();
    }

    QorSorter.prototype = {
        constructor: QorSorter,

        init: function () {
            let options = this.options,
                $this = this.$element,
                $rows = $this.find(SELECTOR_TR),
                orderType = 0,
                count = 0,
                index = 0,
                previousPosition;

            $('body').addClass(CLASS_SORTING);
            $this.find('tbody .qor-table__actions').append(QorSorter.TEMPLATE);

            $rows.each(function (i) {
                let position = $(this).find(options.input).data('position');

                if (i > 0) {
                    if (position > previousPosition) {
                        count++;
                    } else {
                        count--;
                    }
                }

                previousPosition = position;
                index = i;
            });

            if (count === index) {
                orderType = 1;
            } else if (-count === index) {
                orderType = -1;
            }

            this.$rows = $rows;
            this.orderType = orderType;
            this.bind();
        },

        bind: function () {
            let options = this.options;

            this.$element.on(EVENT_CLICK, options.input, function () {
                return false;
            }).on(EVENT_CHANGE, options.input, $.proxy(this.change, this)).on(EVENT_MOUSE_DOWN, options.toggle, $.proxy(this.mousedown, this)).on(EVENT_MOUSE_UP, $.proxy(this.mouseup, this)).on(EVENT_DRAG_START, SELECTOR_TR, $.proxy(this.dragstart, this)).on(EVENT_DRAG_END, SELECTOR_TR, $.proxy(this.dragend, this)).on(EVENT_DRAG_OVER, SELECTOR_TR, $.proxy(this.dragover, this)).on(EVENT_DROP, SELECTOR_TR, $.proxy(this.drop, this));
        },

        unbind: function () {
            this.$element.off(EVENT_CHANGE, this.change).off(EVENT_MOUSE_DOWN, this.mousedown);
        },

        change: function (e) {
            let options = this.options,
                orderType = this.orderType,
                $rows = this.$rows,
                $sourceInput = $(e.currentTarget),
                $source = $sourceInput.closest('tr'),
                $tbody = $source.parent(),
                source = $sourceInput.data(),
                sourcePosition = source.position,
                targetPosition = parseInt($sourceInput.val(), 10),
                largethan = targetPosition > sourcePosition,
                $target;

            e.stopPropagation();

            $rows.each(function () {
                let $this = $(this),
                    $input = $this.find(options.input),
                    position = $input.data('position');

                if (position === targetPosition) {
                    $target = $this;

                    if (largethan) {
                        if (orderType === 1) {
                            $target.after($source);
                        } else if (orderType === -1) {
                            $target.before($source);
                        }
                    } else {
                        if (orderType === 1) {
                            $target.before($source);
                        } else if (orderType === -1) {
                            $target.after($source);
                        }
                    }
                }

                if (largethan) {
                    if (position > sourcePosition && position <= targetPosition) {
                        $input.data('position', --position).val(position);
                    }
                } else {
                    if (position < sourcePosition && position >= targetPosition) {
                        $input.data('position', ++position).val(position);
                    }
                }
            });

            $sourceInput.data('position', targetPosition);

            if (!$target) {
                if (largethan) {
                    if (orderType === 1) {
                        $tbody.append($source);
                    } else if (orderType === -1) {
                        $tbody.prepend($source);
                    }
                } else {
                    if (orderType === 1) {
                        $tbody.prepend($source);
                    } else if (orderType === -1) {
                        $tbody.append($source);
                    }
                }
            }

            this.sort($source, {
                url: source.sortingUrl,
                from: sourcePosition,
                to: targetPosition
            });

            return false;
        },

        mousedown: function (e) {
            this.startY = e.pageY;
            $(e.currentTarget).closest('tr').prop('draggable', true);
        },

        mouseup: function () {
            this.$element.find(SELECTOR_TR).prop('draggable', false);
        },

        dragend: function () {
            $(SELECTOR_TR).removeClass(CLASS_HOVER);
            this.$element.find(SELECTOR_TR).prop('draggable', false);
        },

        dragstart: function (e) {
            let event = e.originalEvent,
                $target = $(e.currentTarget);

            if ($target.prop('draggable') && event.dataTransfer) {
                event.dataTransfer.effectAllowed = 'move';
                this.$source = $target;
            }
        },

        dragover: function (e) {
            let $source = this.$source;

            $(SELECTOR_TR).removeClass(CLASS_HOVER);
            $(e.currentTarget).prev('tr').addClass(CLASS_HOVER);

            if (!$source || e.currentTarget === this.$source[0]) {
                return;
            }

            e.preventDefault();
        },

        drop: function (e) {
            let options = this.options,
                orderType = this.orderType,
                movedown = e.pageY > this.startY,
                $source = this.$source,
                $sourceInput,
                $target,
                source,
                sourcePosition,
                targetPosition,
                largethan;

            $(SELECTOR_TR).removeClass(CLASS_HOVER);

            if (!$source || e.currentTarget === this.$source[0]) {
                return;
            }

            e.preventDefault();

            $target = $(e.currentTarget);

            $sourceInput = $source.find(options.input);
            source = $sourceInput.data();
            sourcePosition = source.position;
            targetPosition = $target.find(options.input).data('position');
            largethan = targetPosition > sourcePosition;


            this.$element.find(SELECTOR_TR).each(function () {
                let $input = $(this).find(options.input),
                    position = $input.data('position');

                if (largethan) {
                    if (position > sourcePosition && position <= targetPosition) {
                        $input.data('position', --position).val(position);
                    }
                } else {
                    if (position < sourcePosition && position >= targetPosition) {
                        $input.data('position', ++position).val(position);
                    }
                }
            });

            $sourceInput.data('position', targetPosition).val(targetPosition);

            if (largethan) {
                if (orderType === 1) {
                    $target.after($source);
                } else if (orderType === -1) {
                    $target.before($source);
                } else {
                    if (movedown) {
                        $target.after($source);
                    } else {
                        $target.before($source);
                    }
                }
            } else {
                if (orderType === 1) {
                    $target.before($source);
                } else if (orderType === -1) {
                    $target.after($source);
                } else {
                    if (movedown) {
                        $target.before($source);
                    } else {
                        $target.after($source);
                    }
                }
            }

            this.sort($source, {
                url: source.sortingUrl,
                from: sourcePosition,
                to: targetPosition
            });
        },

        sort: function ($row, data) {
            let options = this.options;

            if (data.url) {
                this.highlight($row);

                $.ajax(data.url, {
                    method: 'post',
                    data: {
                        from: data.from,
                        to: data.to
                    },
                    success: function (actualPosition, textStatus, xhr) {
                        if (xhr.status === 200) {
                            $row.find(options.input).data('position', actualPosition).val(actualPosition);
                        }
                    },
                    error: function (xhr, textStatus, errorThrown) {
                        if (xhr.status === 422) {
                            window.alert(xhr.responseText);
                        } else {
                            window.alert([textStatus, errorThrown].join(': '));
                        }
                    }
                });
            }
        },

        highlight: function ($row) {
            $row.addClass(CLASS_HIGHLIGHT);

            setTimeout(function () {
                $row.removeClass(CLASS_HIGHLIGHT);
            }, 2000);
        },

        destroy: function () {
            this.unbind();
            this.$element.removeData(NAMESPACE);
        }
    };

    QorSorter.DEFAULTS = {
        toggle: false,
        input: false
    };

    QorSorter.TEMPLATE = '<a class="qor-sorting__toggle"><i class="material-icons">drag_handle</i></a>';

    QorSorter.plugin = function (options) {
        return this.each(function () {
            let $this = $(this),
                data = $this.data(NAMESPACE),
                fn;

            if (!data) {
                if (/destroy/.test(options)) {
                    return;
                }

                $this.data(NAMESPACE, (data = new QorSorter(this, options)));
            }

            if (typeof options === 'string' && $.isFunction(fn = data[options])) {
                fn.apply(data);
            }
        });
    };

    $(function () {
        if (!/sorting\=true/.test(location.search)) {
            return;
        }

        let selector = '.qor-js-table',
            options = {
                toggle: '.qor-sorting__toggle',
                input: '.qor-sorting__position'
            };

        $(document).on(EVENT_DISABLE, function (e) {
            QorSorter.plugin.call($(selector, e.target), 'destroy');
        }).on(EVENT_ENABLE, function (e) {
            QorSorter.plugin.call($(selector, e.target), options);
        }).trigger('disable.qor.slideout').triggerHandler(EVENT_ENABLE);
    });

    return QorSorter;

});
