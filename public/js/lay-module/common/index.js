layui.define(['jquery'], function (exports) {
    var $ = layui.jquery;
    var obj = {
        get_role_ids: function (roles) {
            var role_ids = [];
            if (roles && roles.length > 0) {
                for (x in roles) {
                    role_ids.push(roles[x].value);
                }
            }

            return role_ids;
        },
        get: function (url, isshow) {
            return new Promise(async (resolve, reject) => {
                let ajax_abort = $.ajax({
                    url: url,
                    type: 'get',
                    dataType: 'JSON',
                    success: function (res) {
                        if (isshow) {
                            if (res.status) {
                                layer.msg(res.msg, {
                                    offset: '15px',
                                    icon: 1,
                                    time: 2000,
                                    id: 'Message',
                                })
                            } else {
                                layer.msg(res.msg, {
                                    offset: '15px',
                                    icon: 2,
                                    time: 2000,
                                    id: 'Message',
                                })
                            }
                        }
                        resolve(res)
                    },
                    error: function (error) {
                        if (error.responseJSON) {
                            for (let i in error.responseJSON.errors) {
                                layer.msg(error.responseJSON.errors[i].join('、'), {
                                    offset: '15px',
                                    icon: 2,
                                    time: 2000,
                                    id: 'Message',
                                })
                            }
                        }
                        layer.closeAll('loading');
                        reject(error.responseJSON)
                    },
                    complete: function (XMLHttpRequest, status) {
                        if (status === 'timeout') {
                            ajax_abort.abort();
                            layer.msg('会话请求超时，请重新登录！', {
                                offset: '15px',
                                icon: 2,
                                time: 2000,
                                id: 'Message',
                            })
                        }
                        layer.closeAll('loading');
                        reject(status)
                    },
                })
            })
        },


    };

    //输出接口
    exports('common', obj);

});