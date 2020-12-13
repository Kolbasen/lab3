import { HttpCllient } from '../common/http.util';

import { Dish, Order, Bill } from '../interfaces';


export class Client {
    private httpClient: HttpCllient;

    constructor(url: string) {
        this.httpClient = new HttpCllient(url);
    }

    public getMenuList() {
        return this.httpClient.get<Dish[]>('/dishes/list')
    }

    public createOrder(order: Order) {
        return this.httpClient.post<Bill, Order>('/dishes/order/create', order)
    }
}