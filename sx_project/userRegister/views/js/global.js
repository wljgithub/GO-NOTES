String.prototype.gblen = function() {  
    var len = 0;  
        for (var i=0; i<this.length; i++) {  
            if (this.charCodeAt(i)>127 || this.charCodeAt(i)==94) {  
                len += 2;  
            } else {  
                len ++;  
            }  
        }  
    return len;  
}

function GetRequest() {   
    var url = location.search; //获取url中"?"符后的字串   
    var theRequest = new Object();   
    if (url.indexOf("?") != -1) {   
       var str = url.substr(1);   
       strs = str.split("&");   
       for(var i = 0; i < strs.length; i ++) {   
          theRequest[strs[i].split("=")[0]]=unescape(strs[i].split("=")[1]);   
       }   
    }   
    return theRequest;   
}  

//IE不支持new date()
function NewDate(str) { 
    if(str==null)
     return false;
    str = str.split('-'); 
    var date = new Date(); 
    date.setUTCFullYear(str[0], str[1] - 1, str[2]); 
    date.setUTCHours(0, 0, 0, 0); 
    return date; 
}

function stopPropagation(e) {
    if(document.all){  //只有ie识别
        e.cancelBubble=true;
    }else{
        e.stopPropagation();
    }
}
function CheckLoged(data){
    if(typeof(data) == 'string' && -1 != data.indexOf('<title>登录-有度管理后台</title>')){
        window.location.href = '/userportal/login.html';
    }
}

var Browser = function() {
    var browserInfo = window.navigator.userAgent;
    var browserFlags =  {};

    browserFlags.ISFF = browserInfo.indexOf('Firefox') != -1;
    browserFlags.ISOPERA = browserInfo.indexOf('Opera') != -1;
    browserFlags.ISCHROME = browserInfo.indexOf('Chrome') != -1;
    browserFlags.ISSAFARI = browserInfo.indexOf('Safari') != -1 && !browserFlags.ISCHROME;
    browserFlags.ISWEBKIT = browserInfo.indexOf('WebKit') != -1;

    browserFlags.ISIE = browserInfo.indexOf('Trident') > 0 || navigator.userAgent.indexOf('MSIE') > 0;
    browserFlags.ISIE6 = browserInfo.indexOf('MSIE 6') > 0;
    browserFlags.ISIE7 = browserInfo.indexOf('MSIE 7') > 0;
    browserFlags.ISIE8 = browserInfo.indexOf('MSIE 8') > 0;
    browserFlags.ISIE9 = browserInfo.indexOf('MSIE 9') > 0;
    browserFlags.ISIE10 = browserInfo.indexOf('MSIE 10') > 0;
    browserFlags.ISOLD = browserFlags.ISIE6 || browserFlags.ISIE7 || browserFlags.ISIE8 || browserFlags.ISIE9; // MUST be here

    browserFlags.ISIE11UP = browserInfo.indexOf('MSIE') == -1 && browserInfo.indexOf('Trident') > 0;
    browserFlags.ISIE10UP = browserFlags.ISIE10 || browserFlags.ISIE11UP;
    browserFlags.ISIE9UP = browserFlags.ISIE9 || browserFlags.ISIE10UP;

    return browserFlags;
}();

var num2date = function(data, type){
    var date = new Date(data*1000);
    var year = date.getFullYear();
    var month = "0" + (date.getMonth() + 1);
    var day = "0" + date.getDate();
    // Hours part from the timestamp
    var hours = "0" + date.getHours();
    // Minutes part from the timestamp
    var minutes = "0" + date.getMinutes();
    // Seconds part from the timestamp
    var seconds = "0" + date.getSeconds();

    var month = month.substr(month.length - 2, 2)
    var day = day.substr(day.length - 2, 2)
    var hours = hours.substr(hours.length - 2, 2)
    var minutes = minutes.substr(minutes.length - 2, 2)
    var seconds = seconds.substr(seconds.length - 2, 2)

    // Will display time in 10:30:23 format
    switch(type) {
        case 'full': {
            return year+ '-' + month + '-' + day + ' ' + hours + ':' + minutes + ':' + seconds;
        }break;
        case 'datetime': {
            return year+ '-' + month + '-' + day + ' ' + hours + ':' + minutes
        }break;
        case 'time': {
            return hours + ':' + minutes
        }break;
        default:{
            return year+ '-' + month + '-' + day
        }break;
    }
}

/**
 * 
 * 将字节转化为有单位的字符串
 * @param {int} fileSizeInBytes 
 * @returns 
 */
function parseFileSize(fileSizeInBytes){
    var i = -1;
    var byteUnits = ['MB', 'GB', 'TB', 'PB', 'EB', 'ZB', 'YB'];
    do {
        fileSizeInBytes = fileSizeInBytes / 1024;
        i++;
    } while (fileSizeInBytes > 1024);

    return Number(fileSizeInBytes.toFixed(2)) + byteUnits[i];
}


/**
 * 数组合并去重
 * 
 * @returns 
 */
Array.prototype.unique = function() {
    var a = this.concat();
    for(var i=0; i<a.length; ++i) {
        for(var j=i+1; j<a.length; ++j) {
            if(a[i] === a[j])
                a.splice(j--, 1);
        }
    }
    return a;
};

/**
 * 数组对象合并去重
 * 
 * @param {any} key 根据此key 判断是否重复
 * @returns 
 */
Array.prototype.uniqueObj = function(key) {
    var a = this.concat();
    for(var i=0; i<a.length; ++i) {
        for(var j=i+1; j<a.length; ++j) {
            if(a[i][key] === a[j][key])
                a.splice(j--, 1);
        }
    }
    return a;
};

var CheckReg = {
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
        $("#buin_tis").text("请填写8位数的总机号，没有企业号请写0");
        return "0" == value || reg.test(value);   
    }, 
    usercount : function(value){
        var reg = /^[1-9]\d*$/;
        return  reg.test(value);  
    },
    need : function(value){
        return !!value
    },
    mobile :function(value){
        var reg = /^[1-9][0-9]{10}$/;
        return  reg.test(value);
    },
    email : function(value){
        var reg = /^([\w-_]+(?:\.[\w-_]+)*)@((?:[a-z0-9]+(?:-[a-zA-Z0-9]+)*)+\.[a-z0-9]{2,10}(\.[a-z]{2,10})*)$/i;
        return  reg.test(value);
    },
    code : function(value){
        var reg = /^[0-9]{6}$/;
        return  reg.test(value); 
    },
    same : function(value, ele){ 
        return !!ele && value == $("#"+ele.data("same")).val()
    },
    link : function(value){
        var reg = /(http(s?):\/\/|www[0-9]*\.|wap\.|[0-9]+g\.)([a-zA-Z0-9\-_]+)((:[0-9]+)?)|(\.[a-zA-Z0-9\-_]+)+[a-zA-Z0-9\-_.:\/+%?&=#\(\)|;~!@$\^\\`]*/ig
        return reg.test(value)
    },
    ip: function (value) {
        var reg = /\d+\.\d+\.\d+\.\d+/;
        return reg.test(value)
    },
    addr: function(value) {
        var reg = /^(?=^.{3,255}$)[a-zA-Z0-9][-a-zA-Z0-9]{0,62}(\.[a-zA-Z0-9][-a-zA-Z0-9]{0,62})+$/;
        return reg.test(value)
    },
    date: function(value) {
        var reg = /\d{4}\-\d{2}\-\d{2}/;
        return reg.test(value)
    },
    pcClient: function(value) {
        var reg = /\.exe$/;
        return reg.test(value)
    }
};

