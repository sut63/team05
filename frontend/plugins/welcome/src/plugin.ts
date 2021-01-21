import { createPlugin } from '@backstage/core';
import WelcomePage from './components/WelcomePage';
import SignIn from './components/SignIn'
import Insurance from './components/Insurance';
import Payment from './components/Payment'


export const plugin = createPlugin({
  id: 'welcome',
  register({ router }) {
    router.registerRoute('/signin', SignIn);
    router.registerRoute('/welcompage', WelcomePage);
    router.registerRoute('/insurance', Insurance);
    router.registerRoute('/payment', Payment);
  },
});
