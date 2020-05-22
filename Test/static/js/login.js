$("#login").click(function() {
    //获取用户名
    var username = $("#username").val();
    console.log(46546)
    //获取用户密码
    var password = $("#password").val();
    //是否记住密码
    var persit = $("#persit").is(":checked")
    // console.log(persit)
    $.post("/login",{username:username,password:password},function(data){
        if(data =="suc"){
            //登录成功
            alert("登录成功！")
        }
    })
})
