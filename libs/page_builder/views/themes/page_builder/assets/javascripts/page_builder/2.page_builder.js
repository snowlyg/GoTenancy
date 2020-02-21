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
        NAMESPACE = 'qor.pagebuilder',
        EVENT_ENABLE = 'enable.' + NAMESPACE,
        EVENT_CLICK = 'click.' + NAMESPACE,
        EVENT_DISABLE = 'disable.' + NAMESPACE,
        CLASS_CHOSE = '.select2-selection__choice',
        CLASS_CHOSE_CONTAINER = '.select2-container',
        CLASS_CHOSE_INPUT = '.select2-search__field',
        CLASS_SORTABLE_BODY = '.qor-dragable',
        CLASS_SORTABLE = '.qor-dragable__list',
        CLASS_SORTABLE_HANDLE = '.qor-dragable__list-handle',
        CLASS_SORTABLE_DELETE = '.qor-dragable__list-delete',
        CLASS_SORTABLE_DATA = '.qor-dragable__list-data',
        CLASS_SORTABLE_BUTTON_ADD = '.qor-dragable__button-add',
        CLASS_MANY = 'qor-bottomsheets__select-many qor-bottomsheets__pagebuilder',
        CLASS_SORTABLE_MANY = '[data-select-modal="pagebuilder_many_sortable"]';

    function QorPageBuilder(element, options) {
        this.$element = $(element);
        this.options = $.extend({}, QorPageBuilder.DEFAULTS, $.isPlainObject(options) && options);
        this.init();
    }

    QorPageBuilder.prototype = {
        constructor: QorPageBuilder,

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

            var sortEle = $parent.find(CLASS_SORTABLE)[0];

            this.sortable = window.Sortable.create(sortEle, {
                animation: 150,
                handle: CLASS_SORTABLE_HANDLE,
                filter: CLASS_SORTABLE_DELETE,
                dataIdAttr: 'data-index',

                onFilter: function (e) {
                    var $ele = $(e.item);
                    var eleIndex = $ele.data('index');

                    $ele.remove();
                    self.removeItemsFromList(eleIndex);
                },
                onUpdate: function () {
                    self.renderOption();
                }
            });

            if (select2Data.remoteData) {
                option.ajax = $.fn.select2.ajaxCommonOptions(select2Data);

                option.templateResult = function (data) {
                    var tmpl = $this.parents('.qor-field').find('[name="select2-result-template"]');
                    return $.fn.select2.ajaxFormatResult(data, tmpl);
                };

                option.templateSelection = function (data) {
                    if (data.loading) return data.text;
                    var tmpl = $this.parents('.qor-field').find('[name="select2-selection-template"]');
                    return $.fn.select2.ajaxFormatResult(data, tmpl);
                };
            }

            if ($this.is('select')) {
                $this
                    .on('change', function () {
                        setTimeout(function () {
                            $parent.find(CLASS_CHOSE).hide();
                        }, 1);

                        $(CLASS_CHOSE_INPUT).attr('placeholder', placeholderText);
                    })
                    .on('select2:select', function (e) {
                        self.addItems(e.params.data);
                    })
                    .on('select2:unselect', function (e) {
                        self.removeItems(e.params.data);
                    });

                $this.select2(option);

                $parent.find(CLASS_CHOSE_CONTAINER).hide();
                $parent.find(CLASS_CHOSE).hide();
                $(CLASS_CHOSE_INPUT).attr('placeholder', placeholderText);
            }

            this.bind();
        },

        bind: function () {
            this.$parent.on(EVENT_CLICK, CLASS_SORTABLE_BUTTON_ADD, this.show.bind(this)).on(EVENT_CLICK, CLASS_SORTABLE_MANY, this.openSortable.bind(this));
        },

        unbind: function () {
            this.$parent.off(EVENT_CLICK);
        },

        openSortable: function (e) {
            let data = $(e.target).data();

            this.BottomSheets = $body.data('qor.bottomsheets');
            this.selectListingUrl = data.selectListingUrl;
            data.ingoreSubmit = true;

            data.url = data.selectListingUrl;

            if (data.selectDefaultCreating) {
                data.url = data.selectCreatingUrl;
            }

            this.BottomSheets.open(data, this.handleBottomSelect.bind(this));
        },

        show: function () {
            let $container = this.$parent.find(CLASS_CHOSE_CONTAINER);

            $container.show();
            this.$parent.find(CLASS_SORTABLE_BUTTON_ADD).hide();
            setTimeout(function () {
                $container.find(CLASS_CHOSE_INPUT).click();
            }, 100);
        },

        handleBottomSelect: function ($bottomsheets) {
            let options = {
                onSelect: this.onSelectResults.bind(this), // render selected item after click item lists
                onSubmit: this.onSubmitResults.bind(this) // render new items after new item form submitted
            };

            $bottomsheets.qorSelectCore(options).addClass(CLASS_MANY);
            this.$bottomsheets = $bottomsheets;
            this.initTab();
        },

        onSelectResults: function (data) {
            this.addItems(this.collectData(data));
        },

        onSubmitResults: function (data) {
            this.addItems(this.collectData(data), true);
        },

        collectData: function (data) {
            // Handle data for sortable
            let remoteDataPrimaryKey = this.$element.data('remote-data-primary-key'),
                obj = $.extend({}, data);

            obj.SortableID = data[remoteDataPrimaryKey] || data.primaryKey || data.Id || data.ID;
            obj.SortableValue = data.Name || data.text || data.Text || data.Title || data.Code || obj.SortableID;

            return obj;
        },

        renderHint: function (data) {
            return window.Mustache.render($('[name="qor-pagebuilder-hint"]').html(), data);
        },

        initTab: function () {
            let data = this.$element.data(),
                $bottomsheets = this.$bottomsheets,
                $bottomsheetsBody = $bottomsheets.find('.qor-bottomsheets__body');

            $bottomsheets.addClass('has-tab');
            $bottomsheets.find('.qor-bottomsheets__title').html(data.selectTitle);
            $bottomsheetsBody.html('');
            $bottomsheetsBody.append(
                `<ul class="qor-bottomsheets__tab clearfix">
                    <li class="is-active" data-tab-type="create" data-tab-url="${data.selectCreatingUrl}">${data.selectCreatingTitle}</li>
                    <li data-tab-url="${data.selectListingUrl}" data-tab-type="list">${data.selectListingTitle}</li>
                </ul>
                <div class="qor-bottomsheets__tab-content"></div>`
            );

            $bottomsheets.on(EVENT_CLICK, '.qor-bottomsheets__tab li', this.switchResource.bind(this));
            $bottomsheets.find('.is-active').click();
        },

        switchResource: function (e) {
            let $target = $(e.target),
                url = $target.data('tab-url'),
                $bottomsheets = this.$bottomsheets;

            $bottomsheets.find('.qor-bottomsheets__tab li').removeClass('is-active');
            $target.addClass('is-active');

            $.ajax(url, {
                method: 'GET',
                dataType: 'html',
                success: function (html) {
                    $bottomsheets.find('.qor-bottomsheets__tab-content').attr('content-type', $target.data('tab-type'));
                    $bottomsheets
                        .find('.qor-bottomsheets__tab-content')
                        .html(
                            $(html)
                                .find('.mdl-layout__content.qor-page')
                                .html()
                        )
                        .trigger('enable');
                }
            });
        },

        renderItem: function (data) {
            return window.Mustache.render(QorPageBuilder.LIST_HTML, data);
        },

        renderOption: function () {
            var indexArr = this.sortable.toArray();
            var $selector = this.$parent.find(CLASS_SORTABLE_DATA);

            $selector.empty();

            window._.each(indexArr, function (id) {
                $selector.append(
                    window.Mustache.render(QorPageBuilder.OPTION_HTML, {
                        value: id
                    })
                );
            });
        },

        removeItems: function (data) {
            this.$parent
                .find(CLASS_SORTABLE)
                .find('li[data-index="' + data.id + '"]')
                .remove();
            this.renderOption();
        },

        removeItemsFromList: function () {
            this.renderOption();
        },

        addItems: function (data, isNewData) {
            if (!data.PreviewIcon) {
                data.PreviewIcon = data.$clickElement
                    .find('.qor-preview-icon')
                    .parent()
                    .html();
            }

            data.EditUrl = `${this.selectListingUrl}/${data.SortableID}`;

            this.$sortableList.append(this.renderItem(data));
            this.renderOption();

            if (isNewData) {
                this.$bottomsheets.find('.qor-widget__cancel').click();
            }

            this.$bottomsheets.remove();
            if (!$('.qor-bottomsheets').is(':visible')) {
                $('body').removeClass('qor-bottomsheets-open');
            }
        },

        destroy: function () {
            let $element = this.$element;
            this.sortable.destroy();
            this.unbind();

            if ($element.is('select')) {
                $element.select2('destroy');
            }

            $element.removeData(NAMESPACE);
        }
    };

    QorPageBuilder.DEFAULTS = {};

    QorPageBuilder.LIST_HTML = `<li data-index="[[SortableID]]" data-value="[[SortableValue]]">
                                    [[#PreviewIcon]][[&PreviewIcon]][[/PreviewIcon]]
                                    <span>
                                        <a href="[[EditUrl]]" data-url="[[EditUrl]]" data-ajax-mute="true" data-bottomsheet-classname="qor_pagebuilder--edit_widget">[[SortableValue]]</a></span>
                                    <div><i class="material-icons qor-dragable__list-delete">clear</i><i class="material-icons qor-dragable__list-handle">drag_handle</i></div>
                                </li>`;

    QorPageBuilder.OPTION_HTML = '<option selected value="[[value]]"></option>';

    QorPageBuilder.plugin = function (options) {
        return this.each(function () {
            var $this = $(this);
            var data = $this.data(NAMESPACE);
            var fn;

            if (!data) {
                if (/destroy/.test(options)) {
                    return;
                }

                $this.data(NAMESPACE, (data = new QorPageBuilder(this, options)));
            }

            if (typeof options === 'string' && $.isFunction((fn = data[options]))) {
                fn.apply(data);
            }
        });
    };

    $(function () {
        var selector = '[data-toggle="qor.pagebuilder"]';

        $(document)
            .on(EVENT_DISABLE, function (e) {
                QorPageBuilder.plugin.call($(selector, e.target), 'destroy');
            })
            .on(EVENT_ENABLE, function (e) {
                QorPageBuilder.plugin.call($(selector, e.target));
            })
            .triggerHandler(EVENT_ENABLE);
    });

    return QorPageBuilder;
});
