/**
 * Created by tovkal on 4/12/14.
 */

var dropZone = document.getElementById('drop-zone');
var progress = document.getElementById('photoProgressBar');

function uploadOnDrop(e) {
    e.preventDefault();
    this.className = 'upload-drop-zone';

    startUpload(e.dataTransfer.files)
};

function highlightUploadDropZone() {
    this.className = 'upload-drop-zone drop';
    return false;
};

function unhighlightUploadDropZone() {
    this.className = 'upload-drop-zone';
    return false;
};

function setupUpload() {
    dropZone.ondrop = uploadOnDrop;
    dropZone.ondragover = highlightUploadDropZone;
    dropZone.ondragleave = unhighlightUploadDropZone;
};

function disableUpload() {
    dropZone.ondrop = null;
    dropZone.ondragover = null;
    dropZone.ondragleave = null;
};

$(document).ready(function() {
    'use strict';

    // UPLOAD CLASS DEFINITION
    // ======================
    function startUpload(files) {

        if (files.length == 0) {
            return;
        }

        var data = new FormData();
        data.append('SelectedFile', files[0]);

        var request = new XMLHttpRequest();
        request.onreadystatechange = function(){
            if(request.readyState == 4){
                var resp;
                try {
                    resp = JSON.parse(request.response);
                } catch (e){
                    resp = {
                        status: 'error',
                        data: 'Unknown error occurred: [' + request.responseText + ']'
                    };
                }

                if (resp.status == 'success') {
                    setBackgroundImage('drop-zone', files[0].name);
                } else {
                    alert("Error pujant foto:\n" + resp['data']); // DEBUG
                }
            }
        };

        request.upload.addEventListener('progress', function(e){
            progress.style.width = Math.ceil(e.loaded/e.total) * 100 + '%';
        }, false);

        request.open('POST', '/api/upload');
        request.send(data);
    };

    $("#js-upload-submit").click(function() {
        var uploadFiles = document.getElementById('js-upload-files').files;

        startUpload(uploadFiles)
    });

    setupUpload();
});
