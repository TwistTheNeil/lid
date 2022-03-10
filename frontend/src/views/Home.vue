<template>
  <div class="home">
    <SearchbarInput />
    <FileList :headers="headers" :items="files" />
  </div>
</template>

<script>
import SearchbarInput from "@/components/SearchbarInput.vue";
import FileList from "@/components/FileList.vue";
import { ref, onMounted } from "vue";

export default {
  name: "Home",
  components: {
    SearchbarInput,
    FileList,
  },
  setup() {
    const storageDevices = ref([]);
    const files = ref([]);
    const headers = ["hash", "name"];

    const getStorageDevices = async () => {
      const response = await fetch(`/api/v1/devices`);
      const responsejson = await response.json();
      storageDevices.value = responsejson;
    };
    const getFiles = async () => {
      const response = await fetch(`/api/v1/files`);
      const responsejson = await response.json();
      files.value = responsejson;
    };

    onMounted(async () => {
      await getStorageDevices();
      await getFiles();
    });

    return {
      storageDevices,
      files,
      headers,

      getStorageDevices,
      getFiles,
    };
  },
};
</script>
