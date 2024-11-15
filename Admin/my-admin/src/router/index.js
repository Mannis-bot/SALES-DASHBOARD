import { createRouter, createWebHistory } from "vue-router";
import Dashboard from "../views/Dashboard.vue";

import OrdersManagement from "../views/OrdersManagement.vue";

const routes = [
  {
    path: "/",
    name: "Dashboard",
    component: Dashboard,
  },

  {
    path: "/orders",
    name: "OrdersManagement",
    component: OrdersManagement,
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
