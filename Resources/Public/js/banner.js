getAjax(1)
//请求数据
function getAjax(page)
{
    $.ajax({
        type: "GET",
        dataType: "json",
        url: "/api/v1/getBanners",
        data: {"page":page},
        success: function (result) {
            let _html= "";
            if (Number(result.code) == 200) {
                $.each(result.data.list,function (k,v) {
                    _html += "<tr><td>"+v.id+"</td><td>"+v.bname+"</td><td>"+v.nav.name+"</td><td>"+v.target_link+"</td><td>"+v.info+"</td>"
                    _html += "<td>"+(Number(v.is_show)==1 ? '启用' : '禁用')+"</td><td class='td-manage'>"
                    _html += '<a title="编辑" href="/api/v1/bannerDetail?id='+v.id+'"><i class="fa fa-pencil-square-o" aria-hidden="true"></i>编辑</a>'
                    _html += '&nbsp;&nbsp;<a title="删除" onclick="delBanner('+v.id+')"><i class="fa fa-times" aria-hidden="true"></i>删除</a>'
                    _html +=  '</td></tr>';
                })
                pageList(page,result.data.count)
            } else {
                _html = "<tr><td colspan='9' class='text-center'>暂无内容</td></tr>"
            }
            $(".tbody").empty().append(_html)
        }
    });
}


function delBanner(id) {
    if (id < 1) {
        layer.msg("非法操作");
        return false
    }
    $.get({
        type:"POST",
        dataType: "json",
        url:"/api/v1/delBanner",
        data:{"id":id},
        success:function (result) {
            if (result.code == 200) {
                layer.msg("操作成功");
                return false
            }
            layer.msg("操作失败");
            return false
        }
    })
}
//分页
function pageList(page,PageCount) {
    $('#pageLimit').bootstrapPaginator({
        currentPage: page,//当前页。
        totalPages: PageCount,//总页数。
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