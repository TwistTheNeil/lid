<template>
  <div class="home">
    <SearchbarInput />
    <FileList :headers="headers" :items="files" />
  </div>
</template>

<script>
import SearchbarInput from "@/components/SearchbarInput.vue";
import FileList from "@/components/FileList.vue";

export default {
  name: "Home",
  components: {
    SearchbarInput,
    FileList,
  },
  data() {
    return {
      storageDevices: null,
      files: null,
      headers: ["hash", "name"],
    };
  },
  methods: {
    async getStorageDevices() {
      const response = await fetch(`/api/v1/devices`);
      const responsejson = await response.json();
      this.storageDevices = responsejson;
    },
    async getFiles() {
      const response = await fetch(`/api/v1/files`);
      const responsejson = await response.json();
      this.files = responsejson;
    },
  },
  async created() {
    await this.getStorageDevices();
    await this.getFiles();
  },
};
</script>
