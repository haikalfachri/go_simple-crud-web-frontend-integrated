// home.html
function openAddForm() {
    window.location.href = '/form';
}

function getAllUserData() {
    fetch('http://localhost:8000/biodata', {
        method: 'GET',
        headers: {
            'Content-Type': 'application/json'
        },
    })
        .then(function (response) {
            return response.json();
        })
        .then(function (data) {
            displayUserData(data.data);
        })
        .catch(function (error) {
            console.error('Failed:', error);
        });
}

function displayUserData(userData) {
    var userDataDiv = document.getElementById('userData');
    userDataDiv.innerHTML = '';

    if (userData.length === 0) {
        var emptyMessage = document.createElement('label');
        emptyMessage.textContent = 'No biodata available.';
        userDataDiv.appendChild(emptyMessage);
    } else {
        userData.forEach(function (user) {
            var userItem = document.createElement('div');
            userItem.classList.add('user-item');
            var dob = user.dob.slice(0, 10);
            var gender = user.gender.charAt(0).toUpperCase() + user.gender.slice(1);
            userItem.innerHTML = `
                <div class="entry">
                    <img class="profile-pic" src="${user.url}" alt="Profile Photo">
                    <div class="data">
                        <h3>${user.name}</h3>
                        <p>Date of Birth: ${dob}</p>
                        <p>Phone        : ${user.phone}</p>
                        <p>Address      : ${user.address}</p>
                        <p>Gender       : ${gender}</p>
                    </div>
                </div>
                <div class="actions">
                        <img class="button" src="icon/edit.png" onclick="editUser(${user.ID})" alt="Edit">
                        <img class="button" src="icon/delete.png" onclick="deleteUser(${user.ID})" alt="Delete">
                </div>
                <hr class="divider">
                `;
            userDataDiv.appendChild(userItem);
        });
    }
}

function deleteUser(id) {
    fetch(`http://localhost:8000/biodata/${id}`, {
        method: 'DELETE',
        headers: {
            'Content-Type': 'application/json'
        },
    })
        .then(function (response) {
            return response.json();
        })
        .then(function () {
            console.log('user deleted');
            getAllUserData();
        })
        .catch(function (error) {
            console.error('Failed:', error);
        });
}

function editUser(id) {
    window.location.href = `/edit?${id}`;
}

function fetchUser() {
    var url = window.location.search;
    var userid = url.slice(1)

    fetch(`http://localhost:8000/biodata/${userid}`, {
        method: 'GET',
        headers: {
            'Content-Type': 'application/json'
        },
    })
        .then(function (response) {
            return response.json();
        })
        .then(function (data) {
            console.log('Get a user:', data.data);
            fillEditForm(data.data);
        })
        .catch(function (error) {
            console.error('Failed:', error);
        });
}

function fillEditForm(userData) {
    var userDataDiv = document.getElementById('userData');
    userDataDiv.innerHTML = '';

    var formElement = document.createElement('form');
    formElement.setAttribute('action', '#');
    formElement.onsubmit = function() { return false; };

    createInput(formElement, 'Name', 'text', 'Enter your name', userData.name, true);
    createInput(formElement, 'Date of Birth', 'date', 'Enter your date of birth', userData.dob.slice(0, 10), true);
    createInput(formElement, 'Phone Number', 'number', 'Enter your phone number', userData.phone, true);
    createInput(formElement, 'Address', 'address', 'Enter your address', userData.address, true);

    var profileInputLabel = document.createElement('label');
    profileInputLabel.textContent = 'Profile Picture';
    formElement.appendChild(profileInputLabel);

    var profileInputBox = document.createElement('div');
    profileInputBox.classList.add('input-box');

    var profileInput = document.createElement('input');
    profileInput.setAttribute('type', 'file');
    profileInput.setAttribute('accept', 'image/*');
    profileInput.setAttribute('id', 'profileImage');
    profileInput.required = true;
    profileInputBox.appendChild(profileInput);

    formElement.appendChild(profileInputBox);

    var genderLabel = document.createElement('label');
    genderLabel.textContent = 'Gender';
    formElement.appendChild(genderLabel);

    var genderDiv = document.createElement('div');
    genderDiv.classList.add('gender');

    if (userData.gender === "female"){
        createRadio(genderDiv, 'Male', 'gender', 'male', false);
        createRadio(genderDiv, 'Female', 'gender', 'female', true);
    } else{
        createRadio(genderDiv, 'Male', 'gender', 'male', true);
        createRadio(genderDiv, 'Female', 'gender', 'female', false);
    }
    
    formElement.appendChild(genderDiv);

    var submitButtonBox = document.createElement('div');
    submitButtonBox.classList.add('input-box', 'button');

    var submitButton = document.createElement('input');
    submitButton.setAttribute('type', 'submit');
    submitButton.value = 'Save';
    submitButton.onclick = editBiodata;

    submitButtonBox.appendChild(submitButton);
    formElement.appendChild(submitButtonBox);

    userDataDiv.appendChild(formElement);
}

