layui.define(['jquery'], function (exports) {
    var $ = layui.jquery;
    var obj = {
        // 格式化多选下拉id
        get_ids: function (items) {
            var ids = [];
            if (items && items.length > 0) {
                for (x in items) {
                    ids.push(parseInt(items[x].value));
                }
            }

            return ids;
        },

        //authtree 方法start
        // 全选样例
        checkAll: function (dst) {
            layui.use(['jquery', 'layer', 'authtree'], function () {
                var layer = layui.layer;
                var authtree = layui.authtree;

                authtree.checkAll(dst);
            });
        },

        // 全不选样例
        uncheckAll: function (dst) {
            layui.use(['jquery', 'layer', 'authtree'], function () {
                var layer = layui.layer;
                var authtree = layui.authtree;

                authtree.uncheckAll(dst);
            });
        },

        // 显示全部
        showAll: function (dst) {
            layui.use(['jquery', 'layer', 'authtree'], function () {
                var layer = layui.layer;
                var authtree = layui.authtree;

                authtree.showAll(dst);
            });
        },

        // 隐藏全部
        closeAll: function (dst) {
            layui.use(['jquery', 'layer', 'authtree'], function () {
                var layer = layui.layer;
                var authtree = layui.authtree;

                authtree.closeAll(dst);
            });
        },

        //authtree end

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
            }).catch((e) => {
            });
        },

        //post 请求
        //url 请求地址
        //data 请求数据
        //isshow 是否显示返回 msg
        post: function (url, data, isshow, cb) {
            let callback = function (res) {
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
            };

            if (cb) {
                callback = cb;
            }

            return new Promise(async (resolve, reject) => {
                let ajax_abort = $.ajax({
                    url: url,
                    type: 'POST',
                    data: data,
                    dataType: 'JSON',
                    timeout: 8000,
                    success: callback,
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
            }).catch((e) => {
            });
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
            }).catch((e) => {
            });
        },

        //delete 请求
        //url 请求地址
        delete: function (url, cb) {
            var callback = function (res) {
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
            };
            if (cb) {
                callback = cb;
            }
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
                    success: callback,
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
            }).catch((e) => {
            });
        }
    };

    //输出接口
    exports('common', obj);

});