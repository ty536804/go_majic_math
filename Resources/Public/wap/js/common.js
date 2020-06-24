$(function() {
    $("img.lazy").lazyload({
        effect : "fadeIn"
    });
});


$('.home_nav').css({"height":$(document).height()+"px"});
$('.home_icon').on('click',function () {
    let flag = $(".home_nav").is(":hidden");
    if(flag){
        $('.navicon').attr("src","/static/wap/img/clone.png");
        $(".home_nav").show();
    }else{
        $('.navicon').attr("src","/static/wap/img/home_icon_breadnav_nor@2x.png");
        $(".home_nav").hide();
    }
})

let _currentUrl = 'http://'+window.location.host
let currentUrl = window.location.href.substr(_currentUrl.length,window.location.href.length)

$('.home_nav a').each(function () {
    let _actClass = '/'+$(this).attr('class');
    if (_actClass == currentUrl) {
        $(this).addClass('nav_active').siblings().removeClass('nav_active')
    } else if (currentUrl == '/') {
        $('.home_nav a:eq(0)').addClass('nav_active').siblings().removeClass('nav_active')
    }
})

$('.c-com').val(window.location.href);

$('.f_btn').on('click',function () {
    if ($('.footer_con_right .c-area').val()=="") {
        layer.tips('姓名不能为空', '.footer_con_right .c-area', {
            tips: [1, '#3595CC'],
            time: 4000
        });
        return false;
    }
    if ($('.footer_con_right .c-tel').val()=="") {
        layer.tips('电话不能为空', '.footer_con_right .c-tel', {
            tips: [1, '#3595CC'],
            time: 4000
        });
        return false;
    }
    if ($('.footer_con_right .c-city').val()=="") {
        layer.tips('地区不能为空', '.footer_con_right .c-city', {
            tips: [1, '#3595CC'],
            time: 4000
        });
        return false;
    }
    $.ajax({
        type: "POST",
        dataType: "json",
        url: "/AddMessage",
        data:$('#myform').serialize(),
        success: function (result) {
            if (result.code == 200) {
                $('input').val('');
                layer.alert("留言成功");
                return false
            }
            layer.alert("留言失败");
            return false
        }
    })
    return false;
})
