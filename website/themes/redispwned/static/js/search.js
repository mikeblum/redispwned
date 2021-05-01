document.addEventListener("DOMContentLoaded", function(event) {
    attachFormSubmitEvent("redis-search")
}, {once: true});

function formSubmit(event) {
    event.preventDefault();
    if (document.getElementById("redis-addr").value === "") {
        return;
    }
    let el = document.getElementById("redis-report-done");
    toggleHide(el);
    el = document.getElementById("redis-report-error");
    toggleHide(el);
    el = document.getElementById("redis-report-running");
    toggleShow(el);
    fetch("https://api.redispwned.app/scan/" + document.getElementById("redis-addr").value, {
        method : "POST",
        headers: {
            "Content-Type": 'application/json',
            "X-CSRF": document.getElementById("_csrf").value
        },
    }).then(
        response => response.json()
    ).then(
        json => {
            document.getElementById("redis-report-city").innerText = json.city
            document.getElementById("redis-report-country").innerText = json.country_code
            document.getElementById("redis-report-info").innerText = json.info.info_response
            document.getElementById("redis-report-ping").innerText = json.info.ping_response
            let el = document.getElementById("redis-report-running");
            toggleHide(el);
            el = document.getElementById("redis-report-done");
            toggleShow(el);
        }
    ).catch((error) => {
        let el = document.getElementById("redis-report-running");
        toggleHide(el);
        el = document.getElementById("redis-report-error");
        toggleShow(el);
    });
}

function fetchCsrf() {
    fetch("https://api.redispwned.app/csrf", {
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

function toggleShow(el) {
    el.classList.add('show');
    el.classList.remove('hide');
}

function toggleHide(el) {
    el.classList.add('hide');
    el.classList.remove('show');
}
