import { Bill, Dish } from "../interfaces";

export const printDishesList = (menu: Dish[]) => {
    console.log(`Our menu:\n${menu.map(dish => `name: ${dish.name}, price: ${dish.price}\n`)}`);
}

export const printBill = (bill: Bill) => {
    console.log(`Bill was successfully created:\nYour bill ${bill.total_price_no_tax}\nRecommended tips are ${bill.recommended_tips}`);
}