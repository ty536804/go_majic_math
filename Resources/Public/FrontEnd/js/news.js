//请求数据
function getAjax(page)
{
    $.ajax({
        type: "GET",
        dataType: "json",
        url: "/newList",
        data: {"page":page},
        success: function (result) {
            let _html= "";
            if (Number(result.code) == 200) {
                $.each(result.data.list,function (k,v) {
                    let timeStr = v.created_at;
                    _html += "<a href='/detail?id="+v.id+"'><dl><dt><img src='"+v.thumb_img+"'></dt>"
                    _html += '<dd><h3>'+v.title+'</h3><p>'+v.summary+'</p></dd></dl>';
                    if (timeStr.substring(0,10) == getFormatDate()) {
                        _html += "<span></span>";
                    }
                    _html += "</a>";
                })
                pageList(page)
            }
            $(".new_ul").empty().append(_html)
        }
    });
}
function getFormatDate(){
    var nowDate = new Date();
    var year = nowDate.getFullYear();
    var month = nowDate.getMonth() + 1 < 10 ? "0" + (nowDate.getMonth() + 1) : nowDate.getMonth() + 1;
    var date = nowDate.getDate() < 10 ? "0" + nowDate.getDate() : nowDate.getDate();
    return year + "-" + month + "-" + date;
}
//分页
pageList(1)
function pageList(page) {
    $('#pageLimit').bootstrapPaginator({
        currentPage: page,//当前页。
        totalPages: $('#PageCount').val(),//总页数。
        size:"normal",//应该是页眉的大小。
        bootstrapMajorVersion: 3,//bootstrap的版本要求。
        alignment:"right",
        numberOfPages:5,//显示的页数
        itemTexts: function (type, page, current) {//如下的代码是将页眉显示的中文显示我们自定义的中文。
            switch (type) {
                case "first": return "首页";
                case "prev": return "上一页";
                case "next": return "下一页";
                case "last": return "末页";
                case "page": return page;
            }
        },
        onPageClicked: function (event, originalEvent, type, page) {//给每个页眉绑定一个事件，其实就是ajax请求，其中page变量为当前点击的页上的数字。
            getAjax(page)
        }
    });
}