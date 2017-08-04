function en() {
    var formParam = $("#login-formWrapper").serialize();
    document.getElementById('login-submit').className = "pass-button-full pass-button-full-disabled";
    document.getElementById('login-submit').disabled = "disabled";

    document.getElementById('login-verify').style.display = "none";
    document.getElementById('input-code').style.display = "none";
    document.getElementById('sdaf').style.display = "none";

    $.ajax({
        type: 'post',
        url: "/execBaiduLogin",
        data: formParam,
        dataType: 'json',
        error: function(XMLHttpRequest, textStatus, errorThrown) {
            alert("发送错误: " + XMLHttpRequest.status);
            reborn('login-submit');
        },
        success: function(data) {
            switch (data.errInfo.no) {
                case "0":
                    document.getElementById('sdaf').style.display = "block";
                    printCookies(data)
                    document.getElementById('login-verifyWrapper').style.display = "none";
                    break;
                case "400101": // 开启登录保护
                case "400023": // 需要验证手机或邮箱
                    document.getElementById('login-verify').style.display = "";
                    document.getElementById('my-phone').innerHTML = data.data.phone;
                    document.getElementById('my-email').innerHTML = data.data.email;
                    document.getElementById('token').value = data.data.token;
                    document.getElementById('u').value = data.data.u;
                    document.getElementById('login-verifyWrapper').style.display = "none";
                    break;
                case "500001":
                case "500002":
                    document.getElementById('login-verifycode').value = "";
                    document.getElementById('login-verifyWrapper').style.display = "";
                    document.getElementById('login-vcodestr').value = data.data.codeString;
                    document.getElementById('login-verifyCodeImg').src = "https://wappass.baidu.com/cgi-bin/genimage?" + data.data.codeString;
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
    token = document.getElementById('token').value;
    document.getElementById('verify-type').value = type;

    document.getElementById('phone-submit').className = "pass-button-full pass-button-full-disabled";
    document.getElementById('phone-submit').disabled = "disabled";
    document.getElementById('email-submit').className = "pass-button-full pass-button-full-disabled";
    document.getElementById('email-submit').disabled = "disabled";

    $.ajax({
        type: 'post',
        url: "/sendCode",
        data: {
            type: type,
            token: token,
        },
        dataType: 'json',
        error: function(XMLHttpRequest, textStatus, errorThrown) {
            alert("发送数据时出错: " + XMLHttpRequest.status);
            reborn('phone-submit');
            reborn('email-submit');
        },
        success: function(data) {
            if (data.error == undefined) {
                document.getElementById('input-code').style.display = "";
                document.getElementById('sended-status').innerHTML = data.msg;
            } else {
                alert(data.error)
            }
            reborn('phone-submit');
            reborn('email-submit');
        },
    });

}

function sendCode() {
    type = document.getElementById('verify-type').value;
    var formParam = $("#input-code").serialize();

    $.ajax({
        type: 'post',
        url: "/execVerifiedLogin",
        data: formParam,
        dataType: 'json',
        error: function(XMLHttpRequest, textStatus, errorThrown) {
            alert("发送错误: " + XMLHttpRequest.status);
        },
        success: function(data) {
            switch (data.errInfo.no) {
                case "0":
                    document.getElementById('input-code').style.display = "";
                    document.getElementById('login-verify').style.display = "none";
                    document.getElementById('input-code').style.display = "none";
                    if (data.data.bduss == "") {
                        document.getElementById('code-msg').innerHTML = "验证成功，请在本页面重新登录";
                    } else {
                        document.getElementById('sdaf').style.display = "";
                        printCookies(data);
                    }
                    break;
                default:
                    document.getElementById('code-error').innerHTML = "ErrorCode: " + data.errInfo.no + ", ErrorMsg: " + data.errInfo.msg;
                    break;
            }
            reborn();
        },
    });
}

function printCookies(data) {
    document.getElementById('USERNAME').value = document.getElementById('login-username').value;
    document.getElementById('input-code').style.display = "none";
    document.getElementById('BDUSS').value = data.data.bduss;
    document.getElementById('PTOKEN').value = data.data.ptoken;
    document.getElementById('STOKEN').value = data.data.stoken;
    document.getElementById('COOKIE').value = data.data.cookieString;
}

function reborn(id) {
    document.getElementById(id).className = "pass-button-full";
    document.getElementById(id).disabled = false;
}

function refleshImg() {
    document.getElementById('login-verifyCodeImg').src = "https://wappass.baidu.com/cgi-bin/genimage?" + document.getElementById('login-vcodestr').value + "&v=" + Math.random();
}