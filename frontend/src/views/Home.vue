<template>
  <div class="home">
    {{ storageDevices }}
    {{ files }}
  </div>
</template>

<script>
export default {
  name: "Home",
  data() {
    return {
      storageDevices: null,
      files: null,
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
