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

    var NAMESPACE = 'qor.timepicker';
    var EVENT_ENABLE = 'enable.' + NAMESPACE;
    var EVENT_DISABLE = 'disable.' + NAMESPACE;
    var EVENT_CLICK = 'click.' + NAMESPACE;
    var EVENT_FOCUS = 'focus.' + NAMESPACE;
    var EVENT_KEYDOWN = 'keydown.' + NAMESPACE;
    var EVENT_BLUR = 'blur.' + NAMESPACE;
    var EVENT_CHANGE_TIME = 'selectTime.' + NAMESPACE;

    var CLASS_PARENT = '[data-picker-type]';
    var CLASS_TIME_SELECTED = '.ui-timepicker-selected';

    function QorTimepicker(element, options) {
        this.$element = $(element);
        this.options = $.extend(true, {}, QorTimepicker.DEFAULTS, $.isPlainObject(options) && options);
        this.formatDate = null;
        this.pickerData = this.$element.data();
        this.parent = this.$element.closest(CLASS_PARENT);
        this.isDateTimePicker = this.parent.data('picker-type') == 'datetime';
        this.$targetInput = this.parent.find(this.pickerData.targetInput || (this.isDateTimePicker ? '.qor-datetimepicker__input' : '.qor-datepicker__input'));
        this.init();
    }

    QorTimepicker.prototype = {
        init: function () {
            if (this.$targetInput.is(':disabled')) {
                this.$element.remove();
                return;
            }
            this.bind();
            this.oldValue = this.$targetInput.val();

            var dateNow = new Date();
            var month = dateNow.getMonth() + 1;
            var date = dateNow.getDate();

            month = month < 8 ? '0' + month : month;
            date = date < 10 ? '0' + date : date;

            this.dateValueNow = dateNow.getFullYear() + '-' + month + '-' + date;
        },

        bind: function () {
            var pickerOptions = {
                timeFormat: 'H:i',
                showOn: null,
                wrapHours: false,
                scrollDefault: 'now'
            };

            if (this.isDateTimePicker) {
                this.$targetInput
                    .qorTimepicker(pickerOptions)
                    .on(EVENT_CHANGE_TIME, $.proxy(this.changeTime, this))
                    .on(EVENT_BLUR, $.proxy(this.blur, this))
                    .on(EVENT_FOCUS, $.proxy(this.focus, this))
                    .on(EVENT_KEYDOWN, $.proxy(this.keydown, this));
            }

            this.$element.on(EVENT_CLICK, $.proxy(this.show, this));
        },

        unbind: function () {
            this.$element.off(EVENT_CLICK, this.show);

            if (this.isDateTimePicker) {
                this.$targetInput
                    .off(EVENT_CHANGE_TIME, this.changeTime)
                    .off(EVENT_BLUR, this.blur)
                    .off(EVENT_FOCUS, this.focus)
                    .off(EVENT_KEYDOWN, this.keydown);
            }
        },

        focus: function () {
        },

        blur: function () {
            var inputValue = this.$targetInput.val();
            var inputArr = inputValue.split(' ');
            var inputArrLen = inputArr.length;

            var tempValue;
            var newDateValue;
            var newTimeValue;
            var isDate;
            var isTime;
            var splitSym;

            var timeReg = /\d{1,2}:\d{1,2}/;
            var dateReg = /^\d{4}-\d{1,2}-\d{1,2}/;

            if (!inputValue) {
                return;
            }

            if (inputArrLen == 1) {
                if (dateReg.test(inputArr[0])) {
                    newDateValue = inputArr[0];
                    newTimeValue = '00:00';
                }

                if (timeReg.test(inputArr[0])) {
                    newDateValue = this.dateValueNow;
                    newTimeValue = inputArr[0];
                }
            } else {
                for (var i = 0; i < inputArrLen; i++) {
                    // check for date && time
                    isDate = dateReg.test(inputArr[i]);
                    isTime = timeReg.test(inputArr[i]);

                    if (isDate) {
                        newDateValue = inputArr[i];
                        splitSym = '-';
                    }

                    if (isTime) {
                        newTimeValue = inputArr[i];
                        splitSym = ':';
                    }

                    tempValue = inputArr[i].split(splitSym);

                    for (var j = 0; j < tempValue.length; j++) {
                        if (tempValue[j].length < 2) {
                            tempValue[j] = '0' + tempValue[j];
                        }
                    }

                    if (isDate) {
                        newDateValue = tempValue.join(splitSym);
                    }

                    if (isTime) {
                        newTimeValue = tempValue.join(splitSym);
                    }
                }
            }

            if (this.checkDate(newDateValue) && this.checkTime(newTimeValue)) {
                this.$targetInput.val(newDateValue + ' ' + newTimeValue);
                this.oldValue = this.$targetInput.val();
            } else {
                this.$targetInput.val(this.oldValue);
            }
        },

        keydown: function (e) {
            var keycode = e.keyCode;
            var keys = [48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 8, 37, 38, 39, 40, 27, 32, 20, 189, 16, 186, 96, 97, 98, 99, 100, 101, 102, 103, 104, 105];
            if (keys.indexOf(keycode) == -1) {
                e.preventDefault();
            }
        },

        checkDate: function (value) {
            var regCheckDate = /^(?:(?!0000)[0-9]{4}-(?:(?:0[1-9]|1[0-2])-(?:0[1-9]|1[0-9]|2[0-8])|(?:0[13-9]|1[0-2])-(?:29|30)|(?:0[13578]|1[02])-31)|(?:[0-9]{1,2}(?:0[48]|[2468][048]|[13579][26])|(?:0[48]|[2468][048]|[13579][26])00)-02-29)$/;
            return regCheckDate.test(value);
        },

        checkTime: function (value) {
            var regCheckTime = /^([01]\d|2[0-3]):?([0-5]\d)$/;
            return regCheckTime.test(value);
        },

        changeTime: function () {
            var $targetInput = this.$targetInput;

            var oldValue = this.oldValue;
            var timeReg = /\d{1,2}:\d{1,2}/;
            var hasTime = timeReg.test(oldValue);
            var selectedTime = $targetInput
                .data()
                .timepickerList.find(CLASS_TIME_SELECTED)
                .html();
            var newValue;

            if (!oldValue) {
                newValue = this.dateValueNow + ' ' + selectedTime;
            } else if (hasTime) {
                newValue = oldValue.replace(timeReg, selectedTime);
            } else {
                newValue = oldValue + ' ' + selectedTime;
            }

            $targetInput.val(newValue);
        },

        show: function () {
            if (!this.isDateTimePicker) {
                return;
            }

            this.$targetInput.qorTimepicker('show');
            this.oldValue = this.$targetInput.val();
        },

        destroy: function () {
            this.unbind();
            this.$targetInput.qorTimepicker('remove');
            this.$element.removeData(NAMESPACE);
        }
    };

    QorTimepicker.DEFAULTS = {};

    QorTimepicker.plugin = function (option) {
        return this.each(function () {
            var $this = $(this);
            var data = $this.data(NAMESPACE);
            var options;
            var fn;

            if (!data) {
                if (!$.fn.qorDatepicker) {
                    return;
                }

                if (/destroy/.test(option)) {
                    return;
                }

                options = $.extend(true, {}, $this.data(), typeof option === 'object' && option);
                $this.data(NAMESPACE, (data = new QorTimepicker(this, options)));
            }

            if (typeof option === 'string' && $.isFunction((fn = data[option]))) {
                fn.apply(data);
            }
        });
    };

    $(function () {
        var selector = '[data-toggle="qor.timepicker"]';

        $(document)
            .on(EVENT_DISABLE, function (e) {
                QorTimepicker.plugin.call($(selector, e.target), 'destroy');
            })
            .on(EVENT_ENABLE, function (e) {
                QorTimepicker.plugin.call($(selector, e.target));
            })
            .triggerHandler(EVENT_ENABLE);
    });

    return QorTimepicker;
});
