#include "WiFi.h"
#include <WebServer.h>
#include <uri/UriBraces.h>

const char* WIFI_USER     = "Policial";
const char* WIFI_PASSWORD = "07855005889";

const int LOCKER_O_PIN = 26;
const int LOCKER_1_PIN = 27;

bool LOCKER_0_STATUS = LOW;
bool LOCKER_1_STATUS = LOW;

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
  pinMode(LOCKER_O_PIN, OUTPUT);
  pinMode(LOCKER_1_PIN, OUTPUT);
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
  server.on(UriBraces("/lock/{}"), lockHandler);
  server.on(UriBraces("/unlock/{}"), unlockHandler);

  server.begin();
  Serial.println(" - Server: OK");
}

void homeHandler() {
  sendResponse(200, "Welcome to ValuSafe!");
}

void lockHandler() {
  String lockerID = server.pathArg(0);
s
  if (lockerID == "0") {
    digitalWrite(LOCKER_O_PIN, LOW);
  } else if (lockerID == "1") {
    digitalWrite(LOCKER_1_PIN, LOW);
  } else {
    sendResponse(404, "Locker not found!");
  }

  sendResponse(200, "Locker was successfully locked!");
}

void unlockHandler() {
  String lockerID = server.pathArg(0);

  if (lockerID == "0") {
    digitalWrite(LOCKER_O_PIN, HIGH);
  } else if (lockerID == "1") {
    digitalWrite(LOCKER_1_PIN, HIGH);
  } else {
    sendResponse(404, "Locker not found!");
  }

  sendResponse(200, "Locker was successfully unlocked!");
}

void sendResponse(int code, String message) {
  server.send(code, "text/plain", message);
}

void loop() {
  server.handleClient();
  delay(2);
}