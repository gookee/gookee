<!DOCTYPE html>
<html>
<head>
    <title></title>
    <link rel="stylesheet" type="text/css" href="/css/css.css"/>
    <link rel="stylesheet" type="text/css" href="/themes/bootstrap/easyui.css">
    <link rel="stylesheet" type="text/css" href="/themes/icon.css">
    <script type="text/javascript" src="/js/jquery.js"></script>
    <script type="text/javascript" src="/js/jquery.easyui.min.js"></script>
    <script type="text/javascript" src="/js/easyui-lang-zh_CN.js"></script>
    <script type="text/javascript" src="/js/bll.js"></script>
    <script type="text/javascript">
        if(self != top)
            top.location.href = self.location.href;

        function startExam() {
            var username = $.trim($("#username").val());
            var password = $.trim($("#password").val());
            if (username == "") {
                alert('用户名不能为空！');
                $('#username').focus();
                return;
            }
            if (password == "") {
                alert('密码不能为空！');
                $('#password').focus();
                return;
            }

            $('#dataRow').submit();
        }

        $(document).on('keyup', function(e){
            if($(e.target).attr('id') == "password" && e.keyCode == 13){
                startExam();
            }
        });
    </script>
</head>
<body class="easyui-layout">
    <div data-options="region:'north'" style="height:56px; overflow: hidden;">
        <div class="header" style="height:56px; width: 100%;"></div>
    </div>
    <div data-options="region: 'center'" style="text-align: left; padding-top: 80px;">
        <form method="post" id="dataRow">
            <table border="0" cellpadding="3" cellspacing="5" width="400" align="center">
                <tr>
                    <td style="width: 80px;">用户名：</td>
                    <td><input name="username" id="username" type="text" style="width: 250px; height: 22px;" value=""/></td>
                </tr>
                <tr>
                    <td>密&nbsp; 码：</td>
                    <td><input name="password" id="password" type="password" style="width: 250px; height: 22px;" value=""/>
                    </td>
                </tr>
                <tr>
                    <td></td>
                    <td>
                        <a href="javascript:void(0);" onclick="startExam();" class="easyui-linkbutton"
                        data-options="iconCls:'icon-ok'">登录</a>
                    </td>
                </tr>
            </table>
        </form>
    </div>
</body>
</html>