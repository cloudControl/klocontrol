## kloControl

http://www.hackster.io/hassaku/wifi-toilet-sensor-for-vacancy-and-light
When working at an office with around 15 bright minded professionals, but having only one restroom to share, overlapping needs are a regularly
occurring inconvenience. To overcome this,we wanted to connect our restroom to the office network, blessing our team with the power to find out
online whether or not the throne is free.

We used the Core Wifi- connected microcontroller by Spark.io combined with a photosensor to measure if the lights are on.

Build binary
~~~
go build .
~~~

Release
~~~ bash
 cd $GOPATH/src/github.com/cloudControl/klocontrol
 cctrlapp klocontrol create custom --buildpack https://github.com/cloudControl/buildpack-go
 cctrlapp klocontrol/default push --ship
~~~

