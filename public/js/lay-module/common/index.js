layui.define(['jquery'], function (exports) {
    var $ = layui.jquery;
    var obj = {
        // 格式化多选下拉角色id
        get_role_ids: function (roles) {
            var role_ids = [];
            if (roles && roles.length > 0) {
                for (x in roles) {
                    role_ids.push(roles[x].value);
                }
            }

            return role_ids;
        },

        // get 请求
        // url 请求地址
        //isshow 是否显示返回 msg
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

        //post 请求
        //url 请求地址
        //data 请求数据
        //isshow 是否显示返回 msg
        post: function (url, data, isshow) {
            return new Promise(async (resolve, reject) => {
                let ajax_abort = $.ajax({
                    url: url,
                    type: 'POST',
                    data: data,
                    dataType: 'JSON',
                    timeout: 8000,
                    success: function (res) {
                        if (isshow) {
                            if (res.status) {
                                layer.msg(res.msg, {
                                    offset: '15px',
                                    icon: 1,
                                    time: 1000,
                                    id: 'Message',
                                })
                            } else {
                                layer.msg(res.msg, {
                                    offset: '15px',
                                    icon: 2,
                                    time: 1000,
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

        //patch 请求
        //url 请求地址
        //data 请求数据
        //isshow 是否显示返回 msg
        patch: function (url, data, isshow) {
            return new Promise(async (resolve, reject) => {
                let ajax_abort = $.ajax({
                    url: url,
                    type: 'PATCH',
                    data: data,
                    dataType: 'JSON',
                    timeout: 8000,
                    success: function (res) {
                        if (isshow) {
                            if (res.status) {
                                layer.msg(res.msg, {
                                    offset: '15px',
                                    icon: 1,
                                    time: 1000,
                                    id: 'Message',
                                })
                            } else {
                                layer.msg(res.msg, {
                                    offset: '15px',
                                    icon: 2,
                                    time: 1000,
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

        //delete 请求
        //url 请求地址
        delete: function (url) {
            return new Promise(async (resolve, reject) => {
                let ajax_abort = $.ajax({
                    url: url,
                    type: 'DELETE',
                    headers: {
                        'Content-Type': 'application/json',
                        'X-HTTP-Method-Override': 'DELETE',
                    },
                    dataType: 'JSON',
                    timeout: 8000,
                    success: function (res) {
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