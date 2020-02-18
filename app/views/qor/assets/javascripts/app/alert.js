$(function () {
    'use strict';

    $(document).on('click.qor.alert', '[data-dismiss="alert"]', function () {
        $(this)
            .closest('.qor-alert')
            .removeClass('qor-alert__active');
    });

    setTimeout(function () {
        $('.qor-alert[data-dismissible="true"]').removeClass('qor-alert__active');
    }, 5000);
});
