(function (factory) {
    if (typeof define === "function" && define.amd) {
        // AMD. Register as anonymous module.
        define(["jquery"], factory);
    } else if (typeof exports === "object") {
        // Node / CommonJS
        factory(require("jquery"));
    } else {
        // Browser globals.
        factory(jQuery);
    }
})(function ($) {
    "use strict";

    let _ = window._,
        FormData = window.FormData,
        QOR_Translations = window.QOR_Translations,
        NAMESPACE = "qor.bottomsheets",
        EVENT_CLICK = "click." + NAMESPACE,
        EVENT_SUBMIT = "submit." + NAMESPACE,
        EVENT_SUBMITED = "ajaxSuccessed." + NAMESPACE,
        EVENT_RELOAD = "reload." + NAMESPACE,
        EVENT_RELOADFROMURL = "reloadFromUrl." + NAMESPACE,
        EVENT_BOTTOMSHEET_BEFORESEND = "bottomsheetBeforeSend." + NAMESPACE,
        EVENT_BOTTOMSHEET_LOADED = "bottomsheetLoaded." + NAMESPACE,
        EVENT_BOTTOMSHEET_CLOSED = "bottomsheetClosed." + NAMESPACE,
        EVENT_BOTTOMSHEET_SUBMIT = "bottomsheetSubmitComplete." + NAMESPACE,
        EVENT_HIDDEN = "hidden." + NAMESPACE,
        EVENT_KEYUP = "keyup." + NAMESPACE,
        CLASS_OPEN = "qor-bottomsheets-open",
        CLASS_IS_SHOWN = "is-shown",
        CLASS_IS_SLIDED = "is-slided",
        CLASS_MAIN_CONTENT = ".mdl-layout__content.qor-page",
        CLASS_BODY_CONTENT = ".qor-page__body",
        CLASS_BODY_HEAD = ".qor-page__header",
        CLASS_BOTTOMSHEETS_FILTER = ".qor-bottomsheet__filter",
        CLASS_BOTTOMSHEETS_BUTTON = ".qor-bottomsheets__search-button",
        CLASS_BOTTOMSHEETS_INPUT = ".qor-bottomsheets__search-input",
        URL_GETQOR = "http://www.getqor.com/";

    function getUrlParameter(name, search) {
        name = name.replace(/[\[]/, "\\[").replace(/[\]]/, "\\]");
        var regex = new RegExp("[\\?&]" + name + "=([^&#]*)");
        var results = regex.exec(decodeURIComponent(search));
        return results === null
            ? ""
            : results[1].replace(/\+/g, " ");
    }

    function updateQueryStringParameter(key, value, uri) {
        var escapedkey = String(key).replace(/[\\^$*+?.()|[\]{}]/g, "\\$&"),
            re = new RegExp("([?&])" + escapedkey + "=.*?(&|$)", "i"),
            separator = uri.indexOf("?") !== -1 ? "&" : "?";

        if (uri.match(re)) {
            if (value) {
                return uri.replace(re, "$1" + key + "=" + value + "$2");
            } else {
                if (RegExp.$1 === "?" || RegExp.$1 === RegExp.$2) {
                    return uri.replace(re, "$1");
                } else {
                    return uri.replace(re, "");
                }
            }
        } else if (value) {
            return uri + separator + key + "=" + value;
        }
    }

    function pushArrary($ele, isScript) {
        let array = [],
            prop = "href";

        isScript && (prop = "src");
        $ele.each(function () {
            array.push($(this).attr(prop));
        });
        return _.uniq(array);
    }

    function execSlideoutEvents(url, response) {
        // exec qorSliderAfterShow after script loaded
        var qorSliderAfterShow = $.fn.qorSliderAfterShow;
        for (var name in qorSliderAfterShow) {
            if (
                qorSliderAfterShow.hasOwnProperty(name) &&
                !qorSliderAfterShow[name]["isLoadedInBottomSheet"] &&
                name != "initPublishForm" &&
                name != "qorActivityinit"
            ) {
                qorSliderAfterShow[name]["isLoadedInBottomSheet"] = true;
                qorSliderAfterShow[name].call(this, url, response);
            }
        }
    }

    function loadScripts(srcs, data, callback) {
        let scriptsLoaded = 0;

        for (let i = 0, len = srcs.length; i < len; i++) {
            let script = document.createElement("script");

            script.onload = function () {
                scriptsLoaded++;

                if (scriptsLoaded === srcs.length) {
                    if ($.isFunction(callback)) {
                        callback();
                    }
                }

                if (data && data.url && data.response) {
                    execSlideoutEvents(data.url, data.response);
                }
            };

            script.src = srcs[i];
            document.body.appendChild(script);
        }
    }

    function loadStyles(srcs) {
        let ss = document.createElement("link"),
            src = srcs.shift();

        ss.type = "text/css";
        ss.rel = "stylesheet";
        ss.onload = function () {
            if (srcs.length) {
                loadStyles(srcs);
            }
        };
        ss.href = src;
        document.getElementsByTagName("head")[0].appendChild(ss);
    }

    function compareScripts($scripts) {
        let $currentPageScripts = $("script"),
            slideoutScripts = pushArrary($scripts, true),
            currentPageScripts = pushArrary($currentPageScripts, true),
            scriptDiff = _.difference(slideoutScripts, currentPageScripts);
        return scriptDiff;
    }

    function compareLinks($links) {
        let $currentStyles = $("link"),
            slideoutStyles = pushArrary($links),
            currentStyles = pushArrary($currentStyles),
            styleDiff = _.difference(slideoutStyles, currentStyles);

        return styleDiff;
    }

    function QorBottomSheets(element, options) {
        this.$element = $(element);
        this.options = $.extend(
            {},
            QorBottomSheets.DEFAULTS,
            $.isPlainObject(options) && options
        );
        this.resourseData = {};
        this.init();
    }

    QorBottomSheets.prototype = {
        constructor: QorBottomSheets,

        init: function () {
            this.build();
            this.bind();
        },

        build: function () {
            let $bottomsheets;

            this.$bottomsheets = $bottomsheets = $(QorBottomSheets.TEMPLATE).appendTo(
                "body"
            );
            this.$body = $bottomsheets.find(".qor-bottomsheets__body");
            this.$title = $bottomsheets.find(".qor-bottomsheets__title");
            this.$header = $bottomsheets.find(".qor-bottomsheets__header");
            this.$bodyClass = $("body").prop("class");
            this.filterURL = "";
            this.searchParams = "";
        },

        bind: function () {
            this.$bottomsheets
                .on(EVENT_SUBMIT, "form", this.submit.bind(this))
                .on(EVENT_CLICK, '[data-dismiss="bottomsheets"]', this.hide.bind(this))
                .on(EVENT_CLICK, ".qor-pagination-container a", this.pagination.bind(this))
                .on(EVENT_CLICK, CLASS_BOTTOMSHEETS_BUTTON, this.search.bind(this))
                .on(EVENT_KEYUP, this.keyup.bind(this))
                .on("selectorChanged.qor.selector", this.selectorChanged.bind(this))
                .on("filterChanged.qor.filter", this.filterChanged.bind(this))
                .on(EVENT_RELOADFROMURL, this.reloadFromUrl.bind(this));
        },

        unbind: function () {
            this.$bottomsheets
                .off(EVENT_SUBMIT, "form")
                .off(EVENT_CLICK)
                .off("selectorChanged.qor.selector")
                .off("filterChanged.qor.filter");
        },

        bindActionData: function (actiondData) {
            var $form = this.$body
                .find('[data-toggle="qor-action-slideout"]')
                .find("form");
            for (var i = actiondData.length - 1; i >= 0; i--) {
                $form.prepend(
                    '<input type="hidden" name="primary_values[]" value="' +
                    actiondData[i] +
                    '" />'
                );
            }
        },

        filterChanged: function (e, search, key) {
            // if this event triggered:
            // search: ?locale_mode=locale, ?filters[Color].Value=2
            // key: search param name: locale_mode

            var loadUrl;

            loadUrl = this.constructloadURL(search, key);
            loadUrl && this.reload(loadUrl);
            return false;
        },

        selectorChanged: function (e, url, key) {
            // if this event triggered:
            // url: /admin/!remote_data_searcher/products/Collections?locale=en-US
            // key: search param key: locale

            var loadUrl;

            loadUrl = this.constructloadURL(url, key);
            loadUrl && this.reload(loadUrl);
            return false;
        },

        keyup: function (e) {
            var searchInput = this.$bottomsheets.find(CLASS_BOTTOMSHEETS_INPUT);

            if (e.which === 13 && searchInput.length && searchInput.is(":focus")) {
                this.search();
            }
        },

        search: function () {
            var $bottomsheets = this.$bottomsheets,
                param = "?keyword=",
                baseUrl = $bottomsheets.data().url,
                searchValue = $.trim(
                    $bottomsheets.find(CLASS_BOTTOMSHEETS_INPUT).val()
                ),
                url = baseUrl + param + searchValue;

            if (/\?/g.test(baseUrl)) {
                url = baseUrl + "&keyword=" + searchValue;
            }

            this.reload(url);
        },

        pagination: function (e) {
            var $ele = $(e.target).closest("a"),
                url = $ele.prop("href");
            if (url) {
                this.reload(url);
            }
            return false;
        },

        reload: function (url) {
            var $content = this.$bottomsheets.find(CLASS_BODY_CONTENT);

            this.addLoading($content);
            this.fetchPage(url);
        },

        reloadFromUrl: function (e, url) {
            this.reload(url);
        },

        fetchPage: function (url) {
            var $bottomsheets = this.$bottomsheets,
                _this = this;

            $.get(url, function (response) {
                var $response = $(response).find(CLASS_MAIN_CONTENT),
                    $responseHeader = $response.find(CLASS_BODY_HEAD),
                    $responseBody = $response.find(CLASS_BODY_CONTENT);

                if ($responseBody.length) {
                    $bottomsheets.find(CLASS_BODY_CONTENT).html($responseBody.html());

                    if ($responseHeader.length) {
                        _this.$body
                            .find(CLASS_BODY_HEAD)
                            .html($responseHeader.html())
                            .trigger("enable");
                        _this.addHeaderClass();
                    }
                    // will trigger this event(relaod.qor.bottomsheets) when bottomsheets reload complete: like pagination, filter, action etc.
                    $bottomsheets.trigger(EVENT_RELOAD);
                } else {
                    _this.reload(url);
                }
            }).fail(function () {
                window.alert("server error, please try again later!");
            });
        },

        constructloadURL: function (url, key) {
            var fakeURL,
                value,
                filterURL = this.filterURL,
                bindUrl = this.$bottomsheets.data().url;

            if (!filterURL) {
                if (bindUrl) {
                    filterURL = bindUrl;
                } else {
                    return;
                }
            }

            fakeURL = new URL(URL_GETQOR + url);
            value = getUrlParameter(key, fakeURL.search);
            filterURL = this.filterURL = updateQueryStringParameter(
                key,
                value,
                filterURL
            );

            return filterURL;
        },

        addHeaderClass: function () {
            this.$body.find(CLASS_BODY_HEAD).hide();
            if (
                this.$bottomsheets
                    .find(CLASS_BODY_HEAD)
                    .children(CLASS_BOTTOMSHEETS_FILTER).length
            ) {
                this.$body
                    .addClass("has-header")
                    .find(CLASS_BODY_HEAD)
                    .show();
            }
        },

        addLoading: function ($element) {
            $element.html("");
            $(QorBottomSheets.TEMPLATE_LOADING)
                .appendTo($element)
                .trigger("enable.qor.material");
        },

        loadExtraResource: function (data) {
            let styleDiff = compareLinks(data.$links),
                scriptDiff = compareScripts(data.$scripts);

            styleDiff.length && loadStyles(styleDiff);
            scriptDiff.length && loadScripts(scriptDiff, data);
        },

        loadMedialibraryJS: function ($response) {
            var $script = $response.filter("script"),
                theme = /theme=media_library/g,
                src,
                _this = this;

            $script.each(function () {
                src = $(this).prop("src");
                if (theme.test(src)) {
                    var script = document.createElement("script");
                    script.src = src;
                    document.body.appendChild(script);
                    _this.mediaScriptAdded = true;
                }
            });
        },

        submit: function (e) {
            let form = e.target,
                $form = $(form),
                _this = this,
                url = $form.prop("action"),
                formData,
                $bottomsheets = $form.closest(".qor-bottomsheets"),
                resourseData = $bottomsheets.data(),
                ajaxType = resourseData.ajaxType,
                $submit = $form.find(":submit");

            // will ingore submit event if need handle with other submit event: like select one, many...
            if (resourseData.ingoreSubmit) {
                return;
            }

            // will submit form as normal,
            // if you need download file after submit form or other things, please add
            // data-use-normal-submit="true" to form tag
            // <form action="/admin/products/!action/localize" method="POST" enctype="multipart/form-data" data-normal-submit="true"></form>
            var normalSubmit = $form.data().normalSubmit;

            if (normalSubmit) {
                return;
            }

            $(document).trigger(EVENT_BOTTOMSHEET_BEFORESEND);
            e.preventDefault();

            formData = new FormData(form);

            $.ajax(url, {
                method: $form.prop("method"),
                data: formData,
                dataType: ajaxType ? ajaxType : "html",
                processData: false,
                contentType: false,
                beforeSend: function () {
                    $submit.prop("disabled", true);
                },
                success: function (data, textStatus, jqXHR) {
                    if (resourseData.ajaxMute) {
                        $bottomsheets.remove();
                        return;
                    }

                    if (resourseData.ajaxTakeover) {
                        resourseData.$target
                            .parent()
                            .trigger(EVENT_SUBMITED, [data, $bottomsheets]);
                        return;
                    }

                    // handle file download from form submit
                    var disposition = jqXHR.getResponseHeader("Content-Disposition");
                    if (disposition && disposition.indexOf("attachment") !== -1) {
                        var fileNameRegex = /filename[^;=\n]*=((['"]).*?\2|[^;\n]*)/,
                            matches = fileNameRegex.exec(disposition),
                            contentType = jqXHR.getResponseHeader("Content-Type"),
                            fileName = "";

                        if (matches != null && matches[1]) {
                            fileName = matches[1].replace(/['"]/g, "");
                        }

                        window.QOR.qorAjaxHandleFile(url, contentType, fileName, formData);
                        $submit.prop("disabled", false);

                        return;
                    }

                    $(".qor-error").remove();

                    var returnUrl = $form.data("returnUrl");
                    var refreshUrl = $form.data("refreshUrl");

                    if (refreshUrl) {
                        window.location.href = refreshUrl;
                        return;
                    }

                    if (returnUrl == "refresh") {
                        _this.refresh();
                        return;
                    }

                    if (returnUrl && returnUrl != "refresh") {
                        _this.load(returnUrl);
                    } else {
                        _this.refresh();
                    }

                    $(document).trigger(EVENT_BOTTOMSHEET_SUBMIT);
                },
                error: function (err) {
                    window.QOR.handleAjaxError(err);
                },
                complete: function () {
                    $submit.prop("disabled", false);
                }
            });
        },

        load: function (url, data, callback) {
            var options = this.options,
                method,
                dataType,
                load,
                actionData = data.actionData,
                resourseData = this.resourseData,
                selectModal = resourseData.selectModal,
                ingoreSubmit = resourseData.ingoreSubmit,
                $bottomsheets = this.$bottomsheets,
                $header = this.$header,
                $body = this.$body;

            if (!url) {
                return;
            }

            this.show();
            this.addLoading($body);

            this.filterURL = url;
            $body.removeClass("has-header has-hint");

            data = $.isPlainObject(data) ? data : {};

            method = data.method ? data.method : "GET";
            dataType = data.datatype ? data.datatype : "html";

            load = $.proxy(function () {
                $.ajax(url, {
                    method: method,
                    dataType: dataType,
                    success: $.proxy(function (response) {
                        if (method === "GET") {
                            let $response = $(response),
                                $content,
                                bodyClass,
                                loadExtraResourceData = {
                                    $scripts: $response.filter("script"),
                                    $links: $response.filter("link"),
                                    url: url,
                                    response: response
                                },
                                hasSearch =
                                    selectModal && $response.find(".qor-search-container").length,
                                bodyHtml = response.match(/<\s*body.*>[\s\S]*<\s*\/body\s*>/gi);

                            $content = $response.find(CLASS_MAIN_CONTENT);

                            if (bodyHtml) {
                                bodyHtml = bodyHtml
                                    .join("")
                                    .replace(/<\s*body/gi, "<div")
                                    .replace(/<\s*\/body/gi, "</div");
                                bodyClass = $(bodyHtml).prop("class");
                                $("body").addClass(bodyClass);
                            }

                            if (!$content.length) {
                                return;
                            }

                            this.loadExtraResource(loadExtraResourceData);

                            if (ingoreSubmit) {
                                $content.find(CLASS_BODY_HEAD).remove();
                            }

                            $content
                                .find(".qor-button--cancel")
                                .attr("data-dismiss", "bottomsheets");

                            $body.html($content.html());
                            this.$title.html($response.find(options.title).html());

                            if (data.selectDefaultCreating) {
                                this.$title.append(
                                    `<button class="mdl-button mdl-button--primary" type="button" data-load-inline="true" data-select-nohint="${
                                        data.selectNohint
                                    }" data-select-modal="${
                                        data.selectModal
                                    }" data-select-listing-url="${data.selectListingUrl}">${
                                        data.selectBacktolistTitle
                                    }</button>`
                                );
                            }

                            if (selectModal) {
                                $body
                                    .find(".qor-button--new")
                                    .data("ingoreSubmit", true)
                                    .data("selectId", resourseData.selectId)
                                    .data("loadInline", true);
                                if (
                                    selectModal != "one" &&
                                    !data.selectNohint &&
                                    (typeof resourseData.maxItem === "undefined" ||
                                        resourseData.maxItem != "1")
                                ) {
                                    $body.addClass("has-hint");
                                }
                                if (selectModal == "mediabox" && !this.mediaScriptAdded) {
                                    this.loadMedialibraryJS($response);
                                }
                            }

                            $header.find(".qor-button--new").remove();
                            this.$title.after($body.find(".qor-button--new"));

                            if (hasSearch) {
                                $bottomsheets.addClass("has-search");
                                $header.find(".qor-bottomsheets__search").remove();
                                $header.prepend(QorBottomSheets.TEMPLATE_SEARCH);
                            }

                            if (actionData && actionData.length) {
                                this.bindActionData(actionData);
                            }

                            if (resourseData.bottomsheetClassname) {
                                $bottomsheets.addClass(resourseData.bottomsheetClassname);
                            }

                            $bottomsheets.trigger("enable");

                            $bottomsheets.one(EVENT_HIDDEN, function () {
                                $(this).trigger("disable");
                            });

                            this.addHeaderClass();
                            $bottomsheets.data(data);

                            // handle after opened callback
                            if (callback && $.isFunction(callback)) {
                                callback(this.$bottomsheets);
                            }

                            // callback for after bottomSheets loaded HTML
                            $bottomsheets.trigger(EVENT_BOTTOMSHEET_LOADED, [url, response]);
                        } else {
                            if (data.returnUrl) {
                                this.load(data.returnUrl);
                            } else {
                                this.refresh();
                            }
                        }
                    }, this),

                    error: $.proxy(function () {
                        this.$bottomsheets.remove();
                        if (!$(".qor-bottomsheets").is(":visible")) {
                            $("body").removeClass(CLASS_OPEN);
                        }
                        var errors;
                        if ($(".qor-error span").length > 0) {
                            errors = $(".qor-error span")
                                .map(function () {
                                    return $(this).text();
                                })
                                .get()
                                .join(", ");
                        } else {
                            errors = QOR_Translations.serverError;
                        }
                        window.alert(errors);
                    }, this)
                });
            }, this);

            load();
        },

        open: function (options, callback) {
            if (!options.loadInline) {
                this.init();
            }
            this.resourseData = options;
            this.load(options.url, options, callback);
        },

        show: function () {
            this.$bottomsheets.addClass(CLASS_IS_SHOWN).get(0).offsetHeight;
            this.$bottomsheets.addClass(CLASS_IS_SLIDED);
            $("body").addClass(CLASS_OPEN);
        },

        hide: function (e) {
            let $bottomsheets = $(e.target).closest(".qor-bottomsheets"),
                $datePicker = $(".qor-datepicker").not(".hidden");

            if ($datePicker.length) {
                $datePicker.addClass("hidden");
            }

            $bottomsheets.qorSelectCore("destroy");

            $bottomsheets.trigger(EVENT_BOTTOMSHEET_CLOSED).remove();
            if (!$(".qor-bottomsheets").is(":visible")) {
                $("body").removeClass(CLASS_OPEN);
            }
            return false;
        },

        refresh: function () {
            this.$bottomsheets.remove();
            $("body").removeClass(CLASS_OPEN);

            setTimeout(function () {
                window.location.reload();
            }, 350);
        },

        destroy: function () {
            this.unbind();
            this.$element.removeData(NAMESPACE);
        }
    };

    QorBottomSheets.DEFAULTS = {
        title: ".qor-form-title, .mdl-layout-title",
        content: false
    };

    QorBottomSheets.TEMPLATE_ERROR = `<ul class="qor-error"><li><label><i class="material-icons">error</i><span>[[error]]</span></label></li></ul>`;
    QorBottomSheets.TEMPLATE_LOADING = `<div style="text-align: center; margin-top: 30px;"><div class="mdl-spinner mdl-js-spinner is-active qor-layout__bottomsheet-spinner"></div></div>`;
    QorBottomSheets.TEMPLATE_SEARCH = `<div class="qor-bottomsheets__search">
            <input autocomplete="off" type="text" class="mdl-textfield__input qor-bottomsheets__search-input" placeholder="Search" />
            <button class="mdl-button mdl-js-button mdl-button--icon qor-bottomsheets__search-button" type="button"><i class="material-icons">search</i></button>
        </div>`;

    QorBottomSheets.TEMPLATE = `<div class="qor-bottomsheets">
            <div class="qor-bottomsheets__header">
            <h3 class="qor-bottomsheets__title"></h3>
            <button type="button" class="mdl-button mdl-button--icon mdl-js-button mdl-js-repple-effect qor-bottomsheets__close" data-dismiss="bottomsheets">
            <span class="material-icons">close</span>
            </button>
            </div>
            <div class="qor-bottomsheets__body"></div>
        </div>`;

    QorBottomSheets.plugin = function (options) {
        return this.each(function () {
            var $this = $(this);
            var data = $this.data(NAMESPACE);
            var fn;

            if (!data) {
                if (/destroy/.test(options)) {
                    return;
                }

                $this.data(NAMESPACE, (data = new QorBottomSheets(this, options)));
            }

            if (typeof options === "string" && $.isFunction((fn = data[options]))) {
                fn.apply(data);
            }
        });
    };

    $.fn.qorBottomSheets = QorBottomSheets.plugin;

    return QorBottomSheets;
});
