$(function(){
    $('#content').summernote({
        height: 200,
        tabsize: 2,
        lang: 'zh-CN',
        callbacks:{
            onImageUpload: function(files, editor, $editable) {
                uploadSummerPic(files[0], editor, $editable);
            },
            insertImage:function ($editable,file) {
                console.log(file)
            }
        }
    });

    function uploadSummerPic(file, editor, $editable) {
        let fd = new FormData();
        let fileName = file['name'];
        fd.append("file", file);
        $.ajax({
            type:"POST",
            url:"/api/v1/upload",
            data: fd,
            cache: false,
            contentType: false,
            processData: false,
            success: function (data) {
                let _url = "/static/upload/"+data.data;
                $('#content').summernote('insertImage', _url, fileName);
            }
        });
    }
});