function createInput(formElement, label, type, placeholder, value, required) {
    var inputLabel = document.createElement('label');
    inputLabel.textContent = label;
    formElement.appendChild(inputLabel);

    var inputBox = document.createElement('div');
    inputBox.classList.add('input-box');

    var input = document.createElement('input');
    input.setAttribute('type', type);
    input.setAttribute('placeholder', placeholder);
    input.required = required;
    input.value = value;

    inputBox.appendChild(input);
    formElement.appendChild(inputBox);
}

function createRadio(parentElement, label, name, value, checked) {
    var radioInput = document.createElement('input');
    radioInput.setAttribute('type', 'radio');
    radioInput.setAttribute('name', name);
    if (value === 'female'){
        radioInput.setAttribute('style', 'margin-left: 10px')
    }
    radioInput.value = value;
    radioInput.required = true;
    radioInput.checked = checked;

    parentElement.appendChild(radioInput);
    parentElement.appendChild(document.createTextNode(label));
}

// add.html
function closeForm() {
    window.location.href = '/home';
}

function addBiodata(event) {
    event.preventDefault(); // Prevent form submission

    var form = document.querySelector('form');
    var name = form.querySelector('input[type="text"]').value;
    var dob = form.querySelector('input[type="date"]').value + "T00:00:00.000+00:00";
    var phone = form.querySelector('input[type="number"]').value;
    var address = form.querySelector('input[type="address"]').value;
    var gender = form.querySelector('input[name="gender"]:checked').value;
    var profileImage = form.querySelector('#profileImage').files[0];

    if (name.length === 0) {
        alert('Please fill in all required fields.');
        return;
    }

    var formData = new FormData();
    formData.append('name', name);
    formData.append('dob', dob);
    formData.append('phone', phone);
    formData.append('address', address);
    formData.append('gender', gender);
    formData.append('image', profileImage);

    fetch('http://localhost:8000/biodata', {
        method: 'POST',
        body: formData
    })
        .then(function (response) {
            return response.json();
        })
        .then(function (data) {
            console.log('Success:', data);
        })
        .catch(function (error) {
            console.error('Failed:', error);
        });

    window.location.href = '/home';
}

function editBiodata(event) {
    event.preventDefault(); // Prevent form submission

    var form = document.querySelector('form');
    var name = form.querySelector('input[type="text"]').value;
    var dob = form.querySelector('input[type="date"]').value + "T00:00:00.000+00:00";
    var phone = form.querySelector('input[type="number"]').value;
    var address = form.querySelector('input[type="address"]').value;
    var gender = form.querySelector('input[name="gender"]:checked').value;
    var profileImage = form.querySelector('#profileImage').files[0];

    var formData = new FormData();
    formData.append('name', name);
    formData.append('dob', dob);
    formData.append('phone', phone);
    formData.append('address', address);
    formData.append('gender', gender);
    formData.append('image', profileImage);

    var url = window.location.search;
    var userid = url.slice(1)

    fetch(`http://localhost:8000/biodata/${userid}`, {
        method: 'PUT',
        body: formData
    })
        .then(function (response) {
            return response.json();
        })
        .then(function (data) {
            console.log('Success:', data);
        })
        .catch(function (error) {
            console.error('Failed:', error);
        });

    window.location.href = '/home';
    getAllUserData();
}