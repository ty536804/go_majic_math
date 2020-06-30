$(function(){
    //导航
    $('.lanmu').click(function(){
        $('.nav').stop(true,true).slideToggle();
    })

    $('.nav').css('height',$(window).height()-138);

    $('.ntitle1').click(function(){
        $(this).addClass('xz').siblings().removeClass('xz');
        $(this).parents().siblings('.ntitle_nr').children('.ntitle_nr1').eq($(this).index()).show().siblings().hide();
    })

    //弹窗1
    $('.zx1').height($(window).height());
    $('.tcdiv1').css('margin-top',($(window).height()-$('.tcdiv1').height())/2);
    /*$('.tanchuang1').click(function(){
        $('.zx1').fadeIn();
        $('.tcdiv1').fadeIn();

    })
    $('.zx1,.gb1').click(function(){
        $('.zx1').hide();
        $('.tcdiv1').hide();
    })*/


    //魔力思维 我们的服务
    $('.qmlsw2_1 ul li:odd').addClass('odd');

    /*-----------------------------L. start---------------------------------*/

    $('.about-box-02 li').click(function(){
        var oTxt = $(this).children('.box-txt');
        if (oTxt.is(":hidden")) {
            oTxt.slideDown();
            $(this).siblings('li').find('.box-txt').slideUp();
        } else {
            oTxt.slideUp();
        }
    })


    /*-----------------------------L. end---------------------------------*/

    /*-----------------------------WY. start---------------------------------*/
    /* 弹窗关闭 */ 
    $('.y-close').click(function(){
        $('.y-mc').fadeOut();
        $('.y-public-tc').slideUp();
    });
    /* 登录 */ 
    $('.y-login-up').click(function(){
        $('.nav').fadeOut();
        $('.y-public-tc').hide();
        $('.y-mc').fadeIn();
        $('.y-login-tc').slideDown();
    });
    /* 注册 */ 
    $('.y-design-up').click(function(){
        $('.nav').fadeOut();
        $('.y-public-tc').hide();
        $('.y-mc').fadeIn();
        $('.y-design-tc').slideDown();
    });
    /* 忘记密码 */ 
    $('.y-forget-up').click(function(){
        $('.nav').fadeOut();
        $('.y-public-tc').hide();
        $('.y-mc').fadeIn();
        $('.y-forget-tc').slideDown();
    });
    /* 下载 */ 
    $('.y-download-up').click(function(){
        $('.nav').fadeOut();
        $('.y-public-tc').hide();
        $('.y-mc').fadeIn();
        $('.y-download-tc').slideDown();
    })
    /*-----------------------------WY. end---------------------------------*/

})