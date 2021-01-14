import { createPlugin } from '@backstage/core';
import WelcomePage from './components/WelcomePage';
import SignIn from './components/SignIn'
import Insurance from './components/Insurance';
import TableInsurance from './components/TableInsurance';
import Product from './components/Product';
import Officerlogin from './components/Officerlogin';



export const plugin = createPlugin({
  id: 'welcome',
  register({ router }) {
    router.registerRoute('/', WelcomePage);
    router.registerRoute('/signin', SignIn);
    router.registerRoute('/insurance', Insurance);
    router.registerRoute('/tableinsurance', TableInsurance);
    router.registerRoute('/product', Product);
    router.registerRoute('/officerlogin', Officerlogin);
  },
});
