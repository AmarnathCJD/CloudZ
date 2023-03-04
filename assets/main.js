var ui = {
    upload_button: document.getElementById('upload-button'),
    upload_input: document.getElementById('upload-input'),
    upload_progress: document.getElementById('upload-progress'),

    add_download_button: document.getElementById('add-download-button'),
    add_download_input: document.getElementById('add-download-input'),
}

// Path: assets\main.js

function upload_file() {
    var file = ui.upload_input.files[0];
    var reader = new FileReader();

    reader.onload = function (e) {
        var data = e.target.result;
        var xhr = new XMLHttpRequest();
        xhr.open('POST', 'http://127.0.0.1:8080/upload', true);
        xhr.setRequestHeader('Content-Type', 'multipart/form-data');
        // set multipart field file
        xhr.setRequestHeader('Content-Disposition', 'form-data; name="file"; filename="' + file.name + '"');
        // set multipart field file size
        xhr.setRequestHeader('Content-Length', file.size);
        // set multipart field file type
        // start upload
        
        xhr.upload.onprogress = function (e) {
            if (e.lengthComputable) {
                var percent = (e.loaded / e.total) * 100;
                ui.upload_progress.value = percent;
            }
        };
        xhr.send(data);
    };
    reader.readAsArrayBuffer(file);
}

ui.upload_button.addEventListener('click', upload_file);