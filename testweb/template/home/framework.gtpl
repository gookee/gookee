<!DOCTYPE html>
<html>
<head>
    <title>可多超市加盟商管理系统</title>
    <link rel="stylesheet" type="text/css" href="themes/default/easyui.css">
    <link rel="stylesheet" type="text/css" href="themes/icon.css">
    <link rel="stylesheet" type="text/css" href="css/css.css">
    <script type="text/javascript" src="js/jquery.js"></script>
    <script type="text/javascript" src="js/jquery.easyui.min.js"></script>
    <script type="text/javascript" src="js/easyui-lang-zh_CN.js"></script>
    <script type="text/javascript" src="js/bll.js"></script>
    <script type="text/javascript">
        $(function () {
            $('#menu a[src]').on('click', function () {
                addTab($(this).text(), $(this).attr('src'));
            });

            $('#menu').accordion({
                onSelect: function (title, index) {
                    $('#menuContainer').panel('setTitle', '当前位置：首页>' + title);
                }
            });

            addTab('加盟商查询', 'userinfo', false);
        });

        function addTab(title, src, closable) {
            if ($('#tabs').tabs('getTab', title) == null) {
                $('#tabs').tabs('add', {
                    title: title,
                    content: '<iframe frameborder="0" style="width: 100%; height: ' + ($('#tabs').innerHeight() - 35) + 'px;" src="' + src + '"></iframe>',
                    closable: closable == undefined ? true : closable
                });
            } else {
                selectTab(title);
            }
        }

        function selectTab(title) {
            $('#tabs').tabs('select', title);
        }

        function editUser(){
            $("#frm").showDialog({
                title: '编辑',
                cache: false,
                modal: true,
                onOpen: function () {
                    top.$("#uname").val('{{.}}');
                },
                buttons: [
                    {
                        text: '保存',
                        handler: function () {
                            $.ajax({
                                type: "post",
                                data: top.$("#userData").serialize(),
                                success: function (msg) {
                                    if (msg == "") {
                                        alert("编辑成功！");
                                        $("#frm").hideDialog();
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
                            $("#frm").hideDialog();
                        }
                    }
                ]
            });
        }
    </script>
</head>
<body class="easyui-layout">
<div data-options="region:'north'" style="height:56px; overflow: hidden;">
    <table style="width: 100%; height: 100%;" cellpadding="0" cellspacing="0" class="header">
        <tr>
            <td style="text-align: right; padding-right: 30px; padding-top: 10px;">
                 <a href="javascript:void(0)" onclick="editUser();">密码修改</a>
                &nbsp; &nbsp;
                <a href="javascript:void(0)" onclick="top.location.href='logout';">退出</a>
            </td>
        </tr>
    </table>
</div>
<div id="menuContainer" data-options="region:'west',title:'当前位置：首页&gt;加盟商管理'" style="width:200px;">
    <div id="menu" class="easyui-accordion" data-options="fit:true,border:false">
        <div title="加盟商管理" style="overflow:auto;padding:10px 20px;">
            <div class="li"><a href="javascript:void(0);" src="userinfo">加盟商查询</a></div>
        </div>
    </div>
</div>
<div data-options="region:'center',border:false">
    <div id="tabs" class="easyui-tabs" data-options="fit:true"></div>
</div>
<div style="display: none; width: 370px;" id="frm">
    <form method="post" id="userData">
        <table border="0" cellpadding="3" cellspacing="5" width="100%">
            <tr>
                <td>用户名：</td>
                <td><input name="uname" id="uname" type="text"  style="width: 250px; height: 22px;" value=""/></td>
            </tr>
            <tr>
                <td>密码：</td>
                <td><input name="password" id="password" type="text"  style="width: 250px; height: 22px;" value=""/></td>
            </tr>
        </table>
    </form>
</div>
</body>
</html>