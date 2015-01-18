function enablePhotoUpload() {
	$(".enabledUpload").removeClass('hidden');
	setupUploadDropZone();
}

function disablePhotoUpload() {
	$(".enabledUpload").addClass('hidden');
	disableUploadDropZone();
}

function setBackgroundImage(element, image) {
	$("#" + element).css('background-image', 'url("img/uploads/' + image + '")');
	$("#" + element).addClass('background-image-style');
}

function removeBackgroundImage(element) {
	$("#" + element).css('background-image', '');
	$("#" + element).removeClass('background-image-style');
}

function resetForm($form) {
    $form.find('input:text, input:password, input:file, select, textarea').val('');
    $form.find('input:radio, input:checkbox')
         .removeAttr('checked').removeAttr('selected');
}

function changeButton(element, action, text, functionName, id, row) {
	var btnClass = '';
	var clickAction = '';
	switch(action){
		case 'read':
			btnClass = 'btn btn-default';
			clickAction = 'closeModal()';
			break;
		case 'create':
			btnClass = 'btn btn-success';
			clickAction = functionName + '()';
			break;
		case 'update':
			btnClass = 'btn btn-primary';
			clickAction = functionName + '(' + id + ',' + row + ')';
			break;
		case 'delete':
			break;
	}

	element.attr('class', btnClass).text(text).attr('onclick', clickAction);
}

$.fn.serializeObject = function()
{
    var o = {};
    var a = this.serializeArray();
    $.each(a, function() {
        if (o[this.name] !== undefined) {
            if (!o[this.name].push) {
                o[this.name] = [o[this.name]];
            }
            o[this.name].push(this.value || '');
        } else {
            o[this.name] = this.value || '';
        }
    });
    return o;
};