window.onload = function () {
    document.addEventListener("drop", function (e) {  //拖离
        e.preventDefault();
    })
    document.addEventListener("dragleave", function (e) {  //拖后放
        e.preventDefault();
    })
    document.addEventListener("dragenter", function (e) {  //拖进
        e.preventDefault();
    })
    document.addEventListener("dragover", function (e) {  //拖来拖去
        e.preventDefault();
    })
    var mainBody = document.getElementById("mainBody")
    mainBody.ondragover = function (ev) {
        ev.preventDefault()
    }
    mainBody.ondrop = function (ev) {
        ev.preventDefault();
        var files = ev.dataTransfer.files;
        var fd = new FileReader()
        if (files[0].type.indexOf('image') !== -1) {
            var form = new FormData()
            form.append("image", files[0])
            var xhr = new XMLHttpRequest()
            xhr.open("POST", "/upload")
            //xhr.setRequestHeader("Content-Type", "multipart/form-data;boundary=##AaB03x")
            xhr.onload = function (evt) {
                alert(evt.target.responseText)
            }
            xhr.send(form)
        } else {
            alert("请选择图片上传");
        }
    }
}