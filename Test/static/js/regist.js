$("#regist").click(function() {
	// alert("login")
	//获取用户名
	var username = $("#username").val();
	//获取用户密码
	var password = $("#password").val();
	//获取其他密码
	var repeatpass = $("#repeatpass").val();
	//进行比较
	if(password !== repeatpass){
        var mymessage=confirm("两次输入密码不一致，是否重新注册？");
        if(mymessage==true)
        {
            $("#username").val("")
            $("#password").val("")
            $("#repeatpass").val("")
        }
    }else{
		//提交进行相关验证
		$.post("/regist",{username:username,password:password},function(data){
			if(data == "suc"){
			    //跳转到登录也
                window.location.href = "/login"
            }
		})
	}
})
