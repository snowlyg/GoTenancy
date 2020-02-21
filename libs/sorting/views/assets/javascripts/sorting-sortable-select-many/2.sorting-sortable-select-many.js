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
        NAMESPACE = 'qor.chooser.sortable',
        EVENT_ENABLE = 'enable.' + NAMESPACE,
        EVENT_CLICK = 'click.' + NAMESPACE,
        EVENT_DISABLE = 'disable.' + NAMESPACE,
        CLASS_CHOSE_CONTAINER = '.select2-container',
        CLASS_CHOSE_INPUT = '.select2-search__field',
        CLASS_SORTABLE_BODY = '.qor-dragable',
        CLASS_SORTABLE = '.qor-dragable__list',
        CLASS_SORTABLE_HANDLE = '.qor-dragable__list-handle',
        CLASS_SORTABLE_DELETE = '.qor-dragable__list-delete',
        CLASS_SORTABLE_DATA = '.qor-dragable__list-data',
        CLASS_SORTABLE_BUTTON_ADD = '.qor-dragable__button-add',
        CLASS_BOTTOMSHEETS = '.qor-bottomsheets',
        CLASS_MANY = 'qor-bottomsheets__select-many',
        CLASS_SELECTED = 'is_selected',
        CLASS_SORTABLE_MANY = '[data-select-modal="many_sortable"]';

    function QorChooserSortable(element, options) {
        this.$element = $(element);
        this.options = $.extend({}, QorChooserSortable.DEFAULTS, $.isPlainObject(options) && options);
        this.init();
    }

    QorChooserSortable.prototype = {
        constructor: QorChooserSortable,

        init: function () {
            var $this = this.$element,
                select2Data = $this.data(),
                $parent = $this.parents(CLASS_SORTABLE_BODY),
                placeholderText = $this.data('placeholder'),
                self = this,
                option = {
                    minimumResultsForSearch: 8,
                    dropdownParent: $this.parent()
                };

            this.$selector = $parent.find(CLASS_SORTABLE_DATA);
            this.$sortableList = $parent.find(CLASS_SORTABLE);
            this.$parent = $parent;

            this.hasRemoteImage = select2Data.remoteImage;

            var sortEle = $parent.find(CLASS_SORTABLE)[0];

            this.sortable = window.Sortable.create(sortEle, {
                animation: 150,
                handle: CLASS_SORTABLE_HANDLE,
                filter: CLASS_SORTABLE_DELETE,
                dataIdAttr: 'data-index',

                onFilter: function (e) {
                    var $ele = $(e.item);

                    $ele.remove();
                    self.removeItemsFromList($ele.data());
                },
                onUpdate: function () {
                    self.renderOption();
                }
            });

            if (select2Data.remoteData) {
                let getSelect2AjaxDynamicURL = window.getSelect2AjaxDynamicURL;
                let remoteDataImage = select2Data.remoteImage;

                option.ajax = $.fn.select2.ajaxCommonOptions(select2Data);

                if (getSelect2AjaxDynamicURL && $.isFunction(getSelect2AjaxDynamicURL)) {
                    option.ajax.url = function () {
                        return getSelect2AjaxDynamicURL(select2Data);
                    };
                } else {
                    option.ajax.url = select2Data.remoteUrl;
                }

                option.templateResult = function (data) {
                    var tmpl = $this.parents('.qor-field').find('[name="select2-result-template"]');
                    return $.fn.select2.ajaxFormatResult(data, tmpl, remoteDataImage);
                };

                option.templateSelection = function (data) {
                    if (data.loading) return data.text;
                    var tmpl = $this.parents('.qor-field').find('[name="select2-selection-template"]');
                    return $.fn.select2.ajaxFormatResult(data, tmpl, remoteDataImage);
                };
            }

            if ($this.is('select')) {
                $this
                    .on('change', function () {
                        // setTimeout(function() {
                        //     $parent.find(CLASS_CHOSE).remove();
                        // }, 1);

                        $(CLASS_CHOSE_INPUT).attr('placeholder', placeholderText);
                    })
                    .on('select2:select', function (e) {
                        let data = e.params.data;
                        data.value = data.Name || data.text || data.Text || data.Title || data.Code;
                        self.addItems(data);
                    })
                    .on('select2:unselect', function (e) {
                        self.removeItems(e.params.data);
                    });

                $this.select2(option);

                $parent.find(CLASS_CHOSE_CONTAINER).hide();
                $(CLASS_CHOSE_INPUT).attr('placeholder', placeholderText);
            }

            this.bind();
        },

        bind: function () {
            this.$parent.on(EVENT_CLICK, CLASS_SORTABLE_BUTTON_ADD, this.show.bind(this));
            $(document).on(EVENT_CLICK, CLASS_SORTABLE_MANY, this.openSortable.bind(this));
        },

        unbind: function () {
            this.$parent.off(EVENT_CLICK, CLASS_SORTABLE_BUTTON_ADD, this.show);
            $(document).off(EVENT_CLICK, CLASS_SORTABLE_MANY, this.openSortable);
        },

        openSortable: function (e) {
            let data = $(e.target).data();

            this.BottomSheets = $body.data('qor.bottomsheets');
            this.selectedIconTmpl = $('[name="select-many-selected-icon"]').html();

            data.ingoreSubmit = true;
            data.url = data.selectListingUrl;

            if (data.selectDefaultCreating) {
                data.url = data.selectCreatingUrl;
            }

            this.BottomSheets.open(data, this.handleBottomSelect.bind(this));
        },

        show: function () {
            // $('[data-toggle="qor.chooser.sortable"]').find("option:selected[data-index='11']").prop("selected", false)

            let $container = this.$parent.find(CLASS_CHOSE_CONTAINER);
            let $ele = this.$element;
            let currentVal = $ele.val();
            let targetVal = [];
            let $li = this.$parent.find('.qor-dragable__list > li');

            if (currentVal.length) {
                $li.each(function () {
                    targetVal.push(String($(this).data('index')));
                });

                if (targetVal.length) {
                    currentVal.forEach(function (val) {
                        if (!targetVal.includes(val)) {
                            $ele.find(`option:selected[data-index='${val}']`).prop('selected', false);
                        }
                    });
                }
            }

            this.$element.trigger('change');

            $container.show();
            this.$parent.find(CLASS_SORTABLE_BUTTON_ADD).hide();
            setTimeout(function () {
                $container.find(CLASS_CHOSE_INPUT).click();
            }, 100);
        },

        handleBottomSelect: function () {
            var $bottomsheets = $(CLASS_BOTTOMSHEETS),
                options = {
                    onSelect: this.onSelectResults.bind(this), // render selected item after click item lists
                    onSubmit: this.onSubmitResults.bind(this) // render new items after new item form submitted
                };

            $bottomsheets.qorSelectCore(options).addClass(CLASS_MANY);
            this.initItems();
        },

        onSelectResults: function (data) {
            let $tr = data.$clickElement,
                $td = $tr.find('td:first'),
                obj = this.collectData(data);

            if (!$(CLASS_SORTABLE).find('li[data-index="' + obj.id + '"]').length) {
                this.addItems(obj);
                $tr.addClass(CLASS_SELECTED);
                $td.append(this.selectedIconTmpl);
            } else {
                this.removeItems(obj);
                $tr.removeClass(CLASS_SELECTED);
                $td.find('.qor-select__select-icon').remove();
            }
        },

        onSubmitResults: function (data) {
            this.addItems(this.collectData(data), true);
        },

        collectData: function (data) {
            // Handle data for sortable
            let remoteDataPrimaryKey = this.$element.data('remote-data-primary-key'),
                obj = {};

            obj.id = data[remoteDataPrimaryKey] || data.primaryKey || data.Id || data.ID;
            obj.value = data.Name || data.text || data.Text || data.Title || data.Code || obj.id;

            return obj;
        },

        initItems: function () {
            var $tr = $(CLASS_BOTTOMSHEETS).find('tbody tr'),
                selectedIconTmpl = this.selectedIconTmpl,
                selectedIDs = [],
                primaryKey,
                $selectedItems = this.$sortableList.find('[data-index]');

            $selectedItems.each(function () {
                selectedIDs.push($(this).data('index'));
            });

            $tr.each(function () {
                var $this = $(this),
                    $td = $this.find('td:first');

                primaryKey = $this.data().primaryKey;

                if (selectedIDs.indexOf(primaryKey) != '-1') {
                    $this.addClass(CLASS_SELECTED);
                    $td.append(selectedIconTmpl);
                }
            });
        },

        renderItem: function (data) {
            if (this.hasRemoteImage) {
                return window.Mustache.render(QorChooserSortable.LIST_HTML_WITH_IMAGE, data);
            } else {
                return window.Mustache.render(QorChooserSortable.LIST_HTML, data);
            }
        },

        renderOption: function () {
            var indexArr = this.sortable.toArray();
            var $selector = this.$parent.find(CLASS_SORTABLE_DATA);

            $selector.empty();

            window._.each(indexArr, function (id) {
                $selector.append(
                    window.Mustache.render(QorChooserSortable.OPTION_HTML, {
                        value: id
                    })
                );
            });
        },

        removeItems: function (data) {
            $(CLASS_SORTABLE)
                .find('li[data-index="' + data.id + '"]')
                .remove();
            this.renderOption();
        },

        removeItemsFromList: function (data) {
            this.renderOption();
            var index = data.index;
            var value = data.value;

            if (index && $(`.select2-selection__choice[item-id="${index}"]`).length) {
                $(`.select2-selection__choice[item-id="${index}"]`)
                    .find('.select2-selection__choice__remove')
                    .click();
            } else if (value && $(`.select2-selection__choice[title="${value}"]`).length) {
                $(`.select2-selection__choice[title="${value}"]`)
                    .find('.select2-selection__choice__remove')
                    .click();
            }
        },

        addItems: function (data, isNewData) {
            this.$sortableList.append(this.renderItem(data));
            this.renderOption();

            if (isNewData) {
                this.BottomSheets.hide();
            }
        },

        destroy: function () {
            this.sortable.destroy();
            this.unbind();
            this.$element.select2('destroy').removeData(NAMESPACE);
        }
    };

    QorChooserSortable.DEFAULTS = {};

    QorChooserSortable.LIST_HTML =
        '<li data-index="[[id]]" data-result-id="[[_resultId]]" data-value="[[value]]"><span>[[value]]</span><div><i class="material-icons qor-dragable__list-delete">clear</i><i class="material-icons qor-dragable__list-handle">drag_handle</i></div></li>';

    QorChooserSortable.LIST_HTML_WITH_IMAGE =
        '<li data-index="[[id]]" data-result-id="[[_resultId]]" data-value="[[value]]"><span><img src="[[Image]]">[[value]]</span><div><i class="material-icons qor-dragable__list-delete">clear</i><i class="material-icons qor-dragable__list-handle">drag_handle</i></div></li>';

    QorChooserSortable.OPTION_HTML = '<option selected value="[[value]]"></option>';

    QorChooserSortable.plugin = function (options) {
        return this.each(function () {
            var $this = $(this);
            var data = $this.data(NAMESPACE);
            var fn;

            if (!data) {
                if (/destroy/.test(options)) {
                    return;
                }

                $this.data(NAMESPACE, (data = new QorChooserSortable(this, options)));
            }

            if (typeof options === 'string' && $.isFunction((fn = data[options]))) {
                fn.apply(data);
            }
        });
    };

    $(function () {
        var selector = '[data-toggle="qor.chooser.sortable"]';

        $(document)
            .on(EVENT_DISABLE, function (e) {
                QorChooserSortable.plugin.call($(selector, e.target), 'destroy');
            })
            .on(EVENT_ENABLE, function (e) {
                QorChooserSortable.plugin.call($(selector, e.target));
            })
            .triggerHandler(EVENT_ENABLE);
    });

    return QorChooserSortable;
});
