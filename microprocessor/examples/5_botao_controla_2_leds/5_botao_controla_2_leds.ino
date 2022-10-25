const int LED_0 = 26;
const int LED_1 = 27;

int LED_0_ACTUAL_STATE = LOW;
int LED_1_ACTUAL_STATE = LOW;

const int BUTTON_ON_OFF = 19;
const int BUTTON_SELECT = 18;

void setup() {
  pinMode(LED_0, OUTPUT);
  pinMode(LED_1, OUTPUT);
  
  pinMode(BUTTON_ON_OFF, INPUT);
  pinMode(BUTTON_SELECT, INPUT);
}

void loop() {
  if (digitalRead(BUTTON_SELECT) == LOW) {
    if (digitalRead(BUTTON_ON_OFF) == HIGH) {
      LED_0_ACTUAL_STATE = (LED_0_ACTUAL_STATE + 1)%2;
      digitalWrite(LED_0, LED_0_ACTUAL_STATE);
      delay(200);
    }
  } else {
    if (digitalRead(BUTTON_ON_OFF) == HIGH) {
      LED_1_ACTUAL_STATE = (LED_1_ACTUAL_STATE + 1)%2;
      digitalWrite(LED_1, LED_1_ACTUAL_STATE);
      delay(200);
    }
  }
}


