<template>
  <div class="home">
    <SearchbarInput />
    <FileList :headers="headers" />
  </div>
</template>

<script>
import SearchbarInput from "@/components/SearchbarInput.vue";
import FileList from "@/components/FileList.vue";
import { useFileStore } from "../store/fileStore";
import { useStorageDeviceStore } from "../store/storageDeviceStore";
import { onMounted } from "vue";

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

    const getStorageDevices = async () => {
      const response = await fetch(`/api/v1/devices`);
      const responsejson = await response.json();
      storageDevicesStore.$reset;
      storageDevicesStore.devices = responsejson;
    };
    const getFiles = async () => {
      const response = await fetch(`/api/v1/files`);
      const responsejson = await response.json();
      fileStore.$reset;
      fileStore.files = responsejson;
    };

    onMounted(async () => {
      await getStorageDevices();
      await getFiles();
    });

    return {
      headers,
    };
  },
};
</script>
