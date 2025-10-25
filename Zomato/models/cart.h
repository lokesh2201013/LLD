#define CART_H
#include<bits/stdc++.h>

#include "MenuItem.h"
#include "Restaurant.h"

using namespace std;
class Cart{
    private:
    Restaurant * restaurant;
    vetcor<MenuItem> items;

    public:
    Cart(){
        restaurant=nullptr;
    }

    void addItem(const MenuItem&item){
        if(!restaurant){
            cout<<"Set a restaurant"<<"\n";
            return ;
        }
        items.push_back(item);
    }

    double gettotalCost(){
        double sum=0;
        for(const auto&it:items){
            sum+=it.getPrice();
        }
        return sum;
    }

    bool isEmpty(){
        return (!restaurant||items.empty());
    }
    void clear(){
        items.clear();
        restaurant=nullptr;
    }
    void setRestaurant(Restuarant*r){
        restaurant=r;
    }
     Restaurant* getRestaurant() const {
        return restaurant;
    }

    const vector<MenuItem>& getItems() const {
        return items;
    }
};