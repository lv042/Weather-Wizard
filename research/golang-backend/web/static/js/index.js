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

var font_color = font_color;

//Update the temperature graph in the beginning
updateTemperatureGraph();

var ly_margin = {
    l: 65,
    r: 30,
    t: 50,
    b: 50
}

// Temperature graph
var temp_trace = {
    x: [],
    y: [],
    type: 'scatter'
};
var temp_data = [temp_trace];
var temp_layout = {
    title: 'Temperature',
    xaxis: {
        showgrid: false,
        tickformat: '%Y-%m-%d %H:%M:%S',
        automargin: true,
        gridcolor: font_color
    },
    yaxis: {
        title: 'Temperature (°C)',
        showline: false,
        gridcolor: font_color,
        showgrid: false,
        dtick: 10000,
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
        tickformat: '%Y-%m-%d %H:%M:%S',
        automargin: true,
        gridcolor: font_color
    },
    yaxis: {
        title: 'Humidity (%)',
        showline: false,
        gridcolor: font_color,
        showgrid: false,
        dtick: 10000,
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
        tickformat: '%Y-%m-%d %H:%M:%S',
        automargin: true,
        showline: false
    },
    yaxis: {
        title: 'Light Intensity (lux)',
        showline: false,
        gridcolor: font_color,
        showgrid: false,
        dtick: 10000,
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
        showline: false,
        tickformat: '%Y-%m-%d %H:%M:%S',
        automargin: true,
        gridcolor: font_color,
        color : font_color

    },
    yaxis: {
        title: 'Air Pressure (hPa)',
        showline: false,
        gridcolor: font_color,
        //changes the step size of the y-axis
        dtick: 10000,
        color: font_color,
        showgrid: false,
    },
    margin: ly_margin,
    plot_bgcolor: 'transparent',
    paper_bgcolor: 'transparent',
    color: font_color,
    colorway: [font_color]
};





Plotly.newPlot('pressure', pres_data, pres_layout);

function updateTemperatureGraph() {
    // Send an AJAX request to fetch the latest temperature data
    $.ajax({
        url: 'api/weather',
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
                timestamps.push(jsonData[i].Timestamp);
                temperatures.push(jsonData[i].Temperature);
                humidities.push(jsonData[i].Humidity);
                light.push(jsonData[i].LightIntensity);
                pressure.push(jsonData[i].Pressure);
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
            console.log(temperatureTrace);


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

//mails


// Call the updateTemperatureGraph function every 10 seconds
setInterval(updateTemperatureGraph, 10000);


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
