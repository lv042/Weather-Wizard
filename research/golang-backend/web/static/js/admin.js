
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
    location.reload();
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
setInterval(fetchData, 1000);