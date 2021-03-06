import { createPlugin } from '@backstage/core';
import WelcomePage from './components/WelcomePage';
import SignIn from './components/SignIn'
import Insurance from './components/Insurance';
import TableInsurance from './components/TableInsurance';
import ResultInsurance from './components/ResultInsurance';

import Product from './components/Product';
import Officerlongin from './components/Officerlongin';

import Payment from './components/Payment';
import Inquiry from './components/Inquiry'
import Inquirysearch from './components/Inquirysearch'
import Recordinsurance from './components/Recordinsurance'
import Payback from './components/Payback';
import RecordinsuranceSearch from './components/RecordinsuranceSearch';
import PaymentSearch from './components/PaymentSearch';
import PaybackSearch from './components/PaybackSearch';

import ProductSearch from './components/ProductSearch';
export const plugin = createPlugin({
  id: 'welcome',
  register({ router }) {
    router.registerRoute('/', WelcomePage);
    router.registerRoute('/signin', SignIn);
    router.registerRoute('/payment', Payment);
    router.registerRoute('/insurance', Insurance);
    router.registerRoute('/tableinsurance', TableInsurance);
    router.registerRoute('/resultinsurance', ResultInsurance);
    router.registerRoute('/product', Product);

    router.registerRoute('/officerlongin', Officerlongin);

    router.registerRoute('/inquiry', Inquiry);
    router.registerRoute('/inquirysearch', Inquirysearch);
    router.registerRoute('/payback', Payback);
    router.registerRoute('/recordinsurance', Recordinsurance);
    router.registerRoute('/recordinsurancesearch', RecordinsuranceSearch);
    router.registerRoute('/paymentsearch', PaymentSearch);
    router.registerRoute('/productsearch', ProductSearch);
    router.registerRoute('/paybacksearch', PaybackSearch);

  },
});