var Filetype = ['txt', 'xml', 'pdf', 'zip', 'doc', 'ppt', 'xls', 'docx', 'pptx', 'xlsx']

// JavaScript Document

$(function () {
    var aside = $('.aside');
    var header_menu = $(".header-menu");
    var main_content = $('.main-content');
    var md_pwd = $('#md-pwd');
    var md_pwd_form = md_pwd.find('#md-pwd-form')
    var md_mail = $('#md-mail');
    var md_mail_form = md_mail.find('#md-mail-form')
    var aside_height = aside.height();
    /*窗口拉动*/

    // var init_window = function(){
    //     var h = $(window).height();
    //     if(aside_height + 50 > h){
    //         h = aside_height
    //     } else {
    //         h -= 50;
    //     }
    //     aside_height <= h && ( aside.height(h));
    // }
    // init_window();

    // $(window).resize(init_window)

    aside.find("dt").click(function(){
        var isClose = $(this).parent().hasClass('close');
        if (isClose) {
            $(this).parent().removeClass('close')
            SessStorage.remove($(this).parent().attr('class'))
        } else {
            $(this).parent().addClass('close')
            SessStorage.set($(this).parent().attr('class').split(' ')[0], 1)
        }
    })


    /*修改密码*/
    $('#changeSysPwd').click(function () {
        md_pwd.show();
    });
    md_pwd.find('.close').click(function () {
        md_pwd.find('form')[0].reset();
        md_pwd.hide();
    });
    /*邮箱配置*/
    $('#changeEmail').click(function () {
        md_mail.show();
        md_mail.find("#mail-host").focus();
    });
    md_mail.find('.close').click(function () {
        md_mail.find('form')[0].reset();
        md_mail.hide();
    });

    /*box收缩伸展*/

    $('.show_btn').click(function () {
        if ($(this).text() == "一") {
            $(this).text("+");
            $(this).parents('.infor-box').find('.show-body').slideUp(500);
        } else {
            $(this).text("一");
            $(this).parents('.infor-box').find('.show-body').slideDown(500);
        }
    })


    md_pwd_form.submit(function () {
        var t = $(this);
        var p1 = $.trim(t.find('#pwd_old').val())
        var p2 = $.trim(t.find('#pwd_new').val())
        var p3 = $.trim(t.find('#pwd_new2').val());
        var err = 0;
        if( !p1 ){
            err = 1
            t.find('#pwd_old').parent().siblings('.tis').text('请填写旧密码')
            t.find('#pwd_old').focus();
        } else if( !CheckReg.pwd(p1) ){
            err = 1
            t.find('#pwd_old').parent().siblings('.tis').text('密码由4至16位常规字符')
            t.find('#pwd_old').focus();
        } else {
            t.find('#pwd_old').parent().siblings('.tis').text('')
        }
        if( !p2 ){
            err += 1
            t.find('#pwd_new').parent().siblings('.tis').text('请填写新密码')
            err == 2 || t.find('#pwd_new').focus();
        } else if( !CheckReg.pwd(p2) ){
            err += 1
            t.find('#pwd_new').parent().siblings('.tis').text('密码由4至16位常规字符')
            err == 2 ||  t.find('#pwd_new').focus();
        } else {
            t.find('#pwd_new').parent().siblings('.tis').text('')
        }
        if( !p3 ){
            err += 1
            t.find('#pwd_new2').parent().siblings('.tis').text('请重复新密码')
            err == 3 || t.find('#pwd_new2').focus();
        } else if( p2 !== p3 ){
            err += 1
            t.find('#pwd_new2').parent().siblings('.tis').text('密码不一致')
            err == 3 || t.find('#pwd_new2').focus();
        } else {
            t.find('#pwd_new2').parent().siblings('.tis').text('')
        }
        
        if(!err) {
            var pwd = hex_md5(p1);
            pwd = sha256_digest(pwd);
            npwd = sha256_digest(hex_md5(p2));
            var dat = {
                old_pwd: sha256_digest(pwd),
                new_pwd: npwd
            };
            $.post("/userportal/modifypwd", dat, function (msg) {
                if (msg == "0") {
                    t.find('#pwd_old').parent().siblings('.tis').text('密码错误').show()
                    t.find('#pwd_old').focus();
                } else if (msg == "1") {
                    YdResultBox.show('修改成功')
                    md_pwd.hide();
                    setTimeout(function(){
                        window.location.href = "logout"
                    }, 2000);                    
                }
            }).always(function(data){
                CheckLoged(data)
            });
        }
        return false;
    })

    md_mail_form.submit(function () {
        var t = $(this);
        var host = $.trim(t.find('#mail-host').val())
        var port = $.trim(t.find('#mail-port').val())
        var addr = $.trim(t.find('#mail-addr').val())
        var pwd = $.trim(t.find('#mail-pwd').val())
        var username = $.trim(t.find('#mail-username').val())
        var err = 0;
        if (!host) {
            t.find('#mail-host').parent().siblings('.tis').text('请填写服务器地址')
            err++;
        }
        if (!port) {
            t.find('#mail-port').parent().siblings('.tis').text('请填写服务器端口')
            err++;
        }
        if (!addr) {
            t.find('#mail-addr').parent().siblings('.tis').text('请填写邮箱地址')
            err++;
        } else if (!CheckReg.email(addr)) {
            t.find('#mail-addr').parent().siblings('.tis').text('邮箱格式错误')
            err++;
        }
        if (!pwd) {
            t.find('#mail-pwd').parent().siblings('.tis').text('请填写邮箱密码')
            err++;
        }
        if (!username) {
            t.find('#mail-username').parent().siblings('.tis').text('请填写发件人')
            err++;
        }
        if(!err) {
            var d = {
                host: host,
                port: port,
                addr: addr,
                password: pwd,
                username: username
            }
            $.post("/userportal/configmail", d, function (data) {
                CheckLoged(data)
                if (0 === data.code) {
                    YdResultBox.show('配置成功')
                    md_mail.hide();
                    GMail.Host = host
                    GMail.Port = port
                    GMail.Addr = addr
                    GMail.Username = username
                } else {
                    YdResultBox.show('配置失败')
                }
            }).always(function(data){
                CheckLoged(data)
            });
        }
        return false;
    })

    //menu highline
    var Aside_menu = {
        
        'app': ['app_msg_create', 'app_web_create', 'app_app_create', 'app_base', 
         'app_send', 'app_reply', 'app_menu', 'app_app', 'app_menu_reply', 'app_api', 'app_callback_api',  'sysmsg', 'app_info', 'app_create', 'exmail', 'exmail_connection', 'exmail_alertcallback', 'exmail_org', 'app_tmpl', 'sysfileapp', 'meeting_room', 'report_tpl', 'report_tpls', 'report_tpl_add'],
               
        
        'service': ['service_doc'],
        
        'dashboard': [],
        'org': ['explain'],
        'group_manage': [],
        // 管理工具
        'tools':[],
        'app_org_sync': ['app_org_sync', 'ldap_info', 'app_org'],
        'clean': [],
        'app_plugin': ['sms', 'audio', 'audio-log', 'msg-monitor', 'msg-monitor-power'],
        'mobile_online': ['log', 'log_operate'],
        'upgrade': ['upgrade_mac', 'client_download_page'],
        'conncode': [],
        
        // 我的企业
        'license': [],
        'work_panel': [],
        'role': [],
        'org_visible': ['hide_dept', 'private_dept', 'visible_all', 'custom_userattr'],
        'rca': ['rca_visible', 'rca_visible_rule', 'rca_center'],
        'client_info': [],
        'client_setting': ['permission_mobile', 'msgsize', 'client_power', 'server_push_msg'],
        'server': [],
        'chat_setting': []
    }
    var Menu = {
        'dashboard': [],
        'org': [],
        'group_manage': [],
        'app': [],
        'tools' : ['app_org_sync', 'sysmsg', 'clean', 'app_plugin', 'mobile_online', 'upgrade', 'conncode', 'msg-monitor', 'meeting_room'],
        'ent' : ['license', 'role', 'org_visible', 'rca', 'work_panel', 'server', 'client_setting', 'client_info', 'chat_setting']
    }
    var MenuHighLine = function () {
        var href = location.pathname;
        if( href.indexOf('.html') ){
            href = href.substring(12, href.length - 5).split("/")[0]
        } else {
            href = href.substring(12, href.length).split("/")[0]
        }
        var m = '';
        var hm = '';
        _.forEach(Aside_menu, function(v, k) {
            if(m) { return }
            if (k === href){
                m = k;
                return;
            }
            _.forEach(v, function(v1) {
                if(m) { return }
                if (v1 === href){
                    m = k;
                }
            })
        })
        if(m) { 
            _.forEach(Menu, function(v, k) {
                if(hm) { return }
                if (k === m){
                    hm = k;
                    return;
                }
                _.forEach(v, function(v1) {
                    if(hm) { return }
                    if (v1 === m){
                        hm = k;
                    }
                })
            })
            header_menu.find('.header-menu-item').each(function(index, el) {
                if($(this).data('menu') == hm) {
                    $(this).addClass('on')
                    return false;
                }
            });
            aside.find('dd').each(function(index, el) {
                if($(this).data('menu') == m) {
                    $(this).addClass('on')
                    return false;
                }
            });
        }
    }();
    // var MenuSlideTog = function() {
    //     aside.find('dl').each(function(index, el) {
    //         var klass = $(el).attr('class').split(' ')[0]
    //         if (SessStorage.get(klass)) {
    //             // $(el).removeClass('close')
    //             $(el).addClass('close')
    //         } else {
    //             // $(el).addClass('close')
    //         }
    //     });
    // }()

    Browser.ISOLD && $(".ie-alert").show()
    $("#ie-modal-tip,#ie-modal .close").click(function(){
        $("#ie-modal").hide();
    })
});

