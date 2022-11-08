#include "WiFi.h"
#include <WebServer.h>
#include <ESP32Servo.h>
#include <uri/UriBraces.h>

// const char* WIFI_USER     = "Morada_Passaros Lateral 2.4";
// const char* WIFI_PASSWORD = "maia1234";

// const char* WIFI_USER     = "LAB-DIDATICO";
// const char* WIFI_PASSWORD = "C1-13#lami#2017";

const char* WIFI_USER     = "LAB_DIGITAL";
const char* WIFI_PASSWORD = "C1-17*2018@labdig";

struct Locker {
  const int pin;
  Servo servo;
};

Servo servo0;
Servo servo1;

Locker locker0 = Locker{27, servo0};
Locker locker1 = Locker{26, servo1};

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
  locker0.servo.attach(locker0.pin);
  locker1.servo.attach(locker1.pin);

  locker0.servo.write(0);
  locker1.servo.write(0);

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
  String id = server.pathArg(0);

  if (id == "0") {
    Serial.printf("Serial 0 angle: %d", locker0.servo.read());
    if (locker0.servo.read() == 0) {
      locker0.servo.write(120);
      return sendResponse(200, "Locker unlocked with success.");      
    } else if (locker0.servo.read() == 119) {
      locker0.servo.write(0);
      return sendResponse(200, "Locker locked with success.");
    } else {
      return sendResponse(425, "Locker 0 is changing its state. Try again!");
    }
  } else if (id == "1") {
    Serial.printf("Serial 1 angle: %d", locker1.servo.read());
    if (locker1.servo.read() == 0) {
      locker1.servo.write(120);
      return sendResponse(200, "Locker unlocked with success.");      
    } else if (locker1.servo.read() == 119) {
      locker1.servo.write(0);
      return sendResponse(200, "Locker locked with success.");
    } else {
      return sendResponse(425, "Locker 0 is changing its state. Try again!");
    }
  }

  return sendResponse(404, "Locker not found.");




  // int i = getLockerIndexById(id);

  // if (i == -1) {
  //   return sendResponse(404, "Locker not found.");
  // }
  // Serial.println("chegou aqui");
  // Serial.printf("Servo %s angle: %d\n", i, lockers[i].servo.read());
  // if (lockers[i].servo.read() == 0) {
  //   lockers[i].servo.write(120);
  //   return sendResponse(200, "Locker unlocked with success.");    
  // } else if (lockers[i].servo.read() == 120) {
  //   lockers[i].servo.write(0);
  //   return sendResponse(200, "Locker locked with success.");
  // } else {
  //   return sendResponse(425, "Locker 0 is changing its state. Try again!");    
  // }
}

void sendResponse(int code, String message) {
  String jsonMessage = String("{\"message\": \"" + message + "\"}");

  server.send(code, "application/json", jsonMessage);
}

void loop() {
  server.handleClient();
  delay(2);
}
