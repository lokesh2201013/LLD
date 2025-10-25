#define DELIVERY_H
#include<bits/stdc++.h>
#include "Order.h"
using namespace std;

class DeliveryOrder:public Order{
    private:
    string userAddress;

    public:
    DeliveryOrder(){
        userAddress="";
    }
    string getType()const override{
        return "Delivery";
    }

    void setUserAddress(const string&addr){
        userAddress=addr;
    }

    string getuserAddress()const{
        return userAddress;
    }
}