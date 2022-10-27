#define INTERNAL_LED 2
#define EXTERNAL_LED 13

void setup() {
  pinMode(INTERNAL_LED, OUTPUT);
  pinMode(EXTERNAL_LED, OUTPUT);
}

void loop() {
  digitalWrite(INTERNAL_LED, HIGH);
  delay(500);
  digitalWrite(INTERNAL_LED, LOW);
  delay(500);

  digitalWrite(EXTERNAL_LED, HIGH);
  delay(500);
  digitalWrite(EXTERNAL_LED, LOW);
  delay(500);
}
