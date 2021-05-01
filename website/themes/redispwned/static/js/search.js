document.addEventListener("DOMContentLoaded", function(event) {
    attachFormSubmitEvent("redis-search")
}, {once: true});

function formSubmit(event) {
    event.preventDefault();
    if (document.getElementById("redis-addr").value === "") {
        return;
    }
    fetch("http://localhost:8080/scan/" + document.getElementById("redis-addr").value, {
        method : "POST",
        headers: {
            "Content-Type": 'application/json',
            "X-CSRF": document.getElementById("_csrf").value
        },
    }).then(
        response => response.json()
    ).then(
        json => console.log(json)
    );
}

function fetchCsrf() {
    fetch("http://localhost:8080/csrf", {
        method : "GET",
    }).then(
        response => response.json()
    ).then(
        json => document.getElementById("_csrf").value = json.token
    );
}

function attachFormSubmitEvent(formId){
    // fetch a CSRF token
    fetchCsrf();
    document.getElementById(formId).addEventListener("submit", formSubmit);
}
