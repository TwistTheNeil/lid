import { createRouter, createWebHashHistory } from "vue-router";
import Home from "@/views/Home.vue";
import Devices from "@/views/Devices.vue";
import DeviceFileList from "@/views/DeviceFileList";
import { useSearchStore } from "@/store/searchStore";

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
  {
    path: "/device/:uuid",
    name: "DeviceFileList",
    component: DeviceFileList,
  },
];

const router = createRouter({
  history: createWebHashHistory(process.env.BASE_URL),
  routes,
});

router.beforeEach(() => {
  const searchStore = useSearchStore();
  searchStore.search = "";
});

export default router;
