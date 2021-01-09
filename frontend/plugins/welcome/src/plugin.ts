import { createPlugin } from '@backstage/core';
import WelcomePage from './components/WelcomePage';
import SignIn from './components/SignIn'
<<<<<<< HEAD
import Insurance from './components/Insurance';
=======
import Payment from './components/Payment'

>>>>>>> 3ca1d64 (ทำหน้า UI ของระบบชำระเบี้ยประกัน - fix #25)

export const plugin = createPlugin({
  id: 'welcome',
  register({ router }) {
    router.registerRoute('/', WelcomePage);
    router.registerRoute('/signin', SignIn);
<<<<<<< HEAD
    router.registerRoute('/insurance', Insurance);
=======
    router.registerRoute('/payment', Payment);
>>>>>>> 3ca1d64 (ทำหน้า UI ของระบบชำระเบี้ยประกัน - fix #25)
  },
});
