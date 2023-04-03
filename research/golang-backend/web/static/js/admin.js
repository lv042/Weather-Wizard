
//some settings for vanta js which is responsible for the cool background
VANTA.CLOUDS({
    el: "#main",
    mouseControls: true,
    touchControls: true,
    gyroControls: false,
    minHeight: 1200.00,
    minWidth: 200.00,
    speed: 1.00
});

// Get the h1 element by its ID
var header = document.getElementById('name');

// Add a click event listener to the h1 element
header.addEventListener('click', function() {
    // Reload the page when the h1 element is clicked
    location.href = '/';
});

window.addEventListener('resize', function() {
    location.reload();
});
// Define the API endpoint
const apiUrl = '/api/metrics';

// Define a function to fetch the data asynchronously
async function fetchData() {
    const response = await fetch(apiUrl);
    const data = await response.json();
    console.log(data);

    // Define empty arrays for the x and y data
    var request_x = [];
    var request_y = [];
    var error_x = [];
    var error_y = [];

    // Loop through the requestCount and errorCount objects in the data
    for (var key in data.requestCount) {
        request_x.push(key); // Add the key to the x data for the request count chart
        request_y.push(data.requestCount[key]); // Add the value to the y data for the request count chart
    }
    for (var key in data.errorCount) {
        error_x.push(key); // Add the key to the x data for the error count chart
        error_y.push(data.errorCount[key]); // Add the value to the y data for the error count chart
    }

    // Define the data and layout objects for the request count chart
    var request_data = [
        {
            x: request_x,
            y: request_y,
            type: 'bar',
            marker: {
                color: 'darkred',
                width: 0.1
            },
            width: 0.1
        }
    ];
    var request_layout = {
        plot_bgcolor: 'transparent',
        paper_bgcolor: 'transparent',
        title: 'Request Count',
        xaxis: {
            showgrid: false,
            automargin: true,
            gridcolor: 'white'
        },
        yaxis: {
            showline: false,
            gridcolor: 'white'
        }
    };

    // Define the data and layout objects for the error count chart
    var error_data = [
        {
            x: error_x,
            y: error_y,
            type: 'bar',
            marker: {
                color: 'darkred',
                width: 0.1
            },
            width: 0.1
        }
    ];
    var error_layout = {
        plot_bgcolor: 'transparent',
        paper_bgcolor: 'transparent',
        title: 'Error Count',
        xaxis: {
            showgrid: false,
            automargin: true,
            gridcolor: 'white'
        },
        yaxis: {
            showline: false,
            gridcolor: 'white'
        }
    };

    // Create the request count chart
    Plotly.newPlot('RequestCount', request_data, request_layout);

    // Create the error count chart
    Plotly.newPlot('ErrorCount', error_data, error_layout);
}
fetchData().then(r => console.log(r));

// Call the fetchData function to fetch the data in the background
setInterval(fetchData, 10000);

// Get the submit button element
var submitButton = document.getElementById('submit');

// Add an event listener to the submit button
submitButton.addEventListener('click', function(event) {
    event.preventDefault(); // Prevent the default form submission
    console.log('Button clicked');
    // Get the toggle switch element
    var toggle = document.getElementById('toggle');

    // Get the email input element
    var emailInput = document.getElementById('email');

    // If the toggle switch is checked
    if (toggle.checked) {
        // Get the value of the email input
        var email = emailInput.value;

        // If the email is not empty
        if (email.trim() !== '') {
            // Send a POST request to the backend
            fetch('/api/notifications', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify({
                    email: email,
                    enabled: true
                })
            })
                .then(function(response) {
                    if (response.ok) {
                        console.log('Email registered successfully');
                    } else {
                        console.log('Error registering email');
                    }
                })
                .catch(function(error) {
                    console.log('Error registering email:', error);
                });
        }
    } else {
        // Send a POST request to the backend to disable email notifications
        fetch('/api/notifications', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                enabled: false
            })
        })
            .then(function(response) {
                if (response.ok) {
                    console.log('Email notifications disabled successfully');
                } else {
                    console.log('Error disabling email notifications');
                }
            })
            .catch(function(error) {
                console.log('Error disabling email notifications:', error);
            });
    }
});
