#include <WiFi.h>
#include <HTTPClient.h>

const char* ssid = "YourWiFiSSID";
const char* password = "YourWiFiPassword";

const char* serverName = "http://192.168.0.100/HC05Predict/index.php";

// HC05 Ultrasonic Sensor Pins
const int trigPin = 2;
const int echoPin = 3;

void setup() {
  Serial.begin(115200);
  pinMode(trigPin, OUTPUT);
  pinMode(echoPin, INPUT);
  
  connectToWiFi();
}

void loop() {
  long duration, distance;
  
  // Triggering the Ultrasonic Sensor
  digitalWrite(trigPin, LOW);
  delayMicroseconds(2);
  digitalWrite(trigPin, HIGH);
  delayMicroseconds(10);
  digitalWrite(trigPin, LOW);
  
  // Receiving the Echo signal
  duration = pulseIn(echoPin, HIGH);
  
  // Calculating distance
  distance = duration * 0.034 / 2;
  
  // Print distance to Serial Monitor
  Serial.print("Distance: ");
  Serial.println(distance);
  
  // Sending data to server
  sendData(distance);
  
  delay(5000); // Adjust delay as needed
}

void connectToWiFi() {
  Serial.print("Connecting to WiFi");
  WiFi.begin(ssid, password);
  while (WiFi.status() != WL_CONNECTED) {
    delay(500);
    Serial.print(".");
  }
  Serial.println("WiFi connected");
}

void sendData(long distance) {
  HTTPClient http;
  
  // Constructing the URL with distance as a parameter
  String url = serverName + "?distance=" + String(distance);
  
  Serial.print("Sending data to server: ");
  Serial.println(url);
  
  // Sending GET request
  http.begin(url);
  int httpResponseCode = http.GET();
  
  // Checking for response
  if (httpResponseCode > 0) {
    Serial.print("HTTP Response code: ");
    Serial.println(httpResponseCode);
    String response = http.getString();
    Serial.println(response);
  } else {
    Serial.print("Error code: ");
    Serial.println(httpResponseCode);
  }
  
  http.end();
}