var LoadBox = {
    box : $("#loadBox"),
    show : function (txt, callback) {
        txt && LoadBox.box.find('#loadtxt').text(txt);
        LoadBox.box.show()
        callback && callback();
    },
    hide : function (callback) {
        LoadBox.box.hide()
        callback && callback();
    }
}
var Storage = {
    get : function (name) {
        if (!window.localStorage) {return false}
        return  JSON.parse( window.localStorage.getItem(name))
    },
    set : function (name, value){
        if (!window.localStorage) {return false}
        window.localStorage.setItem(name, value);
    },
    remove: function(name) {
        if (!window.localStorage) {return false}
        window.localStorage.removeItem(name)
    },
    clear: function(){
        if (!window.localStorage) {return false}
        window.localStorage.clear();
    }
}
var SessStorage = {
    get : function (name) {
        if (!window.sessionStorage) {return false}
        return  JSON.parse( window.sessionStorage.getItem(name))
    },
    set : function (name, value){
        if (!window.sessionStorage) {return false}
        window.sessionStorage.setItem(name, value);
    },
    remove: function(name) {
        if (!window.sessionStorage) {return false}
        window.sessionStorage.removeItem(name)
    },
    clear: function(){
        if (!window.sessionStorage) {return false}
        window.sessionStorage.clear();
    }
}

function closeModal(t) {
    $(t).parents('.modal').hide();
}
function closeAlertBox(t) {
    $(t).parents('.alertBox:first').hide()
}


