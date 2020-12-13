import env from './config/env';

import { Client } from './dishes/client';
import { printBill, printDishesList } from './common/log.util';

const client = new Client(env.url);

(async () => {
    try {
        const dishesList = await client.getMenuList();
        printDishesList(dishesList);
    } catch (err) {
        console.log(`Error during request to get dishes list: ${err.message || err}`);
    }

    try {
        const bill = await client.createOrder({
            dish_ids: [1, 2, 3],
            table_id: 1
        });
        printBill(bill);
    } catch (err) {
        console.log(`Error during request to create order: ${err.message || err}`);
    }
})();
