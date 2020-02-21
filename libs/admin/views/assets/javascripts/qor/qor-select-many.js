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
        Mustache = window.Mustache,
        NAMESPACE = 'qor.selectone',
        PARENT_NAMESPACE = 'qor.bottomsheets',
        EVENT_CLICK = 'click.' + NAMESPACE,
        EVENT_ENABLE = 'enable.' + NAMESPACE,
        EVENT_DISABLE = 'disable.' + NAMESPACE,
        EVENT_RELOAD = 'reload.' + PARENT_NAMESPACE,
        CLASS_CLEAR_SELECT = '.qor-selected-many__remove',
        CLASS_UNDO_DELETE = '.qor-selected-many__undo',
        CLASS_DELETED_ITEM = 'qor-selected-many__deleted',
        CLASS_SELECT_FIELD = '.qor-field__selected-many',
        CLASS_SELECT_INPUT = '.qor-field__selectmany-input',
        CLASS_SELECT_ICON = '.qor-select__select-icon',
        CLASS_SELECT_HINT = '.qor-selectmany__hint',
        CLASS_SELECT_ALL = '.qor-selectmany__selectall',
        CLASS_PARENT = '.qor-field__selectmany',
        CLASS_SELECTED = 'is_selected',
        CLASS_MANY = 'qor-bottomsheets__select-many';

    function QorSelectMany(element, options) {
        this.$element = $(element);
        this.options = $.extend({}, QorSelectMany.DEFAULTS, $.isPlainObject(options) && options);
        this.init();
    }

    QorSelectMany.prototype = {
        constructor: QorSelectMany,

        init: function () {
            this.bind();
        },

        bind: function () {
            $document.on(EVENT_RELOAD, `.${CLASS_MANY}`, this.reloadData.bind(this));

            this.$element
                .on(EVENT_CLICK, CLASS_CLEAR_SELECT, this.clearSelect.bind(this))
                .on(EVENT_CLICK, '[data-select-modal="many"]', this.openBottomSheets.bind(this))
                .on(EVENT_CLICK, CLASS_UNDO_DELETE, this.undoDelete.bind(this));
        },

        unbind: function () {
            $document.off(EVENT_CLICK, '[data-select-modal="many"]').off(EVENT_RELOAD, `.${CLASS_MANY}`);
            this.$element.off(EVENT_CLICK, CLASS_CLEAR_SELECT).off(EVENT_CLICK, CLASS_UNDO_DELETE);
        },

        clearSelect: function (e) {
            var $target = $(e.target),
                $selectFeild = $target.closest(CLASS_PARENT);

            $target.closest('[data-primary-key]').addClass(CLASS_DELETED_ITEM);
            this.updateSelectInputData($selectFeild);

            return false;
        },

        undoDelete: function (e) {
            var $target = $(e.target),
                $selectFeild = $target.closest(CLASS_PARENT);

            $target.closest('[data-primary-key]').removeClass(CLASS_DELETED_ITEM);
            this.updateSelectInputData($selectFeild);

            return false;
        },

        openBottomSheets: function (e) {
            let $this = $(e.target),
                data = $this.data();

            this.BottomSheets = $body.data('qor.bottomsheets');
            this.bottomsheetsData = data;

            this.$selector = data.selectId ? $(data.selectId) : $this.closest(CLASS_PARENT).find('select');
            this.$selectFeild = this.$selector.closest(CLASS_PARENT).find(CLASS_SELECT_FIELD);

            // select many templates
            this.SELECT_MANY_SELECTED_ICON = $('[name="select-many-selected-icon"]').html();
            this.SELECT_MANY_UNSELECTED_ICON = $('[name="select-many-unselected-icon"]').html();
            this.SELECT_MANY_HINT = $('[name="select-many-hint"]').html();
            this.SELECT_MANY_TEMPLATE = $('[name="select-many-template"]').html();

            data.url = data.selectListingUrl;

            if (data.selectDefaultCreating) {
                data.url = data.selectCreatingUrl;
            }

            this.BottomSheets.open(data, this.handleSelectMany.bind(this));
        },

        reloadData: function () {
            this.initItems();
        },

        renderSelectMany: function (data) {
            return Mustache.render(this.SELECT_MANY_TEMPLATE, data);
        },

        renderHint: function (data) {
            return Mustache.render(this.SELECT_MANY_HINT, data);
        },

        initItems: function () {
            var $tr = this.$bottomsheets.find('tbody tr'),
                selectedIconTmpl = this.SELECT_MANY_SELECTED_ICON,
                unSelectedIconTmpl = this.SELECT_MANY_UNSELECTED_ICON,
                selectedIDs = [],
                primaryKey,
                $selectedItems = this.$selectFeild.find('[data-primary-key]').not('.' + CLASS_DELETED_ITEM);

            $selectedItems.each(function () {
                selectedIDs.push($(this).data().primaryKey);
            });

            $tr.each(function () {
                var $this = $(this),
                    $td = $this.find('td:first');

                primaryKey = $this.data().primaryKey;

                if (selectedIDs.indexOf(primaryKey) != '-1') {
                    $this.addClass(CLASS_SELECTED);
                    $td.append(selectedIconTmpl);
                } else {
                    $td.append(unSelectedIconTmpl);
                }
            });

            this.updateHint(this.getSelectedItemData());
        },

        getSelectedItemData: function () {
            var selecedItems = this.$selectFeild.find('[data-primary-key]').not('.' + CLASS_DELETED_ITEM);
            return {
                selectedNum: selecedItems.length
            };
        },

        updateHint: function (data) {
            var template;

            $.extend(data, this.bottomsheetsData);
            template = this.renderHint(data);

            this.$bottomsheets.find(CLASS_SELECT_HINT).remove();
            this.$bottomsheets.find('.qor-page__body').before(template);
        },

        updateSelectInputData: function ($selectFeild) {
            var $selectList = $selectFeild ? $selectFeild : this.$selectFeild,
                $selectedItems = $selectList.find('[data-primary-key]').not('.' + CLASS_DELETED_ITEM),
                $selector = $selectFeild ? $selectFeild.find(CLASS_SELECT_INPUT) : this.$selector,
                $options = $selector.find('option'),
                $option,
                data,
                primaryKey;

            $options.prop('selected', false);

            $selectedItems.each(function () {
                primaryKey = $(this).data().primaryKey;
                $option = $options.filter('[value="' + primaryKey + '"]');

                if (!$option.length) {
                    data = {
                        primaryKey: primaryKey,
                        displayName: ''
                    };
                    $option = $(Mustache.render(QorSelectMany.SELECT_MANY_OPTION_TEMPLATE, data));
                    $selector.append($option);
                }

                $option.prop('selected', true);
            });
        },

        changeIcon: function ($ele, template) {
            $ele.find(CLASS_SELECT_ICON).remove();
            $ele.find('td:first').prepend(template);
        },

        removeItem: function (data) {
            var primaryKey = data.primaryKey;

            this.$selectFeild
                .find('[data-primary-key="' + primaryKey + '"]')
                .find(CLASS_CLEAR_SELECT)
                .click();
            this.changeIcon(data.$clickElement, this.SELECT_MANY_UNSELECTED_ICON);
        },

        addItem: function (data, isNewData) {
            var template = this.renderSelectMany(data),
                $option,
                $list = this.$selectFeild.find('[data-primary-key="' + data.primaryKey + '"]');

            if ($list.length) {
                if ($list.hasClass(CLASS_DELETED_ITEM)) {
                    $list.removeClass(CLASS_DELETED_ITEM);
                    this.updateSelectInputData();
                    this.changeIcon(data.$clickElement, this.SELECT_MANY_SELECTED_ICON);
                    return;
                } else {
                    return;
                }
            }

            this.$selectFeild.append(template);

            if (isNewData) {
                $option = $(Mustache.render(QorSelectMany.SELECT_MANY_OPTION_TEMPLATE, data));
                $option.appendTo(this.$selector);
                $option.prop('selected', true);
                this.$bottomsheets.remove();
                if (!$('.qor-bottomsheets').is(':visible')) {
                    $('body').removeClass('qor-bottomsheets-open');
                }
                return;
            }

            this.changeIcon(data.$clickElement, this.SELECT_MANY_SELECTED_ICON);
        },

        handleSelectMany: function ($bottomsheets) {
            let options = {
                onSelect: this.onSelectResults.bind(this), // render selected item after click item lists
                onSubmit: this.onSubmitResults.bind(this) // render new items after new item form submitted
            };

            $bottomsheets.qorSelectCore(options).addClass(CLASS_MANY);
            $bottomsheets.on(EVENT_CLICK, CLASS_SELECT_ALL, this.handleSelectAll.bind(this));
            this.$bottomsheets = $bottomsheets;
            this.initItems();
        },

        handleSelectAll: function () {
            let $trs = this.$bottomsheets.find('.qor-table tbody tr'),
                $trsHasUnselected = $trs.not('.is_selected');

            if ($trsHasUnselected.length) {
                $trsHasUnselected.click();
            } else {
                $trs.click();
            }
        },

        onSelectResults: function (data) {
            this.handleResults(data);
        },

        onSubmitResults: function (data) {
            this.handleResults(data, true);
        },

        handleResults: function (data, isNewData) {
            data.displayName = data.Text || data.Name || data.Title || data.Code || data[Object.keys(data)[0]];

            if (isNewData) {
                this.addItem(data, true);
                return;
            }

            var $element = data.$clickElement,
                isSelected;

            $element.toggleClass(CLASS_SELECTED);
            isSelected = $element.hasClass(CLASS_SELECTED);

            if (isSelected) {
                this.addItem(data);
            } else {
                this.removeItem(data);
            }

            this.updateHint(this.getSelectedItemData());
            this.updateSelectInputData();
        },

        destroy: function () {
            this.unbind();
            this.$element.removeData(NAMESPACE);
        }
    };

    QorSelectMany.SELECT_MANY_OPTION_TEMPLATE = '<option value="[[ primaryKey ]]" >[[ displayName ]]</option>';

    QorSelectMany.plugin = function (options) {
        return this.each(function () {
            var $this = $(this);
            var data = $this.data(NAMESPACE);
            var fn;

            if (!data) {
                if (/destroy/.test(options)) {
                    return;
                }

                $this.data(NAMESPACE, (data = new QorSelectMany(this, options)));
            }

            if (typeof options === 'string' && $.isFunction((fn = data[options]))) {
                fn.apply(data);
            }
        });
    };

    $(function () {
        var selector = '[data-toggle="qor.selectmany"]';
        $(document)
            .on(EVENT_DISABLE, function (e) {
                QorSelectMany.plugin.call($(selector, e.target), 'destroy');
            })
            .on(EVENT_ENABLE, function (e) {
                QorSelectMany.plugin.call($(selector, e.target));
            })
            .triggerHandler(EVENT_ENABLE);
    });

    return QorSelectMany;
});
