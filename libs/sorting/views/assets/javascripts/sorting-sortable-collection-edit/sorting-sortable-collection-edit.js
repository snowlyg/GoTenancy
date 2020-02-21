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

    let NAMESPACE = 'qor.collection.sortable',
        EVENT_ENABLE = 'enable.' + NAMESPACE,
        EVENT_DISABLE = 'disable.' + NAMESPACE,
        EVENT_CLICK = 'click.' + NAMESPACE,
        CLASS_ITEM = '.qor-sortable__item',
        CLASS_CHILDREN_ITEM = '> .qor-field__block > .qor-sortable__item',
        CLASS_BUTTON_CHANGE = '>.qor-sortable__button > .qor-sortable__button-change',
        CLASS_BUTTON_DONE = '>.qor-sortable__button > .qor-sortable__button-done',
        CLASS_BUTTON_ADD = '> .qor-field__block > .qor-sortable__button-add',
        CLASS_BUTTON_DELETE = '> .qor-sortable__button-delete',
        CLASS_BUTTON_MOVE = '>.qor-field__block > .qor-sortable__item > .qor-sortable__change .qor-sortable__button-move',
        CLASS_ACTION = '> .qor-sortable__change > .qor-sortable__action',
        CLASS_ACTION_POSITION = '.qor-sortable__action-position',
        CLASS_ITEM_FILTER = '.is-delete,.qor-fieldset--new',
        IS_LOADED = 'sortable-collection-loaded';

    function moveArrPos(arr, from, to) {
        if (to >= arr.length) {
            var k = to - arr.length;
            while (k-- + 1) {
                arr.push(undefined);
            }
        }
        arr.splice(to, 0, arr.splice(from, 1)[0]);
        return arr;
    }

    function QorCollectionSortable(element, options) {
        this.$element = $(element);
        this.options = $.extend({}, QorCollectionSortable.DEFAULTS, $.isPlainObject(options) && options);
        this.init();
    }

    QorCollectionSortable.prototype = {
        constructor: QorCollectionSortable,

        init: function () {
            this.bind();
            this.initItemOrder();
        },

        bind: function () {
            this.$element
                .on(EVENT_CLICK, CLASS_BUTTON_MOVE, this.moveItem.bind(this))
                .on(EVENT_CLICK, CLASS_BUTTON_DONE, this.finishMoveItem.bind(this))
                .on(EVENT_CLICK, CLASS_BUTTON_CHANGE, this.startMoveItem.bind(this));

            $(document).on('added.qor.replicator.sortableCollection', this.resetOrderArr.bind(this));
        },

        unbind: function () {
            this.$element
                .off(EVENT_CLICK, CLASS_BUTTON_MOVE)
                .off(EVENT_CLICK, CLASS_BUTTON_DONE)
                .off(EVENT_CLICK, CLASS_BUTTON_CHANGE);
            $(document).off('added.qor.replicator.sortableCollection');
        },

        initItemOrder: function (resetResource) {
            let $item = this.$element.find(CLASS_CHILDREN_ITEM).not(CLASS_ITEM_FILTER);

            // return false if no item
            if (!$item.length) {
                return;
            }

            this.$item = $item;

            let $select = $item.find(CLASS_ACTION).find(CLASS_ACTION_POSITION),
                orderData = {},
                orderArr = [],
                itemTotal = $item.length,
                template = $item.first().html(),
                fullResourceName,
                resourceNamePrefix;

            if ($select.length) {
                $select.remove();
            }

            if (!resourceNamePrefix) {
                fullResourceName = template.match(/(\w+)\="(\S*\[\d+\]\S*)"/); // get results : [attribute, name, value]
                if (fullResourceName.length) {
                    fullResourceName = fullResourceName[2];
                    resourceNamePrefix = fullResourceName.match(/^(\S*)\[(\d+)\]([^\[\]]*)$/); // get results : [input, prefix, index, suffix]

                    if (resourceNamePrefix.length) {
                        resourceNamePrefix = resourceNamePrefix[1];
                    }
                }
            }

            $item.each(function () {
                let $this = $(this),
                    $action = $this.find(CLASS_ACTION),
                    itemIndex = parseInt($this.attr('order-index'));

                orderData.itemTotal = itemTotal;
                orderData.itemIndex = itemIndex;
                orderArr.push(itemIndex);

                $action.prepend('<select class="qor-sortable__action-position"></select>');

                for (let i = 1; i <= itemTotal; i++) {
                    let renderData = {},
                        isSelected;

                    itemIndex + 1 == i ? (isSelected = true) : (isSelected = false);
                    renderData = $.extend({}, orderData, {selectorPosition: i, isSelected: isSelected});

                    $action.find('select').append(window.Mustache.render(QorCollectionSortable.OPTION_HTML, renderData));
                }

                // reset form resource name prop
                if (resetResource) {
                    let resourceName,
                        newPosition,
                        resourceNameEnd,
                        hasPrefixPosition = /\[\d+\]/.test(resourceNamePrefix),
                        $resource = $(this).find('[name^="' + resourceNamePrefix + '"]');

                    if (!$resource.length) {
                        return;
                    }

                    $resource.each(function () {
                        resourceName = $(this).prop('name');
                        resourceNameEnd = resourceName.match(/\.\w+$/);
                        newPosition = `[${itemIndex}]`;

                        if (hasPrefixPosition) {
                            // fullResourceName = QorResource.SerializableMeta.Sections[0].Items[0(this number is the position)].Name
                            // resourceNamePrefix = QorResource.SerializableMeta.Sections[0].Items
                            // change the last [0].Name to real position
                            resourceName = resourceName.replace(/\[\d+\]\.\w+$/, `${newPosition}${resourceNameEnd}`);
                        } else {
                            resourceName = resourceName.replace(/\[\d+\]/, newPosition);
                        }

                        $(this).prop('name', resourceName);
                    });
                }

                $this.data(orderData);
            });
            this.itemOrderArr = this.itemOrderArr || orderArr;
        },

        resetOrderArr: function (e, $item) {
            let orderIndex = parseInt($item.attr('order-index'));
            if (!orderIndex) {
                return;
            }

            this.itemOrderArr = this.itemOrderArr || [];
            this.itemOrderArr.push(orderIndex);
            this.$item = this.$element.find(CLASS_CHILDREN_ITEM).not(CLASS_ITEM_FILTER);
        },

        moveItem: function (e) {
            let $current = $(e.target).closest(CLASS_ITEM),
                currentPosition = $current.data().itemIndex,
                targetPosition = parseInt($current.find(CLASS_ACTION_POSITION).val()) - 1,
                $item = this.$item,
                newPosArr = [];

            if (targetPosition == currentPosition) {
                return;
            }

            newPosArr = moveArrPos(this.itemOrderArr, currentPosition, targetPosition);

            newPosArr.forEach(function (arr) {
                let index = newPosArr.indexOf(arr);

                $item
                    .filter(`[order-item="item_${arr}"]`)
                    .attr('order-index', index)
                    .css('order', index);
            });

            this.itemOrderArr = newPosArr;
            this.initItemOrder(true);
        },

        finishMoveItem: function (e) {
            let $target = $(e.target),
                $element = this.$element,
                $item = $element.find(CLASS_CHILDREN_ITEM).not(CLASS_ITEM_FILTER),
                $addButton = $element.find(CLASS_BUTTON_ADD),
                $deleteButton = $item.find(CLASS_BUTTON_DELETE),
                $actionButtons = $item.find(CLASS_ACTION);

            $target.hide();
            $actionButtons.hide();

            $element.find(CLASS_BUTTON_CHANGE).show();
            $addButton.show();
            $deleteButton.show();
        },

        startMoveItem: function (e) {
            let $target = $(e.target),
                $element = this.$element,
                $item = $element.find(CLASS_CHILDREN_ITEM).not(CLASS_ITEM_FILTER),
                $addButton = $element.find(CLASS_BUTTON_ADD),
                $deleteButton = $item.find(CLASS_BUTTON_DELETE),
                $actionButtons = $item.find(CLASS_ACTION);

            if (!$item.length) {
                return false;
            }

            $target.hide();

            $element.find(CLASS_BUTTON_DONE).show();
            $actionButtons.show();
            $addButton.hide();
            $deleteButton.hide();
            this.initItemOrder();
        },

        destroy: function () {
            this.unbind();
            this.$element.removeData(NAMESPACE);
        }
    };

    QorCollectionSortable.DEFAULTS = {};

    QorCollectionSortable.OPTION_HTML = '<option value="[[selectorPosition]]" [[#isSelected]]selected[[/isSelected]]>[[selectorPosition]] of [[itemTotal]]</option>';

    QorCollectionSortable.plugin = function (options) {
        return this.each(function () {
            var $this = $(this);
            var data = $this.data(NAMESPACE);
            var fn;

            if (!data) {
                if (/destroy/.test(options)) {
                    return;
                }

                $this.data(NAMESPACE, (data = new QorCollectionSortable(this, options)));
            }

            if (typeof options === 'string' && $.isFunction((fn = data[options]))) {
                fn.apply(data);
            }
        });
    };

    $(function () {
        var selector = '[data-toggle="qor.collection.sortable"]';

        if ($('body').data(IS_LOADED)) {
            return;
        }
        $('body').data(IS_LOADED, true);

        $(document)
            .on(EVENT_DISABLE, function (e) {
                QorCollectionSortable.plugin.call($(selector, e.target), 'destroy');
            })
            .on(EVENT_ENABLE, function (e) {
                QorCollectionSortable.plugin.call($(selector, e.target));
            })
            .triggerHandler(EVENT_ENABLE);
    });

    return QorCollectionSortable;
});
