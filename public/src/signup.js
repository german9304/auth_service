// get user attributes
function getUserAttributeValues(e) {
    const userAttributes = ['name', 'age', 'password', 'email'];
    return userAttributes.reduce((acc, curr) => {
        const formValue = e.target.querySelector(`input[name="${curr}"]`);
        acc[curr] = curr === 'age'? +formValue.value : formValue.value;
        return acc;
    }, {});
}

/**
 * Creates a user by making an http POST request
 * 
 * @param {object} user with fields name, email, password and age
 * @returns Promise with status code and response data
 */
async function createUser(user) {
    const data = await fetch('/create-user', {
        method: 'POST',
        headers: {
            'Content-type': 'application/json',
        },
        body: JSON.stringify(user),
    });
    const response = await data.json();

    return {
        status: data.status,
        response,
    }
}

/**
 * Handles form submit
 * 
 * @param {HTMLElement} e form element
 */
async function handleSubmit(e) {
    e.preventDefault();
    const userAttributes = getUserAttributeValues(e);
    try {
        const result = await createUser(userAttributes);
        if(result.status === 201) {
            const userCreate = document.querySelector('.user-create');
            const afterSignupElement = document.querySelector('.after-signup');
            userCreate.classList.add('hide');
            afterSignupElement.classList.remove('hide');
        } else {

        }
    } catch (err){
        console.log(err);
    }
}

// if form element is present
const form = document.getElementById('auth-create-user');
if(form) {
    form.addEventListener('submit', handleSubmit);
}