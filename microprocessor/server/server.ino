#include "WiFi.h"
#include <WebServer.h>
#include <uri/UriBraces.h>

const char* WIFI_USER     = "Morada_Passaros Lateral 2.4";
const char* WIFI_PASSWORD = "maia1234";

struct Locker {
  const int pin;
  bool state;
};

struct Button {
  const int pin;
};

struct Locker locker0 = {27, LOW};
struct Locker locker1 = {26, LOW};

struct Button button0 = {18};

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
  pinMode(locker0.pin, OUTPUT);
  pinMode(locker1.pin, OUTPUT);
  pinMode(button0.pin, INPUT);
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
    if (locker0.state == LOW && digitalRead(button0.pin) == HIGH) {
      digitalWrite(locker0.pin, HIGH);
      locker0.state = HIGH;
      return sendResponse(200, "Locker 0 unlocked with success.");
    } else if (locker0.state == HIGH && digitalRead(button0.pin) == HIGH) {
      digitalWrite(locker0.pin, LOW);
      locker0.state = LOW;
      return sendResponse(200, "Locker 0 locked with success.");
    } else {
      sendResponse(406, "Close locker 0's door before locking it.");
    }
  } else if (lockerID == "1") {
    if (locker1.state == LOW && digitalRead(button0.pin) == HIGH) {
      digitalWrite(locker1.pin, HIGH);
      locker1.state = HIGH;
      return sendResponse(200, "Locker 1 unlocked with success.");
    } else if (locker1.state == HIGH && digitalRead(button0.pin) == HIGH) {
      digitalWrite(locker1.pin, LOW);
      locker1.state = LOW;
      return sendResponse(200, "Locker 1 locked with success.");
    } else {
      sendResponse(406, "Close locker 1's door before locking it.");
    }
  }

  sendResponse(404, "Locker not found.");
}

void sendResponse(int code, String message) {
  String jsonMessage = String("{\"message\": \"" + message + "\"}");
  Serial.println(jsonMessage);

  server.send(code, "application/json", jsonMessage);
}

void loop() {
  server.handleClient();
  delay(2);
}