(function ($) { 
    $.fn.youduResultBox = function (options) { 
        var defaults = { 
            //各种属性及其默认值 
            btns: [{
                title: '确定',
                nohide: 0, // 按钮操作是否关闭窗口 0 不关闭  1关闭
                main: 0, // 是否主要按钮
                callback: null
            }],
            main_btn: 'ydbtn-primary',
            tpl: '<div class="modal modal-alert" id="<&id&>"><div class="box_inner"><div class="hd"><h4 class="modal-hd-title"></h4><a href="javascript:;" class="close"></a></div><div class="bd"><h3 class="modal-bd-title"></h3><p class="modal-bd-content"></p></div><div class="ft"><&btns&></div></div></div>'
        }; 
        var icon_type = ['warning', 'success', 'error'];
        var _btn = '<button class="ydbtn <&main&>" type="button" data-bin="<&bin&>"><&title&></button>'
        var opts = $.extend(defaults, options); 
        var id = "yd-rbox-" + new Date().getTime() + (Math.random() * 100000 >> 0);
        var _that = this;
        var show_callback = null;
        
        var bindEvent = function(){
            _that.children("#"+id).find('.close').click(function(){
                _that.hide();
            })
            _that.children("#"+id).find('button').click(function(){
                var bin = $(this).data('bin');
                opts.btns[bin].callback && opts.btns[bin].callback();
                opts.btns[bin].nohide || _that.hide();
                $(this).hasClass(opts.main_btn) && show_callback && show_callback();
            })
        }

        var init = function(){
            var list = opts.btns;
            var bts = '';
            for (var i=0; i<list.length; i++) {
                var self = list[i];
                var on = self.main ? opts.main_btn : '';
                bts += _btn.replace(/<&title&>/ig, self.title).replace(/<&main&>/ig, on).replace(/<&bin&>/ig, i);
            }
            var h = opts.tpl.replace(/<&id&>/ig, id).replace(/<&btns&>/ig, bts)
            _that.append(h);
            setTimeout(bindEvent, 0)
        }();

        this.hide = function(){
            _that.children("#"+id).hide();
        }
        this.show = function(param){
            var p = {
                title: '提示',
                txt: '',
                note: '',
                callback: null,
                type: 0 // 0 warning, 1 success, 2 error
            }
            if (typeof(param) == 'string') {
                p.txt = param
            } else {
                p = $.extend(p, param); 
            }
            show_callback = p.callback 
            _that.children("#"+id).removeClass('warning success error')
            if (p.note) {
                _that.children("#"+id)
                .removeClass('single').addClass(icon_type[p.type])
                .find('.modal-bd-title').html(p.txt || '').end()
                .find('.modal-bd-content').html(p.note).end()
                .find('.modal-hd-title').html(p.title).end()
                .show()
            } else {
                _that.children("#"+id)
                .addClass('single').addClass(icon_type[p.type])
                .find('.modal-bd-title').html(p.txt || '').end()
                .find('.modal-hd-title').html(p.title).end()
                .show()
            }
            
        }
        return this;
    }; 
})(jQuery); 
var YdResultBox = $('body').youduResultBox();
/**
 * TogResultBox.show({
    txt: '确定禁用该应用？',
    callback: function(){
        app_tog.set(0)
        CommonApp.toggleAppSubmit(AppId, app_tog.get())
    }
})
 */
var TogResultBox = $('body').youduResultBox({
    btns: [
        {
            title: '确定',
            nohide: 0, 
            main: 1
        },{
            title: '取消'
        }
    ]
});

(function ($) { 
    $('body').append('<div class="alertAuto" id="alertAuto"></div>')
    $.fn.alertAuto = function (options) { 
        var defaults = { 
            //各种属性及其默认值 
            showtime: 3000, // 显示时间 3s
            success: '<div class="alert alert-success">text</div>',
            error: '<div class="alert alert-error">text</div>',
            normal: '<div class="alert">text</div>'
        }; 
        var timer = null;
        var opts = $.extend(defaults, options); 
        var _that = this;

        this.hidden = function(){
            _that.html('');
        }
        
        this.succeed = function(text){
            var h = opts.success.replace(/text/ig, text)
            $(window).scrollTop(0)
            this.html(h)
            timer && clearTimeout(timer)
            timer = setTimeout(this.hidden, opts.showtime)
        }
        this.failed = function(text){
            var h = opts.error.replace(/text/ig, text)
            this.html(h)
            timer && clearTimeout(timer)
            timer = setTimeout(this.hidden, opts.showtime)
        }
        this.show = function(text){
            var h = opts.normal.replace(/text/ig, text)
            this.html(h)
            timer && clearTimeout(timer)
            timer = setTimeout(this.hidden, opts.showtime)
        }
        return this;
    }; 
})(jQuery); 

(function ($) { 
    $.fn.ListBtn = function (options) { 
        var defaults = { 
            width: 0,
            ulWidth: 0,
            title: '',
            items: [
                {
                    title: "",
                    callback: ""
                }
            ]
        }; 
        var opts = $.extend(defaults, options); 
        var _that = this;
        
        var bindEvent = function(){
            _that.find('.tog, .pbtn').click(function(evt){
                if (_that.children('ul:hidden').length) {
                    $('body').find('.list-btn ul').slideUp(120)
                    _that.children('ul').slideDown(120)
                } else {
                    $('body').find('.list-btn ul').slideUp(120)
                }
                evt.stopPropagation();
            })
            // opts.callback && _that.find('.pbtn').click(function(evt){
            //     evt.stopPropagation();
            //     opts.callback()
            // })
            _that.find("li").each(function($index){
                opts.items[$index].callback && $(this).click(function(){
                    $('body').find('.list-btn ul').slideUp(120)
                    opts.items[$index].callback();
                })
            })
            if (!window.hasListBtn) {
                $('body').click(function(){
                    $('.list-btn').children('ul').slideUp(120)
                })
                window.hasListBtn = 1
            }
        }

        var init = function(){
            var list = opts.items;
            var h = '<a href="javascript:;" class="pbtn">'+opts.title+'</a><i class="tog"></i><ul>'
            for (var i=0; i<list.length; i++) {
                h += '<li><a href="javascript:;">'+ list[i]['title'] +'</a></li>'
            }
            h += '</ul>'
            _that.html(h);
            _that.find("li:eq(0)").addClass('first');
            if (opts.width) {
                _that.css("width", opts.width)
            }
            if (opts.ulWidth) {
                _that.find("ul").css("width", opts.ulWidth)
            } else {
                _that.find("ul").css("width", opts.width || "100%")
            }
            bindEvent();
        }();
        return this;
    }; 
})(jQuery); 


(function ($) { 
    $.fn.youduRadio = function (options) { 
        var defaults = {
            active_class: 'on',
            tpl : '<label class="<&active_class&>"" for="<&for&>">'
                    +'<input type="radio" name="<&name&>" id="<&rd_id&>" value="<&val&>" <&checked&> ><&title&></label>',
            group : [],
            click_callback: null
        }; 
        var opts = $.extend(defaults, options); 
        var _that = this;

        var bindEvent = function(){
            _that.find('input').change(function(){
                $(this).parent().addClass(opts.active_class);
                $(this).parent().siblings().removeClass(opts.active_class);
                opts.click_callback && opts.click_callback();
            })
        }

        var init = function(){
            var list = opts.group;
            var h = ''
            for (var i=0; i<list.length; i++) {
                var chk = list[i].cheched ? 'checked' : ''
                var on = list[i].cheched ? 'on' : ''
                var id = 'ydrd' + i;
                h += opts.tpl.replace(/<&active_class&>/ig, on).replace(/<&rd_id&>/ig, id).replace(/<&for&>/ig, id).replace(/<&name&>/ig, list[i].name).replace(/<&title&>/ig, list[i].title).replace(/<&val&>/ig, list[i].val).replace(/<&checked&>/ig, chk)
            }
            _that.html(h);
            _that.find("label:eq(0)").addClass('first');
            _that.find("label:last").addClass('last');
            bindEvent();
        }();
        
        this.get = function(){
            return _that.find("input:checked").val();
        }
        this.set = function(v){
            _that.find('input').each(function($index, el){
                var _v = $(this).val();
                if (_v == v) {
                    $(this).attr('checked', true).end()
                    $(this).parent().addClass(opts.active_class);
                    $(this).parent().siblings().removeClass(opts.active_class).children().removeAttr('checked')
                    
                    return false;
                }
            })
            opts.click_callback && opts.click_callback();
        }
        
        return this;
    }; 
})(jQuery); 

