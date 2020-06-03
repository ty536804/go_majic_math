$(function () {
    $('.c-com').val(window.location.href);
    $.ajax({
        type: "GET",
        dataType: "json",
        url: "/getNavList",
        success: function (result) {
            let _html= "";
            if (Number(result.code) == 200) {
                let _currentUrl = 'http://'+window.location.host
                let currentUrl = window.location.href.substr(_currentUrl.length,window.location.href.length)
                $.each(result.data,function (k,v) {
                    if (currentUrl == v.base_url) {
                        if (v.base_url == '/') {
                            _html += '<a href="'+v.base_url+'">'+v.name+'</a>';
                        } else {
                            _html += '<a href="'+v.base_url+'">'+v.name+'<p class="nav_line"><span class="nav_line"></span></p></a>';
                        }
                    } else {
                        _html += '<a href="'+v.base_url+'">'+v.name+'</a>';
                    }
                })
            }
            $(".links").empty().append(_html)
        }
    });

    $('.f_btn').on('click',function () {
        if ($('.footer_con_right .c-area').val()=="") {
            layer.tips('姓名不能为空', '.footer_con_right .c-area', {
                tips: [1, '#3595CC'],
                time: 4000
            });
            // layer.alert("姓名不能为空");
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
                    layer.alert("留言成功");
                    return false
                }
                layer.alert("留言失败");
                return false
            }
        })
        return false;
    })
})