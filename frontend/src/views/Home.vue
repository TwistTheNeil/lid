<template>
  <div class="home">
    {{ devices }}
  </div>
</template>

<script>
export default {
  name: "Home",
  data() {
    return {
      devices: null,
    };
  },
  methods: {
    async getStorageDevices() {
      const response = await fetch(`/api/v1/devices`);
      const responsejson = await response.json();
      this.devices = responsejson;
    },
  },
  async created() {
    this.getStorageDevices();
  },
};
</script>
