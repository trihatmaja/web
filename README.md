#README
This is web interface for home automation using raspberry pi and golang project.
Feature:
 1. Turn On and Off.
 2. Dimming 220v light using slider.
 3. Motor speed control.
 4. API for turn on and off.
 5. Feedback for detecting that the device is actually turned on or off.
 6. Sense the current to calculate the estimated power and cost.
 
The feature will be added for the future:
 1. Chart.
 2. Schedule turn on and off.
 3. Sense the voltage value for better calculation of power and cost.
 4. Send voltage and current data to HADOOP for better historical data 
   and analitics using API or Socket programming.

#How to run
Several items that being use in this project, make sure you have install this before 
running:
 1. go language
 2. gorilla mux (http://www.gorillatoolkit.org/pkg/mux)
 3. mrmorphic hwio (http://github.com/mrmorphic/hwio)
 4. pi blaster (https://github.com/sarfata/pi-blaster)

to start the web:
 go run web.go

#Testing
Still no unit testing for this project.

#Electrical Schematic
See Circuit Folder

Component:
 1. Triac BT139
 2. Triac driver MOC3021
 3. Zero Crossing H11A1 or H11AA1
 4. Resistor 1k ohm 1 watt (the value depend on the current but minimum 1 watt)
 5. Fuse 3 amps/ 4 amps
 6. Lamp 5 watt for test
 7. Bridge Rectifier KBJ408G (if using H11A1)

#Beware
Careful with the electrical component, make sure it all safe!!! You playing with 110v/220v voltage line..
Careful..
Careful..

#Contribute
Contact me
