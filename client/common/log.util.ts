import { Bill, Dish } from "../interfaces";

export const printDishesList = (menu: Dish[]) => {
    console.log(`Our menu:\n${menu.map(dish => `name: ${dish.name}, price: ${dish.price}\n`)}`);
}

export const printBill = (bill: Bill) => {
    console.log(`Bill was successfully created: ${bill}`);
}