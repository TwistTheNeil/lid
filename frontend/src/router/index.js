import { createRouter, createWebHashHistory } from "vue-router";
import Home from "@/views/Home.vue";
import Devices from "@/views/Devices.vue";

const routes = [
  {
    path: "/",
    name: "Home",
    component: Home,
  },
  {
    path: "/devices",
    name: "Devices",
    component: Devices,
  },
];

const router = createRouter({
  history: createWebHashHistory(process.env.BASE_URL),
  routes,
});

export default router;
