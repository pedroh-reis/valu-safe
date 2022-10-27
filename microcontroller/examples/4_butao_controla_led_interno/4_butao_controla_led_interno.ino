const int INTERNAL_LED = 2;
const int BUTTON = 15;

void setup() {
  pinMode(INTERNAL_LED, OUTPUT);
  pinMode(BUTTON, INPUT);
}

void loop() {
  digitalWrite(INTERNAL_LED, digitalRead(BUTTON));
}