(function ($) { 
    $.fn.youduToggle = function (options) { 
        var defaults = {
            active_class: 'on',
            tpl : '<a href="javascript:;" class="<&active_class&>"></a><span><&name&></span>',
            toggle : false,
            tis : {
                on: '已开启',
                off: '已关闭'
            },
            tis_dir: "left", // left  right
            click_callback: null,
            click_submit: null
        }; 
        var opts = $.extend(defaults, options); 
        var _that = this;

        var bindEvent = function(){
            _that.find('a').click(opts.click_submit ? function(){
                opts.click_submit()
            } : function(){
                opts.toggle = !opts.toggle;
                $(this).toggleClass(opts.active_class, opts.toggle);
                $(this).siblings('span').text(opts.toggle ? opts.tis.on : opts.tis.off)
                opts.click_callback && opts.click_callback();
            })
        }

        var init = function(){
            var list = opts.group;
            var h = ''
            var klass = opts.toggle ? opts.active_class : '';
            var name = opts.toggle ? opts.tis.on : opts.tis.off;
            h += opts.tpl.replace(/<&active_class&>/ig, klass).replace(/<&name&>/ig, name)
            opts.tis_dir == 'right' && _that.addClass("tis_right")
            _that.html(h);
            bindEvent();
        }();
        
        this.get = function(){
            return opts.toggle;
        }
        this.set = function(tog){
            opts.toggle = tog
            _that.children('a').toggleClass(opts.active_class, opts.toggle);
            _that.children('span').text(opts.toggle ? opts.tis.on : opts.tis.off)
        }
        
        return this;
    }; 
})(jQuery);

(function ($) { 
    $.fn.youduCheckBox = function (options) { 
        var defaults = {
            disabled_class: 'disabled',
            active_class: 'on',
            tpl : '<div class="check-group  <&disabled_class&>"><span class="chk <&active_class&> <&disabled_class&>"><input class="checkbox" type="checkbox" name="<&name&>" id="<&id&>" <&checked&> ></span><label for="<&id&>"><&title&></label></div>',
            group:[],
            click_callback: null
        }; 
        var opts = $.extend(defaults, options); 
        var _that = this;

        var bindEvent = function(){
            _that.find('span').click(function(){
                var self = $(this)
                if (self.parent().hasClass(opts.disabled_class)) { return }
                self.toggleClass(opts.active_class);
                setTimeout(function() {
                    if (self.hasClass(opts.active_class)) {
                        self.children('input').attr('checked', 'checked')
                    } else {
                        self.children('input').removeAttr('checked')
                    }
                }, 20);
                opts.click_callback && opts.click_callback();
            })
            _that.find('label').click(function(evt){
                var dom = $(this).parent();
                // dom.children('span:first').toggleClass(opts.active_class);
            })
        }

        var init = function(){
            var list = opts.group;
            var h = ''
            for (var i=0; i<list.length; i++) {
                var self = list[i];
                var on = self.checked ? opts.active_class : '';
                var disabled = self.disabled ? opts.disabled_class : '';
                var checked = self.checked ? "checked='checked'" : '';
                h += opts.tpl.replace(/<&disabled_class&>/ig, disabled).replace(/<&active_class&>/ig, on).replace(/<&checked&>/ig, checked).replace(/<&name&>/ig, self.name).replace(/<&id&>/ig, self.id).replace(/<&title&>/ig, self.title)
            }
            _that.html(h);
            bindEvent();
        }();
        
        this.get = function(){
            var p = {};
            _that.find('input').each(function(){
                var id = $(this).attr("id");
                p[id] = $(this).is(':checked')
            })
            return p;
        }
        
        return this;
    }; 
})(jQuery); 

(function ($) { 
    /**
     * youdu tooltip plugin
     * author: dico.zhang
     */
    $.fn.youduTooltip = function (options) { 
        var defaults = {
            dir: 'left_top', //left_top  left_bottom right_top right_bottom
            title: '',
            tis: '',
            distance: 4, //距离
            tpl : '<div class="tooltip <&dir_class&>" id="<&id&>">'
                    +'<div class="tooltip-arrow"></div>'
                    +'<div class="tooltip-inner"><&title&><&tis&></div>',
            evt: 'hover', // hover click
            maxWidth: 200
        }; 
        var opts = $.extend(defaults, options); 
        var _that = this;
        var id = "yd-tooltip-" + new Date().getTime() + (Math.random() * 100000 >> 0);
        var dom = null;

        var bindEvent = function(){
            if (opts.evt == 'hover'){
                _that.mouseenter(function(){
                    setPosition()
                })
                _that.mouseleave(function(){
                    hide()
                })
            }
        }

        var hide = function(){
            dom.css({
                left: '-99999px',
                top: '-99999px'
            })
        }

        var setPosition = function(){
            var os = _that.offset();
            var ow = _that.outerWidth(); // 触发事件 物体宽度
            var oh = _that.outerHeight();// 触发事件 物体高度
            var w = dom.outerWidth(); // 提示框宽度
            var h = dom.outerHeight();// 提示框高度
            
            switch(opts.dir){
                case 'left_top':{
                    dom.css({
                        left: os.left,
                        top: os.top - h - (8 + opts.distance)
                    })
                }break;
                case 'left_bottom':{
                    dom.css({
                        left: os.left,
                        top: os.top + oh + (8 + opts.distance)
                    })
                }break;
                case 'right_top':{
                    var lw = 30 > ow ? 30 : ow
                    dom.css({
                        left: os.left - w + lw,
                        top: os.top - h - (8 + opts.distance)
                    })
                }break;
                case 'right_bottom':{
                    var lw = 30 > ow ? 30 : ow
                    dom.css({
                        left: os.left - w + lw,
                        top: os.top + oh + (8 + opts.distance)
                    })
                }break;
                default:
                break
            }
        }        

        _that.destory = function() {
            dom.remove()
        }

        _that.update = function(title, tis){
            dom.find(".tooltip-title").html(title)
            dom.find(".tooltip-inner").html(tis)
        }

        var init = function(){
            var title = '';
            if (opts.title) {
                title = '<div class="tooltip-title">'+ opts.title +'</div>'
            }
            var h = opts.tpl.replace(/<&id&>/ig, id).replace(/<&dir_class&>/ig, opts.dir).replace(/<&tis&>/ig, opts.tis).replace(/<&title&>/ig, title)
            $('body').append(h); 
            dom = $("#" + id);
            dom.css({
                maxWidth: opts.maxWidth
            })
            // setPosition();   
            // $(window).resize(setPosition)        
            bindEvent()
        }();
        return this;
    }; 
})(jQuery); 


