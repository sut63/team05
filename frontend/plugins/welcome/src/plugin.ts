import { createPlugin } from '@backstage/core';
import WelcomePage from './components/WelcomePage';
import SignIn from './components/SignIn'
<<<<<<< HEAD
import Insurance from './components/Insurance';
import TableInsurance from './components/TableInsurance';
=======
import Product from './components/Product'
>>>>>>> 8b87335 (สร้างหน้า UI ของระบบบันทึกข้อมูลผลิตภัณฑ์ - close #28)


export const plugin = createPlugin({
  id: 'welcome',
  register({ router }) {
    router.registerRoute('/', WelcomePage);
    router.registerRoute('/signin', SignIn);
<<<<<<< HEAD
    router.registerRoute('/insurance', Insurance);
    router.registerRoute('/tableinsurance', TableInsurance);

=======
    router.registerRoute('/product', Product);
>>>>>>> 8b87335 (สร้างหน้า UI ของระบบบันทึกข้อมูลผลิตภัณฑ์ - close #28)
  },
});
