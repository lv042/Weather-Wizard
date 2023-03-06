
//Update the temperature graph in the beginning
updateTemperatureGraph();

var ly_margin = {
    l: 70,
    r: 30,
    t: 30,
    b: 30
}

// Temperature graph
var trace1 = {
    x: [],
    y: [],
    type: 'scatter'
};
var temp_data = [trace1];
var temp_layout = {
    title: 'Temperature',
    xaxis: {
        showgrid: false,
        tickformat: '%m/%d %H:%M:%S',
        automargin: true,
        gridcolor: 'white'
    },
    yaxis: {
        title: 'Temperature (°C)',
        showline: false,
        gridcolor: 'white'
    },
    margin: ly_margin,
    plot_bgcolor: 'transparent',
    paper_bgcolor: 'transparent'
};
Plotly.newPlot('temperature', temp_data, temp_layout);

// Humidity graph
var hum_trace = {
    x: [],
    y: [],
    type: 'scatter'
};
var hum_data = [hum_trace];
var hum_layout = {
    title: 'Humidity',
    xaxis: {
        showgrid: false,
        tickformat: '%m/%d %H:%M:%S',
        automargin: true,
        gridcolor: 'white'
    },
    yaxis: {
        title: 'Humidity (%)',
        showline: false,
        gridcolor: 'white'
    },
    margin: ly_margin,
    plot_bgcolor: 'transparent',
    paper_bgcolor: 'transparent'
};
Plotly.newPlot('humidity', hum_data, hum_layout);

// Light intensity graph
var light_trace = {
    x: [],
    y: [],
    type: 'scatter'
};
var light_data = [light_trace];
var light_layout = {
    title: 'Light Intensity',
    xaxis: {
        showgrid: false,
        tickformat: '%m/%d %H:%M:%S',
        automargin: true
        ,
    },
    yaxis: {
        title: 'Light Intensity (lux)',
        showline: false,
        gridcolor: 'white'
    },
    margin: ly_margin,
    plot_bgcolor: 'transparent',
    paper_bgcolor: 'transparent'
};
Plotly.newPlot('light', light_data, light_layout);

// pressure graph
var pres_trace = {
    x: [],
    y: [],
    type: 'scatter',
};
var pres_data = [pres_trace];
var pres_layout = {
    title: 'Air Pressure',
    xaxis: {
        showgrid: false,
        tickformat: '%m/%d %H:%M:%S',
        automargin: true,
        gridcolor: 'white'

    },
    yaxis: {
        title: 'Air Pressure (hPa)',
        showline: false,
        gridcolor: 'white',
        //changes the step size of the y-axis
        dtick: 1
    },
    margin: ly_margin,
    plot_bgcolor: 'transparent',
    paper_bgcolor: 'transparent'
};

Plotly.newPlot('pressure', pres_data, pres_layout);

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
            var humidities = [];
            var light = [];
            var pressure = [];
            for (var i = 0; i < jsonData.length; i++) {

                //Adds the data to the arrays
                timestamps.push(new Date(jsonData[i].timestamp));
                temperatures.push(jsonData[i].temperature);
                humidities.push(jsonData[i].humidity);
                light.push(jsonData[i].light_intensity);
                pressure.push(jsonData[i].pressure);
            }
            //Update temperature graph
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

            Plotly.newPlot('temperature', [temperatureTrace], temp_layout);

            //Update humidity graph
            // Update the humidity trace with the new data
            var humidityTrace = {
                x: timestamps,
                y: humidities,
                type: 'scatter',
                mode: 'lines',
                line: {
                    color: 'red',
                    width: 2
                }
            };

            Plotly.newPlot('humidity', [humidityTrace], hum_layout);

            //Update light graph
            // Update the light trace with the new data
            var lightTrace = {
                x: timestamps,
                y: light,
                type: 'scatter',
                mode: 'lines',
                line: {
                    color: 'red',
                    width: 2
                }
            };

            Plotly.newPlot('light', [lightTrace], light_layout);

            //Update pressure graph
            // Update the pressure trace with the new data
            var pressureTrace = {
                x: timestamps,
                y: pressure,
                type: 'scatter',
                mode: 'lines',
                line: {
                    color: 'red',
                    width: 2
                }
            };

            Plotly.newPlot('pressure', [pressureTrace], pres_layout);

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
