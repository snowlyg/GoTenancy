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

    let NAMESPACE = 'qor.worker',
        EVENT_ENABLE = 'enable.' + NAMESPACE,
        EVENT_DISABLE = 'disable.' + NAMESPACE,
        EVENT_CLICK = 'click.' + NAMESPACE,
        CLASS_NEW_WORKER = '.qor-worker--new',
        CLASS_WORKER_CONTAINER = '.qor-worker-form',
        CLASS_WORKER_LIST = '.qor-worker-form-list',
        CLASS_WORKER_PROGRESS = '.qor-worker--progress',
        CLASS_WORKER_SHOW = '.qor-worker-form--show',
        CLASS_BUTTON_BACK = '.qor-button--back',
        CLASS_TABLE = '.qor-js-table',
        CLASS_SELECT = '.is-selected';

    function updateProgress(progress) {
        let $progress = document.querySelector('#qor-worker--progress');
        if ($progress) {
            document.querySelector('#qor-worker--progress').MaterialProgress.setProgress(progress);
        }
    }

    function QorWorker(element, options) {
        this.$element = $(element);
        this.options = $.extend({}, QorWorker.DEFAULTS, $.isPlainObject(options) && options);
        this.init();
    }

    QorWorker.prototype = {
        constructor: QorWorker,

        init: function () {
            this.bind();
            this.formOpened = false;
            if ($(CLASS_WORKER_SHOW).length) {
                this.formOpened = true;
            }

            updateProgress($('#qor-worker--progress').data('progress'));

            if (!$('.qor-slideout').is(':visible')) {
                $.fn.qorSliderAfterShow.updateWorkerProgress();
            }
        },

        bind: function () {
            this.$element.on(EVENT_CLICK, CLASS_NEW_WORKER, this.showForm.bind(this)).on(EVENT_CLICK, CLASS_BUTTON_BACK, this.hideForm.bind(this));
        },

        unbind: function () {
            this.$element.off(EVENT_CLICK);
        },

        hideForm: function (e) {
            e.preventDefault();

            var $parent = this.$element;
            var $lists = $parent.find(CLASS_WORKER_CONTAINER).find('>li');

            $lists
                .show()
                .removeClass('current')
                .find('form')
                .addClass('hidden');
            $(CLASS_BUTTON_BACK).addClass('hidden');
            $(CLASS_WORKER_LIST).show();

            this.formOpened = false;

            window.onbeforeunload = null;
            $.fn.qorSlideoutBeforeHide = null;
        },

        showForm: function (e) {
            var $target = $(e.target);
            e.preventDefault();

            if (this.formOpened) {
                return;
            }

            var $targetList = $target.closest('li');
            var $parent = $target.closest(CLASS_WORKER_CONTAINER);
            var $parentList = $target.closest(CLASS_WORKER_LIST);
            var $lists = $parent.find('>li');

            $lists.hide().removeClass('current');

            $targetList.addClass('current').show();
            $(CLASS_BUTTON_BACK).removeClass('hidden');
            $targetList.find(CLASS_WORKER_LIST).hide();

            $parentList
                .show()
                .find('form')
                .removeClass('hidden');
            this.formOpened = true;
        },

        destroy: function () {
            this.unbind();
            QorWorker.getWorkerProgressIntervId && window.clearInterval(QorWorker.getWorkerProgressIntervId);
        }
    };

    QorWorker.DEFAULTS = {};
    QorWorker.POPOVERTEMPLATE = `<div class="qor-modal fade qor-modal--worker-errors" tabindex="-1" role="dialog" aria-hidden="true">
          <div class="mdl-card mdl-shadow--2dp" role="document">
            <div class="mdl-card__title">
              <h2 class="mdl-card__title-text">Process Errors</h2>
            </div>
          <div class="mdl-card__supporting-text" id="qor-worker-errors"></div>
            <div class="mdl-card__actions">
              <a class="mdl-button mdl-button--colored mdl-js-button mdl-js-ripple-effect" data-dismiss="modal">close</a>
            </div>
          </div>
        </div>`;

    QorWorker.plugin = function (options) {
        return this.each(function () {
            var $this = $(this);
            var data = $this.data(NAMESPACE);
            var fn;

            if (!data) {
                if (/destroy/.test(options)) {
                    return;
                }

                $this.data(NAMESPACE, (data = new QorWorker(this, options)));
            }

            if (typeof options === 'string' && $.isFunction((fn = data[options]))) {
                fn.apply(data);
            }
        });
    };

    $.fn.qorSliderAfterShow.updateWorkerProgress = function (url) {
        if (!$('.workers-log-output').length) {
            return;
        }
        QorWorker.getWorkerProgressIntervId = window.setInterval(QorWorker.updateWorkerProgress, 2000, url);
    };

    QorWorker.updateTableStatus = function (status) {
        var $selectedItem = $(CLASS_TABLE).find(CLASS_SELECT);
        var statusName = $(CLASS_WORKER_PROGRESS).data().statusName;

        $selectedItem
            .find('td[data-heading="' + statusName + '"]')
            .find('.qor-table__content')
            .html(status);
    };

    QorWorker.isScrollToBottom = function (element) {
        return element.clientHeight + element.scrollTop <= element.scrollHeight;
    };

    QorWorker.updateWorkerProgress = function (url) {
        let progressURL = url,
            $logContainer = $('.workers-log-output'),
            $progressValue = $('.qor-worker--progress-value'),
            $progressStatusStatus = $('.qor-worker--progress-status'),
            $progress = $(CLASS_WORKER_PROGRESS),
            $selectTR = $(CLASS_TABLE).find(CLASS_SELECT),
            status = ['killed', 'exception', 'cancelled', 'scheduled'];

        if (!$logContainer.length) {
            return;
        }
        if ($progress.length) {
            var progressData = $progress.data();
        }

        if ($selectTR.length && progressData && progressData.statusName) {
            var orignialStatus = $selectTR
                .find('td[data-heading="' + progressData.statusName + '"]')
                .find('.qor-table__content')
                .html();
        }

        if (!$progress.length || !$progress.length || status.indexOf(progressData.status) != -1) {
            window.clearInterval(QorWorker.getWorkerProgressIntervId);
            return;
        }

        if (progressData.progress >= 100) {
            window.clearInterval(QorWorker.getWorkerProgressIntervId);
            QorWorker.updateTableStatus(progressData.status);
            $('.qor-workers-abort').addClass('hidden');
            $('.qor-workers-rerun').removeClass('hidden');
            return;
        }

        $.ajax({
            url: progressURL,
            method: 'GET',
            dataType: 'html',
            processData: false,
            contentType: false
        }).done(function (html) {
            let $html = $(html),
                contentData = $html.find(CLASS_WORKER_PROGRESS).data(),
                currentStatus = contentData.progress,
                progressStatusStatus = contentData.status,
                isNotNormalStatus = status.indexOf(progressStatusStatus) != -1;

            $progressValue.html(currentStatus);
            $progressStatusStatus.html(progressStatusStatus);

            if (isNotNormalStatus) {
                $progressStatusStatus.addClass('highlight');
            }

            // set status progress
            updateProgress(currentStatus);

            // update process log
            let log = $.trim($html.find('.workers-log-output').html()),
                $errorTable = $html.find('.workers-error-output');

            if (QorWorker.isScrollToBottom($logContainer[0])) {
                $logContainer.html(log).scrollTop($logContainer[0].scrollHeight);
            } else {
                $logContainer.html(log);
            }

            if ($errorTable.length) {
                $('.workers-error-output').html($errorTable.html());
            }

            if (orignialStatus != progressStatusStatus) {
                QorWorker.updateTableStatus(progressStatusStatus);
            }

            if (currentStatus >= 100 || isNotNormalStatus) {
                window.clearInterval(QorWorker.getWorkerProgressIntervId);
                $('.qor-workers-abort').addClass('hidden');
                $('.qor-workers-rerun').removeClass('hidden');
                $('.qor-worker--progress-result').html($html.find('.qor-worker--progress-result').html());
            }
        });
    };

    $(function () {
        var selector = '[data-toggle="qor.workers"]';

        $(document)
            .on(EVENT_DISABLE, function (e) {
                QorWorker.plugin.call($(selector, e.target), 'destroy');
            })
            .on(EVENT_ENABLE, function (e) {
                QorWorker.plugin.call($(selector, e.target));
            })
            .triggerHandler(EVENT_ENABLE);
    });

    return QorWorker;
});
