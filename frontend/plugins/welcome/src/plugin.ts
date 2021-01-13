import { createPlugin } from '@backstage/core';
import WelcomePage from './components/WelcomePage';
import SignIn from './components/SignIn'
import Insurance from './components/Insurance';
import TableInsurance from './components/TableInsurance';

import Product from './components/Product';
import Officerlongin from './components/Officerlongin';

import Payment from './components/Payment';
import Product from './components/Product'
import Inquiry from './components/Inquiry'

export const plugin = createPlugin({
  id: 'welcome',
  register({ router }) {
    router.registerRoute('/', WelcomePage);
    router.registerRoute('/signin', SignIn);
    router.registerRoute('/payment', Payment);
    router.registerRoute('/insurance', Insurance);
    router.registerRoute('/tableinsurance', TableInsurance);
    router.registerRoute('/product', Product);

    router.registerRoute('/officerlongin', Officerlongin);

    router.registerRoute('/inquiry', Inquiry);
  },
});
