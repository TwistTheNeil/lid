<template>
  <router-view />
</template>

<script>
import { onMounted } from "vue";

import { useFileStore } from "@/store/fileStore";
import { useStorageDeviceStore } from "@/store/storageDeviceStore";

export default {
  name: "Home",
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
