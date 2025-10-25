#include<bits/stdc++.h>
using namespace std;

#include "MenuItem.h"
#include "Restaurant.h"
#include "User.h"

class Order{
    public:
    static int nextOrderId;
    User* user;
    Restaurant *restaurant;
    vector<MenuItem>items;
    PaymentStrategy* paymentStrategy;
    double total;
    string scheduled;

    Order(){
        user=nullptr;
        restaurant=nullptr;
        paymentStartegy=nullptr;
        total=0.0;
        scheduled="";
        orderId=++netOrderId;
    }
    virtual ~Order(){
        delete paymentStrategy;
    }
    
    bool processPayement(){
        if(payementStrategy){
            payementStrategy->pay(total);
            return true;
        }
        else{
            cout<<"Choose a mode first"<<"\n";
            return false;
        }
    }
    virtual string getType()const=0;
    int getOrderId()const{
        return orderId;
    }

    void setUser(User*u){
        user=u;
    }

    User* getUser(){
        return user;
    }

    void setResatuarant(Restaurant*r){
        restaurant=r;
    }

    Restaurant* getRestuarant(){
        return restaurant;
    }
    void setItems(vector<MenuItems>&its){
        items=its;
        total=0;
        for(auto &i:items){
            total+=i.getPrice();
        }
    }

    const vector<MenuItem>&getItems(){
        return items;
    }

    void setPaymentStrategy(PaymentStrategy*p){
        paymentStrategy=p;
    }

    void setScheduled(string&s){
        scheduled=s;
    }
    
    string getScheduled(){
        return scheduled;
    }
    double getTotal(){
        return total;
    }
    void setTotal(int total){
        this->total=total;
    }

};
int Order::nextOrderId=0;