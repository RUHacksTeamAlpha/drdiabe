import Vue from "vue";
import VueRouter from "vue-router";

Vue.use(VueRouter);

const routerOptions = [];

const routes = routerOptions.map((route) => {
  return {
    ...route,
    component: () => import(`../views/${route.component}.vue`),
  };
});

const router = new VueRouter({
  mode: "history",
  routes,
});

export default router;
