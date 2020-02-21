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

    var dirtyForm = function (ele, options) {
        var hasChangedObj = false;

        if (this instanceof jQuery) {
            options = ele;
            ele = this;
        } else if (!(ele instanceof jQuery)) {
            ele = $(ele);
        }

        ele.each(function (item, element) {
            var $ele = $(element);

            if ($ele.is('form')) {
                if ($ele.hasClass('ignore-dirtyform')) {
                    return false;
                }
                hasChangedObj = dirtyForm(
                    $ele.find(
                        'input:not([type="hidden"]):not(".search-field input"):not(".chosen-search input"):not(".ignore-dirtyform"), textarea, select'
                    ),
                    options
                );
                if (hasChangedObj) {
                    return false;
                }
            } else if ($ele.is(':checkbox') || $ele.is(':radio')) {
                if ($ele.hasClass('ignore-dirtyform')) {
                    return false;
                }

                if (element.checked != element.defaultChecked) {
                    hasChangedObj = true;
                    return false;
                }
            } else if ($ele.is('input') || $ele.is('textarea')) {
                if ($ele.hasClass('ignore-dirtyform')) {
                    return false;
                }

                if (element.value != element.defaultValue) {
                    hasChangedObj = true;
                    return false;
                }
            } else if ($ele.is('select')) {
                if ($ele.hasClass('ignore-dirtyform')) {
                    return false;
                }

                var option;
                var defaultSelectedIndex = 0;
                var numberOfOptions = element.options.length;

                for (var i = 0; i < numberOfOptions; i++) {
                    option = element.options[i];
                    hasChangedObj = hasChangedObj || option.selected != option.defaultSelected;
                    if (option.defaultSelected) {
                        defaultSelectedIndex = i;
                    }
                }

                if (hasChangedObj && !element.multiple) {
                    hasChangedObj = defaultSelectedIndex != element.selectedIndex;
                }

                if (hasChangedObj) {
                    return false;
                }
            }
        });

        return hasChangedObj;
    };

    $.fn.extend({
        dirtyForm: dirtyForm
    });

    $(function () {
        $(document).on('submit', 'form', function () {
            $.fn.qorSlideoutBeforeHide = null;
        });

        $(document).on('change', 'form', function () {
            if ($(this).dirtyForm()) {
                $.fn.qorSlideoutBeforeHide = true;
            } else {
                $.fn.qorSlideoutBeforeHide = null;
            }
        });
    });
});
