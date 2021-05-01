document.addEventListener("DOMContentLoaded", function(event) {
    serversByCountry();
    serversByVersion();
}, {once: true});

function serversByCountry() {
    var request = new XMLHttpRequest();
    request.open('GET', 'https://api.redispwned.app/v1/servers-by-country', true);

    request.onload = function() {
        if (this.status >= 200 && this.status < 400) {
            var jsonData = JSON.parse(this.response);
            const data = {
                labels: jsonData.labels,
                datasets: [{
                    label: 'Exposed Redis Servers By Country',
                    // Redis Labs red
                    backgroundColor: '#cd5d57',
                    borderColor: 'rgb(255, 99, 132)',
                    data: jsonData.data,
                }]
            };
            const config = {
                type: 'bar',
                data,
                // horizontal-bar-chart
                options: {
                    indexAxis: 'y',
                }
            };
            if (window.serversByCountryChart) {
                window.serversByCountryChart.destroy();
            }

            window.serversByCountryChart = new Chart(
                document.getElementById('servers-by-country'),
                config
            );
        } else {
            console.log("Failed to render report")
        }
    };

    request.onerror = reportError

    request.send();
}

function serversByVersion() {
    var request = new XMLHttpRequest();
    request.open('GET', 'https://api.redispwned.app/v1/servers-by-version', true);

    request.onload = function() {
        if (this.status >= 200 && this.status < 400) {
            var jsonData = JSON.parse(this.response);
            const data = {
                labels: jsonData.labels,
                datasets: [{
                    label: 'Exposed Redis Servers By Version',
                    // Redis Labs red
                    backgroundColor: '#cd5d57',
                    borderColor: 'rgb(255, 99, 132)',
                    data: jsonData.data,
                }]
            };
            const config = {
                type: 'bar',
                data,
                // horizontal-bar-chart
                options: {
                    indexAxis: 'y',
                }
            };
            if (window.serversByVersionChart) {
                window.serversByVersionChart.destroy();
            }
            window.serversByVersion = new Chart(
                document.getElementById('servers-by-version'),
                config
            );
        } else {
            console.log("Failed to render report")
        }
    };

    request.onerror = reportError

    request.send();
}

function reportError() {
    console.log("Failed to fetch report")
}
