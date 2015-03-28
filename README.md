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
Several items that being use in this, make sure you have install this before 
running:
 1. go language
 2. gorilla mux (http://www.gorillatoolkit.org/pkg/mux)
 3. mrmorphic hwio (http://github.com/mrmorphic/hwio)
 4. pi blaster (https://github.com/sarfata/pi-blaster)

to start the web:
 go run web.go

#Testing
Still no unit testing for this project.

#Electrical Circuit
Still in progress for drawing

#Contribute
Contact me
