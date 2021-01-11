import { createPlugin } from '@backstage/core';
import WelcomePage from './components/WelcomePage';
import SignIn from './components/SignIn'
<<<<<<< HEAD
import Insurance from './components/Insurance';
=======
>>>>>>> 868d6b0 (แก้หน้า UI สำหรับเก็บข้อมูลผลิตภัณฑ์ - close #92)
import TableInsurance from './components/TableInsurance';
import Product from './components/Product'


export const plugin = createPlugin({
  id: 'welcome',
  register({ router }) {
    router.registerRoute('/', WelcomePage);
    router.registerRoute('/signin', SignIn);
<<<<<<< HEAD
    router.registerRoute('/insurance', Insurance);
    router.registerRoute('/tableinsurance', TableInsurance);

=======
    router.registerRoute('/tableinsurance', TableInsurance);
>>>>>>> 868d6b0 (แก้หน้า UI สำหรับเก็บข้อมูลผลิตภัณฑ์ - close #92)
    router.registerRoute('/product', Product);
  },
});
