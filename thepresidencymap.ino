// This #include statement was automatically added by the Particle IDE.
#include "neopixel/neopixel.h"

#include "application.h"

SYSTEM_MODE(AUTOMATIC)

#define PIXEL_PIN D2
#define PIXEL_COUNT 51
#define PIXEL_TYPE WS2812B

enum State { 
  NM,
  CO,
  Wy,
  MT,
  UT,
  NV,
  ID,
  WA,
  OR,
  CA,
  AZ0,
  AK1,
  HI2,
  TX3,
  LA4,
  MS5,
  AR6,
  OK7,
  MO8,
  KS9,
  NE0,
  SD1,
  ND2,
  MN3,
  WI4,
  MI5,
  IA6,
  IL7,
  IN8,
  KY9,
  TN0,
  AL1,
  FL2,
  GA3,
  SC4,
  NC5,
  OH6,
  WV7,
  DC8,
  VA9,
  PA0,
  MD1,
  DE2,
  NJ3,
  NY4,
  MA5,
  CT6,
  RH7,
  NH8,
  VT9,
  ME};


Adafruit_NeoPixel strip = Adafruit_NeoPixel(PIXEL_COUNT, PIXEL_PIN, PIXEL_TYPE);

// Prototypes for local build. 
void rainbow(uint8_t wait);
uint32_t Wheel(byte WheelPos);
int toggleLED (String str);

bool runDemo;
char StateColor[PIXEL_COUNT];
String cs;

void setup() {
    strip.begin();
    strip.show();
    cs.reserve(51);
    
    runDemo = false;
    
    for(int i = 0; i < PIXEL_COUNT; i++){
        StateColor[i] = ' ';
    }
    Particle.variable("colorString", cs);
    Particle.publish("Started");
    Particle.function("demo", demo);
    Particle.function("color", colorString);
}


void loop() {
    //colorAll(333333, 15);
    if(runDemo){
        for(int i = 1; i < PIXEL_COUNT; i++){
            strip.setPixelColor(i-1, 0x5f,0 , 0);
            delay(200);
            strip.setPixelColor(i, 0x5f,0 ,0 );
            delay(200);
        }
        Particle.publish("Finished Demo Cycle");
        runDemo = false;
    }
    generateColorString();
    delay(4000);
    Particle.publish("Colors", cs);
}

void generateColorString(){
    cs = "";
    for(int i = 0; i < PIXEL_COUNT; i++){
        cs +=  StateColor[i];
    }
}

int demo(String str){
    runDemo = true;
    return 0;
}

int colorString(String states){
    int statesColored = 0;
    for(int i = 0; i < PIXEL_COUNT; i++){
        switch(states[i]){
            case 'R':
                strip.setPixelColor(i, 0x5f, 0, 0);
                StateColor[i] = 'R';
                break;
            
            case 'G':
                strip.setPixelColor(i, 0, 0x5f, 0);
                StateColor[i] = 'G';
                break;
                
            case 'B':
                strip.setPixelColor(i, 0, 0, 0x5f);
                StateColor[i] = 'B';
                break;
                
            default:
                strip.setPixelColor(i, 0, 0, 0);
                StateColor[i] = ' ';
        }
        statesColored++;
    }
    strip.show();
    return statesColored;
}

void colorAll(uint32_t c, uint8_t wait) {
  uint16_t i;

  for(i=0; i<strip.numPixels(); i++) {
    strip.setPixelColor(i, c);
  }
  strip.show();
  delay(wait);
}
