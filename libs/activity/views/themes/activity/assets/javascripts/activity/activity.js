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

    let Mustache = window.Mustache,
        NAMESPACE = "qor.activity",
        EVENT_ENABLE = "enable." + NAMESPACE,
        EVENT_DISABLE = "disable." + NAMESPACE,
        EVENT_CLICK = "click." + NAMESPACE,
        EVENT_SUBMIT = "submit." + NAMESPACE,
        CLASS_EDIT_NOTE = ".qor-activity__edit-button",
        CLASS_CONTENT = ".qor-activities__content",
        CLASS_TAB_LISTS = ".qor-tab__activity",
        CLASS_TAB_ACTIVITY = ".qor-activities__header button",
        CLASS_EDIT_NOTE_FORM = ".qor-activity__edit-note_form",
        CLASS_NEW_NOTE_FORM = ".qor-activity__new-note_form",
        CLASS_LISTS = ".qor-activity__lists",
        CLASS_ACTIVITY = "is-active",
        ID_LISTS = "#activity-lists";


    function moveActivityHtml() {
        let $target = $(".qor-slideout > .qor-slideout__body"),
            $ele = $('[data-toggle="qor.activity"]'),
            $publish2 = $(".qor-form-container.qor-pulish2__action"),
            $activityForm = $ele.find("#activity-form"),
            afterQorActivityinit = $.fn.afterQorActivityinit;

        if ($target.length) {
            $ele.appendTo($target);
            $target.find(".qor-page__body,.qor-page__header").appendTo($activityForm);
        }


        if ($publish2.length) {
            $publish2.prependTo($activityForm);
            if (afterQorActivityinit && $.isFunction(afterQorActivityinit)) {
                afterQorActivityinit();
            }
        }
    }

    function QorActivity(element, options) {
        this.$element = $(element);
        this.options = $.extend(
            {},
            QorActivity.DEFAULTS,
            $.isPlainObject(options) && options
        );

        this.init();
    }

    QorActivity.prototype = {
        constructor: QorActivity,

        init: function () {
            this.bind();
            this.initTabs();
        },

        bind: function () {
            this.$element
                .on(EVENT_CLICK, CLASS_EDIT_NOTE, this.click.bind(this))
                .on(EVENT_SUBMIT, "#activity-lists form", this.submit.bind(this))
                .on(EVENT_CLICK, CLASS_TAB_ACTIVITY, this.tabClick.bind(this));
        },

        unbind: function () {
            this.$element.off(EVENT_CLICK).off(EVENT_SUBMIT);
        },

        submit: function (e) {
            let $form = $(e.target),
                FormDatas,
                _this = this,
                NoteTitle = $(ID_LISTS).data().noteTitle;

            e.preventDefault();

            FormDatas = $form.serialize();
            $.ajax($form.prop("action"), {
                method: $form.prop("method"),
                data: FormDatas,
                dataType: "json",
                headers: {
                    Accept: "application/json; charset=utf-8"
                }
            }).done(function (data) {
                if (data.errors) {
                    return;
                }
                data.NoteTitle = NoteTitle;

                if ($form.is(CLASS_EDIT_NOTE_FORM)) {
                    _this.hideEditForm($form);
                    $form.find(".qor-activity__list-note").html(data.Note);
                }

                if ($form.is(CLASS_NEW_NOTE_FORM)) {
                    $(CLASS_LISTS).prepend(_this.renderActivityList(data));
                    _this.clearForm();
                }
            });

            return false;
        },

        renderActivityList: function (data) {
            return Mustache.render(QorActivity.ACTIVITY_LIST_TEMPLATE, data);
        },

        clearForm: function () {
            var $textarea = $(CLASS_NEW_NOTE_FORM).find('textarea[data-toggle="qor.redactor"]');
            if ($textarea.length) {
                $textarea.redactor("code.set", "");
            }

            $(CLASS_NEW_NOTE_FORM)
                .find('[name="QorResource.Content"],[name="QorResource.Note"]')
                .val("");

            $(CLASS_NEW_NOTE_FORM).find(":submit").attr("disabled", true);
        },

        click: function (e) {
            let $target = $(e.target),
                parents = $target.closest(".qor-activity__list");
            this.showEditForm(parents);
            e.stopPropagation();
        },

        tabClick: function (e) {
            let _this = this,
                $element = this.$element,
                $target = $(e.target),
                id = $target.data("id"),
                url = $target.data("resource-url");

            if ($target.is(CLASS_TAB_LISTS) && $target.hasClass(CLASS_ACTIVITY)) {
                return;
            }

            $element.find(CLASS_CONTENT).hide();
            $element.find(id).show();
            $element.find(CLASS_TAB_ACTIVITY).removeClass(CLASS_ACTIVITY);
            $target.addClass(CLASS_ACTIVITY);

            if (url) {
                $.ajax({
                    url: url,
                    method: "GET",
                    dataType: "json",
                    headers: {
                        Accept: "application/json; charset=utf-8"
                    },
                    success: function (data) {
                        if (data.length) {
                            $(CLASS_LISTS).html("");
                            for (var i = data.length - 1; i >= 0; i--) {
                                data[i].NoteTitle = $(ID_LISTS).data().noteTitle;
                                $(CLASS_LISTS).append(_this.renderActivityList(data[i]));
                            }
                        }
                        $(CLASS_LISTS)
                            .find(".mdl-spinner")
                            .remove();
                    }
                });
            } else {
                $(CLASS_LISTS)
                    .find(".mdl-spinner")
                    .remove();
            }
        },

        showEditForm: function (ele) {
            ele
                .find(".qor-activity__list-note,.qor-activity__edit-button")
                .removeClass("show")
                .addClass("hide");
            ele
                .find(".qor-activity__edit-feilds,.qor-activity__edit-save-button")
                .removeClass("hide")
                .addClass("show");
        },

        hideEditForm: function (ele) {
            ele
                .find(".qor-activity__list-note,.qor-activity__edit-button")
                .removeClass("hide")
                .addClass("show");
            ele
                .find(".qor-activity__edit-feilds,.qor-activity__edit-save-button")
                .removeClass("show")
                .addClass("hide");
        },

        initTabs: function () {
            let $parent = this.$element.closest(".mdl-layout__content.qor-page"),
                afterQorActivityinit = $.fn.afterQorActivityinit;

            if ($parent.length) {
                let $ele = $('[data-toggle="qor.activity"]');

                $ele.appendTo($parent);
                $parent
                    .find(".qor-page__body,.qor-page__header")
                    .appendTo($ele.find("#activity-form"));
                $(".qor-form-container.qor-pulish2__action").prependTo(
                    $ele.find("#activity-form")
                );
                if (afterQorActivityinit && $.isFunction(afterQorActivityinit)) {
                    afterQorActivityinit();
                }
            }
        },

        destroy: function () {
            this.unbind();
            this.$element.removeData(NAMESPACE);
        }
    };

    QorActivity.DEFAULTS = {};

    QorActivity.ACTIVITY_LIST_TEMPLATE = `<div class="qor-activity__list">
            <form class="qor-activity__edit-note_form" action="[[ URL ]]" method="POST">
                <input type="hidden" name="QorResource.ID" value="[[ ID ]]">
                <div class="qor-activity__list-title">
                    <strong>[[ CreatorName ]]</strong> <span>[[ Action ]]</span>
                </div>
                <div class="qor-activity__list-date">
                    [[ UpdatedAt ]]
                </div>
                <div class="qor-activity__list-content">[[ &Content ]]</div>
                <div>
                    <strong class="qor-activity__list-note-title">[[ NoteTitle ]]</strong>
                    <span class="qor-activity__list-note">[[ Note ]]</span>
                    <a class="mdl-button mdl-js-button mdl-button--icon qor-activity__edit-button" href="#">
                        <i class="material-icons md-18 qor-activity__edit-button">edit</i>
                    </a>
                    <div class="mdl-textfield mdl-js-textfield qor-activity__edit-feilds">
                        <label class="mdl-textfield__label">[[ Note ]]</label>
                        <input class="mdl-textfield__input" type="text" name="QorResource.Note" value="[[ Note ]]">
                    </div>
                    <button class="mdl-button mdl-js-button mdl-button--icon qor-activity__edit-save-button" type="submit"><i class="material-icons md-24">done</i></button>
                </div>
            </form>
        </div>`;

    QorActivity.plugin = function (options) {
        return this.each(function () {
            var $this = $(this);
            var data = $this.data(NAMESPACE);
            var fn;

            if (!data) {
                if (/destroy/.test(options)) {
                    return;
                }

                $this.data(NAMESPACE, (data = new QorActivity(this, options)));
            }

            if (typeof options === "string" && $.isFunction((fn = data[options]))) {
                fn.apply(data);
            }
        });
    };

    // init activity html after sliderout loaded.
    $.fn.qorSliderAfterShow.qorActivityinit = moveActivityHtml;

    $(function () {
        var selector = '[data-toggle="qor.activity"]';

        $(document)
            .on(EVENT_DISABLE, function (e) {
                QorActivity.plugin.call($(selector, e.target), "destroy");
            })
            .on(EVENT_ENABLE, function (e) {
                QorActivity.plugin.call($(selector, e.target));
            })
            .triggerHandler(EVENT_ENABLE);
    });

    return QorActivity;
});
