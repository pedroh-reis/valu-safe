#include "WiFi.h"
#include <WebServer.h>
#include <ESP32Servo.h>
#include <uri/UriBraces.h>

// const char* WIFI_USER     = "Morada_Passaros Lateral 2.4";
// const char* WIFI_PASSWORD = "maia1234";

const char* WIFI_USER     = "LAB-DIDATICO";
const char* WIFI_PASSWORD = "C1-13#lami#2017";

struct ServoMotor {
  Servo servo;
  const int pin;
}

struct ServoMotor servo0 = {Servo servo, 18};
struct ServoMotor servo1 = {Servo servo, 19};

WebServer server(80);

void setup() {
  setupSerialConnection();
  setupPinModes();
  setupWifiConnection();
  setupWebServer();
}

void setupSerialConnection() {
  Serial.begin(115200);
  Serial.println("### SETUP ###");
  Serial.println(" - Serial connection: OK");
}

void setupPinModes() {
  servo0.servo.attach(servo0.pin);
  servo1.servo.attach(servo1.pin);
  Serial.println(" - Pin modes: OK");
}

void setupWifiConnection() {
  Serial.println("   ### WIFI ###");
  Serial.printf("    - Trying to connect to %s ...", WIFI_USER);
  WiFi.begin(WIFI_USER, WIFI_PASSWORD);
  while (WiFi.status() != WL_CONNECTED) {
    delay(1000);
    Serial.print(".");
  }
  Serial.println();
  Serial.printf("    - Connected to %s\n", WIFI_USER);
  Serial.print("    - IP address: ");
  Serial.println(WiFi.localIP());
}

void setupWebServer() {
  server.on("/", homeHandler);
  server.on(UriBraces("/locker/{}"), changeState);

  server.begin();
  Serial.println(" - Server: OK");
}

void homeHandler() {
  sendResponse(200, "Welcome to ValuSafe Microcontroller Web Server!");
}

void changeState() {
  String lockerID = server.pathArg(0);

  if (lockerID == "0") {
    if (servo0.servo.read() == 0) {
      servo0.servo.write(120);
      return sendResponse(200, "Locker 0 unlocked with success.");
    } else if (servo0.servo.read() == 120) {
      servo0.servo.write(0);
      return sendResponse(200, "Locker 0 locked with success.");
    } else {
      sendResponse(425, "Locker 0 is changing its state. Try again!");
    }
  } else if (lockerID == "1") {
    if (servo1.servo.read() == 0) {
      servo1.servo.write(120);
      return sendResponse(200, "Locker 1 unlocked with success.");
    } else if (servo1.servo.read() == 120) {
      servo1.servo.write(0);
      return sendResponse(200, "Locker 1 locked with success.");
    } else {
      sendResponse(425, "Locker 1 is changing its state. Try again!");
    }
  }

  sendResponse(404, "Locker not found.");
}

void sendResponse(int code, String message) {
  String jsonMessage = String("{\"message\": \"" + message + "\"}");

  server.send(code, "application/json", jsonMessage);
}

void loop() {
  server.handleClient();
  delay(2);
}