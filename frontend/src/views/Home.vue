<template>
  <div class="home">
    <SearchbarInput />
    <FileList :headers="headers" />
  </div>
</template>

<script>
import { onMounted } from "vue";
import SearchbarInput from "@/components/SearchbarInput.vue";
import FileList from "@/components/FileList.vue";

import { useFileStore } from "../store/fileStore";
import { useStorageDeviceStore } from "../store/storageDeviceStore";

export default {
  name: "Home",
  components: {
    SearchbarInput,
    FileList,
  },
  setup() {
    const headers = ["hash", "name"];
    const fileStore = useFileStore();
    const storageDevicesStore = useStorageDeviceStore();

    onMounted(async () => {
      await storageDevicesStore.refresh();
      await fileStore.refresh();
    });

    return {
      headers,
    };
  },
};
</script>
