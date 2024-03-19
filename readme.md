## Home Automation with HC-05 Ultrasonic Sensor

This repository contains code for implementing a home automation system using an HC-05 ultrasonic sensor, written in Go, Arduino (for ESP32), and PHP. The system utilizes the HC-05 ultrasonic sensor to detect distance, which is then sent to a PHP server via Wi-Fi. The PHP server receives the data and stores it in a CSV file for further processing.

### Components:

1. **Go Code (GoLang)**:
   - The Go code performs machine learning tasks on the collected ultrasonic sensor data stored in a CSV file. It uses the Gorgonia library to define, train, and evaluate a neural network model for classification or regression tasks.

2. **Arduino Code (C++)**:
   - The Arduino code, written in C++, is designed to run on an ESP32 microcontroller. It reads distance data from the HC-05 ultrasonic sensor and sends it to the PHP server via Wi-Fi using HTTP requests.

3. **PHP Code**:
   - The PHP code receives the incoming distance data from the ESP32 and stores it in a CSV file named `data.csv`.

### Usage:

1. **Setting up the Hardware**:
   - Connect the HC-05 ultrasonic sensor to the ESP32 microcontroller according to the wiring diagram.
   - Ensure the ESP32 is connected to a Wi-Fi network.

2. **Running the Code**:
   - Upload the Arduino code to the ESP32 microcontroller.
   - Run the Go code on your local machine to perform machine learning tasks on the collected data.
   - Ensure the PHP server is running to receive incoming data from the ESP32.

3. **Data Analysis**:
   - Analyze the stored data in the `data.csv` file to gain insights into occupancy patterns, movement detection, or any other relevant metrics.

### Wiring Diagram:

```
HC-05 Ultrasonic Sensor      ESP32 Microcontroller
    VCC ------------------------ 3.3V
    GND ------------------------ GND
    Trig ------------------------ Digital Pin 2
    Echo ------------------------ Digital Pin 3
```

### Dependencies:

- GoLang (for Go code)
- Arduino IDE (for ESP32 code)
- PHP server with support for HTTP requests

### Notes:

- Adjust the Wi-Fi SSID and password in the ESP32 code (`ssid` and `password` variables) to match your network credentials.
- Ensure the correct URL for the PHP server is specified in the ESP32 code (`serverName` variable).
- Modify the delay in the Arduino code (`delay(5000)`) as needed based on the desired data collection frequency.
- Additional features and functionalities can be implemented based on specific home automation requirements.



## Benefits

- **Custom Home Automation**: The system enables custom home automation solutions tailored to specific requirements, such as occupancy detection, energy management, or security monitoring.
  
- **Real-time Data Analysis**: Data collected from the ultrasonic sensor can be analyzed in real-time to make informed decisions and automate actions based on occupancy or distance measurements.
  
- **Scalability**: The modular architecture allows for easy scalability and integration with additional sensors or actuators to expand the functionality of the home automation system.

## Dependencies

- [Gorgonia](https://pkg.go.dev/gorgonia.org/gorgonia): Library for machine learning and neural network implementation in Go.
- [Arduino IDE](https://www.arduino.cc/en/software): Integrated development environment for programming Arduino and ESP32 microcontrollers.
- [WiFi.h](https://www.arduino.cc/en/Reference/WiFi): Arduino library for connecting to Wi-Fi networks.
- [HTTPClient.h](https://github.com/espressif/arduino-esp32/tree/master/libraries/HTTPClient): Arduino library for making HTTP requests.

## Note

Ensure that the server hosting the `index.php` file is running and accessible from the ESP32 device. Adjust the delay in the Arduino code as needed to control the frequency of data transmission based on your application requirements.


## License
This project is licensed under the [MIT License](LICENSE). You are free to modify and distribute the code for personal or commercial use.
