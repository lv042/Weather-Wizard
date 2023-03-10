#System Requirements

## Wemos sending data to the backend

Embedded device sends real sensordata to the application backend over https.



The Wemos sends requests to the backend with the latest sensordata.


Post request to the backend with the following json format:

```json
{
    "temperature": 20.0,
    "humidity": 50.0,
    "pressure": 1000.0,
    "light": 100.0
}
```

## Backend pinging the Wemos

Embedded device receives or retrieves real status messages from the application backend over https.

## 3-Input sensors

The embedded device contains three or more correctly working input sensors:

- Temperature and humidity sensor
  ```
  

  ```

[Image](docs/../images/wifi_manager.png)

- Light sensor
    ```
    
    
    ```

[Image](docs/../images/wifi_manager.png)

- Button
    ```
    
    
    ```

[Image](docs/../images/wifi_manager.png)

## 3-Output sensors

The embedded device contains three or more correctly working visual and/or sensory outputs.

- LED
    ```
    
    
    ```

[Image](docs/../images/wifi_manager.png)
  
- Buzzer
    ```
    
    
    ```

[Image](docs/../images/wifi_manager.png)
  
- Display
    ```
    
    
    ```

[Image](docs/../images/wifi_manager.png)
  
## Wifi-Manager

The embedded device uses wifi manager to authenticate and connect to the network over https.

```
    
    
```

[Image](docs/../images/wifi_manager.png)