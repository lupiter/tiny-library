import { createRouter, createWebHistory, RouteRecordRaw } from "vue-router";
import Home from "../views/Home.vue";

const routes: Array<RouteRecordRaw> = [
  {
    path: "/",
    name: "Home",
    component: Home
  },
  {
    path: "/about",
    name: "About",
    component: () =>
      import(/* webpackChunkName: "about" */ "../views/About.vue")
  },
  {
    path: "/books",
    name: "Catalogue",
    component: () =>
      import(/* webpackChunkName: "books" */ "../views/Books.vue")
  },
  {
    path: "/people",
    name: "Patrons",
    component: () =>
      import(/* webpackChunkName: "patrons" */ "../views/Patrons.vue")
  },
  {
    path: "/loans",
    name: "Loans",
    component: () =>
      import(/* webpackChunkName: "loans" */ "../views/Loans.vue")
  }
];

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
});

export default router;
