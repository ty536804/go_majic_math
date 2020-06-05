$('.com').val(window.location.href);
$('.lj_btn').on('click',function () {
    if ($('.banner_form .c-area').val()=="") {
        layer.tips('姓名不能为空!', '.banner_form .c-area', {
            tips: [1, '#3595CC'],
            time: 4000
        });
        return false;
    }
    if ($('banner_form .c-tel').val()=="") {
        layer.tips('电话不能为空', '.banner_form .c-tel', {
            tips: [1, '#3595CC'],
            time: 4000
        });
        return false;
    }
    if ($('.banner_form .c-city').val()=="") {
        layer.tips('地区不能为空┖', '.banner_form .c-city', {
            tips: [1, '#3595CC'],
            time: 4000
        });
        return false;
    }

    $.ajax({
        type: "POST",
        dataType: "json",
        url: "/AddMessage",
        data:$('.banner_myform').serialize(),
        success: function (result) {
            if (result.code == 200) {
                layer.alert("留言成功");
                return false
            }
            layer.alert("留言失败");
            return false
        }
    })
    return false;
})

getAjax()
//请求数据
function getAjax()
{
    $.ajax({
        type: "GET",
        dataType: "json",
        url: "/joinData",
        success: function (result) {
            let _banner = "";
            let _oli = "";
            if (Number(result.code) == 200) {
                $.each(result.data.banner,function (k,v) {
                        _banner +='<div class="carousel-item '+(k==0 ? 'active': '')+'" ><img src="/static/upload/'+v.imgurl+'"></div>'
                        _oli += '<li data-target="#myCarousel" data-slide-to="'+k+'" class="active"></li>';
                })
            }
            $('.carousel-inner').empty().append(_banner);
            $('.carousel-indicators').empty().append(_oli);
            $('.area_con').append('<img src="/static/upload/'+result.data.app[0].imgurl+'" />')
            $('.app_img').empty().append('<img src="/static/upload/'+result.data.learn[0].imgurl+'" />')
            $('.mid').empty().append('<img src="/static/upload/'+result.data.mid[0].imgurl+'" />')
        }
    });
}