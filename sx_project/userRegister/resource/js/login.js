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


jQuery(document).ready(function($) {

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

	
});