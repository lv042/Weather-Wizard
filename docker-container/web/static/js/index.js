// Temperature graph
var trace1 = {
    x: [1, 2, 3, 4],
    y: [20, 22, 25, 21],
    type: 'scatter'
};
var data1 = [trace1];
var layout1 = {
    title: 'Temperature'
};
Plotly.newPlot('temperature', data1, layout1);

// Humidity graph
var trace2 = {
    x: [1, 2, 3, 4],
    y: [40, 42, 45, 41],
    type: 'scatter'
};
var data2 = [trace2];
var layout2 = {
    title: 'Humidity'
};
Plotly.newPlot('humidity', data2, layout2);

// Light intensity graph
var trace3 = {
    x: [1, 2, 3, 4],
    y: [60, 62, 65, 61],
    type: 'scatter'
};
var data3 = [trace3];
var layout3 = {
    title: 'Light Intensity'
};
Plotly.newPlot('light', data3, layout3);

// Wind graph
var trace4 = {
    x: [1, 2, 3, 4],
    y: [10, 12, 15, 11],
    type: 'scatter'
};
var data4 = [trace4];
var layout4 = {
    title: 'Wind Speed'
};
Plotly.newPlot('wind', data4, layout4);

function updateTemperatureGraph() {
    // Send an AJAX request to fetch the latest temperature data
    $.ajax({
        url: '/get_temperature_data',
        method: 'GET',
        success: function(data) {
            // Update the temperature trace with the new data
            var temperatureTrace = {
                x: data.time,
                y: data.temperature,
                type: 'scatter',
                mode: 'lines',
                line: {
                    color: 'red',
                    width: 2
                }
            };
            Plotly.update('temperature', temperatureTrace);
        }
    });
}

// Call the updateTemperatureGraph function every 10 seconds
setInterval(updateTemperatureGraph, 10000);
