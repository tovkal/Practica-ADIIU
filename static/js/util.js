// Enable photo upload in a modal
function enablePhotoUpload() {
	$(".enabledUpload").removeClass('hidden');
	setupUploadDropZone();
}

// Disable photo upload in a modal
function disablePhotoUpload() {
	$(".enabledUpload").addClass('hidden');
	disableUploadDropZone();
}

// Set the image as the background of the element
function setBackgroundImage(element, image) {
	$("#" + element).css('background-image', 'url("img/uploads/' + image + '")');
	$("#" + element).addClass('background-image-style');
}

// Remove a background image from the element
function removeBackgroundImage(element) {
	$("#" + element).css('background-image', '');
	$("#" + element).removeClass('background-image-style');
}

// Reset form fields
function resetForm($form) {
    $form.find('input:text, input:password, input:file, select, textarea, :input[type=number]').val('');
    $form.find('input:radio, input:checkbox')
         .removeAttr('checked').removeAttr('selected');
}

// Setup the modal's action button for a specific operation
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
	}

	element.attr('class', btnClass).text(text).attr('onclick', clickAction);
}

// jQuery function for form serialization
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

// Create a option element given an id to be set as value and a name. The label will be "id - name"
function createOption(id, name) {
	return $("<option></option>").attr("value", id).text(id + " - " + name);
}

// Common ajax calls
function read(operation, id, success_fn) {
	$.ajax({
		type: "GET",
		datatype: "json",
		url: "/api/" + operation + (id != null ? "/" + id : ""),
		success: success_fn,
		error: function(jqXHR, textStatus, errorThrown) {
			alert("Error reading Categoria. Got:\n" + errorThrown); // DEBUG
		}
	});
};

function create(operation, success_fn) {
	$.ajax({
		type: "POST",
		datatype: "json",
		url: "/api/" + operation,
		data: JSON.stringify($form.serializeObject()),
		success: success_fn,
		error: function(jqXHR, textStatus, errorThrown) {
			alert("Error creating Categoria. Got:\n" + errorThrown); // DEBUG
		}
	});
};

function update(operation, id, success_fn) {
	$.ajax({
		type: "PUT",
		datatype: "json",
		url: "/api/" + operation + "/" + id,
		data: JSON.stringify($form.serializeObject()),
		success: success_fn,
		error: function(jqXHR, textStatus, errorThrown) {
			alert("Error creating Categoria. Got:\n" + errorThrown); // DEBUG
		}
	});
};

function deleteFn(operation, id, success_fn) {
	$.ajax({
		type: "DELETE",
		datatype: "json",
		url: "/api/" + operation + "/" + id,
		success: success_fn,
		error: function(jqXHR, textStatus, errorThrown) {
			alert("Error deleting Categoria. Got:\n" + errorThrown); // DEBUG
		}
	});
};