(function ($) { 
    /**
     * youdu checkbox list plugin
     * author: dico.zhang
     */
    $.fn.youduCheckboxList = function (options) { 
        var defaults = {
            items : []
        }; 

        var __item_tpl = '<label class="checkbox-list-item">' + 
                            '<div class="checkbox-list-item-icon"><img src="<&src&>"  alt=""></div>' +
                            '<div class="checkbox-list-item-title"><&title&></div>' +
                            '<div class="checkbox-list-item-chk <&checked&> <&disabled&>" data-value="<&value&>"></div>' +
                            '<input type="checkbox" class="chk" name="<&chk_name&>" value="0">'+
                        '</label>'
        var opts = $.extend(defaults, options); 
        var _that = this;
        var dom = null;

        var bindEvent = function(){
            _that.find('.chk').change(function(){
                var input = $(this)
                if(input.is(':checked')) {
                    $(this).prev('.checkbox-list-item-chk').addClass('on')
                    input.val(1)
                } else {
                    $(this).prev('.checkbox-list-item-chk').removeClass('on')
                    input.val(0)
                }
            })
        }    

        _that.reDraw = function () {
            opts = $.extend(defaults, options); 
            var list = opts.items;
            var h = ''
            var name = "youdu-checkbox-list-" + new Date().getTime() + (Math.random() * 100000 >> 0);
            for (var i=0; i<list.length; i++) {
                var on = list[i].cheched ? 'on' : ''
                var disabled = list[i].disabled ? 'disabled' : ''

                h += __item_tpl.replace(/<&disabled&>/ig, disabled)
                    .replace(/<&src&>/ig, list[i].img)
                    .replace(/<&title&>/ig, list[i].title)
                    .replace(/<&value&>/ig, list[i].val)
                    .replace(/<&checked&>/ig, on)
                    .replace(/<&chk_name&>/ig, name)
            }
            _that.html(h);
            bindEvent();
        }

        var init = function(){
            _that.reDraw(options);
        }();
        return this;
    }; 
})(jQuery); 


(function ($) { 
    $.fn.youduRadioList = function (options) { 
        var defaults = {
            tpl : '<div class="yd-radio"><input type="radio" name="<&name&>" value="<&value&>" id="<&id&>" <&checked&> > <b></b><label for="<&id&>"><&title&></label></div>',
            group : [],
            name: 'yd-rd-name-' + new Date().getTime() + (Math.random() * 100000 >> 0),
            click_callback: null
        }; 
        var opts = $.extend(defaults, options); 
        var _that = this;

        var bindEvent = function(){
            _that.find('input').change(function(){
                opts.click_callback && opts.click_callback(_that.get());
            })
            _that.find('b').click(function(){
                $(this).next('label').click()
            })
        }

        var init = function(){
            var list = opts.group;
            var h = ''
            for (var i=0; i<list.length; i++) {
                var chk = list[i].checked ? 'checked' : '';
                var id = 'yd-rd-list-' + new Date().getTime() + (Math.random() * 100000 >> 0);
                h += opts.tpl.replace(/<&name&>/ig, opts.name).replace(/<&id&>/ig, id).replace(/<&for&>/ig, id).replace(/<&value&>/ig, list[i].value).replace(/<&title&>/ig, list[i].title).replace(/<&checked&>/ig, chk)
            }
            _that.html(h);
            bindEvent();
        }();
        
        this.get = function(){
            return _that.find("input:checked").val();
        }
        this.set = function(v){
            _that.find('input').each(function($index, el){
                var _v = $(this).val();
                if (_v == v) {
                    $(this).attr('checked', true)
                    return false;
                }
            })
        }
        
        return this;
    }; 
})(jQuery); 

var AlertAuto = $("#alertAuto").alertAuto();


jQuery.prototype.serializeObject=function(){  
    var obj=new Object();  
    $.each(this.serializeArray(),function(index,param){  
        if(!(param.name in obj)){  
            obj[param.name]=param.value;  
        }  
    });  
    return obj;  
}; 

var DataTableConfig = {
    bFilter : 0,
    bLengthChange : false,
    bDestroy: true,
    bSort : false,
    pagingType: "full_numbers",
    iDisplayLength: 30,
    autoWidth: false,
    bPaginate: 30,
    sProcessing : "正在获取数据，请稍后...", 
    bAutoWidth: false,
    bInfo: false,
    sSearch: "",
    language: {
        lengthMenu: "每页显示_MENU_条",
        zeroRecords: "暂无记录",
        info: "",
        infoEmpty: "",
        paginate: {
            first:"«",
            previous:"<",
            next:">",
            last:"»"
        }
    }
};


(function($) {
    $.extend({       
        urlGet: function(){
            var aQuery = window.location.href.split("?");  //取得Get参数
            var aGET = {};
            if(aQuery.length > 1)
            {
                var aBuf = aQuery[1].split("&");
                for(var i=0, iLoop = aBuf.length; i<iLoop; i++)
                {
                    var aTmp = aBuf[i].split("=");  //分离key与Value
                    aGET[aTmp[0]] = aTmp[1];
                }
            }
            return aGET;
        }
    })

    
    $("#form-contact").submit(function(){
        if (!GMail.Host) {
            $('#md-mail').show();
            return false;
        }
        var person = $(this).find("#contact_person").val();
        var entname = $(this).find("#contact_entname").val();
        var mobile = $(this).find("#contact_mobile").val();
        if(person == ""){
            AlertAuto.failed('请填写联系人')
            $("#contact_person").focus();
            return false;
        }
        if(entname == ""){
            AlertAuto.failed('请填写企业名称')
            $("#contact_entname").focus();
            return false;
        }
        if(!CheckReg.mobile(mobile)){
            AlertAuto.failed('请填写正确手机号码')
            $("#contact_mobile").focus();
            return false;
        }

        var d = {
            modal: $(this).find("#contact_model").val(),
            person: person,
            entname: entname,
            mobile: mobile
        }
        var submit_btn = $(this).find("button[type='submit']");
        $.ajax({
            type: "POST",
            url: "/userportal/contactus",
            data: d,
            beforeSend:function(){
                LoadBox.show('正在发送信息');
                submit_btn.attr("disabled",'disabled');
            },
            success:function(data){
                LoadBox.hide();
                if(data.code == 0){                
                    $("#md-contact").hide();
                    YdResultBox.show({
                        txt: '发送成功',
                        note: '我们会在七个工作日内联系您'
                    })
                }else{
                    YdResultBox.show('发送失败')
                }
            },
            fail:function(){
                YdResultBox.show('发送失败')
            },
            complete:function(){
                LoadBox.hide();
                submit_btn.removeAttr("disabled");
            }
        })
        return false;
    })
})(jQuery);

