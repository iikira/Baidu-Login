function baidu_login() {
    var formParam = $("#login-form-wrapper").serialize();
    $("#login-submit").attr("disabled", true);

    $("#login-verify").css("display", "none");
    $("#verifycode-wrapper").css("display", "none");
    $("#succeed-div").css("display", "none");

    $("#sended-status").html("");
    $("#code-error").html("");

    $.ajax({
        type: 'post',
        url: "/cgi-bin/baidu/login",
        data: formParam,
        dataType: 'json',
        error: function(XMLHttpRequest, textStatus, errorThrown) {
            alert("发送错误: " + XMLHttpRequest.status);
            $("#login-submit").attr("disabled", false);
        },
        success: function(data) {
            switch (data.errInfo.no) {
                case "0":
                    $("#login-verifyWrapper").css("display", "none"); // 隐藏验证码输入框
                    $("#succeed-div").css("display", ""); // 显示登录成功窗口
                    printCookies(data);
                    $("#verifycode-wrapper").css("display", "none");
                    break;
                case "400101": // 开启登录保护
                case "400023": // 需要验证手机或邮箱
                    $("#login-verify").css("display", "");
                    $("#my-phone").html(data.data.phone);
                    $("#my-email").html(data.data.email);
                    $("#token").val(data.data.token);
                    $("#u").val(data.data.u);
                    $("#login-verifyWrapper").css("display", "none");
                    break;
                case "500001": // 请输入验证码
                case "500002": // 验证码错误
                    $("#login-verifycode").val("");
                    $("#login-verifyWrapper").css("display", "");
                    $("#login-vcodestr").val(data.data.codeString);
                    $("#login-verifyCodeImg").attr("src", "https://wappass.baidu.com/cgi-bin/genimage?" + data.data.codeString);
                    break;
                default:
                    alert("ErrorCode:" + data.errInfo.no + ",ErrorMsg:" + data.errInfo.msg);
                    break;
            }
            reborn('login-submit');
        },
    });
}

function sendType(type) {
    token = $("#token").val();
    $("#verify-type").val(type); // 验证类型, phone or email

    $("#phone-submit").attr("disabled", true);
    $("#email-submit").attr("disabled", true);

    $.ajax({
        type: 'post',
        url: "/cgi-bin/baidu/sendcode",
        data: {
            type: type,
            token: token,
        },
        dataType: 'json',
        error: function(XMLHttpRequest, textStatus, errorThrown) {
            alert("发送数据时出错: " + XMLHttpRequest.status);
            $("#phone-submit").attr("disabled", false);
            $("#email-submit").attr("disabled", false);
        },
        success: function(data) {
            if (data.error == undefined) {
                $("#verifycode-wrapper").css("display", "");
                $("#sended-status").html(data.msg);
            } else {
                alert(data.error)
            }
            $("#phone-submit").attr("disabled", false);
            $("#email-submit").attr("disabled", false);
        },
    });
}

function verify_login() {
    type = $("#verify-type").val();
    var formParam = $("#verifycode-wrapper").serialize();
    $("#code-submit").attr("disabled", true);

    $.ajax({
        type: 'post',
        url: "/cgi-bin/baidu/verifylogin",
        data: formParam,
        dataType: 'json',
        error: function(XMLHttpRequest, textStatus, errorThrown) {
            alert("发送错误: " + XMLHttpRequest.status);
        },
        success: function(data) {
            switch (data.errInfo.no) {
                case "0":
                    $("#login-verify").css("display", "none");
                    $("#verifycode-wrapper").css("display", "none");
                    if (data.data.bduss == "") {
                        document.getElementById('code-msg').innerHTML = "验证成功，请在本页面重新登录";
                    } else {
                        $("#succeed-div").css("display", "");
                        printCookies(data);
                    }
                    break;
                default:
                    $("#code-error").html("ErrorCode: " + data.errInfo.no + ", ErrorMsg: " + data.errInfo.msg);
                    break;
            }
        },
    });
    $("#code-submit").attr("disabled", false);
}

function printCookies(data) {
    document.getElementById('verifycode-wrapper').style.display = "none";
    $("#USERNAME").val($("#login-username").val());
    $("#BDUSS").val(data.data.bduss);
    $("#PTOKEN").val(data.data.ptoken);
    $("#STOKEN").val(data.data.stoken);
    $("#COOKIE").val(data.data.cookieString);
}

function reborn(id) {
    document.getElementById(id).disabled = false;
}

function refleshImg() {
    $("#login-verifyCodeImg").attr("src", "https://wappass.baidu.com/cgi-bin/genimage?" + $("#login-vcodestr").val() + "&v=" + Math.random());
}