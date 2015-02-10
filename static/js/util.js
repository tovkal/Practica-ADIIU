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
	$("#" + element).css('background-image', 'url("http://staticadiiu.tovkal.com/img/uploads/' + image + '")');
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
	return $("<option></option>").attr("value", id).text(name);
}

// Common ajax calls
function read(operation, id, success_fn, error_fn) {
	$.ajax({
		type: "GET",
		datatype: "json",
		url: "/api/" + operation + (id != null ? "/" + id : ""),
		success: success_fn,
		error: error_fn == null ? function(jqXHR, textStatus, errorThrown) {
			alert("Error deleting Categoria. Got:\n" + errorThrown); // DEBUG
		} : error_fn
	});
};

function create(operation, element, success_fn, error_fn) {
	$.ajax({
		type: "POST",
		datatype: "json",
		url: "/api/" + operation,
		data: JSON.stringify(element.serializeObject()),
		success: success_fn,
		error: error_fn == null ? function(jqXHR, textStatus, errorThrown) {
			alert("Error creating Categoria. Got:\n" + errorThrown); // DEBUG
		} : error_fn
	});
};

function update(operation, id, element, success_fn, error_fn) {
	$.ajax({
		type: "PUT",
		datatype: "json",
		url: "/api/" + operation + "/" + id,
		data: JSON.stringify(element.serializeObject()),
		success: success_fn,
		error: error_fn == null ? function(jqXHR, textStatus, errorThrown) {
			alert("Error deleting Categoria. Got:\n" + errorThrown); // DEBUG
		} : error_fn
	});
};

function deleteFn(operation, id, success_fn, error_fn) {
	$.ajax({
		type: "DELETE",
		datatype: "json",
		url: "/api/" + operation + "/" + id,
		success: success_fn,
		error: error_fn == null ? function(jqXHR, textStatus, errorThrown) {
			alert("Error deleting Categoria. Got:\n" + errorThrown); // DEBUG
		} : error_fn
	});
};

Date.prototype.formattedDate = function() {
	var yyyy = this.getFullYear().toString();
	var mm = (this.getMonth()+1).toString(); // getMonth() is zero-based
	var dd  = this.getDate().toString();
return yyyy + "-" + (mm[1]?mm:"0"+mm[0]) + "-" + (dd[1]?dd:"0"+dd[0]); // padding
};

function getJSONLength(json) {
	var key, count = 0;
	for(key in json) {
		if(json.hasOwnProperty(key)) {
			count++;
		}
	}
	return count;
}

function createCookie(name, value, days) {
	if (days) {
		var date = new Date();
		date.setTime(date.getTime() + (days * 24 * 60 * 60 * 1000));
		var expires = "; expires=" + date.toGMTString();
	} else {
		var expires = "";
	}		

	document.cookie = name + "=" + value + expires + "; path=/";
}

function readCookie(name) {
	var nameEQ = name + "=";
	var ca = document.cookie.split(';');
	for (var i = 0; i < ca.length; i++) {
		var c = ca[i];
		while (c.charAt(0) == ' ') { 
			c = c.substring(1, c.length);
		}

		if (c.indexOf(nameEQ) == 0) { 
			return c.substring(nameEQ.length, c.length); 
		}
	}
	return null;
}

function eraseCookie(name) {
	createCookie(name, "", -1);
}

// Custom jQuery functions
(function($) {
    $.fn.hideBootstrap = function() {
		this.addClass("hidden");
    };
    $.fn.showBootstrap = function() {
        this.removeClass("hidden");
    };
    $.fn.isEmpty = function() {
        return this.val().length == 0;
    };
    $.fn.isVisible = function() {
        return !this.hasClass('hidden');
    };
})(jQuery);

$(document).ready(function() {
	applyNivel();
});

function applyNivel() {
	if (readCookie("nivel") == 255) {
		$(".loggedon").showBootstrap();
		$(".responsable-magatzem").showBootstrap();
		$(".responsable-farmacia").hideBootstrap();
	} else if (readCookie("nivel") == 0) {
		$(".loggedon").showBootstrap();
		$(".responsable-magatzem").hideBootstrap();
		$(".responsable-farmacia").showBootstrap();
	} else {
		$(".responsable-magatzem").hideBootstrap();
		$(".responsable-farmacia").hideBootstrap();
		$(".loggedon").hideBootstrap();
	}
}
