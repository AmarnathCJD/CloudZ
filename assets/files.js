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
                <div class="flex flex-col">
                    <img src="${file.icon}"
                        class="w-10 h-10" />
                </div>
                <div class="flex flex-col ml-2">
                    <div class="flex flex-row">
                        <div class="flex flex-col">
                            <p class="text-sm font-medium text-gray-900">${file.filename}</p>
                            <p class="text-sm text-gray-500">${file.filesize}</p>
                        </div>
                        <div class="flex flex-col ml-2">
                            <p class="text-sm text-gray-500">${file.fileext}</p>
                            <p class="text-sm text-gray-500">${file.datemodified}</p>
                        </div>
                    </div>
                    <div class="flex flex-row">
                        <div class="flex flex-col">
                            <div class="w-full bg-gray-200 rounded-full h-2.5 dark:bg-gray-700">
                                <div class="bg-blue-600 h-2.5 rounded-full" style="width: 45%">
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
                <div class="flex flex-row mt-2 ml-auto">
                    <div class="flex flex-col">
                        <button type="button"
                            class="text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center inline-flex items-center mr-2 dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800">
                            <svg aria-hidden="true" class="w-5 h-5 mr-2 -ml-1"
                                fill="currentColor" viewBox="0 0 20 20"
                                xmlns="http://www.w3.org/2000/svg">
                                <path
                                    d="M3 1a1 1 0 000 2h1.22l.305 1.222a.997.997 0 00.01.042l1.358 5.43-.893.892C3.74 11.846 4.632 14 6.414 14H15a1 1 0 000-2H6.414l1-1H14a1 1 0 00.894-.553l3-6A1 1 0 0017 3H6.28l-.31-1.243A1 1 0 005 1H3zM16 16.5a1.5 1.5 0 11-3 0 1.5 1.5 0 013 0zM6.5 18a1.5 1.5 0 100-3 1.5 1.5 0 000 3z">
                                </path>
                            </svg>
                            Download
                        </button>
                    </div>
                    <div class="flex flex-col ml-2">
                        <button type="button"
                            class="text-white bg-red-700 hover:bg-red-800 focus:ring-4 focus:outline-none focus:ring-red-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center inline-flex items-center mr-2 dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800">
                            <svg aria-hidden="true" class="w-5 h-5 mr-2 -ml-1"
                                fill="currentColor" viewBox="0 0 20 20"
                                xmlns="http://www.w3.org/2000/svg">
                                <path
                                    d="M3 1a1 1 0 000 2h1.22l.305 1.222a.997.997 0 00.01.042l1.358 5.43-.893.892C3.74 11.846 4.632 14 6.414 14H15a1 1 0 000-2H6.414l1-1H14a1 1 0 00.894-.553l3-6A1 1 0 0017 3H6.28l-.31-1.243A1 1 0 005 1H3zM16 16.5a1.5 1.5 0 11-3 0 1.5 1.5 0 013 0zM6.5 18a1.5 1.5 0 100-3 1.5 1.5 0 000 3z">
                                </path>
                            </svg>
                            Delete
                        </button>
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