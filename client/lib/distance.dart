import 'dart:async';

import 'package:flutter/material.dart';
import 'package:geolocator/geolocator.dart';
import 'package:http/http.dart' as http;

class Distance extends StatefulWidget {
  const Distance({Key? key}) : super(key: key);

  @override
  _DistanceState createState() => _DistanceState();
}

class _DistanceState extends State<Distance> with SingleTickerProviderStateMixin {
  late AnimationController _animationController;
  int limit=100;
   int curDist=1000;
    @override
    void initState() {
      _animationController =
          new AnimationController(vsync: this, duration: Duration(seconds: 1));
      _animationController.repeat(reverse: true);
      super.initState();
    }
 
  void calcDist(startLat, startLong, endLat, endLong) async{
   
      var client = http.Client();
    var url = Uri.parse('http://localhost:4000/trigger');
    double distanceInMeters = Geolocator.distanceBetween(startLat, startLong, endLat, endLong);
    print(distanceInMeters);
    if(distanceInMeters<limit){
      var response = await client.get(url);
       print(response.body);
    }
  }
  void location() async{
    LocationPermission permission;
   permission = await Geolocator.requestPermission();
    final LocationSettings locationSettings = LocationSettings(
        accuracy: LocationAccuracy.high,
        distanceFilter: 0,
);
StreamSubscription<Position> positionStream = Geolocator.getPositionStream(locationSettings: locationSettings).listen(
    (Position? position) {
        if(position!=null){
          calcDist(10.7634132, 78.8124742, position.latitude, position.longitude);
        }
        print(position == null ? 'Unknown' : '${position.latitude.toString()}, ${position.longitude.toString()}');
    });
  }
  void triggerServer(){

  }

  @override
  Widget build(BuildContext context) {
    location();
    return Scaffold(
      body: Container(
        child:  Center(  
      child: Container(  
        width: 300,  
        height: 165,  
        padding: new EdgeInsets.all(10.0),  
        child: Card(  
          shape: RoundedRectangleBorder(  
            borderRadius: BorderRadius.circular(15.0),  
          ),  
          color: curDist<limit? Colors.green:Colors.red,  
          elevation: 10,  
          child: Column(  
            mainAxisSize: MainAxisSize.min,  
            children: <Widget>[
              FadeTransition(opacity: _animationController,
              child: ListTile(  
                contentPadding: EdgeInsets.all(20.0),
                title: Text(  
                  "Distance to destination",  
                  style: TextStyle(fontSize: 22.0) ,
                  
                ), 
                subtitle:
                Padding(padding: EdgeInsets.all(20.0),
                child: Text( 
                  curDist.toString(),  
                  textAlign: TextAlign.center,
                  style: TextStyle(fontSize: 20.0)  
                ), 
                )
                  
              ), 
              )
                
                
            ],  
          ),  
        ),  
      )  
    )  ,
      ),
    );
  }
}