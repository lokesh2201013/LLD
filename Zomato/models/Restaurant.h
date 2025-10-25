#define RESTAURANT_H
#include <bits/stdc++.h>
#include "MenuItem.h"
using namespace std;

class Restaurant{
   
    private:
   //tokeep count of restautent like a uuid
    static int nextRestaurantId:
    int restaurantId:
    string name;
    string location;
    vector<MenuItem>menu;

    public:
    Restaurant(string name,string location){
        this->name=name;
        this->location=location;
        this->restaurantId=++nextRestaurantId;
    }
    ~Restaurant(){
        menu.clear();
    }

     string getName()const{
        return name;
     }

     void setName(string n){
        name=n;
     }
     string getLocation()const{
        return location;
     }

     void setLocation(string loc){
        location=loc;
     }  
     void addMenuItem(MenuItem item){
        menu.push_back(item);
     }
     vector<MenuItem>&getMenu()const{
        return menu;
     }
     
}
int Restauant::nextRestaurantId=0;