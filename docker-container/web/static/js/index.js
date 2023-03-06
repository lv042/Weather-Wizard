// Temperature graph
var trace1 = {
    x: [],
    y: [],
    type: 'scatter'
};
var data1 = [trace1];
var layout1 = {
    title: 'Temperature',
    plot_bgcolor: 'rgba(0,0,0,0)',
    paper_bgcolor: 'rgba(0,0,0,0)'
};
Plotly.newPlot('temperature', data1, layout1);

// Humidity graph
var trace2 = {
    x: [],
    y: [],
    type: 'scatter'
};
var data2 = [trace2];
var layout2 = {
    title: 'Humidity',
    plot_bgcolor: 'rgba(0,0,0,0)',
    paper_bgcolor: 'rgba(0,0,0,0)'
};
Plotly.newPlot('humidity', data2, layout2);

// Light intensity graph
var trace3 = {
    x: [],
    y: [],
    type: 'scatter'
};
var data3 = [trace3];
var layout3 = {
    title: 'Light Intensity',
    plot_bgcolor: 'rgba(0,0,0,0)',
    paper_bgcolor: 'rgba(0,0,0,0)'
};
Plotly.newPlot('light', data3, layout3);

// Wind graph
var trace4 = {
    x: [],
    y: [],
    type: 'scatter'
};
var data4 = [trace4];
var layout4 = {
    title: 'Wind Speed',
    plot_bgcolor: 'rgba(0,0,0,0)',
    paper_bgcolor: 'rgba(0,0,0,0)'
};
Plotly.newPlot('wind', data4, layout4);

function updateTemperatureGraph() {
    // Send an AJAX request to fetch the latest temperature data
    $.ajax({
        url: 'index.php?action=weather_data',
        method: 'GET',
        success: function(data) {
            // Parse the JSON data
            var jsonData = JSON.parse(data);
            console.log(jsonData);

            // Extract the temperature and timestamp data
            var timestamps = [];
            var temperatures = [];
            for (var i = 0; i < jsonData.length; i++) {
                timestamps.push(new Date(jsonData[i].timestamp));
                temperatures.push(jsonData[i].temperature);
            }

            // Update the temperature trace with the new data
            var temperatureTrace = {
                x: timestamps,
                y: temperatures,
                type: 'scatter',
                mode: 'lines',
                line: {
                    color: 'red',
                    width: 2
                }
            };
            var layout = {
                title: 'Temperature (°C)',
                xaxis: {
                    title: 'Time',
                    showgrid: false,
                    tickformat: '%m/%d %H:%M:%S',
                    automargin: true
                },
                yaxis: {
                    title: 'Temperature (°C)',
                    showline: false
                },
                margin: {
                    l: 60,
                    r: 30,
                    t: 30,
                    b: 30
                },
                plot_bgcolor: 'transparent',
                paper_bgcolor: 'transparent'

            };
            Plotly.newPlot('temperature', [temperatureTrace], layout);
        }
    });
}


// Call the updateTemperatureGraph function every 10 seconds
setInterval(updateTemperatureGraph, 3000);


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




//some settings for vanta js which is responsible for the cool background
VANTA.CLOUDS({
    el: "#main",
    mouseControls: true,
    touchControls: true,
    gyroControls: false,
    minHeight: 500.00,
    minWidth: 200.00,
    speed: 1.00
})
