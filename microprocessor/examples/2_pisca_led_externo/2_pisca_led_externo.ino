#define EXTERNAL_LED_1 26
#define EXTERNAL_LED_2 27

void setup() {
  pinMode(EXTERNAL_LED_1, OUTPUT);
  pinMode(EXTERNAL_LED_2, OUTPUT);
}

void loop() {
  digitalWrite(EXTERNAL_LED_1, HIGH);
  digitalWrite(EXTERNAL_LED_2, HIGH);
  delay(500);
  digitalWrite(EXTERNAL_LED_1, LOW);
  digitalWrite(EXTERNAL_LED_2, LOW);
  delay(500);
}
