var ui = {
    upload_button: document.getElementById('upload-button'),
    upload_input: document.getElementById('upload-input'),
    upload_progress: document.getElementById('upload-progress'),
    progress_bar : document.getElementById('progress-bar'),

    add_download_button: document.getElementById('add-download-button'),
    add_download_input: document.getElementById('add-download-input'),
}

// Path: assets\main.js

function upload_file() {
    var file = ui.upload_input.files[0];
    var reader = new FileReader();

    reader.onload = function (e) {
        // upload file to server as multipart/form-data
        var xhr = new XMLHttpRequest();
        xhr.open('POST', '/upload', true);
        ui.upload_progress.hidden = false;
        xhr.upload.onprogress = function (e) {
            if (e.lengthComputable) {
                var percent = (e.loaded / e.total) * 100;
                ui.progress_bar.style.width = percent + '%';
            }
        }
        xhr.onload = function () {
            if (this.status == 200) {
                ui.upload_progress.hidden = true;
                ui.progress_bar.style.width = '0%';
                // upload complete
                refresh_file_manager();
            }
        };
        var formData = new FormData();
        formData.append('file', file);
        formData.append('name', file.name);
        xhr.send(formData);
    };
    reader.readAsArrayBuffer(file);
}

ui.upload_button.addEventListener('click', upload_file);