<template>
  <Navbar />
  <div class="container">
    <router-view />
  </div>
</template>

<script>
import { onMounted } from "vue";

import { useFileStore } from "@/store/fileStore";
import { useStorageDeviceStore } from "@/store/storageDeviceStore";
import Navbar from "@/components/Navbar";

export default {
  name: "Home",
  components: {
    Navbar,
  },
  setup() {
    const fileStore = useFileStore();
    const storageDevicesStore = useStorageDeviceStore();

    onMounted(async () => {
      await storageDevicesStore.refresh();
      await fileStore.refresh();
    });
  },
};
</script>
