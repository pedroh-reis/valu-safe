#define INTERNAL_LED 2

void setup() {
  pinMode(INTERNAL_LED, OUTPUT);
}

void loop() {
  digitalWrite(INTERNAL_LED, HIGH);
  delay(500);
  digitalWrite(INTERNAL_LED, LOW);
  delay(500);
}
