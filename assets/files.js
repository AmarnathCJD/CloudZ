var files = {
    manager: document.getElementById('files-manager'),
}

function init_file_manager() {
    $.ajax({
        url: '/list',
        type: 'GET',
        dataType: 'json',
        success: function (data) {
            html = ''
            if (data.length == 0) {
                html += `<p class="text-sm text-gray-500">No files found</p>`
            } 
            for (let i = 0; i < data.length; i++) {
                var file = data[i]
                html += `<div class="flex flex-row">
                <div class="flex flex-col ml-2">
                    <div class="flex flex-row">
                        <div class="flex flex-col">
                            <p class="text-sm font-medium text-gray-900"><img src="${file.icon}"
                            class="w-10 h-10 rounded-full object-cover" /></p>
                        </div>
                        <div class="flex flex-col ml-2">
                            <a href='/file/${file.filename}'><p class="text-sm font-medium text-blue-900 break-all">${file.filename}</p></a>
                            <p class="text-sm text-gray-500">${file.filesize}</p>
                        </div>
                    </div>  
                </div>
            </div>`
            }
            $('#files-manager').html(html)
        }
    })
}

function refresh_file_manager() {
    init_file_manager()
}

init_file_manager()