<template>
  <Navbar />
  <div class="container">
    <router-view />
  </div>
</template>

<script>
import { onBeforeMount } from "vue";

import { useFileStore } from "@/store/fileStore";
import { useStorageDeviceStore } from "@/store/storageDeviceStore";
import Navbar from "@/components/Navbar.vue";

export default {
  name: "Home",
  components: {
    Navbar,
  },
  setup() {
    const fileStore = useFileStore();
    const storageDevicesStore = useStorageDeviceStore();

    onBeforeMount(async () => {
      await storageDevicesStore.refresh();
      await fileStore.refresh();
    });
  },
};
</script>
