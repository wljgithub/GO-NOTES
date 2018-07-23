var CheckReg = {
	errmsg: {
		normalString: '',
		buin: '请填写正确的企业号',
		usercount: '',
		need: '不能为空',
		mobile: '请填写正确的手机号码',
		email: '请填写正确的邮箱地址',
		code: '请填写正确的验证码',
		same: '密码不一致',
		mobile_email: '手机或邮箱格式不正确'
	},
	normalString : function(value, element){
		var reg = /^[A-Za-z0-9_]{4,16}$/;
		return reg.test(value);  
	},
	pwd : function(value, element){
		var reg = /^[A-Za-z0-9_\-+=`~!@#\$%\^&\*\(\)\.\,\;\:\'\"\{\}\[\]\?\/\\\|<>]{4,16}$/;
		return reg.test(value);  
	},
	buin : function(value){
		var reg = /^[1-9][0-9]{7}$/;
		//$("#buin_tis").text("请填写8位数的总机号，没有企业号请写0");
		return reg.test(value);   
	}, 
	usercount : function(value){
		var reg = /^[1-9]\d*$/;
		return  reg.test(value);  
	},
	need : function(value){
		return !!$.trim(value)
	},
	mobile :function(value){
		var reg = /^[1-9][0-9]{10}$/;
		return  reg.test(value);
	},
	email : function(value){
		var reg = /^([\w-_]+(?:\.[\w-_]+)*)@((?:[a-z0-9]+(?:-[a-zA-Z0-9]+)*)+\.[a-z]{2,10}(\.[a-z]{2,10})*)$/i;
		return  reg.test(value);
	},
	code : function(value){
		var reg = /^[0-9]{6}$/;
		return  reg.test(value); 
	},
	same : function(value, ele){ 
		return !!ele && value == $("#"+ele.data("same")).val()
	},
	mobile_email : function(val) {
		return CheckReg.mobile(val) || CheckReg.email(val);
	}
};

function getCaptcha() {
	$.get("captcha", {_:new Date().getTime()}, function (d) {
		$("#captchaId").val(d.id);
		$("#captchaImg").attr("src", d.src);
	});
}

function FormReg(form){
	var ins = form.find(".input-group:not('.hide')");
	var err = 0;
	$.each(ins, function(){
		var t = $(this);
		var val = $.trim(t.find('.input-item').val())
		var m = t.data('check');
		if(m && !CheckReg[m](val, t)){
			var errmsg = CheckReg.errmsg[m];
			if (t.data('err')) {
				errmsg = t.data('err')
			}
			t.addClass('err').find('.input-item').focus().end().find('.input-err').text(errmsg);
			err++;
		}else{
			t.removeClass('err').find('.input-err').text('');
		}
	});
	return err === 0;
}

var code_getting  = false;
var code_time = 60;
function codeTimerStart(id, t) {
	if (t) {
		$(document.getElementById(id)).text("重新获取(" + t + ")").attr('disabled', true)
		t--;
		code_timeout = setTimeout(function() {
			codeTimerStart(id, t)
		}, 1000)
	} else {
		// document.getElementById(id).innerText = "获取验证码"
		$(document.getElementById(id)).text("获取验证码").attr('disabled', false)
		code_getting = false;
	}
}

$("#btnGetCode").click(function(){
	if (code_getting) {
		return
	}
	code_getting = true;
	var val = $('#check-contact').text();
	if (CheckReg.mobile(val)) {
		$.post('sendsmscode.html', {mobile: val,type:1}, function(data){
			codeTimerStart("btnGetCode", code_time)
			if (data != 0) {
				alert(data)
			}
		})  
	} else if (CheckReg.email(val)) {
		$.post("sendmailcode.html", {email: val}, function(d){
			codeTimerStart("btnGetCode", code_time)
			if( d == "0" ) {
				$("#lb_mail").text("验证码已发送到您的邮箱");
			}else if(d=="1"){
				$("#in_email").focus();
				$("#lb_mail").text("抱歉，您填写的邮箱不存在");
			}else{
				setCodeBtnBeable();
				$("#lb_mail").text("系统错误，请稍后重试或者咨询客服");
			}
		}) 
	}
})

$("#btnGetRegCode").click(function(){
	if (code_getting) {
		return
	}
	code_getting = true;
	// $.get(window.location.origin + '/trial/gencode?mobile=' + checked_mobile, function(){
	var val = $('#f-mobile').val();
	$.post('sendsmscode.html', {mobile: val}, function(){
		codeTimerStart("btnGetRegCode", code_time)
	})
})

$("#cloud-pwd-find-step1").submit(function(){
	if (!FormReg($(this))) {
		return false
	}
	var t = $(this);
	var val = t.find('#contact').val();
	if (CheckReg.mobile(val)) {
		t.find('#contact-type').val('mobile')
	} else if (CheckReg.email(val)) {
		t.find('#contact-type').val('email')
	}
})

$("#cloud-pwd-find-step3").submit(function(){
	if (!FormReg($(this))) {
		return false
	}
	var p = sha256_digest(hex_md5($(this).find("#pwd").val()));
	$(this).find("#pwd").val(p);
})

jQuery(document).ready(function($) {

	// $(".input-item").blur(function(){
	// 	var t = $(this);
	// 	var m = t.parent().data('check');
	// 	if(m && !CheckReg[m](t.val(), t)){
	// 		var errmsg = CheckReg.errmsg[m];
	// 		if (t.parent().data('err')) {
	// 			errmsg = t.parent().data('err')
	// 		}
	// 		t.nextAll('.input-err').text(errmsg).parent().addClass('err')
	// 	}else{
	// 		t.nextAll('.input-err').text('').parent().removeClass('err')
	// 	}
	// });	

	//禁止input password 输入空格
	$("input[type='password']").keypress(function(e){
		var keyValue;
		if (window.event) { //IE
			keyValue = e.keyCode;
		} else if (e.which) { //Netscape/Firefox/Opera
			keyValue = e.which;
		}
		if (keyValue == 32) {
			return false;
		}
	})

	$("#create_admin").submit(function(){
		if (!FormReg($(this))) {
			return false
		}
		var pwd = hex_md5($('#password').val());
		$('#password').val(sha256_digest(pwd));
		//return false;
	});

	$("#cloud-login").submit(function(){
		var t = ($(this))
		if (!FormReg(t)) {
			return false
		}
		
		var pwd = hex_md5(t.find("#password").val());
		pwd = $("#uid").val() + sha256_digest(pwd);
		form = { 
			mobile:$.trim(t.find("#mobile").val()), 
			pwd: sha256_digest(pwd),
			id: t.find("#captchaId").val(),
			digits: t.find("#captcha").val()
		};
		var jqxhr = $.post('cloud_login.html', form, function (d) {
			switch(d.code) {
				case 0:{
					window.location.href= "./";
					return;
				}break;
				case 8:
				case 19:{
					t.find("#mobile").focus().nextAll('.input-err').text('企业号或密码错误').parent().addClass('err')
				}break;
				case 16:{
					t.find("#captcha").focus().nextAll('.input-err').text('验证码错误').parent().addClass('err')
				}break;
				default:{
					t.find("#mobile").nextAll('.input-err').text('系统错误，请稍后登录').parent().addClass('err')
				}break;
			}
			t.find("#captcha").parent().removeClass('hide').next().removeClass('hide')
			getCaptcha();
		});
		jqxhr.fail(function(jgxhr, status){
			if(window["console"]){
				console.error("request error: ", status);
			}
			getCaptcha();
		});
		return false
	});
});