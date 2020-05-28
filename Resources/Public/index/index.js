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
                let _html = "";
                let _banner = "";
                let _oli = "";
                let _dl = "";
                if (Number(result.code) == 200) {
                    $.each(result.data.nav,function (k,v) {
                        _html += '<a href="'+v.base_url+'">'+v.name+'</a>';
                    })
                    $.each(result.data.banner,function (k,v) {
                        _banner +='<div class="carousel-item '+(k==0 ? 'active': '')+'" ><img src="/static/upload/'+v.imgurl+'"></div>'
                        _oli += '<li data-target="#myCarousel" data-slide-to="'+k+'" class="active"></li>';
                    })
                    $.each(result.data.list,function (k,v) {
                        _dl += "<dl><dt><img src='/static/upload/"+v.thumb_img+"'></dt><dd><h5>"+v.title+"</h5><p>"+v.summary+"</p></dd></dl>"
                    })
                }
                $(".links").empty().append(_html)
                $(".carousel-inner").empty().append(_banner)
                $('.carousel-indicators').empty().append(_oli)
                $('.dynamic .six_reason').append(_dl);
            }
        });
    }
})