//form搜索盒子
(function($){
    $.fn.youduSearchBox = function (options) {
        var defaults = {
            inputId: '',
            placeholder: '',
            autocomplete: '',
            tpl : '<div class="form-group search-group"><input type="text" placeholder="<&placeholder&>" autocomplete="<&autocomplete&>" maxlength="20"><button type="submit" class="icon-search"></button><i class="icon-search-close hide"></i></div>',
            submit_callback: null,
            events: []  // 例：{"type":'input',"interval": 300,"callback":function(){}}
        }; 
        var opts = $.extend(defaults, options); 
        var _that = this;
        var bindEvent = function(){
            _that.submit(function(){
                opts.submit_callback && opts.submit_callback();
                return false;
            })

            if(!!opts.events.length){
                opts.events.forEach(function(v){
                    if(v.type && v.type == "input"){
                        _that.find("input").on(v.type,_.debounce(function(){
                            v.callback && v.callback();
                        }, opts.events.interval || 300))
                    }else if(v.type){
                        _that.find("input").on(v.type,function(){
                            v.callback && v.callback();
                        })
                    }
                })
            }

            _that.find(".icon-search-close").click(function(){
                _that.find("input").val('');
                $(this).addClass("hide");
                if(opts.events.length){
                    opts.events.forEach(function(v){
                        if(v.type == "input"){
                            v.callback();
                            return false
                        }
                    })
                }
            })

            _that.find("input").keyup(function(){
                if($(this).val()){
                    _that.find(".icon-search-close").removeClass("hide")
                }else{
                    _that.find(".icon-search-close").addClass("hide")
                }
            })

        }

        var init = function(){
            var h = opts.tpl.replace(/<&placeholder&>/ig, opts.placeholder).replace(/<&autocomplete&>/ig, opts.autocomplete)
            _that.append(h); 
            opts.inputId && _that.find("input").attr("id",opts.inputId)
            bindEvent();
        }();

        this.get = function(){
            return $.trim(_that.find("input").val());
        }

        this.set = function(val){
            var val = val || ''
            if(!val){
                _that.find(".icon-search-close").addClass("hide")
            }
            return _that.find("input").val(val)
        }

        return this;
        
    }
})(jQuery);

(function ($) { 
    /**
     * youdu select  plugin
     * auth dico.zhang
     * demo :
     *  <select id="ydseldemo" placeholder="请选择">
            <option value="">请选择</option>
            <option value="1" selected>苹果</option>
            <option value="2">士多俾梨士多俾梨</option>
            <option value="3">雪梨</option>
        </select>
     * 
     *  var ydseldemo = $("#ydseldemo").youduSelect({
            size: 'sm',
            hasIcon: 1,
            minWidth: 100
        })
    */
    $.fn.youduSelect = function (options) { 
        var defaults = {
            tpl : '<div class="yd-select <&size&>" id="<&id&>"><&old_tpl&><div class="yd-select-tog"><span class="yd-select-res"><&placeholder&></span><i class="yd-select-tri"></i></div><div class="yd-select-options <&okicon&>" style="opacity: 0"><ul><&options&></ul></div></div>',
            size: '', // sm 
            hasIcon : false,
            minWidth : '',
            relativeBox: '',
            callback: null
        }; 
        var opts = $.extend(defaults, options); 
        var _that = this;
        var _dom_select = null;
        var _dom_opts = null;
        var _item_tpl = '<li class="yd-select-item <&selected&>" data-vas="<&value&>"><&title&></li>';
        var _ok_icon = 'ic';
        var _class_sel = 'selected'
        var _class_dis = 'disabled'
        var select_id = "yd-select-" + new Date().getTime() + (Math.random() * 100000 >> 0);
        var res = null;

        function _set(v, isReset) {
            var isSet = false;
            if (res && res.value == v) {
                return
            }
            _dom_select.find('.yd-select-item').each(function(){
                if (!isSet && $(this).data('vas') == v) {
                    isSet = true
                    $(this).addClass(_class_sel).siblings().removeClass(_class_sel).end().parents('.yd-select-options').first().hide();
                    _dom_select.find('.yd-select-res').text($(this).text())
                    res = {
                        value : v,
                        title : $(this).text()
                    }
                    opts.callback && opts.callback(res);
                }
            })
            if (!isSet) {
                _set('', isReset);
            }
        }

        function _setPosition(){
            var page_offset = $(opts.relativeBox).offset()
            var page_height = $(opts.relativeBox).outerHeight()
            var self_offset = _dom_opts.offset()
            var self_height = _dom_opts.outerHeight()
            if ((self_offset.top + self_height + 20) > (page_offset.top + page_height)) {
                _dom_opts.css('top', '-'+(self_height+5)+'px')
            }
            _dom_opts.hide().css('opacity', 1)
        }

        function _buildSel () {
            var h = _buildOpts()
            var ic = opts.hasIcon ? _ok_icon: '';
            h = opts.tpl.replace(/<&size&>/ig, opts.size).replace(/<&id&>/ig, select_id).replace(/<&placeholder&>/ig, placeholder).replace(/<&okicon&>/ig, ic).replace(/<&options&>/ig, h).replace(/<&old_tpl&>/ig, _that.clone().removeAttr('id').css('display', 'none').prop("outerHTML"))

            return h;
        }

        function _buildOpts () {
            var h = '';
            var options = _that.find("option")
            options.each(function(k, v){
                var sel = $(this).attr(_class_sel) ? _class_sel : '';
                var val = $(this).attr('value');
                var title = $(this).text();
                if (!val) {
                    placeholder = title
                }
                if (sel) {
                    placeholder = title
                    if (!res) {
                        res = {
                            value : val,
                            title : title
                        }
                    }
                }
                h += _item_tpl.replace(/<&selected&>/ig, sel).replace(/<&value&>/ig, val).replace(/<&title&>/ig, title)
            })
            return h;
        }

        var bindEvent = function(){
            _dom_select.find('.yd-select-tog').click(function(evt){
                $('.yd-select-options:visible').hide()
                _dom_opts.toggle();
                evt.stopPropagation()
            })
            _dom_select.on('click', '.yd-select-item:not(.'+_class_dis+')', function(evt){
                _set($(this).data('vas'))          
                evt.stopPropagation()    
            })

            $('body').click(function(){
                _dom_opts.hide();
            })
        }


        var init = function(){
            var placeholder = _that.attr("placeholder");
            var h = _buildSel();
            _that.after(h)
            _that.hide();
            _dom_select = _that.next("#"+select_id);
            _dom_opts = _dom_select.children(".yd-select-options");
            if (opts.minWidth) {
                _dom_select.css('min-width', (opts.minWidth - 36) + 'px')
            }
            opts.relativeBox && $(opts.relativeBox).length && _setPosition()
            bindEvent();
        }();
        
        this.get = function(){
            return res;
        }
        this.set = function(v){
            _set(v)
        }

        this.reInit = function(){
            var h = _buildOpts();
            var r = res ? res.value : '';
            _dom_opts.children('ul').html(h)
            setTimeout(function(){
                res = null;
                _dom_opts.removeAttr("style").css('opacity', 0).show()
                _setPosition();
                _set(r, 1)
            }, 10);
        }

        this.disableOpts = function(v){
            if (!v) { return }
            _dom_select.find('.yd-select-item').each(function(){
                if ($(this).data('vas') == v) {
                    $(this).addClass(_class_dis)
                }
            })
        }
        this.ableAllOpts = function(){
            _dom_select.find('.yd-select-item.'+_class_dis).removeClass(_class_dis)
        }
        this.ableOpts = function(v){
            if (!v) { return }
            _dom_select.find('.yd-select-item').each(function(){
                if ($(this).data('vas') == v) {
                    $(this).removeClass(_class_dis)
                }
            })
        }
        
        this.addOptions = function(items) {
            if (!(items instanceof Array)) {
                console.error("youdu-select-plugin: add options params is not array", items)
                return false
            }
            var h = '';
            _.forEach(items, function(m) {
                h += _item_tpl.replace(/<&selected&>/ig, '').replace(/<&value&>/ig, m.value).replace(/<&title&>/ig, m.title)
            })
            _dom_select.children('.yd-select-options').children('ul').append(h)
        }
        
        this.removeOption = function(value) {
            _dom_select.children('.yd-select-options').find("li[data-vas='"+value+"']").remove()
            setTimeout(_that.reInit, 30)
        }
        
        return this;
    }; 
})(jQuery); 


