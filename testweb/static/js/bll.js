$(function () {
    init();
});

function init() {
    if (typeof(initGrid) != "undefined" && $.isFunction(initGrid)) {
        initGrid();
    }
}

function add() {
    $("#dataRow")[0].reset();
    var arg = arguments;
    $("#form").showDialog({
        title: '添加',
        cache: false,
        onOpen: function () {
            if (typeof(addSet) != "undefined" && $.isFunction(addSet))
                addSet(arg);
        },
        buttons: [
            {
                text: '保存',
                handler: function () {
                    if (typeof(addSubmitSet) != "undefined" && $.isFunction(addSubmitSet) && addSubmitSet(arg) == false) return;
                    $.ajax({
                        type: "post",
                        data: top.$("#dataRow").serialize() + "&action=add",
                        cache: false,
                        success: function (msg) {
                            if (msg == "") {
                                alert("添加成功!");
                                $("#form").hideDialog();
                                init();
                            } else {
                                alert(msg);
                            }
                        },
                        error: function () {
                            alert("请求超时，请重试！");
                        }
                    });

                }
            },
            {
                text: '关闭',
                handler: function () {
                    $("#form").hideDialog();
                }
            }
        ]
    });
}

function edit(id) {
    $("#dataRow")[0].reset();
    var arg = arguments;
    $("#form").showDialog({
        title: '编辑',
        cache: false,
        modal: true,
        onOpen: function () {
            $.ajax({
                type: "post",
                dataType: "json",
                data: {"action": "find", "id": id},
                success: function (data) {
                    if (typeof(editSet) != "undefined" && $.isFunction(editSet))
                        editSet(data, arg);
                },
                error: function () {
                    alert("请求超时，请重试！");
                }
            });
        },
        buttons: [
            {
                text: '保存',
                handler: function () {
                    if (typeof(editSubmitSet) != "undefined" && $.isFunction(editSubmitSet) && editSubmitSet(arg) == false) return;
                    $.ajax({
                        type: "post",
                        data: top.$("#dataRow").serialize() + "&action=edit&id=" + id,
                        success: function (msg) {
                            if (msg == "") {
                                alert("编辑成功！");
                                $("#form").hideDialog();
                                init();
                            } else {
                                alert(msg);
                            }
                        },
                        error: function () {
                            alert("请求超时，请重试！");
                        }
                    });
                }
            },
            {
                text: '关闭',
                handler: function () {
                    $("#form").hideDialog();
                }
            }
        ]
    });
}

function del(index) {
    if (window.confirm("确认删除？")) {
        $.ajax({
            type: "post",
            data: {"action": "del", "id": index},
            cache: false,
            success: function (msg) {
                if (msg == "") {
                    alert("删除成功！");
                    init();
                } else {
                    alert(msg);
                }
            },
            error: function () {
                alert("请求超时，请重试！");
            }
        });
    }
}

function batchDel() {
    if (window.confirm("确认删除？")) {
        var batch = $("[name=id]:checked");
        var str = "";
        for (var i = 0; i < batch.length; i++) {
            str += batch[i].value + ",";
        }
        str = str.substr(0, str.length - 1);
        $.ajax({
            type: "post",
            data: {"action": "batchDel", "ids": str},
            success: function (msg) {
                if (msg == "") {
                    alert("删除成功！");
                    init();
                } else {
                    alert(msg);
                }
            },
            error: function () {
                alert("请求超时，请重试！");
            }
        });
    }
}