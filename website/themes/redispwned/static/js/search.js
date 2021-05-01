document.addEventListener("DOMContentLoaded", function(event) {
    attachFormSubmitEvent("redis-search")
}, {once: true});

function formSubmit(event) {
    console.log(document.getElementById("redis-addr").value)
    fetch("http://localhost:8080/scan/" + document.getElementById("redis-addr").value, {
        method : "POST",
    }).then(
        response => response.json()
    ).then(
        json => console.log(json)
    );
    event.preventDefault();
}


function attachFormSubmitEvent(formId){
    document.getElementById(formId).addEventListener("submit", formSubmit);
}
