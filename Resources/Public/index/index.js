$(function () {
    getAjax()
//请求数据
    function getAjax()
    {
        $.ajax({
            type: "GET",
            dataType: "json",
            url: "/index",
            success: function (result) {
                let _banner = "";
                let _oli = "";
                let _dl = "";
                if (Number(result.code) == 200) {
                    if (result.data.banner.length > 1) {
                        $.each(result.data.banner, function (k, v) {
                            if (k >0) {
                                _banner += '<div class="carousel-item" ><img src="/static/upload/' + v.imgurl + '"></div>'
                                _oli += '<li data-target="#myCarousel" data-slide-to="' + k + '" class="active"></li>';
                            }
                        })
                        $(".carousel-inner").append(_banner)
                        $('.carousel-indicators').append(_oli)
                    }
                    $.each(result.data.list, function (k, v) {
                        _dl += "<dl><dt><img src='/static/upload/" + v.thumb_img + "'></dt><dd><h5>" + v.title + "</h5><p>" + v.summary + "</p></dd></dl>"
                    })

                    $.each(result.data.magic, function (k, magicVal) {

                        $.each(magicVal, function (k2, v2) {
                            if (k == "数学思维有什么魔力？") {
                                let _magic_list = '<h4>'+v2.summary+'</h4><p>'+v2.content+'</p>';
                                $('.magic_list dl dd:eq('+k2+')').empty().html(_magic_list)
                            }
                            if (k== "张梅玲教授倾力推荐") {
                                $('.recommend h3').empty().html(v2.name)
                                $('.outborder dl dt').empty().html(v2.summary)
                                $('.outborder dl dd').empty().html(v2.content)
                            }
                            if (k== "选择魔法数学的六个理由") {
                                $('.six_reason dl:eq('+k2+')').html("<dt><p class='celebrity_desc'>"+v2.content+"</p></dt><dd>"+v2.summary+"</dd>");
                            }
                            if (k== "魔法数学与传统数学的区别") {
                                let tradition_ul = v2.content.split(',')
                                let _tradition_ul = '<li class="difference_con_tit mb0">'+v2.summary+'</li>';
                                for (let i =0 ;i< tradition_ul.length;i++) {
                                    _tradition_ul+= '<li '+(i ==(tradition_ul.length-1) ? 'class="mb0"' : '')+'>'+tradition_ul[i]+'</li>'
                                }
                                if (k2 <1) {
                                    $('.tradition_ul').empty().append(_tradition_ul)
                                } else {
                                    $('.magic_ul').empty().append(_tradition_ul)
                                }
                            }
                            if (k== "Hi-Thinking教学法") {
                                let _thinking_ul = '<h4>'+v2.summary+'</h4><p>'+v2.content+'</p>';
                                $('.thinking_ul dl dd:eq('+k2+')').empty().html(_thinking_ul)
                            }
                            if (k== "特色课堂，让孩子流连忘返") {
                                let tradition_ul = v2.content.split(',')
                                for (let i =0 ;i< tradition_ul.length;i++) {
                                    if (k2 <1) {
                                        $('.different_ul_top dl dt:eq('+i+')').empty().html(tradition_ul[i])
                                    } else {
                                        $('.different_ul_bottom dl dt:eq('+i+')').empty().html(tradition_ul[i])
                                    }

                                }
                            }
                        });

                    });
                }
                $('.dynamic .six_reason').append(_dl);
            }
        });
    }
})