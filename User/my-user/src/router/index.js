// src/router/user-router.js
import { createRouter, createWebHistory } from 'vue-router';
import Home from '../views/Home.vue';
import Cart from '../views/Cart.vue';
import OrderHistory from '../views/OrderHistory.vue';
import Login from '../views/Login.vue';
import Signup from '../views/Signup.vue';
import Checkout from '../views/Checkout.vue';

const routes = [
  { path: '/', name: 'Home', component: Home },
  { path: '/cart', name: 'Cart', component: Cart },
  { path: '/order-history', name: 'OrderHistory', component: OrderHistory },
  { path: '/login', name: 'Login', component: Login },
  { path: '/signup', name: 'Signup', component: Signup },
  { path: '/checkout', name: 'Checkout', component: Checkout}
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