//时分选择盒子
(function($) {
    $.fn.youduChooseTime = function (options) {
        var defaults = {
            tpl: '<span class="time-hour click"><&hour&></span><span>:</span><span class="time-min click"><&min&></span><ul class="ul-time-box hour"></ul><ul class="ul-time-box min"></ul>',
            timeHour: 24,
            timeMin: 60,
            hour: "09",
            min: "00",
            callback: null
        }; 
        var opts = $.extend(defaults, options); 
        var _that = this;

        function _buildTpl () {
            var h = opts.tpl
            h = h.replace(/<&hour&>/ig, opts.hour).replace(/<&min&>/ig, opts.min)
            return h;
        }

        var hideUlTime = function () {
            $(".ul-time-box").hide();
            $("body").unbind("click",hideUlTime);
        }

        var bindEvent = function () {
            _that.find(".time-hour").click(function(e){
                e.stopPropagation();
                $(".ul-time-box").not(_that.find(".ul-time-box.hour")).hide();
                _that.find(".ul-time-box.hour").toggle();
                $("body").bind("click", hideUlTime);
            })

            _that.find(".time-min").click(function(e){
                e.stopPropagation();
                $(".ul-time-box").not(_that.find(".ul-time-box.min")).hide();
                _that.find(".ul-time-box.min").toggle();
                $("body").bind("click", hideUlTime);
            })

            _that.find(".ul-time-box.hour li").click(function(){
                var val = $(this).attr("val");
                _that.find(".time-hour").html(val);
                _that.hour = val;
            })
            _that.find(".ul-time-box.min li").click(function(){
                var val = $(this).attr("val");
                _that.find(".time-min").html(val);
                _that.min = val;
            })
           
        }
        

        var init = function () {
            var tmpl = _buildTpl();
            _that.addClass("select-time-box").html(tmpl);
            _that.hour = opts.hour;
            _that.min = opts.min;

            //小时
            for (var i = 0; i < opts.timeHour; i++) {
                i = i < 10 ? "0" + i : i
               _that.find(".ul-time-box.hour").append("<li val="+i+">" + i + "</li>")
            }
            //分
            for (var i = 0; i < opts.timeMin; i++) {
                i = i < 10 ? "0" + i : i
                _that.find(".ul-time-box.min").append("<li val="+i+">" + i + "</li>")
            }

            bindEvent();
        }()

        this.set = function(val){
            switch(typeof(val)){
                case "string":{
                    var hour = val.substr(0,2)
                    var min = val.substr(2,2)
                    this.hour = hour
                    this.min = min
                    $(_that[0]).find(".time-hour").text(hour);
                    $(_that[0]).find(".time-min").text(min);
                }break;
                default:break;
            }
        }

        this.setHour = function (val) {
            if(typeof(val) == "number"){
                val = val < 10 ? "0"+val : String(val)
                this.hour = val
                $(_that[0]).find(".time-hour").text(val);
            }
        }

        this.setMin = function (val) {
            if(typeof(val) == "number"){
                val = val < 10 ? "0"+val : String(val)
                this.min = val
                $(_that[0]).find(".time-min").text(val);
            }
        }

        return this;
     };
})(jQuery);


function showContactUs(model) {
    $("#md-contact").find("#contact_model").val(model);
    $("#md-contact").find("#contact_person").val('');
    $("#md-contact").find("#contact_entname").val('');
    $("#md-contact").find("#contact_mobile").val('');
    $("#md-contact").show();
}


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


var Helper = function(){
    this.utf8_to_b64 = function ( str ) {
        return window.btoa(unescape(encodeURIComponent( str )));
    }

    this.b64_to_utf8 = function ( str ) {
        return decodeURIComponent(escape(window.atob( str )));
    }
    return this
}()

function startPageAction(){
    var param = GetRequest()
    if (param.a) {
        $("#"+param.a).addClass('animate')
        setTimeout(function(){
            $("#"+param.a).removeClass('animate')
        }, 3000)
    }
}
startPageAction()
